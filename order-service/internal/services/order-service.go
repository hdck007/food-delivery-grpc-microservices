package services

import (
	"hdck007.com/order-service/internal/dto"
	"hdck007.com/order-service/internal/models"
	"hdck007.com/order-service/internal/repositories"
)

type OrderService struct {
	repository repositories.IOrderRepository
}

func NewOrderService(orderRepository repositories.IOrderRepository) *OrderService {
	return &OrderService{
		repository: orderRepository,
	}
}

func (orderService *OrderService) CreateOrder(createOrderQuery dto.CreateOrderRequest) *models.Order {
	order := &models.Order{
		UserId: createOrderQuery.UserId,
		Status: "PENDING",
	}

	menuItemsToCreate := []models.OrderItem{}
	for _, menuItem := range createOrderQuery.MenuItemList {
		menuItemsToCreate = append(menuItemsToCreate, models.OrderItem{
			MenuItemId:   menuItem.MenuItemId,
			RestaurantId: menuItem.RestaurantId,
			Quantity:     menuItem.Quantity,
			OrderId:      order.ID,
		})
	}

	order.OrderItems = menuItemsToCreate
	orderService.repository.Create(order)

	return order
}

func (orderService *OrderService) GetOrderById(orderId uint) *models.Order {
	return orderService.repository.GetOrderById(orderId)
}

func (orderService *OrderService) GetOrdersByUserId(userId string) []models.Order {
	return orderService.repository.GetOrdersByUserId(userId)
}

func (orderService *OrderService) UpdateStatus(orderId uint, status string) *models.Order {
	return orderService.repository.UpdateStatus(orderId, status)
}

func (orderService *OrderService) GetOrders() []models.Order {
	return orderService.repository.GetOrders()
}
