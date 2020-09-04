package awsclients

import (
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ServiceDiscoveryClient interface {
    CreateHttpNamespace(ctx workflow.Context, input *servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error)
    CreateHttpNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreateHttpNamespaceInput) *ServicediscoveryCreateHttpNamespaceResult

    CreatePrivateDnsNamespace(ctx workflow.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error)
    CreatePrivateDnsNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) *ServicediscoveryCreatePrivateDnsNamespaceResult

    CreatePublicDnsNamespace(ctx workflow.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) (*servicediscovery.CreatePublicDnsNamespaceOutput, error)
    CreatePublicDnsNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) *ServicediscoveryCreatePublicDnsNamespaceResult

    CreateService(ctx workflow.Context, input *servicediscovery.CreateServiceInput) (*servicediscovery.CreateServiceOutput, error)
    CreateServiceAsync(ctx workflow.Context, input *servicediscovery.CreateServiceInput) *ServicediscoveryCreateServiceResult

    DeleteNamespace(ctx workflow.Context, input *servicediscovery.DeleteNamespaceInput) (*servicediscovery.DeleteNamespaceOutput, error)
    DeleteNamespaceAsync(ctx workflow.Context, input *servicediscovery.DeleteNamespaceInput) *ServicediscoveryDeleteNamespaceResult

    DeleteService(ctx workflow.Context, input *servicediscovery.DeleteServiceInput) (*servicediscovery.DeleteServiceOutput, error)
    DeleteServiceAsync(ctx workflow.Context, input *servicediscovery.DeleteServiceInput) *ServicediscoveryDeleteServiceResult

    DeregisterInstance(ctx workflow.Context, input *servicediscovery.DeregisterInstanceInput) (*servicediscovery.DeregisterInstanceOutput, error)
    DeregisterInstanceAsync(ctx workflow.Context, input *servicediscovery.DeregisterInstanceInput) *ServicediscoveryDeregisterInstanceResult

    DiscoverInstances(ctx workflow.Context, input *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error)
    DiscoverInstancesAsync(ctx workflow.Context, input *servicediscovery.DiscoverInstancesInput) *ServicediscoveryDiscoverInstancesResult

    GetInstance(ctx workflow.Context, input *servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error)
    GetInstanceAsync(ctx workflow.Context, input *servicediscovery.GetInstanceInput) *ServicediscoveryGetInstanceResult

    GetInstancesHealthStatus(ctx workflow.Context, input *servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error)
    GetInstancesHealthStatusAsync(ctx workflow.Context, input *servicediscovery.GetInstancesHealthStatusInput) *ServicediscoveryGetInstancesHealthStatusResult

    GetNamespace(ctx workflow.Context, input *servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error)
    GetNamespaceAsync(ctx workflow.Context, input *servicediscovery.GetNamespaceInput) *ServicediscoveryGetNamespaceResult

    GetOperation(ctx workflow.Context, input *servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error)
    GetOperationAsync(ctx workflow.Context, input *servicediscovery.GetOperationInput) *ServicediscoveryGetOperationResult

    GetService(ctx workflow.Context, input *servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error)
    GetServiceAsync(ctx workflow.Context, input *servicediscovery.GetServiceInput) *ServicediscoveryGetServiceResult

    ListInstances(ctx workflow.Context, input *servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error)
    ListInstancesAsync(ctx workflow.Context, input *servicediscovery.ListInstancesInput) *ServicediscoveryListInstancesResult

    ListNamespaces(ctx workflow.Context, input *servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error)
    ListNamespacesAsync(ctx workflow.Context, input *servicediscovery.ListNamespacesInput) *ServicediscoveryListNamespacesResult

    ListOperations(ctx workflow.Context, input *servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error)
    ListOperationsAsync(ctx workflow.Context, input *servicediscovery.ListOperationsInput) *ServicediscoveryListOperationsResult

    ListServices(ctx workflow.Context, input *servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error)
    ListServicesAsync(ctx workflow.Context, input *servicediscovery.ListServicesInput) *ServicediscoveryListServicesResult

    ListTagsForResource(ctx workflow.Context, input *servicediscovery.ListTagsForResourceInput) (*servicediscovery.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *servicediscovery.ListTagsForResourceInput) *ServicediscoveryListTagsForResourceResult

    RegisterInstance(ctx workflow.Context, input *servicediscovery.RegisterInstanceInput) (*servicediscovery.RegisterInstanceOutput, error)
    RegisterInstanceAsync(ctx workflow.Context, input *servicediscovery.RegisterInstanceInput) *ServicediscoveryRegisterInstanceResult

    TagResource(ctx workflow.Context, input *servicediscovery.TagResourceInput) (*servicediscovery.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *servicediscovery.TagResourceInput) *ServicediscoveryTagResourceResult

    UntagResource(ctx workflow.Context, input *servicediscovery.UntagResourceInput) (*servicediscovery.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *servicediscovery.UntagResourceInput) *ServicediscoveryUntagResourceResult

    UpdateInstanceCustomHealthStatus(ctx workflow.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error)
    UpdateInstanceCustomHealthStatusAsync(ctx workflow.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) *ServicediscoveryUpdateInstanceCustomHealthStatusResult

    UpdateService(ctx workflow.Context, input *servicediscovery.UpdateServiceInput) (*servicediscovery.UpdateServiceOutput, error)
    UpdateServiceAsync(ctx workflow.Context, input *servicediscovery.UpdateServiceInput) *ServicediscoveryUpdateServiceResult
}

type ServicediscoveryCreateHttpNamespaceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryCreateHttpNamespaceResult) Get(ctx workflow.Context) (*servicediscovery.CreateHttpNamespaceOutput, error) {
    var output servicediscovery.CreateHttpNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryCreatePrivateDnsNamespaceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryCreatePrivateDnsNamespaceResult) Get(ctx workflow.Context) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error) {
    var output servicediscovery.CreatePrivateDnsNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryCreatePublicDnsNamespaceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryCreatePublicDnsNamespaceResult) Get(ctx workflow.Context) (*servicediscovery.CreatePublicDnsNamespaceOutput, error) {
    var output servicediscovery.CreatePublicDnsNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryCreateServiceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryCreateServiceResult) Get(ctx workflow.Context) (*servicediscovery.CreateServiceOutput, error) {
    var output servicediscovery.CreateServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryDeleteNamespaceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryDeleteNamespaceResult) Get(ctx workflow.Context) (*servicediscovery.DeleteNamespaceOutput, error) {
    var output servicediscovery.DeleteNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryDeleteServiceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryDeleteServiceResult) Get(ctx workflow.Context) (*servicediscovery.DeleteServiceOutput, error) {
    var output servicediscovery.DeleteServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryDeregisterInstanceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryDeregisterInstanceResult) Get(ctx workflow.Context) (*servicediscovery.DeregisterInstanceOutput, error) {
    var output servicediscovery.DeregisterInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryDiscoverInstancesResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryDiscoverInstancesResult) Get(ctx workflow.Context) (*servicediscovery.DiscoverInstancesOutput, error) {
    var output servicediscovery.DiscoverInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryGetInstanceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryGetInstanceResult) Get(ctx workflow.Context) (*servicediscovery.GetInstanceOutput, error) {
    var output servicediscovery.GetInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryGetInstancesHealthStatusResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryGetInstancesHealthStatusResult) Get(ctx workflow.Context) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
    var output servicediscovery.GetInstancesHealthStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryGetNamespaceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryGetNamespaceResult) Get(ctx workflow.Context) (*servicediscovery.GetNamespaceOutput, error) {
    var output servicediscovery.GetNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryGetOperationResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryGetOperationResult) Get(ctx workflow.Context) (*servicediscovery.GetOperationOutput, error) {
    var output servicediscovery.GetOperationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryGetServiceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryGetServiceResult) Get(ctx workflow.Context) (*servicediscovery.GetServiceOutput, error) {
    var output servicediscovery.GetServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryListInstancesResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryListInstancesResult) Get(ctx workflow.Context) (*servicediscovery.ListInstancesOutput, error) {
    var output servicediscovery.ListInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryListNamespacesResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryListNamespacesResult) Get(ctx workflow.Context) (*servicediscovery.ListNamespacesOutput, error) {
    var output servicediscovery.ListNamespacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryListOperationsResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryListOperationsResult) Get(ctx workflow.Context) (*servicediscovery.ListOperationsOutput, error) {
    var output servicediscovery.ListOperationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryListServicesResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryListServicesResult) Get(ctx workflow.Context) (*servicediscovery.ListServicesOutput, error) {
    var output servicediscovery.ListServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryListTagsForResourceResult) Get(ctx workflow.Context) (*servicediscovery.ListTagsForResourceOutput, error) {
    var output servicediscovery.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryRegisterInstanceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryRegisterInstanceResult) Get(ctx workflow.Context) (*servicediscovery.RegisterInstanceOutput, error) {
    var output servicediscovery.RegisterInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryTagResourceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryTagResourceResult) Get(ctx workflow.Context) (*servicediscovery.TagResourceOutput, error) {
    var output servicediscovery.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryUntagResourceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryUntagResourceResult) Get(ctx workflow.Context) (*servicediscovery.UntagResourceOutput, error) {
    var output servicediscovery.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryUpdateInstanceCustomHealthStatusResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryUpdateInstanceCustomHealthStatusResult) Get(ctx workflow.Context) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error) {
    var output servicediscovery.UpdateInstanceCustomHealthStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicediscoveryUpdateServiceResult struct {
	Result workflow.Future
}

func (r *ServicediscoveryUpdateServiceResult) Get(ctx workflow.Context) (*servicediscovery.UpdateServiceOutput, error) {
    var output servicediscovery.UpdateServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServiceDiscoveryStub struct {
    activities awsactivities.ServiceDiscoveryActivities
}

func NewServiceDiscoveryStub() ServiceDiscoveryClient {
    return &ServiceDiscoveryStub{}
}

func (a *ServiceDiscoveryStub) CreateHttpNamespace(ctx workflow.Context, input *servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error) {
    var output servicediscovery.CreateHttpNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHttpNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) CreateHttpNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreateHttpNamespaceInput) *ServicediscoveryCreateHttpNamespaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHttpNamespace, input)
    return &ServicediscoveryCreateHttpNamespaceResult{Result: future}
}

