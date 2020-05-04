package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"time"
)

func main() {

	mongodbUri := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongodbUri))

	exitIfError(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	exitIfError(err)
	err = client.Ping(ctx, nil)


	exitIfError(err)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "xsdWelcome to this life-changing API.\nIts the best API, its true, all other API's are fake.")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", r)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
