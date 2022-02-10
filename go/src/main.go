package main

import (
    "fmt"
    "encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//task struct
type Task struct {
    ID string `json:"id"`
    Description string `json:"description"`
    Deadline string `json:"deadline"`
    Status string `json:"status"`
}


//slice used to store the tasks (à voir car pas de persistance des données à l'étape 1)
var tasks []Task

//handlers
func handleRequests() {
    r := chi.NewRouter() //creation of the routeur

    r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { //declare new routes to which we pass http handlers
		w.Write([]byte("home page"))
        fmt.Println("Endpoint Hit: homePage")
	})

    //= getAllTasks
    r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) { //voir quel endpoint mettre
        json.NewEncoder(w).Encode(tasks) //we use the writer and write the "items"
        fmt.Println("Endpoint Hit: getAllTasks")
    })

    //= createNewTask
    r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
        var t Task
        
        err := json.NewDecoder(r.Body).Decode(&t)
        if err !=nil {
            fmt.Println(err)
            return
        }
        
        tasks = append(tasks, t)

        w.Write([]byte("Superbe, tâche créée"))
    })

    //update an existing task
    r.Put("/tasks", func(w http.ResponseWriter, r *http.Request)

    http.ListenAndServe("localhost:8080", r)
}

/* func getAllTasks(w http.ResponseWriter, r *http.Request) {
}

//searchTask searches the tasks data for a matching task
func searchTask(w http.ResponseWriter, r *http.Request) {
} */

func main() {
    fmt.Println("Rest API Boris v2.0 ")
	tasks = []Task{
		{ID: "1", Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données", Deadline: "09/02/2022", Status: "Ongoing"},
		{ID: "2", Description: "Faire évoluer l'API : intégrer le routeur go-chi", Deadline: "09/02/2022", Status: "To do"},
 		{ID: "3", Description: "Créer tests unitaires", Deadline: "10/02/2022", Status: "To do"},
 		{ID: "4", Description: "Intégrer une persistance des données", Deadline: "10/02/2022", Status: "To do"},
	}
    handleRequests()
}