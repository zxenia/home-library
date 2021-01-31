package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// our initial migration function
func initialMigration() {
	db, err := gorm.Open("sqlite3", "library.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Genre{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Place{})
}
