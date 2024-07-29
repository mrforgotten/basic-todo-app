package authorservice

import (
	"basic-rest-api-orm/model"
	"basic-rest-api-orm/repository"
	"log"
)

type AuthorService interface {
	Create(author *model.Author) error
	GetAll() ([]model.Author, error)
	GetByID(id int) (model.Author, error)
	Update(id int, update *model.Author) error
}

type AuthorServiceImpl struct {
	authorRepo repository.AuthorRepository
}

func NewProvideAuthorService(r repository.AuthorRepository) AuthorService {
	return &AuthorServiceImpl{authorRepo: r}
}

func (s *AuthorServiceImpl) GetAll() ([]model.Author, error) {
	var authors []model.Author
	authors, err := s.authorRepo.GetAll()
	if err != nil {
		log.Printf("Error while getting all authors: %v", err)
		return nil, err
	}
	return authors, nil
}

func (s *AuthorServiceImpl) GetByID(id int) (model.Author, error) {
	author, err := s.authorRepo.GetById(id)
	if err != nil {
		return model.Author{}, err
	}
	return *author, nil
}

func (s *AuthorServiceImpl) Create(author *model.Author) error {

	err := s.authorRepo.Create(author)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthorServiceImpl) Update(id int, update *model.Author) error {

	update.Id = id

	err := s.authorRepo.Update(update)
	if err != nil {
		return err
	}

	return nil
}
