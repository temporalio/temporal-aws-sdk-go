package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ecs"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ECSClient interface {
       CreateCapacityProvider(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) (*ecs.CreateCapacityProviderOutput, error)
       CreateCapacityProviderAsync(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) *EcsCreateCapacityProviderResult

       CreateCluster(ctx workflow.Context, input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error)
       CreateClusterAsync(ctx workflow.Context, input *ecs.CreateClusterInput) *EcsCreateClusterResult

       CreateService(ctx workflow.Context, input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error)
       CreateServiceAsync(ctx workflow.Context, input *ecs.CreateServiceInput) *EcsCreateServiceResult

       CreateTaskSet(ctx workflow.Context, input *ecs.CreateTaskSetInput) (*ecs.CreateTaskSetOutput, error)
       CreateTaskSetAsync(ctx workflow.Context, input *ecs.CreateTaskSetInput) *EcsCreateTaskSetResult

       DeleteAccountSetting(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) (*ecs.DeleteAccountSettingOutput, error)
       DeleteAccountSettingAsync(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) *EcsDeleteAccountSettingResult

       DeleteAttributes(ctx workflow.Context, input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error)
       DeleteAttributesAsync(ctx workflow.Context, input *ecs.DeleteAttributesInput) *EcsDeleteAttributesResult

       DeleteCapacityProvider(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) (*ecs.DeleteCapacityProviderOutput, error)
       DeleteCapacityProviderAsync(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) *EcsDeleteCapacityProviderResult

       DeleteCluster(ctx workflow.Context, input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error)
       DeleteClusterAsync(ctx workflow.Context, input *ecs.DeleteClusterInput) *EcsDeleteClusterResult

       DeleteService(ctx workflow.Context, input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error)
       DeleteServiceAsync(ctx workflow.Context, input *ecs.DeleteServiceInput) *EcsDeleteServiceResult

       DeleteTaskSet(ctx workflow.Context, input *ecs.DeleteTaskSetInput) (*ecs.DeleteTaskSetOutput, error)
       DeleteTaskSetAsync(ctx workflow.Context, input *ecs.DeleteTaskSetInput) *EcsDeleteTaskSetResult

       DeregisterContainerInstance(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error)
       DeregisterContainerInstanceAsync(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) *EcsDeregisterContainerInstanceResult

       DeregisterTaskDefinition(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error)
       DeregisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) *EcsDeregisterTaskDefinitionResult

       DescribeCapacityProviders(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) (*ecs.DescribeCapacityProvidersOutput, error)
       DescribeCapacityProvidersAsync(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) *EcsDescribeCapacityProvidersResult

       DescribeClusters(ctx workflow.Context, input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error)
       DescribeClustersAsync(ctx workflow.Context, input *ecs.DescribeClustersInput) *EcsDescribeClustersResult

       DescribeContainerInstances(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error)
       DescribeContainerInstancesAsync(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) *EcsDescribeContainerInstancesResult

       DescribeServices(ctx workflow.Context, input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error)
       DescribeServicesAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *EcsDescribeServicesResult

       DescribeTaskDefinition(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error)
       DescribeTaskDefinitionAsync(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) *EcsDescribeTaskDefinitionResult

       DescribeTaskSets(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) (*ecs.DescribeTaskSetsOutput, error)
       DescribeTaskSetsAsync(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) *EcsDescribeTaskSetsResult

       DescribeTasks(ctx workflow.Context, input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error)
       DescribeTasksAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *EcsDescribeTasksResult

       DiscoverPollEndpoint(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error)
       DiscoverPollEndpointAsync(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) *EcsDiscoverPollEndpointResult

       ListAccountSettings(ctx workflow.Context, input *ecs.ListAccountSettingsInput) (*ecs.ListAccountSettingsOutput, error)
       ListAccountSettingsAsync(ctx workflow.Context, input *ecs.ListAccountSettingsInput) *EcsListAccountSettingsResult

       ListAttributes(ctx workflow.Context, input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error)
       ListAttributesAsync(ctx workflow.Context, input *ecs.ListAttributesInput) *EcsListAttributesResult

       ListClusters(ctx workflow.Context, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error)
       ListClustersAsync(ctx workflow.Context, input *ecs.ListClustersInput) *EcsListClustersResult

       ListContainerInstances(ctx workflow.Context, input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error)
       ListContainerInstancesAsync(ctx workflow.Context, input *ecs.ListContainerInstancesInput) *EcsListContainerInstancesResult

       ListServices(ctx workflow.Context, input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error)
       ListServicesAsync(ctx workflow.Context, input *ecs.ListServicesInput) *EcsListServicesResult

       ListTagsForResource(ctx workflow.Context, input *ecs.ListTagsForResourceInput) (*ecs.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *ecs.ListTagsForResourceInput) *EcsListTagsForResourceResult

       ListTaskDefinitionFamilies(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error)
       ListTaskDefinitionFamiliesAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) *EcsListTaskDefinitionFamiliesResult

       ListTaskDefinitions(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error)
       ListTaskDefinitionsAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) *EcsListTaskDefinitionsResult

       ListTasks(ctx workflow.Context, input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error)
       ListTasksAsync(ctx workflow.Context, input *ecs.ListTasksInput) *EcsListTasksResult

       PutAccountSetting(ctx workflow.Context, input *ecs.PutAccountSettingInput) (*ecs.PutAccountSettingOutput, error)
       PutAccountSettingAsync(ctx workflow.Context, input *ecs.PutAccountSettingInput) *EcsPutAccountSettingResult

       PutAccountSettingDefault(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) (*ecs.PutAccountSettingDefaultOutput, error)
       PutAccountSettingDefaultAsync(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) *EcsPutAccountSettingDefaultResult

       PutAttributes(ctx workflow.Context, input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error)
       PutAttributesAsync(ctx workflow.Context, input *ecs.PutAttributesInput) *EcsPutAttributesResult

       PutClusterCapacityProviders(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) (*ecs.PutClusterCapacityProvidersOutput, error)
       PutClusterCapacityProvidersAsync(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) *EcsPutClusterCapacityProvidersResult

       RegisterContainerInstance(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error)
       RegisterContainerInstanceAsync(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) *EcsRegisterContainerInstanceResult

       RegisterTaskDefinition(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error)
       RegisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) *EcsRegisterTaskDefinitionResult

       RunTask(ctx workflow.Context, input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error)
       RunTaskAsync(ctx workflow.Context, input *ecs.RunTaskInput) *EcsRunTaskResult

       StartTask(ctx workflow.Context, input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error)
       StartTaskAsync(ctx workflow.Context, input *ecs.StartTaskInput) *EcsStartTaskResult

       StopTask(ctx workflow.Context, input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error)
       StopTaskAsync(ctx workflow.Context, input *ecs.StopTaskInput) *EcsStopTaskResult

       SubmitAttachmentStateChanges(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) (*ecs.SubmitAttachmentStateChangesOutput, error)
       SubmitAttachmentStateChangesAsync(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) *EcsSubmitAttachmentStateChangesResult

       SubmitContainerStateChange(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error)
       SubmitContainerStateChangeAsync(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) *EcsSubmitContainerStateChangeResult

       SubmitTaskStateChange(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error)
       SubmitTaskStateChangeAsync(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) *EcsSubmitTaskStateChangeResult

       TagResource(ctx workflow.Context, input *ecs.TagResourceInput) (*ecs.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *ecs.TagResourceInput) *EcsTagResourceResult

       UntagResource(ctx workflow.Context, input *ecs.UntagResourceInput) (*ecs.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *ecs.UntagResourceInput) *EcsUntagResourceResult

       UpdateClusterSettings(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) (*ecs.UpdateClusterSettingsOutput, error)
       UpdateClusterSettingsAsync(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) *EcsUpdateClusterSettingsResult

       UpdateContainerAgent(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error)
       UpdateContainerAgentAsync(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) *EcsUpdateContainerAgentResult

       UpdateContainerInstancesState(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error)
       UpdateContainerInstancesStateAsync(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) *EcsUpdateContainerInstancesStateResult

       UpdateService(ctx workflow.Context, input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error)
       UpdateServiceAsync(ctx workflow.Context, input *ecs.UpdateServiceInput) *EcsUpdateServiceResult

       UpdateServicePrimaryTaskSet(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) (*ecs.UpdateServicePrimaryTaskSetOutput, error)
       UpdateServicePrimaryTaskSetAsync(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) *EcsUpdateServicePrimaryTaskSetResult

       UpdateTaskSet(ctx workflow.Context, input *ecs.UpdateTaskSetInput) (*ecs.UpdateTaskSetOutput, error)
       UpdateTaskSetAsync(ctx workflow.Context, input *ecs.UpdateTaskSetInput) *EcsUpdateTaskSetResult

       WaitUntilServicesInactive(ctx workflow.Context, input *ecs.DescribeServicesInput) error
       WaitUntilServicesStable(ctx workflow.Context, input *ecs.DescribeServicesInput) error
       WaitUntilTasksRunning(ctx workflow.Context, input *ecs.DescribeTasksInput) error
       WaitUntilTasksStopped(ctx workflow.Context, input *ecs.DescribeTasksInput) error}

