resource "aws_dynamodb_table" "lifts" {
  name = "lifts"
  
  billing_mode = "PAY_PER_REQUEST"
  table_class = "STANDARD"

  hash_key = "pk"
  range_key = "sk"

  attribute {
    name = "pk"
    type = "S"
  }

  attribute {
    name = "sk"
    type = "S"
  }

}
