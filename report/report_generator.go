package report

import "file-reader-challenge/models"

type ReportGenerator interface {
	GenerateReportData(totalBalance float64, summary map[string]map[string]float64) models.ReportData
}
