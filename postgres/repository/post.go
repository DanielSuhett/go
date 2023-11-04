package repository

import "log"

type Post struct {
	id    int
	Title string
	Text  string
}

func (db *Database) GetPosts() ([]*Post, error) {
	rows, err := db.Query("select * from posts;")

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
