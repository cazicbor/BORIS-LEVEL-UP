package mongostore

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"github.com/cazicbor/BORIS_LEVEL_UP/db"
	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DefaultDatabase = "todolist"

const collectionName = "task"

type MongoHandler struct {
	C *mongo.Collection
}

//MongoHandler Constructor, to init the repo
func NewMongoTaskStore() *MongoHandler {
	db := db.GetDB()
	mh := &MongoHandler{
		C: db.Collection(collectionName),
	}
	return mh
}

func (mh *MongoHandler) GetTaskByID(id int) (*model.Task, error) {

	var task *model.Task

	stringID := strconv.Itoa(id)
	hxID := hex.EncodeToString([]byte(stringID))

	objectId, err := primitive.ObjectIDFromHex(hxID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(objectId)

	filter := bson.M{
		"_id": objectId,
	}

	err = mh.C.FindOne(context.TODO(), filter).Decode(task)
	if err != nil {
		log.Fatal(err)
	}

	return task, nil
}

func (mh *MongoHandler) GetAllTasksByID() []*model.Task {

	cur, err := mh.C.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
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

	result, err := mh.C.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}
	t.ID, _ = strconv.Atoi(result.InsertedID)
	return t, err
}

func (mh *MongoHandler) UpdateTaskByID(t *model.Task) (*model.Task, error) {

	filter := bson.M{
		"_id": t.ID,
	}

	task := mh.C.FindOne(context.TODO(), filter)

	update, err := mh.C.UpdateOne(context.TODO(), filter, task)
	if err != nil {
		return nil, fmt.Errorf("ID not found")
	}

	t.ID = update.UpsertedID.(int)
	return t, nil
}

func (mh *MongoHandler) DeleteTaskByID(id int) error {

	filter := bson.M{
		"_id": id,
	}

	_, err := mh.C.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("could not delete task : %v", id)
	}

	return nil
}
