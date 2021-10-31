package request

import "catalog/bussiness/product"

type UpdateReqProduct struct {
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (request *UpdateReqProduct) UpdateExistProduct() *product.ProductField {

	var updateExistProduct product.ProductField

	updateExistProduct.CategoryId = request.CategoryId
	updateExistProduct.Name = request.Name
	updateExistProduct.Price = request.Price
	updateExistProduct.Description = request.Description
	updateExistProduct.IsActive = request.IsActive

	return &updateExistProduct

}
