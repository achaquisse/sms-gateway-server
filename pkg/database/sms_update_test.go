package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"testing"
	"time"
)

type mockDBGetAndWriteAPI struct{}

func (m mockDBGetAndWriteAPI) GetItem(_ context.Context, _ *dynamodb.GetItemInput, _ ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	sms := Sms{
		Pk:       SmsStatusPending,
		Sk:       "payment#1",
		To:       258841234567,
		Message:  "Pending payment",
		StatusAt: time.Now().Unix(),
	}

	item, _ := attributevalue.MarshalMap(sms)

	return &dynamodb.GetItemOutput{
		Item: item,
	}, nil
}

func (m mockDBGetAndWriteAPI) TransactWriteItems(_ context.Context, _ *dynamodb.TransactWriteItemsInput, _ ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error) {
	return &dynamodb.TransactWriteItemsOutput{}, nil
}

func TestUpdateSuccess(t *testing.T) {
	err := SmsUpdate(mockDBGetAndWriteAPI{}, "sms-test-table", SmsUpdateRequest{
		Sk:     "payment#1",
		Status: SmsStatusSuccess,
	})

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdateError(t *testing.T) {
	err := SmsUpdate(mockDBGetAndWriteAPI{}, "sms-test-table", SmsUpdateRequest{
		Sk:     "payment#1",
		Status: "unknown",
	})

	if err == nil {
		t.Fatal("should return an error")
	}

	expectedErr := "'status' should be Success or Pending"

	if err.Error() != expectedErr {
		t.Fatalf("error message different from expected. Got %s, expected: %s", err.Error(), expectedErr)
	}
}
