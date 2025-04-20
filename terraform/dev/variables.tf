variable "project_id" {
  type = string
}

variable "region" {
  type    = string
  default = "asia-northeast1"

  validation {
    condition     = var.region == "asia-northeast1"
    error_message = "The region must be asia-northeast1."
  }
}

variable "repository_id" {
  type    = string
  default = "hello-world-web-api"
}
