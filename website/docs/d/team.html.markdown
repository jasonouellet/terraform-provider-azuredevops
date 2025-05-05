---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_team"
description: |-
  Use this data source to access information about an existing Team in a Project within Azure DevOps.
---

# Data Source: azuredevops_team

Use this data source to access information about an existing Team in a Project within Azure DevOps.

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  work_item_template = "Agile"
  version_control    = "Git"
  visibility         = "private"
  description        = "Managed by Terraform"
}

data "azuredevops_team" "example" {
  project_id = azuredevops_project.example.id
  name       = "Example Project Team"
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The Project ID.

* `name` - (Required) The name of the Team.

---

* `top` - (Optional) The maximum number of teams to return. Defaults to `100`. This property is deprecated and will be removed in the feature

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Team.

* `descriptor` - The descriptor of the Team.

* `description` - The description of the team.

* `administrators` - List of subject descriptors for `administrators` of the team.

* `members` - List of subject descriptors for `members` of the team.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Teams - Get](https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minute) Used when retrieving the Team.

## PAT Permissions Required

- **vso.project**:	Grants the ability to read projects and teams.
