package request

import (
	"catalog/bussiness/category"
)

type InsertReqCategory struct {
	Name      string `json:"name"`
	Is_active bool   `json:"is_active"`
}

func (request *InsertReqCategory) InsertNewCategory() *category.CategoryField {

	var InsertReqCategory category.CategoryField

	InsertReqCategory.Name = request.Name
	InsertReqCategory.Is_active = request.Is_active

	return &InsertReqCategory

}
