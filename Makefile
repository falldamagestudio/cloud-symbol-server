.PHONY: deploy deploy-core deploy-download-api deploy-admin-api deploy-firebase-and-frontend
.PHONY: destroy
.PHONY: test test-download-api test-admin-api

.PHONY: run-local-firebase-emulators
.PHONY: run-local-download-api run-local-admin-api
.PHONY: test-local test-local-download-api test-local-admin-api

.PHONY: generate-apis generate-server-api generate-client-api

.PHONY: test-cli build-cli

ifndef ENV
ENV:=test
endif

#########################################################
# Remote (connected to GCP) commands
#########################################################

deploy-core:
	cd environments/$(ENV)/core && terraform init && terraform apply -auto-approve

deploy-download-api:
	cd environments/$(ENV)/download_api && terraform init && terraform apply -auto-approve

deploy-admin-api:
	cd environments/$(ENV)/admin_api && terraform init && terraform apply -auto-approve

deploy-firebase-and-frontend:
	cd firebase/frontend \
		&&	VUE_APP_FIREBASE_CONFIG='$(shell cat environments/$(ENV)/firebase/frontend/firebase-config.json)' \
			VUE_APP_DOWNLOAD_API_PROTOCOL="$(shell jq -r ".downloadAPIProtocol" < environments/$(ENV)/config.json)" \
			VUE_APP_DOWNLOAD_API_HOST="$(shell jq -r ".downloadAPIHost" < environments/$(ENV)/config.json)" \
			npm run build
	cd firebase && firebase deploy --project="$(shell jq -r ".gcpProjectId" < environments/$(ENV)/firebase/config.json)"

deploy: deploy-core deploy-download-api deploy-admin-api deploy-firebase-and-frontend

destroy:
	cd environments/$(ENV)/core && terraform destroy

test-download-api:
	cd download-api \
	&&	DOWNLOAD_API_PROTOCOL="$(shell jq -r ".downloadAPIProtocol" < environments/$(ENV)/config.json)" \
		DOWNLOAD_API_HOST="$(shell jq -r ".downloadAPIHost" < environments/$(ENV)/config.json)" \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/download-api

test-admin-api:
	cd admin-api \
	&&	ADMIN_API_PROTOCOL="$(shell jq -r ".adminAPIProtocol" < environments/$(ENV)/config.json)" \
		ADMIN_API_HOST="$(shell jq -r ".adminAPIHost" < environments/$(ENV)/config.json)" \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/admin-api

test: test-download-api test-admin-api

#########################################################
# Local (emulator) commands
#########################################################

run-local-firebase-emulators:
	cd firebase && firebase emulators:start --project=demo-cloud-symbol-server --import state --export-on-exit

run-local-download-api:
	cd download-api/cmd \
	&&	GCP_PROJECT_ID=demo-cloud-symbol-server \
		FIRESTORE_EMULATOR_HOST=localhost:8082 \
		STORAGE_EMULATOR_HOST=localhost:9199 \
		SYMBOL_STORE_BUCKET_NAME=default-bucket \
		SYMBOL_SERVER_STORES=[\"example\"] \
		PORT=8083 \
		go run main.go

run-local-admin-api:
	cd admin-api/cmd \
	&&	GCP_PROJECT_ID=demo-cloud-symbol-server \
		FIRESTORE_EMULATOR_HOST=localhost:8082 \
		STORAGE_EMULATOR_HOST=localhost:9199 \
		SYMBOL_STORE_BUCKET_NAME=default-bucket \
		SYMBOL_SERVER_STORES=[\"example\"] \
		PORT=8084 \
		go run main.go

run-local-frontend:
	cd firebase/frontend \
	&&	VUE_APP_FIREBASE_CONFIG='$(shell cat environments/local/firebase/frontend/firebase-config.json)' \
		VUE_APP_FIRESTORE_EMULATOR_PORT=8082 \
		VUE_APP_AUTH_EMULATOR_URL=http://localhost:9099 \
		VUE_APP_DOWNLOAD_API_PROTOCOL="$(shell jq -r ".downloadAPIProtocol" < environments/local/config.json)" \
		VUE_APP_DOWNLOAD_API_HOST="$(shell jq -r ".downloadAPIHost" < environments/local/config.json)" \
		npm run serve

test-local-download-api:
	cd download-api \
	&&	DOWNLOAD_API_PROTOCOL=http \
		DOWNLOAD_API_HOST=localhost:8083 \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/download-api

test-local-admin-api:
	cd admin-api \
	&&	ADMIN_API_PROTOCOL=http \
		ADMIN_API_HOST=localhost:8084 \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/admin-api -count=1

test-local: test-local-download-api test-local-admin-api

#########################################################
# API regeneration commands
#########################################################

generate-apis: generate-server-api generate-client-api

generate-server-api:

	rm -r admin-api/generated/go
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli \
		generate \
		--git-user-id=falldamagestudio \
		--git-repo-id=cloud-symbol-server/admin-api \
		-i /local/admin-api/admin-api.yaml \
		-g go-server \
		-o /local/admin-api/generated

generate-client-api:

	rm -r cli/generated/BackendAPI/src
	rm -r cli/generated/BackendAPI/docs
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli \
		generate \
		-i /local/admin-api/admin-api.yaml \
		-g csharp-netcore \
		--additional-properties=netCoreProjectFile=true,library=httpclient,packageName=BackendAPI \
		-o /local/cli/generated/BackendAPI

#########################################################
# CLI commands
#########################################################

test-cli:
	cd cli \
	&& dotnet test

build-cli: test-cli
	cd cli \
	&& dotnet publish \
		--runtime linux-x64 \
		--self-contained \
		--configuration Release \
	&& dotnet publish \
		--runtime win-x64 \
		--self-contained \
		--configuration Release

