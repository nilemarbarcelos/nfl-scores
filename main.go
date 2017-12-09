package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nilemarbarcelos/nfl-scores/config"
	"github.com/nilemarbarcelos/nfl-scores/controller"
)

func main() {
	nfl := mux.NewRouter()
	nfl.Path("/nfl-scores/{season}/{week}").Methods(http.MethodGet).HandlerFunc(controller.FindGames)

	if err := http.ListenAndServe(":"+config.Get().Port, handlers.CORS()(nfl)); err != nil {
		log.Fatal(err)
	}
}
