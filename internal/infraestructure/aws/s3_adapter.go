package adapters

import (
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mreyeswilson/prueba_stori/internal/domain/interfaces"
)

type S3Adapter struct {
	config *Config
}

func NewS3Adapter() interfaces.IStorageAdapter {
	return &S3Adapter{
		config: NewConfig(),
	}
}

func (s *S3Adapter) GetObject(bucketName string, key string) (io.ReadCloser, error) {
	sess, err := s.config.GetSession()
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, errors.New("failed to get object")
	}

	return obj.Body, nil
}
