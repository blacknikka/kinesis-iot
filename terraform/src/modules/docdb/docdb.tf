resource "aws_docdb_subnet_group" "service" {
  name       = "docdb-${var.base_name}-subnet"
  subnet_ids = [var.subnet_for_docdb1.id, var.subnet_for_docdb2.id]
}

resource "aws_docdb_cluster_instance" "service" {
  count              = 1
  identifier         = format("docdb-%s-instnace-%02d", var.base_name, count.index + 1)
  cluster_identifier = aws_docdb_cluster.service.id
  instance_class     = var.docdb_instance_class
}

resource "aws_docdb_cluster" "service" {
  skip_final_snapshot             = true
  db_subnet_group_name            = aws_docdb_subnet_group.service.name
  cluster_identifier              = "docdb-cluster-${var.base_name}"
  engine                          = "docdb"
  master_username                 = "tf_${replace(var.base_name, "-", "_")}_admin"
  master_password                 = var.docdb_password
  db_cluster_parameter_group_name = aws_docdb_cluster_parameter_group.service.name
  vpc_security_group_ids          = [aws_security_group.docdb.id]
}

resource "aws_docdb_cluster_parameter_group" "service" {
  family = "docdb3.6"
  name   = "docdb-parameter-group-${var.base_name}"
}
