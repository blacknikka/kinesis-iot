
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

resource "aws_iot_certificate" "iot_certificate" {
  active = true
}

resource "local_file" "cert_pem" {
  sensitive_content = aws_iot_certificate.iot_certificate.certificate_pem
  filename = "${path.module}/cert/iot-cert.pem"
}

resource "local_file" "cert_public_key" {
  sensitive_content = aws_iot_certificate.iot_certificate.public_key
  filename = "${path.module}/cert/iot-motor.public.key"
}

resource "local_file" "cert_private_key" {
  sensitive_content = aws_iot_certificate.iot_certificate.private_key
  filename = "${path.module}/cert/iot-motor.private.key"
}
