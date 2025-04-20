# terraform

## env

```
❯ terraform version
Terraform v1.11.4
on darwin_arm64
```

## auth

```
gcloud auth login
gcloud auth application-default login
gcloud auth application-default set-quota-project [project_id]
```

※ `gcloud config list` で現在接続対象としている GCP プロジェクトを要確認！

## gcloud config

```
gcloud config configurations create [name]
gcloud config set core/account [mail]
gcloud config set core/project [project_id]
gcloud config set compute/region asia-northeast1
gcloud config set compute/zone asia-northeast1-a
```

## create remote store

```
gsutil mb -l asia-northeast1 gs://tf-state-my-test-project-production-73177136639b
```

## operation

```
terraform init --upgrade
```

```
terraform fmt --recursive
terraform validate
tflint
```

- tflint : https://github.com/terraform-linters/tflint

```
terraform plan --out plan.tfplan
```

```
terraform apply "plan.tfplan"
```
