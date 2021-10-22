package request

import "catalog/bussiness/product"

type UpdateReqProduct struct {
	Category_id int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Is_active   bool   `json:"is_active"`
}

func (request *UpdateReqProduct) UpdateExistProduct() *product.ProductField {

	var updateExistProduct product.ProductField

	updateExistProduct.Category_id = request.Category_id
	updateExistProduct.Name = request.Name
	updateExistProduct.Price = request.Price
	updateExistProduct.Description = request.Description
	updateExistProduct.Image = request.Image
	updateExistProduct.Is_active = request.Is_active

	return &updateExistProduct

}