type EcsCreateCapacityProviderResult struct {
	Result workflow.Future
}

func (r *EcsCreateCapacityProviderResult) Get(ctx workflow.Context) (*ecs.CreateCapacityProviderOutput, error) {
    var output ecs.CreateCapacityProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsCreateClusterResult struct {
	Result workflow.Future
}

func (r *EcsCreateClusterResult) Get(ctx workflow.Context) (*ecs.CreateClusterOutput, error) {
    var output ecs.CreateClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsCreateServiceResult struct {
	Result workflow.Future
}

func (r *EcsCreateServiceResult) Get(ctx workflow.Context) (*ecs.CreateServiceOutput, error) {
    var output ecs.CreateServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsCreateTaskSetResult struct {
	Result workflow.Future
}

func (r *EcsCreateTaskSetResult) Get(ctx workflow.Context) (*ecs.CreateTaskSetOutput, error) {
    var output ecs.CreateTaskSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeleteAccountSettingResult struct {
	Result workflow.Future
}

func (r *EcsDeleteAccountSettingResult) Get(ctx workflow.Context) (*ecs.DeleteAccountSettingOutput, error) {
    var output ecs.DeleteAccountSettingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeleteAttributesResult struct {
	Result workflow.Future
}

func (r *EcsDeleteAttributesResult) Get(ctx workflow.Context) (*ecs.DeleteAttributesOutput, error) {
    var output ecs.DeleteAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeleteCapacityProviderResult struct {
	Result workflow.Future
}

func (r *EcsDeleteCapacityProviderResult) Get(ctx workflow.Context) (*ecs.DeleteCapacityProviderOutput, error) {
    var output ecs.DeleteCapacityProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeleteClusterResult struct {
	Result workflow.Future
}

func (r *EcsDeleteClusterResult) Get(ctx workflow.Context) (*ecs.DeleteClusterOutput, error) {
    var output ecs.DeleteClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeleteServiceResult struct {
	Result workflow.Future
}

func (r *EcsDeleteServiceResult) Get(ctx workflow.Context) (*ecs.DeleteServiceOutput, error) {
    var output ecs.DeleteServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeleteTaskSetResult struct {
	Result workflow.Future
}

func (r *EcsDeleteTaskSetResult) Get(ctx workflow.Context) (*ecs.DeleteTaskSetOutput, error) {
    var output ecs.DeleteTaskSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeregisterContainerInstanceResult struct {
	Result workflow.Future
}

func (r *EcsDeregisterContainerInstanceResult) Get(ctx workflow.Context) (*ecs.DeregisterContainerInstanceOutput, error) {
    var output ecs.DeregisterContainerInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDeregisterTaskDefinitionResult struct {
	Result workflow.Future
}

func (r *EcsDeregisterTaskDefinitionResult) Get(ctx workflow.Context) (*ecs.DeregisterTaskDefinitionOutput, error) {
    var output ecs.DeregisterTaskDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDescribeCapacityProvidersResult struct {
	Result workflow.Future
}

func (r *EcsDescribeCapacityProvidersResult) Get(ctx workflow.Context) (*ecs.DescribeCapacityProvidersOutput, error) {
    var output ecs.DescribeCapacityProvidersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDescribeClustersResult struct {
	Result workflow.Future
}

func (r *EcsDescribeClustersResult) Get(ctx workflow.Context) (*ecs.DescribeClustersOutput, error) {
    var output ecs.DescribeClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDescribeContainerInstancesResult struct {
	Result workflow.Future
}

func (r *EcsDescribeContainerInstancesResult) Get(ctx workflow.Context) (*ecs.DescribeContainerInstancesOutput, error) {
    var output ecs.DescribeContainerInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDescribeServicesResult struct {
	Result workflow.Future
}

func (r *EcsDescribeServicesResult) Get(ctx workflow.Context) (*ecs.DescribeServicesOutput, error) {
    var output ecs.DescribeServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDescribeTaskDefinitionResult struct {
	Result workflow.Future
}

func (r *EcsDescribeTaskDefinitionResult) Get(ctx workflow.Context) (*ecs.DescribeTaskDefinitionOutput, error) {
    var output ecs.DescribeTaskDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDescribeTaskSetsResult struct {
	Result workflow.Future
}

func (r *EcsDescribeTaskSetsResult) Get(ctx workflow.Context) (*ecs.DescribeTaskSetsOutput, error) {
    var output ecs.DescribeTaskSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDescribeTasksResult struct {
	Result workflow.Future
}

func (r *EcsDescribeTasksResult) Get(ctx workflow.Context) (*ecs.DescribeTasksOutput, error) {
    var output ecs.DescribeTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsDiscoverPollEndpointResult struct {
	Result workflow.Future
}

func (r *EcsDiscoverPollEndpointResult) Get(ctx workflow.Context) (*ecs.DiscoverPollEndpointOutput, error) {
    var output ecs.DiscoverPollEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListAccountSettingsResult struct {
	Result workflow.Future
}

func (r *EcsListAccountSettingsResult) Get(ctx workflow.Context) (*ecs.ListAccountSettingsOutput, error) {
    var output ecs.ListAccountSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListAttributesResult struct {
	Result workflow.Future
}

func (r *EcsListAttributesResult) Get(ctx workflow.Context) (*ecs.ListAttributesOutput, error) {
    var output ecs.ListAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListClustersResult struct {
	Result workflow.Future
}

func (r *EcsListClustersResult) Get(ctx workflow.Context) (*ecs.ListClustersOutput, error) {
    var output ecs.ListClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListContainerInstancesResult struct {
	Result workflow.Future
}

func (r *EcsListContainerInstancesResult) Get(ctx workflow.Context) (*ecs.ListContainerInstancesOutput, error) {
    var output ecs.ListContainerInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListServicesResult struct {
	Result workflow.Future
}

func (r *EcsListServicesResult) Get(ctx workflow.Context) (*ecs.ListServicesOutput, error) {
    var output ecs.ListServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *EcsListTagsForResourceResult) Get(ctx workflow.Context) (*ecs.ListTagsForResourceOutput, error) {
    var output ecs.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListTaskDefinitionFamiliesResult struct {
	Result workflow.Future
}

func (r *EcsListTaskDefinitionFamiliesResult) Get(ctx workflow.Context) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
    var output ecs.ListTaskDefinitionFamiliesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListTaskDefinitionsResult struct {
	Result workflow.Future
}

func (r *EcsListTaskDefinitionsResult) Get(ctx workflow.Context) (*ecs.ListTaskDefinitionsOutput, error) {
    var output ecs.ListTaskDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsListTasksResult struct {
	Result workflow.Future
}

func (r *EcsListTasksResult) Get(ctx workflow.Context) (*ecs.ListTasksOutput, error) {
    var output ecs.ListTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsPutAccountSettingResult struct {
	Result workflow.Future
}

func (r *EcsPutAccountSettingResult) Get(ctx workflow.Context) (*ecs.PutAccountSettingOutput, error) {
    var output ecs.PutAccountSettingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsPutAccountSettingDefaultResult struct {
	Result workflow.Future
}

func (r *EcsPutAccountSettingDefaultResult) Get(ctx workflow.Context) (*ecs.PutAccountSettingDefaultOutput, error) {
    var output ecs.PutAccountSettingDefaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsPutAttributesResult struct {
	Result workflow.Future
}

func (r *EcsPutAttributesResult) Get(ctx workflow.Context) (*ecs.PutAttributesOutput, error) {
    var output ecs.PutAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsPutClusterCapacityProvidersResult struct {
	Result workflow.Future
}

func (r *EcsPutClusterCapacityProvidersResult) Get(ctx workflow.Context) (*ecs.PutClusterCapacityProvidersOutput, error) {
    var output ecs.PutClusterCapacityProvidersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsRegisterContainerInstanceResult struct {
	Result workflow.Future
}

func (r *EcsRegisterContainerInstanceResult) Get(ctx workflow.Context) (*ecs.RegisterContainerInstanceOutput, error) {
    var output ecs.RegisterContainerInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsRegisterTaskDefinitionResult struct {
	Result workflow.Future
}

func (r *EcsRegisterTaskDefinitionResult) Get(ctx workflow.Context) (*ecs.RegisterTaskDefinitionOutput, error) {
    var output ecs.RegisterTaskDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsRunTaskResult struct {
	Result workflow.Future
}

func (r *EcsRunTaskResult) Get(ctx workflow.Context) (*ecs.RunTaskOutput, error) {
    var output ecs.RunTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsStartTaskResult struct {
	Result workflow.Future
}

func (r *EcsStartTaskResult) Get(ctx workflow.Context) (*ecs.StartTaskOutput, error) {
    var output ecs.StartTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsStopTaskResult struct {
	Result workflow.Future
}

func (r *EcsStopTaskResult) Get(ctx workflow.Context) (*ecs.StopTaskOutput, error) {
    var output ecs.StopTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsSubmitAttachmentStateChangesResult struct {
	Result workflow.Future
}

func (r *EcsSubmitAttachmentStateChangesResult) Get(ctx workflow.Context) (*ecs.SubmitAttachmentStateChangesOutput, error) {
    var output ecs.SubmitAttachmentStateChangesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsSubmitContainerStateChangeResult struct {
	Result workflow.Future
}

func (r *EcsSubmitContainerStateChangeResult) Get(ctx workflow.Context) (*ecs.SubmitContainerStateChangeOutput, error) {
    var output ecs.SubmitContainerStateChangeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsSubmitTaskStateChangeResult struct {
	Result workflow.Future
}

func (r *EcsSubmitTaskStateChangeResult) Get(ctx workflow.Context) (*ecs.SubmitTaskStateChangeOutput, error) {
    var output ecs.SubmitTaskStateChangeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsTagResourceResult struct {
	Result workflow.Future
}

func (r *EcsTagResourceResult) Get(ctx workflow.Context) (*ecs.TagResourceOutput, error) {
    var output ecs.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsUntagResourceResult struct {
	Result workflow.Future
}

func (r *EcsUntagResourceResult) Get(ctx workflow.Context) (*ecs.UntagResourceOutput, error) {
    var output ecs.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsUpdateClusterSettingsResult struct {
	Result workflow.Future
}

func (r *EcsUpdateClusterSettingsResult) Get(ctx workflow.Context) (*ecs.UpdateClusterSettingsOutput, error) {
    var output ecs.UpdateClusterSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsUpdateContainerAgentResult struct {
	Result workflow.Future
}

func (r *EcsUpdateContainerAgentResult) Get(ctx workflow.Context) (*ecs.UpdateContainerAgentOutput, error) {
    var output ecs.UpdateContainerAgentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsUpdateContainerInstancesStateResult struct {
	Result workflow.Future
}

func (r *EcsUpdateContainerInstancesStateResult) Get(ctx workflow.Context) (*ecs.UpdateContainerInstancesStateOutput, error) {
    var output ecs.UpdateContainerInstancesStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsUpdateServiceResult struct {
	Result workflow.Future
}

func (r *EcsUpdateServiceResult) Get(ctx workflow.Context) (*ecs.UpdateServiceOutput, error) {
    var output ecs.UpdateServiceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsUpdateServicePrimaryTaskSetResult struct {
	Result workflow.Future
}

func (r *EcsUpdateServicePrimaryTaskSetResult) Get(ctx workflow.Context) (*ecs.UpdateServicePrimaryTaskSetOutput, error) {
    var output ecs.UpdateServicePrimaryTaskSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcsUpdateTaskSetResult struct {
	Result workflow.Future
}

func (r *EcsUpdateTaskSetResult) Get(ctx workflow.Context) (*ecs.UpdateTaskSetOutput, error) {
    var output ecs.UpdateTaskSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ECSStub struct {
    activities awsactivities.ECSActivities
}

func NewECSStub() ECSClient {
    return &ECSStub{}
}

func (a *ECSStub) CreateCapacityProvider(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) (*ecs.CreateCapacityProviderOutput, error) {
    var output ecs.CreateCapacityProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCapacityProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) CreateCapacityProviderAsync(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) *EcsCreateCapacityProviderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCapacityProvider, input)
    return &EcsCreateCapacityProviderResult{Result: future}
}

func (a *ECSStub) CreateCluster(ctx workflow.Context, input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
    var output ecs.CreateClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) CreateClusterAsync(ctx workflow.Context, input *ecs.CreateClusterInput) *EcsCreateClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input)
    return &EcsCreateClusterResult{Result: future}
}

func (a *ECSStub) CreateService(ctx workflow.Context, input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
    var output ecs.CreateServiceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateService, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) CreateServiceAsync(ctx workflow.Context, input *ecs.CreateServiceInput) *EcsCreateServiceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateService, input)
    return &EcsCreateServiceResult{Result: future}
}

func (a *ECSStub) CreateTaskSet(ctx workflow.Context, input *ecs.CreateTaskSetInput) (*ecs.CreateTaskSetOutput, error) {
    var output ecs.CreateTaskSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTaskSet, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) CreateTaskSetAsync(ctx workflow.Context, input *ecs.CreateTaskSetInput) *EcsCreateTaskSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTaskSet, input)
    return &EcsCreateTaskSetResult{Result: future}
}

func (a *ECSStub) DeleteAccountSetting(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) (*ecs.DeleteAccountSettingOutput, error) {
    var output ecs.DeleteAccountSettingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccountSetting, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeleteAccountSettingAsync(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) *EcsDeleteAccountSettingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAccountSetting, input)
    return &EcsDeleteAccountSettingResult{Result: future}
}

func (a *ECSStub) DeleteAttributes(ctx workflow.Context, input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error) {
    var output ecs.DeleteAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeleteAttributesAsync(ctx workflow.Context, input *ecs.DeleteAttributesInput) *EcsDeleteAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAttributes, input)
    return &EcsDeleteAttributesResult{Result: future}
}

func (a *ECSStub) DeleteCapacityProvider(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) (*ecs.DeleteCapacityProviderOutput, error) {
    var output ecs.DeleteCapacityProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCapacityProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeleteCapacityProviderAsync(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) *EcsDeleteCapacityProviderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCapacityProvider, input)
    return &EcsDeleteCapacityProviderResult{Result: future}
}

func (a *ECSStub) DeleteCluster(ctx workflow.Context, input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
    var output ecs.DeleteClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeleteClusterAsync(ctx workflow.Context, input *ecs.DeleteClusterInput) *EcsDeleteClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input)
    return &EcsDeleteClusterResult{Result: future}
}

func (a *ECSStub) DeleteService(ctx workflow.Context, input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error) {
    var output ecs.DeleteServiceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteService, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeleteServiceAsync(ctx workflow.Context, input *ecs.DeleteServiceInput) *EcsDeleteServiceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteService, input)
    return &EcsDeleteServiceResult{Result: future}
}

func (a *ECSStub) DeleteTaskSet(ctx workflow.Context, input *ecs.DeleteTaskSetInput) (*ecs.DeleteTaskSetOutput, error) {
    var output ecs.DeleteTaskSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTaskSet, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeleteTaskSetAsync(ctx workflow.Context, input *ecs.DeleteTaskSetInput) *EcsDeleteTaskSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTaskSet, input)
    return &EcsDeleteTaskSetResult{Result: future}
}

func (a *ECSStub) DeregisterContainerInstance(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error) {
    var output ecs.DeregisterContainerInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterContainerInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeregisterContainerInstanceAsync(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) *EcsDeregisterContainerInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterContainerInstance, input)
    return &EcsDeregisterContainerInstanceResult{Result: future}
}

func (a *ECSStub) DeregisterTaskDefinition(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error) {
    var output ecs.DeregisterTaskDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterTaskDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DeregisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) *EcsDeregisterTaskDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterTaskDefinition, input)
    return &EcsDeregisterTaskDefinitionResult{Result: future}
}

func (a *ECSStub) DescribeCapacityProviders(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) (*ecs.DescribeCapacityProvidersOutput, error) {
    var output ecs.DescribeCapacityProvidersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCapacityProviders, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DescribeCapacityProvidersAsync(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) *EcsDescribeCapacityProvidersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCapacityProviders, input)
    return &EcsDescribeCapacityProvidersResult{Result: future}
}

func (a *ECSStub) DescribeClusters(ctx workflow.Context, input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error) {
    var output ecs.DescribeClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DescribeClustersAsync(ctx workflow.Context, input *ecs.DescribeClustersInput) *EcsDescribeClustersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input)
    return &EcsDescribeClustersResult{Result: future}
}

func (a *ECSStub) DescribeContainerInstances(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
    var output ecs.DescribeContainerInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeContainerInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DescribeContainerInstancesAsync(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) *EcsDescribeContainerInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeContainerInstances, input)
    return &EcsDescribeContainerInstancesResult{Result: future}
}

func (a *ECSStub) DescribeServices(ctx workflow.Context, input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error) {
    var output ecs.DescribeServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServices, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DescribeServicesAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *EcsDescribeServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeServices, input)
    return &EcsDescribeServicesResult{Result: future}
}

func (a *ECSStub) DescribeTaskDefinition(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
    var output ecs.DescribeTaskDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTaskDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DescribeTaskDefinitionAsync(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) *EcsDescribeTaskDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTaskDefinition, input)
    return &EcsDescribeTaskDefinitionResult{Result: future}
}

func (a *ECSStub) DescribeTaskSets(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) (*ecs.DescribeTaskSetsOutput, error) {
    var output ecs.DescribeTaskSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTaskSets, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DescribeTaskSetsAsync(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) *EcsDescribeTaskSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTaskSets, input)
    return &EcsDescribeTaskSetsResult{Result: future}
}

func (a *ECSStub) DescribeTasks(ctx workflow.Context, input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
    var output ecs.DescribeTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DescribeTasksAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *EcsDescribeTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTasks, input)
    return &EcsDescribeTasksResult{Result: future}
}

func (a *ECSStub) DiscoverPollEndpoint(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
    var output ecs.DiscoverPollEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DiscoverPollEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) DiscoverPollEndpointAsync(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) *EcsDiscoverPollEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DiscoverPollEndpoint, input)
    return &EcsDiscoverPollEndpointResult{Result: future}
}

func (a *ECSStub) ListAccountSettings(ctx workflow.Context, input *ecs.ListAccountSettingsInput) (*ecs.ListAccountSettingsOutput, error) {
    var output ecs.ListAccountSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAccountSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListAccountSettingsAsync(ctx workflow.Context, input *ecs.ListAccountSettingsInput) *EcsListAccountSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAccountSettings, input)
    return &EcsListAccountSettingsResult{Result: future}
}

func (a *ECSStub) ListAttributes(ctx workflow.Context, input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error) {
    var output ecs.ListAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListAttributesAsync(ctx workflow.Context, input *ecs.ListAttributesInput) *EcsListAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAttributes, input)
    return &EcsListAttributesResult{Result: future}
}

func (a *ECSStub) ListClusters(ctx workflow.Context, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
    var output ecs.ListClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListClustersAsync(ctx workflow.Context, input *ecs.ListClustersInput) *EcsListClustersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListClusters, input)
    return &EcsListClustersResult{Result: future}
}

func (a *ECSStub) ListContainerInstances(ctx workflow.Context, input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
    var output ecs.ListContainerInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListContainerInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListContainerInstancesAsync(ctx workflow.Context, input *ecs.ListContainerInstancesInput) *EcsListContainerInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListContainerInstances, input)
    return &EcsListContainerInstancesResult{Result: future}
}

func (a *ECSStub) ListServices(ctx workflow.Context, input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
    var output ecs.ListServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServices, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListServicesAsync(ctx workflow.Context, input *ecs.ListServicesInput) *EcsListServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServices, input)
    return &EcsListServicesResult{Result: future}
}

func (a *ECSStub) ListTagsForResource(ctx workflow.Context, input *ecs.ListTagsForResourceInput) (*ecs.ListTagsForResourceOutput, error) {
    var output ecs.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListTagsForResourceAsync(ctx workflow.Context, input *ecs.ListTagsForResourceInput) *EcsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &EcsListTagsForResourceResult{Result: future}
}

func (a *ECSStub) ListTaskDefinitionFamilies(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
    var output ecs.ListTaskDefinitionFamiliesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTaskDefinitionFamilies, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListTaskDefinitionFamiliesAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) *EcsListTaskDefinitionFamiliesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTaskDefinitionFamilies, input)
    return &EcsListTaskDefinitionFamiliesResult{Result: future}
}

func (a *ECSStub) ListTaskDefinitions(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
    var output ecs.ListTaskDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTaskDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListTaskDefinitionsAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) *EcsListTaskDefinitionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTaskDefinitions, input)
    return &EcsListTaskDefinitionsResult{Result: future}
}

func (a *ECSStub) ListTasks(ctx workflow.Context, input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
    var output ecs.ListTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) ListTasksAsync(ctx workflow.Context, input *ecs.ListTasksInput) *EcsListTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTasks, input)
    return &EcsListTasksResult{Result: future}
}

func (a *ECSStub) PutAccountSetting(ctx workflow.Context, input *ecs.PutAccountSettingInput) (*ecs.PutAccountSettingOutput, error) {
    var output ecs.PutAccountSettingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountSetting, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) PutAccountSettingAsync(ctx workflow.Context, input *ecs.PutAccountSettingInput) *EcsPutAccountSettingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAccountSetting, input)
    return &EcsPutAccountSettingResult{Result: future}
}

func (a *ECSStub) PutAccountSettingDefault(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) (*ecs.PutAccountSettingDefaultOutput, error) {
    var output ecs.PutAccountSettingDefaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountSettingDefault, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) PutAccountSettingDefaultAsync(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) *EcsPutAccountSettingDefaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAccountSettingDefault, input)
    return &EcsPutAccountSettingDefaultResult{Result: future}
}

func (a *ECSStub) PutAttributes(ctx workflow.Context, input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error) {
    var output ecs.PutAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) PutAttributesAsync(ctx workflow.Context, input *ecs.PutAttributesInput) *EcsPutAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAttributes, input)
    return &EcsPutAttributesResult{Result: future}
}

func (a *ECSStub) PutClusterCapacityProviders(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) (*ecs.PutClusterCapacityProvidersOutput, error) {
    var output ecs.PutClusterCapacityProvidersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutClusterCapacityProviders, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) PutClusterCapacityProvidersAsync(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) *EcsPutClusterCapacityProvidersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutClusterCapacityProviders, input)
    return &EcsPutClusterCapacityProvidersResult{Result: future}
}

func (a *ECSStub) RegisterContainerInstance(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
    var output ecs.RegisterContainerInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterContainerInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) RegisterContainerInstanceAsync(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) *EcsRegisterContainerInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterContainerInstance, input)
    return &EcsRegisterContainerInstanceResult{Result: future}
}

