project_id = "test-cloud-symbol-server"

symbol_store_bucket_name = "test-cloud-symbol-server-symbols"
symbol_store_bucket_location = "europe-west1"

symbol_server_stores = [
    "example"
]

# Reference: https://cloud.google.com/firestore/docs/locations
#  europe-west is a multi-region location, comprised of europe-west1 + europe-west4, with automatic replication
firestore_location = "europe-west"

