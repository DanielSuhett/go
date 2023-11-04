package main

import (
	"fmt"
	"log"
	"postgres/database"
	"postgres/models"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	postService := &models.Database{DB: db}

	post, _ := postService.GetPost(1)

	fmt.Println(post)
}
