.PHONY: deploy deploy-core deploy-database deploy-db-migrations deploy-backend-api deploy-firebase-and-frontend
.PHONY: remove-db-migrations
.PHONY: destroy
.PHONY: test test-backend-api test-cli

.PHONY: run-local-firebase-emulators
.PHONY: run-local-backend-api
.PHONY: test-local test-local-backend-api test-local-cli

.PHONY: generate-db-models
.PHONY: generate-apis generate-go-server-api generate-go-client-api generate-csharp-client-api generate-typescript-client-api

.PHONY: build-cli

OPENAPI_GENERATOR_VERSION:=v6.2.1

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
	./binaries/cloud_sql_proxy -instances "$(shell jq -r ".cloudSQLInstance" < $(ENV)/config.json)=tcp:5430" -fd_rlimit 1024 -enable_iam_login -credential_file=$(ENV)/database/google_application_credentials.json & echo "$$!" > db_migration_proxy.pid
	# We are not sure how long it takes for the proxy to start; we guess that 2 seconds should be enough
	sleep 2
	# Both fail and success paths from migration result in killing the proxy as well
	migrate -source "file://./db-migrations" -database "$(shell jq -r ".dbMigrationEndpoint" < $(ENV)/config.json)" -verbose up || (cat db_migration_proxy.pid | xargs kill && rm db_migration_proxy.pid && exit 1) && (cat db_migration_proxy.pid | xargs kill && rm db_migration_proxy.pid && exit 0)

remove-db-migrations:
	# Run Cloud SQL proxy in the background
	# We don't know whether or not the user already has a proxy running
	# If this is a first-time deploy via the 'deploy' Makefile target, the user has not been able to start a proxy
	# Because of this, we run a short-lived proxy just for this command
	./binaries/cloud_sql_proxy -instances "$(shell jq -r ".cloudSQLInstance" < $(ENV)/config.json)=tcp:5430" -fd_rlimit 1024 -enable_iam_login -credential_file=$(ENV)/database/google_application_credentials.json & echo "$$!" > db_migration_proxy.pid
	# We are not sure how long it takes for the proxy to start; we guess that 2 seconds should be enough
	sleep 2
	# Both fail and success paths from migration result in killing the proxy as well
	migrate -source "file://./db-migrations" -database "$(shell jq -r ".dbMigrationEndpoint" < $(ENV)/config.json)" -verbose down -all || (cat db_migration_proxy.pid | xargs kill && rm db_migration_proxy.pid && exit 1) && (cat db_migration_proxy.pid | xargs kill && rm db_migration_proxy.pid && exit 0)

deploy-database:
	cd $(ENV)/database && terraform init && terraform apply -auto-approve

deploy-backend-api:
	cd $(ENV)/backend_api && terraform init && terraform apply -auto-approve

inject-cli-binaries-into-frontend:
	cp cli/cloud-symbol-server-cli/bin/Release/net6.0/linux-x64/publish/cloud-symbol-server-cli \
		firebase/frontend/public/cloud-symbol-server-cli-linux
	cp cli/cloud-symbol-server-cli/bin/Release/net6.0/win-x64/publish/cloud-symbol-server-cli.exe \
		firebase/frontend/public/cloud-symbol-server-cli-win64.exe

deploy-firebase-and-frontend: build-cli inject-cli-binaries-into-frontend
	cd firebase/frontend \
		&&	npm install \
		&&	VUE_APP_FIREBASE_CONFIG='$(shell cat $(ENV)/firebase/frontend/firebase-config.json)' \
			VUE_APP_BACKEND_API_ENDPOINT="$(shell jq -r ".backendAPIEndpoint" < $(ENV)/config.json)" \
			VUE_APP_VERSION="$(VERSION)" \
			npm run build
	cd firebase && firebase deploy --project="$(shell jq -r ".gcpProjectId" < $(ENV)/firebase/config.json)"

deploy: deploy-core deploy-database deploy-db-migrations deploy-backend-api deploy-firebase-and-frontend

destroy:
	cd $(ENV)/core && terraform destroy

test-backend-api:
	cd backend-api/test \
	&&	BACKEND_API_ENDPOINT="$(shell jq -r ".backendAPIEndpoint" < $(ENV)/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < $(ENV)/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < $(ENV)/test-credentials.json)" \
		go test -timeout 120s github.com/falldamagestudio/cloud-symbol-server/backend-api/test -count=1

test-cli:
	cd cli \
	&&	BACKEND_API_ENDPOINT="$(shell jq -r ".backendAPIEndpoint" < $(ENV)/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < $(ENV)/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < $(ENV)/test-credentials.json)" \
		dotnet test

