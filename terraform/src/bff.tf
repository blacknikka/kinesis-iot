# bff task
data "template_file" "bff_task_definition" {
  template = file("${path.module}/bff-task.json")

  vars = {
    image_url          = "${module.bff_ecr.ecr_repo.repository_url}:latest"
    name               = "bff"
    region             = data.aws_region.current.name
    log_group_name     = aws_cloudwatch_log_group.for_ecs.name
    CLUSTER_ENDPOINT   = "${module.docdb.docdb_cluster.endpoint}:27017"
    CLUSTER_USERNAME   = var.docdb_admin_user
    CLUSTER_PASSWORD   = var.docdb_password
    CLUSTER_OPTIONS    = "/?ssl=true&sslcertificateauthorityfile=rds-combined-ca-bundle.pem&replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false"
    DATABSE_NAME       = "db"
    NORMAL_COLLECTION  = "col"
    SUMMARY_COLLECTION = "summary"
  }
}

