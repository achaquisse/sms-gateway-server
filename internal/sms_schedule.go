package internal

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"sms-gateway-server/pkg/database"
	"sms-gateway-server/pkg/helper"
)

var dbClient = database.DbClient(os.Getenv("AWS_REGION"))
var tableName = os.Getenv("DYNAMODB_TABLE_NAME")

func SmsSchedule(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var smsRequest database.SmsCreateRequest

	errMarshal := json.Unmarshal([]byte(request.Body), &smsRequest)
	if errMarshal != nil {
		errorMessage := fmt.Sprintf("Unable to unmarshall request body. cause: %s", errMarshal.Error())
		log.Warn().Msg(errorMessage)
		return helper.ApiGwError(http.StatusBadRequest, errorMessage)
	}

	errDb := database.SmsCreate(dbClient, tableName, smsRequest)
	if errDb != nil {
		errorMessage := fmt.Sprintf("Unable persist sms. cause: %s", errDb.Error())
		log.Warn().Msg(errorMessage)
		return helper.ApiGwError(http.StatusBadRequest, errorMessage)
	}

	return helper.ApiGwSuccess(http.StatusOK, true)
}
