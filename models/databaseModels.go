package models

type Session struct {
	ClientCode   string `json:"client_code"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	SessionToken string `json:"session_token"`
	SessionKey   string `json:"session_key"`
}
