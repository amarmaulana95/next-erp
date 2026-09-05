resource "kubernetes_namespace" "next_erp_dev" {
  metadata {
    name = var.namespace

    labels = {
      managed-by = "terraform"
    }
  }
}
