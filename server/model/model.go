package model

type ErrInfo struct {
	Line     int    `json:"line"`
	ErrorMsg string `json:"errorMsg"`
	FixMsg   string `json:"fixMsg"`
}

type Response struct {
	Success  bool      `json:"success"`
	Msg      string    `json:"msg"`
	ErrInfos []ErrInfo `json:"errInfos"`
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
