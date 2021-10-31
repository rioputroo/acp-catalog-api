package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryId  int
	Name        string
	Price       int
	Description string
	IsActive    bool
}

func NewProduct(
	newCategoryId int,
	newName string,
	newPrice int,
	newDescription string,
	newIsActive bool,
) Product {
	return Product{
		CategoryId:  newCategoryId,
		Name:        newName,
		Price:       newPrice,
		Description: newDescription,
		IsActive:    newIsActive,
	}
}

func UpdateProduct(
	newCategoryId int,
	newName string,
	newPrice int,
	newDescription string,
	newIsActive bool,
) Product {
	return Product{
		CategoryId:  newCategoryId,
		Name:        newName,
		Price:       newPrice,
		Description: newDescription,
		IsActive:    newIsActive,
	}
}
