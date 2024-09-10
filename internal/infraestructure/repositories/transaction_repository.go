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
	Sender     interfaces.ISenderService
}

func NewTransactionRepository(
	storageAdapter interfaces.IStorageAdapter,
	calculatorService interfaces.ICalculatorService,
	senderService interfaces.ISenderService,
) interfaces.ITransactionRepository {
	return &TransactionRepository{
		Calculator: calculatorService,
		S3Adapter:  storageAdapter,
		Sender:     senderService,
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

	summary, err := t.Calculator.MakeSummary(&reader)

	if err != nil {
		return models.Summary{}, err
	}

	t.Sender.SendEmail(
		"test@devwil.com",
		"Summary",
		&map[string]interface{}{
			"customer_name": "John Doe",
			"transactions":  summary.Transactions,
			"total_balance": summary.TotalBalance,
			"total_credits": summary.CreditSum,
			"total_debits":  summary.DebitSum,
			"company_name":  "Stori Test",
		},
	)

	return summary, nil
}