BEGIN;

-- The schema_migrations table has been created by the migration tool,
--  and the migration tool gives it the current login user as owner
-- We change this to the group role to ensure consistency with how we handle
--  ownership for other objects in the database
ALTER TABLE cloud_symbol_server.schema_migrations OWNER TO cloud_symbol_server_admin;

END;