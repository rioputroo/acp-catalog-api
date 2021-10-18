package handler

import (
	"catalog/handler/api"

	"github.com/labstack/echo"
)

func HandlerApi() *echo.Echo {
	e := echo.New()

	e.GET("/catalog/products", api.GetProducts)
	e.GET("/catalog/products/:productId", api.GetProductsById)
	e.POST("/catalog/products", api.CreateProduct)
	e.PUT("/catalog/products/:productId", api.GetProductsById)
	e.DELETE("/catalog/products/:productId", api.DeleteProduct)

	return e
}
