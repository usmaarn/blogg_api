package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Errors     any    `json:"errors,omitempty"`
}

func ErrorResponse(w http.ResponseWriter, status int, message string, erorrMap any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&Response{status, message, nil, erorrMap})
}

func SuccessResponse(w http.ResponseWriter, status int, message string, data any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&Response{status, message, data, nil})
}
