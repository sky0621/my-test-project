resource "google_project_service" "service_usage" {
  project = var.project_id
  service = "serviceusage.googleapis.com"
}

resource "google_project_service" "iam_credentials" {
  project            = var.project_id
  service            = "iamcredentials.googleapis.com"
  disable_on_destroy = false
  depends_on         = ["google_project_service.service_usage"]
}

resource "google_project_service" "artifact_registry" {
  project            = var.project_id
  service            = "artifactregistry.googleapis.com"
  disable_on_destroy = false
  depends_on         = ["google_project_service.service_usage"]
}

resource "google_project_service" "container" {
  project            = var.project_id
  service            = "container.googleapis.com"
  disable_on_destroy = false
  depends_on         = ["google_project_service.service_usage"]
}
