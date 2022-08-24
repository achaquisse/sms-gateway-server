package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
	"time"
)

type SmsUpdateRequest struct {
	Sk     string
	Status string
}

func SmsUpdate(client DBGetAndWriteAPI, tableName string, request SmsUpdateRequest) error {
	log.Debug().Msgf("SmsUpdate request: %v", request)

	sms, errSms := readSms(client, tableName, request)
	if errSms != nil {
		return errSms
	}

	if request.Status == SmsStatusSuccess {
		sms.Pk = SmsStatusSuccess
	}
	if request.Status == SmsStatusPending {
		sms.Pk = SmsStatusPending
	}
	sms.StatusAt = time.Now().Unix()

	smsNewMarshaled, _ := attributevalue.MarshalMap(sms)

	out, errTransaction := client.TransactWriteItems(context.Background(), &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					Item:      smsNewMarshaled,
					TableName: &tableName,
				},
			},
			{
				Delete: &types.Delete{
					Key: map[string]types.AttributeValue{
						"pk": &types.AttributeValueMemberS{Value: SmsStatusPending},
						"sk": &types.AttributeValueMemberS{Value: request.Sk},
					},
					TableName: &tableName,
				},
			},
		},
	})

	if errTransaction != nil {
		log.Error().Msgf("Failed to update item. request: %v, cause: %s", request, errTransaction.Error())
		return errTransaction
	}

	log.Debug().Msgf("SmsUpdate success: request: %v. out: %v", request, out)
	return nil
}

func readSms(client DBGetAndWriteAPI, tableName string, request SmsUpdateRequest) (Sms, error) {
	if request.Sk == "" {
		return Sms{}, errors.New("'sk' is required")
	}

	if request.Status != SmsStatusSuccess && request.Status != SmsStatusFailed {
		return Sms{}, errors.New(fmt.Sprintf("'status' should be %s or %s", SmsStatusSuccess, SmsStatusPending))
	}

	item, errGet := client.GetItem(context.Background(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: SmsStatusPending},
			"sk": &types.AttributeValueMemberS{Value: request.Sk},
		},
		TableName: &tableName,
	})

	if errGet != nil {
		log.Error().Msgf("Failed to fetch %s=%s. cause: %s", SmsStatusPending, request.Sk, errGet.Error())
		return Sms{}, errGet
	}

	if len(item.Item) == 0 {
		msg := fmt.Sprintf("Not found %s=%s", SmsStatusPending, request.Sk)
		log.Warn().Msg(msg)
		return Sms{}, errors.New(msg)
	}

	var sms Sms
	errUnmarshal := attributevalue.UnmarshalMap(item.Item, &sms)

	if errUnmarshal != nil {
		log.Error().Msgf("Failed to unmarshal %s. cause: %s", request.Sk, errGet.Error())
		return Sms{}, errUnmarshal
	}

	return sms, nil
}
