package entities

import "time"

type Post struct {
	ID        int        `json:"id"`
	Title     string     `db:"title" json:"title"`
	Text      string     `db:"text" json:"text"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
