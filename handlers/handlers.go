package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"github.com/go-chi/chi"
)

//handlers
func HandleRequests() {
	r := chi.NewRouter() //creation of the router

	r.Get("/", Index)
	r.Get("/tasks", GetAllTasks)
	r.Post("/task", CreateNewTask)
	r.Put("/task", UpdateTask)
	r.Delete("/task", DeleteTask)
	http.ListenAndServe("localhost:8080", r)
}

func Index(w http.ResponseWriter, r *http.Request) { //declare new routes to which we pass http handlers
	w.Write([]byte("home page"))
	fmt.Println("Endpoint Hit: homePage")
}

func GetTask(w http.ResponseWriter, r *http.Request) {

}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(repository.Tasks) //we use the writer and write the "items"
	if err != nil {
		log.Printf("Body encoding error, %v", err)
		w.WriteHeader(http.StatusInternalServerError) //internal server error
		return
	}

	fmt.Println("Endpoint Hit: getAllTasks")
	w.Write([]byte("Here are all the tasks"))
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var t repository.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(http.StatusInternalServerError) //internal server error
		return
	}

	repository.Tasks[len(repository.Tasks)] = &repository.Task{ //we append the repository.Task t to the map
		ID:          t.ID,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}

	w.Write([]byte("Great, new task created"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var t repository.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, ok := repository.Tasks[t.ID]; !ok {
		fmt.Println(err, "Element not found")
		return
	}
	repository.Tasks[t.ID] = &repository.Task{ //we append the repository.Task t to the map
		ID:          t.ID,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}

	w.Write([]byte("task updated"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var t repository.Task

	err := json.NewDecoder(r.Body).Decode(&t) //we decode the request body from byte format to JSON, in order to satisfy the interface followed by t

	if err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(http.StatusBadRequest) //bad request
		return
	}

	for index := range repository.Tasks {
		if index == t.ID {
			delete(repository.Tasks, index) //we delete the repository.Task
		} else {
			fmt.Printf("id %v not found", t.ID)
		}
	}
	w.Write([]byte("repository.Task deleted"))
}
