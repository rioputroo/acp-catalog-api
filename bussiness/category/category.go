package category

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func NewCategory(
	newName string,
	newIsActive bool,
) Category {
	return Category{
		Name:     newName,
		IsActive: newIsActive,
	}
}

func UpdateCategory(
	new_name string,
	new_is_active bool,
) Category {
	return Category{
		Name:     new_name,
		IsActive: new_is_active,
	}
}
