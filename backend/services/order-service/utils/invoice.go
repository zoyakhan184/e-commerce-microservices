package utils

import (
	"fmt"
	"order-service/models"
)

func GenerateInvoice(order models.Order) string {
	invoice := fmt.Sprintf(`
==============================
         INVOICE
==============================
Order ID    : %s
User ID     : %s
Date        : %s
Status      : %s
Payment     : %s

Items:
`, order.ID, order.UserID, order.CreatedAt.Format("2006-01-02 15:04"), order.OrderStatus, order.PaymentStatus)

	var total float64 = 0
	for _, item := range order.OrderItems {
		line := fmt.Sprintf(" - %s | Size: %s | Color: %s | Qty: %d | ₹%.2f\n",
			item.ProductID, item.Size, item.Color, item.Quantity, item.Price)
		invoice += line
		total += float64(item.Quantity) * item.Price
	}

	invoice += fmt.Sprintf("\nTotal Amount: ₹%.2f\n", total)
	invoice += "==============================\n"
	return invoice
}
