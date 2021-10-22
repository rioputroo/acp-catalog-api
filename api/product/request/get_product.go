package request

import "catalog/bussiness/product"

type ResProductById struct {
	Id          int    `json:"id"`
	Category_id int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Is_active   bool   `json:"is_active"`
}

func GetProductById(product product.Product) *ResProductById {

	var resProductById ResProductById
	resProductById.Id = product.Id
	resProductById.Category_id = product.Category_id
	resProductById.Name = product.Name
	resProductById.Price = product.Price
	resProductById.Description = product.Description
	resProductById.Image = product.Image
	resProductById.Is_active = product.Is_active

	return &resProductById

}
