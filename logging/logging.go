package logging

import (
	"context"
	"log/slog"
	"os"
)

func init() {
	logger := slog.New(&Handler{slog.NewJSONHandler(os.Stderr, nil)})
	slog.SetDefault(logger)
}

type contextKey string

const (
	// RequestIDKey リクエストIDのキー
	RequestIDKey contextKey = "request_id"
)

var keys = []contextKey{RequestIDKey}

// Handler ハンドラー
type Handler struct {
	slog.Handler
}

// Handle ハンドル
func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	for _, key := range keys {
		if v := ctx.Value(key); v != nil {
			r.AddAttrs(slog.Attr{Key: string(key), Value: slog.AnyValue(v)})
		}
	}
	return h.Handler.Handle(ctx, r)
}
