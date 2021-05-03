package models

type Product struct {
	SKU   string `gorm:"primaryKey;type:varchar(150);not null"`
	Stock int    `gorm:"not null"`
}
