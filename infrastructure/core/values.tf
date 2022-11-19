output "db_admin_user_passsword" {
    value = module.database.db_admin_user_passsword
    sensitive = true
}
