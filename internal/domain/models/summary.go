package models

type Summary struct {
	TotalBalance string                   `json:"total_balance"`
	CreditSum    string                   `json:"credit_sum"`
	DebitSum     string                   `json:"debit_sum"`
	Transactions []map[string]interface{} `json:"transactions"`
}
