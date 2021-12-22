.PHONY: default-env-to-test
.PHONY: test
.PHONY: deploy deploy-core
.PHONY: destroy

default-env-to-test:
ifndef ENV
ENV:=test
endif

test: default-env-to-test

deploy-core: default-env-to-test
	cd environments/$(ENV)/core && terraform init && terraform apply -auto-approve

deploy: default-env-to-test deploy-core

destroy: default-env-to-test
	cd environments/$(ENV)/core && terraform destroy
