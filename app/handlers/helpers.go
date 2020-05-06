package handlers

import (
	"encoding/json"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, payload interface{}) http.ResponseWriter{
	body, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return w
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
	return w
}
