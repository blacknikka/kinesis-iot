variable "base_name" {
  description = "base name for this project"
  type        = string
  default     = "iot"
}

variable "docdb_admin_user" {
  description = "admin user name for docdb"
  type        = string
  default     = "tf_iot_admin"
}

variable "docdb_password" {
  description = "password for docdb"
  type        = string
  default     = "secretsecret"
}

