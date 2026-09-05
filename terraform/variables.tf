variable "kubernetes_context" {
  description = "Kubernetes context yang digunakan Terraform"
  type        = string
  default     = "kind-next-erp-dev-cluster"
}

variable "namespace" {
  description = "Namespace aplikasi Next ERP"
  type        = string
  default     = "next-erp-dev"
}
