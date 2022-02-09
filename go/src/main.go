package main

import (
	"fmt"
	"log"
	"net/http"
    "encoding/json"
    "bytes"
    "io/ioutil"
)

//task struct

type Task struct {
    ID string `json:"id"`
    Description string `json:"description"`
    Deadline string `json:"deadline"`
    Status string `json:"status"`
}

type taskHandlers struct {
    store map[string]Task
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
    http.HandleFunc("/post", createTask)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//handlers

func getAllTasks(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: getAllTasks")
    json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
    postBody, _ := json.Marshal(map[string]string{
        "ID":  "5",
        "Description": "jsp",
        "Deadline" : "01/01/01",
        "Status" : "To do",
     })
     responseBody := bytes.NewBuffer(postBody)
  //Leverage Go's HTTP Post function to make request
     resp, err := http.Post("localhost:8080/post", "application/json", responseBody)
  //Handle Error
     if err != nil {
        log.Fatalf("An Error Occured %v", err)
     }
     defer resp.Body.Close()
  //Read the response body
     body, err := ioutil.ReadAll(resp.Body)
     if err != nil {
        log.Fatalln(err)
     }
     sb := string(body)
     log.Printf(sb)
}

// func createTask(w http.ResponseWriter, r *http.Request) {
//     jsonData := []Task{Task{"5", "jsp", "12/02/2022", "ToDo"}}
//     jsonValue, _ := json.Marshal(jsonData)
//     request, _ := http.NewRequest("POST", "localhost:8080/post", bytes.NewBuffer(jsonValue))
//     request.Header.Set("Content-Type", "application/json")
//     client := &http.Client{}
//     response, err := client.Do(request)

//     if err!= nil {
//         fmt.Println("Request has failed with %s", err)
//     } else {
//         data, _ := ioutil.ReadAll(response.Body)
//         fmt.Println(string(data))
//     }
// }


// func createTask(w http.ResponseWriter, r *http.Request) error {
// 	var t task
//     if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
//         internalServerError(w, r)
//         return
//     }
//     h.store.Lock()
//     h.store.m[t.ID] = t
//     h.store.Unlock()
//     jsonBytes, err := json.Marshal(t)
//     if err != nil {
//         internalServerError(w, r)
//         return
//     }
//     w.WriteHeader(http.StatusOK)
//     w.Write(jsonBytes)
// }


// func updateTask(w http.ResponseWriter, r *http.Request) error {
// 	u := new(user)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	users[id].Name = u.Name
// 	return c.JSON(http.StatusOK, users[id])
// }

// func deleteTask(w http.ResponseWriter, r *http.Request) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	delete(users, id)
// 	return c.NoContent(http.StatusNoContent)
// }

func internalServerError(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("not found"))
}


func main() {
	tasks = []Task{
		{ID: "1", Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données", Deadline: "09/02/2022", Status: "Ongoing"},
		{ID: "2", Description: "Faire évoluer l'API : intégrer le routeur go-chi", Deadline: "09/02/2022", Status: "To do"},
 		{ID: "3", Description: "Créer tests unitaires", Deadline: "10/02/2022", Status: "To do"},
 		{ID: "4", Description: "Intégrer une persistance des données", Deadline: "10/02/2022", Status: "To do"},
	}
    handleRequests()
}
