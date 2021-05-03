package entities

type Stock struct {
	StoreNumber     int
	WareHouseCode   string
	SkuCode         string
	Quantity        float64
	EngagedQuantity float64
	PendingQuantity float64
}
