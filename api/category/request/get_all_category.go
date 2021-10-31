package request

import (
	"catalog/bussiness/category"
	"gorm.io/gorm"
)

type ResCategory struct {
	*gorm.Model
	Category []ResCategoryById `json:"category"`
}

func GetAllCategory(categories []category.Category) ResCategory {
	var resCategoryById ResCategoryById

	resCategory := ResCategory{}

	for i := 0; i < len(categories); i++ {
		resCategoryById.Id = int(categories[i].ID)
		resCategoryById.Name = categories[i].Name
		resCategoryById.IsActive = categories[i].IsActive
		resCategory.Category = append(resCategory.Category, resCategoryById)
	}
	return resCategory
}