test: test-backend-api test-cli

#########################################################
# Local (emulator) commands
#########################################################

run-local-sql-auth-proxy:
	./binaries/cloud_sql_proxy -instances "$(shell jq -r ".cloudSQLInstance" < $(ENV)/config.json)=tcp:5432" -fd_rlimit 1024 -enable_iam_login -credential_file=$(ENV)/database/google_application_credentials.json

run-local-psql:
	psql "host=127.0.0.1 sslmode=disable dbname=cloud_symbol_server user=$(shell jq -r ".psqlUser" < $(ENV)/config.json)"

run-local-backend-api:
	cd backend-api/cmd \
	&&	GCP_PROJECT_ID=test-cloud-symbol-server \
		SYMBOL_STORE_BUCKET_NAME="$(shell jq -r ".symbol_store_bucket_name" < environments/local/core/config.json)" \
		CLOUD_SQL_INSTANCE="$(shell jq -r ".cloudSQLInstance" < $(ENV)/config.json)" \
		CLOUD_SQL_USER="$(shell jq -r ".cloudSQLAdminUser" < $(ENV)/config.json)" \
		PORT=8084 \
		GOOGLE_APPLICATION_CREDENTIALS="../../environments/local/backend_api/google_application_credentials.json" \
		go run main.go

run-local-frontend:
	cd firebase/frontend \
	&&	npm install \
	&&	VUE_APP_FIREBASE_CONFIG='$(shell cat environments/local/firebase/frontend/firebase-config.json)' \
		VUE_APP_BACKEND_API_ENDPOINT="$(shell jq -r ".backendAPIEndpoint" < environments/local/config.json)" \
		VUE_APP_VERSION="$(VERSION)" \
		npm run serve

test-local-backend-api:
	cd backend-api/test \
	&&	BACKEND_API_ENDPOINT="$(shell jq -r ".backendAPIEndpoint" < environments/local/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < environments/local/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < environments/local/test-credentials.json)" \
		go test -timeout 120s github.com/falldamagestudio/cloud-symbol-server/backend-api/test -count=1

test-local-cli:
	cd cli \
	&&	BACKEND_API_ENDPOINT="$(shell jq -r ".backendAPIEndpoint" < environments/local/config.json)" \
		TEST_EMAIL="$(shell jq -r ".email" < environments/local/test-credentials.json)" \
		TEST_PAT="$(shell jq -r ".pat" < environments/local/test-credentials.json)" \
		dotnet test

test-local: test-local-backend-api test-local-cli

#########################################################
# DB model regeneration commands
#########################################################

generate-db-models:
	# This will connect to the SQL db
	# The user should have the SQL Auth proxy running (via `make run-local-sql-auth-proxy`) first
	sqlboiler psql --output backend-api/generated/sql-db-models --wipe

#########################################################
# API regeneration commands
#########################################################

generate-apis: generate-go-server-api generate-go-client-api generate-csharp-client-api generate-typescript-client-api

generate-go-server-api:

	rm -rf backend-api/generated/go-server/go
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		--git-user-id=falldamagestudio \
		--git-repo-id=cloud-symbol-server/backend-api \
		-i /local/backend-api/admin-api.yaml \
		-g go-server \
		--additional-properties=enumClassPrefix=true,generateAliasAsModel=false \
		-o /local/backend-api/generated/go-server

generate-go-client-api:

	rm -rf backend-api/generated/go-client/docs
	rm -rf backend-api/generated/go-client/*.go
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		--git-user-id=falldamagestudio \
		--git-repo-id=cloud-symbol-server/backend-api \
		-i /local/backend-api/admin-api.yaml \
		-g go \
		--additional-properties=enumClassPrefix=true,generateAliasAsModel=false \
		-o /local/backend-api/generated/go-client

generate-csharp-client-api:

	rm -rf cli/generated/BackendAPI/src
	rm -rf cli/generated/BackendAPI/docs
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		-i /local/backend-api/admin-api.yaml \
		-g csharp-netcore \
		--additional-properties=netCoreProjectFile=true,library=httpclient,packageName=BackendAPI,generateAliasAsModel=false \
		-o /local/cli/generated/BackendAPI

generate-typescript-client-api:

#	rm -rf firebase/frontend/src/generated/
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
		generate \
		-i /local/backend-api/admin-api.yaml \
		-g typescript-axios \
		--additional-properties=generateAliasAsModel=false \
		-o /local/firebase/frontend/src/generated

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

