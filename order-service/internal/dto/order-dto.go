package dto

type MenuItemRequest struct {
	MenuItemId   string
	RestaurantId string
	Quantity     int
}

type CreateOrderRequest struct {
	UserId       string
	MenuItemList []MenuItemRequest
}

type UpdateOrderRequest struct {
	OrderId string
	Status  string
}
