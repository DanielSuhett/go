package models

import (
	"log"
	"postgres/database"
)

type Post struct {
	id    int
	Title string
	Text  string
}

type Database struct {
	DB *database.Database
}

func (d *Database) CreatePost(post *Post) (bool, error) {
	result, err := d.DB.Exec("INSERT INTO posts(title, text) VALUES($1, $2)", post.Title, post.Text)

	if err != nil {
		log.Fatal(err)
	}

	created, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return created > 0, nil
}

func (d *Database) GetPosts() ([]*Post, error) {
	rows, err := d.DB.Query("select title, text from posts;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	posts := make([]*Post, 0)

	for rows.Next() {
		post := new(Post)
		err = rows.Scan(&post.id, &post.Title, &post.Text)

		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (d *Database) GetPost(id int) (*Post, error) {
	rows, err := d.DB.Query("select title, text from posts where id = $1;", id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	post := new(Post)

	for rows.Next() {
		err = rows.Scan(&post.Title, &post.Text)
	}

	if err != nil {
		log.Fatal(err)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return post, nil
}

func (d *Database) UpdatePost(id int, post *Post) (bool, error) {
	result, err := d.DB.Exec("UPDATE posts SET title = $1, text = $2 WHERE id = $3", post.Title, post.Text, id)

	if err != nil {
		log.Fatal(err)
	}

	updated, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return updated > 0, nil
}
