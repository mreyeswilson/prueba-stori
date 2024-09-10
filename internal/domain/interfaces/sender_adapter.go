package interfaces

type ISenderAdapter interface {
	SendEmail(from string, to []*string, subject string, html string) error
	GetIdentities() ([]*string, error)
	GetTemplate(templateName string) string
}
