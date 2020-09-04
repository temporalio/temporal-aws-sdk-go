
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/emr/emriface"
)

type EMRActivities struct {
	client emriface.EMRAPI
}

func NewEMRActivities(client emriface.EMRAPI) *EMRActivities {
    return &EMRActivities{client: client}
}

func (a *EMRActivities) AddInstanceFleet(input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error) {
    return a.client.AddInstanceFleet(input)
}

func (a *EMRActivities) AddInstanceGroups(input *emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error) {
    return a.client.AddInstanceGroups(input)
}

func (a *EMRActivities) AddJobFlowSteps(input *emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error) {
    return a.client.AddJobFlowSteps(input)
}

func (a *EMRActivities) AddTags(input *emr.AddTagsInput) (*emr.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *EMRActivities) CancelSteps(input *emr.CancelStepsInput) (*emr.CancelStepsOutput, error) {
    return a.client.CancelSteps(input)
}

func (a *EMRActivities) CreateSecurityConfiguration(input *emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error) {
    return a.client.CreateSecurityConfiguration(input)
}

func (a *EMRActivities) DeleteSecurityConfiguration(input *emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error) {
    return a.client.DeleteSecurityConfiguration(input)
}

func (a *EMRActivities) DescribeCluster(input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error) {
    return a.client.DescribeCluster(input)
}

func (a *EMRActivities) DescribeJobFlows(input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error) {
    return a.client.DescribeJobFlows(input)
}

func (a *EMRActivities) DescribeNotebookExecution(input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error) {
    return a.client.DescribeNotebookExecution(input)
}

func (a *EMRActivities) DescribeSecurityConfiguration(input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error) {
    return a.client.DescribeSecurityConfiguration(input)
}

func (a *EMRActivities) DescribeStep(input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error) {
    return a.client.DescribeStep(input)
}

func (a *EMRActivities) GetBlockPublicAccessConfiguration(input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
    return a.client.GetBlockPublicAccessConfiguration(input)
}

func (a *EMRActivities) GetManagedScalingPolicy(input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error) {
    return a.client.GetManagedScalingPolicy(input)
}

func (a *EMRActivities) ListBootstrapActions(input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error) {
    return a.client.ListBootstrapActions(input)
}

func (a *EMRActivities) ListClusters(input *emr.ListClustersInput) (*emr.ListClustersOutput, error) {
    return a.client.ListClusters(input)
}

func (a *EMRActivities) ListInstanceFleets(input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error) {
    return a.client.ListInstanceFleets(input)
}

func (a *EMRActivities) ListInstanceGroups(input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error) {
    return a.client.ListInstanceGroups(input)
}

func (a *EMRActivities) ListInstances(input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error) {
    return a.client.ListInstances(input)
}

func (a *EMRActivities) ListNotebookExecutions(input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error) {
    return a.client.ListNotebookExecutions(input)
}

func (a *EMRActivities) ListSecurityConfigurations(input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error) {
    return a.client.ListSecurityConfigurations(input)
}

func (a *EMRActivities) ListSteps(input *emr.ListStepsInput) (*emr.ListStepsOutput, error) {
    return a.client.ListSteps(input)
}

func (a *EMRActivities) ModifyCluster(input *emr.ModifyClusterInput) (*emr.ModifyClusterOutput, error) {
    return a.client.ModifyCluster(input)
}

func (a *EMRActivities) ModifyInstanceFleet(input *emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error) {
    return a.client.ModifyInstanceFleet(input)
}

func (a *EMRActivities) ModifyInstanceGroups(input *emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error) {
    return a.client.ModifyInstanceGroups(input)
}

func (a *EMRActivities) PutAutoScalingPolicy(input *emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error) {
    return a.client.PutAutoScalingPolicy(input)
}

func (a *EMRActivities) PutBlockPublicAccessConfiguration(input *emr.PutBlockPublicAccessConfigurationInput) (*emr.PutBlockPublicAccessConfigurationOutput, error) {
    return a.client.PutBlockPublicAccessConfiguration(input)
}

func (a *EMRActivities) PutManagedScalingPolicy(input *emr.PutManagedScalingPolicyInput) (*emr.PutManagedScalingPolicyOutput, error) {
    return a.client.PutManagedScalingPolicy(input)
}

func (a *EMRActivities) RemoveAutoScalingPolicy(input *emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error) {
    return a.client.RemoveAutoScalingPolicy(input)
}

func (a *EMRActivities) RemoveManagedScalingPolicy(input *emr.RemoveManagedScalingPolicyInput) (*emr.RemoveManagedScalingPolicyOutput, error) {
    return a.client.RemoveManagedScalingPolicy(input)
}

func (a *EMRActivities) RemoveTags(input *emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error) {
    return a.client.RemoveTags(input)
}

func (a *EMRActivities) RunJobFlow(input *emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error) {
    return a.client.RunJobFlow(input)
}

func (a *EMRActivities) SetTerminationProtection(input *emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error) {
    return a.client.SetTerminationProtection(input)
}

func (a *EMRActivities) SetVisibleToAllUsers(input *emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error) {
    return a.client.SetVisibleToAllUsers(input)
}

func (a *EMRActivities) StartNotebookExecution(input *emr.StartNotebookExecutionInput) (*emr.StartNotebookExecutionOutput, error) {
    return a.client.StartNotebookExecution(input)
}

func (a *EMRActivities) StopNotebookExecution(input *emr.StopNotebookExecutionInput) (*emr.StopNotebookExecutionOutput, error) {
    return a.client.StopNotebookExecution(input)
}

func (a *EMRActivities) TerminateJobFlows(input *emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error) {
    return a.client.TerminateJobFlows(input)
}

func (a *EMRActivities) WaitUntilClusterRunning(input *emr.DescribeClusterInput) error {
    return a.client.WaitUntilClusterRunning(input)
}

func (a *EMRActivities) WaitUntilClusterTerminated(input *emr.DescribeClusterInput) error {
    return a.client.WaitUntilClusterTerminated(input)
}

func (a *EMRActivities) WaitUntilStepComplete(input *emr.DescribeStepInput) error {
    return a.client.WaitUntilStepComplete(input)
}
