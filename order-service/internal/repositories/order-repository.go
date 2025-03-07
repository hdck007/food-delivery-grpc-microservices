package repositories

import (
	"hdck007.com/order-service/internal/database"
	"hdck007.com/order-service/internal/models"
)

type OrderRepository struct {
	databaseEngine *database.DatabaseEngine
}

func NewOrderRepository(databaseEngine *database.DatabaseEngine) IOrderRepository {
	return &OrderRepository{
		databaseEngine: databaseEngine,
	}
}

func (orderRepository *OrderRepository) Create(
	order *models.Order,
) {
	orderRepository.databaseEngine.Db.Create(order)
}

func (orderRepository *OrderRepository) UpdateStatus(orderId uint, status string) *models.Order {
	order := orderRepository.GetOrderById(orderId)
	order.Status = status
	orderRepository.databaseEngine.Db.Save(order)
	return order
}

func (orderRepository *OrderRepository) GetOrderById(orderId uint) *models.Order {
	var order models.Order
	orderRepository.databaseEngine.Db.Preload("OrderItems").Find(&order).Where(&models.Order{ID: orderId})
	return &order
}

func (orderRepository *OrderRepository) GetOrders() []models.Order {
	var orders []models.Order
	orderRepository.databaseEngine.Db.Find(&orders)
	return orders
}

func (orderRepository *OrderRepository) GetOrdersByUserId(userId string) []models.Order {
	var orders []models.Order
	orderRepository.databaseEngine.Db.Find(orders).Where(&models.Order{UserId: userId})
	return orders
}
