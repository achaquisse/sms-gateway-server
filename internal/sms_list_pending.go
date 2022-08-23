package internal

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
	"net/http"
	"sms-gateway-server/pkg/database"
	"sms-gateway-server/pkg/helper"
)

func SmsListPending(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	out, errDb := database.SmsListPending(dbClient, tableName)

	if errDb != nil {
		errorMessage := fmt.Sprintf("Unable list pending sms. cause: %s", errDb.Error())
		log.Warn().Msg(errorMessage)
		return helper.ApiGwError(http.StatusBadRequest, errorMessage)
	}

	return helper.ApiGwSuccess(http.StatusOK, out)
}
