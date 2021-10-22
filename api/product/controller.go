package product

import (
	"catalog/api/product/request"
	"catalog/bussiness/product"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) GetProductsById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("productId"))

	product, err := controller.service.GetProductById(id)
	fmt.Println("controller-id", id)
	fmt.Println("controller", product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "successfully fetch product",
		"products": product,
	})

}

func (controller *Controller) GetAllProducts(c echo.Context) error {

	products, err := controller.service.GetAllProducts()

	fmt.Println("Controller", &products)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "successfully fetch list of products",
		"products": products,
	})
}

func (controller *Controller) CreateProduct(c echo.Context) error {
	insertProduct := new(request.InsertReqProduct)

	fmt.Println("controller", insertProduct)

	if err := c.Bind(insertProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	err := controller.service.CreateProduct(*insertProduct.InsertNewProduct())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Product Created",
	})
}

func (controller *Controller) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("productId"))
	err := controller.service.DeleteProduct(id)

	fmt.Println("Controller", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"message": "successfully delete product",
	})

}

func (controller *Controller) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("productId"))

	product := new(request.UpdateReqProduct)

	fmt.Println(product)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	controller.service.UpdateProduct(*product.UpdateExistProduct(), id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully update product",
	})
}
