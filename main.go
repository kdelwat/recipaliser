package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kdelwat/recipaliser/cmd"
	"github.com/kdelwat/recipaliser/db"
	"github.com/kdelwat/recipaliser/model"
)

func main() {
	var err error
	db.Db, err = gorm.Open("sqlite3", "./recipaliser.db")

	if err != nil {
		fmt.Println("Failed to create database")
		fmt.Println(err)
	}

	defer db.Db.Close()

	db.Db.AutoMigrate(&model.Ingredient{})
	cmd.Execute()
}
