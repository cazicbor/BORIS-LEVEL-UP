package repository

import (
	"fmt"
	"strconv"
)

//Task struct
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}

//modification de la structure de données : slice -> map pour pouvoir utiliser la métode Put
var Tasks = make(map[int]*Task)

func GetAllIDs() map[int]*Task {
	return Tasks
}

func GetTaskByID(id string) (*Task, error) {
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Print(err)
	}
	if _, ok := Tasks[idToInt]; !ok {
		return nil, fmt.Errorf("IDs not matching")
	}
	return Tasks[idToInt], nil
}

func AddTaskToDB() *Task {
	var t Task

	Tasks[len(Tasks)] = &Task{ //we append the Task t to the map
		ID:          t.ID,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}
	return &t
}

func UpdateTaskByID() *Task {
	var t Task

	Tasks[t.ID] = &Task{
		ID:          t.ID,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}

	if _, ok := Tasks[t.ID]; !ok {
		fmt.Println("Task not found")
		return nil
	}
	return &t
}

func DeleteTaskByID() *Task {
	var t Task

	for index := range Tasks {
		if index == t.ID {
			delete(Tasks, index) //we delete the Task
		} else {
			fmt.Printf("id %v not found", t.ID)
		}
	}
	return &t
}

func InitRepo() map[int]*Task {
	Tasks[1] = &Task{
		ID:          1,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}
	Tasks[2] = &Task{
		ID:          2,
		Description: "Faire évoluer l'API : intégrer le routeur go-chi",
		Deadline:    "09/02/2022",
		Status:      "To do",
	}
	Tasks[3] = &Task{
		ID:          3,
		Description: "Créer tests unitaires",
		Deadline:    "10/02/2022",
		Status:      "To do",
	}
	Tasks[4] = &Task{
		ID:          4,
		Description: "Intégrer une persistance des données",
		Deadline:    "10/02/2022",
		Status:      "To do",
	}
	return Tasks
}
