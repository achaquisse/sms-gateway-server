package helper

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
	"net/http"
)

func ApiGwSuccess(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	jsonBody, errBody := json.Marshal(body)
	if errBody != nil {
		log.Error().Msgf("Unable to marshal api response. cause: %s", errBody.Error())
		return ApiGwError(http.StatusInternalServerError, errBody.Error())
	}
	apiResponse := events.APIGatewayProxyResponse{
		Body:       string(jsonBody),
		StatusCode: statusCode,
	}
	return apiResponse, nil
}

func ApiGwError(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	apiResponse := events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: statusCode,
	}
	return apiResponse, nil
}
