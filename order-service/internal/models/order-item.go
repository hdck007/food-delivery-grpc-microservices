package models

import "time"

type OrderItem struct {
	ID           uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	OrderId      uint `gorm:"foreignKey:OrderId"`
	MenuItemId   string
	RestaurantId string
	Quantity     int
	CreatedAt    time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:false"`
}
