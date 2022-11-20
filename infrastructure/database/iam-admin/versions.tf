terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0.0"
    }

    postgresql = {
      source = "cyrilgdn/postgresql"
      version = "~> 1.17.0"
    }
  }
}
