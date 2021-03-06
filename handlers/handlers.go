package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository/mongostore"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mh *mongostore.MongoHandler

//handlers
func HandleRequests() {

	r := chi.NewRouter() //creation of the router

	r.Route("/", func(r chi.Router) {
		r.Get("/", Index)
		r.Get("/task/{id}", GetTask)
		r.Get("/tasks", GetAllTasks)
		r.Post("/task", CreateNewTask)
		r.Put("/task", UpdateTask)
		r.Delete("/task", DeleteTask)
	})
	http.ListenAndServe("localhost:8080", r)
}

func Index(w http.ResponseWriter, r *http.Request) { //declare new routes to which we pass http handlers
	w.Write([]byte("Home page"))
	fmt.Println("Endpoint Hit: homePage")
}

func GetTask(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		fmt.Println("Error")
		return
	}
	task, err := repository.GetRepository().GetTaskByID(id)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(task)

	fmt.Println("Endpoint Hit: GetTask")
	w.Write([]byte("Here's your task"))
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(repository.GetRepository().GetAllTasksByID()) //we use the writer and write the "items"
	if err != nil {
		log.Printf("Body encoding error, %v", err)
		w.WriteHeader(http.StatusInternalServerError) //internal server error
		return
	}

	fmt.Println("Endpoint Hit: GetAllTasks")
	w.Write([]byte("Here are all the tasks"))
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {

	var t *model.Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("Body encoding error, %v", err)
		w.WriteHeader(http.StatusBadRequest) //internal server error
		return
	}

	_, err = repository.GetRepository().AddTaskToDB(t)
	if err != nil {
		log.Printf("Body encoding error, %v", err)
		w.WriteHeader(http.StatusBadRequest) //internal server error
		return
	}
	fmt.Println("Endpoint Hit: CreateNewTask")

	w.Write([]byte("Great, new task created"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	var t model.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = repository.GetRepository().UpdateTaskByID(&t)
	if err == repository.ErrNotFound {
		w.WriteHeader((http.StatusBadRequest))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Endpoint Hit: UpdateTask")
	w.Write([]byte("Task updated"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	var t model.Task
	err := json.NewDecoder(r.Body).Decode(&t) //we decode the request body from byte format to JSON, in order to satisfy the interface followed by t
	if err != nil {
		fmt.Println(err)
		return
	}
	id := t.ID.(primitive.ObjectID).Hex()
	err = repository.GetRepository().DeleteTaskByID(id)
	if err == repository.ErrNotFound {
		w.WriteHeader((http.StatusBadRequest))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Endpoint Hit: DeleteTask")
	w.Write([]byte("Task deleted"))
}
