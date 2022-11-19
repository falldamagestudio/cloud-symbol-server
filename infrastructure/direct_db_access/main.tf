# Create a service account. This is used for direct DB acess
resource "google_service_account" "db_user_service_account" {
  account_id   = "direct-db-access"
  display_name = "Service Account used for direct DB access"
}

locals {
  # Cloud SQL IAM requires DB usernames to be max 63 characters in length
  # Service account email addresses are often longer than this; therefore, Cloud SQL
  #   want usernames that are truncated, like:
  #   name@project.iam.gserviceaccount.com => name@project.iam
  db_username = trimsuffix(google_service_account.db_user_service_account.email, ".gserviceaccount.com")
}

# Create an SQL IAM account for direct DB access's service account
# Reference: https://binx.io/2021/05/19/how-to-connect-to-a-cloudsql-with-iam-authentication/
resource "google_sql_user" "db_user" {
  name     = local.db_username
  instance = var.database_name
  type     = "CLOUD_IAM_SERVICE_ACCOUNT"
}

# Grant the direct DB access's service account access to the database
resource "google_project_iam_member" "db_user_cloudsql_client" {
  project = var.project_id
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.db_user_service_account.email}"
}

# Grant the direct DB access's service account access to the database
resource "google_project_iam_member" "db_user_cloudsql_instance_user" {
  project = var.project_id
  role    = "roles/cloudsql.instanceUser"
  member  = "serviceAccount:${google_service_account.db_user_service_account.email}"
}
