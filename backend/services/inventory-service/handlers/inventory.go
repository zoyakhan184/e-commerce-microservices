package handlers

import (
	"context"
	inventorypb "inventory-service/proto"
	"inventory-service/repository"
	"inventory-service/rabbitmq" // ✅ add this

)

type InventoryService struct {
	Repo repository.InventoryRepository
	inventorypb.UnimplementedInventoryServiceServer
}

func (s *InventoryService) GetStock(ctx context.Context, req *inventorypb.StockRequest) (*inventorypb.StockResponse, error) {
	inv, err := s.Repo.GetStock(req.ProductId, req.Size, req.Color)
	if err != nil {
		return nil, err
	}
	return &inventorypb.StockResponse{
		SkuId:     inv.SkuID,
		ProductId: inv.ProductID,
		Size:      inv.Size,
		Color:     inv.Color,
		Quantity:  int32(inv.Quantity),
	}, nil
}

func (s *InventoryService) UpdateStockOnOrder(ctx context.Context, req *inventorypb.StockUpdateRequest) (*inventorypb.GenericResponse, error) {
	err := s.Repo.UpdateStock(req.SkuId, -int(req.QuantityChange))
	if err != nil {
		return &inventorypb.GenericResponse{Message: "Failed to update stock"}, err
	}

	// ✅ Check updated stock and emit low-stock alert if threshold is breached
	inv, err := s.Repo.GetBySKU(req.SkuId)
	if err == nil && inv.Quantity <= 5 {
		rabbitmq.EmitLowStockEvent(inv)
	}

	return &inventorypb.GenericResponse{Message: "Stock updated"}, nil
}

func (s *InventoryService) Restock(ctx context.Context, req *inventorypb.StockUpdateRequest) (*inventorypb.GenericResponse, error) {
	err := s.Repo.Restock(req.SkuId, int(req.QuantityChange))
	if err != nil {
		return &inventorypb.GenericResponse{Message: "Failed to restock"}, err
	}
	return &inventorypb.GenericResponse{Message: "Restocked successfully"}, nil
}

func (s *InventoryService) ListLowStock(ctx context.Context, req *inventorypb.LowStockRequest) (*inventorypb.LowStockList, error) {
	invs, err := s.Repo.ListLowStock(int(req.Threshold))
	if err != nil {
		return nil, err
	}

	var items []*inventorypb.StockResponse
	for _, inv := range invs {
		items = append(items, &inventorypb.StockResponse{
			SkuId:     inv.SkuID,
			ProductId: inv.ProductID,
			Size:      inv.Size,
			Color:     inv.Color,
			Quantity:  int32(inv.Quantity),
		})
	}
	return &inventorypb.LowStockList{Items: items}, nil
}
