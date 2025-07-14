package handlers

import (
  "bff-service/clients"
  orderpb "bff-service/proto/order"
  "bff-service/utils"
  "github.com/gin-gonic/gin"
  "net/http"
)

func PlaceOrder(c *gin.Context) {
  uid := c.GetString("user_id")
  resp, err := clients.OrderClient().PlaceOrder(c, &orderpb.PlaceOrderRequest{UserId: uid})
  if err != nil {
    utils.RespondWithError(c, http.StatusInternalServerError, "Error")
    return
  }
  utils.RespondWithJSON(c, http.StatusOK, resp)
}

func GetOrders(c *gin.Context) {
  uid := c.GetString("user_id")
  resp, err := clients.OrderClient().GetOrders(c, &orderpb.GetOrdersRequest{UserId: uid})
  if err != nil {
    utils.RespondWithError(c, http.StatusInternalServerError, "Error")
    return
  }
  utils.RespondWithJSON(c, http.StatusOK, resp.Orders)
}

func GetOrderDetails(c *gin.Context) {
	orderID := c.Param("order_id") // expect /api/orders/:order_id

	resp, err := clients.OrderClient().GetOrderDetails(c, &orderpb.GetOrderDetailsRequest{
		OrderId: orderID,
	})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch order details")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp.Order)
}

func UpdateOrderStatus(c *gin.Context) {
	var req struct {
		OrderId string `json:"order_id" binding:"required"`
		Status  string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	resp, err := clients.OrderClient().UpdateOrderStatus(c, &orderpb.UpdateOrderStatusRequest{
		OrderId: req.OrderId,
		Status:  req.Status,
	})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update status")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func GenerateInvoice(c *gin.Context) {
	orderID := c.Param("order_id") // expect /api/orders/:order_id/invoice

	resp, err := clients.OrderClient().GenerateInvoice(c, &orderpb.GenerateInvoiceRequest{
		OrderId: orderID,
	})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to generate invoice")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, gin.H{
		"invoice": resp.InvoiceText,
	})
}
