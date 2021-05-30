package price

type Request struct {
	SKU   string  `json:"sku" binding:"required"`
	Price float64 `json:"price" binding:"omitempty"`
}
