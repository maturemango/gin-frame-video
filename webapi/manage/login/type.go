package login

import "time"

type SysLoginCaptcha struct {
	Id       string   `json:"id"`
	B64s     string   `json:"base64image"`
	Answer   string   `json:"answer"`
}

type requestInfo struct {
	count          int
	fristRequest   time.Time
}