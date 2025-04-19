output "vpc_name" {
  value = google_compute_network.default.name
}

output "subnet_name" {
  value = google_compute_subnetwork.default.name
}
