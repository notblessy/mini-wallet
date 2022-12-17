package utils

var (
	RespStatusSuccess = "success"
	RespStatusFail    = "fail"
	RespStatusError   = "error"
)

type DefaultResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
