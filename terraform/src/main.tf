terraform {
  required_version = ">= 0.12.0"
  #   backend "s3" {
  #     region  = "ap-northeast-1"
  #     encrypt = true

  #     bucket = "terraform-bucket-fortfstate"
  #     key    = "terraform.tfstate"
  #   }
}

provider "aws" {
  region = "ap-northeast-1"
  #   assume_role {
  #     role_arn = var.assume_role
  #   }
}

module "network" {
  source = "./modules/network"

  base_name = var.base_name
}

module "iot" {
  source = "./modules/iot"

  base_name   = var.base_name
  kinesis_iot = module.kinesis.kinesis_iot
}

module "kinesis" {
  source = "./modules/kinesis"

  base_name = var.base_name
}

module "lambda" {
  source = "./modules/lambda"

  base_name              = var.base_name
  kinesis_iot            = module.kinesis.kinesis_iot
  influx_dns_name        = module.ecs_influx.lb_dns_name
  docdb_cluster_endpoint = module.docdb.docdb_cluster.endpoint
  docdb_admin_user       = var.docdb_admin_user
  docdb_password         = var.docdb_password

  vpc_main           = module.network.vpc_main
  subnet_for_lambda1 = module.network.subnet_for_app
  subnet_for_lambda2 = module.network.subnet_for_app2
}

# bff ecr
module "bff_ecr" {
  source = "./modules/ecr"

  app_name = "bff"
}

# influx CloudWatch Logs
resource "aws_cloudwatch_log_group" "for_ecs" {
  name = "tf-ecs-group"
}

# execution role (ecs)
resource "aws_iam_role" "ecs_execution_role" {
  name = "ecs_execution_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "ecs-tasks.amazonaws.com"
        ]
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ecs_task" {
  role       = aws_iam_role.ecs_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_cluster" "main" {
  name = var.base_name
}

module "ecs_influx" {
  source = "./modules/ecs"

  app_name             = "influx"
  base_name            = var.base_name
  vpc_main             = module.network.vpc_main
  subnet_for_app       = module.network.subnet_for_app
  subnet_for_app2      = module.network.subnet_for_app2
  ecs_cluster          = aws_ecs_cluster.main
  ecs_task_definition  = data.template_file.influx_task_definition.rendered
  ecs_container_name   = "influx"
  ecs_container_port   = 8086
  lb_health_check_path = "/health"
  ecs_execution_role   = aws_iam_role.ecs_execution_role

}

module "ecs_bff" {
  source = "./modules/ecs"

  app_name             = "bff"
  base_name            = var.base_name
  vpc_main             = module.network.vpc_main
  subnet_for_app       = module.network.subnet_for_app
  subnet_for_app2      = module.network.subnet_for_app2
  ecs_cluster          = aws_ecs_cluster.main
  ecs_task_definition  = data.template_file.bff_task_definition.rendered
  ecs_container_name   = "bff"
  ecs_container_port   = 8080
  lb_health_check_path = "/current"
  ecs_execution_role   = aws_iam_role.ecs_execution_role

}

module "docdb" {
  source = "./modules/docdb"

  base_name              = var.base_name
  vpc_main               = module.network.vpc_main
  subnet_for_docdb1      = module.network.subnet_for_app
  subnet_for_docdb2      = module.network.subnet_for_app2
  docdb_admin_user       = var.docdb_admin_user
  docdb_password         = var.docdb_password
  allowed_security_group = [
      module.lambda.secutiry_group_for_lambda.id,
      module.ecs_bff.security_group.id,
  ]
}

module "s3" {
  source = "./modules/s3"

  app_name = var.base_name
}
