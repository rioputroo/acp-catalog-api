package product

import (
	"catalog/api/product/request"
	"catalog/api/response"
	"catalog/bussiness/product"
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
	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}

	responseProduct := request.GetProductById(*product)
	return c.JSON(response.NewSuccessResponse(responseProduct))
}

func (controller *Controller) GetAllProducts(c echo.Context) error {

	products, err := controller.service.GetAllProducts()
	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}

	responseProduct := request.GetProducts(products)
	return c.JSON(response.NewSuccessResponse(responseProduct))
}

func (controller *Controller) CreateProduct(c echo.Context) error {
	insertProduct := new(request.InsertReqProduct)

	if err := c.Bind(insertProduct); err != nil {
		return c.JSON(response.NewBadRequestResponse())
	}

	err := controller.service.CreateProduct(*insertProduct.InsertNewProduct())

	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}

	return c.JSON(response.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("productId"))
	err := controller.service.DeleteProduct(id)

	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}

	return c.JSON(response.NewSuccessResponseNoContent())

}

func (controller *Controller) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("productId"))
	product := new(request.UpdateReqProduct)

	if err := c.Bind(product); err != nil {
		return c.JSON(response.NewBadRequestResponse())
	}

	err := controller.service.UpdateProduct(*product.UpdateExistProduct(), id)

	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}
	return c.JSON(response.NewSuccessResponseWithoutData())
}
