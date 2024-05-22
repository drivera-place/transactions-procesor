provider "aws" {
  region  = "us-east-1"
}

resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "Transactions"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Id"
  range_key      = "Date"

  attribute {
    name = "Id"
    type = "N"
  }

  attribute {
    name = "Date"
    type = "S"
  }

  ttl {
    attribute_name = "TimeToExist"
    enabled        = false
  }

  tags = {
    Name        = "dynamodb-table"
    Environment = "Training"
  }
}