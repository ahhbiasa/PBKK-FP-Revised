package models

import (
	"time"
)

// Product represents the "products" table.
type Product struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	CategoryID  int       `gorm:"not null" json:"category_id"`
	ShopID      int       `gorm:"not null" json:"shop_id"`
	Stock       int       `gorm:"not null" json:"stock"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// A product belongs to a category
	Category Category `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	Shop     Shop     `gorm:"foreignKey:ShopID;references:ID" json:"shop"`
}
