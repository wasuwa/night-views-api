package apierrors

import "net/http"

// APIError APIエラー
type APIError struct {
	StatusCode int `json:"-"`
	Code       string
	Message    string
}

var (
	// InternalServerError サーバーエラー
	InternalServerError = &APIError{
		StatusCode: http.StatusInternalServerError,
		Code:       "A100",
		Message:    "予期せぬエラーが発生しました。もう一度お試しください。",
	}
	// NightViewNotFound 夜景情報が見つからない
	NightViewNotFound = &APIError{
		StatusCode: http.StatusNotFound,
		Code:       "N100",
		Message:    "夜景情報が見つかりません。",
	}
)
