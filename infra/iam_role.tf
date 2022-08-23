resource "aws_iam_role" "lambda_exec_iam_role" {
  name = "${local.resource_prefix}-lambda-iam-role"

  assume_role_policy = jsonencode({
    "Version" : "2012-10-17"
    "Statement" : [
      {
        "Action" : "sts:AssumeRole",
        "Principal" : {
          "Service" : ["lambda.amazonaws.com"]
        },
        "Effect" : "Allow",
        "Sid" : ""
      }
    ]
  })
}

resource "aws_iam_policy" "iam_policy" {
  name = "${local.resource_prefix}-iam-policy"

  policy = jsonencode({
    "Version" : "2012-10-17"
    "Statement" : [
      {
        "Action" : [
          "dynamodb:*",
        ],
        "Effect" : "Allow",
        "Resource" : aws_dynamodb_table.sms_table.arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_basic_exec_policy" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda_exec_iam_role.name
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb_exec_policy" {
  policy_arn = aws_iam_policy.iam_policy.arn
  role       = aws_iam_role.lambda_exec_iam_role.name
}