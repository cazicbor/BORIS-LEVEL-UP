package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func InitMockTasks() {
	MockTasks[1] = &Task{
		ID:          1,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}
	MockTasks[2] = &Task{
		ID:          2,
		Description: "Faire évoluer l'API : intégrer le routeur go-chi",
		Deadline:    "09/02/2022",
		Status:      "To do",
	}
	MockTasks[3] = &Task{
		ID:          3,
		Description: "Créer tests unitaires",
		Deadline:    "10/02/2022",
		Status:      "To do",
	}
	MockTasks[4] = &Task{
		ID:          4,
		Description: "Intégrer une persistance des données",
		Deadline:    "10/02/2022",
		Status:      "To do",
	}
} */

func TestGetTaskByID(t *testing.T) {
	testTask := &Task{
		ID:          1,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}
	Tasks[testTask.ID] = testTask

	res, err := GetTaskByID(strconv.Itoa(testTask.ID))
	assert.Nil(t, err)
	assert.Equal(t, testTask, res, "The two IDs should be the same.")

	res, err = GetTaskByID("64")
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

/* func TestAddTaskToDB(t *testing.T) {
	var newMockTask *Task
	MockTasks[newMockTask.ID] = &Task{
		ID:          newMockTask.ID,
		Description: newMockTask.Description,
		Deadline:    newMockTask.Deadline,
		Status:      newMockTask.Status,
	}

	res := AddTaskToDB()
	if res != newMockTask { //we compare both results
		t.Errorf("Error in creating a new task, got: %v, want: %v", res, newMockTask)
	}
}

func TestUpdateTaskByID(t *testing.T) {

}

func TestDeleteTaskByID(t *testing.T) {
	deletedTask := DeleteTaskByID()
	for _, mockTask := range MockTasks {
		if deletedTask.ID == mockTask.ID {
			t.Errorf("Error in deleting the task %v", deletedTask.ID)
		}
	}
}
*/
