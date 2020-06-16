# aws_vpc
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
    Name = var.network_name
  }
}

# aws_internet_gateway
resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = var.base_name
  }
}

# # aws_eip
# resource "aws_eip" "nat_a" {
#   vpc        = true
#   depends_on = [aws_internet_gateway.igw]
#   tags = {
#     Name = "${var.network_name}-a"
#   }
# }

# # aws_nat_gateway
# resource "aws_nat_gateway" "gw_nat_a" {
#   allocation_id = aws_eip.nat_a.id
#   subnet_id     = aws_subnet.public_a.id
#   depends_on    = [aws_internet_gateway.igw]
#   tags = {
#     Name = "${var.network_name}-nat-a"
#   }
# }

# aws_route_table
resource "aws_route_table" "rt_public_a" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${var.base_name}-table-pub-a"
  }
}

# resource "aws_route_table" "rt_private_a" {
#   vpc_id = aws_vpc.main.id
#   tags = {
#     Name = "${var.base_name}-table-private-a"
#   }
# }

# aws_route_table_association
resource "aws_route_table_association" "rta_public_a" {
  route_table_id = aws_route_table.rt_public_a.id
  subnet_id      = aws_subnet.public_a.id
}

# resource "aws_route_table_association" "rta_private_a" {
#   route_table_id = aws_route_table.rt_private_a.id
#   subnet_id      = aws_subnet.private_a.id
# }

# aws_route
resource "aws_route" "route_public_a" {
  route_table_id         = aws_route_table.rt_public_a.id
  gateway_id             = aws_internet_gateway.igw.id
  destination_cidr_block = "0.0.0.0/0"
}

# resource "aws_route" "route_private_a" {
#   route_table_id         = aws_route_table.rt_private_a.id
#   nat_gateway_id         = aws_nat_gateway.gw_nat_a.id
#   destination_cidr_block = "0.0.0.0/0"
# }

# subnet
resource "aws_subnet" "public_a" {
  availability_zone       = "${data.aws_region.current.name}a"
  cidr_block              = "10.0.1.0/24"
  vpc_id                  = aws_vpc.main.id
  map_public_ip_on_launch = true
  tags = {
    Name = "${var.base_name}-subnet-pub-a"
  }
}

# resource "aws_subnet" "private_a" {
#   availability_zone       = "${data.aws_region.current.name}a"
#   cidr_block              = "10.0.128.0/24"
#   vpc_id                  = aws_vpc.main.id
#   map_public_ip_on_launch = true
#   tags = {
#     Name = "${var.base_name}-subnet-pri-a"
#   }
# }
