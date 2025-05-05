---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_teams"
description: |-
  Use this data source to access information about existing Teams in a Project or globally within an Azure DevOps organization
---

# Data Source: azuredevops_teams

Use this data source to access information about existing Teams in a Project or globally within an Azure DevOps organization

## Example Usage

```hcl
data "azuredevops_teams" "example" {
}

output "project_id" {
  value = data.azuredevops_teams.example.teams.*.project_id
}

output "name" {
  value = data.azuredevops_teams.example.teams.*.name
}

output "all_administrators" {
  value = data.azuredevops_teams.example.teams.*.administrators
}

output "administrators" {
  value = data.azuredevops_teams.example.teams.*.members
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Optional) The Project ID. If no project ID all teams of the organization will be returned.

* `top` - (Optional) The maximum number of teams to return. Defaults to `100`.

## Attributes Reference

The following attributes are exported:

* `teams` - A list of `teams` blocks as documented below. A list of existing projects in your Azure DevOps Organization with details about every project which includes:

---

A `teams` block supports the following:

* `project_id` - The ID of the Project.

* `id` - The ID of the Team.

* `name` - The name of the team.

* `description` - Team description.

* `administrators` - List of subject descriptors for `administrators` of the team.

* `members` - List of subject descriptors for `members` of the team.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `read` - (Defaults to 30 minute) Used when retrieving the Teams.

## PAT Permissions Required

- **vso.project**:	Grants the ability to read projects and teams.
