terraform {
  backend "gcs" {
    bucket = "test-cloud-symbol-server-state"
    prefix = "database"
  }
}
