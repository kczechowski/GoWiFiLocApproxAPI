package container

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	Router        *mux.Router
	MongoClient   *mongo.Client
	MongoDatabase func() *mongo.Database
}
