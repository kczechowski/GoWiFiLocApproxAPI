package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type NetworkFilter struct {
	Mac string `json:"mac"`
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	DeviceId string `json:"device_id"`
}

func (networkFilter NetworkFilter) ToMongoFilter() bson.M {
	filter := bson.M{}

	v := reflect.ValueOf(networkFilter)
	for i := 0; i < v.NumField(); i++ {
		jsonTag := v.Type().Field(i).Tag.Get("json")
		fieldValue := v.Field(i)
		if !fieldValue.IsZero() {
			filter[jsonTag] = fieldValue.String()
		}
	}

	return filter
}