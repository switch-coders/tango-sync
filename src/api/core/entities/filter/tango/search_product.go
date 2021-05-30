package tango

import "time"

type SearchProduct struct {
	OnlyEnable string
	LastUpdate string
}

func NewSearchProduct(lastUpdate *time.Time) SearchProduct {
	var date time.Time
	if lastUpdate != nil {
		date = lastUpdate.UTC()
	} else {
		date = time.Now().UTC().Add(-time.Minute * 120)
	}

	return SearchProduct{
		OnlyEnable: "true",
		LastUpdate: date.Format("2006-01-02T15:04:05"),
	}
}
