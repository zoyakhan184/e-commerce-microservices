package handlers

import (
	"bff-service/clients"
	productpb "bff-service/proto/product"
	"bff-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List all products with optional filters
func ListProducts(c *gin.Context) {
	filter := &productpb.ProductFilter{
		CategoryId: c.Query("category_id"),
		Brand:      c.Query("brand"),
	}

	resp, err := clients.ProductClient().ListProducts(c, filter)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list products")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp.Products)
}

// Get product details by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	resp, err := clients.ProductClient().GetProduct(c, &productpb.ProductIdRequest{Id: id})
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Product not found")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// Add a new product
func AddProduct(c *gin.Context) {
	var req productpb.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	resp, err := clients.ProductClient().AddProduct(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to add product")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// Edit an existing product
func EditProduct(c *gin.Context) {
	var req productpb.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	resp, err := clients.ProductClient().EditProduct(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update product")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	resp, err := clients.ProductClient().DeleteProduct(c, &productpb.ProductIdRequest{Id: id})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// Add a category (or subcategory)
func AddCategory(c *gin.Context) {
	var req productpb.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid category input")
		return
	}

	resp, err := clients.ProductClient().AddCategory(c, &req)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to add category")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

// List all categories
func ListCategories(c *gin.Context) {
	resp, err := clients.ProductClient().ListCategories(c, &productpb.Empty{})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp.Categories)
}
