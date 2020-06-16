# aws_instance
resource "aws_instance" "app" {
  ami                    = data.aws_ssm_parameter.amzn2_ami.value
  instance_type          = "t2.micro"
  subnet_id              = var.subnet_for_ec2.id
  vpc_security_group_ids = [aws_security_group.for_app_server.id]

  tags = {
    Name = "${var.base_name}-ec2"
  }
}

resource "aws_security_group" "for_app_server" {
  name        = "for_app_server"
  description = "for application server."
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

  tags = {
    Name = "allow only 80"
  }
}
