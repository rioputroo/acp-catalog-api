package category

import (
	"catalog/bussiness/category"
	"time"

	"gorm.io/gorm"
)

type DbRepository struct {
	DB *gorm.DB
}

type CategoryTable struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Is_active  bool   `json:"is_active"`
	Created_At time.Time
	Updated_At time.Time
}

//get field category form bussiness
func newProduct(categoryTemp category.Category) *CategoryTable {
	return &CategoryTable{
		categoryTemp.Id,
		categoryTemp.Name,
		categoryTemp.Is_active,
		categoryTemp.Created_At,
		categoryTemp.Updated_At,
	}
}

func (field *CategoryTable) ToCategory() category.Category {

	var category category.Category
	category.Id = field.Id
	category.Name = field.Name
	category.Is_active = field.Is_active

	return category
}

func NewCategoryRepository(db *gorm.DB) *DbRepository {
	return &DbRepository{
		db,
	}
}

func (temp *DbRepository) GetCategoryById(id int) (*category.Category, error) {

	var data CategoryTable
	err := temp.DB.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	category := data.ToCategory()
	return &category, nil
}

func (temp *DbRepository) GetAllCategory() ([]category.Category, error) {

	var data []CategoryTable
	err := temp.DB.Find(&data).Error
	if err != nil {
		return nil, err
	}

	var result []category.Category
	for _, value := range data {
		result = append(result, value.ToCategory())
	}

	return result, nil

}

func (temp *DbRepository) CreateCategory(category category.Category) error {

	data := newProduct(category)
	err := temp.DB.Save(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (temp *DbRepository) UpdateCategory(category category.Category, id int) error {

	data := newProduct(category)
	err := temp.DB.Where("id = ?", &id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil

}

func (temp *DbRepository) DeleteCategory(id int) error {

	var category CategoryTable

	err := temp.DB.Delete(&category, id).Error
	if err != nil {
		return err
	}
	return nil
}