func (a *ECSStub) RegisterTaskDefinition(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
    var output ecs.RegisterTaskDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterTaskDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) RegisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) *EcsRegisterTaskDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterTaskDefinition, input)
    return &EcsRegisterTaskDefinitionResult{Result: future}
}

func (a *ECSStub) RunTask(ctx workflow.Context, input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
    var output ecs.RunTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RunTask, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) RunTaskAsync(ctx workflow.Context, input *ecs.RunTaskInput) *EcsRunTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RunTask, input)
    return &EcsRunTaskResult{Result: future}
}

func (a *ECSStub) StartTask(ctx workflow.Context, input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error) {
    var output ecs.StartTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartTask, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) StartTaskAsync(ctx workflow.Context, input *ecs.StartTaskInput) *EcsStartTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartTask, input)
    return &EcsStartTaskResult{Result: future}
}

func (a *ECSStub) StopTask(ctx workflow.Context, input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
    var output ecs.StopTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopTask, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) StopTaskAsync(ctx workflow.Context, input *ecs.StopTaskInput) *EcsStopTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopTask, input)
    return &EcsStopTaskResult{Result: future}
}

func (a *ECSStub) SubmitAttachmentStateChanges(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) (*ecs.SubmitAttachmentStateChangesOutput, error) {
    var output ecs.SubmitAttachmentStateChangesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubmitAttachmentStateChanges, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) SubmitAttachmentStateChangesAsync(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) *EcsSubmitAttachmentStateChangesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SubmitAttachmentStateChanges, input)
    return &EcsSubmitAttachmentStateChangesResult{Result: future}
}

