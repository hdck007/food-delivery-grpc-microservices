package repositories

import (
	"hdck007.com/order-service/internal/database"
	"hdck007.com/order-service/internal/models"
)

type MenuItemRepository struct {
	databaseEngine *database.DatabaseEngine
}

func NewMenuItemRepository(databaseEngine *database.DatabaseEngine) IMenuItemRepository {
	return &MenuItemRepository{
		databaseEngine: databaseEngine,
	}
}

func (menuItemRepository *MenuItemRepository) Create(menuItem *models.MenuItem) {
	menuItemRepository.databaseEngine.Db.Create(&menuItem)
}

func (menuItemRepository *MenuItemRepository) GetMenuItemById(menuItemId uint) *models.MenuItem {
	var menuItem models.MenuItem
	menuItemRepository.databaseEngine.Db.Preload("MenuItems").Find(&menuItem).Where(&models.MenuItem{ID: menuItemId})
	return &menuItem
}

func (menuItemRepository *MenuItemRepository) GetMenuItemsByRestaurantId(restaurantId uint) []models.MenuItem {
	var menuItems []models.MenuItem
	menuItemRepository.databaseEngine.Db.Preload("MenuItems").Find(&menuItems).Where(&models.MenuItem{RestaurantId: restaurantId})
	return menuItems
}
