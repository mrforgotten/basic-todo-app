package todoservice

import (
	"basic-rest-api-orm/model"
	"basic-rest-api-orm/repository"
	"log"
)

type TodoService interface {
	Create(todo *model.Todo) error
	GetAll() ([]model.Todo, error)
	GetByID(id int) (model.Todo, error)
	Update(id int, update *model.Todo) error
}

type TodoServiceImpl struct {
	todoRepo repository.TodoRepository
}

func ProvideTodoService(r repository.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepo: r}
}

func (s *TodoServiceImpl) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	todos, err := s.todoRepo.GetAll()
	if err != nil {
		log.Printf("Error while getting all todos: %v", err)
		return nil, err
	}
	return todos, nil
}

func (s *TodoServiceImpl) GetByID(id int) (model.Todo, error) {
	todo, err := s.todoRepo.GetById(id)
	if err != nil {
		return model.Todo{}, err
	}
	return *todo, nil
}

func (s *TodoServiceImpl) Create(todo *model.Todo) error {

	err := s.todoRepo.Create(todo)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoServiceImpl) Update(id int, update *model.Todo) error {

	update.Id = id

	err := s.todoRepo.Update(update)
	if err != nil {
		return err
	}

	return nil
}
