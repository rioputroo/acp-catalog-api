package request

import "catalog/bussiness/category"

type ResCategory struct {
	Category []ResCategoryById
}

func GetAllCategory(product []category.Category) ResCategory {

	resCategory := ResCategory{}

	for i := 0; i < len(product); i++ {
		var resCategoryById ResCategoryById
		resCategoryById.Id = product[i].Id
		resCategoryById.Name = product[i].Name
		resCategoryById.Is_active = product[i].Is_active
		resCategory.Category = append(resCategory.Category, resCategoryById)
	}
	return resCategory
}
