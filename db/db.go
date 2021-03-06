package db

import (
	"context"
	"log"
	"time"

	"github.com/cazicbor/BORIS_LEVEL_UP/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

//DB database instance

//SetUpDB sets up the mongo database with configuration parameters
func SetUpDB() {
	//var db = DB
	var usr string
	configuration := config.GetConfig()

	if configuration.DB.MongoDBUser != "" {
		usr = configuration.DB.MongoDBUser + ":" + configuration.DB.MongoDBPwd
	}
	mongoURI := "mongodb://" + usr + "@" + configuration.DB.MongoDBHost + ":" + configuration.DB.MongoDBPort

	Client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("[InitDB] : %s\n", err)
	}
	//TODO cancel func
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatalf("[InitDB] : %s\n", err)
	}
	if err = Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("[InitDB] : %s\n", err)
	}
	if configuration.DB.Mode == "prod" {
		DB = Client.Database(configuration.DB.Database)
	} else {
		//fmt.Println(configuration.DB.Mode)
		DB = Client.Database(configuration.DB.TestDatabase)
	}

}

//GetDB returns the DB instance
func GetDB() *mongo.Database {
	return DB
}
