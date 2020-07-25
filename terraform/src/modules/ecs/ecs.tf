resource "aws_ecs_cluster" "main" {
  name = var.base_name
}

# grafana task
data "template_file" "grafana_task_definition" {
  template = file("${path.module}/grafana-task.json")

  vars = {
    image_url      = "grafana/grafana:latest"
    name           = "grafana"
    region         = data.aws_region.current.name
    log_group_name = aws_cloudwatch_log_group.app.name
  }
}

# influx task
data "template_file" "influx_task_definition" {
  template = file("${path.module}/influx-task.json")

  vars = {
    image_url      = "influxdb:latest"
    name           = "influx"
    region         = data.aws_region.current.name
    log_group_name = aws_cloudwatch_log_group.app.name
  }
}

# grafana task definition
resource "aws_ecs_task_definition" "grafana_task" {
  family                = "grafana_task"
  network_mode          = "awsvpc"
  cpu                   = 256
  memory                = 512
  container_definitions = data.template_file.grafana_task_definition.rendered
  execution_role_arn    = aws_iam_role.ecs_execution_role.arn

  requires_compatibilities = [
    "FARGATE",
  ]
}

# influxdb task definition
resource "aws_ecs_task_definition" "influx_task" {
  family                = "influx_task"
  network_mode          = "awsvpc"
  cpu                   = 256
  memory                = 512
  container_definitions = data.template_file.influx_task_definition.rendered
  execution_role_arn    = aws_iam_role.ecs_execution_role.arn

  requires_compatibilities = [
    "FARGATE",
  ]
}

resource "aws_ecs_service" "grafana" {
  name            = "grafana"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.grafana_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  depends_on = [
    aws_alb_listener.influx_lb,
  ]

  load_balancer {
    target_group_arn = aws_alb_target_group.alb_target.id
    container_name   = "grafana"
    container_port   = 3000
  }

  network_configuration {
    subnets = [
      var.subnet_for_app.id,
      var.subnet_for_app2.id,
    ]

    security_groups = [
      aws_security_group.load_balancers.id,
    ]

    assign_public_ip = true
  }
}

resource "aws_ecs_service" "influx" {
  name            = "influx"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.influx_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  depends_on = [
    aws_alb_listener.influx_listner_lb,
  ]

  load_balancer {
    target_group_arn = aws_alb_target_group.alb_target_influx.id
    container_name   = "influx"
    container_port   = 8086
  }

  network_configuration {
    subnets = [
      var.subnet_for_app.id,
      var.subnet_for_app2.id,
    ]

    security_groups = [
      aws_security_group.load_balancers_influx.id,
    ]

    assign_public_ip = true
  }
}
