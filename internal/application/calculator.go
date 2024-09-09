package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	models "github.com/mreyeswilson/prueba_stori/internal/domain"
)

type CalculatorService struct{}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (c *CalculatorService) ParseInfo(reader *csv.Reader) []models.Transaction {
	transactions := []models.Transaction{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("failed to read record: %v", err)
		}

		id := row[0]
		date, _ := time.Parse("2/1", row[1])
		value, _ := strconv.ParseFloat(row[2], 64)

		transactions = append(transactions, models.Transaction{
			ID:    id,
			Date:  date,
			Value: value,
		})
	}
	return transactions
}

func (c *CalculatorService) MakeSummary(transactions []models.Transaction) models.Summary {
	var totalBalance float64
	creditSum, debitSum := 0.0, 0.0
	creditCount, debitCount := 0, 0
	transactionsByMonth := make(map[string]int)

	for _, t := range transactions {
		totalBalance += t.Value

		// Agrupar transacciones por mes (usamos solo mes y año)
		monthYear := t.Date.Format("January 2006")
		transactionsByMonth[monthYear]++

		// Calcular totales para créditos y débitos
		if t.Value > 0 {
			creditSum += t.Value
			creditCount++
		} else {
			debitSum += t.Value
			debitCount++
		}
	}

	// Calcular promedios
	avgCredit := 0.0
	if creditCount > 0 {
		avgCredit = creditSum / float64(creditCount)
	}

	avgDebit := 0.0
	if debitCount > 0 {
		avgDebit = debitSum / float64(debitCount)
	}

	summary := models.Summary{
		TotalBalance:        totalBalance,
		CreditSum:           creditSum,
		DebitSum:            debitSum,
		AvgCredit:           avgCredit,
		AvgDebit:            avgDebit,
		CreditCount:         creditCount,
		DebitCount:          debitCount,
		TransactionsByMonth: transactionsByMonth,
	}

	return summary
}
