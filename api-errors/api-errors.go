package apierrors

import "net/http"

// APIError APIエラー
type APIError struct {
	StatusCode int `json:"-"`
	Code       string
	Message    string
}

var (
	// InternalServerError 500エラー
	InternalServerError = &APIError{
		StatusCode: http.StatusInternalServerError,
		Code:       "A100",
		Message:    "予期せぬエラーが発生しました。もう一度お試しください。",
	}
)
