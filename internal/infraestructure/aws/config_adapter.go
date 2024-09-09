package adapters

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")), // Cambia esto a tu región
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}
	return sess, nil
}
