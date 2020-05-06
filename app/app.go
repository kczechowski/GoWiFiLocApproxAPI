package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/container"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Container *container.Container
}

func (app *App) Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Container = &container.Container{}

	mongoClient, err := app.getMongo()

	if err != nil {
		log.Fatal(err.Error())
	}
	app.Container.Mongo = mongoClient

	app.Container.Router = mux.NewRouter()
	app.setRoutes()

}

func (app *App) setRoutes() {
	router := app.Container.Router
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to this life-changing API.\nIts the best API, its true, all other API's are fake.")
	})
}

func (app *App) getMongo() (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("mongodb://%s", os.Getenv("MONGODB_URI"))
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return client, err
}


type RequestHandlerFunc func(container *container.Container, w http.ResponseWriter, r *http.Request)

func (app *App) handleFunc(handler RequestHandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler(app.Container, writer, request)
	}
}

func (app *App) Run(host string) {
	fmt.Println("Server listening!")
	log.Fatal(http.ListenAndServe(host, app.Container.Router))
}