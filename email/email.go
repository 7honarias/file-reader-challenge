package email

import (
	"file-reader-challenge/models"
	"file-reader-challenge/report"
	"fmt"
	"net/smtp"
	"os"
)

type OAuth2EmailSender struct{}

func (s *OAuth2EmailSender) SendMailWithOAuth2(repo models.ReportData) error {
	headers := fmt.Sprintf("MIME-version: 1.0;\r\n"+
		"Content-Type: text/html; charset=\"UTF-8\";\r\n"+
		"Subject: Stori challenge mail report\r\n"+
		"From: %s\r\n"+
		"To: %s\r\n", os.Getenv("FROM"), "c8251812@gmail.com")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	password := os.Getenv("PASSWORD")

	body, err := report.GenerateReportHtml(repo)
	if err != nil {
		return err
	}

	msg := []byte(headers + "\r\n" + body)

	auth := smtp.PlainAuth("", os.Getenv("FROM"), password, host)

	err = smtp.SendMail(fmt.Sprintf("%s:%s", host, port), auth, os.Getenv("FROM"), []string{"c8251812@gmail.com"}, msg)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
