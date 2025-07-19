package handlers

import (
	"bff-service/clients"
	adminpb "bff-service/proto/admin"
	"bff-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	resp, err := clients.AdminClient().GetDashboardData(c, &adminpb.Empty{})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func ListUsers(c *gin.Context) {
	resp, err := clients.AdminClient().ListAllUsers(c, &adminpb.Empty{})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp.Users)
}

func ListAllOrders(c *gin.Context) {
	resp, err := clients.AdminClient().ViewAllOrders(c, &adminpb.Empty{})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp.Orders)
}

func GetRecentActivity(c *gin.Context) {
	resp, err := clients.AdminClient().GetRecentActivity(c, &adminpb.Empty{})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch recent activity")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, resp.Activities)
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("userId")

	_, err := clients.AdminClient().DeleteUser(c, &adminpb.DeleteUserRequest{
		UserId: userID,
	})

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}