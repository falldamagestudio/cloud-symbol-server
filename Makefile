.PHONY: deploy deploy-core deploy-download-api deploy-admin-api deploy-firebase-and-frontend
.PHONY: destroy
.PHONY: test test-download-api test-admin-api

.PHONY: run-local-firebase-emulators
.PHONY: run-local-download-api run-local-admin-api
.PHONY: test-local test-local-download-api test-local-admin-api

.PHONY: generate-apis generate-go-server-api generate-go-client-api generate-csharp-client-api

.PHONY: test-cli build-cli

ifndef ENV
ENV:=environments/test
endif

# Handle version numbers
# The default version is going to be a number on the form 1.2.5-local
#   where 1.2.5 comes from version.json
# When necessary, both prefix and suffix can be overridden (this is typically done in CI pipelines)

ifndef VERSION_PREFIX
VERSION_PREFIX:=$(shell jq -r ".version" < "version.json")
endif

ifndef VERSION_SUFFIX
VERSION_SUFFIX:=local
endif

ifneq "$(VERSION_SUFFIX)" ""
VERSION:=$(VERSION_PREFIX)-$(VERSION_SUFFIX)
else
VERSION:=$(VERSION_PREFIX)
endif

#########################################################
# Remote (connected to GCP) commands
#########################################################

deploy-core:
	cd $(ENV)/core && terraform init && terraform apply -auto-approve

deploy-download-api:
	cd $(ENV)/download_api && terraform init && terraform apply -auto-approve

deploy-admin-api:
	cd $(ENV)/admin_api && terraform init && terraform apply -auto-approve

inject-cli-binaries-into-frontend:
	cp cli/cloud-symbol-server-cli/bin/Release/net6.0/linux-x64/publish/cloud-symbol-server-cli \
		firebase/frontend/public/cloud-symbol-server-cli-linux
	cp cli/cloud-symbol-server-cli/bin/Release/net6.0/win-x64/publish/cloud-symbol-server-cli.exe \
		firebase/frontend/public/cloud-symbol-server-cli-win64.exe

deploy-firebase-and-frontend: build-cli inject-cli-binaries-into-frontend
	cd firebase/frontend \
		&&	npm install \
		&&	VUE_APP_FIREBASE_CONFIG='$(shell cat $(ENV)/firebase/frontend/firebase-config.json)' \
			VUE_APP_ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < $(ENV)/config.json)" \
			VUE_APP_DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < $(ENV)/config.json)" \
			VUE_APP_VERSION="$(VERSION)" \
			npm run build
	cd firebase && firebase deploy --project="$(shell jq -r ".gcpProjectId" < $(ENV)/firebase/config.json)"

deploy: deploy-core deploy-download-api deploy-admin-api deploy-firebase-and-frontend

destroy:
	cd $(ENV)/core && terraform destroy

test-download-api:
	cd download-api \
	&&	DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < $(ENV)/config.json)" \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/download-api

test-admin-api:
	cd admin-api \
	&&	ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < $(ENV)/config.json)" \
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
		PORT=8083 \
		go run main.go

run-local-admin-api:
	cd admin-api/cmd \
	&&	GCP_PROJECT_ID=demo-cloud-symbol-server \
		FIRESTORE_EMULATOR_HOST=localhost:8082 \
		STORAGE_EMULATOR_HOST=localhost:9199 \
		SYMBOL_STORE_BUCKET_NAME=default-bucket \
		PORT=8084 \
		go run main.go

run-local-frontend:
	cd firebase/frontend \
	&&	npm install \
	&&	VUE_APP_FIREBASE_CONFIG='$(shell cat environments/local/firebase/frontend/firebase-config.json)' \
		VUE_APP_FIRESTORE_EMULATOR_PORT=8082 \
		VUE_APP_AUTH_EMULATOR_URL=http://localhost:9099 \
		VUE_APP_ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < environments/local/config.json)" \
		VUE_APP_DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < environments/local/config.json)" \
		VUE_APP_VERSION="$(VERSION)" \
		npm run serve

test-local-download-api:
	cd download-api \
	&&	DOWNLOAD_API_ENDPOINT=http://localhost:8083 \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/download-api

test-local-admin-api:
	cd admin-api \
	&&	ADMIN_API_ENDPOINT=http://localhost:8084 \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/admin-api -count=1

test-local: test-local-download-api test-local-admin-api

#########################################################
# API regeneration commands
#########################################################

generate-apis: generate-go-server-api generate-go-client-api generate-csharp-client-api

generate-go-server-api:

	rm -rf admin-api/generated/go-server/go
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
		--additional-properties=generateAliasAsModel=true \
		-o /local/admin-api/generated/go-server

generate-go-client-api:

	rm -rf admin-api/generated/go-client/docs
	rm -rf admin-api/generated/go-client/*.go
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli \
		generate \
		--git-user-id=falldamagestudio \
		--git-repo-id=cloud-symbol-server/admin-api \
		-i /local/admin-api/admin-api.yaml \
		-g go \
		--additional-properties=generateAliasAsModel=true \
		-o /local/admin-api/generated/go-client

generate-csharp-client-api:

	rm -rf cli/generated/BackendAPI/src
	rm -rf cli/generated/BackendAPI/docs
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli \
		generate \
		-i /local/admin-api/admin-api.yaml \
		-g csharp-netcore \
		--additional-properties=netCoreProjectFile=true,library=httpclient,packageName=BackendAPI,generateAliasAsModel=true \
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
		/p:VersionPrefix=$(VERSION_PREFIX) \
		/p:VersionSuffix=$(VERSION_SUFFIX) \
	&& dotnet publish \
		--runtime win-x64 \
		--self-contained \
		--configuration Release \
		/p:VersionPrefix=$(VERSION_PREFIX) \
		/p:VersionSuffix=$(VERSION_SUFFIX)

