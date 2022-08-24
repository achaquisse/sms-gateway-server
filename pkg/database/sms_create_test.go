package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"testing"
)

type mockDBQueryAPI func(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)

func (m mockDBQueryAPI) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return m(ctx, params, optFns...)
}

type mockDBPutItemAPI func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)

func (m mockDBPutItemAPI) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m(ctx, params, optFns...)
}

var mockDBPutItemClient = func(t *testing.T) DBPutItemAPI {
	return mockDBPutItemAPI(func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
		if params.TableName == nil {
			t.Fatal("table name is required")
		}
		return &dynamodb.PutItemOutput{}, nil
	})
}

func TestSmsCreateSuccess(t *testing.T) {
	out := SmsCreate(mockDBPutItemClient(t), "sms-test-table", SmsCreateRequest{
		Sk:      "STU-0001:payment#1",
		To:      258841234567,
		Message: "Payment received",
	})

	if out != nil {
		t.Errorf("Failed to create sms request. %s", out.Error())
	}
}

func TestSmsCreateFailure_Sk(t *testing.T) {
	out := SmsCreate(mockDBPutItemClient(t), "sms-test-table", SmsCreateRequest{
		Sk:      "",
		To:      258841234567,
		Message: "Payment received",
	})

	expected := "'sk' is required"

	if out.Error() != expected {
		t.Errorf("Failed. Expected: %s, Got: %s", expected, out.Error())
	}
}

func TestSmsCreateFailure_To(t *testing.T) {
	out := SmsCreate(mockDBPutItemClient(t), "sms-test-table", SmsCreateRequest{
		Sk:      "STU-0001:payment#1",
		To:      0,
		Message: "Payment received",
	})

	expected := "'to' must be a valid phone number"

	if out.Error() != expected {
		t.Errorf("Failed. Expected: %s, Got: %s", expected, out.Error())
	}
}
