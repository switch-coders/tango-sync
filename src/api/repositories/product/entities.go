package product

import (
	"time"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
)

type product struct {
	SKU       string    `gorm:"type:varchar(150);not null"`
	Stock     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

func (p product) toEntity() *entities.Product {
	return &entities.Product{
		Stock: p.Stock,
		Sku:   p.SKU,
	}
}
