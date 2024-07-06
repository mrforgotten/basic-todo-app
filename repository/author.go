package repository

import (
	"basic-rest-api-orm/model"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

type AuthorRepository interface {
	Create(author *model.Author) error
	GetById(id int) (*model.Author, error)
	GetAll() ([]model.Author, error)
	Update(update *model.Author) error
	Delete(id int) error
}

type AuthorRepositoryImpl struct {
	db *pg.DB

	logger *log.Logger
}

func ProvideAuthorRepository(db *pg.DB) AuthorRepository {
	var logger = log.New(os.Stderr, "AuthorRepository: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &AuthorRepositoryImpl{db: db, logger: logger}
}

func (r *AuthorRepositoryImpl) GetAll() ([]model.Author, error) {
	var authors []model.Author
	err := r.db.Model(&authors).Select()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepositoryImpl) AuthorNameIsExist(tx *pg.Tx, name string) (bool, error) {
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

func (r *AuthorRepositoryImpl) GetById(id int) (*model.Author, error) {
	var author *model.Author
	err := r.db.Model(&author).Where("id = ?", id).Select()
	if err != nil {
		return &model.Author{}, err
	}
	return author, nil
}

func (r *AuthorRepositoryImpl) Create(author *model.Author) error {
	// make transaction for create author
	tx, err := r.db.Begin()
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
	isExist, err := r.AuthorNameIsExist(tx, author.Name)

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

func (r *AuthorRepositoryImpl) Update(update *model.Author) error {

	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Close()
	}()

	var prev *model.Author = new(model.Author)
	// Check if author is exist
	if err := tx.Model(prev).Where("id = ?", update.Id).Select(); err != nil {
		return err
	}

	update.UpdatedAt = time.Now().UTC()

	_, err = tx.Model(update).Column("name", "updated_at").WherePK().Update(update)
	if err != nil {
		return err
	}

	tx.Commit()

	r.db.Model(update).WherePK().Select()

	return nil
}

func (r *AuthorRepositoryImpl) Delete(id int) error {
	author, err := r.GetById(id)
	if err != nil {
		return err
	}

	_, err = r.db.Model(&author).Delete()

	if err != nil {
		return err
	}

	return err
}
