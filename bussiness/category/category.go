package category

import "time"

type Category struct {
	Id         int
	Name       string
	Is_active  bool
	Created_At time.Time
	Updated_At time.Time
}

func NewCategory(
	new_name string,
	new_is_active bool,
) Category {
	return Category{
		Name:      new_name,
		Is_active: new_is_active,
	}
}
