package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

const (
	accessPointTweet  = "https://api.twitter.com/2/tweets"
	accessPointMedia  = "https://upload.twitter.com/1.1/media/upload.json"
	accessPointSearch = "https://api.twitter.com/1.1/users/search.json"
	MaxImgInTweet     = 4
	MaxImgDataSize    = 1 << 20
)

func checkTwitterRights(twitterUid string, r *http.Request) (*database.TwUserAccessToken, error) {
	if len(twitterUid) == 0 {
		util.LogInst().Warn().Msg("no twitter id for ninja user:" + twitterUid)
		return nil, fmt.Errorf("bind twitter first")
	}

	var ut, err = getAccessTokenFromSession(r)
	if err == nil {
		return ut, nil
	}
	ut, err = database.DbInst().GetTwAccessToken(twitterUid)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("access token not in db")
		return nil, err
	}
	return ut, nil
}

func twitterApiPost(url string, token *oauth1.Token,
	reqBody io.Reader, contentType string, response any) error {

	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		util.LogInst().Err(err).Msg("create http post request failed")
		return err
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := httpClient.Do(req)
	if err != nil {
		util.LogInst().Err(err).Msg(" http do  failed")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bts, _ := io.ReadAll(resp.Body)
		util.LogInst().Warn().Int("status", resp.StatusCode).Msg("post tweet failed:" + string(bts))
		return fmt.Errorf("status:%s err:%s", resp.Status, string(bts))
	}
	if response == nil {
		return nil
	}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		util.LogInst().Err(err).Msg("decode response failed")
		return err
	}
	return nil
}

