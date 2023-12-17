package server

const (
	DefaultProjectID = "dessage"
)

type TwitterAPIResponse struct {
	Data TwitterBasicInfo `json:"data"`
}

type TwitterBasicInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	SignUpAt     uint64 `json:"sign_up_at"`
	RefreshToken string `json:"refresh_token"`
}
