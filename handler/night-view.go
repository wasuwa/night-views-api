package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wasuwa/night-view-api/usecase"
)

type NightViewHandler interface {
	FetchNightViewByID(http.ResponseWriter, *http.Request)
}

type nightViewHandler struct {
	nightViewUsecase usecase.NightViewUsecase
}

// NewNightViewHandler NightViewHandlerの実装を初期化する
func NewNightViewHandler(us usecase.NightViewUsecase) NightViewHandler {
	return &nightViewHandler{nightViewUsecase: us}
}

// FetchNightViewByID IDに紐づく夜景を取得する
func (nh *nightViewHandler) FetchNightViewByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	nightView, err := nh.nightViewUsecase.FetchNightViewByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(nightView)
}
