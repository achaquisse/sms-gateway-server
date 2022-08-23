package internal

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
	"net/http"
	"sms-gateway-server/pkg/database"
	"sms-gateway-server/pkg/helper"
)

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
