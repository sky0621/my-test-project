resource "kubernetes_service_account" "default" {
  metadata {
    name      = var.name
    namespace = "default"
    annotations = {
      "iam.gke.io/gcp-service-account" = var.email
    }
  }
}
