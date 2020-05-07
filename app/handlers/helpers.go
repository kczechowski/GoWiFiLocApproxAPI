package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func respondWithJson(w http.ResponseWriter, payload interface{}) http.ResponseWriter {
	body, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, err)
		return w
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
	return w
}

type JsonError struct {
	ErrorType string `json:"type"`
	Message   string `json:"message"`
}

func respondWithError(w http.ResponseWriter, err error, status int) http.ResponseWriter {
	w.WriteHeader(status)

	m := make(map[string]JsonError)
	m["error"] = JsonError{
		ErrorType: reflect.TypeOf(err).String(),
		Message:   err.Error(),
	}
	w = respondWithJson(w, m)
	return w
}
