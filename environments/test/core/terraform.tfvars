project_id = "test-cloud-symbol-server"

serverless_vpc_connector_name = "vpc-access-connector"
serverless_vpc_connector_region = "europe-west1"
serverless_vpc_connector_ip_cidr_range = "10.8.0.0/28"

symbol_store_bucket_name = "test-cloud-symbol-server-symbols"
symbol_store_bucket_location = "europe-west1"

# Reference: https://cloud.google.com/firestore/docs/locations
#  europe-west is a multi-region location, comprised of europe-west1 + europe-west4, with automatic replication
firestore_location = "europe-west"

database_name = "db"
database_region = "europe-west1"
database_tier = "db-g1-small"
database_disk_size_gb = 10
database_enable_public_ip = false
