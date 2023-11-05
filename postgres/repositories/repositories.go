package repositories

import (
	"postgres/repositories/post"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	Post post.PostRepository
}

type Options struct {
	Writer *sqlx.DB
	Reader *sqlx.DB
}

func New(opts Options) *Container {
	return &Container{
		Post: post.NewSqlxRepository(opts.Writer, opts.Reader),
	}
}
