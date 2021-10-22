package migration

import (
	"catalog/modules/category"
	"catalog/modules/product"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&product.ProductTable{}, &category.CategoryTable{})
}
