output "project_id" {
    value = var.project_id
}

output "symbol_store_bucket_name" {
    value = var.symbol_store_bucket_name
}

output "database_name" {
    value = var.database_name
}

output "serverless_vpc_connector_name" {
    value = var.serverless_vpc_connector_name
}

output "db_admin_user_passsword" {
    value = module.core.db_admin_user_passsword
    sensitive = true
}