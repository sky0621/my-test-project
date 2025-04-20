resource "google_service_account" "default" {
  project    = var.project_id
  account_id = var.account_id
}

resource "google_project_iam_member" "default" {
  for_each = var.roles

  project = var.project_id
  member  = "serviceAccount:${google_service_account.default.email}"
  role    = "roles/${each.value}"
}
