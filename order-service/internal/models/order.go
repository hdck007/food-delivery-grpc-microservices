package models

import (
	"time"
)

type Order struct {
	ID         uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserId     string
	Status     string
	CreatedAt  time.Time   `gorm:"autoCreateTime:false"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime:false"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`
}
