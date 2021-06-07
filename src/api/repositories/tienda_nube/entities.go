package tienda_nube

import (
	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"strconv"
	"strings"
)

type account struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	UserID      int64  `json:"user_id"`
}

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

func (a account) toEntity() *entities.TnAccount {
	return &entities.TnAccount{
		AccessToken: a.AccessToken,
		UserID:      a.UserID,
	}
}
