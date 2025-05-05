---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_build_folder_permissions"
description: |-
  Manages permissions for a AzureDevOps Build Folder
---

# azuredevops_build_folder_permissions

Manages permissions for a Build Folder

~> **Note** Permissions can be assigned to group principals and not to single user principals.

## Example Usage
### Set specific folder permissions

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  work_item_template = "Agile"
  version_control    = "Git"
  visibility         = "private"
  description        = "Managed by Terraform"
}

data "azuredevops_group" "example-readers" {
  project_id = azuredevops_project.example.id
  name       = "Readers"
}

resource "azuredevops_build_folder" "example" {
  project_id  = azuredevops_project.example.id
  path        = "\\ExampleFolder"
  description = "ExampleFolder description"
}

resource "azuredevops_build_folder_permissions" "example" {
  project_id = azuredevops_project.example.id
  path       = "\\ExampleFolder"
  principal  = data.azuredevops_group.example-readers.id

  permissions = {
    "ViewBuilds" : "Allow",
    "EditBuildQuality" : "Allow",
    "RetainIndefinitely" : "Allow",
    "DeleteBuilds" : "Deny",
    "ManageBuildQualities" : "Deny",
    "DestroyBuilds" : "Deny",
    "UpdateBuildInformation" : "Deny",
    "QueueBuilds" : "Allow",
    "ManageBuildQueue" : "Deny",
    "StopBuilds" : "Allow",
    "ViewBuildDefinition" : "Allow",
    "EditBuildDefinition" : "Deny",
    "DeleteBuildDefinition" : "Deny",
    "AdministerBuildPermissions" : "NotSet"
  }
}
```
### Set root folder permissions
```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  work_item_template = "Agile"
  version_control    = "Git"
  visibility         = "private"
  description        = "Managed by Terraform"
}

data "azuredevops_group" "example-readers" {
  project_id = azuredevops_project.example.id
  name       = "Readers"
}

resource "azuredevops_build_folder_permissions" "example" {
  project_id = azuredevops_project.example.id
  path       = "\\"
  principal  = data.azuredevops_group.example-readers.id

  permissions = {
    "RetainIndefinitely" : "Allow"
  }
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project to assign the permissions.

* `principal` - (Required) The **group** principal to assign the permissions.

* `path` - (Required) The folder path to assign the permissions.

* `permissions` - (Required) the permissions to assign. The following permissions are available.

    | Permission                     | Description                           |
    |--------------------------------|---------------------------------------|
    | ViewBuilds                     | View builds                           |
    | EditBuildQuality               | Edit build quality                    |
    | RetainIndefinitely             | Retain indefinitely                   |
    | DeleteBuilds                   | Delete builds                         |
    | ManageBuildQualities           | Manage build qualities                |
    | DestroyBuilds                  | Destroy builds                        |
    | UpdateBuildInformation         | Update build information              |
    | QueueBuilds                    | Queue builds                          |
    | ManageBuildQueue               | Manage build queue                    |
    | StopBuilds                     | Stop builds                           |
    | ViewBuildDefinition            | View build pipeline                   |
    | EditBuildDefinition            | Edit build pipeline                   |
    | DeleteBuildDefinition          | Delete build pipeline                 |
    | OverrideBuildCheckInValidation | Override check-in validation by build |
    | AdministerBuildPermissions     | Administer build permissions          |
    | CreateBuildDefinition          | Create build pipeline                 |

---

* `replace` - (Optional) Replace (`true`) or merge (`false`) the permissions. Default: `true`.

## Relevant Links

* [Azure DevOps Service REST API 7.0 - Security](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the Build Folder Permission.
* `read` - (Defaults to 5 minute) Used when retrieving the Build Folder Permission.
* `update` - (Defaults to 10 minutes) Used when updating the Build Folder Permission.
* `delete` - (Defaults to 10 minutes) Used when deleting the Build Folder Permission.

## Import

The resource does not support import.

## PAT Permissions Required

- **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
