package main

import (
	"fmt"
	"log"
	"skillfactory/31_DB_APPS/pkg/storage"
	"skillfactory/31_DB_APPS/pkg/storage/postgres"
)

func main() {
	var db storage.Data

	db, err := postgres.New("postgres://postgres:postgres@localhost/tasks")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer db.Close()

	tasks, err := db.Tasks()
	if err != nil {
		log.Println("Not Tasks", err)
	}

	for _, t := range tasks {
		fmt.Println(t.Title, t.Content, t.AuthorID)
	}

}
