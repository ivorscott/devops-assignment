provider "aws" {
  region = local.region
  profile = local.profile
}

terraform {
  required_version = ">= 1.13"

  required_providers {
    aws = {
        source  = "hashicorp/aws"
        version = "~> 5.100"
    }
  }
}