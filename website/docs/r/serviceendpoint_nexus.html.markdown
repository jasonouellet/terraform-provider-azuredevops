---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_nexus"
description: |-
  Manages a Service Connection for Nexus IQ.
---

# azuredevops_serviceendpoint_nexus

Manages a Nexus IQ service endpoint within Azure DevOps, which can be used as a resource in YAML pipelines to connect to a Nexus IQ instance.
Nexus IQ is not supported by default, to manage a nexus service connection resource, it is necessary to install the [Nexus Extension](https://marketplace.visualstudio.com/items?itemName=SonatypeIntegrations.nexus-iq-azure-extension) in Azure DevOps.

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

resource "azuredevops_serviceendpoint_nexus" "example" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "nexus-example"
  description           = "Service Endpoint for 'Nexus IQ' (Managed by Terraform)"
  url                   = "https://example.com"

  username = "username"
  password = "password"
}
```

## Arguments Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project. Changing this forces a new Service Connection Nexus to be created.

* `service_endpoint_name` - (Required) The name of the service endpoint. Changing this forces a new Service Connection Nexus to be created.

* `url` - (Required) The Service Endpoint url.

* `username` - (Required) The Service Endpoint username to authenticate at the Nexus IQ Instance.

* `password` - (Required) The Service Endpoint password to authenticate at the Nexus IQ Instance.

---

* `description` - (Optional) The Service Endpoint description. Defaults to Managed by Terraform.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the service endpoint.
* `project_id` - The ID of the project.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 2 minutes) Used when creating the Nexus Service Endpoint.
* `read` - (Defaults to 1 minute) Used when retrieving the Nexus Service Endpoint.
* `update` - (Defaults to 2 minutes) Used when updating the Nexus Service Endpoint.
* `delete` - (Defaults to 2 minutes) Used when deleting the Nexus Service Endpoint.

## Import

Azure DevOps Nexus Service Connection can be imported using the `projectId/id` or or `projectName/id`, e.g.

```shell
terraform import azuredevops_serviceendpoint_nexus.example projectName/00000000-0000-0000-0000-000000000000
```
