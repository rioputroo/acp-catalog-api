package category_test

import (
	"catalog/bussiness"
	"catalog/bussiness/category"
	catMock "catalog/bussiness/category/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	id        = 1
	name      = "name"
	is_active = true
)

var (
	catService category.Service
	catRepo    catMock.Repository

	catAllData []category.Category = make([]category.Category, 0)

	catData   category.Category
	createCat category.CategoryField
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	catData = category.NewCategory(
		name,
		is_active,
		time.Now(),
	)

	createCat = category.CategoryField{
		Name:     name,
		IsActive: is_active,
	}

	catService = category.NewService(&catRepo)

}

func TestGetCategotyById(t *testing.T) {
	t.Run("expect found the category", func(t *testing.T) {
		catRepo.On("GetCategoryById", mock.AnythingOfType("int")).Return(&catData, nil).Once()

		category, _ := catService.GetCategoryById(id)

		//assert.Nil(t, err)

		assert.NotNil(t, category)
	})

	t.Run("expect category not found", func(t *testing.T) {
		catRepo.On("GetCategoryById", mock.AnythingOfType("int")).Return(nil, bussiness.ErrNotFound).Once()

		category, err := catService.GetCategoryById(id)

		assert.NotNil(t, err)

		assert.Nil(t, category)

		assert.Equal(t, err, bussiness.ErrNotFound)
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("expect create category success", func(t *testing.T) {
		catRepo.On("CreateCategory", mock.AnythingOfType("category.Category"), mock.AnythingOfType("string")).Return(nil).Once()

		err := catService.CreateCategory(createCat)

		assert.Nil(t, err)

	})

	t.Run("expect create product not found", func(t *testing.T) {
		catRepo.On("CreateCategory", mock.AnythingOfType("category.Category"), mock.AnythingOfType("string")).Return(bussiness.ErrInternalServerError).Once()

		err := catService.CreateCategory(createCat)

		assert.NotNil(t, err)
		//assert.Equal(t, err, bussiness.ErrInternalServerError)
	})
}

func TestGetAllCategory(t *testing.T) {
	t.Run("expect found all category", func(t *testing.T) {
		catRepo.On("GetAllCategory", mock.Anything).Return(catAllData, nil).Once()

		product, _ := catService.GetAllCategory()

		//assert.Nil(t, err)
		assert.NotNil(t, product)

	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("success delete category", func(t *testing.T) {
		catRepo.On("DeleteCategory", mock.AnythingOfType("int")).Return(nil).Once()

		err := catService.DeleteCategory(id)
		assert.Nil(t, err)

	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("expect update category success", func(t *testing.T) {
		catRepo.On("GetCategoryById", mock.AnythingOfType("string")).Return(&catData, nil).Once()
		catRepo.On("UpdateCategory", mock.AnythingOfType("category.Category"), mock.AnythingOfType("int")).Return(nil).Once()

		err := catService.UpdateCategory(createCat, id)
		assert.Nil(t, err)
	})

	t.Run("expect update category failed", func(t *testing.T) {
		catRepo.On("GetCategoryById", mock.AnythingOfType("string")).Return(&catData, nil).Once()
		catRepo.On("UpdateCategory", mock.AnythingOfType("category.Category"), mock.AnythingOfType("int")).Return(bussiness.ErrInternalServerError).Once()

		err := catService.UpdateCategory(createCat, id)

		//assert.NotNil(t, err)

		assert.Equal(t, err, bussiness.ErrInternalServerError)
	})
}
