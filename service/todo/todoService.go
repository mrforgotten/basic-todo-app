package todoservice

import (
	"basic-rest-api-orm/model"
	"basic-rest-api-orm/repository"
	"log"
)

type TodoService struct {
	todoRepo repository.TodoRepository
}

func ProvideTodoService(r repository.TodoRepository) TodoService {
	return TodoService{todoRepo: r}
}

func (s *TodoService) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	todos, err := s.todoRepo.GetAll()
	if err != nil {
		log.Printf("Error while getting all todos: %v", err)
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetByID(id int) (model.Todo, error) {
	todo, err := s.todoRepo.GetById(id)
	if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Create(todo model.Todo) (model.Todo, error) {
	if todo.Title == "" {

	}
	data, err := s.todoRepo.Create(todo)
	if err != nil {
		return model.Todo{}, err
	}
	return data, nil
}

func (s *TodoService) Update(todo model.Todo) (model.Todo, error) {
	data, err := s.todoRepo.Create(todo)
	if err != nil {
		return model.Todo{}, err
	}
	return data, nil
}
