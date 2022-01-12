module "download_api" {

    source = "../../../infrastructure/admin_api"

    project_id = data.terraform_remote_state.core.outputs.project_id
    # database_location = data.terraform_remote_state.core.outputs.database_location
    source_path = "../../../admin-api"
    source_bucket_name = var.source_bucket_name
    source_bucket_location = var.source_bucket_location
    function_region = var.function_region

    symbol_store_bucket_name = data.terraform_remote_state.core.outputs.symbol_store_bucket_name
}
