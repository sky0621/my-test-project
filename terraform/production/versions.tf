terraform {
  required_version = "~> 1.11.4"
  required_providers {
    google-beta = {
      source  = "hashicorp/google-beta"
      version = ">= 6.30.0"
    }
    google = {
      source  = "hashicorp/google"
      version = ">= 6.30.0"
    }
  }
}
