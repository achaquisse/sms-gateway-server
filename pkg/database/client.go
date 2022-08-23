package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DbClient(region string) *dynamodb.Client {
	var cfg, _ = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	return dynamodb.NewFromConfig(cfg)
}

type DBPutItemAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}
