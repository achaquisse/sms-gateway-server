resource "aws_lambda_function" "lambda_update_status" {
  filename      = var.zip_lambda.update_status
  function_name = "${local.resource_prefix}-update-status"
  role          = aws_iam_role.lambda_exec_iam_role.arn
  handler       = "update-status"

  runtime       = "go1.x"
  memory_size   = 128
  timeout       = 10
  architectures = ["arm64"]

  environment {
    variables = {
      DYNAMODB_TABLE_NAME : aws_dynamodb_table.sms_table.name
    }
  }
}

resource "aws_cloudwatch_log_group" "log_update_status" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_update_status.function_name}"
  retention_in_days = 14
}