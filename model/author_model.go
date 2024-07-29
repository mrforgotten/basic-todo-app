package model

import "time"

type Author struct {
	Id        int       `pg:"id,pk"`
	Name      string    `pg:"name,unique" binding:"required"`
	CreatedAt time.Time `pg:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `pg:"updated_at" json:"updated_at,omitempty"`
	TodoList  []*Todo   `pg:"rel:has-many" json:"omitempty"`
}
