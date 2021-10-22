package request

import "catalog/bussiness/product"

type ResProducts struct {
	Products []ResProductById
}

func GetProducts(product []product.Product) ResProducts {

	resProducts := ResProducts{}

	for i := 0; i < len(product); i++ {
		var resProductById ResProductById
		resProductById.Id = product[i].Id
		resProductById.Category_id = product[i].Category_id
		resProductById.Name = product[i].Name
		resProductById.Price = product[i].Price
		resProductById.Description = product[i].Description
		resProductById.Image = product[i].Image
		resProductById.Is_active = product[i].Is_active

		resProducts.Products = append(resProducts.Products, resProductById)
	}
	return resProducts
}
