package handlers

import (
	"bff-service/clients"
	paymentpb "bff-service/proto/payment"
	"bff-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Initiate payment for an order
func InitiatePayment(c *gin.Context) {
	var req paymentpb.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	// Ensure required fields are present
	if req.OrderId == "" || req.UserId == "" || req.Gateway == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Missing required fields: order_id, user_id, gateway")
		return
	}

	resp, err := clients.PaymentClient().InitiatePayment(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error initiating payment: "+err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// Verify payment status
func VerifyPayment(c *gin.Context) {
	var req paymentpb.VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.PaymentId == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Missing payment_id")
		return
	}

	resp, err := clients.PaymentClient().VerifyPayment(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error verifying payment: "+err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// Refund a payment
func RefundPayment(c *gin.Context) {
	var req paymentpb.RefundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.OrderId == "" && req.PaymentId == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Missing order_id or payment_id")
		return
	}

	resp, err := clients.PaymentClient().RefundPayment(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error processing refund: "+err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, resp)
}
