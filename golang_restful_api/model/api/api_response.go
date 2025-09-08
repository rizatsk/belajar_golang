package api

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ApiResponseError struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Error   interface{} `json:"error"`
}
