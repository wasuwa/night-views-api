package main

import (
	"net/http"

	"github.com/wasuwa/night-view-api/handler"
	"github.com/wasuwa/night-view-api/infrastructure/datastore"
	"github.com/wasuwa/night-view-api/usecase"
)


func main() {
	db := datastore.NewDB()
	nightViewRepository := datastore.NewNightViewStore(db)
	nightViewUsecase := usecase.NewNightViewUsecase(nightViewRepository)
	nightViewHandler := handler.NewNightViewHandler(nightViewUsecase)

	http.HandleFunc("/night-view", nightViewHandler.FetchNightViewByID)
	http.ListenAndServe(":8080", nil)
}
