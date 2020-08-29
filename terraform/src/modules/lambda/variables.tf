# vpc
variable "vpc_main" {}
variable "subnet_for_lambda1" {}
variable "subnet_for_lambda2" {}


# base name
variable "base_name" {}

# kinesis stream resource
variable "kinesis_iot" {}

# influxDB DNS name
variable "influx_dns_name" {}

# docdb Cluster endpoint
variable "docdb_cluster_endpoint" {}

# docdb admin user
variable "docdb_admin_user" {}

# docdb admin password
variable "docdb_password" {}
