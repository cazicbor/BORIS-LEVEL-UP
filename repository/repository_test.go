package repository

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRepo(t *testing.T) {
	res := InitRepo()
	assert.NotNil(t, res)
	assert.Equal(t, len(Tasks), len(res), "The two lenghts should be equal")

}

func TestGetAllIDs(t *testing.T) {

}

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

	res, err = GetTaskByID("mes couilles")
	assert.Nil(t, res)
	assert.NotNil(t, err)

}

/* func TestAddTaskToDB(t *testing.T) {
	var newTestTask *Task
	Tasks[newTestTask.ID] = &Task{
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
