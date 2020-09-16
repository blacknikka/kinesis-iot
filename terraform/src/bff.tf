# bff task
data "template_file" "bff_task_definition" {
  template = file("${path.module}/bff-task.json")

  vars = {
    image_url      = "${module.bff_ecr.ecr_repo.repository_url}:latest"
    name           = "bff"
    region         = data.aws_region.current.name
    log_group_name = aws_cloudwatch_log_group.for_ecs.name
  }
}

