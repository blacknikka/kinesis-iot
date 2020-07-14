resource "aws_alb" "app" {
  name            = "elb-for-app"
  security_groups = [aws_security_group.load_balancers.id]
  subnets         = [
      var.subnet_for_app.id,
      var.subnet_for_app2.id,
  ]
}

resource "aws_alb_target_group" "alb_target" {
  name     = "alb-target-for-app"
  port     = 3000
  protocol = "HTTP"
  target_type = "ip"
  vpc_id   = var.vpc_main.id
}

resource "aws_alb_listener" "front_end" {
  load_balancer_arn = aws_alb.app.id
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.alb_target.id
    type             = "forward"
  }
}

resource "aws_security_group" "load_balancers" {
  name        = "load_balancers"
  description = "Allows all traffic"
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
}

