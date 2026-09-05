terraform {
  required_version = ">= 1.15.0"

  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.38"
    }
  }
}

provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "kind-next-erp-dev-cluster"
}

resource "kubernetes_namespace" "next_erp_dev" {
  metadata {
    name = "next-erp-dev"
  }
}
