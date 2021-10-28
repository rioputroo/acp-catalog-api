package mocks

import (
	"catalog/bussiness/category"

	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (t *Service) GetAllCategory() ([]category.Category, error) {
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

func (t *Service) GetCategoryById(id int) (*category.Category, error) {
	ret := t.Called(id)

	var tCatSuccess *category.Category
	if rf, ok := ret.Get(0).(func(int) *category.Category); ok {
		tCatSuccess = rf(id)
	} else {
		if ret.Get(0) != nil {
			tCatSuccess = ret.Get(0).(*category.Category)
		}
	}

	var tCatError error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		tCatError = rf(id)
	} else {
		tCatError = ret.Error(1)
	}
	return tCatSuccess, tCatError
}

func (t *Service) CreateCategory(categoryField category.CategoryField) error {
	ret := t.Called(categoryField)

	var tCatError error
	if rf, ok := ret.Get(0).(func(category.CategoryField) error); ok {
		tCatError = rf(categoryField)
	} else {
		tCatError = ret.Error(0)
	}

	return tCatError
}

func (t *Service) DeleteCategory(id int) error {
	ret := t.Called(id)

	var tCatError error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		tCatError = rf(id)
	} else {
		tCatError = ret.Error(0)
	}
	return tCatError
}

func (t *Service) UpdateCategory(categoryField category.CategoryField, id int) error {
	ret := t.Called(categoryField, id)

	var tCatError error
	if rf, ok := ret.Get(0).(func(category.CategoryField, int) error); ok {
		tCatError = rf(categoryField, id)
	} else {
		tCatError = ret.Error(0)
	}

	return tCatError
}
