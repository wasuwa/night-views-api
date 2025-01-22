package main

import (
	"log"

	"github.com/labstack/echo/v4"
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
	api.RegisterHandlers(e, server)

	log.Fatal(e.Start(":8080"))
}
