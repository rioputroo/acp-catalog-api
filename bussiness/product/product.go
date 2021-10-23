package product

import "time"

type Product struct {
	Id          int
	Category_id int
	Name        string
	Price       int
	Description string
	Image       string
	Is_active   bool
	Created_At  time.Time
	Updated_At  time.Time
}

func NewProduct(
	new_category_id int,
	new_name string,
	new_price int,
	new_description string,
	new_image string,
	new_is_active bool,
	new_create_at time.Time,
) Product {
	return Product{
		Category_id: new_category_id,
		Name:        new_name,
		Price:       new_price,
		Description: new_description,
		Image:       new_image,
		Is_active:   new_is_active,
		Created_At:  time.Now().UTC(),
	}
}

func UpdateProduct(
	new_category_id int,
	new_name string,
	new_price int,
	new_description string,
	new_image string,
	new_is_active bool,
	new_update_at time.Time,
) Product {
	return Product{
		Category_id: new_category_id,
		Name:        new_name,
		Price:       new_price,
		Description: new_description,
		Image:       new_image,
		Is_active:   new_is_active,
		Updated_At:  time.Now().UTC(),
	}
}
