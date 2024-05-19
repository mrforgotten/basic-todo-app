package model

type Todo struct {
	Id        int    `pg:"id" json:"id"`
	Title     string `pg:"title" json:"title"`
	Completed bool   `pg:"completed" json:"completed"`
	AuthorId  int    `pg:"author_id" json:"author_id"`
	Author    Author `pg:"rel:has-one" json:"author"`
}
