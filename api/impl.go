package api

import (
	"github.com/wasuwa/night-view-api/handler"
)

// Server ServerInterfaceの実装
type Server struct {
	handler.NightViewHandler
}

// NewServer サーバーを初期化する
func NewServer(nightViewHandler handler.NightViewHandler) Server {
	return Server{
		NightViewHandler: nightViewHandler,
	}
}
