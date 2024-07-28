package repository

import (
	"basic-rest-api-orm/model"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	GetAll() ([]model.Todo, error)
	GetById(id int) (*model.Todo, error)
	Update(update *model.Todo) error
	Delete(id int) error
}

type TodoRepositoryImpl struct {
	db *pg.DB

	logger *log.Logger
}

func NewProvideTodoRepository(db *pg.DB) TodoRepository {
	var logger = log.New(os.Stderr, "TodoRepository: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &TodoRepositoryImpl{db: db, logger: logger}
}

func (r *TodoRepositoryImpl) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Model(&todos).Select()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepositoryImpl) GetById(id int) (*model.Todo, error) {
	var todo *model.Todo
	err := r.db.Model(&todo).Where("id = ?", id).Select()
	if err != nil {
		return &model.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) Create(todo *model.Todo) error {
	// make transaction for create todo
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			log.Println("Error while creating todo: ", err)
			tx.Rollback()
		}
		tx.Close()
	}()

	// Check if todo name is exist

	if err != nil {
		return err
	}

	// Insert todo
	if _, err := tx.Model(todo).Returning("*").Insert(todo); err != nil {
		return err
	}
	log.Println("Todo created")
	tx.Commit()

	return nil
}

func (r *TodoRepositoryImpl) Update(update *model.Todo) error {

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

	var prev *model.Todo = new(model.Todo)
	// Check if todo is exist
	if err := tx.Model(prev).Where("id = ?", update.Id).Select(); err != nil {
		return err
	}

	update.UpdatedAt = time.Now().UTC()

	_, err = tx.Model(update).Column("title", "updated_at").WherePK().Update(update)
	if err != nil {
		return err
	}

	tx.Commit()

	r.db.Model(update).WherePK().Select()

	return nil
}

func (r *TodoRepositoryImpl) Delete(id int) error {
	todo, err := r.GetById(id)
	if err != nil {
		return err
	}

	_, err = r.db.Model(&todo).Delete()

	if err != nil {
		return err
	}

	return err
}
