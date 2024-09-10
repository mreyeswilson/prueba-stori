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

	for _, t := range transactions {
		totalBalance += t.Value

		// Calcular totales para créditos y débitos
		if t.Value > 0 {
			creditSum += t.Value
			creditCount++
		} else {
			debitSum += t.Value
			debitCount++
		}
	}

	trxToMap := func() []map[string]interface{} {
		trxMap := []map[string]interface{}{}
		for _, t := range transactions {
			trxMap = append(trxMap, map[string]interface{}{
				"id":    t.ID,
				"date":  t.Date.Format("2006-01-02 15:04"),
				"value": fmt.Sprintf("$%.2f", t.Value),
			})
		}
		return trxMap
	}

	summary := models.Summary{
		TotalBalance: fmt.Sprintf("$%.2f", totalBalance),
		CreditSum:    fmt.Sprintf("$%.2f", creditSum),
		DebitSum:     fmt.Sprintf("$%.2f", debitSum),
		Transactions: trxToMap(),
	}

	return summary, nil
}
