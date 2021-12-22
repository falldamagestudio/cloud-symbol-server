terraform {
  backend "gcs" {
    bucket = "test-cloud-symbol-store-state"
    prefix = "core"
  }
}
