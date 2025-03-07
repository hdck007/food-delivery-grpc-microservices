package models

import (
	"time"
)

type Delivery struct {
	ID                   uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	OrderId              uint `gorm:"foreignKey:OrderId"`
	RiderId              uint `gorm:"foreignKey:RiderId"`
	RestaurantId         uint `gorm:"foreignKey:RestaurantId"`
	DestinationLongitude float64
	DestinationLatitude  float64
	Status               string
	CreatedAt            time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime:false"`
}