func (a *ECSStub) SubmitContainerStateChange(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
    var output ecs.SubmitContainerStateChangeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubmitContainerStateChange, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) SubmitContainerStateChangeAsync(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) *EcsSubmitContainerStateChangeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SubmitContainerStateChange, input)
    return &EcsSubmitContainerStateChangeResult{Result: future}
}

func (a *ECSStub) SubmitTaskStateChange(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
    var output ecs.SubmitTaskStateChangeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubmitTaskStateChange, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) SubmitTaskStateChangeAsync(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) *EcsSubmitTaskStateChangeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SubmitTaskStateChange, input)
    return &EcsSubmitTaskStateChangeResult{Result: future}
}

func (a *ECSStub) TagResource(ctx workflow.Context, input *ecs.TagResourceInput) (*ecs.TagResourceOutput, error) {
    var output ecs.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) TagResourceAsync(ctx workflow.Context, input *ecs.TagResourceInput) *EcsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &EcsTagResourceResult{Result: future}
}

func (a *ECSStub) UntagResource(ctx workflow.Context, input *ecs.UntagResourceInput) (*ecs.UntagResourceOutput, error) {
    var output ecs.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) UntagResourceAsync(ctx workflow.Context, input *ecs.UntagResourceInput) *EcsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &EcsUntagResourceResult{Result: future}
}

