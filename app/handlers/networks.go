package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/container"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"reflect"
)

func GetNetworks(container *container.Container, w http.ResponseWriter, r *http.Request) {

	networkFilter := models.Network{}

	if filters := r.URL.Query()["filter"]; len(filters) > 0 {
		var filterString = filters[0]

		err := json.Unmarshal([]byte(filterString), &networkFilter)
		if err != nil {
			respondWithError(w, err, http.StatusBadRequest)
			return
		}
	}

	filter := bson.M{}

	// foreach on parsed filter
	v := reflect.ValueOf(networkFilter)
	for i := 0; i < v.NumField(); i++ {
		jsonTag := v.Type().Field(i).Tag.Get("json")
		fieldValue := v.Field(i)
		if !fieldValue.IsZero() {
			filter[jsonTag] = fieldValue.String()
		}
	}

	fmt.Println(filter)

	ctx := r.Context()
	collection := container.MongoDatabase().Collection("networks")

	cursor, err := collection.Find(ctx, filter)
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

	keys := r.URL.Query()["d"]

	var deviceid string

	if len(keys) > 0 {
		deviceid = keys[0]
	}

	network := models.Network{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&network); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	collection := container.MongoDatabase().Collection("networks")

	res, err := collection.InsertOne(ctx, bson.M{
		"mac": network.Mac,
		"lat": network.Lat,
		"lon": network.Lon,
		"device_id": deviceid,
	})

	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
	}

	result := collection.FindOne(ctx, bson.M{
		"_id": res.InsertedID,
	})

	var addedNetwork bson.M

	if err = result.Decode(&addedNetwork); err != nil {
		respondWithError(w, err, http.StatusInternalServerError)
	}

	respondWithJson(w, addedNetwork)
}
