package controller

import (
	"encoding/json"
	"net/http"
	"../util"
)

func FindGames(w http.ResponseWriter, r *http.Request) {
	games := util.Parse()
	json.NewEncoder(w).Encode(games)
}