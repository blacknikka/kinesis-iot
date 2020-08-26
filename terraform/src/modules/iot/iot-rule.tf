resource "aws_iot_topic_rule" "iot_rule" {
  name        = "RuleForwardToKinesis"
  description = "Forward to kinesis from IoT Core"
  enabled     = true
  sql         = "SELECT * FROM 'iot/#'"
  sql_version = "2016-03-23"

  kinesis {
    partition_key = "$${topic()}"
    role_arn      = aws_iam_role.rorle_to_write_kinesis.arn
    stream_name   = var.kinesis_iot.name
  }
}

resource "aws_iam_role" "rorle_to_write_kinesis" {
  name = "role_to_write_kinesis"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "iot.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "iam_policy_for_kinesis" {
  name = "role-to-access-kinesis"
  role = aws_iam_role.rorle_to_write_kinesis.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
        "Effect": "Allow",
        "Action": [
            "kinesis:*"
        ],
        "Resource": "${var.kinesis_iot.arn}"
    }
  ]
}
EOF
}
