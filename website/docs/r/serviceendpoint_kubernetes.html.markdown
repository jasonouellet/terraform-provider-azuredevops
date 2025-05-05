---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_kubernetes"
description: |-
  Manages a Kubernetes service endpoint Azure DevOps organization.
---

# azuredevops_serviceendpoint_kubernetes

Manages a Kubernetes service endpoint within Azure DevOps.

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

resource "azuredevops_serviceendpoint_kubernetes" "example-azure" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example Kubernetes"
  apiserver_url         = "https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io"
  authorization_type    = "AzureSubscription"

  azure_subscription {
    subscription_id   = "00000000-0000-0000-0000-000000000000"
    subscription_name = "Example"
    tenant_id         = "00000000-0000-0000-0000-000000000000"
    resourcegroup_id  = "example-rg"
    namespace         = "default"
    cluster_name      = "example-aks"
  }
}

resource "azuredevops_serviceendpoint_kubernetes" "example-kubeconfig" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example Kubernetes"
  apiserver_url         = "https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io"
  authorization_type    = "Kubeconfig"

  kubeconfig {
    kube_config            = <<EOT
                              apiVersion: v1
                              clusters:
                              - cluster:
                                  certificate-authority: fake-ca-file
                                  server: https://1.2.3.4
                                name: development
                              contexts:
                              - context:
                                  cluster: development
                                  namespace: frontend
                                  user: developer
                                name: dev-frontend
                              current-context: dev-frontend
                              kind: Config
                              preferences: {}
                              users:
                              - name: developer
                                user:
                                  client-certificate: fake-cert-file
                                  client-key: fake-key-file
                             EOT
    accept_untrusted_certs = true
    cluster_context        = "dev-frontend"
  }
}

resource "azuredevops_serviceendpoint_kubernetes" "example-service-account" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example Kubernetes"
  apiserver_url         = "https://sample-kubernetes-cluster.hcp.westeurope.azmk8s.io"
  authorization_type    = "ServiceAccount"

  service_account {
    token   = "000000000000000000000000"
    ca_cert = "0000000000000000000000000000000"
  }
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project.

* `service_endpoint_name` - (Required) The Service Endpoint name.

* `apiserver_url` - (Required) The hostname (in form of URI) of the Kubernetes API.

* `authorization_type` - (Required) The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.

---

* `azure_subscription` - (Optional) An `azure_subscription` block as defined below.

* `kubeconfig` - (Optional) A `kubeconfig` block as defined below.

* `service_account` - (Optional)  A `service_account` block as defined below.

---

An `azure_subscription` block supports the following:

The configuration for `authorization_type=AzureSubscription`.

* `azure_environment` - (Optional) Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.

* `cluster_name` - (Required) The name of the Kubernetes cluster.

* `subscription_id` - (Required) The id of the Azure subscription.

* `subscription_name` - (Required) The name of the Azure subscription.

* `tenant_id` - (Required) The id of the tenant used by the subscription.

* `resourcegroup_id` - (Required) The resource group name, to which the Kubernetes cluster is deployed.

* `namespace` - (Optional) The Kubernetes namespace. Default value is "default".

* `cluster_admin` - (Optional) Set this option to allow use cluster admin credentials.

---

A `kubeconfig` block supports the following: 

The configuration for `authorization_type=Kubeconfig`. 

* `kube_config` - (Required) The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.

* `accept_untrusted_certs` - (Optional) Set this option to allow clients to accept a self-signed certificate.

* `cluster_context` - (Optional) Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.

---

A `service_account` block supports the following:  

The configuration for `authorization_type=ServiceAccount`. This type uses the credentials of a service account currently deployed to the cluster.

* `token` - (Required) The token from a Kubernetes secret object.

* `ca_cert` - (Required) The certificate from a Kubernetes secret object.

* `accept_untrusted_certs` - (Optional) Set this option to allow clients to accept a self-signed certificate. Defaults to `false`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the service endpoint.
* `project_id` - The ID of the project.
* `service_endpoint_name` - The Service Endpoint name.

## Relevant Links

- [Azure DevOps Service REST API 7.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-7.0)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 2 minutes) Used when creating the Kubernetes Service Endpoint.
* `read` - (Defaults to 1 minute) Used when retrieving the Kubernetes Service Endpoint.
* `update` - (Defaults to 2 minutes) Used when updating the Kubernetes Service Endpoint.
* `delete` - (Defaults to 2 minutes) Used when deleting the Kubernetes Service Endpoint.

## Import

Azure DevOps Kubernetes Service Endpoint can be imported using **projectID/serviceEndpointID** or **projectName/serviceEndpointID**

```sh
terraform import azuredevops_serviceendpoint_kubernetes.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
```
