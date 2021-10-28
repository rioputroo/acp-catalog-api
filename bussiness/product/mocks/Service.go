package mocks

import (
	"catalog/bussiness/product"

	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (t *Service) GetAllProducts() ([]product.Product, error) {
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

func (t *Service) GetProductById(id int) (*product.Product, error) {
	ret := t.Called(id)

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

func (t *Service) CreateProduct(productField product.ProductField) error {
	ret := t.Called(productField)

	var tProductError error
	if rf, ok := ret.Get(0).(func(product.ProductField) error); ok {
		tProductError = rf(productField)
	} else {
		tProductError = ret.Error(0)
	}

	return tProductError
}

func (t *Service) DeleteProduct(id int) error {
	ret := t.Called(id)

	var tProductError error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		tProductError = rf(id)
	} else {
		tProductError = ret.Error(0)
	}
	return tProductError
}

func (t *Service) UpdateProduct(productField product.ProductField, id int) error {
	ret := t.Called(productField, id)

	var tProductError error
	if rf, ok := ret.Get(0).(func(product.ProductField, int) error); ok {
		tProductError = rf(productField, id)
	} else {
		tProductError = ret.Error(0)
	}

	return tProductError
}
