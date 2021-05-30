package tango

import (
	"github.com/switch-coders/tango-sync/src/api/core/entities"
)

type stock struct {
	Paging struct {
		PageNumber int  `json:"PageNumber"`
		PageSize   int  `json:"PageSize"`
		MoreData   bool `json:"MoreData"`
	} `json:"Paging"`
	Data []struct {
		StoreNumber     int     `json:"StoreNumber"`
		WareHouseCode   string  `json:"WarehouseCode"`
		SkuCode         string  `json:"SKUCode"`
		Quantity        float64 `json:"Quantity"`
		EngagedQuantity float64 `json:"EngagedQuantity"`
		PendingQuantity float64 `json:"PendingQuantity"`
	} `json:"Data"`
}

type price struct {
	Paging struct {
		PageNumber int  `json:"PageNumber"`
		PageSize   int  `json:"PageSize"`
		MoreData   bool `json:"MoreData"`
	} `json:"Paging"`
	Data []struct {
		PriceListNumber int     `json:"PriceListNumber"`
		SkuCode         string  `json:"SKUCode"`
		Price           float64 `json:"Price"`
	} `json:"Data"`
}

func (s *stock) GetEntity() []entities.Stock {
	stocks := make([]entities.Stock, len(s.Data))

	for i := 0; i < len(s.Data); i++ {
		data := s.Data[i]

		stocks[i] = entities.Stock{
			StoreNumber:     data.StoreNumber,
			WareHouseCode:   data.WareHouseCode,
			SkuCode:         data.SkuCode,
			Quantity:        data.Quantity,
			EngagedQuantity: data.EngagedQuantity,
			PendingQuantity: data.PendingQuantity,
		}

		if stocks[i].Quantity < 0 {
			stocks[i].Quantity = 0
		}
	}

	return stocks
}

func (s *price) GetPriceEntity() []entities.Price {
	stocks := make([]entities.Price, len(s.Data))

	for i := 0; i < len(s.Data); i++ {
		data := s.Data[i]

		stocks[i] = entities.Price{
			PriceListNumber: data.PriceListNumber,
			SkuCode:         data.SkuCode,
			Price:           data.Price,
		}

		if stocks[i].Price < 0 {
			stocks[i].Price = 0
		}
	}

	return stocks
}
