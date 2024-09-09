package repositories

import (
	"io"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mreyeswilson/prueba_stori/internal/domain/interfaces"
	"github.com/mreyeswilson/prueba_stori/internal/domain/models"
)

type TransactionRepository struct {
	Calculator interfaces.ICalculatorService
	S3Adapter  interfaces.IStorageAdapter
}

func NewTransactionRepository(
	calculatorService interfaces.ICalculatorService,
	storageAdapter interfaces.IStorageAdapter,
) interfaces.ITransactionRepository {
	return &TransactionRepository{
		Calculator: calculatorService,
		S3Adapter:  storageAdapter,
	}
}

func (t *TransactionRepository) GetSummary(event events.S3EventRecord) (models.Summary, error) {

	bucket := event.S3.Bucket.Name
	key := event.S3.Object.Key

	body, err := t.S3Adapter.GetObject(bucket, key)

	if err != nil {
		return models.Summary{}, err
	}

	defer body.Close()

	reader := io.Reader(body)

	return t.Calculator.MakeSummary(&reader)
}
