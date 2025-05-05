---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_users"
description: |-
  Use this data source to access information about an existing users within Azure DevOps.
---

# Data Source: azuredevops_users

Use this data source to access information about an existing users within Azure DevOps.

~> **NOTE:** This resource will try to get all the users within the organization which may result in poor performance. `azuredevops_identity_user`, `azuredevops_user` can be used to replace this resource.


## Example Usage

```hcl
# Load single user by using it's principal name
data "azuredevops_users" "example" {
  principal_name = "contoso-user@contoso.onmicrosoft.com"
}

# Load all users know inside an organization
data "azuredevops_users" "example-all-users" {
}

# Load all users know inside an organization with concurrent processing
data "azuredevops_users" "example-all-users" {
  features {
    concurrent_workers = 10
  }
}

# Load all users know inside an organization originating from a specific source (origin)
data "azuredevops_users" "example-all-from-origin" {
  origin = "aad"
}

# Load all users know inside an organization filtered by their subject types
data "azuredevops_users" "example-all-from-subject_types" {
  subject_types = ["aad", "msa"]
}

# Load a single user by origin and origin ID
data "azuredevops_users" "example-all-from-origin-id" {
  origin    = "aad"
  origin_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

~> **NOTE:** DataSource without specifying any arguments will return all users inside an organization.


* `principal_name` - (Optional) The PrincipalName of this graph member from the source provider.

* `subject_types` - (Optional) A list of user subject subtypes to reduce the retrieved results, e.g. `msa`, `aad`, `svc` (service identity), `imp` (imported identity), etc. The supported subject types are listed below.
  <pre>List of possible subject types
  ```hcl
  AadUser                 = "aad" # Azure Active Directory Tenant
  MsaUser                 = "msa" # Windows Live
  UnknownUser             = "unusr"
  BindPendingUser         = "bnd" # Invited user with pending redeem status
  WindowsIdentity         = "win" # Windows Active Directory user
  UnauthenticatedIdentity = "uauth"
  ServiceIdentity         = "svc"
  AggregateIdentity       = "agg"
  ImportedIdentity        = "imp"
  ServerTestIdentity      = "tst"
  GroupScopeType          = "scp"
  CspPartnerIdentity      = "csp"
  SystemServicePrincipal  = "s2s"
  SystemLicense           = "slic"
  SystemScope             = "sscp"
  SystemCspPartner        = "scsp"
  SystemPublicAccess      = "spa"
  SystemAccessControl     = "sace"
  AcsServiceIdentity      = "acs"
  Unknown                 = "ukn"
  ```
  </pre>

* `origin` - (Optional) The type of source provider for the `origin_id` parameter (ex:AD, AAD, MSA) The supported origins are listed below.
  <pre>List of possible origins
  ```hcl
  ActiveDirectory          = "ad"   # Windows Active Directory
  AzureActiveDirectory     = "aad"  # Azure Active Directory
  MicrosoftAccount         = "msa"  # Windows Live Account
  VisualStudioTeamServices = "vsts" # DevOps
  GitHubDirectory          = "ghb"  # GitHub
  ```
  </pre>
  

* `origin_id` - (Optional) The unique identifier from the system of origin.

* `features` - (Optional) A `features` block as defined below.

---

A `features` block supports the following:

* `concurrent_workers` - (Optional) Number of workers to process user data concurrently.

-> **Note** Setting `concurrent_workers` to a value greater than 1 can greatly decrease the time it takes to read the data source.


## Attributes Reference

The following attributes are exported:

* `users` - A `users` block as defined below. A set of existing users in your Azure DevOps Organization with details about every single user.


---

A `users` block supports the following:

* `id` - The ID of the User.

* `descriptor` - The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.

* `principal_name` - This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.

* `origin` - The type of source provider for the origin identifier (ex:AD, AAD, MSA)

* `origin_id` - The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.

* `display_name` - This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.

* `mail_address` - The email address of record for a given graph member. This may be different than the principal name.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Graph Users API](https://docs.microsoft.com/en-us/rest/api/azure/devops/graph/users?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `read` - (Defaults to 30 minute) Used when retrieving the Users.
