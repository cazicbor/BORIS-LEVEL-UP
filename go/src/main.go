package main

import (
	"fmt"
	"log"
	"net/http"
    "encoding/json"
    "io/ioutil"
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

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Benvenue sur la page d'accueil !")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/tasks", getAllTasks)
    http.HandleFunc("/task", createTask)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//handlers
func getAllTasks(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: getAllTasks")
    json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Endpoint Hit: createTask")
    fmt.Fprintf(w, "%v", string(reqBody))

    var t Task 
    err = json.NewDecoder(r.Body).Decode(&t)
    if err !=nil {
        fmt.Println(err)
        return
    }
    tasks = append(tasks, t)

    json.NewEncoder(w).Encode(t)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
    var t Task 
    err := json.NewDecoder(r.Body).Decode(&t)
    if err !=nil {
        fmt.Println(err)
        return
    }
    for _, taskk := range tasks {
        if taskk.ID == t.ID {
            taskk = t
        }
    }
    fmt.Println("Endpoint Hit: updateTask")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
    var t Task 
    err := json.NewDecoder(r.Body).Decode(&t) //we decode the request body from byte format to JSON, in order to satisfy the interface followed by t 
    if err !=nil {
        fmt.Println(err)
        return
    }
    for index, taskk := range tasks {
        if taskk.ID == t.ID {  
            tasks = append(tasks[:index], tasks[index+1:]...) //we delete the task 
        }
    }
    fmt.Println("Endpoint Hit: deleteTask")
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("not found"))
}

func main() {
    fmt.Println("Rest API Boris v1.0 ")
	tasks = []Task{
		{ID: "1", Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données", Deadline: "09/02/2022", Status: "Ongoing"},
		{ID: "2", Description: "Faire évoluer l'API : intégrer le routeur go-chi", Deadline: "09/02/2022", Status: "To do"},
 		{ID: "3", Description: "Créer tests unitaires", Deadline: "10/02/2022", Status: "To do"},
 		{ID: "4", Description: "Intégrer une persistance des données", Deadline: "10/02/2022", Status: "To do"},
	}
    handleRequests()
}
