# ----------------------
# grafana
# ----------------------
resource "aws_alb" "app" {
  name            = "elb-for-app"
  security_groups = [aws_security_group.load_balancers.id]
  subnets = [
    var.subnet_for_app.id,
    var.subnet_for_app2.id,
  ]
}

resource "aws_alb_target_group" "alb_target" {
  name        = "alb-target-for-app"
  port        = 3000
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.vpc_main.id

  health_check {
    matcher = "200,302"
    port    = 3000
  }
}

resource "aws_alb_listener" "influx_lb" {
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

# ----------------------
# influx
# ----------------------
resource "aws_alb" "influx" {
  name            = "elb-for-influx"
  security_groups = [aws_security_group.load_balancers_influx.id]
  subnets = [
    var.subnet_for_app.id,
    var.subnet_for_app2.id,
  ]
}

resource "aws_alb_target_group" "alb_target_influx" {
  name        = "alb-target-for-influx"
  port        = 8086
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.vpc_main.id

  health_check {
    matcher = 200
    port    = 8086
    path    = "/health"
  }
}

resource "aws_alb_listener" "influx_listner_lb" {
  load_balancer_arn = aws_alb.influx.id
  port              = 8086
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.alb_target_influx.id
    type             = "forward"
  }
}

resource "aws_security_group" "load_balancers_influx" {
  name        = "load_balancers_for_influx"
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

