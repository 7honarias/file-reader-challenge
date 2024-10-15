package main

import (
	"context"
	"log"

	"file-reader-challenge/db"
	"file-reader-challenge/email"
	"file-reader-challenge/errors"
	"file-reader-challenge/reader"
	"file-reader-challenge/report"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	db.InitDB()
}

func handleRequest(ctx context.Context, name string) (string, error) {
	var transactionReader reader.TransactionReader = &reader.TransactionReaderImpl{}
	var ReportGenerator report.ReportGenerator = &report.ReportGeneratorImpl{}
	var sender email.Sender = &email.OAuth2EmailSender{}


	transactions, err := transactionReader.ReadTransactions(db.Conn, name)
	if err != nil {
		log.Fatalf(errors.TransactionFileNotFound, err)
		return "", err
	}

	totalBalance, summary := reader.CalculateSummary(transactions)

	repo := ReportGenerator.GenerateReportData(totalBalance, summary)

	if err := sender.SendMailWithOAuth2(repo); err != nil {
		log.Fatalf(errors.EmailFaild, err)
		return "", err
	}

	return "Reporte enviado con éxito", nil
}

func main() {
	defer db.CloseDB()
	lambda.Start(handleRequest)
}
