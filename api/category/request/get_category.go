package request

import "catalog/bussiness/category"

type ResCategoryById struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func GetCategoryById(category category.Category) *ResCategoryById {

	var resCategoryById ResCategoryById
	resCategoryById.Id = int(category.ID)
	resCategoryById.Name = category.Name
	resCategoryById.IsActive = category.IsActive

	return &resCategoryById

}
