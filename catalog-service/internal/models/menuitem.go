package models

import (
	"time"
)

type MenuItem struct {
	ID           uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name         string
	RestaurantId uint `gorm:"foreignKey:RestaurantId"`
	Price        float64
	Stock        int
	CreatedAt    time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:false"`
}
