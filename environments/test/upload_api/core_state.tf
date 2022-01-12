data "terraform_remote_state" "core" {
  backend = "gcs"
  config = {
    bucket = "test-cloud-symbol-server-state"
    prefix = "core"
  }
}
