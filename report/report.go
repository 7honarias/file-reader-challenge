package report

import (
	"file-reader-challenge/models"
	"file-reader-challenge/errors"
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

func GenerateReportData(totalBalance float64, summary map[string]map[string]float64) models.ReportData {
	var report models.ReportData
	report.TotalBalance = totalBalance

	for month, data := range summary {
		numTransactions := int(data["creditCount"] + data["debitCount"])
		avgCredit := 0.0
		avgDebit := 0.0

		if data["creditCount"] > 0 {
			avgCredit = data["credits"] / data["creditCount"]
		}

		if data["debitCount"] > 0 {
			avgDebit = data["debits"] / data["debitCount"]
		}

		monthlySummary := models.MonthlySummary{
			Month:           month,
			NumTransactions: numTransactions,
			AverageCredit:   avgCredit,
			AverageDebit:    avgDebit,
		}

		report.Summary = append(report.Summary, monthlySummary)
	}

	return report
}

func GenerateReportHtml(reporData models.ReportData) (string, error) {
	htmlFilePath := filepath.Join("", "email_template.html")
	htmlContent, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		return "", fmt.Errorf(errors.HTMLTemplateFileNotFount, err)
	}
	var body bytes.Buffer

	htmlContent.Execute(&body, reporData)
	return body.String(), nil
}
