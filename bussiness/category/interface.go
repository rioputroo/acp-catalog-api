package category

type Service interface {
	GetAllCategory() ([]Category, error)

	GetCategoryById(id int) (*Category, error)

	CreateCategory(categoryField CategoryField) error

	UpdateCategory(categoryField CategoryField, id int) error

	DeleteCategory(id int) error
}

type Repository interface {
	GetAllCategory() ([]Category, error)

	GetCategoryById(id int) (*Category, error)

	CreateCategory(category Category) error

	UpdateCategory(category Category, id int) error

	DeleteCategory(id int) error
}
