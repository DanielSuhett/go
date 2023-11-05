package post

import (
	"context"
	"errors"
	"fmt"
	"postgres/entities"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxRepository(writer, reader *sqlx.DB) *repository {
	return &repository{writer, reader}
}

func (repo *repository) Create(ctx context.Context, post *entities.Post) error {

	_, err := repo.writer.ExecContext(ctx, `
		INSERT INTO posts(
			title, 
			text
		) 
		VALUES(
			$1, 
			$2
		)
	`, post.Title, post.Text)

	if err != nil {
		return errors.New("fail to create post")
	}

	return nil
}

func (repo *repository) Update(ctx context.Context, id int, post *entities.Post) error {
	_, err := repo.writer.ExecContext(ctx, `
		UPDATE posts SET 
			title = $1, 
			text = $2 
		WHERE 
			id = $3
	`, post.Title, post.Text, id)

	if err != nil {
		return errors.New("fail to update post")
	}

	return nil
}

func (repo *repository) Get(ctx context.Context, id int) (*entities.Post, error) {

	var post entities.Post

	err := repo.reader.GetContext(ctx, &post, `
		select 
			id,
			title, 
			text, 
			created_at, 
			updated_at 
		from posts 
		where id = $1;
	`, id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("post not found")
	}

	return &post, nil

}

func (repo *repository) GetAll(ctx context.Context) ([]entities.Post, error) {

	var posts []entities.Post

	err := repo.reader.SelectContext(ctx, &posts, `
		SELECT 
			id,
			title,
			text,
			created_at,
			updated_at
		FROM posts 
	`)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("posts not found")
	}

	return posts, nil

}
