package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilemarbarcelos/nfl-scores/parser"
)

func FindGames(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	games := parser.Parse(vars["week"])
	json.NewEncoder(w).Encode(games)
}
