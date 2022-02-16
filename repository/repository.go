package repository

import (
	"errors"
	"fmt"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
)

// type RepositoryProvider interface {
// 	GetAllIDs() []*model.Task
// 	GetTaskByID(id int, err error) (*Task, error)
// 	AddTaskToDB(t *Task) (*Task, error)
// 	UpdateTaskByID(t *Task) (*Task, error)
// 	DeleteTaskByID(id int) error
// }

var ( //custom errors
	ErrNotFound = errors.New("ID not found")
	//...
)

var tasks *localTasks //singleton : une seule instanciation de cette structure, interne au package, qu'on transmet à l'extérieur via GetRepository()

type localTasks struct { //strcuture locale au package repository dans laquelle on stocke les tâches
	db     map[int]*model.Task
	indice int
}

func InitRepo() {
	tasks = &localTasks{
		db:     make(map[int]*model.Task),
		indice: 1,
	}
}

func GetRepository() *localTasks { //méthode permettant d'accéder aux tâches via l'extérieur (getter)
	return tasks
}

func (repo *localTasks) GetAllIDs() []*model.Task { //(repo *localTasks) signifie qu'on travaille avec la structure de données localTasks
	var sliceTasks []*model.Task //slice in which we store the tasks which are in the db
	for _, task := range repo.db {
		sliceTasks = append(sliceTasks, task)
	}
	return sliceTasks
}

func (repo *localTasks) GetTaskByID(id int, err error) (*model.Task, error) {
	if _, ok := tasks.db[id]; !ok {
		return nil, fmt.Errorf("ID not found")
	}
	return tasks.db[id], nil
}

func (repo *localTasks) AddTaskToDB(t *model.Task) (*model.Task, error) {
	tasks.db[tasks.indice] = &model.Task{ //we append the Task t to the map
		ID:          tasks.indice,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}
	tasks.indice++
	return t, nil
}

func (repo *localTasks) UpdateTaskByID(t *model.Task) (*model.Task, error) {
	if _, ok := tasks.db[t.ID]; !ok {
		return nil, ErrNotFound
	}
	tasks.db[t.ID] = t
	return t, nil
}

func (repo *localTasks) DeleteTaskByID(id int) error {
	if _, ok := tasks.db[id]; !ok {
		return ErrNotFound
	}
	delete(tasks.db, id)
	return nil
}
