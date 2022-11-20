variable "project_id" {
  type = string
}

variable "serverless_vpc_connector_name" {
  type = string
}

variable "serverless_vpc_connector_region" {
  type = string
}

variable "serverless_vpc_connector_ip_cidr_range" {
  type = string
}

variable "symbol_store_bucket_name" {
  type = string
}

variable "symbol_store_bucket_location" {
  type = string
}

variable "firestore_location" {
  type = string
}

variable "database_instance_name" {
  type = string
}

variable "database_region" {
  type = string
}

variable "database_tier" {
  type = string
}

variable "database_disk_size_gb" {
  type = number
}

variable "database_enable_public_ip" {
  type = bool
}
