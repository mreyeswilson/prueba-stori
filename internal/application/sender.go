package services

import (
	"github.com/hoisie/mustache"
	"github.com/mreyeswilson/prueba_stori/internal/domain/interfaces"
)

type SenderService struct {
	SenderAdapter interfaces.ISenderAdapter
}

func NewSenderService(
	sender_adapter interfaces.ISenderAdapter,
) interfaces.ISenderService {
	return &SenderService{
		SenderAdapter: sender_adapter,
	}
}

func (s *SenderService) SendEmail(from string, subject string, data *map[string]interface{}) error {
	identities, err := s.SenderAdapter.GetIdentities()

	if err != nil {
		return err
	}

	template := s.SenderAdapter.GetTemplate("TransactionSummaryTemplate")

	renderedTemplate := mustache.Render(template, data)

	return s.SenderAdapter.SendEmail(from, identities, subject, renderedTemplate)
}
