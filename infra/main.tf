terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.0.0"
    }
  }

  required_version = ">= 1.1.0"

  cloud {
    organization = "exodus-mz"

    workspaces {
      name = "sms-gateway-server"
    }
  }
}

provider "aws" {
  region = local.aws_region
}