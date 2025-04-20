# 環境ごとに変わる
terraform {
  backend "gcs" {
    bucket = "tf-state-my-test-project-prd-73177136639b"
  }
}
