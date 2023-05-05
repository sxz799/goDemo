package model

type ErrMsg struct {
	Msg string `json:"msg"`
}

type Response struct {
	Success bool     `json:"success"`
	Msg     string   `json:"msg"`
	ErrMsgs []ErrMsg `json:"errMsgs"`
}
