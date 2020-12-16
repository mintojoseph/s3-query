resource "aws_s3_bucket" "mintos-test-bucket" {
  bucket = "mintos-test-bucket"
  acl    = "private"

  tags = {
    Name        = "Mintos"
    Environment = "test"
  }
}
