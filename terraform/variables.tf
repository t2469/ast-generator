variable "project_name" {
  description = "project_name"
  type        = string
  default     = "ast-generator"
}

variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "ap-northeast-1"
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidr_a" {
  description = "CIDR block for the public subnet in AZ a"
  type        = string
  default     = "10.0.1.0/24"
}

variable "public_subnet_cidr_c" {
  description = "CIDR block for the public subnet in AZ b"
  type        = string
  default     = "10.0.3.0/24"
}

variable "private_subnet_cidr_a" {
  description = "CIDR block for the private subnet in AZ a"
  type        = string
  default     = "10.0.2.0/24"
}

variable "private_subnet_cidr_c" {
  description = "CIDR block for the private subnet in AZ c"
  type        = string
  default     = "10.0.4.0/24"
}

variable "key_pair_name" {
  description = "EC2 key pair name for NAT instance"
  type        = string
  sensitive   = true
}

variable "db_username" {
  description = "Username for RDS database"
  type        = string
  sensitive   = true
}

variable "db_password" {
  description = "Password for RDS database"
  type        = string
  sensitive   = true
}

variable "db_name" {
  description = "Database name for RDS"
  type        = string
  sensitive   = true
}

variable "admin_ip" {
  description = "Administrator's public IP address"
  type        = string
  sensitive   = true
}

variable "google_client_id" {
  description = "Google OAuth Client ID"
  type        = string
  sensitive   = true
}

variable "google_client_secret" {
  description = "Google OAuth Client Secret"
  type        = string
  sensitive   = true
}

variable "google_redirect_url" {
  description = "Google OAuth Redirect URL"
  type        = string
  sensitive   = true
}

variable "jwt_secret" {
  description = "JWT secret key"
  type        = string
  sensitive   = true
}

variable "frontend_url" {
  type        = string
  description = "Frontend URL used by the backend service"
}

variable "vite_api_url" {
  type        = string
  description = "Frontend URL used by the backend service"
}

variable "root_domain" {
  description = "The root domain in Route53"
  type        = string
  default     = "t2469.com"
}

variable "frontend_subdomain" {
  description = "Subdomain for the frontend"
  type        = string
  default     = "ast-generator"
}

variable "api_subdomain" {
  description = "Subdomain for API ALB"
  type        = string
  default     = "api.ast-generator"
}

variable "image_tag" {
  description = "Docker image tag for versioning the api image"
  type        = string
}