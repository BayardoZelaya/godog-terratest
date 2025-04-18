variable "cidr" {
  description = "The CIDR block for the VPC"
  type        = string
}

variable "name" {
  description = "The name tag for the VPC"
  type        = string
  default     = "go-bdd-vpc"
}

variable "region" {
  description = "AWS region for resources"
  type        = string
  default     = "us-east-1"
}
