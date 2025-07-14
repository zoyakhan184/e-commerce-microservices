package handlers

import (
	"context"
	"order-service/models"
	orderpb "order-service/proto"
	"order-service/rabbitmq"
	"order-service/repository"
	"order-service/utils"

	"github.com/google/uuid"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
	Repo *repository.OrderRepo
}

func (s *OrderService) PlaceOrder(ctx context.Context, req *orderpb.PlaceOrderRequest) (*orderpb.PlaceOrderResponse, error) {
	order := models.Order{
		ID:            uuid.New().String(),
		UserID:        req.UserId,
		OrderStatus:   "pending",
		PaymentStatus: "unpaid",
		TotalAmount:   0, // Replace after integrating cart service
	}

	err := s.Repo.CreateOrder(&order)
	if err == nil {
		rabbitmq.EmitOrderPlaced(order)
	}

	return &orderpb.PlaceOrderResponse{
		OrderId: order.ID,
		Status:  "created",
	}, err
}

func (s *OrderService) GetOrders(ctx context.Context, req *orderpb.GetOrdersRequest) (*orderpb.GetOrdersResponse, error) {
	orders, _ := s.Repo.GetOrdersByUser(req.UserId)
	var res []*orderpb.Order
	for _, o := range orders {
		res = append(res, &orderpb.Order{
			Id:            o.ID,
			UserId:        o.UserID,
			OrderStatus:   o.OrderStatus,
			PaymentStatus: o.PaymentStatus,
			TotalAmount:   o.TotalAmount,
			CreatedAt:     o.CreatedAt.String(),
		})
	}
	return &orderpb.GetOrdersResponse{Orders: res}, nil
}

func (s *OrderService) GetOrderDetails(ctx context.Context, req *orderpb.GetOrderDetailsRequest) (*orderpb.GetOrderDetailsResponse, error) {
	o, _ := s.Repo.GetOrderByID(req.OrderId)
	return &orderpb.GetOrderDetailsResponse{
		Order: &orderpb.Order{
			Id:            o.ID,
			UserId:        o.UserID,
			OrderStatus:   o.OrderStatus,
			PaymentStatus: o.PaymentStatus,
			TotalAmount:   o.TotalAmount,
			CreatedAt:     o.CreatedAt.String(),
		},
	}, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.UpdateOrderStatusResponse, error) {
	err := s.Repo.UpdateOrderStatus(req.OrderId, req.Status)
	return &orderpb.UpdateOrderStatusResponse{Status: "updated"}, err
}

func (s *OrderService) GenerateInvoice(ctx context.Context, req *orderpb.GenerateInvoiceRequest) (*orderpb.GenerateInvoiceResponse, error) {
	order, err := s.Repo.GetOrderWithItems(req.OrderId)
	if err != nil {
		return nil, err
	}

	invoice := utils.GenerateInvoice(order)

	return &orderpb.GenerateInvoiceResponse{
		InvoiceText: invoice,
	}, nil
}
