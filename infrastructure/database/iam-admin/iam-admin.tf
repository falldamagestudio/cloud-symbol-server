# Create a service account. This is used for admin DB acess
resource "google_service_account" "iam_admin_user_service_account" {
  account_id   = "iam-admin"
  display_name = "Service Account used for admin DB access"
}

locals {
  # Cloud SQL IAM requires DB usernames to be max 63 characters in length
  # Service account email addresses are often longer than this; therefore, Cloud SQL
  #   want usernames that are truncated, like:
  #   name@project.iam.gserviceaccount.com => name@project.iam
  iam_admin_username = trimsuffix(google_service_account.iam_admin_user_service_account.email, ".gserviceaccount.com")
}

# Create an SQL IAM account for admin DB access's service account
# Reference: https://binx.io/2021/05/19/how-to-connect-to-a-cloudsql-with-iam-authentication/
resource "google_sql_user" "iam_admin_user" {
  name     = local.iam_admin_username
  instance = var.database_instance_name
  type     = "CLOUD_IAM_SERVICE_ACCOUNT"
}

# Grant the admin DB access's service account permission to connect to the database instance via Cloud SQL Auth proxy
resource "google_project_iam_member" "iam_admin_user_cloudsql_client" {
  project = var.project_id
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.iam_admin_user_service_account.email}"
}

# Grant the admin DB access's service account permission to log in to the database instance
resource "google_project_iam_member" "iam_admin_user_cloudsql_instance_user" {
  project = var.project_id
  role    = "roles/cloudsql.instanceUser"
  member  = "serviceAccount:${google_service_account.iam_admin_user_service_account.email}"
}

# Allow the admin DB access's service account to create & use objects within schema
# Reference: https://registry.terraform.io/providers/cyrilgdn/postgresql/latest/docs/resources/postgresql_grant#examples
# Reference: https://www.postgresql.org/docs/current/ddl-priv.html
# Reference: https://dba.stackexchange.com/questions/117109/how-to-manage-default-privileges-for-users-on-a-database-vs-schema
resource "postgresql_grant_role" "function_db_user_allow_readwrite" {
  role        = google_sql_user.iam_admin_user.name
  # TODO: source role name from state
  grant_role  = "cloud_symbol_server_admin"
}
