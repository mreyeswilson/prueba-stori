package models

import "time"

type Transaction struct {
	ID    string    `json:"id"`
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}
