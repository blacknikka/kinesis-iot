# -----------------
# Policy
# -----------------
resource "aws_iot_policy" "allow_all_iot" {
  name = "AllowAllOfIoT"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "iot:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}
