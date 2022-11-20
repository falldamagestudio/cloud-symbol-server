.PHONY: deploy deploy-core deploy-database deploy-db-migrations deploy-download-api deploy-admin-api deploy-firebase-and-frontend
.PHONY: destroy
.PHONY: test test-download-api test-admin-api test-cli

.PHONY: run-local-firebase-emulators
.PHONY: run-local-download-api run-local-admin-api
.PHONY: test-local test-local-download-api test-local-admin-api test-local-cli

.PHONY: generate-db-models
.PHONY: generate-apis generate-go-server-api generate-go-client-api generate-csharp-client-api

.PHONY: build-cli

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

deploy-db-migrations:
	# Run Cloud SQL proxy in the background
	# We don't know whether or not the user already has a proxy running
	# If this is a first-time deploy via the 'deploy' Makefile target, the user has not been able to start a proxy
	# Because of this, we run a short-lived proxy just for this command
	./binaries/cloud_sql_proxy -instances "$(shell jq -r ".cloudSQLProxyEndpoint" < $(ENV)/config.json)=tcp:5430" -fd_rlimit 1024 -enable_iam_login -credential_file=$(ENV)/database/google_application_credentials.json & echo "$$!" > db_migration_proxy.pid
	# We are not sure how long it takes for the proxy to start; we guess that 2 seconds should be enough
	sleep 2
	# Both fail and success paths from migration result in killing the proxy as well
	migrate -source "file://./db-migrations" -database "$(shell jq -r ".dbMigrationEndpoint" < $(ENV)/config.json)" -verbose up || (cat db_migration_proxy.pid | xargs kill && rm db_migration_proxy.pid && exit 1) && (cat db_migration_proxy.pid | xargs kill && rm db_migration_proxy.pid && exit 0)

deploy-database:
	cd $(ENV)/database && terraform init && terraform apply -auto-approve

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

deploy: deploy-core deploy-database deploy-db-migrations deploy-download-api deploy-admin-api deploy-firebase-and-frontend

destroy:
	cd $(ENV)/core && terraform destroy

test-download-api:
	cd download-api/test \
	&&	ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < $(ENV)/config.json)" \
		DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < $(ENV)/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < $(ENV)/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < $(ENV)/test-credentials.json)" \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/download-api/test -count=1

test-admin-api:
	cd admin-api/test \
	&&	ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < $(ENV)/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < $(ENV)/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < $(ENV)/test-credentials.json)" \
		go test -timeout 60s github.com/falldamagestudio/cloud-symbol-server/admin-api/test -count=1

test-cli:
	cd cli \
	&&	ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < $(ENV)/config.json)" \
		DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < $(ENV)/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < $(ENV)/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < $(ENV)/test-credentials.json)" \
		dotnet test

test: test-download-api test-admin-api test-cli

#########################################################
# Local (emulator) commands
#########################################################

run-local-sql-auth-proxy:
	./binaries/cloud_sql_proxy -instances "$(shell jq -r ".cloudSQLProxyEndpoint" < $(ENV)/config.json)=tcp:5432" -fd_rlimit 1024 -enable_iam_login -credential_file=$(ENV)/database/google_application_credentials.json

run-local-psql:
	psql "host=127.0.0.1 sslmode=disable dbname=cloud_symbol_server user=$(shell jq -r ".psqlUser" < $(ENV)/config.json)"

run-local-download-api:
	cd download-api/cmd \
	&&	GCP_PROJECT_ID=test-cloud-symbol-server \
		SYMBOL_STORE_BUCKET_NAME="$(shell jq -r ".symbol_store_bucket_name" < environments/local/core/config.json)" \
		PORT=8083 \
		GOOGLE_APPLICATION_CREDENTIALS="../../environments/local/download_api/google_application_credentials.json" \
		go run main.go

run-local-admin-api:
	cd admin-api/cmd \
	&&	GCP_PROJECT_ID=test-cloud-symbol-server \
		SYMBOL_STORE_BUCKET_NAME="$(shell jq -r ".symbol_store_bucket_name" < environments/local/core/config.json)" \
		PORT=8084 \
		GOOGLE_APPLICATION_CREDENTIALS="../../environments/local/admin_api/google_application_credentials.json" \
		go run main.go

run-local-frontend:
	cd firebase/frontend \
	&&	npm install \
	&&	VUE_APP_FIREBASE_CONFIG='$(shell cat environments/local/firebase/frontend/firebase-config.json)' \
		VUE_APP_ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < environments/local/config.json)" \
		VUE_APP_DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < environments/local/config.json)" \
		VUE_APP_VERSION="$(VERSION)" \
		npm run serve

test-local-download-api:
	cd download-api/test \
	&&	ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < environments/local/config.json)" \
		DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < environments/local/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < environments/local/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < environments/local/test-credentials.json)" \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-server/download-api/test -count=1

test-local-admin-api:
	cd admin-api/test \
	&&	ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < environments/local/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < environments/local/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < environments/local/test-credentials.json)" \
		go test -timeout 60s github.com/falldamagestudio/cloud-symbol-server/admin-api/test -count=1

test-local-cli:
	cd cli \
	&&	ADMIN_API_ENDPOINT="$(shell jq -r ".adminAPIEndpoint" < environments/local/config.json)" \
		DOWNLOAD_API_ENDPOINT="$(shell jq -r ".downloadAPIEndpoint" < environments/local/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < environments/local/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < environments/local/test-credentials.json)" \
		dotnet test

test-local: test-local-download-api test-local-admin-api test-local-cli

#########################################################
# DB model regeneration commands
#########################################################

generate-db-models:
	# This will connect to the SQL db
	# The user should have the SQL Auth proxy running (via `make run-local-sql-auth-proxy`) first
	sqlboiler psql --output admin-api/generated/sql-db-models --wipe

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

build-cli:
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

