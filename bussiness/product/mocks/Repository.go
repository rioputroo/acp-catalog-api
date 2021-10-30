package mocks

import (
	"catalog/bussiness/product"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (t *Repository) GetAllProducts(categoryId int) ([]product.Product, error) {
	ret := t.Called()
	var tProductSuccess []product.Product

	if rf, ok := ret.Get(0).(func() []product.Product); ok {
		tProductSuccess = rf()
	} else {
		if ret.Get(0) != nil {
			tProductSuccess = ret.Get(0).([]product.Product)
		}
	}

	// var tProductError error

	// if rf, ok := ret.Get(1).(func() error); ok {
	// 	tProductError = rf()
	// } else {
	// 	tProductError = ret.Error(1)
	// }

	return tProductSuccess, nil
}

func (t *Repository) GetProductById(id int) (*product.Product, error) {
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

func (t *Repository) CreateProduct(_t product.Product) error {
	ret := t.Called(_t)

	var tProductError error
	if rf, ok := ret.Get(0).(func(product.Product) error); ok {
		tProductError = rf(_t)
	} else {
		tProductError = ret.Error(0)
	}
	return tProductError
}

func (t *Repository) UpdateProduct(_t product.Product, id int) error {
	ret := t.Called(_t, id)

	var tProductError error
	if rf, ok := ret.Get(0).(func(product.Product, int) error); ok {
		tProductError = rf(_t, id)
	} else {
		tProductError = ret.Error(0)
	}

	return tProductError

}

func (t *Repository) DeleteProduct(id int) error {
	ret := t.Called(id)
	var tProductError error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		tProductError = rf(id)
	} else {
		tProductError = ret.Error(0)
	}

	return tProductError
}
