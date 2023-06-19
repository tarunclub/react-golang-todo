package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Task struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

var db *gorm.DB

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	db.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	db.First(&task, params["id"])
	task.Status = !task.Status
	db.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	db.First(&task, params["id"])
	db.Delete(&task)
	var tasks []Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todoapp password=secret sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Task{})

	router := mux.NewRouter()

	// Define allowed origins, methods, and headers
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
