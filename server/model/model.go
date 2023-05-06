package model

type ErrMsg struct {
	Msg string `json:"msg"`
}

type Response struct {
	Success bool     `json:"success"`
	Msg     string   `json:"msg"`
	ErrMsgs []ErrMsg `json:"errMsgs"`
}

type Dept struct {
	Name string
	Code string
	Mkt  string
}
type User struct {
	Name string
	Mkt  string
}
