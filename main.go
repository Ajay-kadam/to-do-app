package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Ajay-kadam/to-do-app/metrics"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("taskdb")
	Init(db)
	metrics.Init()

	r := mux.NewRouter()
	r.HandleFunc("/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/tasks", GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	r.Handle("/metrics", promhttp.Handler())

	log.Println("Serving at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
