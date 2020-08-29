# Archive
data "archive_file" "layer_zip" {
  type        = "zip"
  source_dir  = "${path.module}/build/layer"
  output_path = "${path.module}/lambda/layer.zip"
}
data "archive_file" "function_zip" {
  type        = "zip"
  source_dir  = "${path.module}/build/function"
  output_path = "${path.module}/lambda/function.zip"
}

# Layer
resource "aws_lambda_layer_version" "lambda_layer" {
  layer_name       = "${var.base_name}_lambda_layer"
  filename         = data.archive_file.layer_zip.output_path
  source_code_hash = data.archive_file.layer_zip.output_base64sha256
}

# Function
resource "aws_lambda_function" "from_kinesis" {
  function_name = "${var.base_name}_from_kinesis"

  handler          = "src/postData.lambda_handler"
  filename         = data.archive_file.function_zip.output_path
  runtime          = "python3.6"
  role             = aws_iam_role.lambda_iam_role.arn
  source_code_hash = data.archive_file.function_zip.output_base64sha256
  layers           = [aws_lambda_layer_version.lambda_layer.arn]

  vpc_config {
    subnet_ids = [
      var.subnet_for_lambda1.id,
      var.subnet_for_lambda2.id
    ]
    security_group_ids = [aws_security_group.for_lambda.id]
  }

  environment {
    variables = {
      INFLUX_ENDPOINT        = var.influx_dns_name
      DOCDB_CLUSTER_ENDPOINT = var.docdb_cluster_endpoint
      DOCDB_ADMIN_USER       = var.docdb_admin_user
      DOCDB_ADMIN_PASSWORD   = var.docdb_password
    }
  }
}

# lambda event
resource "aws_lambda_event_source_mapping" "kinesis_to_lambda" {
  event_source_arn  = var.kinesis_iot.arn
  function_name     = aws_lambda_function.from_kinesis.arn
  starting_position = "LATEST"
}

# security group for lambda
resource "aws_security_group" "for_lambda" {
  name        = "for_lambda"
  description = "for lambda"
  vpc_id      = var.vpc_main.id

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


