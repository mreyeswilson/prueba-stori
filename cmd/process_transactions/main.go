package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	services "github.com/mreyeswilson/prueba_stori/internal/application"
	adapters "github.com/mreyeswilson/prueba_stori/internal/infraestructure/aws"
	"github.com/mreyeswilson/prueba_stori/internal/infraestructure/repositories"
)

func lambdaHandler(ctx context.Context, event events.S3Event) (string, error) {
	// Crear una sesión de AWS

	repo := repositories.NewTransactionRepository(
		adapters.NewS3Adapter(),
		services.NewCalculatorService(),
		services.NewSenderService(adapters.NewSESAdapter()),
	)

	summary, err := repo.GetSummary(event.Records[0])

	if err != nil {
		log.Println("Error getting summary: ", err)
		return "", err
	}

	fmt.Println("Total Balance:", summary.TotalBalance)

	return "", nil
}

func main() {
	lambda.Start(lambdaHandler)
}
