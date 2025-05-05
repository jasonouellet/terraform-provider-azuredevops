---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_servicefabric"
description: |-
  Manages a Service Fabric service endpoint Azure DevOps organization.
---

# azuredevops_serviceendpoint_servicefabric

Manages a Service Fabric service endpoint within Azure DevOps.

## Example Usage

### Client Certificate Authentication

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

resource "azuredevops_serviceendpoint_servicefabric" "example" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example Service Fabric"
  description           = "Managed by Terraform"
  cluster_endpoint      = "tcp://test"

  certificate {
    server_certificate_lookup     = "Thumbprint"
    server_certificate_thumbprint = "0000000000000000000000000000000000000000"
    client_certificate            = filebase64("certificate.pfx")
    client_certificate_password   = "password"
  }
}
```

### Azure Active Directory Authentication

```hcl
resource "azuredevops_project" "project" {
  name               = "Sample Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

resource "azuredevops_serviceendpoint_servicefabric" "test" {
  project_id            = azuredevops_project.project.id
  service_endpoint_name = "Sample Service Fabric"
  description           = "Managed by Terraform"
  cluster_endpoint      = "tcp://test"

  azure_active_directory {
    server_certificate_lookup     = "Thumbprint"
    server_certificate_thumbprint = "0000000000000000000000000000000000000000"
    username                      = "username"
    password                      = "password"
  }
}
```

### Windows Authentication

```hcl
resource "azuredevops_project" "project" {
  name               = "Sample Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

resource "azuredevops_serviceendpoint_servicefabric" "test" {
  project_id            = azuredevops_project.project.id
  service_endpoint_name = "Sample Service Fabric"
  description           = "Managed by Terraform"
  cluster_endpoint      = "tcp://test"

  none {
    unsecured   = false
    cluster_spn = "HTTP/www.contoso.com"
  }
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project.

* `service_endpoint_name` - (Required) The Service Endpoint name.

* `cluster_endpoint` - (Required) Client connection endpoint for the cluster. Prefix the value with `tcp://`;. This value overrides the publish profile.

---

* `azure_active_directory` -(Optional) An `azure_active_directory` block as documented below.

* `certificate` -(Optional) A `certificate` block as documented below.

* `description` - (Optional) The Service Endpoint description. Defaults to `Managed by Terraform`.

* `none` -(Optional) A `none` block as documented below.

---

A `certificate` block supports the following:

* `server_certificate_lookup` - (Required) Verification mode for the cluster. Possible values are: `Thumbprint`, `CommonName`.

* `client_certificate` - (Required) Base64 encoding of the cluster's client certificate file.

* `server_certificate_thumbprint` - (Optional) The thumbprint(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (',')

* `server_certificate_common_name` - (Optional) The common name(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (',')

* `client_certificate_password` - (Optional) Password for the certificate.

---

An `azure_active_directory` block supports the following:

* `server_certificate_lookup` - (Required) Verification mode for the cluster. Possible values are: `Thumbprint`, `CommonName`.

* `username` - (Required) - Specify an Azure Active Directory account.

* `password` - (Required) - Password for the Azure Active Directory account.

* `server_certificate_thumbprint` - (Optional) The thumbprint(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple thumbprints with a comma (',')

* `server_certificate_common_name` - (Optional) The common name(s) of the cluster's certificate(s). This is used to verify the identity of the cluster. This value overrides the publish profile. Separate multiple common names with a comma (',')

---

A `none` block supports the following:

* `unsecured` - (Optional) Skip using windows security for authentication.

* `cluster_spn` - (Optional) Fully qualified domain SPN for gMSA account. This is applicable only if `unsecured` option is disabled.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the service endpoint.
* `project_id` - The ID of the project.
* `service_endpoint_name` - The Service Endpoint name.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 2 minutes) Used when creating the Service Fabric Service Endpoint.
* `read` - (Defaults to 1 minute) Used when retrieving the Service Fabric Service Endpoint.
* `update` - (Defaults to 2 minutes) Used when updating the Service Fabric Service Endpoint.
* `delete` - (Defaults to 2 minutes) Used when deleting the Service Fabric Service Endpoint.

## Import

Azure DevOps Service Fabric Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

```sh
terraform import azuredevops_serviceendpoint_servicefabric.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
```
