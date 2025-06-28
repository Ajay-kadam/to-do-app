package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Ajay-kadam/to-do-app/metrics"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func Init(db *mongo.Database) {
	collection = db.Collection("tasks")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, newTask)
	if err != nil {
		http.Error(w, "Insert Failed", http.StatusInternalServerError)
		return
	}
	metrics.TasksCreated.Inc()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Query Failed", http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)
	var tasks []Task
	for cur.Next(ctx) {
		var t Task
		cur.Decode(&t)
		tasks = append(tasks, t)
	}
	json.NewEncoder(w).Encode(tasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, "errorin deleteing task", http.StatusInternalServerError)
		return
	}
	metrics.TasksDeleted.Inc()
	w.WriteHeader(http.StatusNoContent)

}
