package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilemarbarcelos/nfl-scores/controller"
)

func main() {
	nfl := mux.NewRouter()
	nfl.Path("/nfl-scores").Methods(http.MethodGet).HandlerFunc(controller.FindGames)

	if err := http.ListenAndServe(":3000", nfl); err != nil {
		log.Fatal(err)
	}
}
