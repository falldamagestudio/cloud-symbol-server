module "database" {

    source = "../../../infrastructure/database"

    project_id = data.terraform_remote_state.core.outputs.project_id

    database_instance_name = data.terraform_remote_state.core.outputs.database_instance_name
    database_region = data.terraform_remote_state.core.outputs.database_region

    admin_user_name = data.terraform_remote_state.core.outputs.db_terraform_admin_user_name
    admin_user_password = data.terraform_remote_state.core.outputs.db_terraform_admin_user_password
}
