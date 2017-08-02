package view

// ResponseCommon General response structure
type ResponseCommon struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
