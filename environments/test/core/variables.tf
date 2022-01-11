variable "project_id" {
  type = string
}

variable "symbol_store_bucket_name" {
  type = string
}

variable "symbol_store_bucket_location" {
  type = string
}

variable "symbol_store_local_stores" {
  type = list(string)
}

variable "firestore_location" {
  type = string
}
