package repository

import (
	"basic-rest-api-orm/model"

	"github.com/go-pg/pg/v10"
)

type AuthorRepository struct {
	db *pg.DB
}

func ProvideAuthorRepository(db *pg.DB) AuthorRepository {
	return AuthorRepository{db: db}
}

func (a AuthorRepository) GetAll() ([]model.Author, error) {
	var authors []model.Author
	err := a.db.Model(&authors).Select()
	if err != nil {
		return nil, err
	}
	return authors, nil
}
