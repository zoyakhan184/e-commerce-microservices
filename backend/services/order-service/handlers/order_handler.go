package handlers

import (
	"context"
	"order-service/models"
	orderpb "order-service/proto"
	"order-service/rabbitmq"
	"order-service/repository"
	"order-service/utils"

	productpb "product-service/proto"

	"github.com/google/uuid"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
	Repo          *repository.OrderRepo
	ProductClient productpb.ProductServiceClient
}

func (s *OrderService) PlaceOrder(ctx context.Context, req *orderpb.PlaceOrderRequest) (*orderpb.PlaceOrderResponse, error) {
	orderID := uuid.New().String()

	var totalAmount float64
	var orderItems []models.OrderItem

	for _, item := range req.Items {
		itemID := uuid.New().String()

		// Fix: Use ProductIdRequest instead of GetProductRequest
		productResp, err := s.ProductClient.GetProduct(ctx, &productpb.ProductIdRequest{Id: item.ProductId})
		if err != nil {
			return nil, err
		}

		itemTotal := productResp.Price * float64(item.Quantity)
		totalAmount += itemTotal

		orderItems = append(orderItems, models.OrderItem{
			ID:          itemID,
			OrderID:     orderID,
			ProductID:   item.ProductId,
			ProductName: productResp.Name,
			Quantity:    int(item.Quantity),
			Price:       productResp.Price,
			// Remove Size and Color fields since they don't exist in ProductResponse
			// Size:        productResp.Size,
			// Color:       productResp.Color,
		})
	}

	order := models.Order{
		ID:            orderID,
		UserID:        req.UserId,
		OrderStatus:   "pending",
		PaymentStatus: "unpaid",
		TotalAmount:   totalAmount,
		OrderItems:    orderItems,
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
		var items []*orderpb.OrderItem
		for _, i := range o.OrderItems {
			items = append(items, &orderpb.OrderItem{
				ProductId:   i.ProductID,
				ProductName: i.ProductName,
				Quantity:    int32(i.Quantity),
				Price:       i.Price,
				// Remove Size and Color if they don't exist in your proto
				// Size:        i.Size,
				// Color:       i.Color,
			})
		}
		res = append(res, &orderpb.Order{
			Id:            o.ID,
			UserId:        o.UserID,
			OrderStatus:   o.OrderStatus,
			PaymentStatus: o.PaymentStatus,
			TotalAmount:   o.TotalAmount,
			CreatedAt:     o.CreatedAt.String(),
			Items:         items,
		})
	}
	return &orderpb.GetOrdersResponse{Orders: res}, nil
}

func (s *OrderService) GetOrderDetails(ctx context.Context, req *orderpb.GetOrderDetailsRequest) (*orderpb.GetOrderDetailsResponse, error) {
	o, _ := s.Repo.GetOrderByID(req.OrderId)
	var items []*orderpb.OrderItem
	for _, i := range o.OrderItems {
		items = append(items, &orderpb.OrderItem{
			ProductId:   i.ProductID,
			ProductName: i.ProductName,
			Quantity:    int32(i.Quantity),
			Price:       i.Price,
			Size:        i.Size,
			Color:       i.Color,
		})
	}
	return &orderpb.GetOrderDetailsResponse{
		Order: &orderpb.Order{
			Id:            o.ID,
			UserId:        o.UserID,
			OrderStatus:   o.OrderStatus,
			PaymentStatus: o.PaymentStatus,
			TotalAmount:   o.TotalAmount,
			CreatedAt:     o.CreatedAt.String(),
			Items:         items,
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

func (s *OrderService) ListAllOrders(ctx context.Context, _ *orderpb.ListAllOrdersRequest) (*orderpb.ListAllOrdersResponse, error) {
	orders, _ := s.Repo.GetAllOrders()
	var res []*orderpb.Order
	for _, o := range orders {
		var items []*orderpb.OrderItem
		for _, i := range o.OrderItems {
			items = append(items, &orderpb.OrderItem{
				ProductId:   i.ProductID,
				ProductName: i.ProductName,
				Quantity:    int32(i.Quantity),
				Price:       i.Price,
				Size:        i.Size,
				Color:       i.Color,
			})
		}
		res = append(res, &orderpb.Order{
			Id:            o.ID,
			UserId:        o.UserID,
			OrderStatus:   o.OrderStatus,
			PaymentStatus: o.PaymentStatus,
			TotalAmount:   o.TotalAmount,
			CreatedAt:     o.CreatedAt.String(),
			Items:         items,
		})
	}
	return &orderpb.ListAllOrdersResponse{Orders: res}, nil
}
