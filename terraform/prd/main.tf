provider "google" {
  project = var.project_id
  region  = var.region
}

module "vpc" {
  source     = "../_modules/vpc"
  project_id = var.project_id
  region     = var.region
}

module "artifact_registry" {
  source        = "../_modules/artifact_registry"
  project_id    = var.project_id
  region        = var.region
  repository_id = var.repository_id
}

module "google_container_cluster" {
  source      = "../_modules/gke"
  project_id  = var.project_id
  region      = var.region
  vpc_name    = module.vpc.vpc_name
  subnet_name = module.vpc.subnet_name
}
