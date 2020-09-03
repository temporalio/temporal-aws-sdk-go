package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/emr"
	"go.temporal.io/sdk/workflow"
)

type EMRClient interface {
    AddInstanceFleet(ctx workflow.Context, input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error)
    AddInstanceFleetAsync(ctx workflow.Context, input *emr.AddInstanceFleetInput) *EmrAddInstanceFleetResult

    AddInstanceGroups(ctx workflow.Context, input *emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error)
    AddInstanceGroupsAsync(ctx workflow.Context, input *emr.AddInstanceGroupsInput) *EmrAddInstanceGroupsResult

    AddJobFlowSteps(ctx workflow.Context, input *emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error)
    AddJobFlowStepsAsync(ctx workflow.Context, input *emr.AddJobFlowStepsInput) *EmrAddJobFlowStepsResult

    AddTags(ctx workflow.Context, input *emr.AddTagsInput) (*emr.AddTagsOutput, error)
    AddTagsAsync(ctx workflow.Context, input *emr.AddTagsInput) *EmrAddTagsResult

    CancelSteps(ctx workflow.Context, input *emr.CancelStepsInput) (*emr.CancelStepsOutput, error)
    CancelStepsAsync(ctx workflow.Context, input *emr.CancelStepsInput) *EmrCancelStepsResult

    CreateSecurityConfiguration(ctx workflow.Context, input *emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error)
    CreateSecurityConfigurationAsync(ctx workflow.Context, input *emr.CreateSecurityConfigurationInput) *EmrCreateSecurityConfigurationResult

    DeleteSecurityConfiguration(ctx workflow.Context, input *emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error)
    DeleteSecurityConfigurationAsync(ctx workflow.Context, input *emr.DeleteSecurityConfigurationInput) *EmrDeleteSecurityConfigurationResult

    DescribeCluster(ctx workflow.Context, input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error)
    DescribeClusterAsync(ctx workflow.Context, input *emr.DescribeClusterInput) *EmrDescribeClusterResult

    DescribeJobFlows(ctx workflow.Context, input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error)
    DescribeJobFlowsAsync(ctx workflow.Context, input *emr.DescribeJobFlowsInput) *EmrDescribeJobFlowsResult

    DescribeNotebookExecution(ctx workflow.Context, input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error)
    DescribeNotebookExecutionAsync(ctx workflow.Context, input *emr.DescribeNotebookExecutionInput) *EmrDescribeNotebookExecutionResult

    DescribeSecurityConfiguration(ctx workflow.Context, input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error)
    DescribeSecurityConfigurationAsync(ctx workflow.Context, input *emr.DescribeSecurityConfigurationInput) *EmrDescribeSecurityConfigurationResult

    DescribeStep(ctx workflow.Context, input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error)
    DescribeStepAsync(ctx workflow.Context, input *emr.DescribeStepInput) *EmrDescribeStepResult

    GetBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error)
    GetBlockPublicAccessConfigurationAsync(ctx workflow.Context, input *emr.GetBlockPublicAccessConfigurationInput) *EmrGetBlockPublicAccessConfigurationResult

    GetManagedScalingPolicy(ctx workflow.Context, input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error)
    GetManagedScalingPolicyAsync(ctx workflow.Context, input *emr.GetManagedScalingPolicyInput) *EmrGetManagedScalingPolicyResult

    ListBootstrapActions(ctx workflow.Context, input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error)
    ListBootstrapActionsAsync(ctx workflow.Context, input *emr.ListBootstrapActionsInput) *EmrListBootstrapActionsResult

    ListClusters(ctx workflow.Context, input *emr.ListClustersInput) (*emr.ListClustersOutput, error)
    ListClustersAsync(ctx workflow.Context, input *emr.ListClustersInput) *EmrListClustersResult

    ListInstanceFleets(ctx workflow.Context, input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error)
    ListInstanceFleetsAsync(ctx workflow.Context, input *emr.ListInstanceFleetsInput) *EmrListInstanceFleetsResult

    ListInstanceGroups(ctx workflow.Context, input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error)
    ListInstanceGroupsAsync(ctx workflow.Context, input *emr.ListInstanceGroupsInput) *EmrListInstanceGroupsResult

    ListInstances(ctx workflow.Context, input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error)
    ListInstancesAsync(ctx workflow.Context, input *emr.ListInstancesInput) *EmrListInstancesResult

    ListNotebookExecutions(ctx workflow.Context, input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error)
    ListNotebookExecutionsAsync(ctx workflow.Context, input *emr.ListNotebookExecutionsInput) *EmrListNotebookExecutionsResult

    ListSecurityConfigurations(ctx workflow.Context, input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error)
    ListSecurityConfigurationsAsync(ctx workflow.Context, input *emr.ListSecurityConfigurationsInput) *EmrListSecurityConfigurationsResult

    ListSteps(ctx workflow.Context, input *emr.ListStepsInput) (*emr.ListStepsOutput, error)
    ListStepsAsync(ctx workflow.Context, input *emr.ListStepsInput) *EmrListStepsResult

    ModifyCluster(ctx workflow.Context, input *emr.ModifyClusterInput) (*emr.ModifyClusterOutput, error)
    ModifyClusterAsync(ctx workflow.Context, input *emr.ModifyClusterInput) *EmrModifyClusterResult

    ModifyInstanceFleet(ctx workflow.Context, input *emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error)
    ModifyInstanceFleetAsync(ctx workflow.Context, input *emr.ModifyInstanceFleetInput) *EmrModifyInstanceFleetResult

    ModifyInstanceGroups(ctx workflow.Context, input *emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error)
    ModifyInstanceGroupsAsync(ctx workflow.Context, input *emr.ModifyInstanceGroupsInput) *EmrModifyInstanceGroupsResult

    PutAutoScalingPolicy(ctx workflow.Context, input *emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error)
    PutAutoScalingPolicyAsync(ctx workflow.Context, input *emr.PutAutoScalingPolicyInput) *EmrPutAutoScalingPolicyResult

    PutBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.PutBlockPublicAccessConfigurationInput) (*emr.PutBlockPublicAccessConfigurationOutput, error)
    PutBlockPublicAccessConfigurationAsync(ctx workflow.Context, input *emr.PutBlockPublicAccessConfigurationInput) *EmrPutBlockPublicAccessConfigurationResult

    PutManagedScalingPolicy(ctx workflow.Context, input *emr.PutManagedScalingPolicyInput) (*emr.PutManagedScalingPolicyOutput, error)
    PutManagedScalingPolicyAsync(ctx workflow.Context, input *emr.PutManagedScalingPolicyInput) *EmrPutManagedScalingPolicyResult

    RemoveAutoScalingPolicy(ctx workflow.Context, input *emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error)
    RemoveAutoScalingPolicyAsync(ctx workflow.Context, input *emr.RemoveAutoScalingPolicyInput) *EmrRemoveAutoScalingPolicyResult

    RemoveManagedScalingPolicy(ctx workflow.Context, input *emr.RemoveManagedScalingPolicyInput) (*emr.RemoveManagedScalingPolicyOutput, error)
    RemoveManagedScalingPolicyAsync(ctx workflow.Context, input *emr.RemoveManagedScalingPolicyInput) *EmrRemoveManagedScalingPolicyResult

    RemoveTags(ctx workflow.Context, input *emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error)
    RemoveTagsAsync(ctx workflow.Context, input *emr.RemoveTagsInput) *EmrRemoveTagsResult

    RunJobFlow(ctx workflow.Context, input *emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error)
    RunJobFlowAsync(ctx workflow.Context, input *emr.RunJobFlowInput) *EmrRunJobFlowResult

    SetTerminationProtection(ctx workflow.Context, input *emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error)
    SetTerminationProtectionAsync(ctx workflow.Context, input *emr.SetTerminationProtectionInput) *EmrSetTerminationProtectionResult

    SetVisibleToAllUsers(ctx workflow.Context, input *emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error)
    SetVisibleToAllUsersAsync(ctx workflow.Context, input *emr.SetVisibleToAllUsersInput) *EmrSetVisibleToAllUsersResult

    StartNotebookExecution(ctx workflow.Context, input *emr.StartNotebookExecutionInput) (*emr.StartNotebookExecutionOutput, error)
    StartNotebookExecutionAsync(ctx workflow.Context, input *emr.StartNotebookExecutionInput) *EmrStartNotebookExecutionResult

    StopNotebookExecution(ctx workflow.Context, input *emr.StopNotebookExecutionInput) (*emr.StopNotebookExecutionOutput, error)
    StopNotebookExecutionAsync(ctx workflow.Context, input *emr.StopNotebookExecutionInput) *EmrStopNotebookExecutionResult

    TerminateJobFlows(ctx workflow.Context, input *emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error)
    TerminateJobFlowsAsync(ctx workflow.Context, input *emr.TerminateJobFlowsInput) *EmrTerminateJobFlowsResult

    WaitUntilClusterRunning(ctx workflow.Context, input *emr.DescribeClusterInput) error
    WaitUntilClusterTerminated(ctx workflow.Context, input *emr.DescribeClusterInput) error
    WaitUntilStepComplete(ctx workflow.Context, input *emr.DescribeStepInput) error}
