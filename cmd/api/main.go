package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wasuwa/night-view-api/api"
	"github.com/wasuwa/night-view-api/handler"
	"github.com/wasuwa/night-view-api/infrastructure/datastore"
	"github.com/wasuwa/night-view-api/usecase"
)

func main() {
	db := datastore.NewDB()
	nightViewRepository := datastore.NewNightViewStore(db)
	nightViewUsecase := usecase.NewNightViewUsecase(nightViewRepository)
	nightViewHandler := handler.NewNightViewHandler(nightViewUsecase)

	server := api.NewServer(nightViewHandler)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(
		middleware.CORS(),
		middleware.Logger(),
		middleware.Recover(),
		middleware.CSRF(),
		middleware.RequestID(),
		middleware.Secure(),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Skipper:      middleware.DefaultSkipper,
			Timeout:      time.Minute,
			ErrorMessage: "timeout",
		}),
	)

	api.RegisterHandlers(e, server)

	log.Fatal(e.Start(":8080"))
}
