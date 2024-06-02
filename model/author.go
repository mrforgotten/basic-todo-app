package model

import "time"

type Author struct {
	tableName struct{}  `pg:",discard_unknown_columns"`
	Id        int       `pg:"id,pk" json:"id"`
	Name      string    `pg:"name,unique" json:"name" binding:"required"`
	CreatedAt time.Time `pg:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `pg:"updated_at" json:"updated_at,omitempty"`
}

type AuthorUpdate struct {
	Name      string
	UpdatedAt time.Time
}
