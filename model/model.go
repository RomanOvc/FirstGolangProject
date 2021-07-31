package model

type Resp struct {
	Result int `json:"answer"`
}

type ErrorResp struct {
	Err string `json:"err"`
}
