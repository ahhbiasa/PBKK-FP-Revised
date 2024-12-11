package models

import "time"

type Shop struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	Address   string    `gorm:"type:varchar(100);not null" json:"address" form:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// A category can have many products
	Products []Product `gorm:"foreignKey:ShopID;references:ID" json:"products"`
}
