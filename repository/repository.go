package repository

import (
	"errors"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
)

type RepositoryProvider interface {
	GetAllTasksByID() []*model.Task
	GetTaskByID(id int) (*model.Task, error)
	AddTaskToDB(t *model.Task) (*model.Task, error)
	UpdateTaskByID(t *model.Task) (*model.Task, error)
	DeleteTaskByID(id int) error
}

var ( //custom errors
	ErrNotFound = errors.New("ID not found")
	//...
)

var store RepositoryProvider

func InitRepository(r RepositoryProvider) {
	store = r
}

func GetRepository() RepositoryProvider {
	return store
}
