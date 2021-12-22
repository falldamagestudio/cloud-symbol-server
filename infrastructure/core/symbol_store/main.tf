resource "google_storage_bucket" "store" {
  name          = var.bucket_name
  location      = var.location
  force_destroy = true

  uniform_bucket_level_access = true
}

# Allow anyone to read from symbol store
resource "google_storage_bucket_iam_member" "store_read_access" {

  depends_on = [ google_storage_bucket.store ]

  bucket   = var.bucket_name
  role     = "roles/storage.objectViewer"
  member = "allUsers"
}

