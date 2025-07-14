package handlers

import (
	"bff-service/clients"
	"net/http"
	reviewpb "bff-service/proto/review"

	"github.com/gin-gonic/gin"
)

// ---------- BFF HANDLERS ----------

func GetReviews(c *gin.Context) {
	productID := c.Param("product_id")

	res, err := clients.ReviewClient.GetReviews(c, &reviewpb.GetReviewsRequest{
		ProductId: productID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Reviews)
}

func AddReview(c *gin.Context) {
	var req reviewpb.AddReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	uid := c.GetString("user_id")
	req.UserId = uid

	res, err := clients.ReviewClient.AddReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func EditReview(c *gin.Context) {
	var req reviewpb.EditReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	res, err := clients.ReviewClient.EditReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteReview(c *gin.Context) {
	reviewID := c.Param("review_id")

	res, err := clients.ReviewClient.DeleteReview(c, &reviewpb.DeleteReviewRequest{
		ReviewId: reviewID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
