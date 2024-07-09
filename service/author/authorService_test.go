package authorservice_test

import (
	"basic-rest-api-orm/model"
	mock_repository_test "basic-rest-api-orm/repository/testdata"
	authorservice "basic-rest-api-orm/service/author"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeNow1 time.Time = time.Now().UTC()
var timeNow2 time.Time = time.Now().UTC()
var timeNow3 time.Time = time.Now().UTC()

var AuthorMockData []model.Author = []model.Author{
	{Id: 1, Name: "Test Subject 1", CreatedAt: timeNow1, UpdatedAt: timeNow1},
	{Id: 2, Name: "Test Subject 2", CreatedAt: timeNow2, UpdatedAt: timeNow2},
	{Id: 3, Name: "Test Subject 3", CreatedAt: timeNow3, UpdatedAt: timeNow3},
}

func PrepareService() (*mock_repository_test.MockAuthorRepository, authorservice.AuthorService) {
	mockRepo := new(mock_repository_test.MockAuthorRepository)
	service := authorservice.NewProvideAuthorService(mockRepo)

	return mockRepo, service
}

func TestAuthorCreate(t *testing.T) {
	timeNow := time.Now().UTC()
	AuthorUser4 := &model.Author{
		Id: 1, Name: "Test Create",
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	mockRepo, service := PrepareService()
	mockRepo.On("Create", AuthorUser4).Return(nil)

	err := service.Create(AuthorUser4)

	assert.Equal(t, nil, err, "Should create author without error")
	assert.Equal(t, *AuthorUser4, model.Author{
		Id: 1, Name: "Test Create",
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}, "Should expected result from create")
}
