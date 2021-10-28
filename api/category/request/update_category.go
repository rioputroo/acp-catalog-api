package request

import "catalog/bussiness/category"

type UpdateReqCategory struct {
	Name      string `json:"name"`
	Is_active bool   `json:"is_active"`
}

func (request *UpdateReqCategory) UpdateExistCategory() *category.CategoryField {

	var updateExistCategory category.CategoryField

	updateExistCategory.Name = request.Name
	updateExistCategory.Is_active = request.Is_active

	return &updateExistCategory

}
