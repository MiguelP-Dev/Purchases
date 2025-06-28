package models

import "time"

// ShoppingList representa una lista de compras
type ShoppingList struct {
	ID        uint       `gorm:"primary_key"`
	Title     string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	Items     []ListItem `gorm:"constraint:OnDelete:CASCADE;"`
}
