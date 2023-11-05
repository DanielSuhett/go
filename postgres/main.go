package main

import (
	"context"
	"fmt"
	"postgres/config"
	"postgres/repositories"
)

func main() {
	repo := repositories.New(repositories.Options{
		Writer: config.GetWriterSqlx(),
		Reader: config.GetReaderSqlx(),
	})

	v, _ := repo.Post.Get(context.Background(), 1)

	fmt.Println(v)
}
