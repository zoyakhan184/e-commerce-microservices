package handlers

import (
	"context"
	"review-service/models"
	reviewpb "review-service/proto"
	"review-service/repository"

	"github.com/google/uuid"
)

type ReviewService struct {
	reviewpb.UnimplementedReviewServiceServer
	Repo *repository.ReviewRepo
}

func (s *ReviewService) AddReview(ctx context.Context, req *reviewpb.AddReviewRequest) (*reviewpb.GenericResponse, error) {
	review := models.Review{
		ID:        uuid.New().String(),
		UserID:    req.UserId,
		ProductID: req.ProductId,
		Rating:    int(req.Rating),
		Comment:   req.Comment,
	}
	err := s.Repo.AddReview(&review)
	return &reviewpb.GenericResponse{Status: "created"}, err
}

func (s *ReviewService) GetReviews(ctx context.Context, req *reviewpb.GetReviewsRequest) (*reviewpb.GetReviewsResponse, error) {
	reviews, _ := s.Repo.GetReviews(req.ProductId)
	var res []*reviewpb.Review
	for _, r := range reviews {
		res = append(res, &reviewpb.Review{
			Id:        r.ID,
			UserId:    r.UserID,
			ProductId: r.ProductID,
			Rating:    int32(r.Rating),
			Comment:   r.Comment,
		})
	}
	return &reviewpb.GetReviewsResponse{Reviews: res}, nil
}

func (s *ReviewService) EditReview(ctx context.Context, req *reviewpb.EditReviewRequest) (*reviewpb.GenericResponse, error) {
	err := s.Repo.UpdateReview(req.ReviewId, int(req.Rating), req.Comment)
	return &reviewpb.GenericResponse{Status: "updated"}, err
}

func (s *ReviewService) DeleteReview(ctx context.Context, req *reviewpb.DeleteReviewRequest) (*reviewpb.GenericResponse, error) {
	err := s.Repo.DeleteReview(req.ReviewId)
	return &reviewpb.GenericResponse{Status: "deleted"}, err
}
