output "db_admin_user_passsword" {
    value = random_id.admin_password.hex
    sensitive = true
}
