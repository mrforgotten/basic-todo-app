package repository

import (
	"basic-rest-api-orm/model"
	"errors"
	"log"

	"github.com/go-pg/pg/v10"
)

type AuthorRepository struct {
	db *pg.DB
}

func ProvideAuthorRepository(db *pg.DB) AuthorRepository {
	return AuthorRepository{db: db}
}

func (a *AuthorRepository) GetAll() ([]model.Author, error) {
	var authors []model.Author
	err := a.db.Model(&authors).Select()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *AuthorRepository) AuthorNameIsExist(tx *pg.Tx, name string) (bool, error) {
	author := new(model.Author)
	err := tx.Model(author).Where("name =?", name).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return false, nil
		}
		log.Println("Error while checking author name is exist: ", err)
		return false, err
	}

	return true, nil
}

func (a *AuthorRepository) GetById(id int) (*model.Author, error) {
	var author *model.Author
	err := a.db.Model(&author).Where("id = ?", id).Select()
	if err != nil {
		return &model.Author{}, err
	}
	return author, nil
}

func (a *AuthorRepository) Create(author *model.Author) error {
	// make transaction for create author
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Close()

	// Check if author name is exist
	isExist, err := a.AuthorNameIsExist(tx, author.Name)

	if err != nil {
		return err
	} else if isExist {
		return errors.New("duplicate unique value for column name")
	}

	// Insert author
	if _, err := tx.Model(author).Returning("*").Insert(author); err != nil {
		tx.Rollback()
		return err
	}
	log.Println("Author created")
	tx.Commit()

	return nil
}

func (a *AuthorRepository) Update(prev, update *model.Author) error {
	_, err := a.db.Model(&prev).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
