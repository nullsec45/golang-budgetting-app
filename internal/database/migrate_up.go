package main

import (
	"budgetting-app/internal/models"
	"budgetting-app/common"
	"log"
)

func main() {
	db, err := common.NewMySQL()

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic(err)
	}

	log.Println("Migration completed....")
}