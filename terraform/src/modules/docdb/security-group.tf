resource "aws_security_group" "docdb" {
  name   = "docudb-${var.base_name}"
  vpc_id = var.vpc_main.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "allowd_secutiry" {
  for_each = toset(var.allowed_security_group)
  type                     = "ingress"
  from_port                = 0
  to_port                  = 0
  protocol                 = "-1"
  source_security_group_id = each.value
  security_group_id        = aws_security_group.docdb.id
}