func (a *ECSStub) UpdateClusterSettings(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) (*ecs.UpdateClusterSettingsOutput, error) {
    var output ecs.UpdateClusterSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateClusterSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) UpdateClusterSettingsAsync(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) *EcsUpdateClusterSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateClusterSettings, input)
    return &EcsUpdateClusterSettingsResult{Result: future}
}

func (a *ECSStub) UpdateContainerAgent(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error) {
    var output ecs.UpdateContainerAgentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateContainerAgent, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) UpdateContainerAgentAsync(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) *EcsUpdateContainerAgentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateContainerAgent, input)
    return &EcsUpdateContainerAgentResult{Result: future}
}

func (a *ECSStub) UpdateContainerInstancesState(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error) {
    var output ecs.UpdateContainerInstancesStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateContainerInstancesState, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) UpdateContainerInstancesStateAsync(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) *EcsUpdateContainerInstancesStateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateContainerInstancesState, input)
    return &EcsUpdateContainerInstancesStateResult{Result: future}
}

func (a *ECSStub) UpdateService(ctx workflow.Context, input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
    var output ecs.UpdateServiceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateService, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) UpdateServiceAsync(ctx workflow.Context, input *ecs.UpdateServiceInput) *EcsUpdateServiceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateService, input)
    return &EcsUpdateServiceResult{Result: future}
}

