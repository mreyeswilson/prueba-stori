package models

type Summary struct {
	TotalBalance float64
	CreditSum    float64
	DebitSum     float64
	AvgCredit    float64
	AvgDebit     float64
	CreditCount  int
	DebitCount   int
	TransactionsByMonth map[string]int
}