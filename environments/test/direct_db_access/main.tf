module "direct_db_access" {

    source = "../../../infrastructure/direct_db_access"

    project_id = data.terraform_remote_state.core.outputs.project_id

    database_name = data.terraform_remote_state.core.outputs.database_name
}
