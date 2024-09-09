package models

type Summary struct {
	TotalBalance        float64        `json:"total_balance"`
	CreditSum           float64        `json:"credit_sum"`
	DebitSum            float64        `json:"debit_sum"`
	AvgCredit           float64        `json:"avg_credit"`
	AvgDebit            float64        `json:"avg_debit"`
	CreditCount         int            `json:"credit_count"`
	DebitCount          int            `json:"debit_count"`
	TransactionsByMonth map[string]int `json:"transactions_by_month"`
}
