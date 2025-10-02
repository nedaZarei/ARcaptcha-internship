package repositories

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nedaZarei/arcaptcha-internship-2025/neda-arcaptcha-internship-2025/internal/models"
)

const (
	CREATE_BILLS_TABLE = `CREATE TABLE IF NOT EXISTS bills(
		id SERIAL PRIMARY KEY,
        apartment_id INTEGER NOT NULL REFERENCES apartments(id),
        bill_type VARCHAR(50) NOT NULL,
        total_amount DECIMAL(10,2) NOT NULL,
        due_date DATE NOT NULL,
        billing_deadline DATE,
        description TEXT,
        image_url VARCHAR(2000),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
)

type BillRepository interface {
	CreateBill(ctx context.Context, bill models.Bill) (int, error)
	GetBillByID(id int) (*models.Bill, error)
	GetBillsByApartmentID(apartmentID int) ([]models.Bill, error)
	UpdateBill(ctx context.Context, bill models.Bill) error
	DeleteBill(id int) error
	GetPaymentByBillAndUser(billID, userID int) (*models.Payment, error)
	GetUndividedBillsByTypeAndApartment(apartmentID int, billType models.BillType) ([]models.Bill, error)
	GetUndividedBillsByApartment(apartmentID int) ([]models.Bill, error)
}

type billRepositoryImpl struct {
	db *sqlx.DB
}

func NewBillRepository(autoCreate bool, db *sqlx.DB) BillRepository {
	if autoCreate {
		if _, err := db.Exec(CREATE_BILLS_TABLE); err != nil {
			log.Fatalf("failed to create bills table: %v", err)
		}
	}
	return &billRepositoryImpl{db: db}
}

func (r *billRepositoryImpl) CreateBill(ctx context.Context, bill models.Bill) (int, error) {
	query := `INSERT INTO bills (apartment_id, bill_type, total_amount, due_date, billing_deadline, description, image_url)
 				VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query,
		bill.ApartmentID,
		bill.BillType,
		bill.TotalAmount,
		bill.DueDate,
		bill.BillingDeadline,
		bill.Description,
		bill.ImageURL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *billRepositoryImpl) GetBillByID(id int) (*models.Bill, error) {
	var bill models.Bill
	query := `SELECT id, apartment_id, bill_type, total_amount, due_date, billing_deadline, description, image_url, created_at, updated_at 
			  FROM bills WHERE id = $1`
	err := r.db.Get(&bill, query, id)
	if err != nil {
		return nil, err
	}
	return &bill, nil
}

func (r *billRepositoryImpl) GetBillsByApartmentID(apartmentID int) ([]models.Bill, error) {
	var bills []models.Bill
	query := `SELECT id, apartment_id, bill_type, total_amount, due_date, billing_deadline, description, image_url, created_at, updated_at 
			  FROM bills WHERE apartment_id = $1`
	err := r.db.Select(&bills, query, apartmentID)
	if err != nil {
		return nil, err
	}
	return bills, nil
}

func (r *billRepositoryImpl) UpdateBill(ctx context.Context, bill models.Bill) error {
	query := `UPDATE bills
				SET apartment_id = $1, bill_type = $2, total_amount = $3,
				due_date = $4, billing_deadline = $5, description = $6,
				updated_at = CURRENT_TIMESTAMP
				WHERE id = $7`
	_, err := r.db.ExecContext(ctx, query,
		bill.ApartmentID,
		bill.BillType,
		bill.TotalAmount,
		bill.DueDate,
		bill.BillingDeadline,
		bill.Description,
		bill.ImageURL,
		bill.ID)
	return err
}

func (r *billRepositoryImpl) DeleteBill(id int) error {
	query := `DELETE FROM bills WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *billRepositoryImpl) GetPaymentByBillAndUser(billID, userID int) (*models.Payment, error) {
	var payment models.Payment
	query := `SELECT id, bill_id, user_id, amount, paid_at, payment_status, created_at, updated_at 
              FROM payments WHERE bill_id = $1 AND user_id = $2`
	err := r.db.Get(&payment, query, billID, userID)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *billRepositoryImpl) GetUndividedBillsByTypeAndApartment(apartmentID int, billType models.BillType) ([]models.Bill, error) {
	query := `
    SELECT b.id, b.apartment_id, b.bill_type, b.total_amount, b.due_date,
           b.billing_deadline, b.description, b.image_url, b.created_at, b.updated_at
    FROM bills b
    WHERE b.apartment_id = $1 
      AND b.bill_type = $2
      AND NOT EXISTS (
          SELECT 1 FROM payments p WHERE p.bill_id = b.id
      )
    ORDER BY b.created_at ASC`

	rows, err := r.db.Query(query, apartmentID, billType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bills []models.Bill
	for rows.Next() {
		var bill models.Bill
		err := rows.Scan(
			&bill.ID, &bill.ApartmentID, &bill.BillType, &bill.TotalAmount,
			&bill.DueDate, &bill.BillingDeadline, &bill.Description,
			&bill.ImageURL, &bill.CreatedAt, &bill.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}
	return bills, nil
}

// gets all bills that don't have payment records yet
func (r *billRepositoryImpl) GetUndividedBillsByApartment(apartmentID int) ([]models.Bill, error) {
	query := `
    SELECT b.id, b.apartment_id, b.bill_type, b.total_amount, b.due_date,
           b.billing_deadline, b.description, b.image_url, b.created_at, b.updated_at
    FROM bills b
    WHERE b.apartment_id = $1
      AND NOT EXISTS (
          SELECT 1 FROM payments p WHERE p.bill_id = b.id
      )
    ORDER BY b.created_at ASC`

	rows, err := r.db.Query(query, apartmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bills []models.Bill
	for rows.Next() {
		var bill models.Bill
		err := rows.Scan(
			&bill.ID, &bill.ApartmentID, &bill.BillType, &bill.TotalAmount,
			&bill.DueDate, &bill.BillingDeadline, &bill.Description,
			&bill.ImageURL, &bill.CreatedAt, &bill.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}
	return bills, nil
}
