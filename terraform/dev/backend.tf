# 環境ごとに変わる
terraform {
  backend "gcs" {
    bucket = "tf-state-my-test-project-dev-84fa0505feba"
  }
}
