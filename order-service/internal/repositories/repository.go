package repositories

import (
	"hdck007.com/order-service/internal/models"
)

type IOrderRepository interface {
	Create(order *models.Order)
	UpdateStatus(orderId uint, status string) *models.Order
	GetOrders() []models.Order
	GetOrdersByUserId(userId string) []models.Order
	GetOrderById(orderId uint) *models.Order
}
