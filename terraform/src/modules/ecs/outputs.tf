
# ELB DNS name
output "lb_dns_name" {
  value = aws_alb.app.dns_name
}
