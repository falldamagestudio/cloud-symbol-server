variable "project_id" {
  type = string
}

variable "source_path" {
  type = string
}

variable "source_bucket_name" {
  type = string
}

variable "source_bucket_location" {
  type = string
}

variable "function_region" {
  type = string
}

variable "symbol_store_bucket_name" {
  type = string
}

variable "symbol_server_stores" {
  type = list(string)
}
