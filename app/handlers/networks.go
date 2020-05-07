package handlers

import (
	"fmt"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/container"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetNetworks(container *container.Container, w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	collection := container.MongoDatabase().Collection("networks")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var networks []bson.M

	if err = cursor.All(ctx, &networks); err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}
	if err := cursor.Err(); err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
		return
	}

	//if found records == 0 then initialize networks slice
	if networks == nil {
		networks = make([]bson.M, 0)
	}

	w = respondWithJson(w, networks)
}

func PostNetwork(container *container.Container, w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
