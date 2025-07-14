package handlers

import (
	"bff-service/clients"
	inventorypb "bff-service/proto/inventory"
	"bff-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get stock for a specific product+size+color
func GetStock(c *gin.Context) {
	productID := c.Query("product_id")
	size := c.Query("size")
	color := c.Query("color")

	res, err := clients.InventoryClient.GetStock(c, &inventorypb.StockRequest{
		ProductId: productID,
		Size:      size,
		Color:     color,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Decrease stock on successful order
func UpdateStockOnOrder(c *gin.Context) {
	var req inventorypb.StockUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := clients.InventoryClient.UpdateStockOnOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Increase stock on restock or cancel
func Restock(c *gin.Context) {
	var req inventorypb.StockUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := clients.InventoryClient.Restock(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Admin: List items with low stock
func ListLowStock(c *gin.Context) {
	threshold := c.Query("threshold")

	res, err := clients.InventoryClient.ListLowStock(c, &inventorypb.LowStockRequest{
		Threshold: int32(utils.ParseInt(threshold, 10)), // utility to safely convert string to int32
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.Items)
}
