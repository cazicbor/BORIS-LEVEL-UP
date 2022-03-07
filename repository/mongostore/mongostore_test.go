package mongostore

import (
	"context"
	"fmt"
	"testing"

	"github.com/cazicbor/BORIS_LEVEL_UP/db"
	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const taskCollection = "task"

type MongoHandlerSuite struct {
	suite.Suite
	taskStore repository.RepositoryProvider
	db        *mongo.Database
}

func (s *MongoHandlerSuite) SetupSuite() { //here we run everything that is "global", common to each test
	initEnvTest()
	store := NewMongoTaskStore()
	s.taskStore = store
	s.db = db.GetDB()
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
		ID:          primitive.NewObjectID(),
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	insert, err := s.db.Collection(taskCollection).InsertOne(context.TODO(), testTask)
	assert.Nil(s.T(), err)
	fmt.Println(err)

	testTask.ID = insert.InsertedID.(primitive.ObjectID).Hex()
	result, err := s.taskStore.GetTaskByID(insert.InsertedID.(primitive.ObjectID).Hex())
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), testTask, result)
}

func (s *MongoHandlerSuite) TestGetAllTasksByID() {

	//items to be tested
	testTask1 := &model.Task{
		ID:          primitive.NewObjectID(),
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	_, err := s.db.Collection(taskCollection).InsertOne(context.TODO(), testTask1)
	assert.Nil(s.T(), err)

	slice := s.taskStore.GetAllTasksByID()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(slice))

}

func (s *MongoHandlerSuite) TestUpdateTaskByID() {

	//items to be tested
	testTask := &model.Task{
		ID:          primitive.NewObjectID(),
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

	testTask := &model.Task{
		ID:          primitive.NewObjectID(),
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	var result *model.Task

	insert, err := s.db.Collection(taskCollection).InsertOne(context.TODO(), testTask)
	assert.Nil(s.T(), err)

	err = s.taskStore.DeleteTaskByID(insert.InsertedID.(primitive.ObjectID).Hex())
	assert.Nil(s.T(), err)

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: insert.InsertedID.(primitive.ObjectID),
		},
	}
	err = s.db.Collection(taskCollection).FindOne(context.TODO(), filter).Decode(result)
	assert.ErrorIs(s.T(), err, mongo.ErrNoDocuments)
}

func (s *MongoHandlerSuite) TestAddTaskToDB() {

	//items to be tested: we need to init db with data : tests can't be done without data!
	testTask := &model.Task{
		ID:          primitive.NewObjectID(),
		Description: "test1",
		Deadline:    "test1",
		Status:      "test1",
	}

	var task *model.Task

	id, err := s.taskStore.AddTaskToDB(testTask)
	assert.Nil(s.T(), err)
	objectID, err := primitive.ObjectIDFromHex(id.ID.(string))

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: objectID,
		},
	}

	err = s.db.Collection(taskCollection).FindOne(context.TODO(), filter).Decode(task)
	assert.Nil(s.T(), err)

}
