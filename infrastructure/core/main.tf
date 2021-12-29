module "google_apis" {
  source = "./google_apis"
}

module "symbol_store" {
    depends_on = [ module.google_apis ]

    source = "./symbol_store"

    bucket_name = var.symbol_store_bucket_name
    location = var.symbol_store_bucket_location
}

module "firestore" {
  depends_on = [ module.google_apis ]

  source = "./firestore"

  location = var.firestore_location

}