package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"os"
	"sms-gateway-server/internal"
)

func main() {
	logLevel := os.Getenv("LOG_LEVEL")

	if logLevel == "INFO" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	if logLevel == "DEBUG" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	lambda.Start(internal.SmsSchedule)
}
