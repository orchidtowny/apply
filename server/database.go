package main

import (
	"apply/definition"
	"apply/util"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func setupDatabase() {
	db, openErr := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if openErr != nil {
		panic(openErr)
	}

	database = db

	fmt.Println("Connected to database")

	util.MustNotError(
		db.AutoMigrate(
			&definition.Application{},
			&definition.Invite{},
		),
	)

	fmt.Println("Migrated database")
}
