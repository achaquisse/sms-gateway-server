package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type mockDBPutItemAPI func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)

func (m mockDBPutItemAPI) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m(ctx, params, optFns...)
}
