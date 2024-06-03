package repository

import (
	"basic-rest-api-orm/model"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

var logger = log.New(os.Stderr, "AuthorRepository: ", log.Ldate|log.Ltime|log.Lshortfile)

type AuthorRepository struct {
	db *pg.DB

	logger *log.Logger
}

func ProvideAuthorRepository(db *pg.DB) AuthorRepository {
	return AuthorRepository{db: db, logger: logger}
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

	defer func() {
		if err != nil {
			log.Println("Error while creating author: ", err)
			tx.Rollback()
		}
		tx.Close()
	}()

	// Check if author name is exist
	isExist, err := a.AuthorNameIsExist(tx, author.Name)

	if err != nil {
		return err
	} else if isExist {
		return errors.New("duplicate unique value for column name")
	}

	// Insert author
	if _, err := tx.Model(author).Returning("*").Insert(author); err != nil {
		return err
	}
	log.Println("Author created")
	tx.Commit()

	return nil
}

func (a *AuthorRepository) Update(prev *model.Author, update *model.Author) error {

	tx, err := a.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Close()
	}()

	// Check if author is exist
	if err := tx.Model(prev).WherePK().Select(); err != nil {
		return err
	}

	// update author
	update.Id = prev.Id
	update.UpdatedAt = time.Now()

	_, err = tx.Model(update).WherePK().Update()
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
