package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
)

func SmsListPending(client DBQueryAPI, tableName string) (*[]Sms, error) {
	log.Debug().Msg("Received request to list pending sms")

	resp, errQuery := client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("pk = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: SmsStatusPending},
		},
	})

	if errQuery != nil {
		log.Error().Msgf("Failed to query pending sms. cause: %s", errQuery.Error())
		return nil, errQuery
	}

	var smsList []Sms

	errUnmarshal := attributevalue.UnmarshalListOfMaps(resp.Items, &smsList)
	if errUnmarshal != nil {
		log.Error().Msgf("Failed to unmarshal pending sms. cause: %s", errUnmarshal.Error())
		return nil, errUnmarshal
	}

	log.Debug().Msgf("Found sms: %d", len(smsList))

	return &smsList, nil
}
