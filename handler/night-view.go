package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wasuwa/night-view-api/usecase"
)

// NightViewHandler 夜景のハンドラ
type NightViewHandler interface {
	FetchNightViewByID(ctx echo.Context, id string) error
}

// NightViewHandlerImpl NightViewHandlerの実装
type NightViewHandlerImpl struct {
	nightViewUsecase usecase.NightViewUsecase
}

// NewNightViewHandler NightViewHandlerの実装を初期化する
func NewNightViewHandler(us usecase.NightViewUsecase) NightViewHandler {
	return &NightViewHandlerImpl{nightViewUsecase: us}
}

// FetchNightViewByID IDに紐づく夜景を取得する
func (h *NightViewHandlerImpl) FetchNightViewByID(ctx echo.Context, id string) error {
	nightView, err := h.nightViewUsecase.FetchNightViewByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, nightView)
}
