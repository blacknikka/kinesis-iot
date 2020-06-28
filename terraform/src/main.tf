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

module "ec2" {
  source = "./modules/ec2"

  base_name      = var.base_name
  vpc_main       = module.network.vpc_main
  subnet_for_ec2 = module.network.subnet_for_ec2
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
