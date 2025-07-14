package handlers

import (
	"context"
	"time"

	"payment-service/gateway"
	"payment-service/models"
	paymentpb "payment-service/proto"
	"payment-service/rabbitmq"
	"payment-service/repository"

	"github.com/google/uuid"
)

type PaymentService struct {
	Repo *repository.PaymentRepo
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *PaymentService) InitiatePayment(ctx context.Context, req *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error) {
	id := uuid.New().String()

	payment := &models.Payment{
		ID:      id,
		OrderID: req.OrderId,
		UserID:  req.UserId,
		Gateway: req.Gateway,
		Status:  "pending",
	}

	_ = s.Repo.Create(payment)

	if req.Gateway == "stripe" {
		url, err := gateway.CreateStripeCheckout(req.OrderId, float64(req.Amount), req.Currency, req.SuccessUrl, req.CancelUrl)
		if err != nil {
			return nil, err
		}
		return &paymentpb.PaymentResponse{Url: url, Message: "Redirect to Stripe"}, nil
	}

	// Handle COD
	now := time.Now()
	_ = s.Repo.UpdateStatus(id, "success", "cod_ref_"+id)
	payment.Status = "success"
	payment.TxnRef = "cod_ref_" + id
	payment.PaidAt = &now

	// Emit event
	rabbitmq.EmitPaymentEvent("payment.success", map[string]interface{}{
		"payment_id": id,
		"txn_ref":    payment.TxnRef,
		"user_id":    payment.UserID,
		"order_id":   payment.OrderID,
		"timestamp":  now.Format(time.RFC3339),
	})

	return &paymentpb.PaymentResponse{Url: "", Message: "COD order confirmed"}, nil
}

func (s *PaymentService) VerifyPayment(ctx context.Context, req *paymentpb.VerifyRequest) (*paymentpb.PaymentStatus, error) {
	payment, _ := s.Repo.GetByID(req.PaymentId)
	return &paymentpb.PaymentStatus{Status: payment.Status, TxnRef: payment.TxnRef}, nil
}

func (s *PaymentService) RefundPayment(ctx context.Context, req *paymentpb.RefundRequest) (*paymentpb.RefundStatus, error) {
	payment, err := s.Repo.GetByOrderID(req.OrderId)
	if err != nil {
		return nil, err
	}

	if payment.Gateway == "stripe" {
		err := gateway.ProcessStripeRefund(payment.TxnRef, float64(req.Amount))
		if err != nil {
			return &paymentpb.RefundStatus{Status: "failed", Message: err.Error()}, nil
		}
		s.Repo.UpdateStatus(payment.ID, "refunded", payment.TxnRef)
		return &paymentpb.RefundStatus{Status: "success", Message: "Refunded via Stripe"}, nil
	}

	if payment.Gateway == "cod" {
		s.Repo.UpdateStatus(payment.ID, "refunded", payment.TxnRef)
		return &paymentpb.RefundStatus{Status: "success", Message: "Refund recorded for COD"}, nil
	}

	return &paymentpb.RefundStatus{Status: "failed", Message: "Unsupported gateway"}, nil
}
