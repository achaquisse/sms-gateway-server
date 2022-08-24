package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
	"testing"
)

var mockDBQueryClient = func(t *testing.T) DBQueryAPI {
	return mockDBQueryAPI(func(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
		if params.TableName == nil {
			t.Fatal("table name is required")
		}

		var results []map[string]types.AttributeValue

		item1, _ := attributevalue.MarshalMap(Sms{
			Pk:       SmsStatusPending,
			Sk:       "item1",
			To:       258841234567,
			Message:  "Hello world",
			StatusAt: 0,
		})
		results = append(results, item1)

		item2, _ := attributevalue.MarshalMap(Sms{
			Pk:       SmsStatusPending,
			Sk:       "item2",
			To:       258841234567,
			Message:  "Hello world",
			StatusAt: 0,
		})
		results = append(results, item2)

		return &dynamodb.QueryOutput{
			Items: results,
		}, nil
	})
}

func TestSmsListPending(t *testing.T) {
	out, err := SmsListPending(mockDBQueryClient(t), "sms-test-table")

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(*out) <= 0 {
		t.Fatal("Array should return at least one object")
	} else {
		log.Printf("Found items: %d", len(*out))
	}
}
