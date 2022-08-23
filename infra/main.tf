terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.0.0"
    }
  }

  backend "s3" {
    bucket         = "exodus-tf-state"
    encrypt        = true
    region         = "af-south-1"
    dynamodb_table = "exodus-tf-state-locking"
  }

  required_version = ">= 1.1.0"
}

provider "aws" {
  region = local.aws_region
}