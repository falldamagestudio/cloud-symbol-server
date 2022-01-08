.PHONY: deploy deploy-core deploy-download-api deploy-upload-api deploy-firebase-and-frontend
.PHONY: destroy
.PHONY: test test-download-api test-upload-api

.PHONY: run-local-firebase-emulators
.PHONY: run-local-download-api run-local-upload-api
.PHONY: test-local test-local-download-api test-local-upload-api

.PHONY: build-csharp-cli

ifndef ENV
ENV:=test
endif

#########################################################
# Remote (connected to GCP) commands
#########################################################

deploy-core:
	echo apa3
	echo $(ENV)
	cd environments/$(ENV)/core && terraform init && terraform apply -auto-approve

deploy-download-api:
	cd environments/$(ENV)/download_api && terraform init && terraform apply -auto-approve

deploy-upload-api:
	cd environments/$(ENV)/upload_api && terraform init && terraform apply -auto-approve

deploy-firebase-and-frontend:
	cd firebase/frontend \
		&&	VUE_APP_FIREBASE_CONFIG='$(shell cat environments/$(ENV)/firebase/frontend/firebase-config.json)' \
			VUE_APP_DOWNLOAD_API_PROTOCOL="$(shell jq -r ".downloadAPIProtocol" < environments/$(ENV)/config.json)" \
			VUE_APP_DOWNLOAD_API_HOST="$(shell jq -r ".downloadAPIHost" < environments/$(ENV)/config.json)" \
			npm run build
	cd firebase && firebase deploy --project="$(shell jq -r ".gcpProjectId" < environments/$(ENV)/firebase/config.json)"

deploy: deploy-core deploy-download-api deploy-upload-api deploy-firebase-and-frontend

destroy:
	cd environments/$(ENV)/core && terraform destroy

test-download-api:
	cd download-api \
	&&	DOWNLOAD_API_PROTOCOL="$(shell jq -r ".downloadAPIProtocol" < environments/$(ENV)/config.json)" \
		DOWNLOAD_API_HOST="$(shell jq -r ".downloadAPIHost" < environments/$(ENV)/config.json)" \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-store/download-api

test-upload-api:
	cd upload-api \
	&&	UPLOAD_API_PROTOCOL="$(shell jq -r ".uploadAPIProtocol" < environments/$(ENV)/config.json)" \
		UPLOAD_API_HOST="$(shell jq -r ".uploadAPIHost" < environments/$(ENV)/config.json)" \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-store/upload-api

test: test-download-api test-upload-api

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

run-local-upload-api:
	cd upload-api/cmd \
	&&	GCP_PROJECT_ID=demo-cloud-symbol-store \
		FIRESTORE_EMULATOR_HOST=localhost:8082 \
		STORAGE_EMULATOR_HOST=localhost:9199 \
		SYMBOL_STORE_BUCKET_NAME=default-bucket \
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
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-store/download-api

test-local-upload-api:
	cd upload-api \
	&&	UPLOAD_API_PROTOCOL=http \
		UPLOAD_API_HOST=localhost:8084 \
		go test -timeout 30s github.com/falldamagestudio/cloud-symbol-store/upload-api -count=1

test-local: test-local-download-api test-local-upload-api

build-cli:
	cd cli \
	&& GOOS=windows GOARCH=amd64 go build -o cloud-symbol-store-cli.exe ./cmd \
	&& GOOS=linux GOARCH=amd64 go build -o cloud-symbol-store-cli ./cmd

###

generate-apis:
	rm -r upload-api/generated/go
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli \
		generate \
		--git-user-id=falldamagestudio \
		--git-repo-id=cloud-symbol-store/upload-api \
		-i /local/upload-api/upload-api.yaml \
		-g go-server \
		-o /local/upload-api/generated

	rm cli/generated/*.go
	rm -r cli/generated/docs
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli \
		generate \
		--git-user-id=falldamagestudio \
		--git-repo-id=cloud-symbol-store/cli \
		-i /local/upload-api/upload-api.yaml \
		-g go \
		-o /local/cli/generated

generate-apis-2:
#	rm csharp-cli/BackendAPI/*.go
#	rm -r csharp-cli/GeneratedAPI/docs
	docker run \
		--rm \
		-v "${PWD}:/local" \
		--user $(shell id -u):$(shell id -g) \
		openapitools/openapi-generator-cli \
		generate \
		-i /local/upload-api/upload-api.yaml \
		-g csharp-netcore \
		--additional-properties=netCoreProjectFile=true,library=httpclient,packageName=BackendAPI \
		-o /local/csharp-cli/generated/BackendAPI


build-csharp-cli:
	cd csharp-cli \
	&& dotnet publish \
		--runtime linux-x64 \
		--self-contained \
		--configuration Release \
	&& dotnet publish \
		--runtime win-x64 \
		--self-contained \
		--configuration Release
