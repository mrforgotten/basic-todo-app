package model

type Author struct {
	tableName struct{} `pg:"author"`
	Id        int      `json:"id"`
	Name      string   `json:"name"`
}
