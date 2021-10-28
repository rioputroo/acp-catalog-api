package mocks

import (
	"catalog/bussiness/category"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (t *Repository) GetAllCategory() ([]category.Category, error) {
	ret := t.Called()
	var tCatSuccess []category.Category

	if rf, ok := ret.Get(0).(func() []category.Category); ok {
		tCatSuccess = rf()
	} else {
		if ret.Get(0) != nil {
			tCatSuccess = ret.Get(0).([]category.Category)
		}
	}

	// var tCatError error

	// if rf, ok := ret.Get(1).(func() error); ok {
	// 	tCatError = rf()
	// } else {
	// 	tCatError = ret.Error(1)
	// }

	return tCatSuccess, nil
}

func (t *Repository) GetCategoryById(id int) (*category.Category, error) {
	ret := t.Called(id)

	var tCatCategory *category.Category
	if rf, ok := ret.Get(0).(func(int) *category.Category); ok {
		tCatCategory = rf(id)
	} else {
		if ret.Get(0) != nil {
			tCatCategory = ret.Get(0).(*category.Category)
		}
	}

	var tCatError error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		tCatError = rf(id)
	} else {
		tCatError = ret.Error(1)
	}

	return tCatCategory, tCatError
}

func (t *Repository) CreateCategory(_t category.Category) error {
	ret := t.Called(_t)

	var tCatError error
	if rf, ok := ret.Get(0).(func(category.Category) error); ok {
		tCatError = rf(_t)
	} else {
		tCatError = ret.Error(0)
	}
	return tCatError
}

func (t *Repository) UpdateCategory(_t category.Category, id int) error {
	ret := t.Called(_t, id)

	var tCatError error
	if rf, ok := ret.Get(0).(func(category.Category, int) error); ok {
		tCatError = rf(_t, id)
	} else {
		tCatError = ret.Error(0)
	}

	return tCatError

}

func (t *Repository) DeleteCategory(id int) error {
	ret := t.Called(id)
	var tCatError error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		tCatError = rf(id)
	} else {
		tCatError = ret.Error(0)
	}

	return tCatError
}
