package repositories_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	services "github.com/mreyeswilson/prueba_stori/internal/application"
	"github.com/mreyeswilson/prueba_stori/internal/domain/models"
	adapters "github.com/mreyeswilson/prueba_stori/internal/infraestructure/aws"
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
			mockS3Adapter,
			services.NewCalculatorService(),
			services.NewSenderService(adapters.NewSESAdapter()),
		)

		summary, err := repo.GetSummary(event)

		assert.Nil(t, err)
		assert.IsType(t, models.Summary{}, summary)
		assert.Equal(t, "$-300.00", summary.TotalBalance)
		assert.Equal(t, "$0.00", summary.CreditSum)

	})

	t.Run("Should GetSummary failed", func(t *testing.T) {

		sender := "mreyeswilson@gmail.com"
		recipientList := []*string{aws.String(sender)}
		mockS3Adapter := new(mocks.IStorageAdapter)
		mockS3Adapter.On("GetObject", "test-bucket", "test-key").Return(nil, errors.New("error"))

		mockSESAdapter := new(mocks.ISenderAdapter)
		mockSESAdapter.On("SendEmail", "test@devwil.com", recipientList, "Summary", "<html></html>").Return(nil)
		mockSESAdapter.On("GetIdentities").Return(recipientList, nil)

		mockSESAdapter.On("GetTemplate", "TransactionSummaryTemplate").Return("<html></html>")

		repo := repositories.NewTransactionRepository(
			mockS3Adapter,
			services.NewCalculatorService(),
			services.NewSenderService(mockSESAdapter),
		)

		_, err := repo.GetSummary(event)

		assert.NotNil(t, err)
	})

	t.Run("Should CalculatorService failed", func(t *testing.T) {

		mockS3Adapter := new(mocks.IStorageAdapter)
		mockCalculatorService := new(mocks.ICalculatorService)
		mockReadCloser := io.NopCloser(strings.NewReader("ID,Date,Value\n1,2024/08,-300"))

		mockReader := io.Reader(mockReadCloser)

		mockS3Adapter.On("GetObject", "test-bucket", "test-key").Return(
			mockReadCloser, nil,
		)

		mockCalculatorService.On("MakeSummary", &mockReader).Return(models.Summary{}, errors.New("error"))

		repo := repositories.NewTransactionRepository(
			mockS3Adapter,
			mockCalculatorService,
			services.NewSenderService(adapters.NewSESAdapter()),
		)

		_, err := repo.GetSummary(event)

		assert.NotNil(t, err)

	})
}
