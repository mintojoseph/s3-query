# TODO: ideally terraform state should be remote.

provider "aws" {
  version                 = "~> 3.0"
  region                  = var.aws_region
  shared_credentials_file = file(var.aws_cred_file)
  profile                 = var.aws_profile

}

