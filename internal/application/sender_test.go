package services_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	services "github.com/mreyeswilson/prueba_stori/internal/application"
	"github.com/mreyeswilson/prueba_stori/mocks"
	"github.com/stretchr/testify/assert"
)


func TestSender(t *testing.T) {
	
	t.Run("Should SendEmail ok", func(t *testing.T) {

		recipientList := []*string{aws.String("abc@abc.com")}

		mockSESAdapter := new(mocks.ISenderAdapter)

		
		mockSESAdapter.On("SendEmail", "test@test.com", recipientList, "Summary", "<html></html>").Return(nil)
		mockSESAdapter.On("GetIdentities").Return(recipientList, nil)
		
		mockSESAdapter.On("GetTemplate", "TransactionSummaryTemplate").Return("<html></html>")
		svc := services.NewSenderService(mockSESAdapter)

		err := svc.SendEmail("test@test.com", "Summary", &map[string]interface{}{
			"test": "test",
		})

		assert.Nil(t, err)
	})

	t.Run("Should GetIdentities failed", func(t *testing.T) {

		recipientList := []*string{aws.String("abc@abc.com")}

		mockSESAdapter := new(mocks.ISenderAdapter)

		
		mockSESAdapter.On("SendEmail", "test@test.com", recipientList, "Summary", "<html></html>").Return(nil)
		mockSESAdapter.On("GetIdentities").Return(nil, errors.New("error"))
		
		mockSESAdapter.On("GetTemplate", "TransactionSummaryTemplate").Return("<html></html>")
		svc := services.NewSenderService(mockSESAdapter)

		err := svc.SendEmail("test@test.com", "Summary", &map[string]interface{}{
			"test": "test",
		})

		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})

}