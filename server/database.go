package main

import (
	"apply/definition"
	"apply/util"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase() {
	db, openErr := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if openErr != nil {
		panic(openErr)
	}

	util.Database = db

	fmt.Println("Connected to database")

	util.MustNotError(
		db.AutoMigrate(
			&definition.Application{},
		),
	)

	fmt.Println("Migrated database")
}
