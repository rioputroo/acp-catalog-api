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
	new_create_at time.Time,
) Category {
	return Category{
		Name:       new_name,
		Is_active:  new_is_active,
		Created_At: time.Now().UTC(),
	}
}

func UpdateCategory(
	new_name string,
	new_is_active bool,
	new_update_at time.Time,
) Category {
	return Category{
		Name:       new_name,
		Is_active:  new_is_active,
		Updated_At: time.Now().UTC(),
	}
}
