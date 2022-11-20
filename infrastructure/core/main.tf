module "google_apis" {
  source = "./google_apis"
}

module "network" {
  depends_on = [ module.google_apis ]

  source = "./network"

  serverless_vpc_connector_name = var.serverless_vpc_connector_name
  serverless_vpc_connector_region = var.serverless_vpc_connector_region
  serverless_vpc_connector_ip_cidr_range = var.serverless_vpc_connector_ip_cidr_range
}

module "symbol_store" {
    depends_on = [ module.google_apis ]

    source = "./symbol_store"

    bucket_name = var.symbol_store_bucket_name
    location = var.symbol_store_bucket_location
}

module "firestore" {
  depends_on = [ module.google_apis ]

  source = "./firestore"

  location = var.firestore_location

}

module "database_instance" {
  depends_on = [ module.google_apis, module.network ]

  source = "./database_instance"

  name = var.database_instance_name
  region = var.database_region
  tier = var.database_tier
  disk_size_gb = var.database_disk_size_gb
  enable_public_ip = var.database_enable_public_ip
  private_network_id = module.network.private_network_id
}
