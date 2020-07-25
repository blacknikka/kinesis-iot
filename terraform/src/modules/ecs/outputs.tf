
# ELB DNS name
output "influx_dns_name" {
  value = aws_alb.influx.dns_name
}
