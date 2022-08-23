package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"sms-gateway-server/internal"
)

func main() {
	lambda.Start(internal.SmsListPending)
}
