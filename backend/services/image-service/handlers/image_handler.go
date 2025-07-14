package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"image-service/models"
	"image-service/proto"
	"image-service/repository"

	"github.com/google/uuid"
)

type ImageService struct {
	proto.UnimplementedImageServiceServer
	Repo *repository.ImageRepo
}

func (s *ImageService) UploadImage(ctx context.Context, req *proto.UploadImageRequest) (*proto.UploadImageResponse, error) {
	id := uuid.New().String()
	ext := ".jpg" // You can enhance this by parsing from req.FileType
	fileName := id + ext
	filePath := filepath.Join("uploads", fileName)

	// Save to disk
	if err := os.WriteFile(filePath, req.ImageData, 0644); err != nil {
		log.Printf("❌ Failed to save image: %v", err)
		return nil, err
	}

	image := &models.Image{
		ID:         id,
		EntityID:   req.EntityId,
		EntityType: req.EntityType,
		FileType:   req.FileType,
		FileName:   fileName, // ✅ Save only filename
	}

	if err := s.Repo.Save(image); err != nil {
		log.Printf("❌ Failed to save image metadata: %v", err)
		return nil, err
	}

	log.Printf("✅ Uploaded image: ID=%s, saved to %s", id, filePath)
	return &proto.UploadImageResponse{ImageId: id}, nil
}

func (s *ImageService) GetImage(ctx context.Context, req *proto.GetImageRequest) (*proto.GetImageResponse, error) {
	img, err := s.Repo.Get(req.ImageId)
	if err != nil {
		log.Printf("❌ Failed to fetch image: %v", err)
		return nil, err
	}

	fullPath := filepath.Join("uploads", img.FileName)
	data, err := os.ReadFile(fullPath)
	fmt.Println("image data", data)
	if err != nil {
		log.Printf("❌ Failed to read image from disk: %v", err)
		return nil, err
	}

	return &proto.GetImageResponse{
		ImageData: data,
		FileType:  img.FileType,
	}, nil
}

func (s *ImageService) DeleteImage(ctx context.Context, req *proto.DeleteImageRequest) (*proto.DeleteImageResponse, error) {
	if err := s.Repo.Delete(req.ImageId); err != nil {
		log.Printf("❌ Failed to delete image: %v", err)
		return nil, err
	}

	return &proto.DeleteImageResponse{Status: "deleted"}, nil
}
