package repositories

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nedaZarei/arcaptcha-internship-2025/neda-arcaptcha-internship-2025/internal/models"
)

const (
	CREATE_PAYMENTS_TABLE = `CREATE TABLE IF NOT EXISTS payments(
		id SERIAL PRIMARY KEY,
		bill_id INTEGER REFERENCES bills(id) ON DELETE CASCADE,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		amount DECIMAL(12, 2) NOT NULL,
		paid_at TIMESTAMP WITH TIME ZONE,
		payment_status VARCHAR(50) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment models.Payment) (int, error)
	GetPaymentByID(id int) (*models.Payment, error)
	GetPaymentByBillAndUser(billID, userID int) (*models.Payment, error)
	GetPaymentsByUser(userID int) ([]models.Payment, error)
	GetPendingPaymentsByUser(userID int) ([]models.Payment, error)
	GetPaymentsByBill(billID int) ([]models.Payment, error)
	UpdatePaymentStatus(ctx context.Context, payment models.Payment) error
	UpdatePaymentsStatus(ctx context.Context, payments []models.Payment) error
	DeletePayment(id int) error
}

type paymentRepositoryImpl struct {
	db *sqlx.DB
}

func NewPaymentRepository(autoCreate bool, db *sqlx.DB) PaymentRepository {
	if autoCreate {
		if _, err := db.Exec(CREATE_PAYMENTS_TABLE); err != nil {
			log.Fatalf("failed to create payments table: %v", err)
		}
	}
	return &paymentRepositoryImpl{db: db}
}

func (r *paymentRepositoryImpl) CreatePayment(ctx context.Context, payment models.Payment) (int, error) {
	query := `INSERT INTO payments (bill_id, user_id, amount, paid_at, payment_status) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING id`
	var id int
	if err := r.db.QueryRowContext(ctx, query,
		payment.BillID,
		payment.UserID,
		payment.Amount,
		payment.PaidAt,
		payment.PaymentStatus).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *paymentRepositoryImpl) GetPaymentByID(id int) (*models.Payment, error) {
	var payment models.Payment
	query := `SELECT id, bill_id, user_id, amount, paid_at, payment_status, created_at, updated_at 
			  FROM payments WHERE id = $1`
	err := r.db.Get(&payment, query, id)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepositoryImpl) GetPaymentByBillAndUser(billID, userID int) (*models.Payment, error) {
	var payment models.Payment
	query := `SELECT id, bill_id, user_id, amount, paid_at, payment_status, created_at, updated_at 
			  FROM payments WHERE bill_id = $1 AND user_id = $2`
	err := r.db.Get(&payment, query, billID, userID)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepositoryImpl) GetPaymentsByUser(userID int) ([]models.Payment, error) {
	var payments []models.Payment
	query := `SELECT id, bill_id, user_id, amount, paid_at, payment_status, created_at, updated_at 
			  FROM payments WHERE user_id = $1`
	err := r.db.Select(&payments, query, userID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepositoryImpl) GetPendingPaymentsByUser(userID int) ([]models.Payment, error) {
	var payments []models.Payment
	query := `SELECT id, bill_id, user_id, amount, paid_at, payment_status, created_at, updated_at 
			  FROM payments WHERE user_id = $1 and payment_status = 'pending'`
	err := r.db.Select(&payments, query, userID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepositoryImpl) GetPaymentsByBill(billID int) ([]models.Payment, error) {
	var payments []models.Payment
	query := `SELECT id, bill_id, user_id, amount, paid_at, payment_status, created_at, updated_at 
			  FROM payments WHERE bill_id = $1`
	err := r.db.Select(&payments, query, billID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepositoryImpl) UpdatePaymentStatus(ctx context.Context, payment models.Payment) error {
	query := `UPDATE payments SET 
			  payment_status = :payment_status,
			  paid_at = :paid_at,
			  updated_at = CURRENT_TIMESTAMP
			  WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, payment)
	return err
}

func (r *paymentRepositoryImpl) UpdatePaymentsStatus(ctx context.Context, payments []models.Payment) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	query := `UPDATE payments SET 
			  payment_status = :payment_status,
			  paid_at = :paid_at,
			  updated_at = CURRENT_TIMESTAMP
			  WHERE id = :id`

	for _, payment := range payments {
		_, err = tx.NamedExecContext(ctx, query, payment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *paymentRepositoryImpl) DeletePayment(id int) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
