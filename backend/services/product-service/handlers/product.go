package handlers

import (
	"context"
	// "encoding/base64"
	"fmt"
	// "os"
	// "path/filepath"
	"product-service/models"
	productpb "product-service/proto"
	"product-service/repository"
	"time"

	"github.com/google/uuid"
)

type ProductService struct {
	Repo repository.ProductRepository
	productpb.UnimplementedProductServiceServer
}

func (s *ProductService) AddProduct(ctx context.Context, req *productpb.ProductRequest) (*productpb.GenericResponse, error) {
	product := &models.Product{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
		CategoryId:  req.CategoryId,
		Price:       req.Price,
		Brand:       req.Brand,
		CreatedAt:   time.Now(),
	}
	fmt.Println("Adding image URL before", req.ImageUrls)
	for _, url := range req.ImageUrls {
		fmt.Println("Adding image URL:", url)
		product.Images = append(product.Images, models.ProductImage{
			ID:        uuid.NewString(),
			ImageURL:  url,
			ProductID: product.ID,
		})
	}

	err := s.Repo.AddProduct(product)
	if err != nil {
		return &productpb.GenericResponse{Message: "Failed to add product"}, err
	}
	return &productpb.GenericResponse{Message: "Product added successfully"}, nil
}

func (s *ProductService) EditProduct(ctx context.Context, req *productpb.ProductUpdateRequest) (*productpb.GenericResponse, error) {
	product := &models.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		CategoryId: req.CategoryId,
		Price:       req.Price,
		Brand:       req.Brand,
	}
	err := s.Repo.UpdateProduct(product)
	if err != nil {
		return &productpb.GenericResponse{Message: "Update failed"}, err
	}
	return &productpb.GenericResponse{Message: "Product updated"}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *productpb.ProductIdRequest) (*productpb.GenericResponse, error) {
	err := s.Repo.DeleteProduct(req.Id)
	if err != nil {
		return &productpb.GenericResponse{Message: "Delete failed"}, err
	}
	return &productpb.GenericResponse{Message: "Product deleted"}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *productpb.ProductIdRequest) (*productpb.ProductResponse, error) {
	product, err := s.Repo.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}

	image_urls := []string{}
	for _, img := range product.Images {
		dataURI := fmt.Sprintf("data:image/jpeg;base64,%s", img.ImageURL)
		image_urls = append(image_urls, dataURI)
	}

	return &productpb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Brand:       product.Brand,
		ImageUrls:   image_urls,
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *productpb.ProductFilter) (*productpb.ProductList, error) {
	filters := make(map[string]interface{})
	if req.CategoryId != "" {
		filters["category_id"] = req.CategoryId
	}
	if req.Brand != "" {
		filters["brand"] = req.Brand
	}

	products, err := s.Repo.ListProducts(filters)
	if err != nil {
		return nil, err
	}

	var list []*productpb.ProductResponse
	for _, p := range products {
		var urls = []string{}
		for _, img := range p.Images {
			dataURI := fmt.Sprintf("data:image/jpeg;base64,%s", img.ImageURL)
			urls = append(urls, dataURI)
			// urls = append(urls, img.ImageURL)
		}
		list = append(list, &productpb.ProductResponse{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			CategoryId:  p.CategoryId,
			Price:       p.Price,
			Brand:       p.Brand,
			ImageUrls:  urls,
		})
	}

	return &productpb.ProductList{Products: list}, nil
}

func (s *ProductService) AddCategory(ctx context.Context, req *productpb.CategoryRequest) (*productpb.GenericResponse, error) {
	var parentID *string
	if req.ParentId != "" {
		parentID = &req.ParentId
	}

	category := &models.Category{
		ID:       uuid.NewString(),
		Name:     req.Name,
		Gender:   req.Gender,
		ParentID: parentID,
	}

	err := s.Repo.AddCategory(category)
	if err != nil {
		return &productpb.GenericResponse{Message: "Failed to add category"}, err
	}
	return &productpb.GenericResponse{Message: "Category added"}, nil
}

func (s *ProductService) ListCategories(ctx context.Context, _ *productpb.Empty) (*productpb.CategoryList, error) {
	cats, err := s.Repo.ListCategories()
	if err != nil {
		return nil, err
	}

	var res []*productpb.CategoryResponse
	for _, c := range cats {
		var pid string
		if c.ParentID != nil {
			pid = *c.ParentID
		}
		res = append(res, &productpb.CategoryResponse{
			Id:       c.ID,
			Name:     c.Name,
			Gender:   c.Gender,
			ParentId: pid,
		})
	}

	return &productpb.CategoryList{Categories: res}, nil
}
