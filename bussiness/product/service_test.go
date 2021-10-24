package product_test

import (
	"catalog/bussiness"
	"catalog/bussiness/product"
	productMock "catalog/bussiness/product/mocks"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id          = 1
	category_id = "category_id"
	name        = "name"
	price       = "price"
	description = "description"
	image       = "image url"
	is_active   = "is_active"
	// created_At  = "time"
	// updated_At  = "time"
)

var (
	productService product.Service
	productRepo    productMock.Repository

	productData product.Product
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetProductById(t *testing.T) {
	t.Run("Expect found the product", func(t *testing.T) {
		productRepo.On("GetProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()

		product, err := productService.GetProductById(1)
		fmt.Println(product)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, id, product.Id)
		assert.Equal(t, category_id, product.Category_id)
		assert.Equal(t, name, product.Name)
		assert.Equal(t, price, product.Price)
		assert.Equal(t, price, product.Description)
		assert.Equal(t, price, product.Image)
		// assert.Equal(t, price, product.Created_At)
		// assert.Equal(t, price, product.Updated_At)

	})

	t.Run("Expect product not found", func(t *testing.T) {
		productRepo.On("GetProductById", mock.AnythingOfType("string")).Return(nil).Once()

		product, err := productService.GetProductById(id)

		assert.NotNil(t, err)

		assert.Nil(t, product)

		assert.Equal(t, err, bussiness.ErrNotFound)
	})
}
