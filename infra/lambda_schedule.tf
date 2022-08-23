resource "aws_lambda_function" "lambda_schedule" {
  filename      = var.zip_lambda.schedule
  function_name = "${local.resource_prefix}-schedule"
  role          = aws_iam_role.lambda_exec_iam_role.arn
  handler       = "schedule"

  runtime     = "go1.x"
  memory_size = 128
  timeout     = 10

  environment {
    variables = {
      DYNAMODB_TABLE_NAME : aws_dynamodb_table.sms_table.name
    }
  }
}

resource "aws_cloudwatch_log_group" "log_schedule" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_schedule.function_name}"
  retention_in_days = 14
}