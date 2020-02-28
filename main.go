package main

import (
	"fmt"
	"github.com/kravcs/db-interface/mongo"
)

type DataStore interface {
	GetEmployee(id int) string
}

type App struct {
	ds DataStore
}

func main()  {
	app1 := App{ds: &mongo.MongoStore{}}
	fmt.Print(app1.ds.GetEmployee(12))
}
