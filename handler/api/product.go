package api

import (
	"catalog/database"
	"catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {

	var products []models.Product

	database.DB.Find(&products)

	if len(products) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "successfully fetch list of products",
			"products": products,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Internal Server Error",
	})
}

func GetProductsById(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("productId"))

	var productsById []models.Product

	database.DB.First(&productsById, id)

	if len(productsById) > 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "successfully fetch product",
			"products": productsById,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Internal Server Error",
	})
}

func CreateProduct(c echo.Context) error {

	var products models.Product
	var category []models.Category

	database.DB.Find(&category)

	category_id, _ := strconv.Atoi(c.FormValue("category_id"))
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	description := c.FormValue("description")
	image := c.FormValue("image")
	is_active, _ := strconv.ParseBool(c.FormValue("is_active"))

	products.Category_id = category_id
	products.Name = name
	products.Price = int(price)
	products.Description = description
	products.Image = image
	products.Is_active = bool(is_active)

	c.Bind(&products)

	if err := database.DB.Save(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":  "Product Created",
		"products": products,
	})
}

func UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("productId"))
	var productById []models.Category
	var products models.Product
	database.DB.First(&productById, id)

	category_id, _ := strconv.Atoi(c.FormValue("category_id"))
	name := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	description := c.FormValue("description")
	image := c.FormValue("image")
	is_active, _ := strconv.ParseBool(c.FormValue("is_active"))

	products.Category_id = category_id
	products.Name = name
	products.Price = int(price)
	products.Description = description
	products.Image = image
	products.Is_active = is_active

	database.DB.Model(&productById).Updates(&products)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Successfully update product",
		"products": products,
	})
}

func DeleteProduct(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("productId"))
	var products []models.Product

	database.DB.Delete(&products, id)

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"message": "Successfully delete product",
	})
}
