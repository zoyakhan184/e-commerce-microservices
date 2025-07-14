package models

type Inventory struct {
	SKU       string `json:"sku"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Size      string `json:"size"`
	Color     string `json:"color"`
}
