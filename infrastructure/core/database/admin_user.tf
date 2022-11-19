# The admin account will be granted a password in the form of a long hexadecimal string
resource "random_id" "admin_password" {
  byte_length = 16
}

# Create a built-in SQL account; this will automatically be given the cloudsuperuser role
# It can be used to manage databases, roles and permissions for IAM accounts
resource "google_sql_user" "admin_user" {
  name     = "admin"
  password = random_id.admin_password.hex
  instance = var.name
  type     = "BUILT_IN"
}

# # Grant the cloud function's service account access to the database
# resource "google_project_iam_member" "admin_user_cloudsql_client" {
#   project = var.project_id
#   role    = "roles/cloudsql.client"
#   member  = "serviceAccount:${google_sql_user.admin_user.name}"
# }

# # Grant the cloud function's service account access to the database
# resource "google_project_iam_member" "admin_user_cloudsql_instance_user" {
#   project = var.project_id
#   role    = "roles/cloudsql.instanceUser"
#   member  = "serviceAccount:${google_sql_user.admin_user.name}"
# }
