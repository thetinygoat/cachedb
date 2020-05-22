package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, r *http.Request, res Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}
