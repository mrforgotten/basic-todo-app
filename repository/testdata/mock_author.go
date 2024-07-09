package mock_repository_test

import (
	"basic-rest-api-orm/model"

	"github.com/stretchr/testify/mock"
)

type MockAuthorRepository struct {
	mock.Mock
}

func (r *MockAuthorRepository) Create(author *model.Author) error {
	args := r.Called(author)
	return args.Error(0)
}

func (r *MockAuthorRepository) GetById(id int) (*model.Author, error) {
	args := r.Called(id)
	return args.Get(0).(*model.Author), args.Error(0)
}

func (r *MockAuthorRepository) GetAll() ([]model.Author, error) {
	args := r.Called()
	return args.Get(0).([]model.Author), args.Error(0)
}

func (r *MockAuthorRepository) Update(update *model.Author) error {
	args := r.Called(update)
	return args.Error(0)
}

func (r *MockAuthorRepository) Delete(id int) error {
	args := r.Called(id)
	return args.Error(0)
}
