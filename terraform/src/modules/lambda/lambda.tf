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

  handler          = "src/get_unixtime.lambda_handler"
  filename         = data.archive_file.function_zip.output_path
  runtime          = "python3.6"
  role             = aws_iam_role.lambda_iam_role.arn
  source_code_hash = data.archive_file.function_zip.output_base64sha256
  layers           = [aws_lambda_layer_version.lambda_layer.arn]
}

# lambda event
resource "aws_lambda_event_source_mapping" "kinesis_to_lambda" {
  event_source_arn  = var.kinesis_iot.arn
  function_name     = aws_lambda_function.from_kinesis.arn
  starting_position = "LATEST"
}
