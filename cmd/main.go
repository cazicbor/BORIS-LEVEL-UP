package main

import (
	"fmt"

	"github.com/cazicbor/BORIS_LEVEL_UP/handlers"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
)

func main() {
	fmt.Println("Rest API Boris v2.0 ")
	repository.InitRepo()
	handlers.HandleRequests() //launch handlers
}
