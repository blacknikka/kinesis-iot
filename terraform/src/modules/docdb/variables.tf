# name
variable "base_name" {}

# vpc information.
variable "vpc_main" {}

# subnet for docdb
variable "subnet_for_docdb1" {}
variable "subnet_for_docdb2" {}

# instance class
variable "docdb_instance_class" {
  default = "db.t3.medium"
}

# password
variable "docdb_password" {}

