output "vpc_main" {
  value = aws_vpc.main
}

output "subnet_for_ec2" {
  value = aws_subnet.public_a
}
