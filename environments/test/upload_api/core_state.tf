data "terraform_remote_state" "core" {
  backend = "gcs"
  config = {
    bucket = "test-cloud-symbol-store-state"
    prefix = "core"
  }
}
