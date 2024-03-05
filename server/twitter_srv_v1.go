package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"unicode/utf8"
)

const (
	accessPointTweet  = "https://api.twitter.com/2/tweets"
	accessPointMedia  = "https://upload.twitter.com/1.1/media/upload.json"
	accessPointSearch = "https://api.twitter.com/1.1/users/search.json"
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

func prepareTweet(njTweet *database.NinjaTweet, ut *database.TwUserAccessToken) (*TweetRequest, error) {

	var appendStr = _globalCfg.GetNjProtocolAd(njTweet.CreateAt)
	var combinedTxt = njTweet.Txt + appendStr
	if !util.IsOverTwitterLimit(combinedTxt) {
		return &TweetRequest{
			Text: combinedTxt,
		}, nil
	}

	var token = ut.GetToken()
	var txtLen = len(njTweet.Txt)
	count := utf8.RuneCountInString(njTweet.Txt)
	util.LogInst().Debug().Int("txt-len", txtLen).Int("txt-count", count).Send()

	splitTxt := util.SplitIntoChunks(njTweet.Txt, _globalCfg.MaxTxtPerImg)
	if len(splitTxt) > 4 {
		return nil, fmt.Errorf("txt is too long")
	}
	mediaIDs := make([]string, 0)
	for _, content := range splitTxt {
		txtImg, err := util.ConvertLongTweetToImg(content, _globalCfg.imgFont, _globalCfg.FontSize)
		if err != nil {
			util.LogInst().Err(err).Msg("convert txt to img failed:" + njTweet.String())
			return nil, err
		}

		mediaID, err := uploadMedia(token, txtImg)
		if err != nil {
			util.LogInst().Err(err).Msg("convert txt to img failed:" + njTweet.String())
			return nil, err
		}
		mediaIDs = append(mediaIDs, mediaID)
	}

	combinedTxt = util.TruncateString(njTweet.Txt, appendStr)
	var req = &TweetRequest{
		Text: combinedTxt,
		Media: &Media{
			MediaIDs: mediaIDs,
		},
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

func queryTwitterByName() {

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
