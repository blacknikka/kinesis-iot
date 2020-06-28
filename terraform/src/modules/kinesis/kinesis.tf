# -----------------
# Kinesis stream
# -----------------
resource "aws_kinesis_stream" "iot_stream" {
  name             = var.base_name
  shard_count      = 1
  retention_period = 24
}
