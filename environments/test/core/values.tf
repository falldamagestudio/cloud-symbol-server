output "project_id" {
    value = var.project_id
}

output "symbol_store_bucket_name" {
    value = var.symbol_store_bucket_name
}

output "database_instance_name" {
    value = var.database_instance_name
}

output "database_region" {
    value = var.database_region
}

output "serverless_vpc_connector_name" {
    value = var.serverless_vpc_connector_name
}

output "db_terraform_admin_user_name" {
    value = module.core.db_terraform_admin_user_name
}

output "db_terraform_admin_user_password" {
    value = module.core.db_terraform_admin_user_password
    sensitive = true
}