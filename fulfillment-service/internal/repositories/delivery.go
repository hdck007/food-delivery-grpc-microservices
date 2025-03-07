package repositories

import (
	"hdck007.com/order-service/internal/database"
	"hdck007.com/order-service/internal/models"
)

type DeliveryRepository struct {
	databaseEngine *database.DatabaseEngine
}

func NewDeliveryRepository(databaseEngine *database.DatabaseEngine) IDeliveryRepository {
	return &DeliveryRepository{
		databaseEngine: databaseEngine,
	}
}

func (deliveryRepository *DeliveryRepository) Create(delivery *models.Delivery) *models.Delivery {
	deliveryRepository.databaseEngine.Db.Create(&delivery)
	return delivery
}

func (deliveryRepository *DeliveryRepository) UpdateDeliveryStatus(deliveryId uint, deliveryStatus string) *models.Delivery {
	delivery := &models.Delivery{}
	deliveryRepository.databaseEngine.Db.Model(&models.Delivery{}).Where("id = ?", deliveryId).Updates(map[string]interface{}{"status": deliveryStatus})
	return delivery
}
