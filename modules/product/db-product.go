package product

import (
	"catalog/bussiness/product"
	"gorm.io/gorm"
)

type DbRepository struct {
	DB *gorm.DB
}

type ProductTable struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"id;primaryKey;autoIncrement"`
	CategoryId  int    `json:"category_id" gorm:"category_id"`
	Name        string `json:"name" gorm:"name"`
	Price       int    `json:"price" gorm:"price"`
	Description string `json:"description" gorm:"description"`
	IsActive    bool   `json:"is_active" gorm:"is_active"`
}

//get field product form bussiness
func newProduct(productTemp product.Product) *ProductTable {
	return &ProductTable{
		productTemp.Model,
		productTemp.ID,
		productTemp.CategoryId,
		productTemp.Name,
		productTemp.Price,
		productTemp.Description,
		productTemp.IsActive,
	}
}

func (field *ProductTable) ToProduct() product.Product {

	var product product.Product
	product.ID = field.ID
	product.CategoryId = field.CategoryId
	product.Name = field.Name
	product.Price = field.Price
	product.Description = field.Description
	product.IsActive = field.IsActive
	return product
}

func NewProductRepository(db *gorm.DB) *DbRepository {
	return &DbRepository{
		db,
	}
}

func (temp *DbRepository) GetProductById(id int) (*product.Product, error) {

	var data ProductTable
	err := temp.DB.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	product := data.ToProduct()
	return &product, nil
}

func (temp *DbRepository) GetProductsByCategoryId(categoryId int) ([]product.Product, error) {
	var data []ProductTable
	var result []product.Product

	err := temp.DB.Find(&data, ProductTable{Category_id: categoryId}).Error

	if err != nil {
		return nil, err
	}

	for _, value := range data {
		result = append(result, value.ToProduct())
	}

	return result, nil

}

func (temp *DbRepository) GetAllProducts() ([]product.Product, error) {

	var data []ProductTable
	var result []product.Product

	err := temp.DB.Find(&data).Error
	if err != nil {
		return nil, err
	if &categoryId != nil {
		err := temp.DB.Find(&data, ProductTable{CategoryId: categoryId}).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := temp.DB.Find(&data).Error
		if err != nil {
			return nil, err
		}
	}

	for _, value := range data {
		result = append(result, value.ToProduct())
	}

	return result, nil

}

func (temp *DbRepository) CreateProduct(product product.Product) error {

	data := newProduct(product)

	err := temp.DB.Save(data).Error
	if err != nil {
		return err
	}
	return nil

}

func (temp *DbRepository) UpdateProduct(product product.Product, id int) error {

	data := newProduct(product)
	err := temp.DB.Where("id = ?", &id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil

}

func (temp *DbRepository) DeleteProduct(id int) error {

	var product ProductTable

	err := temp.DB.Delete(&product, id).Error

	if err != nil {
		return err
	}
	return nil

}
