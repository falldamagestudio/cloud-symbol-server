output "db_terraform_admin_user_name" {
    value = module.database_instance.admin_user_name
}

output "db_terraform_admin_user_password" {
    value = module.database_instance.admin_user_password
    sensitive = true
}
