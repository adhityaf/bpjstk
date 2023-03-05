package repositories

import (
	"database/sql"

	"github.com/adhityaf/bpjstk/models"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) (error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) Create(transaction *models.Transaction) (error) {
	query := `INSERT INTO "transactions"("fullname", "quantity", "price", "created_at", "updated_at")
        VALUES($1, $2, $3, $4, $5)`

	statement, err := t.db.Prepare(query)
	defer statement.Close()

	if err != nil {
		return err
	}
	
	_, err = statement.Exec(transaction.Fullname, transaction.Quantity, transaction.Price, transaction.CreatedAt, transaction.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
