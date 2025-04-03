package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wasuwa/night-view-api/api"
	"github.com/wasuwa/night-view-api/config"
	"github.com/wasuwa/night-view-api/handler"
	"github.com/wasuwa/night-view-api/infrastructure/database"
	"github.com/wasuwa/night-view-api/infrastructure/store"
	"github.com/wasuwa/night-view-api/middlewares"
	"github.com/wasuwa/night-view-api/usecase"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	db := database.NewDB()
	nightViewRepository := store.NewNightViewStore(db)
	nearestStationRepository := store.NewNearestStationStore(db)
	nightViewUsecase := usecase.NewNightViewUsecase(nightViewRepository, nearestStationRepository)
	nightViewHandler := handler.NewNightViewHandler(nightViewUsecase)

	server := api.NewServer(nightViewHandler)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(
		middleware.CORS(),
		middleware.Logger(),
		middleware.Recover(),
		middleware.CSRF(),
		middlewares.WithRequestID,
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
