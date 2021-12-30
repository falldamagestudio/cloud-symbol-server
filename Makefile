.PHONY: default-env-to-test
.PHONY: deploy deploy-core deploy_download_api
.PHONY: destroy
.PHONY: start-local-cloud-storage stop-local-cloud-storage
.PHONY: run-local-firebase-emulators
.PHONY: run-local-download-api
.PHONY: test-local test-local-download-api

default-env-to-test:
ifndef ENV
ENV:=test
endif

#########################################################
# Remote (connected to GCP) commands
#########################################################

deploy-core: default-env-to-test
	cd environments/$(ENV)/core && terraform init && terraform apply -auto-approve

deploy-download-api: default-env-to-test
	cd environments/$(ENV)/download_api && terraform init && terraform apply -auto-approve

deploy: default-env-to-test deploy-core deploy-download-api

destroy: default-env-to-test
	cd environments/$(ENV)/core && terraform destroy

#########################################################
# Local (emulator) commands
#########################################################

run-local-firebase-emulators:
	cd firebase && firebase emulators:start --project=demo-cloud-symbol-store --import state --export-on-exit

run-local-download-api:
	cd download-api/cmd \
	&&	GCP_PROJECT_ID=demo-cloud-symbol-store \
		FIRESTORE_EMULATOR_HOST=localhost:8082 \
		STORAGE_EMULATOR_HOST=localhost:9199 \
		SYMBOL_STORE_BUCKET_NAME=default-bucket \
		PORT=8083 \
		go run main.go

test-local-download-api:
	cd download-api \
	&&	PORT=8083 \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-store/download-api

test-local: test-local-download-api
