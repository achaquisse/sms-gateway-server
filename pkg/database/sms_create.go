package database

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/rs/zerolog/log"
	"time"
)

type SmsCreateRequest struct {
	Sk      string
	To      int
	Message string
}

func SmsCreate(client DBPutItemAPI, tableName string, request SmsCreateRequest) error {
	log.Debug().Msgf("SmsCreate request: %v", request)

	if request.Sk == "" {
		return errors.New("'sk' is required")
	}

	if request.To <= 0 {
		return errors.New("'to' must be a valid phone number")
	}

	unmarshalledSms := Sms{
		Pk:       SmsStatusPending,
		Sk:       request.Sk,
		To:       request.To,
		Message:  request.Message,
		StatusAt: time.Now().Unix(),
	}

	item, errMarshal := attributevalue.MarshalMap(unmarshalledSms)
	if errMarshal != nil {
		log.Error().Msgf("Failed to marshal. request: %v, cause: %s", request, errMarshal.Error())
		return errMarshal
	}

	out, errPut := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		Item:      item,
		TableName: &tableName,
	})

	if errPut != nil {
		log.Error().Msgf("Failed to put item. request: %v, cause: %s", request, errPut.Error())
		return errMarshal
	}

	log.Debug().Msgf("SmsCreate success: request: %v. out: %v", request, out)
	return nil
}
