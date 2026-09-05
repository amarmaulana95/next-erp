output "namespace_name" {
  description = "Nama namespace yang dikelola Terraform"
  value       = kubernetes_namespace.next_erp_dev.metadata[0].name
}

output "namespace_uid" {
  description = "UID namespace Kubernetes"
  value       = kubernetes_namespace.next_erp_dev.metadata[0].uid
}
