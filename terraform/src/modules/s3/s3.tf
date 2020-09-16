resource "aws_s3_bucket" "bucket" {
  bucket = "${var.app_name}-app-s3-endpoint"
  acl    = "public-read"

  website {
    index_document = "index.html"
  }
}

resource "aws_s3_bucket_policy" "source" {
    bucket = aws_s3_bucket.bucket.id
    policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AddPerm",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "${aws_s3_bucket.bucket.arn}/*",
            "Condition": {}
        }
    ]
}
EOF
}
