package request

import "catalog/bussiness/category"

type ResCategoryById struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Is_active bool   `json:"is_active"`
}

func GetCategoryById(category category.Category) *ResCategoryById {

	var resCategoryById ResCategoryById
	resCategoryById.Id = category.Id
	resCategoryById.Name = category.Name
	resCategoryById.Is_active = category.Is_active

	return &resCategoryById

}
