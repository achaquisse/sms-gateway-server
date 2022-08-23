# API Gateway

resource "aws_apigatewayv2_api" "lambda" {
  name          = "${local.resource_prefix}-http-gateway"
  protocol_type = "HTTP"
}


# Default stage

resource "aws_apigatewayv2_stage" "lambda" {
  api_id = aws_apigatewayv2_api.lambda.id

  name        = "$default"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gw.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
    }
    )
  }
}


# Integrations

resource "aws_apigatewayv2_integration" "integration_list_pending" {
  api_id = aws_apigatewayv2_api.lambda.id

  integration_uri    = aws_lambda_function.lambda_list_pending.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_integration" "integration_schedule" {
  api_id = aws_apigatewayv2_api.lambda.id

  integration_uri    = aws_lambda_function.lambda_schedule.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_integration" "integration_update_status" {
  api_id = aws_apigatewayv2_api.lambda.id

  integration_uri    = aws_lambda_function.lambda_update_status.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}


# Routes

resource "aws_apigatewayv2_route" "route_list_pending" {
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "GET /pending"
  target    = "integrations/${aws_apigatewayv2_integration.integration_list_pending.id}"
}

resource "aws_apigatewayv2_route" "route_schedule" {
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "POST /pending"
  target    = "integrations/${aws_apigatewayv2_integration.integration_schedule.id}"
}

resource "aws_apigatewayv2_route" "route_update_status" {
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "PUT /pending/{id}"
  target    = "integrations/${aws_apigatewayv2_integration.integration_update_status.id}"
}


# Give API Gateway permission to invoke lambda

resource "aws_lambda_permission" "api_gw_list_pending" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_list_pending.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}

resource "aws_lambda_permission" "api_gw_schedule" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_schedule.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}

resource "aws_lambda_permission" "api_gw_update_status" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_update_status.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}


# API Gateway log group

resource "aws_cloudwatch_log_group" "api_gw" {
  name = "/aws/api_gw/${aws_apigatewayv2_api.lambda.name}"

  retention_in_days = 14
}