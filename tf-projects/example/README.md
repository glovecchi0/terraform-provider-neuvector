# How to create resources

- Copy `./terraform.tfvars.example` to `./terraform.tfvars`
- Edit `./terraform.tfvars`
  - Update the required variables:
    -  `kubeconfig_path` It's the only variable really necessary; It allows you to release the NeuVector Helm Chart on a specific K8s cluster

**NB: IF YOU WANT TO USE ALL THE CONFIGURABLE VARIABLES IN THE `terraform.tfvars` FILE, YOU WILL NEED TO UNCOMMENT THEM THERE AND IN THE `variables.tf` AND `main.tf` FILES.**

## Sample Terraform files for deployment on a GKE cluster (by configuring a custom password for the NeuVector administration user)

#### main.tf

```console
resource "neuvector_application" "example" {
  dynamic "controller_env" {
    for_each = var.controller_env
    content {
      name  = controller_env.value.name
      value = controller_env.value.value
    }
  }

  controller_secret_enabled  = var.controller_secret_enabled
  controller_secret_password = var.controller_secret_password
  manager_svc_type           = var.manager_svc_type
  kubeconfig_path            = var.kubeconfig_path
}
```

#### variables.tf

```console
#variable "name" {}

#variable "namespace" {}

#variable "app_version" {}

#variable "controller_replicas" {}

variable "controller_env" {
  type = list(object({
    name  = string
    value = string
  }))
  #  default = []
}

#variable "controller_node_selector" {}

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

#variable "cve_scanner_replicas" {}

#variable "cve_scanner_node_selector" {}

#variable "resources_limits_cpu" {}

#variable "resources_limits_memory" {}

#variable "resources_requests_cpu" {}

#variable "resources_requests_memory" {}

#variable "containerd_enabled" {}

#variable "containerd_path" {}

variable "kubeconfig_path" {
  type = string
}
```

#### terraform.tfvars

```console
controller_env = [
  {
    name  = "NV_PLATFORM_INFO"
    value = "platform=Kubernetes:GKE"
  }
]
controller_secret_enabled  = true
controller_secret_password = "YourPassword.123"
manager_svc_type           = "LoadBalancer"
kubeconfig_path            = "path/to/your/kubeconfig/file"
```

## Sample Terraform files for deployment on an EKS cluster (by configuring a custom password for the NeuVector administration user)

**NB: THE TERRAFORM FILES FOUND IN THIS DIRECTORY HAVE BEEN TESTED ON EKS.**

#### main.tf

```console
resource "neuvector_application" "example" {
  controller_secret_enabled  = var.controller_secret_enabled
  controller_secret_password = var.controller_secret_password
  manager_svc_type           = var.manager_svc_type
  kubeconfig_path            = var.kubeconfig_path
}
```

#### variables.tf

```console
#variable "name" {}

#variable "namespace" {}

#variable "app_version" {}

#variable "controller_replicas" {}

#variable "controller_env" {
#  type = list(object({
#    name  = string
#    value = string
#  }))
#  default = []
#}

#variable "controller_node_selector" {}

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

#variable "cve_scanner_replicas" {}

#variable "cve_scanner_node_selector" {}

#variable "resources_limits_cpu" {}

#variable "resources_limits_memory" {}

#variable "resources_requests_cpu" {}

#variable "resources_requests_memory" {}

#variable "containerd_enabled" {}

#variable "containerd_path" {}

variable "kubeconfig_path" {
  type = string
}
```

#### terraform.tfvars

```console
controller_secret_enabled  = true
controller_secret_password = "YourPassword.123"
manager_svc_type           = "LoadBalancer"
kubeconfig_path            = "path/to/your/kubeconfig/file"
```

# How to deploy resources using Terraform
```bash
terraform init --upgrade
terraform apply
```

# How to destroy resources using Terraform
```bash
terraform destroy
```

# How to deploy resources using OpenTofu
```bash
tofu init --upgrade
tofu apply
```

# How to destroy resources using OpenTofu
```bash
tofu destroy
```
