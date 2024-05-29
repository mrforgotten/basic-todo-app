package model

type Author struct {
	Id        int    `pg:"id,pk" json:"id"`
	Name      string `pg:"name,unique" json:"name" binding:"required"`
	CreatedAt string `pg:"created_at" json:"created_at"`
}

func (t Author) TableName() string {
	return "author"
}
