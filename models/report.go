package models

type MonthlySummary struct {
	Month              string
	NumTransactions    int
	AverageCredit      float64
	AverageDebit       float64
}

type ReportData struct {
	TotalBalance float64
	Summary      []MonthlySummary
}
