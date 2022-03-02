package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoHandlerSuite struct {
	suite.Suite
	db *mongo.Database
}

func (s *MongoHandlerSuite) SetupSuite() { //here we run everything that is "global", common to each test
	//ici mettre fonction qui init l'environnement : NewMongoRepo, quid de address ?
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

func TestMongoHandlerSuite(t *testing.T) {
	s := new(MongoHandlerSuite)
	suite.Run(t, s)
}

func (s *MongoHandlerSuite) TestNewMongoRepo() {

}

func (s *MongoHandlerSuite) TestGetTaskByID() {
}

func (s *MongoHandlerSuite) TestGetAllTasksByID() {

}

func (s *MongoHandlerSuite) TestAddTaskToDB() {

}

func (s *MongoHandlerSuite) TestUpdateTaskByID() {

}

func (s *MongoHandlerSuite) TestDeleteTaskByID() {
}
