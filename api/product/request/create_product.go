package request

import "catalog/bussiness/product"

type InsertReqProduct struct {
	Category_id int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Is_active   bool   `json:"is_active"`
}

func (request *InsertReqProduct) InsertNewProduct() *product.ProductField {

	var insertReqProduct product.ProductField

	insertReqProduct.Category_id = request.Category_id
	insertReqProduct.Name = request.Name
	insertReqProduct.Price = request.Price
	insertReqProduct.Description = request.Description
	insertReqProduct.Image = request.Image
	insertReqProduct.Is_active = request.Is_active

	return &insertReqProduct

}
