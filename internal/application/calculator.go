package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/mreyeswilson/prueba_stori/internal/domain/interfaces"
	"github.com/mreyeswilson/prueba_stori/internal/domain/models"
)

type CalculatorService struct{}

func NewCalculatorService() interfaces.ICalculatorService {
	return &CalculatorService{}
}

func (c *CalculatorService) parseInfo(reader *io.Reader) ([]models.Transaction, error) {
	transactions := []models.Transaction{}

	csvReader := csv.NewReader(*reader)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			println(err)
			return transactions, fmt.Errorf("failed to read csv: %v", err)
		}

		id := row[0]
		date, _ := time.Parse("2006/01", row[1])
		value, _ := strconv.ParseFloat(row[2], 64)

		transactions = append(transactions, models.Transaction{
			ID:    id,
			Date:  date,
			Value: value,
		})
	}
	return transactions[1:], nil
}

func (c *CalculatorService) MakeSummary(reader *io.Reader) (models.Summary, error) {

	transactions, err := c.parseInfo(reader)

	if err != nil {
		return models.Summary{}, fmt.Errorf("failed to parse info: %v", err)
	}

	var totalBalance float64
	creditSum, debitSum := 0.0, 0.0
	creditCount, debitCount := 0, 0
	averageCredit, averageDebit := 0.0, 0.0
	transactionsByMonth := []models.MonthTransaction{}

	for _, t := range transactions {
		totalBalance += t.Value

		month := t.Date.Format("January, 2006")

		// Verificar si el mes ya existe en el slice
		found := false
		for i, mt := range transactionsByMonth {
			if mt.Month == month {
				transactionsByMonth[i].Count++
				found = true
				break
			}
		}

		// Si no existe, agregarlo
		if !found {
			transactionsByMonth = append(transactionsByMonth, models.MonthTransaction{
				Month: month,
				Count: 1,
			})
		}

		// Calcular totales para créditos y débitos
		if t.Value > 0 {
			creditSum += t.Value
			creditCount++
		} else {
			debitSum += t.Value
			debitCount++
		}

		// Calcular promedios para créditos y débitos

		if creditCount > 0 {
			averageCredit = creditSum / float64(creditCount)
		}

		if debitCount > 0 {
			averageDebit = debitSum / float64(debitCount)
		}
	}

	summary := models.Summary{
		TotalBalance:        fmt.Sprintf("$%.2f", totalBalance),
		CreditSum:           fmt.Sprintf("$%.2f", creditSum),
		CreditAvg:           fmt.Sprintf("$%.2f", averageCredit),
		DebitAvg:            fmt.Sprintf("$%.2f", averageDebit),
		DebitSum:            fmt.Sprintf("$%.2f", debitSum),
		TransactionsByMonth: transactionsByMonth,
	}

	return summary, nil
}
