
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

type ECSActivities struct {
    client ecsiface.ECSAPI
}

func NewECSActivities(session *session.Session, config ...*aws.Config) *ECSActivities {
    client := ecs.New(session, config...)
    return &ECSActivities{client: client}
}

func (a *ECSActivities) CreateCapacityProvider(input *ecs.CreateCapacityProviderInput) (*ecs.CreateCapacityProviderOutput, error) {
    return a.client.CreateCapacityProvider(input)
}

func (a *ECSActivities) CreateCluster(input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
    return a.client.CreateCluster(input)
}

func (a *ECSActivities) CreateService(input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
    return a.client.CreateService(input)
}

func (a *ECSActivities) CreateTaskSet(input *ecs.CreateTaskSetInput) (*ecs.CreateTaskSetOutput, error) {
    return a.client.CreateTaskSet(input)
}

func (a *ECSActivities) DeleteAccountSetting(input *ecs.DeleteAccountSettingInput) (*ecs.DeleteAccountSettingOutput, error) {
    return a.client.DeleteAccountSetting(input)
}

func (a *ECSActivities) DeleteAttributes(input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error) {
    return a.client.DeleteAttributes(input)
}

func (a *ECSActivities) DeleteCapacityProvider(input *ecs.DeleteCapacityProviderInput) (*ecs.DeleteCapacityProviderOutput, error) {
    return a.client.DeleteCapacityProvider(input)
}

func (a *ECSActivities) DeleteCluster(input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
    return a.client.DeleteCluster(input)
}

func (a *ECSActivities) DeleteService(input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error) {
    return a.client.DeleteService(input)
}

func (a *ECSActivities) DeleteTaskSet(input *ecs.DeleteTaskSetInput) (*ecs.DeleteTaskSetOutput, error) {
    return a.client.DeleteTaskSet(input)
}

func (a *ECSActivities) DeregisterContainerInstance(input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error) {
    return a.client.DeregisterContainerInstance(input)
}

func (a *ECSActivities) DeregisterTaskDefinition(input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error) {
    return a.client.DeregisterTaskDefinition(input)
}

func (a *ECSActivities) DescribeCapacityProviders(input *ecs.DescribeCapacityProvidersInput) (*ecs.DescribeCapacityProvidersOutput, error) {
    return a.client.DescribeCapacityProviders(input)
}

func (a *ECSActivities) DescribeClusters(input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error) {
    return a.client.DescribeClusters(input)
}

func (a *ECSActivities) DescribeContainerInstances(input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
    return a.client.DescribeContainerInstances(input)
}

func (a *ECSActivities) DescribeServices(input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error) {
    return a.client.DescribeServices(input)
}

func (a *ECSActivities) DescribeTaskDefinition(input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
    return a.client.DescribeTaskDefinition(input)
}

func (a *ECSActivities) DescribeTaskSets(input *ecs.DescribeTaskSetsInput) (*ecs.DescribeTaskSetsOutput, error) {
    return a.client.DescribeTaskSets(input)
}

func (a *ECSActivities) DescribeTasks(input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
    return a.client.DescribeTasks(input)
}

func (a *ECSActivities) DiscoverPollEndpoint(input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
    return a.client.DiscoverPollEndpoint(input)
}

func (a *ECSActivities) ListAccountSettings(input *ecs.ListAccountSettingsInput) (*ecs.ListAccountSettingsOutput, error) {
    return a.client.ListAccountSettings(input)
}

func (a *ECSActivities) ListAttributes(input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error) {
    return a.client.ListAttributes(input)
}

func (a *ECSActivities) ListClusters(input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
    return a.client.ListClusters(input)
}

func (a *ECSActivities) ListContainerInstances(input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
    return a.client.ListContainerInstances(input)
}

func (a *ECSActivities) ListServices(input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
    return a.client.ListServices(input)
}

func (a *ECSActivities) ListTagsForResource(input *ecs.ListTagsForResourceInput) (*ecs.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ECSActivities) ListTaskDefinitionFamilies(input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
    return a.client.ListTaskDefinitionFamilies(input)
}

