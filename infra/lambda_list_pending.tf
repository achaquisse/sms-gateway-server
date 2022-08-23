resource "aws_lambda_function" "lambda_list_pending" {
  filename      = var.zip_lambda.list_pending
  function_name = "${local.resource_prefix}-list_pending"
  role          = aws_iam_role.lambda_exec_iam_role.arn
  handler       = "list_pending"

  source_code_hash = filebase64sha256(var.zip_lambda.list_pending)

  runtime     = "go1.x"
  memory_size = 128
  timeout     = 10

  environment {
    variables = {
      DYNAMODB_TABLE_NAME : aws_dynamodb_table.sms_table.name
    }
  }
}

resource "aws_cloudwatch_log_group" "log_list_pending" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_list_pending.function_name}"
  retention_in_days = 14
}