# 環境ごとに変わる
terraform {
  backend "gcs" {
    bucket = "tf-state-my-test-project-production-73177136639b"
  }
}
