package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	fmt.Println("Home Library API")

	// Run migrations
	initialMigration()

	// Handle Subsequent requests
	handleRequests()
}
