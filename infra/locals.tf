locals {
  aws_region      = "us-east-1"
  aws_profile     = "default"
  project_name    = "sms-gateway-server"
  resource_prefix = "${local.project_name}-${var.environment}"
}