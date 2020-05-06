package handlers

import (
	"fmt"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/container"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func GetNetworks(container *container.Container, w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	collection := container.MongoDatabase().Collection("networks")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var networks []bson.M

	if err = cursor.All(ctx, &networks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
