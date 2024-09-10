package interfaces

type ISenderService interface {
	SendEmail(from string, subject string, data *map[string]interface{}) error
}
