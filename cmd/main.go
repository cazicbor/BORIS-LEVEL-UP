package main

import (
	"fmt"

	"github.com/cazicbor/BORIS_LEVEL_UP/handlers"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository/mongostore"
)

var mh *mongostore.MongoHandler

func main() {

	mongoDbConnection := "mongodb://localhost:27017"
	mh = mongostore.NewMongoRepo(mongoDbConnection)
	//r := db.InitDBRepo()
	//repository.InitRepository(r) //init the repo with data
	fmt.Println("Rest API Boris v2.0 ")
	handlers.HandleRequests() //launch handlers

}