type EmrAddInstanceFleetResult struct {
	Result workflow.Future
}

func (r *EmrAddInstanceFleetResult) Get(ctx workflow.Context) (*emr.AddInstanceFleetOutput, error) {
    var output emr.AddInstanceFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrAddInstanceGroupsResult struct {
	Result workflow.Future
}

func (r *EmrAddInstanceGroupsResult) Get(ctx workflow.Context) (*emr.AddInstanceGroupsOutput, error) {
    var output emr.AddInstanceGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrAddJobFlowStepsResult struct {
	Result workflow.Future
}

func (r *EmrAddJobFlowStepsResult) Get(ctx workflow.Context) (*emr.AddJobFlowStepsOutput, error) {
    var output emr.AddJobFlowStepsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrAddTagsResult struct {
	Result workflow.Future
}

func (r *EmrAddTagsResult) Get(ctx workflow.Context) (*emr.AddTagsOutput, error) {
    var output emr.AddTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrCancelStepsResult struct {
	Result workflow.Future
}

func (r *EmrCancelStepsResult) Get(ctx workflow.Context) (*emr.CancelStepsOutput, error) {
    var output emr.CancelStepsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrCreateSecurityConfigurationResult struct {
	Result workflow.Future
}

func (r *EmrCreateSecurityConfigurationResult) Get(ctx workflow.Context) (*emr.CreateSecurityConfigurationOutput, error) {
    var output emr.CreateSecurityConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrDeleteSecurityConfigurationResult struct {
	Result workflow.Future
}

func (r *EmrDeleteSecurityConfigurationResult) Get(ctx workflow.Context) (*emr.DeleteSecurityConfigurationOutput, error) {
    var output emr.DeleteSecurityConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrDescribeClusterResult struct {
	Result workflow.Future
}

func (r *EmrDescribeClusterResult) Get(ctx workflow.Context) (*emr.DescribeClusterOutput, error) {
    var output emr.DescribeClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrDescribeJobFlowsResult struct {
	Result workflow.Future
}

func (r *EmrDescribeJobFlowsResult) Get(ctx workflow.Context) (*emr.DescribeJobFlowsOutput, error) {
    var output emr.DescribeJobFlowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrDescribeNotebookExecutionResult struct {
	Result workflow.Future
}

func (r *EmrDescribeNotebookExecutionResult) Get(ctx workflow.Context) (*emr.DescribeNotebookExecutionOutput, error) {
    var output emr.DescribeNotebookExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrDescribeSecurityConfigurationResult struct {
	Result workflow.Future
}

func (r *EmrDescribeSecurityConfigurationResult) Get(ctx workflow.Context) (*emr.DescribeSecurityConfigurationOutput, error) {
    var output emr.DescribeSecurityConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrDescribeStepResult struct {
	Result workflow.Future
}

func (r *EmrDescribeStepResult) Get(ctx workflow.Context) (*emr.DescribeStepOutput, error) {
    var output emr.DescribeStepOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrGetBlockPublicAccessConfigurationResult struct {
	Result workflow.Future
}

func (r *EmrGetBlockPublicAccessConfigurationResult) Get(ctx workflow.Context) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
    var output emr.GetBlockPublicAccessConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrGetManagedScalingPolicyResult struct {
	Result workflow.Future
}

func (r *EmrGetManagedScalingPolicyResult) Get(ctx workflow.Context) (*emr.GetManagedScalingPolicyOutput, error) {
    var output emr.GetManagedScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListBootstrapActionsResult struct {
	Result workflow.Future
}

func (r *EmrListBootstrapActionsResult) Get(ctx workflow.Context) (*emr.ListBootstrapActionsOutput, error) {
    var output emr.ListBootstrapActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListClustersResult struct {
	Result workflow.Future
}

func (r *EmrListClustersResult) Get(ctx workflow.Context) (*emr.ListClustersOutput, error) {
    var output emr.ListClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListInstanceFleetsResult struct {
	Result workflow.Future
}

func (r *EmrListInstanceFleetsResult) Get(ctx workflow.Context) (*emr.ListInstanceFleetsOutput, error) {
    var output emr.ListInstanceFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListInstanceGroupsResult struct {
	Result workflow.Future
}

func (r *EmrListInstanceGroupsResult) Get(ctx workflow.Context) (*emr.ListInstanceGroupsOutput, error) {
    var output emr.ListInstanceGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListInstancesResult struct {
	Result workflow.Future
}

func (r *EmrListInstancesResult) Get(ctx workflow.Context) (*emr.ListInstancesOutput, error) {
    var output emr.ListInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListNotebookExecutionsResult struct {
	Result workflow.Future
}

func (r *EmrListNotebookExecutionsResult) Get(ctx workflow.Context) (*emr.ListNotebookExecutionsOutput, error) {
    var output emr.ListNotebookExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListSecurityConfigurationsResult struct {
	Result workflow.Future
}

func (r *EmrListSecurityConfigurationsResult) Get(ctx workflow.Context) (*emr.ListSecurityConfigurationsOutput, error) {
    var output emr.ListSecurityConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrListStepsResult struct {
	Result workflow.Future
}

func (r *EmrListStepsResult) Get(ctx workflow.Context) (*emr.ListStepsOutput, error) {
    var output emr.ListStepsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrModifyClusterResult struct {
	Result workflow.Future
}

func (r *EmrModifyClusterResult) Get(ctx workflow.Context) (*emr.ModifyClusterOutput, error) {
    var output emr.ModifyClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrModifyInstanceFleetResult struct {
	Result workflow.Future
}

func (r *EmrModifyInstanceFleetResult) Get(ctx workflow.Context) (*emr.ModifyInstanceFleetOutput, error) {
    var output emr.ModifyInstanceFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrModifyInstanceGroupsResult struct {
	Result workflow.Future
}

func (r *EmrModifyInstanceGroupsResult) Get(ctx workflow.Context) (*emr.ModifyInstanceGroupsOutput, error) {
    var output emr.ModifyInstanceGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrPutAutoScalingPolicyResult struct {
	Result workflow.Future
}

func (r *EmrPutAutoScalingPolicyResult) Get(ctx workflow.Context) (*emr.PutAutoScalingPolicyOutput, error) {
    var output emr.PutAutoScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrPutBlockPublicAccessConfigurationResult struct {
	Result workflow.Future
}

func (r *EmrPutBlockPublicAccessConfigurationResult) Get(ctx workflow.Context) (*emr.PutBlockPublicAccessConfigurationOutput, error) {
    var output emr.PutBlockPublicAccessConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrPutManagedScalingPolicyResult struct {
	Result workflow.Future
}

func (r *EmrPutManagedScalingPolicyResult) Get(ctx workflow.Context) (*emr.PutManagedScalingPolicyOutput, error) {
    var output emr.PutManagedScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrRemoveAutoScalingPolicyResult struct {
	Result workflow.Future
}

func (r *EmrRemoveAutoScalingPolicyResult) Get(ctx workflow.Context) (*emr.RemoveAutoScalingPolicyOutput, error) {
    var output emr.RemoveAutoScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrRemoveManagedScalingPolicyResult struct {
	Result workflow.Future
}

func (r *EmrRemoveManagedScalingPolicyResult) Get(ctx workflow.Context) (*emr.RemoveManagedScalingPolicyOutput, error) {
    var output emr.RemoveManagedScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrRemoveTagsResult struct {
	Result workflow.Future
}

func (r *EmrRemoveTagsResult) Get(ctx workflow.Context) (*emr.RemoveTagsOutput, error) {
    var output emr.RemoveTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrRunJobFlowResult struct {
	Result workflow.Future
}

func (r *EmrRunJobFlowResult) Get(ctx workflow.Context) (*emr.RunJobFlowOutput, error) {
    var output emr.RunJobFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrSetTerminationProtectionResult struct {
	Result workflow.Future
}

func (r *EmrSetTerminationProtectionResult) Get(ctx workflow.Context) (*emr.SetTerminationProtectionOutput, error) {
    var output emr.SetTerminationProtectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrSetVisibleToAllUsersResult struct {
	Result workflow.Future
}

func (r *EmrSetVisibleToAllUsersResult) Get(ctx workflow.Context) (*emr.SetVisibleToAllUsersOutput, error) {
    var output emr.SetVisibleToAllUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrStartNotebookExecutionResult struct {
	Result workflow.Future
}

func (r *EmrStartNotebookExecutionResult) Get(ctx workflow.Context) (*emr.StartNotebookExecutionOutput, error) {
    var output emr.StartNotebookExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrStopNotebookExecutionResult struct {
	Result workflow.Future
}

func (r *EmrStopNotebookExecutionResult) Get(ctx workflow.Context) (*emr.StopNotebookExecutionOutput, error) {
    var output emr.StopNotebookExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EmrTerminateJobFlowsResult struct {
	Result workflow.Future
}

func (r *EmrTerminateJobFlowsResult) Get(ctx workflow.Context) (*emr.TerminateJobFlowsOutput, error) {
    var output emr.TerminateJobFlowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type EMRStub struct {
    activities EMRClient
}

func NewEMRStub() EMRClient {
    return &EMRStub{}
}

func (a *EMRStub) AddInstanceFleet(ctx workflow.Context, input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error) {
    var output emr.AddInstanceFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddInstanceFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) AddInstanceGroups(ctx workflow.Context, input *emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error) {
    var output emr.AddInstanceGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddInstanceGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) AddJobFlowSteps(ctx workflow.Context, input *emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error) {
    var output emr.AddJobFlowStepsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddJobFlowSteps, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) AddTags(ctx workflow.Context, input *emr.AddTagsInput) (*emr.AddTagsOutput, error) {
    var output emr.AddTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) CancelSteps(ctx workflow.Context, input *emr.CancelStepsInput) (*emr.CancelStepsOutput, error) {
    var output emr.CancelStepsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelSteps, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) CreateSecurityConfiguration(ctx workflow.Context, input *emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error) {
    var output emr.CreateSecurityConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSecurityConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) DeleteSecurityConfiguration(ctx workflow.Context, input *emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error) {
    var output emr.DeleteSecurityConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSecurityConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) DescribeCluster(ctx workflow.Context, input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error) {
    var output emr.DescribeClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) DescribeJobFlows(ctx workflow.Context, input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error) {
    var output emr.DescribeJobFlowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJobFlows, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) DescribeNotebookExecution(ctx workflow.Context, input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error) {
    var output emr.DescribeNotebookExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNotebookExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) DescribeSecurityConfiguration(ctx workflow.Context, input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error) {
    var output emr.DescribeSecurityConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) DescribeStep(ctx workflow.Context, input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error) {
    var output emr.DescribeStepOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStep, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) GetBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
    var output emr.GetBlockPublicAccessConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBlockPublicAccessConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) GetManagedScalingPolicy(ctx workflow.Context, input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error) {
    var output emr.GetManagedScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetManagedScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListBootstrapActions(ctx workflow.Context, input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error) {
    var output emr.ListBootstrapActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBootstrapActions, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListClusters(ctx workflow.Context, input *emr.ListClustersInput) (*emr.ListClustersOutput, error) {
    var output emr.ListClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListInstanceFleets(ctx workflow.Context, input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error) {
    var output emr.ListInstanceFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInstanceFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListInstanceGroups(ctx workflow.Context, input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error) {
    var output emr.ListInstanceGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInstanceGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListInstances(ctx workflow.Context, input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error) {
    var output emr.ListInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListNotebookExecutions(ctx workflow.Context, input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error) {
    var output emr.ListNotebookExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNotebookExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListSecurityConfigurations(ctx workflow.Context, input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error) {
    var output emr.ListSecurityConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSecurityConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ListSteps(ctx workflow.Context, input *emr.ListStepsInput) (*emr.ListStepsOutput, error) {
    var output emr.ListStepsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSteps, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ModifyCluster(ctx workflow.Context, input *emr.ModifyClusterInput) (*emr.ModifyClusterOutput, error) {
    var output emr.ModifyClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ModifyInstanceFleet(ctx workflow.Context, input *emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error) {
    var output emr.ModifyInstanceFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) ModifyInstanceGroups(ctx workflow.Context, input *emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error) {
    var output emr.ModifyInstanceGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) PutAutoScalingPolicy(ctx workflow.Context, input *emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error) {
    var output emr.PutAutoScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAutoScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) PutBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.PutBlockPublicAccessConfigurationInput) (*emr.PutBlockPublicAccessConfigurationOutput, error) {
    var output emr.PutBlockPublicAccessConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutBlockPublicAccessConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) PutManagedScalingPolicy(ctx workflow.Context, input *emr.PutManagedScalingPolicyInput) (*emr.PutManagedScalingPolicyOutput, error) {
    var output emr.PutManagedScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutManagedScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) RemoveAutoScalingPolicy(ctx workflow.Context, input *emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error) {
    var output emr.RemoveAutoScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveAutoScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) RemoveManagedScalingPolicy(ctx workflow.Context, input *emr.RemoveManagedScalingPolicyInput) (*emr.RemoveManagedScalingPolicyOutput, error) {
    var output emr.RemoveManagedScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveManagedScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) RemoveTags(ctx workflow.Context, input *emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error) {
    var output emr.RemoveTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) RunJobFlow(ctx workflow.Context, input *emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error) {
    var output emr.RunJobFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RunJobFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) SetTerminationProtection(ctx workflow.Context, input *emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error) {
    var output emr.SetTerminationProtectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetTerminationProtection, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) SetVisibleToAllUsers(ctx workflow.Context, input *emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error) {
    var output emr.SetVisibleToAllUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetVisibleToAllUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) StartNotebookExecution(ctx workflow.Context, input *emr.StartNotebookExecutionInput) (*emr.StartNotebookExecutionOutput, error) {
    var output emr.StartNotebookExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartNotebookExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) StopNotebookExecution(ctx workflow.Context, input *emr.StopNotebookExecutionInput) (*emr.StopNotebookExecutionOutput, error) {
    var output emr.StopNotebookExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopNotebookExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *EMRStub) TerminateJobFlows(ctx workflow.Context, input *emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error) {
    var output emr.TerminateJobFlowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateJobFlows, input).Get(ctx, &output)
    return &output, err
}


func (a *EMRStub) WaitUntilClusterRunning(ctx workflow.Context, input *emr.DescribeClusterInput) error {
    return a.client.WaitUntilClusterRunning(input)
}


func (a *EMRStub) WaitUntilClusterTerminated(ctx workflow.Context, input *emr.DescribeClusterInput) error {
    return a.client.WaitUntilClusterTerminated(input)
}


func (a *EMRStub) WaitUntilStepComplete(ctx workflow.Context, input *emr.DescribeStepInput) error {
    return a.client.WaitUntilStepComplete(input)
}
