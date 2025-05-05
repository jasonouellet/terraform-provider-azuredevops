---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_azurecr"
description: |-
  Manages a Azure Container Registry service endpoint within Azure DevOps organization.
---

# azuredevops_serviceendpoint_azurecr

Manages a Azure Container Registry service endpoint within Azure DevOps.

## Example Usage

### Authorize with Service Principal
```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

# azure container registry service connection
resource "azuredevops_serviceendpoint_azurecr" "example" {
  project_id                = azuredevops_project.example.id
  service_endpoint_name     = "Example AzureCR"
  resource_group            = "example-rg"
  azurecr_spn_tenantid      = "00000000-0000-0000-0000-000000000000"
  azurecr_name              = "ExampleAcr"
  azurecr_subscription_id   = "00000000-0000-0000-0000-000000000000"
  azurecr_subscription_name = "subscription name"
}
```

### Authorize with WorkloadIdentityFederation

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

resource "azurerm_resource_group" "identity" {
  name     = "identity"
  location = "UK South"
}

resource "azurerm_user_assigned_identity" "example" {
  location            = azurerm_resource_group.identity.location
  name                = "example-identity"
  resource_group_name = azurerm_resource_group.identity.name
}

# azure container registry service connection
resource "azuredevops_serviceendpoint_azurecr" "example" {
  project_id                             = azuredevops_project.example.id
  resource_group                         = "Example AzureCR ResourceGroup"
  service_endpoint_name                  = "Example AzureCR"
  service_endpoint_authentication_scheme = "WorkloadIdentityFederation"
  azurecr_spn_tenantid                   = "00000000-0000-0000-0000-000000000000"
  azurecr_name                           = "ExampleAcr"
  azurecr_subscription_id                = "00000000-0000-0000-0000-000000000000"
  azurecr_subscription_name              = "subscription name"
  credentials {
    serviceprincipalid = azurerm_user_assigned_identity.example.client_id
  }
}

resource "azurerm_federated_identity_credential" "example" {
  name                = "example-federated-credential"
  resource_group_name = azurerm_resource_group.identity.name
  parent_id           = azurerm_user_assigned_identity.example.id
  audience            = ["api://AzureADTokenExchange"]
  issuer              = azuredevops_serviceendpoint_azurecr.example.workload_identity_federation_issuer
  subject             = azuredevops_serviceendpoint_azurecr.example.workload_identity_federation_subject
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project.

* `service_endpoint_name` - (Required) The name you will use to refer to this service connection in task inputs.

* `resource_group` - (Required) The resource group to which the container registry belongs.

* `azurecr_spn_tenantid` - (Required) The tenant id of the service principal.

* `service_endpoint_authentication_scheme` - (Optional) Specifies the type of azurerm endpoint, either `WorkloadIdentityFederation`, `ManagedServiceIdentity` or `ServicePrincipal`. Defaults to `ServicePrincipal` for backwards compatibility. `ManagedServiceIdentity` has not yet been implemented for this resource.

* `azurecr_name` - (Required) The Azure container registry name.

* `azurecr_subscription_id` - (Required) The subscription id of the Azure targets.

* `azurecr_subscription_name` - (Required) The subscription name of the Azure targets.

---

* `description` - (Optional) The Service Endpoint description. Defaults to `Managed by Terraform`.

* `credentials` - (Optional) A `credentials` block as defined below.

---

A `credentials` block supports the following:

* `serviceprincipalid` - (Required) The ID of the Service Principal Application.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the service endpoint.
* `project_id` - The ID of the project.
* `service_endpoint_name` - The Service Endpoint name.
* `service_principal_id` - The Application(Client) ID of the Service Principal.
* `workload_identity_federation_issuer` - The issuer if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `https://vstoken.dev.azure.com/00000000-0000-0000-0000-000000000000`, where the GUID is the Organization ID of your Azure DevOps Organisation.
* `workload_identity_federation_subject` - The subject if `service_endpoint_authentication_scheme` is set to `WorkloadIdentityFederation`. This looks like `sc://<organisation>/<project>/<service-connection-name>`.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
- [Azure Container Registry REST API](https://docs.microsoft.com/en-us/rest/api/containerregistry/)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 2 minutes) Used when creating the Azure Container Registry Service Endpoint.
* `read` - (Defaults to 1 minute) Used when retrieving the Azure Container Registry Service Endpoint.
* `update` - (Defaults to 2 minutes) Used when updating the Azure Container Registry Service Endpoint.
* `delete` - (Defaults to 2 minutes) Used when deleting the Azure Container Registry Service Endpoint.

## Import

Azure DevOps Azure Container Registry Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

```sh
terraform import azuredevops_serviceendpoint_azurecr.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
```
