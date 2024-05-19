package authorservice

import (
	"basic-rest-api-orm/model"
	"basic-rest-api-orm/repository"
)

type AuthorService struct {
	authorRepo repository.AuthorRepository
}

func ProvideAuthorService(r repository.AuthorRepository) AuthorService {
	return AuthorService{authorRepo: r}
}

func (a AuthorService) GetAll() ([]model.Author, error) {
	var authors []model.Author
	authors, err := a.authorRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return authors, nil
}
