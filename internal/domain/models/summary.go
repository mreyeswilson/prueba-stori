package models

type Summary struct {
	TotalBalance        string             `json:"total_balance"`
	CreditSum           string             `json:"credit_sum"`
	DebitSum            string             `json:"debit_sum"`
	CreditAvg           string             `json:"credit_avg"`
	DebitAvg            string             `json:"debit_avg"`
	TransactionsByMonth []MonthTransaction `json:"transactions_by_month"`
}

type MonthTransaction struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}
