package tienda_nube

import "github.com/switch-coders/tango-sync/src/api/core/entities"

type products []struct {
	Variants []variants `json:"variants"`
}

type variants struct {
	ID              int    `json:"id"`
	ProductID       int    `json:"product_id"`
	Price           string `json:"price"`
	StockManagement bool   `json:"stock_management"`
	Stock           int    `json:"stock"`
	Sku             string `json:"sku"`
}

func (p products) toEntity() *entities.Product {
	product := p[0].Variants
	variant := product[0]
	return &entities.Product{
		ID:              variant.ID,
		ProductID:       variant.ProductID,
		Price:           variant.Price,
		StockManagement: variant.StockManagement,
		Stock:           variant.Stock,
		Sku:             variant.Sku,
	}
}
