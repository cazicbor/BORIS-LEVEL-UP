package repository

import (
	"testing"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/stretchr/testify/assert"
)

func TestInitRepo(t *testing.T) {
	InitRepo()
	assert.Equal(t, 1, tasks.indice, "The two lenghts should be equal")
}

func TestGetRepository(t *testing.T) {
	InitRepo()
	res := GetRepository()
	assert.NotNil(t, res)
	assert.Equal(t, tasks, res)
}

func TestGetAllIDs(t *testing.T) {
	InitRepo()
	res := GetRepository().GetAllIDs()
	assert.NotNil(t, res)

	//assert.Equal(t, tasks, res)
}

func TestGetTaskByID(t *testing.T) {
	var err error
	testTask := &model.Task{
		ID:          1,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}
	tasks.db[testTask.ID] = testTask

	res, err := GetRepository().GetTaskByID(testTask.ID, err)
	assert.Nil(t, err)
	assert.Equal(t, testTask, res, "The two IDs should be the same.")

	res, err = GetRepository().GetTaskByID(66, err)
	assert.Nil(t, res)
	assert.NotNil(t, err)

}

func TestAddTaskToDB(t *testing.T) {
	InitRepo()
	testTask := &model.Task{
		ID:          5,
		Description: "de toute manière je serai bientôt au chômage",
		Deadline:    "16/02/2022",
		Status:      "TO DO",
	}
	tasks.db[testTask.ID] = testTask

	res, err := GetRepository().AddTaskToDB(testTask)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testTask, res, "The two tasks should be the same")
}

func TestUpdateTaskByID(t *testing.T) {
	InitRepo()
	testTask := &model.Task{
		ID:          tasks.indice,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}

	tasks.db[tasks.indice] = testTask

	updatedTask := &model.Task{
		ID:          tasks.indice,
		Description: "nouveau blabla",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}

	res, err := GetRepository().UpdateTaskByID(tasks.db[tasks.indice])
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, updatedTask, res, "The two tasks should be the same")

	res, err = GetRepository().UpdateTaskByID(tasks.db[24])
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, res)

}

func TestDeleteTaskByID(t *testing.T) {
	InitRepo()
	testTask := &model.Task{
		ID:          1,
		Description: "Construire une API REST en utilisant uniquement la librairie standard, sans persistance des données",
		Deadline:    "09/02/2022",
		Status:      "Ongoing",
	}
	tasks.db[testTask.ID] = testTask

	err := GetRepository().DeleteTaskByID(1)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(tasks.db))

	err = GetRepository().DeleteTaskByID(7)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
}
