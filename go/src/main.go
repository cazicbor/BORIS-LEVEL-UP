package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//task struct
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}

//slice used to store the tasks (à voir car pas de persistance des données à l'étape 1)
var tasks []Task

func index(w http.ResponseWriter, r *http.Request) { //declare new routes to which we pass http handlers
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) //method not allowed
		return
	}
	w.Write([]byte("home page"))
	fmt.Println("Endpoint Hit: homePage")
	w.WriteHeader(http.StatusOK)
}

//handlers
func handleRequests() {
	r := chi.NewRouter() //creation of the router

	r.Get("/", index)

	//= getAllTasks
	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) { //voir quel endpoint mettre
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := json.NewEncoder(w).Encode(tasks) //we use the writer and write the "items"
		if err != nil {
			log.Printf("Body encoding error, %v", err)
			w.WriteHeader(http.StatusInternalServerError) //internal server error
			return
		}

		fmt.Println("Endpoint Hit: getAllTasks")
		w.Write([]byte("Et voici les tâches"))
		w.WriteHeader(http.StatusOK) //tout va bien
	})

	//= createNewTask
	r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed) //method not allowed
			return
		}
		var t Task

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Printf("Body parse error, %v", err)
			w.WriteHeader(http.StatusInternalServerError) //internal server error (?)
			return
		}
		tasks = append(tasks, t)

		w.Write([]byte("Superbe, tâche créée"))
		w.WriteHeader(http.StatusOK)
	})

	//update an existing task
	r.Put("/tasks", func(w http.ResponseWriter, r *http.Request) { //A FAIRE

	})

	//delete existing task
	r.Delete("/tasks", func(w http.ResponseWriter, r *http.Request) {
		var t Task

		err := json.NewDecoder(r.Body).Decode(&t) //we decode the request body from byte format to JSON, in order to satisfy the interface followed by t

		if err != nil {
			log.Printf("Body parse error, %v", err)
			w.WriteHeader(http.StatusBadRequest) //bad request
			return
		}

		for index, taskk := range tasks {
			if taskk.ID == t.ID {
				tasks = append(tasks[:index], tasks[index+1:]...) //we delete the task
			} else {
				fmt.Println("id %v not found", taskk.ID)
			}
		}
		w.Write([]byte("Tâche bien supprimée"))
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe("localhost:8080", r)
}

func main() {
	fmt.Println("Rest API Boris v2.0 ")
	tasks = []Task{
		{ID: 1, Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données", Deadline: "09/02/2022", Status: "Ongoing"},
		{ID: 2, Description: "Faire évoluer l'API : intégrer le routeur go-chi", Deadline: "09/02/2022", Status: "To do"},
		{ID: 3, Description: "Créer tests unitaires", Deadline: "10/02/2022", Status: "To do"},
		{ID: 4, Description: "Intégrer une persistance des données", Deadline: "10/02/2022", Status: "To do"},
	}
	handleRequests()
}
