package reader

import (
	"file-reader-challenge/models"

	"github.com/jackc/pgx/v5"
)

type TransactionReader interface {
	ReadTransactions(conn *pgx.Conn, fileName string) ([]models.Transaction, error)
}
