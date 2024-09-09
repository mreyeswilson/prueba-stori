package services_test

import (
	"io"
	"strings"
	"testing"

	services "github.com/mreyeswilson/prueba_stori/internal/application"
	"github.com/mreyeswilson/prueba_stori/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestCalculatorService(t *testing.T) {

	service := services.NewCalculatorService()

	t.Run("Should MakeSummary ok", func(t *testing.T) {

		csvData := `ID,Date,Value
		1,01/01,100
		2,02/01,-50`

		reader := io.Reader(strings.NewReader(csvData))

		summary, err := service.MakeSummary(&reader)

		assert.Nil(t, err)
		assert.IsType(t, models.Summary{}, summary)
		assert.Equal(t, 50.0, summary.TotalBalance)
		assert.Equal(t, 100.0, summary.CreditSum)

	})

	t.Run("Should MakeSumary failed", func(t *testing.T) {

		csvData := `ID,Date,Value
		1,01/01,,100`

		reader := io.Reader(strings.NewReader(csvData))

		_, err := service.MakeSummary(&reader)

		assert.NotNil(t, err)
	})

}
