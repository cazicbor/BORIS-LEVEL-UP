package mongostore

import (
	"context"
	"strconv"
	"testing"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const TaskCollection = "task"

type MongoHandlerSuite struct {
	suite.Suite
	taskStore repository.RepositoryProvider
	db        *mongo.Database
}

func (s *MongoHandlerSuite) SetupSuite() { //here we run everything that is "global", common to each test
	initEnvTest()
	//ici mettre fonction qui init l'environnement
	//...
}

//SetupTest runs before every unit test, in order to "clean" the DB, not to write over existing data
func (s *MongoHandlerSuite) SetupTest() {
	//ici intitaliser le MongoHandler de test?
	err := s.db.Drop(context.TODO())
	if err != nil {
		s.T().Fatal(err)
	}
}

//TearDownSuite disconnects from db
func (s *MongoHandlerSuite) TearDownSuite() {

	err := s.db.Client().Disconnect(context.TODO())
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestMongoRepoSuite(t *testing.T) {

	s := new(MongoHandlerSuite)
	suite.Run(t, s)
}

func (s *MongoHandlerSuite) TestGetTaskByID() {
	//items to be tested
	testTask := &model.Task{
		ID:          1,
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	insert, err := s.db.Collection(TaskCollection).InsertOne(context.TODO(), testTask)
	assert.Nil(s.T(), err)

	testTask.ID, _ = strconv.Atoi(insert.InsertedID.(primitive.ObjectID).Hex())
	result, err := s.taskStore.GetTaskByID(testTask.ID) //à corriger
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), testTask, result)

}

func (s *MongoHandlerSuite) TestGetAllTasksByID() {

	//items to be tested
	testTask1 := &model.Task{
		ID:          1,
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}
	testTask2 := &model.Task{
		ID:          2,
		Description: "test2",
		Deadline:    "test2",
		Status:      "test2",
	}

	_, err := s.db.Collection(TaskCollection).InsertOne(context.TODO(), testTask1)
	assert.Nil(s.T(), err)

	_, err = s.db.Collection(TaskCollection).InsertOne(context.TODO(), testTask2)
	assert.Nil(s.T(), err)

	slice := s.taskStore.GetAllTasksByID() //à corriger
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(slice))

}

func (s *MongoHandlerSuite) TestAddTaskToDB() {

	//items to be tested: we need to init db with data : tests can't be done without data!
	testTask := &model.Task{
		ID:          1,
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	var result *model.Task

	task, err := s.taskStore.AddTaskToDB(testTask) //à corriger
	assert.Nil(s.T(), err)
	IDToString := strconv.Itoa(task.ID)

	objectID, err := primitive.ObjectIDFromHex(IDToString)
	assert.Nil(s.T(), err)

	err = s.db.Collection(TaskCollection).FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objectID}}).Decode(result) //bson.D == slice
	assert.Nil(s.T(), err)
}

func (s *MongoHandlerSuite) TestUpdateTaskByID() {

	//items to be tested
	testTask := &model.Task{
		ID:          1,
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	updatedTask := &model.Task{
		ID:          testTask.ID,
		Description: "testupdate",
		Deadline:    "testupdate",
		Status:      "testupdate",
	}

	//insert, err := s.db.Collection(TaskCollection).InsertOne(context.TODO(), testTask)
	//assert.Nil(s.T(), err)

	res, err := s.taskStore.UpdateTaskByID(updatedTask)
	assert.NotNil(s.T(), res)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), updatedTask, res)

	res, err = s.taskStore.UpdateTaskByID(&model.Task{ID: 10})
	assert.NotNil(s.T(), err)
	assert.ErrorIs(s.T(), err, repository.ErrNotFound)
	assert.Nil(s.T(), res)
}

func (s *MongoHandlerSuite) TestDeleteTaskByID() {
	//items to be tested
	testTask := &model.Task{
		ID:          1,
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	var result *model.Task

	insert, err := s.db.Collection(TaskCollection).InsertOne(context.TODO(), testTask)
	assert.Nil(s.T(), err)

	testTask.ID, _ = strconv.Atoi(insert.InsertedID.(primitive.ObjectID).Hex())

	err = s.taskStore.DeleteTaskByID(testTask.ID) //à corriger
	assert.Nil(s.T(), err)

	err = s.db.Collection(TaskCollection).FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: insert.InsertedID.(primitive.ObjectID)}}).Decode(result)
	assert.ErrorIs(s.T(), err, mongo.ErrNoDocuments)
}