---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_argocd"
description: |-
  Manages a ArgoCD server endpoint within Azure DevOps organization.
---

# azuredevops_serviceendpoint_argocd
Manages a ArgoCD service endpoint within Azure DevOps. Using this service endpoint requires you to first install [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd).

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

resource "azuredevops_serviceendpoint_argocd" "example" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example ArgoCD"
  description           = "Managed by Terraform"
  url                   = "https://argocd.my.com"
  authentication_token {
    token = "0000000000000000000000000000000000000000"
  }
}
```
Alternatively a username and password may be used.

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

resource "azuredevops_serviceendpoint_argocd" "example" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example ArgoCD"
  description           = "Managed by Terraform"
  url                   = "https://argocd.my.com"
  authentication_basic {
    username = "username"
    password = "password"
  }
}
```
## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project.

* `service_endpoint_name` - (Required) The Service Endpoint name.

* `url` - (Required) URL of the ArgoCD server to connect with.

---

* `description` - (Optional) The Service Endpoint description.

* `authentication_token` - (Optional) An `authentication_token` block for the ArgoCD as documented below.

* `authentication_basic` - (Optional) An `authentication_basic` block for the ArgoCD as documented below.

~> **NOTE:** `authentication_basic` and `authentication_token` conflict with each other, only one is required.

---

A `authentication_token` block supports the following:

* `token` - (Required)  Authentication Token generated through ArgoCD.

---

A `authentication_basic` block supports the following:

* `username` - (Required) The Username of the ArgoCD.

* `password` - (Required) The Password of the ArgoCD.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the service endpoint.
* `project_id` - The ID of the project.
* `service_endpoint_name` - The Service Endpoint name.

## Relevant Links
- [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
- [ArgoCD Project/User Token](https://argo-cd.readthedocs.io/en/stable/user-guide/commands/argocd_account_generate-token/)
- [Argo CD Extension](https://marketplace.visualstudio.com/items?itemName=scb-tomasmortensen.vsix-argocd)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 2 minutes) Used when creating the Argo CD Service Endpoint.
* `read` - (Defaults to 1 minute) Used when retrieving the Argo CD Service Endpoint.
* `update` - (Defaults to 2 minutes) Used when updating the Argo CD Service Endpoint.
* `delete` - (Defaults to 2 minutes) Used when deleting the Argo CD Service Endpoint.

## Import
Azure DevOps Argo CD Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.


```sh
terraform import azuredevops_serviceendpoint_argocd.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
```
