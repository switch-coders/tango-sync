package tango

import "time"

type SearchPrice struct {
	ListPriceNumber string
	LastUpdate      string
}

func NewSearchPriceEMC(lastUpdate *time.Time) SearchPrice {
	var date time.Time
	if lastUpdate != nil {
		date = lastUpdate.UTC()
	} else {
		date = time.Now().UTC().Add(-time.Minute * 120)
	}

	return SearchPrice{
		ListPriceNumber: "1000",
		LastUpdate:      date.Format("2006-01-02T15:04:05"),
	}
}
