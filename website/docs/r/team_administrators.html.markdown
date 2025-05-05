*---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_team_administrators"
description: |-
  Manages administrators of a team within a project in a Azure DevOps organization.
---


# azuredevops_team_administrators

Manages administrators of a team within a project in a Azure DevOps organization.

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  work_item_template = "Agile"
  version_control    = "Git"
  visibility         = "private"
  description        = "Managed by Terraform"
}

data "azuredevops_group" "example-project-contributors" {
  project_id = azuredevops_project.example.id
  name       = "Contributors"
}

resource "azuredevops_team" "example" {
  project_id = azuredevops_project.example.id
  name       = "${azuredevops_project.example.name} Team 2"
}

resource "azuredevops_team_administrators" "example-team-administrators" {
  project_id = azuredevops_team.example.project_id
  team_id    = azuredevops_team.example.id
  mode       = "overwrite"
  administrators = [
    data.azuredevops_group.example-project-contributors.descriptor
  ]
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The Project ID.

* `team_id` - (Required) The ID of the Team.

* `administrators` - (Required) List of subject descriptors to define administrators of the team.

  ~> **NOTE** It's possible to define team administrators both within the
   `azuredevops_team` resource via the `administrators` block and by using the
   `azuredevops_team_administrators` resource. However it's not possible to use
   both methods to manage team administrators, since there'll be conflicts.

---

* `mode` - (Optional) The mode how the resource manages team administrators. Possible values: `add`, `overwrite`. Defaults to `add`.
  
  ~> **NOTE:** 1. `mode = add`: the resource will ensure that all specified administrators will be part of the referenced team
    <br> 2. `mode = overwrite`: the resource will replace all existing administrators with the administrators specified within the `administrators` block

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A random ID for this resource. There is no "natural" ID, so a random one is assigned.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Teams - Update](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/update?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the Team administrators.
* `read` - (Defaults to 5 minute) Used when retrieving the Team administrators.
* `update` - (Defaults to 10 minutes) Used when updating the Team administrators.
* `delete` - (Defaults to 10 minutes) Used when deleting the Team administrators.

## Import

The resource does not support import.

## PAT Permissions Required

- **vso.project_write**:	Grants the ability to read and update projects and teams. 
