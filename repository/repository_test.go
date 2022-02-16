package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRepo(t *testing.T) {
	InitRepo()
	assert.Equal(t, 1, id, "The two lenghts should be equal")
}

func TestGetAllIDs(t *testing.T) {
	InitRepo()
	res := GetAllIDs()
	assert.NotNil(t, res)
	assert.Equal(t, Tasks, res, "The two maps should be equal")
}

func TestGetTaskByID(t *testing.T) {
	var err error
	testTask := &Task{
		ID:          1,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}
	Tasks[testTask.ID] = testTask

	res, err := GetTaskByID(testTask.ID, err)
	assert.Nil(t, err)
	assert.Equal(t, testTask, res, "The two IDs should be the same.")

	res, err = GetTaskByID(66, err)
	assert.Nil(t, res)
	assert.NotNil(t, err)

}

func TestAddTaskToDB(t *testing.T) {
	InitRepo()
	testTask := &Task{
		ID:          5,
		Description: "de toute manière je serai bientôt au chômage",
		Deadline:    "16/02/2022",
		Status:      "TO DO",
	}
	Tasks[testTask.ID] = testTask

	res, err := AddTaskToDB(testTask)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testTask, res, "The two tasks should be the same")
}

func TestUpdateTaskByID(t *testing.T) {
	InitRepo()
	testTask := &Task{
		ID:          id,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}

	Tasks[id] = testTask

	updatedTask := &Task{
		ID:          id,
		Description: "nouveau blabla",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}

	res, err := UpdateTaskByID(Tasks[id])
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, updatedTask, res, "The two tasks should be the same")

	res, err = UpdateTaskByID(Tasks[24])
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, res)

}

func TestDeleteTaskByID(t *testing.T) {
	InitRepo()
	testTask := &Task{
		ID:          1,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}
	Tasks[testTask.ID] = testTask

	err := DeleteTaskByID(1)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(Tasks))

	err = DeleteTaskByID(7)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
}
