# -----------------
# IoT Core
# -----------------
resource "aws_iot_thing_type" "type_motor" {
  name = "type_motor"

  properties {
    description  = "type for motor sensor"
  }
}

resource "aws_iot_thing" "iot_sens" {
  name = "${var.base_name}-thing"

  thing_type_name = aws_iot_thing_type.type_motor.name
}

# -----------------
# IoT Certificate
# -----------------
resource "aws_iot_certificate" "iot_certificate" {
  active = true
}

# -----------------
# Create Certificate files by local files
#
# These certificate files are made by shell script `MakeKey.sh`.
# -----------------
resource "local_file" "cert_pem" {
  sensitive_content = aws_iot_certificate.iot_certificate.certificate_pem
  filename = "${path.module}/cert/iot-cert.cert.pem"
}

resource "local_file" "cert_public_key" {
  sensitive_content = aws_iot_certificate.iot_certificate.public_key
  filename = "${path.module}/cert/iot-motor.public.key"
}

resource "local_file" "cert_private_key" {
  sensitive_content = aws_iot_certificate.iot_certificate.private_key
  filename = "${path.module}/cert/iot-motor.private.key"
}

# -----------------
# Policy Attachment(IoT)
# -----------------
resource "aws_iot_policy_attachment" "iot_attachment" {
  policy = aws_iot_policy.allow_all_iot.name
  target = aws_iot_certificate.iot_certificate.arn
}