package dto

type CreateRestaurantRequest struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type CreateMenuItemRequest struct {
	Name         string
	RestaurantId uint
	Price        float64
	Stock        int
}