func (a *ECSStub) UpdateServicePrimaryTaskSet(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) (*ecs.UpdateServicePrimaryTaskSetOutput, error) {
    var output ecs.UpdateServicePrimaryTaskSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServicePrimaryTaskSet, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) UpdateServicePrimaryTaskSetAsync(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) *EcsUpdateServicePrimaryTaskSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateServicePrimaryTaskSet, input)
    return &EcsUpdateServicePrimaryTaskSetResult{Result: future}
}

func (a *ECSStub) UpdateTaskSet(ctx workflow.Context, input *ecs.UpdateTaskSetInput) (*ecs.UpdateTaskSetOutput, error) {
    var output ecs.UpdateTaskSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTaskSet, input).Get(ctx, &output)
    return &output, err
}

func (a *ECSStub) UpdateTaskSetAsync(ctx workflow.Context, input *ecs.UpdateTaskSetInput) *EcsUpdateTaskSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTaskSet, input)
    return &EcsUpdateTaskSetResult{Result: future}
}

func (a *ECSStub) WaitUntilServicesInactive(ctx workflow.Context, input *ecs.DescribeServicesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilServicesInactive, input).Get(ctx, nil)
}

func (a *ECSStub) WaitUntilServicesInactiveAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilServicesInactive, input)
}


func (a *ECSStub) WaitUntilServicesStable(ctx workflow.Context, input *ecs.DescribeServicesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilServicesStable, input).Get(ctx, nil)
}

func (a *ECSStub) WaitUntilServicesStableAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilServicesStable, input)
}


func (a *ECSStub) WaitUntilTasksRunning(ctx workflow.Context, input *ecs.DescribeTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTasksRunning, input).Get(ctx, nil)
}

func (a *ECSStub) WaitUntilTasksRunningAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTasksRunning, input)
}


func (a *ECSStub) WaitUntilTasksStopped(ctx workflow.Context, input *ecs.DescribeTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTasksStopped, input).Get(ctx, nil)
}

func (a *ECSStub) WaitUntilTasksStoppedAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTasksStopped, input)
}

