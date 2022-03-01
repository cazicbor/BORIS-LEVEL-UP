package db

import (
	"testing"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewHandler(t *testing.T) {

}

func TestGetTaskByID(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		//taskCollection := mt.Coll
		expectedData := &model.Task{
			ID:          1,
			Description: "test1",
			Deadline:    "test1",
			Status:      "test1",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"id", expectedData.ID},
			{"description", expectedData.Description},
			{"deadline", expectedData.Deadline},
			{"status", expectedData.Status},
		}))

		taskResponse, err := GetTaskByID(expectedData.ID)

		assert.Nil(t, err)
		assert.Equal(t, &expectedData, taskResponse)
	})
}

func TestGetAllTasksByID(t *testing.T) {

	var sliceTestTasks []*model.Task

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		taskCollection = mt.Coll
		firstExpectedData := &model.Task{
			ID:          1,
			Description: "test1",
			Deadline:    "test1",
			Status:      "test1",
		}
		secondExpectedData := &model.Task{
			ID:          2,
			Description: "test2",
			Deadline:    "test2",
			Status:      "test2",
		}

		/*first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"id", firstExpectedData.ID},
			{"description", firstExpectedData.Description},
			{"deadline", firstExpectedData.Deadline},
			{"status", firstExpectedData.Status},
		})

		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"id", secondExpectedData.ID},
			{"description", secondExpectedData.Description},
			{"deadline", secondExpectedData.Deadline},
			{"status", secondExpectedData.Status},
		})

		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		tasks, err :=
		*/
	})
}

func TestAddTaskToDB(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
	})
}

func TestUpdateTaskByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
	})
}

func TestDeleteTaskByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
	})
}
