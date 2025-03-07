package repositories

import (
	"hdck007.com/order-service/internal/models"
)

type IDeliveryRepository interface {
	Create(delivery *models.Delivery) *models.Delivery
	UpdateDeliveryStatus(deliveryId uint, deliveryStatus string) *models.Delivery
}
