# Development and debugging

Development and debugging can be done in two configurations: either with the Cloud Functions logic running locally + database & storage in GCP ("local"), or with everything deployed to GCP ("test").

There are local emulators for Firebase, Cloud Firestore, and GCS, but they do not work well enough for this use case (where there is both Firebase-access and traditional GCP API access to the resources). Therefore, both local and test have database & storage in GCP.

Before you develop/debug, deploy this to GCP once. That creates all required resources in GCP.

# Set up credentials

* Create a key for the `download-api` Service Account. Place a copy of it at `environments/local/download_api/google_account_credentials.json`.
* Create a key for the `admin-api` Service Account. Place a copy of it at `environments/local/admin_api/google_account_credentials.json`.
* Create a user called `testuser` in Cloud Firestore, and define a PAT for that user. Place the email/pat settings in `environments/local/test-configuration.json`.

# Local development

* Edit code.
* Start admin & download APIs: `make run-local-admin-api`, `make run-local-development-api` in two shells.
* Perform manual tests against APIs, or run testsuite in a third shell: `make test-local`.
* Stop admin & download APIs

# Test development

* Edit code.
* Deploy to test environment in GCP: `make deploy`.
* Run testsuite: `make test`.
