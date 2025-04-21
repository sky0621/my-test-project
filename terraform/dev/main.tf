provider "google" {
  project = var.project_id
  region  = var.region
}
provider "google-beta" {
  project = var.project_id
  region  = var.region
}
provider "kubernetes" {
  config_path    = pathexpand("~/.kube/config")
  config_context = "gke_my-study-project-dev_asia-northeast1-a_my-study-project-dev-gke"
}

module "random" {
  source = "../_modules/random"
}

# ----------------------------------------------------------------------------
# Enable APIs
# ----------------------------------------------------------------------------
module "service" {
  source     = "../_modules/service"
  project_id = var.project_id
}

# ----------------------------------------------------------------------------
# GitHub Actions Service Account
# ----------------------------------------------------------------------------
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

module "iam_workload_identity_pool_for_github_actions" {
  source                             = "../_modules/iam_workload_identity_pool_provider"
  project_id                         = var.project_id
  workload_identity_pool_id          = "gha-pool-${module.random.util_random_id}"
  workload_identity_pool_provider_id = "gha-prov-${module.random.util_random_id}"
}

module "service_account_iam_binding_for_github_actions" {
  source             = "../_modules/service_account_iam_binding"
  service_account_id = module.service_account_for_github_actions.name
  members            = ["principalSet://iam.googleapis.com/${module.iam_workload_identity_pool_for_github_actions.name}/attribute.repository/sky0621/my-test-project"]
}

# ----------------------------------------------------------------------------
# GKE Node Pool Service Account
# ----------------------------------------------------------------------------
module "service_account_for_gke_node" {
  source     = "../_modules/service_account"
  project_id = var.project_id
  account_id = "gke-node-sa"
  roles = [
    "artifactregistry.reader",
  ]
}

module "iam_workload_identity_pool_for_gke_node" {
  source                             = "../_modules/iam_workload_identity_pool_provider"
  project_id                         = var.project_id
  workload_identity_pool_id          = "gke-node-pool-${module.random.util_random_id}"
  workload_identity_pool_provider_id = "gke-node-prov-${module.random.util_random_id}"
}

module "service_account_iam_binding_for_gke_node" {
  source             = "../_modules/service_account_iam_binding"
  service_account_id = module.service_account_for_gke_node.name
  members            = ["principalSet://iam.googleapis.com/${module.iam_workload_identity_pool_for_gke_node.name}/attribute.repository/sky0621/my-test-project"]
}

# ----------------------------------------------------------------------------
# Kubernetes Service Account
# ----------------------------------------------------------------------------
module "service_account_for_k8s" {
  source     = "../_modules/service_account"
  project_id = var.project_id
  account_id = "k8s-sa"
  roles = [
    "artifactregistry.reader",
  ]
}

module "kubernetes_service_account" {
  source = "../_modules/kubernetes_service_account"
  name   = "ar-puller"
  email  = module.service_account_for_k8s.email
}

resource "google_service_account_iam_member" "workload_identity_binding" {
  service_account_id = module.service_account_for_k8s.name
  role               = "roles/iam.workloadIdentityUser"
  member = format(
    "serviceAccount:%s.svc.id.goog[%s/%s]",
    var.project_id,
    "default",
    module.kubernetes_service_account.name,
  )
}

# ----------------------------------------------------------------------------
# Network
# ----------------------------------------------------------------------------
module "vpc" {
  source     = "../_modules/vpc"
  project_id = var.project_id
  region     = var.region
}

# ----------------------------------------------------------------------------
# Artifact Registry
# ----------------------------------------------------------------------------
module "artifact_registry" {
  source        = "../_modules/artifact_registry"
  project_id    = var.project_id
  region        = var.region
  repository_id = var.repository_id
  depends_on    = ["module.service"]
}

# ----------------------------------------------------------------------------
# GKE
# ----------------------------------------------------------------------------
module "google_container_cluster" {
  source       = "../_modules/gke"
  project_id   = var.project_id
  location     = var.zone
  vpc_name     = module.vpc.vpc_name
  subnet_name  = module.vpc.subnet_name
  node_count   = 2
  machine_type = "n1-standard-1"
  depends_on   = ["module.service"]
}
