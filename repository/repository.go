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

func GetTaskByID(id string) *Task {
	idToInt, err := strconv.Atoi(id)
	for _, task := range Tasks {
		if err != nil {
			fmt.Print(err)
		}
		if task.ID == idToInt {
			return task
		}
	}
	return nil
}

func InitRepo() {
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
}