func searchTwitterUsr(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var ut, err = checkTwitterRights(nu.TwID, r)
	if err != nil {
		util.LogInst().Err(err).Msg("load access token failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var keyWords = r.URL.Query().Get("q")

	var token = ut.GetToken()
	var users = make([]database.TWUserInfo, 0)
	users, err = queryTwitterByName(token, keyWords)
	if err != nil {
		util.LogInst().Err(err).Msg("search twitter user failed")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(users)
	w.Write(bts)

	util.LogInst().Debug().Str("q", keyWords).Int("user-no", len(users)).Msg("search twitter user")
}
func resizeImage(img image.Image, targetWidth, targetHeight int) (image.Image, error) {
	srcBounds := img.Bounds()

	// 创建一个新的目标尺寸的空图片
	dstImg := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	// 使用高质量的缩放算法缩放图片
	draw.CatmullRom.Scale(dstImg, dstImg.Bounds(), img, srcBounds, draw.Over, nil)

	return dstImg, nil
}

// 处理base64编码的图片字符串，获取图片对象，并缩放图片
func processBase64Image(hash, base64Str string) (image.Image, error) {
	// 移除数据URI方案的前缀（如果存在）
	prefixIdx := strings.IndexByte(base64Str, ',') + 1
	base64Data := base64Str[prefixIdx:]
	prefix := base64Str[:prefixIdx]

	// 解码base64字符串
	imgData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, fmt.Errorf("decode image base64 failed: %v", err)
	}

	// 解码图片数据
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, fmt.Errorf("decode image failed: %v", err)
	}

	util.LogInst().Info().Int("img-len", len(base64Str)).Msg("raw image is fine")
	if len(base64Str) <= MaxImgDataSize {
		if len(hash) > 20 {
			err = database.DbInst().SaveRawImg(hash, base64Str)
		}
		return img, err
	}

	factor := float64(MaxImgDataSize) / float64(len(base64Str))
	srcBounds := img.Bounds()
	targetWidth := int(float64(srcBounds.Dx()) * factor)
	targetHeight := int(float64(srcBounds.Dy()) * factor)

	resizedImg, err := resizeImage(img, targetWidth, targetHeight)
	if err != nil {
		return nil, fmt.Errorf("resize image failed: %v", err)
	}
	var buf2 bytes.Buffer
	if err := jpeg.Encode(&buf2, resizedImg, nil); err != nil {
		return nil, fmt.Errorf("encode resized image failed: %v", err)
	}

	resizedBase64 := prefix + base64.StdEncoding.EncodeToString(buf2.Bytes())
	util.LogInst().Info().Int("img-len", len(resizedBase64)).Msg("resize image success")
	if len(hash) > 20 {
		err = database.DbInst().SaveRawImg(hash, resizedBase64)
	}
	return resizedImg, err
}

func prepareTweet(njTweet *database.NinjaTweet, ut *database.TwUserAccessToken) (*TweetRequest, error) {

	var token = ut.GetToken()

	var combinedTxt = njTweet.TxtWithSlogan
	mediaIDs := make([]string, 0)
	if len(njTweet.Images) > MaxImgInTweet {
		return nil, fmt.Errorf("images must be less than 5")
	}
	var finalHash []string
	var finalImages []string
	for i, base64Data := range njTweet.ImageRaw {
		hash := njTweet.ImageHash[i]
		if len(hash) > 20 {
			finalHash = append(finalHash, hash)
			finalImages = append(finalImages, njTweet.Images[i])
		}

		img, err := processBase64Image(hash, base64Data)
		if err != nil {
			util.LogInst().Err(err).Msg("parse tweet image to image object failed")
			return nil, err
		}
		mediaID, err := uploadMedia(token, img)
		if err != nil {
			util.LogInst().Err(err).Msg("upload media for tweet failed")
			return nil, err
		}
		mediaIDs = append(mediaIDs, mediaID)

	}
	njTweet.ImageHash = finalHash
	njTweet.Images = finalImages

	req := &TweetRequest{
		Text: combinedTxt,
	}
	if len(mediaIDs) > 0 {
		req.Media = &Media{
			MediaIDs: mediaIDs,
		}
	}
	return req, nil
}

func postTweets(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var ut, err = checkTwitterRights(nu.TwID, r)
	if err != nil {
		util.LogInst().Err(err).Msg("load access token failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	param := &SignDataByEth{}
	err = util.ReadRequest(r, param)
	if err != nil {
		util.LogInst().Err(err).Msg("Error parsing sign data ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	njTweet, err := param.ParseNinjaTweet()
	if err != nil {
		util.LogInst().Err(err).Msg("parse ninja tweet failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tweetBody, err := prepareTweet(njTweet, ut)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var tweetResponse TweetResponse
	bts, _ := json.Marshal(tweetBody)

	err = twitterApiPost(accessPointTweet, ut.GetToken(), bytes.NewBuffer(bts), "application/json", &tweetResponse)
	if err != nil {
		util.LogInst().Err(err).Msg(" posted tweet failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	njTweet.TweetId = tweetResponse.Data.ID
	err = database.DbInst().SaveTweet(njTweet)
	if err != nil {
		util.LogInst().Err(err).Msg("save posted tweet failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(njTweet)
	util.LogInst().Debug().Str("tweet-id", njTweet.TweetId).
		Int64("create_time", njTweet.CreateAt).
		Msg("Tweet posted successfully")
}

func uploadMedia(token *oauth1.Token, img image.Image) (string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	part, err := writer.CreateFormFile("media", "image.jpg")
	if err != nil {
		util.LogInst().Err(err).Msg("create form file failed")
		return "", err
	}

	err = jpeg.Encode(part, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		util.LogInst().Err(err).Msg("jpeg Encode failed")
		return "", err
	}

	writer.Close()
	var result map[string]interface{}
	err = twitterApiPost(accessPointMedia, token, &buffer, writer.FormDataContentType(), &result)
	if err != nil {
		util.LogInst().Err(err).Msg("twitterApiPost  failed")
		return "", err
	}

	mediaID, ok := result["media_id_string"].(string)
	if !ok {
		bts, _ := json.Marshal(result)
		util.LogInst().Warn().Msg("upload media failed:" + string(bts))
		return "", fmt.Errorf("error getting media ID")
	}

	return mediaID, nil
}

func queryTwitterByName(token *oauth1.Token, query string) ([]database.TWUserInfo, error) {
	httpClient := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret).Client(oauth1.NoContext, token)

	queryParams := url.Values{}
	queryParams.Add("q", query)
	//queryParams.Add("page", "3")
	//queryParams.Add("count", "5")

	requestURL := fmt.Sprintf("%s?%s", accessPointSearch, queryParams.Encode())

	resp, err := httpClient.Get(requestURL)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to query Twitter API")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		buf, _ := io.ReadAll(resp.Body)
		util.LogInst().Warn().Str("body", string(buf)).Int("status", resp.StatusCode).Msg("Twitter API request failed")
		return nil, fmt.Errorf("twitter API request failed with status code: %d", resp.StatusCode)
	}

	var users []database.TWUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		util.LogInst().Err(err).Msg("Failed to decode response from Twitter API")
		return nil, err
	}

	return users, nil
}

func shareVoteAction(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {

	vote := &database.TweetVoteAction{}
	var err = util.ReadRequest(r, vote)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing payment status param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ut, err := checkTwitterRights(nu.TwID, r)
	if err != nil {
		util.LogInst().Err(err).Msg("load access token failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req = &TweetRequest{
		Text: _globalCfg.GetNjVoteAd(vote.CreateTime, nu.EthAddr, vote.Slogan),
	}

	bts, _ := json.Marshal(req)
	var tweetResponse TweetResponse
	err = twitterApiPost(accessPointTweet, ut.GetToken(), bytes.NewBuffer(bts), "application/json", &tweetResponse)
	if err != nil {
		util.LogInst().Err(err).Msg(" posted tweet failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.LogInst().Debug().Str("web3-id", nu.EthAddr).
		Str("tweet-id", tweetResponse.Data.ID).
		Int64("create_time", vote.CreateTime).
		Msg("share vote tweet successfully")
}
