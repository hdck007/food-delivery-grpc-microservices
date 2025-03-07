package models

import "time"

type Restaurant struct {
	ID        uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name      string
	Latitude  float64
	Longitude float64
	MenuItems []MenuItem `gorm:"foreignKey:RestaurantId"`
	CreatedAt time.Time  `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime:false"`
}
