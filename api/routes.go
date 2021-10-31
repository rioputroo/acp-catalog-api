package api

import (
	"catalog/api/category"
	"catalog/api/product"

	"github.com/labstack/echo"
)

func HandlerApi(e *echo.Echo, productController *product.Controller, categoryController *category.Controller) {
	//e := echo.New()

	//Product
	e.GET("/catalog/products", productController.GetAllProducts)
	e.GET("/catalog/product/:productId", productController.GetProductsById)
	e.GET("/catalog/filterproduct", productController.GetProductsByCategoryId)
	e.POST("/catalog/product", productController.CreateProduct)
	e.PUT("/catalog/product/:productId", productController.UpdateProduct)
	e.DELETE("/catalog/product/:productId", productController.DeleteProduct)

	//Category
	e.GET("/catalog/category", categoryController.GetAllcategory)
	e.GET("/catalog/category/:categoryId", categoryController.GetCategoryById)
	e.POST("/catalog/category", categoryController.CreateCategory)
	e.PUT("/catalog/category/:categoryId", categoryController.UpdateCategory)
	e.DELETE("/catalog/category/:categoryId", categoryController.DeleteCategory)

	//return e
}
