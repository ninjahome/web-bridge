package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	accessPointTweet = "https://api.twitter.com/2/tweets"
	accessPointMedia = "https://upload.twitter.com/1.1/media/upload.json"
)

func checkTwitterRights(w http.ResponseWriter, r *http.Request) (*TwUserAccessToken, error) {
	var ninjaUsr, err = validateUsrRights(r)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return nil, err
	}

	var twitterUid = ninjaUsr.TwID
	if len(twitterUid) == 0 {
		util.LogInst().Warn().Msg("no twitter id for ninja user:" + ninjaUsr.EthAddr)
		return nil, fmt.Errorf("bind twitter first")
	}
	var ut, errToken = getAccessTokenFromSession(r)
	if errToken == nil {
		return ut, nil
	}
	ut, err = DbInst().GetTwAccessToken(twitterUid)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("access token not in db")
		return nil, err
	}
	return ut, nil
}

func twitterApiPost(token *oauth1.Token, reqBody any, contentType string, response any) error {

	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	bts, err := json.Marshal(reqBody)
	if err != nil {
		util.LogInst().Err(err).Msg("marshal request failed")
		return err
	}

	req, err := http.NewRequest("POST", accessPointTweet, bytes.NewBuffer(bts))
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
		util.LogInst().Warn().Int("status", resp.StatusCode).Msg("post tweet failed" + string(bts))
		return fmt.Errorf("http response status err:%s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		util.LogInst().Err(err).Msg("decode response failed")
		return err
	}
	return nil
}

func prepareTweet(njTweet *NinjaTweet, ut *TwUserAccessToken) (*TweetRequest, error) {
	var appendStr = _globalCfg.GetNjProtocolAd(njTweet.CreateAt)
	var combinedTxt = njTweet.Txt + appendStr
	if !util.IsOverTwitterLimit(combinedTxt) {
		return &TweetRequest{
			Text: combinedTxt,
		}, nil
	}

	var token = ut.GetToken()
	txtImg, err := util.ConvertLongTweetToImg(njTweet.Txt, _globalCfg.imgFont, _globalCfg.FontSize)
	if err != nil {
		util.LogInst().Err(err).Msg("convert txt to img failed:" + njTweet.String())
		return nil, err
	}

	mediaID, err := uploadMedia(token, txtImg)
	if err != nil {
		util.LogInst().Err(err).Msg("convert txt to img failed:" + njTweet.String())
		return nil, err
	}
	combinedTxt = util.TruncateString(njTweet.Txt, 80) + "..." + appendStr
	var req = &TweetRequest{
		Text: combinedTxt,
		Media: &Media{
			MediaIDs: []string{mediaID},
		},
	}
	return req, nil
}

func postTweets(w http.ResponseWriter, r *http.Request) {
	var ut, err = checkTwitterRights(w, r)
	if err != nil {
		util.LogInst().Err(err).Msg("load access token failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	err = twitterApiPost(ut.GetToken(), tweetBody, "application/json", &tweetResponse)
	if err != nil {
		util.LogInst().Err(err).Msg(" posted tweet failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	njTweet.TweetId = tweetResponse.Data.ID
	err = DbInst().SaveTweet(njTweet)
	if err != nil {
		util.LogInst().Err(err).Msg("save posted tweet failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweetResponse)
	util.LogInst().Debug().Msg("Tweet posted successfully")
}

func uploadMedia(token *oauth1.Token, img image.Image) (string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	part, err := writer.CreateFormFile("media", "image.jpg")
	if err != nil {
		return "", err
	}

	err = jpeg.Encode(part, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		return "", err
	}
	writer.Close()

	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	req, err := http.NewRequest("POST", accessPointMedia, &buffer)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bts, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.NewDecoder(bytes.NewBuffer(bts)).Decode(&result) // 使用 bts 构建新的 buffer，因为原 buffer 已被读取

	mediaID, ok := result["media_id_string"].(string)
	if !ok {
		bts, _ := json.Marshal(result)
		util.LogInst().Warn().Msg("upload media failed:" + string(bts))
		return "", fmt.Errorf("error getting media ID")
	}

	return mediaID, nil
}
