package logger

type LoggerDebug struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoggerError struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type Logger struct {
	LogMessage string      `json:"log_message"`
	Data       interface{} `json:"data"`
	Err        interface{} `json:"err"`
}
