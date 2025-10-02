package payment

import (
	"context"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Payment interface {
	PayBills(paymentIDs []int, idempotentKey string) error
}

type paymentImpl struct {
	redis *redis.Client
}

func NewPayment(redis *redis.Client) Payment {
	return &paymentImpl{
		redis: redis,
	}
}

func (p *paymentImpl) PayBills(paymentIDs []int, idempotentKey string) error {
	for _, billID := range paymentIDs {
		key := billPaymentKey(billID, idempotentKey)
		if p.redis.Exists(context.Background(), key).Val() > 0 {
			log.Printf("Payment for bill %d with idempotent key %s already processed", billID, idempotentKey)
			continue
		}
		log.Printf("Processing payment for bill %d with idempotent key %s", billID, idempotentKey)

		//marking the payment as processed in redis
		if err := p.redis.Set(context.Background(), key, "processed", 0).Err(); err != nil {
			return err
		}
	}

	return nil
}

func billPaymentKey(paymentID int, idempotentKey string) string {
	return "bill_payment:" + strconv.Itoa(paymentID) + ":" + idempotentKey
}
