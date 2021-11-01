package request

import "catalog/bussiness/product"

type ResProductById struct {
	Id          int    `json:"id"`
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func GetProductById(product product.Product) *ResProductById {

	var resProductById ResProductById
	resProductById.Id = int(product.ID)
	resProductById.CategoryId = product.CategoryId
	resProductById.Name = product.Name
	resProductById.Price = product.Price
	resProductById.Description = product.Description
	resProductById.IsActive = product.IsActive

	return &resProductById

}
