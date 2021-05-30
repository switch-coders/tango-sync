package tienda_nube

import (
	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"strconv"
	"strings"
)

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

func (p products) toEntity(sku string) *entities.Product {
	for _, products := range p {
		for _, v := range products.Variants {
			if strings.EqualFold(v.Sku, sku) {
				price, _ := strconv.ParseFloat(v.Price, 64)

				return &entities.Product{
					ID:              v.ID,
					ProductID:       v.ProductID,
					Price:           price,
					StockManagement: v.StockManagement,
					Stock:           v.Stock,
					Sku:             v.Sku,
				}
			}
		}
	}

	return nil
}
