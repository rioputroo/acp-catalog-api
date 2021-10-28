package category

import (
	"catalog/api/category/request"
	"catalog/api/response"
	"catalog/bussiness/category"
	"strconv"

	"github.com/labstack/echo"
)

type Controller struct {
	service category.Service
}

func NewController(service category.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) GetCategoryById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("categoryId"))

	category, err := controller.service.GetCategoryById(id)

	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}

	resCategory := request.GetCategoryById(*category)

	return c.JSON(response.NewSuccessResponse(resCategory))
}

func (controller *Controller) GetAllcategory(c echo.Context) error {

	allCategory, err := controller.service.GetAllCategory()

	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}

	resCategory := request.GetAllCategory(allCategory)

	return c.JSON(response.NewSuccessResponse(resCategory))
}

func (controller *Controller) CreateCategory(c echo.Context) error {
	insertCategory := new(request.InsertReqCategory)

	if err := c.Bind(insertCategory); err != nil {
		return c.JSON(response.NewBadRequestResponse())
	}
	err := controller.service.CreateCategory(*insertCategory.InsertNewCategory())
	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}
	return c.JSON(response.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("categoryId"))
	err := controller.service.DeleteCategory(id)
	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}
	return c.JSON(response.NewSuccessResponseNoContent())

}

func (controller *Controller) UpdateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("categoryId"))
	category := new(request.UpdateReqCategory)

	if err := c.Bind(category); err != nil {
		return c.JSON(response.NewBadRequestResponse())
	}

	err := controller.service.UpdateCategory(*category.UpdateExistCategory(), id)

	if err != nil {
		return c.JSON(response.NewErrorBusinessResponse(err))
	}
	return c.JSON(response.NewSuccessResponseWithoutData())
}
