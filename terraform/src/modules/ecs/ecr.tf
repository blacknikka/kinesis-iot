resource "aws_ecr_repository" "grafana_repo" {
  name                 = "grafana_repo"
  image_tag_mutability = "MUTABLE"
}
