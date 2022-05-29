package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var serviceCategory = CategoryService{Repository: categoryRepository}

func TestCateoryService_GetNotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := serviceCategory.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{Id: "2", Name: "category 1"}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := serviceCategory.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

