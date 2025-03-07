package repositories

import (
	"hdck007.com/order-service/internal/models"
)

type IRestaurantRepository interface {
	Create(restaurant *models.Restaurant)
	GetRestaurantById(restaurantId uint) *models.Restaurant
	GetRestaurants() []models.Restaurant
	GetRestaurantsByName(name string) []models.Restaurant
}

type IMenuItemRepository interface {
	Create(createMenuItemRequest *models.MenuItem)
	GetMenuItemsByRestaurantId(restaurantId uint) []models.MenuItem
	GetMenuItemById(itemId uint) *models.MenuItem
}
