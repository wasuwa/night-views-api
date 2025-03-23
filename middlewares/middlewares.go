package middlewares

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/wasuwa/night-view-api/logging"
)

// WithRequestID コンテキストにリクエストIDを含める
func WithRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID, err := uuid.NewRandom()
		if err != nil {
			return next(c)
		}
		req := c.Request()
		ctx := context.WithValue(req.Context(), logging.RequestIDKey, requestID.String())
		c.SetRequest(req.WithContext(ctx))
		return next(c)
	}
}
