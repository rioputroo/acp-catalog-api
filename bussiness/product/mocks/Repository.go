package mocks

import (
	"catalog/bussiness/product"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (t *Repository) GetProducts() ([]product.Product, error) {
	ret := t.Called()
	var tProductSuccess []product.Product

	if rf, ok := ret.Get(0).(func() []product.Product); ok {
		tProductSuccess = rf()
	} else {
		if ret.Get(0) != nil {
			tProductSuccess = ret.Get(0).([]product.Product)
		}
	}

	var tProductError error

	if rf, ok := ret.Get(1).(func() error); ok {
		tProductError = rf()
	} else {
		tProductError = ret.Error(1)
	}

	return tProductSuccess, tProductError
}

func (_m *Repository) GetProductById(id int) (*product.Product, error) {
	ret := _m.Called(id)

	var tProductSuccess *product.Product
	if rf, ok := ret.Get(0).(func(int) *product.Product); ok {
		tProductSuccess = rf(id)
	} else {
		if ret.Get(0) != nil {
			tProductSuccess = ret.Get(0).(*product.Product)
		}
	}

	var tProductError error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		tProductError = rf(id)
	} else {
		tProductError = ret.Error(1)
	}

	return tProductSuccess, tProductError
}
