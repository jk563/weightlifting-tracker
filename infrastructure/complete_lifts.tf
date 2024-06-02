resource "aws_lambda_function_url" "complete_lifts" {
  function_name      = aws_lambda_function.complete_lifts.function_name
  authorization_type = "NONE"
}

resource "aws_lambda_function" "complete_lifts" {
  function_name = "lifts_complete"
  role = aws_iam_role.complete_lifts.arn

  filename = "${path.module}/../application/lifts.zip"
  source_code_hash = filebase64sha256("${path.module}/../application/lifts.zip")
  handler = "main"

  runtime = "provided.al2"
  architectures = ["arm64"]

  environment {
    variables = {
      RequestType = "CompleteWorkout"
    }
  }
}

resource "aws_lambda_permission" "complete_lifts_no_auth" {
  function_name = aws_lambda_function.complete_lifts.id
  action = "lambda:InvokeFunctionUrl"
  principal = "*"
  function_url_auth_type = aws_lambda_function_url.complete_lifts.authorization_type
}


resource "aws_iam_role_policy_attachment" "complete_lifts" {
  role = aws_iam_role.complete_lifts.name
  policy_arn = aws_iam_policy.complete_lifts.arn
}

resource "aws_iam_role" "complete_lifts" {
  name = "complete_lifts"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource "aws_iam_policy" "complete_lifts" {
  name = "CompleteLifts"
  policy = data.aws_iam_policy_document.complete_lifts.json
}

data "aws_iam_policy_document" "complete_lifts" {
  statement {
    actions = [
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:Query",
    ]
    resources = [aws_dynamodb_table.lifts.arn]
  }
}

