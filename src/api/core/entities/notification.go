package entities

type Notification struct {
	Message map[string]interface{}
	Topic   string
}

func NewSkuStockNotification(sku string, stock float64) Notification {
	return Notification{
		Message: map[string]interface{}{
			"sku":   sku,
			"stock": stock,
		},
		Topic: "tn-stock-1",
	}
}

func NewSkuPriceNotification(sku string, price float64) Notification {
	return Notification{
		Message: map[string]interface{}{
			"sku":   sku,
			"price": price,
		},
		Topic: "tn-price-1",
	}
}
