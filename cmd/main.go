package main

import (
	"fmt"

	"github.com/cazicbor/BORIS_LEVEL_UP/handlers"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository/local"
)

func main() {
	r := local.InitLocalRepo()
	repository.InitRepository(r) //init the repo with data
	fmt.Println("Rest API Boris v2.0 ")
	handlers.HandleRequests() //launch handlers

}
