resource "google_compute_network" "default" {
  name                    = "${var.project_id}-vpc"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "${var.project_id}-subnet"
  region        = var.region
  network       = google_compute_network.default.name
  ip_cidr_range = "10.10.0.0/24"
}
