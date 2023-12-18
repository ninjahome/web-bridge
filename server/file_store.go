package server

import "encoding/json"

const (
	DefaultProjectID = "dessage"
)

type TwitterAPIResponse struct {
	Data TWUserInfo `json:"data"`
}

type TWUserInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	SignUpAt     uint64 `json:"sign_up_at"`
	RefreshToken string `json:"refresh_token"`
}

func (t *TWUserInfo) String() string {
	bts, _ := json.Marshal(t)
	return string(bts)
}

func TWUsrInfoMust(str string) *TWUserInfo {
	t := &TWUserInfo{}
	err := json.Unmarshal([]byte(str), t)
	if err != nil {
		return t
	}
	return t
}
