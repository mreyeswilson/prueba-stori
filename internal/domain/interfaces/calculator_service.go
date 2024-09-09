package interfaces

import (
	"io"

	"github.com/mreyeswilson/prueba_stori/internal/domain/models"
)

type ICalculatorService interface {
	MakeSummary(reader *io.Reader) (models.Summary, error)
}