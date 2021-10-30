package product

import (
	"time"
)

type ProductField struct {
	Category_id int
	Name        string
	Price       int
	Description string
	Image       string
	Is_active   bool
	Created_At  time.Time
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) GetProductById(id int) (*Product, error) {

	return s.repository.GetProductById(id)

}

func (s *service) GetAllProducts(categoryId int) ([]Product, error) {

	if &categoryId != nil {
		return s.repository.GetAllProducts(categoryId)
	}

	product, err := s.repository.GetAllProducts(categoryId)

	if err != nil {
		return []Product{}, err
	}
	return product, err

}

func (s *service) CreateProduct(productField ProductField) error {

	product := NewProduct(
		productField.Category_id,
		productField.Name,
		productField.Price,
		productField.Description,
		productField.Image,
		productField.Is_active,
		productField.Created_At,
	)

	err := s.repository.CreateProduct(product)

	if err != nil {
		return err
	}
	return nil

}

func (s *service) UpdateProduct(productField ProductField, id int) error {

	product := UpdateProduct(
		productField.Category_id,
		productField.Name,
		productField.Price,
		productField.Description,
		productField.Image,
		productField.Is_active,
		productField.Created_At,
	)

	err := s.repository.UpdateProduct(product, id)

	if err != nil {
		return err
	}
	return nil

}

func (s *service) DeleteProduct(id int) error {

	err := s.repository.DeleteProduct(id)

	if err != nil {
		return err
	}
	return nil

}
