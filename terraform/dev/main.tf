provider "google" {
  project = var.project_id
  region  = var.region
}

module "service" {
  source     = "../_modules/service"
  project_id = var.project_id
}

module "random" {
  source = "../_modules/random"
}

module "iam_workload_identity_pool_for_github_actions" {
  source                             = "../_modules/iam_workload_identity_pool_provider"
  project_id                         = var.project_id
  workload_identity_pool_id          = "gha-pool-${module.random.util_random_id}"
  workload_identity_pool_provider_id = "gha-prov-${module.random.util_random_id}"
}

module "service_account_for_github_actions" {
  source     = "../_modules/service_account"
  project_id = var.project_id
  account_id = "github-actions-sa"
  roles = [
    "iam.serviceAccountUser",
    "artifactregistry.writer",
    "cloudbuild.builds.builder",
    "clouddeploy.operator",
    "storage.admin",
  ]
}

module "service_account_iam_binding_for_github_actions" {
  source             = "../_modules/service_account_iam_binding"
  service_account_id = module.service_account_for_github_actions.name
  members            = ["principalSet://iam.googleapis.com/${module.iam_workload_identity_pool_for_github_actions.name}/attribute.repository/sky0621/my-test-project"]
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
  depends_on    = ["module.service"]
}

# module "google_container_cluster" {
#   source      = "../_modules/gke"
#   project_id  = var.project_id
#   region      = var.region
#   vpc_name    = module.vpc.vpc_name
#   subnet_name = module.vpc.subnet_name
# }
