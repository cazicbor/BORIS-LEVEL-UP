package repository

import (
	"strconv"
	"testing"
)

func TestGetTaskByID(t *testing.T) {

	var MockTasks = make(map[int]*Task)

	MockTasks[1] = &Task{
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

	for _, mocktask := range MockTasks {
		res := GetTaskByID(strconv.Itoa(mocktask.ID))
		if res != mocktask {
			t.Errorf("Error in getting the task, got: %v, want: %v", res, mocktask)
		}
	}
}

func TestAddTaskToDB(t *testing.T) {

}
