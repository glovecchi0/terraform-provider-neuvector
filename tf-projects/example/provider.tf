terraform {
  required_providers {
    neuvector = {
      source  = "local/neuvector"
      version = "0.1.0"
    }
  }
}

provider "neuvector" {}
