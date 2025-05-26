package response

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func ErrorResponse(message string, erorrMap any) *Response {
	return &Response{false, message, nil, erorrMap}
}

func SuccessResponse(message string, data any) *Response {
	return &Response{true, message, data, nil}
}
