package product_test

import (
	"catalog/bussiness"
	"catalog/bussiness/product"
	productMock "catalog/bussiness/product/mocks"
	"os"
	"testing"

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

	productsData []product.Product = make([]product.Product, 0)

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
		is_active,
	)

	createProduct = product.ProductField{
		CategoryId:  category_id,
		Name:        name,
		Price:       price,
		Description: description,
		IsActive:    is_active,
	}

	productService = product.NewService(&productRepo, nil)

}

func TestCreateProduct(t *testing.T) {
	t.Run("Expect create product success", func(t *testing.T) {
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

func TestGetProductById(t *testing.T) {
	t.Run("Expect found the product", func(t *testing.T) {
		productRepo.On("GetProductById", mock.AnythingOfType("int")).Return(&productData, nil).Once()

		product, err := productService.GetProductById(id)

		assert.Nil(t, err)
		assert.NotNil(t, product)
	})

	t.Run("Expect product not found", func(t *testing.T) {
		productRepo.On("GetProductById", mock.AnythingOfType("int")).Return(nil, bussiness.ErrNotFound).Once()

		product, err := productService.GetProductById(id)

		assert.NotNil(t, err)
		assert.Nil(t, product)
		assert.Equal(t, err, bussiness.ErrNotFound)
	})
}

func TestGetAllProducts(t *testing.T) {
	t.Run("Expect found all product", func(t *testing.T) {
		productRepo.On("GetAllProducts", mock.Anything).Return(productsData, nil).Once()

		product, err := productService.GetAllProducts(category_id)

		assert.Nil(t, err)
		assert.NotNil(t, product)

	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Success delete product", func(t *testing.T) {
		productRepo.On("DeleteProduct", mock.AnythingOfType("int")).Return(nil).Once()

		err := productService.DeleteProduct(id)

		assert.Nil(t, err)

	})
}

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
