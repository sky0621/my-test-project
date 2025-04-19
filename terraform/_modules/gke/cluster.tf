resource "google_container_cluster" "default" {
  project  = var.project_id
  location = var.region
  name     = "${var.project_id}-gke"

  remove_default_node_pool = true
  initial_node_count       = 1

  network    = var.vpc_name
  subnetwork = var.subnet_name
}