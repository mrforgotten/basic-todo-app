package authorservice

import (
	"basic-rest-api-orm/model"
	"basic-rest-api-orm/repository"
	"log"
)

type AuthorService struct {
	authorRepo repository.AuthorRepository
}

func ProvideAuthorService(r repository.AuthorRepository) AuthorService {
	return AuthorService{authorRepo: r}
}

func (s *AuthorService) GetAll() ([]model.Author, error) {
	var authors []model.Author
	authors, err := s.authorRepo.GetAll()
	if err != nil {
		log.Printf("Error while getting all authors: %v", err)
		return nil, err
	}
	return authors, nil
}

func (s *AuthorService) GetByID(id int) (model.Author, error) {
	author, err := s.authorRepo.GetById(id)
	if err != nil {
		return model.Author{}, err
	}
	return *author, nil
}

func (s *AuthorService) Create(author *model.Author) error {

	err := s.authorRepo.Create(author)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthorService) Update(prev, update *model.Author) error {

	err := s.authorRepo.Update(prev, update)
	if err != nil {
		return err
	}

	return nil
}
