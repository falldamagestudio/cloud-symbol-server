locals {
    zip_filename = "${path.module}/cloud_function_source.zip"
}

# Create a zip archive with the cloud function's source code
data "archive_file" "cloud_function_source_zip" {
  type        = "zip"
  source_dir  = var.source_path
  excludes    = [".git"]
  output_path = local.zip_filename
}

# Create a storage bucket for the cloud function's source code
resource "google_storage_bucket" "cloud_function_source_bucket" {
  name     = var.source_bucket_name
  location = var.source_bucket_location
}

# Upload the cloud function's source code to the storage bucket
resource "google_storage_bucket_object" "cloud_function_bucket_object" {
  name   = format("cloud_function_source.%s.zip", data.archive_file.cloud_function_source_zip.output_md5)
  bucket = google_storage_bucket.cloud_function_source_bucket.name
  source = local.zip_filename
}

# Deploy the cloud function
resource "google_cloudfunctions_function" "function" {
  depends_on = [google_storage_bucket_iam_member.function_symbol_store_access]

  name                  = "AdminAPI"
  description           = "Upload API"
  runtime               = "go113"
  region                = var.function_region
  service_account_email = google_service_account.function_service_account.email

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.cloud_function_source_bucket.name
  source_archive_object = google_storage_bucket_object.cloud_function_bucket_object.name
  trigger_http          = true
  entry_point           = "AdminAPI"
  environment_variables = {
    GCP_PROJECT_ID           = var.project_id
    SYMBOL_STORE_BUCKET_NAME = var.symbol_store_bucket_name
  }
}

# Create a service account. The cloud function will run in the context of this service account
resource "google_service_account" "function_service_account" {
  account_id   = "admin-api"
  display_name = "Service Account used to run the Upload API Cloud Function"
}

# Grant the cloud function's service account admin permissions to symbol store bucket
resource "google_storage_bucket_iam_member" "function_symbol_store_access" {
  depends_on = [google_service_account.function_service_account]
  bucket     = var.symbol_store_bucket_name
  role       = "roles/storage.admin"
  member     = "serviceAccount:${google_service_account.function_service_account.email}"
}

# Grant the cloud function's service account permission to create tokens
# This includes the permission to sign blobs (iam.serviceAccounts.signBlob)
#  and that is required in order to create signed URLs
# Reference: https://stackoverflow.com/a/57565326
resource "google_project_iam_member" "function_token_creation_permission" {
  depends_on = [google_service_account.function_service_account]
  project    = var.project_id
  role       = "roles/iam.serviceAccountTokenCreator"
  member     = "serviceAccount:${google_service_account.function_service_account.email}"
}

# Grant the cloud function's service account access to Firestore
# TODO: make permissions narrower - perhaps only roles/datastore.viewer?
resource "google_project_iam_member" "function_firestore_access" {
  depends_on = [google_service_account.function_service_account]
  project    = var.project_id
  role       = "roles/datastore.owner"
  member     = "serviceAccount:${google_service_account.function_service_account.email}"
}

# Create an IAM entry for invoking the function
# This IAM entry allows anyone to invoke the function via HTTP, without being authenticated
resource "google_cloudfunctions_function_iam_member" "allow_unauthenticated_invocation" {
  project        = google_cloudfunctions_function.function.project
  region         = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}