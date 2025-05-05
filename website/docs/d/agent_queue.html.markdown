---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_agent_queue"
description: |-
  Use this data source to access information about an existing Agent Queue within Azure DevOps.
---

# Data Source: azuredevops_agent_queue

Use this data source to access information about an existing Agent Queue within Azure DevOps.

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  work_item_template = "Agile"
  version_control    = "Git"
  visibility         = "private"
  description        = "Managed by Terraform"
}
data "azuredevops_agent_queue" "example" {
  project_id = azuredevops_project.example.id
  name       = "Example Agent Queue"
}

output "name" {
  value = data.azuredevops_agent_queue.example.name
}

output "pool_id" {
  value = data.azuredevops_agent_queue.example.agent_pool_id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The Project Id.

* `name` - (Required) Name of the Agent Queue.

## Attributes Reference

The following attributes are exported:

* `id`  - The id of the agent queue.

* `name` - The name of the agent queue.

* `project_id` - The ID of the Project to which the agent queue belongs.

* `agent_pool_id` - The ID of the Agent pool to which the agent queue belongs.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Agent Queues - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/distributedtask/queues/get?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minute) Used when retrieving the Agent Queue.
