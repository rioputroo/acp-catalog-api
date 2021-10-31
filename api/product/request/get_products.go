package request

import "catalog/bussiness/product"

type ResProducts struct {
	Products []ResProductById
}

func GetProducts(product []product.Product) ResProducts {

	resProducts := ResProducts{}

	for i := 0; i < len(product); i++ {
		var resProductById ResProductById
		resProductById.Id = int(product[i].ID)
		resProductById.CategoryId = product[i].CategoryId
		resProductById.Name = product[i].Name
		resProductById.Price = product[i].Price
		resProductById.Description = product[i].Description
		resProductById.IsActive = product[i].IsActive

		resProducts.Products = append(resProducts.Products, resProductById)
	}
	return resProducts
}
