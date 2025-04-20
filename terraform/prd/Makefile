TF=terraform

.PHONY: tf-init
tf-init:
	@echo "Terraformの初期化を実行中..."
	$(TF) init --upgrade
	@echo "Terraformの初期化が完了しました。"

.PHONY: tf-validate
tf-validate:
	@echo "Terraformのバリデーションを実行中..."
	$(TF) validate
	@echo "Terraformのバリデーションが完了しました。"

.PHONY: tf-plan
tf-plan:
	@echo "Terraformのプランを実行中..."
	$(TF) plan -out=plan.tfplan
	@echo "Terraformのプランが完了しました。"

.PHONY: tf-apply
tf-apply:
	@echo "Terraformの適用を実行中..."
	$(TF) apply -auto-approve plan.tfplan
	@echo "Terraformの適用が完了しました。"

.PHOHY: tf-destroy
tf-destroy:
	@echo "Terraformの破棄を実行中..."
	$(TF) destroy -auto-approve
	@echo "Terraformの破棄が完了しました。"
