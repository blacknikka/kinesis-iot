
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

# resource "aws_iot_certificate" "cert" {
#   active = true
# }
