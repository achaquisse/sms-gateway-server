package internal

import "github.com/aws/aws-lambda-go/events"

func SmsSchedule(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil
}
