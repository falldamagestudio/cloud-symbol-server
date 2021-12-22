module "core" {

    source = "../../../infrastructure/core"

    project_id = var.project_id

    symbol_store_bucket_name = var.symbol_store_bucket_name
    symbol_store_bucket_location = var.symbol_store_bucket_location
}
