package handlers

import (
	"context"
	"log"
	"time"
	"user-service/models"
	userpb "user-service/proto"
	"user-service/repository"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer
	Repo repository.UserRepository
}

// --- RPC Methods ---

func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.UserProfile, error) {
	log.Printf("GetUser called with user_id: %s", req.UserId)
	profile, err := s.Repo.GetProfile(req.UserId)
	if err != nil {
		log.Printf("Error fetching user profile for user_id %s: %v", req.UserId, err)
		return nil, err
	}
	log.Printf("Successfully fetched user profile for user_id: %s", req.UserId)
	return &userpb.UserProfile{
		UserId:    profile.UserID,
		FullName:  profile.FullName,
		Phone:     profile.Phone,
		Gender:    profile.Gender,
		Dob:       profile.DOB,
		AvatarUrl: profile.AvatarURL,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userpb.UserProfile) (*userpb.GenericResponse, error) {
	if req.UserId == "" {
		log.Printf("[UpdateUser] Empty user_id in request")
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	log.Printf("[UpdateUser] called with user_id: %s", req.UserId)

	profile := &models.Profile{
		UserID:    req.UserId,
		FullName:  req.FullName,
		Phone:     req.Phone,
		Gender:    req.Gender,
		DOB:       req.Dob,
		AvatarURL: req.AvatarUrl,
	}

	// Attempt to update, if no rows are affected, insert
	err := s.Repo.UpsertProfile(req.UserId, profile)
	if err != nil {
		log.Printf("Error upserting profile for user_id %s: %v", req.UserId, err)
		return nil, err
	}

	log.Printf("Successfully upserted profile for user_id: %s", req.UserId)
	return &userpb.GenericResponse{Message: "User profile saved"}, nil
}

func (s *UserService) AddAddress(ctx context.Context, req *userpb.AddressRequest) (*userpb.AddressResponse, error) {
	log.Printf("AddAddress called for user_id: %s", req.UserId)

	id := req.Id
	if id == "" {
		id = uuid.New().String()
	}

	address := &models.Address{
		ID:          id,
		UserID:      req.UserId,
		Name:        req.Name,
		Phone:       req.Phone,
		AddressLine: req.AddressLine,
		City:        req.City,
		State:       req.State,
		Zip:         req.Zip,
		Country:     req.Country,
		IsDefault:   req.IsDefault,
	}

	if address.IsDefault {
		s.Repo.ClearDefaultAddress(req.UserId) // remove old default
	}

	err := s.Repo.AddAddress(address)
	if err != nil {
		log.Printf("Error adding address for user_id %s: %v", req.UserId, err)
		return nil, err
	}

	return &userpb.AddressResponse{
		Id:          address.ID,
		UserId:      address.UserID,
		Name:        address.Name,
		Phone:       address.Phone,
		AddressLine: address.AddressLine,
		City:        address.City,
		State:       address.State,
		Zip:         address.Zip,
		Country:     address.Country,
		IsDefault:   address.IsDefault,
	}, nil
}

func (s *UserService) UpdateAddress(ctx context.Context, req *userpb.AddressRequest) (*userpb.GenericResponse, error) {
	log.Printf("UpdateAddress called for address_id: %s, user_id: %s", req.Id, req.UserId)
	err := s.Repo.UpdateAddress(&models.Address{
		ID:          req.Id,
		UserID:      req.UserId,
		AddressLine: req.AddressLine,
		City:        req.City,
		State:       req.State,
		Zip:         req.Zip,
		Country:     req.Country,
	})
	if err != nil {
		log.Printf("Error updating address for address_id %s: %v", req.Id, err)
		return nil, err
	}
	log.Printf("Successfully updated address for address_id: %s", req.Id)
	return &userpb.GenericResponse{Message: "Address updated"}, nil
}

func (s *UserService) GetAddresses(ctx context.Context, req *userpb.UserRequest) (*userpb.AddressList, error) {
	log.Printf("GetAddresses called for user_id: %s", req.UserId)
	addrs, err := s.Repo.GetAddresses(req.UserId)
	if err != nil {
		log.Printf("Error fetching addresses for user_id %s: %v", req.UserId, err)
		return nil, err
	}

	var protoAddrs []*userpb.AddressRequest
	for _, a := range addrs {
		protoAddrs = append(protoAddrs, &userpb.AddressRequest{
			Id:          a.ID,
			UserId:      a.UserID,
			AddressLine: a.AddressLine,
			City:        a.City,
			State:       a.State,
			Zip:         a.Zip,
			Country:     a.Country,
		})
	}
	log.Printf("Successfully fetched %d addresses for user_id: %s", len(protoAddrs), req.UserId)
	return &userpb.AddressList{Addresses: protoAddrs}, nil
}

func (s *UserService) AddToWishlist(ctx context.Context, req *userpb.WishlistRequest) (*userpb.GenericResponse, error) {
	if req.UserId == "" || req.ProductId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user_id and product_id are required")
	}

	log.Printf("AddToWishlist called with user_id: %s, product_id: %s", req.UserId, req.ProductId)

	if err := s.Repo.AddToWishlist(req.UserId, req.ProductId); err != nil {
		log.Printf("Error adding to wishlist: %v", err)
		return nil, err
	}

	log.Printf("Successfully added to wishlist: user_id=%s, product_id=%s", req.UserId, req.ProductId)
	return &userpb.GenericResponse{Message: "Added to wishlist"}, nil
}

func (s *UserService) RemoveFromWishlist(ctx context.Context, req *userpb.WishlistRequest) (*userpb.GenericResponse, error) {
	log.Printf("RemoveFromWishlist called with user_id: %s, product_id: %s", req.UserId, req.ProductId)
	err := s.Repo.RemoveFromWishlist(req.UserId, req.ProductId)
	if err != nil {
		log.Printf("Error removing from wishlist: %v", err)
		return nil, err
	}
	log.Printf("Successfully removed from wishlist: user_id=%s, product_id=%s", req.UserId, req.ProductId)
	return &userpb.GenericResponse{Message: "Removed from wishlist"}, nil
}

func (s *UserService) GetWishlist(ctx context.Context, req *userpb.UserRequest) (*userpb.WishlistResponse, error) {
	log.Printf("GetWishlist called for user_id: %s", req.UserId)
	products, err := s.Repo.GetWishlist(req.UserId)
	if err != nil {
		log.Printf("Error fetching wishlist for user_id %s: %v", req.UserId, err)
		return nil, err
	}
	log.Printf("Successfully fetched wishlist for user_id: %s", req.UserId)
	return &userpb.WishlistResponse{ProductIds: products}, nil
}

func (s *UserService) ListAllProfiles(ctx context.Context, _ *userpb.Empty) (*userpb.ListProfilesResponse, error) {
	log.Println("ListAllProfiles called")

	users, err := s.Repo.GetAllUsersWithProfiles()
	if err != nil {
		log.Printf("Error fetching users with profiles: %v", err)
		return nil, err
	}

	var protoProfiles []*userpb.UserProfile
	for _, u := range users {
		protoProfiles = append(protoProfiles, &userpb.UserProfile{
			UserId:    u.UserID,
			FullName:  u.FullName,
			Email:     u.Email,
			Phone:     u.Phone,
			Gender:    u.Gender,
			Dob:       u.DOB,
			AvatarUrl: u.AvatarURL,
			Role:      u.Role,
			CreatedAt: u.CreatedAt.Format(time.RFC3339), // ✅ Add this
		})
	}

	log.Printf("Successfully fetched %d user profiles", len(protoProfiles))
	return &userpb.ListProfilesResponse{Profiles: protoProfiles}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *userpb.UserProfile) (*userpb.GenericResponse, error) {
	log.Printf("CreateUser called with user_id: %s", req.UserId)

	err := s.Repo.CreateProfile(&models.Profile{
		UserID:    req.UserId,
		FullName:  req.FullName,
		Phone:     req.Phone,
		Gender:    req.Gender,
		DOB:       req.Dob,
		AvatarURL: req.AvatarUrl,
	})
	if err != nil {
		log.Printf("Error creating profile for user_id %s: %v", req.UserId, err)
		return nil, err
	}

	log.Printf("Successfully created profile for user_id: %s", req.UserId)
	return &userpb.GenericResponse{Message: "User profile created"}, nil
}
