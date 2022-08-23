package internal

import (
	"os"
	"sms-gateway-server/pkg/database"
)

var dbClient = database.DbClient(os.Getenv("AWS_REGION"))
var tableName = os.Getenv("DYNAMODB_TABLE_NAME")
