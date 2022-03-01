package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cazicbor/BORIS_LEVEL_UP/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultDatabase = "todolist"

const CollectionName = "task"

type MongoHandler struct {
	client   *mongo.Client
	database string
}

//MongoHandler Constructor
func NewMongoRepo(address string) *MongoHandler {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(address))

	mh := &MongoHandler{
		client:   cl,
		database: DefaultDatabase,
	}

	return mh
}

func (mh *MongoHandler) GetTaskByID(id int) (*model.Task, error) {

	var task *model.Task

	collection := mh.client.Database(mh.database).Collection(CollectionName)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(task)
	if err != nil {
		fmt.Errorf("ID not found")
	}

	return task, nil
}

func (mh *MongoHandler) GetAllTasksByID() []*model.Task {

	collection := mh.client.Database(mh.database).Collection(CollectionName)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	var sliceTasks []*model.Task

	for cur.Next(ctx) {
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

	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.InsertOne(ctx, t)
	if err != nil {
		log.Fatal(err)
	}
	t.ID = result.InsertedID.(int)
	return t, err
}

func (mh *MongoHandler) UpdateTaskByID(t *model.Task) (*model.Task, error) {

	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": t.ID}

	task := collection.FindOne(ctx, filter)

	update, err := collection.UpdateOne(ctx, filter, task)
	if err != nil {
		return nil, fmt.Errorf("ID not found")
	}
	t.ID = update.UpsertedID.(int)
	return t, nil //ici convertir update en *model.Task, COMMENT FAIRE PUTAIN
}

func (mh *MongoHandler) DeleteTaskByID(id int) error {

	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": id}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("Could not delete task : %v", id)
	}

	return nil
}
