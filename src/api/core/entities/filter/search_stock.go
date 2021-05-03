package filter

import "time"

type SearchStock struct {
	WareHouseCode         string
	DiscountPendingOrders string
	LastUpdate            string
}

func NewSearchStocktango(lastUpdate *time.Time) SearchStock {
	var date time.Time
	if lastUpdate != nil {
		date = lastUpdate.UTC()
	} else {
		date = time.Now().UTC().Add(-time.Minute * 120)
	}

	return SearchStock{
		WareHouseCode:         "01",
		DiscountPendingOrders: "true",
		LastUpdate:            date.Format("2006-01-02T15:04:05"),
	}
}