func (a *ServiceDiscoveryStub) CreatePrivateDnsNamespace(ctx workflow.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error) {
    var output servicediscovery.CreatePrivateDnsNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePrivateDnsNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) CreatePrivateDnsNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) *ServicediscoveryCreatePrivateDnsNamespaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePrivateDnsNamespace, input)
    return &ServicediscoveryCreatePrivateDnsNamespaceResult{Result: future}
}

func (a *ServiceDiscoveryStub) CreatePublicDnsNamespace(ctx workflow.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) (*servicediscovery.CreatePublicDnsNamespaceOutput, error) {
    var output servicediscovery.CreatePublicDnsNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePublicDnsNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) CreatePublicDnsNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) *ServicediscoveryCreatePublicDnsNamespaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePublicDnsNamespace, input)
    return &ServicediscoveryCreatePublicDnsNamespaceResult{Result: future}
}

func (a *ServiceDiscoveryStub) CreateService(ctx workflow.Context, input *servicediscovery.CreateServiceInput) (*servicediscovery.CreateServiceOutput, error) {
    var output servicediscovery.CreateServiceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateService, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) CreateServiceAsync(ctx workflow.Context, input *servicediscovery.CreateServiceInput) *ServicediscoveryCreateServiceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateService, input)
    return &ServicediscoveryCreateServiceResult{Result: future}
}

