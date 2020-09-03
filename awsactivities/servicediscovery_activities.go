package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
)

type ServiceDiscoveryActivities struct {
	client servicediscoveryiface.ServiceDiscoveryAPI
}

func NewServiceDiscoveryActivities(client servicediscoveryiface.ServiceDiscoveryAPI) *ServiceDiscoveryActivities {
    return &ServiceDiscoveryActivities{client: client}
}

func (a *ServiceDiscoveryActivities) CreateHttpNamespace(input *servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error) {
    return a.client.CreateHttpNamespace(input)
}

func (a *ServiceDiscoveryActivities) CreatePrivateDnsNamespace(input *servicediscovery.CreatePrivateDnsNamespaceInput) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error) {
    return a.client.CreatePrivateDnsNamespace(input)
}

func (a *ServiceDiscoveryActivities) CreatePublicDnsNamespace(input *servicediscovery.CreatePublicDnsNamespaceInput) (*servicediscovery.CreatePublicDnsNamespaceOutput, error) {
    return a.client.CreatePublicDnsNamespace(input)
}

func (a *ServiceDiscoveryActivities) CreateService(input *servicediscovery.CreateServiceInput) (*servicediscovery.CreateServiceOutput, error) {
    return a.client.CreateService(input)
}

func (a *ServiceDiscoveryActivities) DeleteNamespace(input *servicediscovery.DeleteNamespaceInput) (*servicediscovery.DeleteNamespaceOutput, error) {
    return a.client.DeleteNamespace(input)
}

func (a *ServiceDiscoveryActivities) DeleteService(input *servicediscovery.DeleteServiceInput) (*servicediscovery.DeleteServiceOutput, error) {
    return a.client.DeleteService(input)
}

func (a *ServiceDiscoveryActivities) DeregisterInstance(input *servicediscovery.DeregisterInstanceInput) (*servicediscovery.DeregisterInstanceOutput, error) {
    return a.client.DeregisterInstance(input)
}

func (a *ServiceDiscoveryActivities) DiscoverInstances(input *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error) {
    return a.client.DiscoverInstances(input)
}

func (a *ServiceDiscoveryActivities) GetInstance(input *servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error) {
    return a.client.GetInstance(input)
}

func (a *ServiceDiscoveryActivities) GetInstancesHealthStatus(input *servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
    return a.client.GetInstancesHealthStatus(input)
}

func (a *ServiceDiscoveryActivities) GetNamespace(input *servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error) {
    return a.client.GetNamespace(input)
}

func (a *ServiceDiscoveryActivities) GetOperation(input *servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error) {
    return a.client.GetOperation(input)
}

func (a *ServiceDiscoveryActivities) GetService(input *servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error) {
    return a.client.GetService(input)
}

func (a *ServiceDiscoveryActivities) ListInstances(input *servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error) {
    return a.client.ListInstances(input)
}

func (a *ServiceDiscoveryActivities) ListNamespaces(input *servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error) {
    return a.client.ListNamespaces(input)
}

func (a *ServiceDiscoveryActivities) ListOperations(input *servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error) {
    return a.client.ListOperations(input)
}

func (a *ServiceDiscoveryActivities) ListServices(input *servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error) {
    return a.client.ListServices(input)
}

func (a *ServiceDiscoveryActivities) ListTagsForResource(input *servicediscovery.ListTagsForResourceInput) (*servicediscovery.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ServiceDiscoveryActivities) RegisterInstance(input *servicediscovery.RegisterInstanceInput) (*servicediscovery.RegisterInstanceOutput, error) {
    return a.client.RegisterInstance(input)
}

func (a *ServiceDiscoveryActivities) TagResource(input *servicediscovery.TagResourceInput) (*servicediscovery.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ServiceDiscoveryActivities) UntagResource(input *servicediscovery.UntagResourceInput) (*servicediscovery.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ServiceDiscoveryActivities) UpdateInstanceCustomHealthStatus(input *servicediscovery.UpdateInstanceCustomHealthStatusInput) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error) {
    return a.client.UpdateInstanceCustomHealthStatus(input)
}

func (a *ServiceDiscoveryActivities) UpdateService(input *servicediscovery.UpdateServiceInput) (*servicediscovery.UpdateServiceOutput, error) {
    return a.client.UpdateService(input)
}
