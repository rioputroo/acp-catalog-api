package category

type CategoryField struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) GetCategoryById(id int) (*Category, error) {
	return s.repository.GetCategoryById(id)
}

func (s *service) GetAllCategory() ([]Category, error) {

	product, err := s.repository.GetAllCategory()
	if err != nil {
		return []Category{}, err
	}
	return product, err

}

func (s *service) CreateCategory(categoryField CategoryField) error {

	category := NewCategory(
		categoryField.Name,
		categoryField.IsActive,
	)

	err := s.repository.CreateCategory(category)

	if err != nil {
		return err
	}
	return nil

}

func (s *service) UpdateCategory(categoryField CategoryField, id int) error {

	category := UpdateCategory(
		categoryField.Name,
		categoryField.IsActive,
	)

	err := s.repository.UpdateCategory(category, id)

	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteCategory(id int) error {
	err := s.repository.DeleteCategory(id)

	if err != nil {
		return err
	}
	return nil
}
