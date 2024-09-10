package adapters

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	ses "github.com/aws/aws-sdk-go/service/ses"
)

type SESAdapter struct {
	Config *Config
}

func NewSESAdapter() *SESAdapter {
	return &SESAdapter{
		Config: NewConfig(),
	}
}

func (s *SESAdapter) GetSES() *ses.SES {

	session, err := s.Config.GetSession()

	if err != nil {
		log.Fatalf("failed to create session: %v", err)
		return nil
	}

	svc := ses.New(session)

	return svc
}

func (s *SESAdapter) GetTemplate(templateName string) string {
	session, err := s.Config.GetSession()

	if err != nil {
		log.Fatalf("failed to create session: %v", err)
		return ""
	}

	svc := ses.New(session)

	template, err := svc.GetTemplate(&ses.GetTemplateInput{
		TemplateName: aws.String(templateName),
	})

	if err != nil {
		log.Fatalf("failed to get template: %v", err)
		return ""
	}

	return *template.Template.HtmlPart
}

func (s *SESAdapter) SendEmail(from string, to []*string, subject string, html string) error {
	session, err := s.Config.GetSession()

	if err != nil {
		log.Fatalf("failed to create session: %v", err)
		return err
	}

	svc := ses.New(session)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: to,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(html),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(from),
	}

	_, err = svc.SendEmail(input)

	if err != nil {
		log.Fatalf("failed to send email: %v", err)
		return err
	}

	return nil
}

func (s *SESAdapter) GetIdentities() ([]*string, error) {
	session, err := s.Config.GetSession()

	if err != nil {
		log.Fatalf("failed to create session: %v", err)
		return nil, err
	}

	svc := ses.New(session)

	identities, err := svc.ListIdentities(&ses.ListIdentitiesInput{
		IdentityType: aws.String("EmailAddress"),
	})

	if err != nil {
		log.Fatalf("failed to get identities: %v", err)
		return nil, err
	}

	return identities.Identities, nil
}
