package product_test

import (
	"catalog/bussiness"
	"catalog/bussiness/product"
	productMock "catalog/bussiness/product/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	id          = 1
	category_id = 2
	name        = "name"
	price       = 12000
	description = "description"
	image       = "image url"
	is_active   = true
)

var (
	productService product.Service
	productRepo    productMock.Repository

	productData   product.Product
	createProduct product.ProductField
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	productData = product.NewProduct(
		category_id,
		name,
		price,
		description,
		image,
		is_active,
		time.Now(),
	)

	createProduct = product.ProductField{
		Category_id: category_id,
		Name:        name,
		Price:       price,
		Description: description,
		Image:       image,
		Is_active:   is_active,
	}

	productService = product.NewService(&productRepo)

}

func TestCreateProduct(t *testing.T) {
	t.Run("Expect create user success", func(t *testing.T) {
		productRepo.On("CreateProduct", mock.AnythingOfType("product.Product"), mock.AnythingOfType("string")).Return(nil).Once()

		err := productService.CreateProduct(createProduct)

		assert.Nil(t, err)

	})

	t.Run("Expect create product not found", func(t *testing.T) {
		productRepo.On("CreateProduct", mock.AnythingOfType("product.Product"), mock.AnythingOfType("string")).Return(bussiness.ErrInternalServerError).Once()

		err := productService.CreateProduct(createProduct)

		assert.NotNil(t, err)

		assert.Equal(t, err, bussiness.ErrInternalServerError)
	})
}

// func TestGetProductById(t *testing.T) {
// 	t.Run("Expect found the product", func(t *testing.T) {
// 		productRepo.On("GetProductById", mock.AnythingOfType("int")).Return(&productData, nil).Once()

// 		product, err := productService.GetProductById(id)

// 		assert.Nil(t, err)

// 		assert.NotNil(t, product)

// 		assert.Equal(t, id, product.Id)
// 		assert.Equal(t, category_id, product.Category_id)
// 		assert.Equal(t, name, product.Name)
// 		assert.Equal(t, price, product.Price)
// 		assert.Equal(t, description, product.Description)
// 		assert.Equal(t, image, product.Image)
// 		assert.Equal(t, is_active, product.Is_active)

// 	})

// 	t.Run("Expect product not found", func(t *testing.T) {
// 		productRepo.On("GetProductById", mock.AnythingOfType("int")).Return(nil, bussiness.ErrNotFound).Once()

// 		product, err := productService.GetProductById(id)

// 		assert.NotNil(t, err)

// 		assert.Nil(t, product)

// 		assert.Equal(t, err, bussiness.ErrNotFound)
// 	})
// }

func TestUpdateProduct(t *testing.T) {
	t.Run("Expect update product success", func(t *testing.T) {
		productRepo.On("GetProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()
		productRepo.On("UpdateProduct", mock.AnythingOfType("product.Product"), mock.AnythingOfType("int")).Return(nil).Once()

		err := productService.UpdateProduct(createProduct, id)

		assert.Nil(t, err)

	})

	t.Run("Expect update product failed", func(t *testing.T) {
		productRepo.On("GetProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()
		productRepo.On("UpdateProduct", mock.AnythingOfType("product.Product"), mock.AnythingOfType("int")).Return(bussiness.ErrInternalServerError).Once()

		err := productService.UpdateProduct(createProduct, id)

		assert.NotNil(t, err)

		assert.Equal(t, err, bussiness.ErrInternalServerError)
	})
}
