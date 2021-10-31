package category

import (
	"catalog/bussiness/category"
	"gorm.io/gorm"
)

type DbRepository struct {
	DB *gorm.DB
}

type CategoryTable struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"id;primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"name"`
	IsActive bool   `json:"is_active" gorm:"is_active"`
}

//get field category form bussiness
func newProduct(categoryTemp category.Category) *CategoryTable {
	return &CategoryTable{
		categoryTemp.Model,
		categoryTemp.ID,
		categoryTemp.Name,
		categoryTemp.IsActive,
	}
}

func (field *CategoryTable) ToCategory() category.Category {

	var category category.Category
	category.ID = field.ID
	category.Name = field.Name
	category.IsActive = field.IsActive

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
