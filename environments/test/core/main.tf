module "core" {

    source = "../../../infrastructure/core"

    project_id = var.project_id

    serverless_vpc_connector_name = var.serverless_vpc_connector_name
    serverless_vpc_connector_region = var.serverless_vpc_connector_region
    serverless_vpc_connector_ip_cidr_range = var.serverless_vpc_connector_ip_cidr_range

    symbol_store_bucket_name = var.symbol_store_bucket_name
    symbol_store_bucket_location = var.symbol_store_bucket_location

    firestore_location = var.firestore_location

    database_name = var.database_name
    database_region = var.database_region
    database_tier = var.database_tier
    database_disk_size_gb = var.database_disk_size_gb
    database_enable_public_ip = var.database_enable_public_ip
}
