package models

import (
	"time"
)

// Category represents the "categories" table.
type Category struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// A category can have many products
	Products []Product `gorm:"foreignKey:CategoryID;references:ID" json:"products"`
}
