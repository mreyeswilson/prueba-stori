package repositories_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	services "github.com/mreyeswilson/prueba_stori/internal/application"
	"github.com/mreyeswilson/prueba_stori/internal/domain/models"
	"github.com/mreyeswilson/prueba_stori/internal/infraestructure/repositories"
	"github.com/mreyeswilson/prueba_stori/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository(t *testing.T) {

	event := events.S3EventRecord{
		S3: events.S3Entity{
			Bucket: events.S3Bucket{
				Name: "test-bucket",
			},
			Object: events.S3Object{
				Key: "test-key",
			},
		},
	}

	
	t.Run("Should GetSummary ok", func(t *testing.T) {
		
		mockS3Adapter := new(mocks.IStorageAdapter)
		mockS3Adapter.On("GetObject", "test-bucket", "test-key").Return(
			io.NopCloser(strings.NewReader("ID,Date,Value\n1,2024/08,-300")), nil,
		)
		
		repo := repositories.NewTransactionRepository(
			services.NewCalculatorService(),
			mockS3Adapter,
		)
		
		summary, err := repo.GetSummary(event)
		
		assert.Nil(t, err)
		assert.IsType(t, models.Summary{}, summary)
		assert.Equal(t, -300.0, summary.TotalBalance)
		assert.Equal(t, 0.0, summary.CreditSum)
		
	})
	
	t.Run("Should GetSummary failed", func(t *testing.T) {
		
		mockS3Adapter := new(mocks.IStorageAdapter)
		mockS3Adapter.On("GetObject", "test-bucket", "test-key").Return(nil, errors.New("error"))

		repo := repositories.NewTransactionRepository(
			services.NewCalculatorService(),
			mockS3Adapter,
		)

		_, err := repo.GetSummary(event)

		assert.NotNil(t, err)
	})

}
