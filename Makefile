.PHONY: default-env-to-test
.PHONY: test test-download-api
.PHONY: deploy deploy-core deploy_download_api
.PHONY: destroy
.PHONY: start-local-cloud-storage stop-local-cloud-storage

default-env-to-test:
ifndef ENV
ENV:=test
endif

test-download-api:
	cd download-api && go test -timeout 30s download-api

test: default-env-to-test test-download-api

deploy-core: default-env-to-test
	cd environments/$(ENV)/core && terraform init && terraform apply -auto-approve

deploy-download-api: default-env-to-test
	cd environments/$(ENV)/download_api && terraform init && terraform apply -auto-approve

deploy: default-env-to-test deploy-core deploy-download-api

destroy: default-env-to-test
	cd environments/$(ENV)/core && terraform destroy

start-local-cloud-storage:
	docker run -d --name fake-gcs-server --rm -p 9000:4443 -v ${PWD}/example-store:/data fsouza/fake-gcs-server -scheme http

stop-local-cloud-storage:
	docker stop fake-gcs-server
