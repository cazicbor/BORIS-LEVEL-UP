package mongostore

import (
	"context"
	"log"

	"github.com/cazicbor/BORIS_LEVEL_UP/db"
	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoHandler struct {
	C *mongo.Collection
}

//MongoHandler Constructor, to init the repo
func NewMongoTaskStore(c string) *MongoHandler {
	db := db.GetDB()
	mh := &MongoHandler{
		C: db.Collection(c),
	}
	return mh
}

func (mh *MongoHandler) GetTaskByID(id string) (*model.Task, error) {

	var task *model.Task

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: objectID,
		},
	}

	err = mh.C.FindOne(context.TODO(), filter).Decode(&task)

	task.ID = id

	return task, err
}

func (mh *MongoHandler) GetAllTasksByID() []*model.Task {

	cur, err := mh.C.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil
	}

	defer cur.Close(context.TODO())

	var sliceTasks []*model.Task

	for cur.Next(context.TODO()) {
		task := &model.Task{}
		err := cur.Decode(task)

		if err != nil {
			log.Fatal(err)
		}

		sliceTasks = append(sliceTasks, task)
	}

	return sliceTasks
}

func (mh *MongoHandler) AddTaskToDB(t *model.Task) (*model.Task, error) {

	id := primitive.NewObjectID()
	t.ID = id
	_, err := mh.C.InsertOne(context.TODO(), t)
	if err != nil {
		return nil, repository.ErrNotFound
	}
	t.ID = t.ID.(primitive.ObjectID).Hex()
	return t, err
}

func (mh *MongoHandler) UpdateTaskByID(t *model.Task) (*model.Task, error) {

	id, err := primitive.ObjectIDFromHex(t.ID.(string))
	if err != nil {
		return nil, err
	}
	t.ID = id
	filter := bson.M{
		"_id": id,
	}

	old := mh.C.FindOneAndReplace(context.TODO(), filter, t)
	if old.Err() != nil {
		return nil, repository.ErrNotFound
	}

	return t, nil
}

func (mh *MongoHandler) DeleteTaskByID(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: oid,
		},
	}

	_, err = mh.C.DeleteOne(context.TODO(), filter)
	return err
}
