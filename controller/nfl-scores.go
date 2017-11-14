package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nilemarbarcelos/nfl-scores/parser"
)

func FindGames(w http.ResponseWriter, r *http.Request) {
	games := parser.Parse()
	json.NewEncoder(w).Encode(games)
}
