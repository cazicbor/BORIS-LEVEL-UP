package local

import (
	"fmt"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
)

var tasks *localTasks //singleton : une seule instanciation de cette structure, interne au package, qu'on transmet à l'extérieur via GetRepository()

type indice int

type localTasks struct { //strcuture locale au package repository dans laquelle on stocke les tâches
	db     map[indice]*model.Task
	indice int
}

func InitLocal() *localTasks {
	tasks = &localTasks{
		db:     make(map[indice]*model.Task),
		indice: 1,
	}
	return tasks
}

func GetRepository() *localTasks { //méthode permettant d'accéder aux tâches via l'extérieur (getter)
	return tasks
}

func (repo *localTasks) GetAllTasksByID() []*model.Task { //(repo *localTasks) signifie qu'on travaille avec la structure de données localTasks
	var sliceTasks []*model.Task //slice in which we store the tasks which are in the db
	for _, task := range repo.db {
		sliceTasks = append(sliceTasks, task)
	}
	return sliceTasks
}

func (repo *localTasks) GetTaskByID(id int) (*model.Task, error) {
	i := indice(id)
	if _, ok := tasks.db[i]; !ok {
		return nil, fmt.Errorf("ID not found")
	}
	return tasks.db[indice(id)], nil
}

func (repo *localTasks) AddTaskToDB(t *model.Task) (*model.Task, error) {
	tasks.db[indice(tasks.indice)] = &model.Task{ //we append the model.Task t to the map
		ID:          tasks.indice,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}
	tasks.indice++
	return t, nil
}

func (repo *localTasks) UpdateTaskByID(t *model.Task) (*model.Task, error) {
	if _, ok := tasks.db[indice(t.ID.(int))]; !ok {
		return nil, repository.ErrNotFound
	}
	tasks.db[indice(t.ID.(int))] = t
	return t, nil
}

func (repo *localTasks) DeleteTaskByID(id int) error {
	if _, ok := tasks.db[indice(id)]; !ok {
		return repository.ErrNotFound
	}
	delete(tasks.db, indice(id))
	return nil
}
