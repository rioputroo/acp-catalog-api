package product

import (
	"catalog/bussiness/product"
	"time"

	"gorm.io/gorm"
)

type DbRepository struct {
	DB *gorm.DB
}

type ProductTable struct {
	Id          int    `json:"id;primaryKey;autoIncrement"`
	Category_id int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Is_active   bool   `json:"is_active"`
	Created_At  time.Time
	Updated_At  time.Time
}

//get field product form bussiness
func newProduct(productTemp product.Product) *ProductTable {
	return &ProductTable{
		productTemp.Id,
		productTemp.Category_id,
		productTemp.Name,
		productTemp.Price,
		productTemp.Description,
		productTemp.Image,
		productTemp.Is_active,
		productTemp.Created_At,
		productTemp.Updated_At,
	}
}

func (field *ProductTable) ToProduct() product.Product {

	var product product.Product
	product.Id = field.Id
	product.Category_id = field.Category_id
	product.Name = field.Name
	product.Price = field.Price
	product.Description = field.Description
	product.Image = field.Image
	product.Is_active = field.Is_active
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
