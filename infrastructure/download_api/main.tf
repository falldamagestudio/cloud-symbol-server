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

  name                  = "DownloadFile"
  description           = "Download File"
  runtime               = "go113"
  region                = var.function_region
  service_account_email = google_service_account.function_service_account.email

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.cloud_function_source_bucket.name
  source_archive_object = google_storage_bucket_object.cloud_function_bucket_object.name
  trigger_http          = true
  entry_point           = "DownloadFile"
  environment_variables = {
    GCP_PROJECT_ID           = var.project_id
    SYMBOL_STORE_BUCKET_HOST = "https://storage.googleapis.com/"
    SYMBOL_STORE_BUCKET_NAME = var.symbol_store_bucket_name
  }
}

# Create a service account. The cloud function will run in the context of this service account
resource "google_service_account" "function_service_account" {
  account_id   = "transfer"
  display_name = "Service Account used to run the transfer Cloud Function"
}

# Grant the cloud function's service account admin permissions to symbol store bucket
resource "google_storage_bucket_iam_member" "function_symbol_store_access" {
  depends_on = [google_service_account.function_service_account]
  bucket     = var.symbol_store_bucket_name
  role       = "roles/storage.admin"
  member     = "serviceAccount:${google_service_account.function_service_account.email}"
}

# Create a service account. This account can be used to invoke the function via HTTP.
resource "google_service_account" "invoke_function_service_account" {
  account_id   = "invoke-transfer"
  display_name = "Service account used to invoke the transfer Cloud Function via HTTP"
}

# Grant the cloud function's invocation service account permissions to launch the function via HTTP
resource "google_cloudfunctions_function_iam_member" "function_invoker" {
  depends_on = [google_service_account.invoke_function_service_account]

  project        = google_cloudfunctions_function.function.project
  region         = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name

  role   = "roles/cloudfunctions.invoker"
  member = "serviceAccount:${google_service_account.invoke_function_service_account.email}"
}

# Create an IAM entry for invoking the function
# This IAM entry allows anyone to invoke the function via HTTP, without being authenticated
#
# It would be ideal to have the function always require authentication, but that will be for later
# The problematic bit is: how do we embed a suitable authentication into the Unreal client?
# If we solve that, we can remove this IAM entry.
#
# The main risk posed by allowing this to be called unauthenticated is that an external party
#   could flood our internal database, driving up cost and making it hard for us to understand
#   our own data
resource "google_cloudfunctions_function_iam_member" "allow_unauthenticated_invocation" {
  project        = google_cloudfunctions_function.function.project
  region         = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}