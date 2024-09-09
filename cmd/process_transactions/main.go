package main

import (
	"context"
	"encoding/csv"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	services "github.com/mreyeswilson/prueba_stori/internal/application"
)

func lambdaHandler(ctx context.Context, event events.S3Event) (string, error) {
	// Crear una sesión de AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Cambia esto a tu región
	})
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	s3Service := s3.New(sess)
	s3Record := event.Records[0]

	record := s3Record.S3
	bucket := record.Bucket.Name
	key := record.Object.Key

	// process the file
	resp, err := s3Service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return "", fmt.Errorf("failed to get object from S3: %v", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	svc := services.NewCalculatorService()
	transactions := svc.ParseInfo(reader)

	summary := svc.MakeSummary(transactions)

	println(summary)

	return "", nil
}

func main() {
	lambda.Start(lambdaHandler)
}
