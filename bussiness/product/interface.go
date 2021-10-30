package product

type Service interface {
	GetAllProducts(categoryId int) ([]Product, error)

	GetProductById(id int) (*Product, error)

	CreateProduct(productField ProductField) error

	UpdateProduct(productField ProductField, id int) error

	DeleteProduct(id int) error
}

type Repository interface {
	GetAllProducts(categoryId int) ([]Product, error)

	GetProductById(id int) (*Product, error)

	CreateProduct(product Product) error

	UpdateProduct(product Product, id int) error

	DeleteProduct(id int) error
}
