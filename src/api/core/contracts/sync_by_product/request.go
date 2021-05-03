package sync_by_product

type Request struct {
	SKU   string  `json:"sku" binding:"required"`
	Stock float64 `json:"stock" binding:"omitempty"`
}