func (a *ServiceDiscoveryStub) DeleteNamespace(ctx workflow.Context, input *servicediscovery.DeleteNamespaceInput) (*servicediscovery.DeleteNamespaceOutput, error) {
    var output servicediscovery.DeleteNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) DeleteNamespaceAsync(ctx workflow.Context, input *servicediscovery.DeleteNamespaceInput) *ServicediscoveryDeleteNamespaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNamespace, input)
    return &ServicediscoveryDeleteNamespaceResult{Result: future}
}

func (a *ServiceDiscoveryStub) DeleteService(ctx workflow.Context, input *servicediscovery.DeleteServiceInput) (*servicediscovery.DeleteServiceOutput, error) {
    var output servicediscovery.DeleteServiceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteService, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) DeleteServiceAsync(ctx workflow.Context, input *servicediscovery.DeleteServiceInput) *ServicediscoveryDeleteServiceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteService, input)
    return &ServicediscoveryDeleteServiceResult{Result: future}
}

func (a *ServiceDiscoveryStub) DeregisterInstance(ctx workflow.Context, input *servicediscovery.DeregisterInstanceInput) (*servicediscovery.DeregisterInstanceOutput, error) {
    var output servicediscovery.DeregisterInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) DeregisterInstanceAsync(ctx workflow.Context, input *servicediscovery.DeregisterInstanceInput) *ServicediscoveryDeregisterInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterInstance, input)
    return &ServicediscoveryDeregisterInstanceResult{Result: future}
}

func (a *ServiceDiscoveryStub) DiscoverInstances(ctx workflow.Context, input *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error) {
    var output servicediscovery.DiscoverInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DiscoverInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) DiscoverInstancesAsync(ctx workflow.Context, input *servicediscovery.DiscoverInstancesInput) *ServicediscoveryDiscoverInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DiscoverInstances, input)
    return &ServicediscoveryDiscoverInstancesResult{Result: future}
}

func (a *ServiceDiscoveryStub) GetInstance(ctx workflow.Context, input *servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error) {
    var output servicediscovery.GetInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) GetInstanceAsync(ctx workflow.Context, input *servicediscovery.GetInstanceInput) *ServicediscoveryGetInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetInstance, input)
    return &ServicediscoveryGetInstanceResult{Result: future}
}

func (a *ServiceDiscoveryStub) GetInstancesHealthStatus(ctx workflow.Context, input *servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
    var output servicediscovery.GetInstancesHealthStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInstancesHealthStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) GetInstancesHealthStatusAsync(ctx workflow.Context, input *servicediscovery.GetInstancesHealthStatusInput) *ServicediscoveryGetInstancesHealthStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetInstancesHealthStatus, input)
    return &ServicediscoveryGetInstancesHealthStatusResult{Result: future}
}

func (a *ServiceDiscoveryStub) GetNamespace(ctx workflow.Context, input *servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error) {
    var output servicediscovery.GetNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) GetNamespaceAsync(ctx workflow.Context, input *servicediscovery.GetNamespaceInput) *ServicediscoveryGetNamespaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetNamespace, input)
    return &ServicediscoveryGetNamespaceResult{Result: future}
}

func (a *ServiceDiscoveryStub) GetOperation(ctx workflow.Context, input *servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error) {
    var output servicediscovery.GetOperationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOperation, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) GetOperationAsync(ctx workflow.Context, input *servicediscovery.GetOperationInput) *ServicediscoveryGetOperationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetOperation, input)
    return &ServicediscoveryGetOperationResult{Result: future}
}

func (a *ServiceDiscoveryStub) GetService(ctx workflow.Context, input *servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error) {
    var output servicediscovery.GetServiceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetService, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) GetServiceAsync(ctx workflow.Context, input *servicediscovery.GetServiceInput) *ServicediscoveryGetServiceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetService, input)
    return &ServicediscoveryGetServiceResult{Result: future}
}

func (a *ServiceDiscoveryStub) ListInstances(ctx workflow.Context, input *servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error) {
    var output servicediscovery.ListInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) ListInstancesAsync(ctx workflow.Context, input *servicediscovery.ListInstancesInput) *ServicediscoveryListInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListInstances, input)
    return &ServicediscoveryListInstancesResult{Result: future}
}