func (a *ECSActivities) ListTaskDefinitions(input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
    return a.client.ListTaskDefinitions(input)
}

func (a *ECSActivities) ListTasks(input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
    return a.client.ListTasks(input)
}

func (a *ECSActivities) PutAccountSetting(input *ecs.PutAccountSettingInput) (*ecs.PutAccountSettingOutput, error) {
    return a.client.PutAccountSetting(input)
}

func (a *ECSActivities) PutAccountSettingDefault(input *ecs.PutAccountSettingDefaultInput) (*ecs.PutAccountSettingDefaultOutput, error) {
    return a.client.PutAccountSettingDefault(input)
}

func (a *ECSActivities) PutAttributes(input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error) {
    return a.client.PutAttributes(input)
}

func (a *ECSActivities) PutClusterCapacityProviders(input *ecs.PutClusterCapacityProvidersInput) (*ecs.PutClusterCapacityProvidersOutput, error) {
    return a.client.PutClusterCapacityProviders(input)
}

func (a *ECSActivities) RegisterContainerInstance(input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
    return a.client.RegisterContainerInstance(input)
}

func (a *ECSActivities) RegisterTaskDefinition(input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
    return a.client.RegisterTaskDefinition(input)
}

func (a *ECSActivities) RunTask(input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
    return a.client.RunTask(input)
}

func (a *ECSActivities) StartTask(input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error) {
    return a.client.StartTask(input)
}

func (a *ECSActivities) StopTask(input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
    return a.client.StopTask(input)
}

func (a *ECSActivities) SubmitAttachmentStateChanges(input *ecs.SubmitAttachmentStateChangesInput) (*ecs.SubmitAttachmentStateChangesOutput, error) {
    return a.client.SubmitAttachmentStateChanges(input)
}

func (a *ECSActivities) SubmitContainerStateChange(input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
    return a.client.SubmitContainerStateChange(input)
}

func (a *ECSActivities) SubmitTaskStateChange(input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
    return a.client.SubmitTaskStateChange(input)
}

func (a *ECSActivities) TagResource(input *ecs.TagResourceInput) (*ecs.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ECSActivities) UntagResource(input *ecs.UntagResourceInput) (*ecs.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ECSActivities) UpdateClusterSettings(input *ecs.UpdateClusterSettingsInput) (*ecs.UpdateClusterSettingsOutput, error) {
    return a.client.UpdateClusterSettings(input)
}

func (a *ECSActivities) UpdateContainerAgent(input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error) {
    return a.client.UpdateContainerAgent(input)
}

func (a *ECSActivities) UpdateContainerInstancesState(input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error) {
    return a.client.UpdateContainerInstancesState(input)
}

func (a *ECSActivities) UpdateService(input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
    return a.client.UpdateService(input)
}

func (a *ECSActivities) UpdateServicePrimaryTaskSet(input *ecs.UpdateServicePrimaryTaskSetInput) (*ecs.UpdateServicePrimaryTaskSetOutput, error) {
    return a.client.UpdateServicePrimaryTaskSet(input)
}

func (a *ECSActivities) UpdateTaskSet(input *ecs.UpdateTaskSetInput) (*ecs.UpdateTaskSetOutput, error) {
    return a.client.UpdateTaskSet(input)
}

func (a *ECSActivities) WaitUntilServicesInactive(input *ecs.DescribeServicesInput) error {
    return a.client.WaitUntilServicesInactive(input)
}

func (a *ECSActivities) WaitUntilServicesStable(input *ecs.DescribeServicesInput) error {
    return a.client.WaitUntilServicesStable(input)
}

func (a *ECSActivities) WaitUntilTasksRunning(input *ecs.DescribeTasksInput) error {
    return a.client.WaitUntilTasksRunning(input)
}

func (a *ECSActivities) WaitUntilTasksStopped(input *ecs.DescribeTasksInput) error {
    return a.client.WaitUntilTasksStopped(input)
}
