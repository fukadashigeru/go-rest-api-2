package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/domain/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Scuccessfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
