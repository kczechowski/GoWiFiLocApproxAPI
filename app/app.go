package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/container"
	"github.com/kczechowski/GoWiFiLocApproxAPI/app/handlers"
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
	app.Container.MongoClient = mongoClient

	//Use closure to retrieve mongo database instance
	app.Container.MongoDatabase = func() *mongo.Database {
		return mongoClient.Database(os.Getenv("MONGODB_DATABASE"))
	}

	app.Container.Router = mux.NewRouter()
	app.setRoutes()

}

func (app *App) setRoutes() {
	app.Get("/", app.handleFunc(handlers.GetIndex))
	app.Get("/networks/", app.handleFunc(handlers.GetNetworks))
	app.Post("/networks/", app.handleFunc(handlers.PostNetwork))
}

func (app *App) getMongo() (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("mongodb://%s", os.Getenv("MONGODB_URI"))
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI).SetAuth(options.Credential{
		AuthSource: os.Getenv("MONGODB_DATABASE"), Username: os.Getenv("MONGODB_USERNAME"), Password: os.Getenv("MONGODB_PASSWORD"),
	}))

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

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Container.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Container.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Container.Router.HandleFunc(path, f).Methods("PUT")
}

func (app *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Container.Router.HandleFunc(path, f).Methods("PATCH")
}

func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Container.Router.HandleFunc(path, f).Methods("DELETE")
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