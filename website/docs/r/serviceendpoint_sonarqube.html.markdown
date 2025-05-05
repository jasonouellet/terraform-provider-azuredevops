---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_sonarqube"
description: |-
  Manages a SonarQube Server service endpoint within Azure DevOps organization.
---

# azuredevops_serviceendpoint_sonarqube
Manages a SonarQube Server service endpoint within Azure DevOps. 

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

resource "azuredevops_serviceendpoint_sonarqube" "example" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example SonarQube"
  url                   = "https://sonarqube.my.com"
  token                 = "0000000000000000000000000000000000000000"
  description           = "Managed by Terraform"
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project.

* `service_endpoint_name` - (Required) The Service Endpoint name.

* `url` - (Required) URL of the SonarQube server to connect with.

* `token` - (Required) The Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).

---

* `description` - (Optional) The Service Endpoint description.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the service endpoint.
* `project_id` - The ID of the project.
* `service_endpoint_name` - The Service Endpoint name.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)
- [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
- [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 2 minutes) Used when creating the SonarQube Server Service Endpoint.
* `read` - (Defaults to 1 minute) Used when retrieving the SonarQube Server Service Endpoint.
* `update` - (Defaults to 2 minutes) Used when updating the SonarQube Server Service Endpoint.
* `delete` - (Defaults to 2 minutes) Used when deleting the SonarQube Server Service Endpoint.

## Import
Azure DevOps SonarQube Server Service Endpoint can be imported using the **projectID/serviceEndpointID**, e.g.

```sh
terraform import azuredevops_serviceendpoint_sonarqube.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
```
