terraform {
  required_version = "~> 1.11.4"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 6.30.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = ">= 6.30.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.36.0"
    }
  }
}
