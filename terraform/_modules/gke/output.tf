output "endpoint" {
  value = google_container_cluster.default.endpoint
}

output "identity_namespace" {
  value = google_container_cluster.default.workload_identity_config[0].workload_pool
}

output "name" {
  value = google_container_cluster.default.name
}
