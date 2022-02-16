package repository

import (
	"errors"
	"fmt"
)

var id int

//Task struct
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}

var (
	ErrNotFound = errors.New("ID not found")
)

//modification de la structure de données : slice -> map pour pouvoir utiliser la métode Put
var Tasks = make(map[int]*Task)

func InitRepo() {
	id = 1
}

func GetAllIDs() map[int]*Task {
	return Tasks
}

func GetTaskByID(id int, err error) (*Task, error) {
	if _, ok := Tasks[id]; !ok {
		return nil, fmt.Errorf("IDs not matching")
	}
	return Tasks[id], nil
}

func AddTaskToDB(t *Task) (*Task, error) {
	Tasks[id] = &Task{ //we append the Task t to the map
		ID:          id,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}
	id++
	return t, nil
}

func UpdateTaskByID(t *Task) (*Task, error) {
	if _, ok := Tasks[t.ID]; !ok {
		return nil, ErrNotFound
	}
	Tasks[t.ID] = t
	return Tasks[id], nil
}

func DeleteTaskByID(id int) error {
	if _, ok := Tasks[id]; !ok {
		return ErrNotFound
	}
	delete(Tasks, id)
	return nil
}
