package repository

import (
	"basic-rest-api-orm/model"

	"github.com/go-pg/pg/v10"
)

type TodoRepository struct {
	db *pg.DB
}

func ProvideTodoRepository(db *pg.DB) TodoRepository {
	return TodoRepository{db: db}
}

func (t *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := t.db.Model(&todos).Select()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodoRepository) GetById(id int) (model.Todo, error) {
	var todo model.Todo
	err := t.db.Model(&todo).Where("id = ?", id).Select()
	if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (t *TodoRepository) Create(todo model.Todo) (model.Todo, error) {
	_, err := t.db.Model(&todo).Insert()
	if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (t *TodoRepository) Update(prev, update model.Todo) (model.Todo, error) {
	_, err := t.db.Model(&prev).WherePK().Update()
	if err != nil {
		return model.Todo{}, err
	}
	return update, nil
}
