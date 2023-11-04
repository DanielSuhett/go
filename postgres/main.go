package main

import (
	"cron-postgres/repository"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Env struct {
	db *repository.Database
}

func main() {
	db, err := repository.Connect()

	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	env.Search()

	repository.Close(db)
}

func (env *Env) Search() {
	posts, err := env.db.GetPosts()

	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		fmt.Printf("%v - %v\n", post.Title, post.Text)
	}
}
