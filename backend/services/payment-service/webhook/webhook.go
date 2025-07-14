package webhook

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"payment-service/rabbitmq"
	"payment-service/repository"
	"time"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/webhook"
)

func StartStripeWebhookServer(repo *repository.PaymentRepo) {
	http.HandleFunc("/stripe/webhook", func(w http.ResponseWriter, r *http.Request) {
		const MaxBodyBytes = int64(65536)
		r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Request body read error", http.StatusServiceUnavailable)
			return
		}

		endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
		event, err := webhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), endpointSecret)
		if err != nil {
			http.Error(w, "Signature verification failed", http.StatusBadRequest)
			log.Println("Signature error:", err)
			return
		}

		switch event.Type {
		case "checkout.session.completed":
			var session stripe.CheckoutSession
			err := json.Unmarshal(event.Data.Raw, &session)
			if err != nil {
				http.Error(w, "JSON parse error", http.StatusBadRequest)
				return
			}

			orderID := session.ClientReferenceID
			txnRef := session.PaymentIntent.ID

			payment, _ := repo.GetByOrderID(orderID)
			repo.UpdateStatus(payment.ID, "success", txnRef)

			rabbitmq.EmitPaymentEvent("payment.success", map[string]interface{}{
				"payment_id": payment.ID,
				"txn_ref":    txnRef,
				"user_id":    payment.UserID,
				"order_id":   payment.OrderID,
				"timestamp":  time.Now().Format(time.RFC3339),
			})
			log.Println("✅ Stripe payment success for order:", orderID)
		}

		w.WriteHeader(http.StatusOK)
	})

	go func() {
		port := os.Getenv("STRIPE_WEBHOOK_PORT")
		if port == "" {
			port = "50063"
		}
		log.Println("🔔 Stripe webhook listening on port", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Webhook server failed:", err)
		}
	}()
}
