package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	//	id          int64
	Category_id int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Is_active   bool   `json:"is_active"`
}

type Category struct {
	gorm.Model
	//S	id         int
	Name      string `json:"name"`
	Is_active bool   `json:"is_active"`
}
