output "get_lifts_url" {
  value = aws_lambda_function_url.get_lifts.function_url
  description = "URL for the GetLifts function" 
}

resource "aws_lambda_function_url" "get_lifts" {
  function_name      = aws_lambda_function.get_lifts.function_name
  authorization_type = "NONE"
}

resource "aws_lambda_function" "get_lifts" {
  function_name = "lifts"
  role = aws_iam_role.get_lifts.arn

  filename = "${path.module}/../application/lifts.zip"
  source_code_hash = filebase64sha256("${path.module}/../application/lifts.zip")
  handler = "main"

  runtime = "provided.al2"
  architectures = ["arm64"]

  environment {
    variables = {
      RequestType = "GetWorkout"
    }
  }
}

resource "aws_lambda_permission" "get_lifts_no_auth" {
  function_name = aws_lambda_function.get_lifts.id
  action = "lambda:InvokeFunctionUrl"
  principal = "*"
  function_url_auth_type = aws_lambda_function_url.get_lifts.authorization_type
}

resource "aws_iam_role_policy_attachment" "get_lifts" {
  role = aws_iam_role.get_lifts.name
  policy_arn = aws_iam_policy.get_lifts.arn
}

resource "aws_iam_role" "get_lifts" {
  name = "get_lifts"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource "aws_iam_policy" "get_lifts" {
  name = "GetLifts"
  policy = data.aws_iam_policy_document.get_lifts.json
}

resource "aws_cloudwatch_log_group" "get_lifts" {
  name = "/aws/lambda/lifts"
}

data "aws_iam_policy_document" "get_lifts" {
  statement {
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
    resources = ["${aws_cloudwatch_log_group.get_lifts.arn}:*"]
  }

  statement {
    actions = [
      "dynamodb:GetItem",
      "dynamodb:Query",
    ]
    resources = [aws_dynamodb_table.lifts.arn]
  }
}
