package services

import (
	"hdck007.com/order-service/internal/dto"
	"hdck007.com/order-service/internal/models"
	"hdck007.com/order-service/internal/repositories"
)

type MenuItemService struct {
	repository repositories.IMenuItemRepository
}

func NewMenuItemService(menuItemRepository repositories.IMenuItemRepository) *MenuItemService {
	return &MenuItemService{
		repository: menuItemRepository,
	}
}

func (menuItemService *MenuItemService) CreateMenuItem(createMenuItemQuery dto.CreateMenuItemRequest) *models.MenuItem {
	menuItem := &models.MenuItem{
		Name: createMenuItemQuery.Name,
	}

	menuItemService.repository.Create(menuItem)

	return menuItem
}

func (menuItemService *MenuItemService) GetMenuItemById(menuItemId uint) *models.MenuItem {
	return menuItemService.repository.GetMenuItemById(menuItemId)
}

func (menuItemService *MenuItemService) GetMenuItemsByRestaurantId(restaurantId uint) []models.MenuItem {
	return menuItemService.repository.GetMenuItemsByRestaurantId(restaurantId)
}
