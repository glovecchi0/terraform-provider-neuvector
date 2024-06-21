resource "neuvector_application" "example" {
  controller_secret_enabled  = var.controller_secret_enabled
  controller_secret_password = var.controller_secret_password
  manager_svc_type           = var.manager_svc_type
  kubeconfig_path            = var.kubeconfig_path
}
