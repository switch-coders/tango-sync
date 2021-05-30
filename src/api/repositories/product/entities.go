package product

import (
	"github.com/switch-coders/tango-sync/src/api/core/entities"
)

type product struct {
	SKU   string   `gorm:"PRIMARY_KEY;type:varchar(150);not null"`
	Stock *int     `gorm:"null"`
	Price *float64 `gorm:"null"`
}

func (p product) toEntity() *entities.Product {
	var product = entities.Product{
		Sku: p.SKU,
	}

	if p.Price != nil {
		product.Price = *p.Price
	}

	if p.Stock != nil {
		product.Stock = *p.Stock
	}

	return &product
}
