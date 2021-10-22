package category

import (
	"catalog/api/category/request"
	"catalog/bussiness/category"
	"fmt"
	"net/http"
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "successfully fetch category",
		"category": category,
	})
}

func (controller *Controller) GetAllcategory(c echo.Context) error {

	allCategory, err := controller.service.GetAllCategory()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "successfully fetch list of category",
		"category": allCategory,
	})
}

func (controller *Controller) CreateCategory(c echo.Context) error {
	insertCategory := new(request.InsertReqCategory)

	if err := c.Bind(insertCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	err := controller.service.CreateCategory(*insertCategory.InsertNewCategory())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Category Created",
	})
}

func (controller *Controller) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("categoryId"))
	err := controller.service.DeleteCategory(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"message": "successfully delete product",
	})

}

func (controller *Controller) UpdateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("categoryId"))

	product := new(request.UpdateReqCategory)

	fmt.Println(product)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	controller.service.UpdateCategory(*product.UpdateExistCategory(), id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully update category",
	})
}
