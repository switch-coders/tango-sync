package models

import "time"

type Product struct {
	SKU   string `gorm:"PRIMARY_KEY;type:varchar(150);not null"`
	Stock int
	Price float64
}

type Audit struct {
	ID        int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Sku       string `gorm:"type:varchar(100);not null"`
	Job       string `gorm:"type:varchar(20);not null"`
	Value     string `gorm:"type:varchar(100);not null"`
	CreatedAt *time.Time
}

type Account struct {
	TangoKey string `gorm:"PRIMARY_KEY;type:varchar(100);not null"`
	Name     string `gorm:"type:varchar(40);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	JobSync  bool
	JobPrice bool
}
