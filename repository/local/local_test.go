package local

import (
	"testing"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"github.com/stretchr/testify/assert"
)

func TestInitRepo(t *testing.T) {

	InitLocalRepo()

	assert.Equal(t, 1, tasks.Indice, "The two lenghts should be equal")
}

func TestGetRepository(t *testing.T) {

	testTasks := &localTasks{
		db:     make(map[Indice]*model.Task),
		Indice: 1,
	}
	tasks = testTasks

	tt := GetRepository()

	assert.NotNil(t, tt)
	assert.Equal(t, tasks, tt)
}

func TestGetAllTasks(t *testing.T) {

	var sliceTestTasks []*model.Task

	maptasks := map[Indice]*model.Task{ //map that we build for testing with test data
		1: &model.Task{
			ID:          1,
			Description: "test1",
			Deadline:    "test1",
			Status:      "test1",
		},
	}
	testTasks := &localTasks{
		db:     maptasks,
		Indice: 1,
	}
	tasks = testTasks

	for _, testTask := range testTasks.db {
		sliceTestTasks = append(sliceTestTasks, testTask)
	}

	res := tasks.GetAllTasksByID()

	assert.NotNil(t, res)
	assert.Equal(t, sliceTestTasks, res)
}

func TestGetTaskByID(t *testing.T) {

	var err error

	maptasks := map[Indice]*model.Task{ //map that we build for testing with test data
		1: &model.Task{
			ID:          1,
			Description: "test1",
			Deadline:    "test1",
			Status:      "test1",
		},
	}
	testTasks := &localTasks{
		db:     maptasks,
		Indice: 1,
	}
	tasks = testTasks

	tasks.db[Indice(testTasks.Indice)] = maptasks[Indice(testTasks.Indice)]
	test := maptasks[Indice(testTasks.Indice)]

	res, err := GetRepository().GetTaskByID(testTasks.Indice)
	assert.Nil(t, err)
	assert.Equal(t, test, res, "The two tasks should be the same.")

	res, err = GetRepository().GetTaskByID(66)

	assert.Nil(t, res)
	assert.NotNil(t, err)

}

func TestAddTaskToDB(t *testing.T) {

	maptasks := map[Indice]*model.Task{ //map that we build for testing with test data
		1: &model.Task{
			ID:          1,
			Description: "test1",
			Deadline:    "test1",
			Status:      "test1",
		},
	}
	testTasks := &localTasks{
		db:     maptasks,
		Indice: 1,
	}
	tasks = testTasks

	//tasks.db[Indice(testTasks.Indice)] = maptasks[Indice(testTasks.Indice)]
	newTask := maptasks[Indice(testTasks.Indice)]

	res, err := GetRepository().AddTaskToDB(maptasks[Indice(testTasks.Indice)])
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, newTask, res, "The two tasks should be the same")
}

func TestUpdateTaskByID(t *testing.T) {

	maptasks := map[Indice]*model.Task{ //map that we build for testing with test data
		1: &model.Task{
			ID:          1,
			Description: "test1",
			Deadline:    "test1",
			Status:      "test1",
		},
	}
	testTasks := &localTasks{
		db:     maptasks,
		Indice: 1,
	}
	tasks = testTasks

	//tasks.db[Indice(testTasks.Indice)] = maptasks[Indice(testTasks.Indice)]

	updatedTask := &model.Task{
		ID:          tasks.Indice,
		Description: "testupdate",
		Deadline:    "testupdate",
		Status:      "testupdate",
	}

	res, err := GetRepository().UpdateTaskByID(updatedTask)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, updatedTask, res, "The two tasks should be the same")

	res, err = GetRepository().UpdateTaskByID(&model.Task{ID: 48})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrNotFound)
	assert.Nil(t, res)

}

func TestDeleteTaskByID(t *testing.T) {

	maptasks := map[Indice]*model.Task{ //map that we build for testing with test data
		1: &model.Task{
			ID:          1,
			Description: "test1",
			Deadline:    "test1",
			Status:      "test1",
		},
	}
	testTasks := &localTasks{
		db:     maptasks,
		Indice: 1,
	}
	tasks = testTasks

	tasks.db[Indice(testTasks.Indice)] = maptasks[Indice(testTasks.Indice)]

	err := GetRepository().DeleteTaskByID(1)

	assert.Nil(t, err)
	assert.Equal(t, len(tasks.db), 0)

	err = GetRepository().DeleteTaskByID(7)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrNotFound)
}
