package services

import (
	"hdck007.com/order-service/internal/dto"
	"hdck007.com/order-service/internal/models"
	"hdck007.com/order-service/internal/repositories"
)

type DeliveryService struct {
	repository repositories.IDeliveryRepository
}

func NewDeliveryService(deliveryRepository repositories.IDeliveryRepository) *DeliveryService {
	return &DeliveryService{
		repository: deliveryRepository,
	}
}

func (deliveryService *DeliveryService) CreateDelivery(createDeliveryRequest dto.CreateDeliveryRequest) *models.Delivery {

	delivery := &models.Delivery{
		OrderId: createDeliveryRequest.OrderId,
		RiderId: createDeliveryRequest.RiderId,
		Status:  "PENDING",
	}

	return deliveryService.repository.Create(delivery)
}

func (deliveryService *DeliveryService) UpdateDeliveryStatus(deliveryId uint, deliveryStatus string) *models.Delivery {
	return deliveryService.repository.UpdateDeliveryStatus(deliveryId, deliveryStatus)
}
