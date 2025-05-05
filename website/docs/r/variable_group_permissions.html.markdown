---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_variable_group_permissions"
description: |-
  Manages permissions for aa Azure DevOps Variable Group
---

# azuredevops_variable_group_permissions

Manages permissions for a Variable Group


## Example Usage

```hcl
resource "azuredevops_project" "project" {
  name               = "Testing"
  description        = "Testing-description"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

resource "azuredevops_variable_group" "example" {
  project_id   = azuredevops_project.project.id
  name         = "test"
  description  = "Test Description"
  allow_access = true

  variable {
    name  = "key1"
    value = "val1"
  }
}

data "azuredevops_group" "tf-project-readers" {
  project_id = azuredevops_project.project.id
  name       = "Readers"
}

resource "azuredevops_variable_group_permissions" "permissions" {
  project_id        = azuredevops_project.project.id
  variable_group_id = azuredevops_variable_group.example.id
  principal         = data.azuredevops_group.tf-project-readers.id
  permissions = {
    "View" : "allow",
    "Administer" : "allow",
    "Use" : "allow",
  }
}
```

## Roles

The Azure DevOps UI uses roles to assign permissions for variable groups.

| Role          | Allow Permissions     |
|---------------|-----------------------|
| Reader        | View                  |
| User          | View, Use             |
| Administrator | View, Use, Administer |


## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project.

* `principal` - (Required) The **group** principal to assign the permissions.

* `permissions` - (Required) the permissions to assign. The following permissions are available.

  | Permission  | Description               |
  |-------------|---------------------------|
  | View        | View library item         |
  | Administer  | Administer library item   |
  | Create      | Create library item       |
  | ViewSecrets | View library item secrets |
  | Use         | Use library item          |
  | Owner       | Owner library item        |

* `variable_group_id` - (Required) The id of the variable group to assign the permissions.

---

* `replace` - (Optional) Replace (`true`) or merge (`false`) the permissions. Default: `true`

## Relevant Links

* [Azure DevOps Service REST API 7.1 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.1)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the Variable Group Permissions.
* `read` - (Defaults to 5 minute) Used when retrieving the Variable Group Permissions.
* `update` - (Defaults to 10 minutes) Used when updating the Variable Group Permissions.
* `delete` - (Defaults to 10 minutes) Used when deleting the Variable Group Permissions.

## Import

The resource does not support import.

## PAT Permissions Required

- **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
