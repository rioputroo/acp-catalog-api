package request

import "catalog/bussiness/product"

type InsertReqProduct struct {
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (request *InsertReqProduct) InsertNewProduct() *product.ProductField {

	var insertReqProduct product.ProductField

	insertReqProduct.CategoryId = request.CategoryId
	insertReqProduct.Name = request.Name
	insertReqProduct.Price = request.Price
	insertReqProduct.Description = request.Description
	insertReqProduct.IsActive = request.IsActive

	return &insertReqProduct

}
