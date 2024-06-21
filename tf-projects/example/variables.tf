#variable "name" {
#  type    = string
#  default = "neuvector"
#}

#variable "namespace" {
#  type    = string
#  default = "cattle-neuvector-system"
#}

#variable "app_version" {
#  type    = string
#  default = ""
#}

#variable "controller_replicas" {
#  type = number
#  default = 3
#}

#variable "controller_env" {
#  type = list(object({
#    name  = string
#    value = string
#  }))
#  default = []
#}

#variable "controller_node_selector" {
#  type = map(any)
#  default = {}
#}

variable "controller_secret_enabled" {
  type = bool
  #  default = false
}

variable "controller_secret_password" {
  type = string
  #  default = ""
}

variable "manager_svc_type" {
  type = string
  #  default = "ClusterIP"
}

#variable "cve_scanner_replicas" {
#  type    = number
#  default = 2
#}

#variable "cve_scanner_node_selector" {
#  type = map(any)
#  default = {}
#}

#variable "resources_limits_cpu" {
#  type    = string
#  default = "400m"
#}

#variable "resources_limits_memory" {
#  type    = string
#  default = "2792Mi"
#}

#variable "resources_requests_cpu" {
#  type    = string
#  default = "100m"
#}

#variable "resources_requests_memory" {
#  type    = string
#  default = "2280Mi"
#}

#variable "containerd_enabled" {
#  type    = bool
#  default = true
#}

#variable "containerd_path" {
#  type    = string
#  default = "/var/run/containerd/containerd.sock"
#}

variable "kubeconfig_path" {
  type = string
}
