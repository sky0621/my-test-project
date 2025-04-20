variable "project_id" {
  type = string
}

variable "account_id" {
  type = string
}

variable "roles" {
  type    = set(string)
  default = []
}
