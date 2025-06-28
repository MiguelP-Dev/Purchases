package models

import "time"

// ListItem representa un item de una lista de compras
type ListItem struct {
	ID             uint      `gorm:"primary_key"`
	ShoppingListID uint      `gorm:"not null"`
	ProductName    string    `gorm:"not null"`
	IsPurchased    bool      `gorm:"default:false"`
	Price          float64   `gorm:"default:0.0"`
	Place          string    `gorm:"default:'supermercado'"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
