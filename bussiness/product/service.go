package product

import "fmt"

type ProductField struct {
	Category_id int
	Name        string
	Price       int
	Description string
	Image       string
	Is_active   bool
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

	product, _ := s.repository.GetProductById(id)

	fmt.Println("Service", product)

	return product, nil
}

func (s *service) GetAllProducts() ([]Product, error) {

	product, err := s.repository.GetAllProducts()

	fmt.Println("Service", product)

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
	)

	fmt.Println("service", &product)

	err := s.repository.CreateProduct(product)

	if err != nil {
		return err
	}
	return nil

}

func (s *service) UpdateProduct(productField ProductField, id int) error {

	product := NewProduct(
		productField.Category_id,
		productField.Name,
		productField.Price,
		productField.Description,
		productField.Image,
		productField.Is_active,
	)

	fmt.Println(productField)

	err := s.repository.UpdateProduct(product, id)

	if err != nil {
		return err
	}
	return nil

}

func (s *service) DeleteProduct(id int) error {

	err := s.repository.DeleteProduct(id)

	fmt.Println(&err)
	if err != nil {
		return err
	}
	return nil

}
