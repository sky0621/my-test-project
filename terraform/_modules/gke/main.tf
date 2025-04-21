resource "google_container_cluster" "default" {
  project  = var.project_id
  location = var.location
  name     = "${var.project_id}-gke"

  workload_identity_config {
    workload_pool = "${var.project_id}.svc.id.goog"
  }

  remove_default_node_pool = true
  initial_node_count       = 1

  network    = var.vpc_name
  subnetwork = var.subnet_name

  deletion_protection = false
}

data "google_container_engine_versions" "gke_version" {
  location       = var.location
  version_prefix = "1.27."
}
resource "google_container_node_pool" "default_pool" {
  project  = var.project_id
  location = var.location
  name     = "${var.project_id}-gke-pool"
  cluster  = google_container_cluster.default.name

  version    = data.google_container_engine_versions.gke_version.release_channel_default_version["STABLE"]
  node_count = var.node_count

  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]

    labels = {
      env = var.project_id
    }

    # preemptible  = true
    machine_type = var.machine_type
    tags         = ["gke-node", "${var.project_id}-gke"]
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }
}