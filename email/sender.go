package email

import "file-reader-challenge/models"

type Sender interface {
	SendMailWithOAuth2(repo models.ReportData) error
}