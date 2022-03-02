package main

import (
	"fmt"

	"github.com/cazicbor/BORIS_LEVEL_UP/handlers"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository/db"
)

var mh *db.MongoHandler

func main() {

	mongoDbConnection := "mongodb://localhost:8080"
	mh = db.NewMongoRepo(mongoDbConnection)
	//r := db.InitDBRepo()
	//repository.InitRepository(r) //init the repo with data
	fmt.Println("Rest API Boris v2.0 ")
	handlers.HandleRequests() //launch handlers

}
