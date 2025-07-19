package handlers

import (
	"context"

	adminpb "admin-service/proto" // ✅ correct import alias
	"admin-service/repository"
	userpb "user-service/proto" // already available if you’ve generated it
)

type AdminService struct {
	adminpb.UnimplementedAdminServiceServer // ✅ use the correct prefix
	Repo                                    *repository.AdminRepo
	UserClient                              userpb.UserServiceClient // ✅ New field
}

func (s *AdminService) GetDashboardData(ctx context.Context, _ *adminpb.Empty) (*adminpb.DashboardData, error) {
	userCount, _ := s.Repo.CountUsers()
	orderCount, _ := s.Repo.CountOrders()
	revenue, _ := s.Repo.SumRevenue()
	lowStock, _ := s.Repo.ListLowStockThreshold(10)

	var items []*adminpb.LowStockItem
	for _, i := range lowStock {
		items = append(items, &adminpb.LowStockItem{
			ProductId: i.ProductID,
			Quantity:  int32(i.Quantity),
		})
	}

	return &adminpb.DashboardData{
		TotalUsers:    int32(userCount),
		TotalOrders:   int32(orderCount),
		TotalRevenue:  revenue,
		LowStockItems: items,
	}, nil
}

func (s *AdminService) ListAllUsers(ctx context.Context, _ *adminpb.Empty) (*adminpb.ListUsersResponse, error) {
	// Get basic user info (ID, name, email) from local DB
	users, err := s.Repo.ListUsers()
	if err != nil {
		return nil, err
	}

	// Fetch all profile data from user-service
	profilesResp, err := s.UserClient.ListAllProfiles(ctx, &userpb.Empty{})
	if err != nil {
		return nil, err
	}

	// Map profile data for fast lookup
	profileMap := make(map[string]*userpb.UserProfile)
	for _, p := range profilesResp.Profiles {
		profileMap[p.UserId] = p
	}

	// Merge data from auth DB + profile service
	var result []*adminpb.User
	for _, u := range users {
		p := profileMap[u.UserID]

		result = append(result, &adminpb.User{
			UserId:    u.UserID,
			Name:      u.Name,
			Email:     u.Email,
			Phone:     p.GetPhone(),
			Gender:    p.GetGender(),
			Dob:       p.GetDob(),
			AvatarUrl: p.GetAvatarUrl(),
		})
	}

	return &adminpb.ListUsersResponse{Users: result}, nil
}

func (s *AdminService) ViewAllOrders(ctx context.Context, _ *adminpb.Empty) (*adminpb.ListOrdersResponse, error) {
	orders, _ := s.Repo.ListOrders()
	var res []*adminpb.Order
	for _, o := range orders {
		res = append(res, &adminpb.Order{
			OrderId:     o.ID,
			UserId:      o.UserID,
			Status:      o.OrderStatus,
			TotalAmount: o.TotalAmount,
			CreatedAt:   o.CreatedAt,
		})
	}
	return &adminpb.ListOrdersResponse{Orders: res}, nil
}

func (s *AdminService) GetRecentActivity(ctx context.Context, _ *adminpb.Empty) (*adminpb.RecentActivityResponse, error) {
	activities := s.Repo.GetRecentActivity()

	var res []*adminpb.ActivityItem
	for _, a := range activities {
		res = append(res, &adminpb.ActivityItem{
			Type:      a.Type,
			Message:   a.Message,
			Timestamp: a.Timestamp,
		})
	}

	return &adminpb.RecentActivityResponse{Activities: res}, nil
}

func (s *AdminService) DeleteUser(ctx context.Context, req *adminpb.DeleteUserRequest) (*adminpb.DeleteUserResponse, error) {
	userID := req.GetUserId()

	// Step 1: Delete from local DB (auth DB)
	err := s.Repo.DeleteUser(userID)
	if err != nil {
		return nil, err
	}
	return &adminpb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
