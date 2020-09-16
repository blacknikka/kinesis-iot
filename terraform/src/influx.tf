# influx task
data "template_file" "influx_task_definition" {
  template = file("${path.module}/influx-task.json")

  vars = {
    image_url      = "influxdb:latest"
    name           = "influx"
    region         = data.aws_region.current.name
    log_group_name = aws_cloudwatch_log_group.for_ecs.name
  }
}

