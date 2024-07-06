package model

import "time"

type Todo struct {
	Id        int       `pg:",pk" json:"id"`
	Title     string    `pg:"title" json:"title" binding:"required,min=3,max=100"`
	Completed bool      `pg:"completed" json:"completed"`
	AuthorId  int       `pg:"author_id" json:"author_id"`
	Author    Author    `pg:"rel:has-one" json:"author"`
	CreatedAt time.Time `pg:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `pg:"updated_at" json:"updated_at,omitempty"`
}
