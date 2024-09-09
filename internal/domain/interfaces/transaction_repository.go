package interfaces

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/mreyeswilson/prueba_stori/internal/domain/models"
)

type ITransactionRepository interface {
	GetSummary(event events.S3EventRecord) (models.Summary, error)
}
