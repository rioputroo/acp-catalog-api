package product

import "catalog/modules/product/rabbitmq/consumer"

type ProductField struct {
	CategoryId  int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
type service struct {
	repository Repository
	rabbit     consumer.RabbitRepository
}

func NewService(repository Repository, rabbitmq *consumer.RabbitRepository) Service {
	return &service{
		repository,
		*rabbitmq,
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
		productField.CategoryId,
		productField.Name,
		productField.Price,
		productField.Description,
		productField.IsActive,
	)

	err := s.repository.CreateProduct(product)

	if err != nil {
		return err
	}
	return nil

}

func (s *service) UpdateProduct(productField ProductField, id int) error {

	product := UpdateProduct(
		productField.CategoryId,
		productField.Name,
		productField.Price,
		productField.Description,
		productField.IsActive,
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
