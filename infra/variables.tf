variable "environment" {
  default = "dev"
}

variable "zip_lambda" {
  type = map(string)
  default = {
    list_pending = "../.out/list_pending.zip"
    schedule = "../.out/schedule.zip"
    update_status = "../.out/update_status.zip"
  }
}