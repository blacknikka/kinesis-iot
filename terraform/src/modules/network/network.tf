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

# aws_route_table
resource "aws_route_table" "rt_public_a" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${var.base_name}-table-pub-a"
  }
}

# aws_route_table_association
resource "aws_route_table_association" "rta_public_a" {
  route_table_id = aws_route_table.rt_public_a.id
  subnet_id      = aws_subnet.public_a.id
}

# aws_route
resource "aws_route" "route_public_a" {
  route_table_id         = aws_route_table.rt_public_a.id
  gateway_id             = aws_internet_gateway.igw.id
  destination_cidr_block = "0.0.0.0/0"
}

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