func (a *ServiceDiscoveryStub) ListNamespaces(ctx workflow.Context, input *servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error) {
    var output servicediscovery.ListNamespacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNamespaces, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) ListNamespacesAsync(ctx workflow.Context, input *servicediscovery.ListNamespacesInput) *ServicediscoveryListNamespacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListNamespaces, input)
    return &ServicediscoveryListNamespacesResult{Result: future}
}

func (a *ServiceDiscoveryStub) ListOperations(ctx workflow.Context, input *servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error) {
    var output servicediscovery.ListOperationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOperations, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) ListOperationsAsync(ctx workflow.Context, input *servicediscovery.ListOperationsInput) *ServicediscoveryListOperationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOperations, input)
    return &ServicediscoveryListOperationsResult{Result: future}
}

func (a *ServiceDiscoveryStub) ListServices(ctx workflow.Context, input *servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error) {
    var output servicediscovery.ListServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServices, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) ListServicesAsync(ctx workflow.Context, input *servicediscovery.ListServicesInput) *ServicediscoveryListServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServices, input)
    return &ServicediscoveryListServicesResult{Result: future}
}

func (a *ServiceDiscoveryStub) ListTagsForResource(ctx workflow.Context, input *servicediscovery.ListTagsForResourceInput) (*servicediscovery.ListTagsForResourceOutput, error) {
    var output servicediscovery.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) ListTagsForResourceAsync(ctx workflow.Context, input *servicediscovery.ListTagsForResourceInput) *ServicediscoveryListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &ServicediscoveryListTagsForResourceResult{Result: future}
}

func (a *ServiceDiscoveryStub) RegisterInstance(ctx workflow.Context, input *servicediscovery.RegisterInstanceInput) (*servicediscovery.RegisterInstanceOutput, error) {
    var output servicediscovery.RegisterInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) RegisterInstanceAsync(ctx workflow.Context, input *servicediscovery.RegisterInstanceInput) *ServicediscoveryRegisterInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterInstance, input)
    return &ServicediscoveryRegisterInstanceResult{Result: future}
}

func (a *ServiceDiscoveryStub) TagResource(ctx workflow.Context, input *servicediscovery.TagResourceInput) (*servicediscovery.TagResourceOutput, error) {
    var output servicediscovery.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) TagResourceAsync(ctx workflow.Context, input *servicediscovery.TagResourceInput) *ServicediscoveryTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &ServicediscoveryTagResourceResult{Result: future}
}

func (a *ServiceDiscoveryStub) UntagResource(ctx workflow.Context, input *servicediscovery.UntagResourceInput) (*servicediscovery.UntagResourceOutput, error) {
    var output servicediscovery.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) UntagResourceAsync(ctx workflow.Context, input *servicediscovery.UntagResourceInput) *ServicediscoveryUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &ServicediscoveryUntagResourceResult{Result: future}
}

func (a *ServiceDiscoveryStub) UpdateInstanceCustomHealthStatus(ctx workflow.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error) {
    var output servicediscovery.UpdateInstanceCustomHealthStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateInstanceCustomHealthStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) UpdateInstanceCustomHealthStatusAsync(ctx workflow.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) *ServicediscoveryUpdateInstanceCustomHealthStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateInstanceCustomHealthStatus, input)
    return &ServicediscoveryUpdateInstanceCustomHealthStatusResult{Result: future}
}

func (a *ServiceDiscoveryStub) UpdateService(ctx workflow.Context, input *servicediscovery.UpdateServiceInput) (*servicediscovery.UpdateServiceOutput, error) {
    var output servicediscovery.UpdateServiceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateService, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceDiscoveryStub) UpdateServiceAsync(ctx workflow.Context, input *servicediscovery.UpdateServiceInput) *ServicediscoveryUpdateServiceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateService, input)
    return &ServicediscoveryUpdateServiceResult{Result: future}
}
