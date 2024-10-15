package reader

import (
	"context"
	"encoding/csv"
	"file-reader-challenge/errors"
	"file-reader-challenge/models"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

type TransactionReaderImpl struct{}

func (t *TransactionReaderImpl) ReadTransactions(conn *pgx.Conn, fileName string) ([]models.Transaction, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction

	for _, record := range records[1:] {
		id, _ := strconv.Atoi(record[0])
		amount, _ := strconv.ParseFloat(record[2], 64)

		transaction := models.Transaction{
			Id:          id,
			Date:        record[1],
			Transaction: amount,
		}

		_, err := conn.Exec(context.Background(),
			"INSERT INTO transaction (id, date, transaction) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING;",
			transaction.Id, transaction.Date, transaction.Transaction)
		if err != nil {
			fmt.Printf(errors.DBInsertionError, err)
			continue
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func CalculateSummary(transactions []models.Transaction) (float64, map[string]map[string]float64) {
	summary := make(map[string]map[string]float64)
	totalBalance := 0.0

	for _, transaction := range transactions {
		month := getMonth(transaction.Date)

		if _, exists := summary[month]; !exists {
			summary[month] = map[string]float64{
				"credits":     0,
				"debits":      0,
				"creditCount": 0,
				"debitCount":  0,
			}
		}

		if transaction.Transaction > 0 {
			summary[month]["credits"] += transaction.Transaction
			summary[month]["creditCount"]++
		} else {
			summary[month]["debits"] += transaction.Transaction
			summary[month]["debitCount"]++
		}

		totalBalance += transaction.Transaction
	}

	return totalBalance, summary
}

// getMonth devuelve el mes a partir de una fecha en formato YYYY-MM-DD.
func getMonth(date string) string {
	parts := strings.Split(date, "-")
	return parts[0]
}
