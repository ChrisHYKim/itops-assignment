package util

import (
	"encoding/json"
	"net/http"
)

// API Error 발생 시 출력
type APIError struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
}

func (e APIError) Error() string {
	return e.Message
}

func NewError(statusCode int, message string) error {
	return APIError{Message: message, Code: statusCode}
}
func SendErrorResponse(w http.ResponseWriter, err error) {
	apiErr, ok := err.(APIError)
	if !ok {
		apiErr = APIError{Message: "Internal Server Error", Code: http.StatusInternalServerError}
	}
	// HTTP Header json 형식으로 Frontend 에 전달
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(apiErr.Code)
	// 에러 메시지와 상태 코드 반환
	json.NewEncoder(w).Encode(apiErr)
}
