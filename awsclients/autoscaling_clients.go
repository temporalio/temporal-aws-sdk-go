package awsclients

import (
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type AutoScalingClient interface {
    AttachInstances(ctx workflow.Context, input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error)
    AttachInstancesAsync(ctx workflow.Context, input *autoscaling.AttachInstancesInput) *AutoscalingAttachInstancesResult

    AttachLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error)
    AttachLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) *AutoscalingAttachLoadBalancerTargetGroupsResult

    AttachLoadBalancers(ctx workflow.Context, input *autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error)
    AttachLoadBalancersAsync(ctx workflow.Context, input *autoscaling.AttachLoadBalancersInput) *AutoscalingAttachLoadBalancersResult

    BatchDeleteScheduledAction(ctx workflow.Context, input *autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error)
    BatchDeleteScheduledActionAsync(ctx workflow.Context, input *autoscaling.BatchDeleteScheduledActionInput) *AutoscalingBatchDeleteScheduledActionResult

    BatchPutScheduledUpdateGroupAction(ctx workflow.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error)
    BatchPutScheduledUpdateGroupActionAsync(ctx workflow.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) *AutoscalingBatchPutScheduledUpdateGroupActionResult

    CancelInstanceRefresh(ctx workflow.Context, input *autoscaling.CancelInstanceRefreshInput) (*autoscaling.CancelInstanceRefreshOutput, error)
    CancelInstanceRefreshAsync(ctx workflow.Context, input *autoscaling.CancelInstanceRefreshInput) *AutoscalingCancelInstanceRefreshResult

    CompleteLifecycleAction(ctx workflow.Context, input *autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error)
    CompleteLifecycleActionAsync(ctx workflow.Context, input *autoscaling.CompleteLifecycleActionInput) *AutoscalingCompleteLifecycleActionResult

    CreateAutoScalingGroup(ctx workflow.Context, input *autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error)
    CreateAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.CreateAutoScalingGroupInput) *AutoscalingCreateAutoScalingGroupResult

    CreateLaunchConfiguration(ctx workflow.Context, input *autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error)
    CreateLaunchConfigurationAsync(ctx workflow.Context, input *autoscaling.CreateLaunchConfigurationInput) *AutoscalingCreateLaunchConfigurationResult

    CreateOrUpdateTags(ctx workflow.Context, input *autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error)
    CreateOrUpdateTagsAsync(ctx workflow.Context, input *autoscaling.CreateOrUpdateTagsInput) *AutoscalingCreateOrUpdateTagsResult

    DeleteAutoScalingGroup(ctx workflow.Context, input *autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error)
    DeleteAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.DeleteAutoScalingGroupInput) *AutoscalingDeleteAutoScalingGroupResult

    DeleteLaunchConfiguration(ctx workflow.Context, input *autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error)
    DeleteLaunchConfigurationAsync(ctx workflow.Context, input *autoscaling.DeleteLaunchConfigurationInput) *AutoscalingDeleteLaunchConfigurationResult

    DeleteLifecycleHook(ctx workflow.Context, input *autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error)
    DeleteLifecycleHookAsync(ctx workflow.Context, input *autoscaling.DeleteLifecycleHookInput) *AutoscalingDeleteLifecycleHookResult

    DeleteNotificationConfiguration(ctx workflow.Context, input *autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error)
    DeleteNotificationConfigurationAsync(ctx workflow.Context, input *autoscaling.DeleteNotificationConfigurationInput) *AutoscalingDeleteNotificationConfigurationResult

    DeletePolicy(ctx workflow.Context, input *autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error)
    DeletePolicyAsync(ctx workflow.Context, input *autoscaling.DeletePolicyInput) *AutoscalingDeletePolicyResult

    DeleteScheduledAction(ctx workflow.Context, input *autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error)
    DeleteScheduledActionAsync(ctx workflow.Context, input *autoscaling.DeleteScheduledActionInput) *AutoscalingDeleteScheduledActionResult

    DeleteTags(ctx workflow.Context, input *autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error)
    DeleteTagsAsync(ctx workflow.Context, input *autoscaling.DeleteTagsInput) *AutoscalingDeleteTagsResult

    DescribeAccountLimits(ctx workflow.Context, input *autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error)
    DescribeAccountLimitsAsync(ctx workflow.Context, input *autoscaling.DescribeAccountLimitsInput) *AutoscalingDescribeAccountLimitsResult

    DescribeAdjustmentTypes(ctx workflow.Context, input *autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error)
    DescribeAdjustmentTypesAsync(ctx workflow.Context, input *autoscaling.DescribeAdjustmentTypesInput) *AutoscalingDescribeAdjustmentTypesResult

    DescribeAutoScalingGroups(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
    DescribeAutoScalingGroupsAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) *AutoscalingDescribeAutoScalingGroupsResult

    DescribeAutoScalingInstances(ctx workflow.Context, input *autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error)
    DescribeAutoScalingInstancesAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingInstancesInput) *AutoscalingDescribeAutoScalingInstancesResult

    DescribeAutoScalingNotificationTypes(ctx workflow.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error)
    DescribeAutoScalingNotificationTypesAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) *AutoscalingDescribeAutoScalingNotificationTypesResult

    DescribeInstanceRefreshes(ctx workflow.Context, input *autoscaling.DescribeInstanceRefreshesInput) (*autoscaling.DescribeInstanceRefreshesOutput, error)
    DescribeInstanceRefreshesAsync(ctx workflow.Context, input *autoscaling.DescribeInstanceRefreshesInput) *AutoscalingDescribeInstanceRefreshesResult

    DescribeLaunchConfigurations(ctx workflow.Context, input *autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
    DescribeLaunchConfigurationsAsync(ctx workflow.Context, input *autoscaling.DescribeLaunchConfigurationsInput) *AutoscalingDescribeLaunchConfigurationsResult

    DescribeLifecycleHookTypes(ctx workflow.Context, input *autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error)
    DescribeLifecycleHookTypesAsync(ctx workflow.Context, input *autoscaling.DescribeLifecycleHookTypesInput) *AutoscalingDescribeLifecycleHookTypesResult

    DescribeLifecycleHooks(ctx workflow.Context, input *autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error)
    DescribeLifecycleHooksAsync(ctx workflow.Context, input *autoscaling.DescribeLifecycleHooksInput) *AutoscalingDescribeLifecycleHooksResult

    DescribeLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error)
    DescribeLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) *AutoscalingDescribeLoadBalancerTargetGroupsResult

    DescribeLoadBalancers(ctx workflow.Context, input *autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error)
    DescribeLoadBalancersAsync(ctx workflow.Context, input *autoscaling.DescribeLoadBalancersInput) *AutoscalingDescribeLoadBalancersResult

    DescribeMetricCollectionTypes(ctx workflow.Context, input *autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error)
    DescribeMetricCollectionTypesAsync(ctx workflow.Context, input *autoscaling.DescribeMetricCollectionTypesInput) *AutoscalingDescribeMetricCollectionTypesResult

    DescribeNotificationConfigurations(ctx workflow.Context, input *autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error)
    DescribeNotificationConfigurationsAsync(ctx workflow.Context, input *autoscaling.DescribeNotificationConfigurationsInput) *AutoscalingDescribeNotificationConfigurationsResult

    DescribePolicies(ctx workflow.Context, input *autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error)
    DescribePoliciesAsync(ctx workflow.Context, input *autoscaling.DescribePoliciesInput) *AutoscalingDescribePoliciesResult

    DescribeScalingActivities(ctx workflow.Context, input *autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error)
    DescribeScalingActivitiesAsync(ctx workflow.Context, input *autoscaling.DescribeScalingActivitiesInput) *AutoscalingDescribeScalingActivitiesResult

    DescribeScalingProcessTypes(ctx workflow.Context, input *autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error)
    DescribeScalingProcessTypesAsync(ctx workflow.Context, input *autoscaling.DescribeScalingProcessTypesInput) *AutoscalingDescribeScalingProcessTypesResult

    DescribeScheduledActions(ctx workflow.Context, input *autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error)
    DescribeScheduledActionsAsync(ctx workflow.Context, input *autoscaling.DescribeScheduledActionsInput) *AutoscalingDescribeScheduledActionsResult

    DescribeTags(ctx workflow.Context, input *autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error)
    DescribeTagsAsync(ctx workflow.Context, input *autoscaling.DescribeTagsInput) *AutoscalingDescribeTagsResult

    DescribeTerminationPolicyTypes(ctx workflow.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error)
    DescribeTerminationPolicyTypesAsync(ctx workflow.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) *AutoscalingDescribeTerminationPolicyTypesResult

    DetachInstances(ctx workflow.Context, input *autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error)
    DetachInstancesAsync(ctx workflow.Context, input *autoscaling.DetachInstancesInput) *AutoscalingDetachInstancesResult

    DetachLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error)
    DetachLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) *AutoscalingDetachLoadBalancerTargetGroupsResult

    DetachLoadBalancers(ctx workflow.Context, input *autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error)
    DetachLoadBalancersAsync(ctx workflow.Context, input *autoscaling.DetachLoadBalancersInput) *AutoscalingDetachLoadBalancersResult

    DisableMetricsCollection(ctx workflow.Context, input *autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error)
    DisableMetricsCollectionAsync(ctx workflow.Context, input *autoscaling.DisableMetricsCollectionInput) *AutoscalingDisableMetricsCollectionResult

    EnableMetricsCollection(ctx workflow.Context, input *autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error)
    EnableMetricsCollectionAsync(ctx workflow.Context, input *autoscaling.EnableMetricsCollectionInput) *AutoscalingEnableMetricsCollectionResult

    EnterStandby(ctx workflow.Context, input *autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error)
    EnterStandbyAsync(ctx workflow.Context, input *autoscaling.EnterStandbyInput) *AutoscalingEnterStandbyResult

    ExecutePolicy(ctx workflow.Context, input *autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error)
    ExecutePolicyAsync(ctx workflow.Context, input *autoscaling.ExecutePolicyInput) *AutoscalingExecutePolicyResult

    ExitStandby(ctx workflow.Context, input *autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error)
    ExitStandbyAsync(ctx workflow.Context, input *autoscaling.ExitStandbyInput) *AutoscalingExitStandbyResult

    PutLifecycleHook(ctx workflow.Context, input *autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error)
    PutLifecycleHookAsync(ctx workflow.Context, input *autoscaling.PutLifecycleHookInput) *AutoscalingPutLifecycleHookResult

    PutNotificationConfiguration(ctx workflow.Context, input *autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error)
    PutNotificationConfigurationAsync(ctx workflow.Context, input *autoscaling.PutNotificationConfigurationInput) *AutoscalingPutNotificationConfigurationResult

    PutScalingPolicy(ctx workflow.Context, input *autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error)
    PutScalingPolicyAsync(ctx workflow.Context, input *autoscaling.PutScalingPolicyInput) *AutoscalingPutScalingPolicyResult

    PutScheduledUpdateGroupAction(ctx workflow.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error)
    PutScheduledUpdateGroupActionAsync(ctx workflow.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) *AutoscalingPutScheduledUpdateGroupActionResult

    RecordLifecycleActionHeartbeat(ctx workflow.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error)
    RecordLifecycleActionHeartbeatAsync(ctx workflow.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) *AutoscalingRecordLifecycleActionHeartbeatResult

    ResumeProcesses(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error)
    ResumeProcessesAsync(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) *AutoscalingResumeProcessesResult

    SetDesiredCapacity(ctx workflow.Context, input *autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error)
    SetDesiredCapacityAsync(ctx workflow.Context, input *autoscaling.SetDesiredCapacityInput) *AutoscalingSetDesiredCapacityResult

    SetInstanceHealth(ctx workflow.Context, input *autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error)
    SetInstanceHealthAsync(ctx workflow.Context, input *autoscaling.SetInstanceHealthInput) *AutoscalingSetInstanceHealthResult

    SetInstanceProtection(ctx workflow.Context, input *autoscaling.SetInstanceProtectionInput) (*autoscaling.SetInstanceProtectionOutput, error)
    SetInstanceProtectionAsync(ctx workflow.Context, input *autoscaling.SetInstanceProtectionInput) *AutoscalingSetInstanceProtectionResult

    StartInstanceRefresh(ctx workflow.Context, input *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error)
    StartInstanceRefreshAsync(ctx workflow.Context, input *autoscaling.StartInstanceRefreshInput) *AutoscalingStartInstanceRefreshResult

    SuspendProcesses(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error)
    SuspendProcessesAsync(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) *AutoscalingSuspendProcessesResult

    TerminateInstanceInAutoScalingGroup(ctx workflow.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error)
    TerminateInstanceInAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) *AutoscalingTerminateInstanceInAutoScalingGroupResult

    UpdateAutoScalingGroup(ctx workflow.Context, input *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error)
    UpdateAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.UpdateAutoScalingGroupInput) *AutoscalingUpdateAutoScalingGroupResult

    WaitUntilGroupExists(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error
    WaitUntilGroupInService(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error
    WaitUntilGroupNotExists(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error}
type AutoscalingAttachInstancesResult struct {
	Result workflow.Future
}

func (r *AutoscalingAttachInstancesResult) Get(ctx workflow.Context) (*autoscaling.AttachInstancesOutput, error) {
    var output autoscaling.AttachInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingAttachLoadBalancerTargetGroupsResult struct {
	Result workflow.Future
}

func (r *AutoscalingAttachLoadBalancerTargetGroupsResult) Get(ctx workflow.Context) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error) {
    var output autoscaling.AttachLoadBalancerTargetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingAttachLoadBalancersResult struct {
	Result workflow.Future
}

func (r *AutoscalingAttachLoadBalancersResult) Get(ctx workflow.Context) (*autoscaling.AttachLoadBalancersOutput, error) {
    var output autoscaling.AttachLoadBalancersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingBatchDeleteScheduledActionResult struct {
	Result workflow.Future
}

func (r *AutoscalingBatchDeleteScheduledActionResult) Get(ctx workflow.Context) (*autoscaling.BatchDeleteScheduledActionOutput, error) {
    var output autoscaling.BatchDeleteScheduledActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingBatchPutScheduledUpdateGroupActionResult struct {
	Result workflow.Future
}

func (r *AutoscalingBatchPutScheduledUpdateGroupActionResult) Get(ctx workflow.Context) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error) {
    var output autoscaling.BatchPutScheduledUpdateGroupActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingCancelInstanceRefreshResult struct {
	Result workflow.Future
}

func (r *AutoscalingCancelInstanceRefreshResult) Get(ctx workflow.Context) (*autoscaling.CancelInstanceRefreshOutput, error) {
    var output autoscaling.CancelInstanceRefreshOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingCompleteLifecycleActionResult struct {
	Result workflow.Future
}

func (r *AutoscalingCompleteLifecycleActionResult) Get(ctx workflow.Context) (*autoscaling.CompleteLifecycleActionOutput, error) {
    var output autoscaling.CompleteLifecycleActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingCreateAutoScalingGroupResult struct {
	Result workflow.Future
}

func (r *AutoscalingCreateAutoScalingGroupResult) Get(ctx workflow.Context) (*autoscaling.CreateAutoScalingGroupOutput, error) {
    var output autoscaling.CreateAutoScalingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingCreateLaunchConfigurationResult struct {
	Result workflow.Future
}

func (r *AutoscalingCreateLaunchConfigurationResult) Get(ctx workflow.Context) (*autoscaling.CreateLaunchConfigurationOutput, error) {
    var output autoscaling.CreateLaunchConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingCreateOrUpdateTagsResult struct {
	Result workflow.Future
}

func (r *AutoscalingCreateOrUpdateTagsResult) Get(ctx workflow.Context) (*autoscaling.CreateOrUpdateTagsOutput, error) {
    var output autoscaling.CreateOrUpdateTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDeleteAutoScalingGroupResult struct {
	Result workflow.Future
}

func (r *AutoscalingDeleteAutoScalingGroupResult) Get(ctx workflow.Context) (*autoscaling.DeleteAutoScalingGroupOutput, error) {
    var output autoscaling.DeleteAutoScalingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDeleteLaunchConfigurationResult struct {
	Result workflow.Future
}

func (r *AutoscalingDeleteLaunchConfigurationResult) Get(ctx workflow.Context) (*autoscaling.DeleteLaunchConfigurationOutput, error) {
    var output autoscaling.DeleteLaunchConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDeleteLifecycleHookResult struct {
	Result workflow.Future
}

func (r *AutoscalingDeleteLifecycleHookResult) Get(ctx workflow.Context) (*autoscaling.DeleteLifecycleHookOutput, error) {
    var output autoscaling.DeleteLifecycleHookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDeleteNotificationConfigurationResult struct {
	Result workflow.Future
}

func (r *AutoscalingDeleteNotificationConfigurationResult) Get(ctx workflow.Context) (*autoscaling.DeleteNotificationConfigurationOutput, error) {
    var output autoscaling.DeleteNotificationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDeletePolicyResult struct {
	Result workflow.Future
}

func (r *AutoscalingDeletePolicyResult) Get(ctx workflow.Context) (*autoscaling.DeletePolicyOutput, error) {
    var output autoscaling.DeletePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDeleteScheduledActionResult struct {
	Result workflow.Future
}

func (r *AutoscalingDeleteScheduledActionResult) Get(ctx workflow.Context) (*autoscaling.DeleteScheduledActionOutput, error) {
    var output autoscaling.DeleteScheduledActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDeleteTagsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDeleteTagsResult) Get(ctx workflow.Context) (*autoscaling.DeleteTagsOutput, error) {
    var output autoscaling.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeAccountLimitsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeAccountLimitsResult) Get(ctx workflow.Context) (*autoscaling.DescribeAccountLimitsOutput, error) {
    var output autoscaling.DescribeAccountLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeAdjustmentTypesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeAdjustmentTypesResult) Get(ctx workflow.Context) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
    var output autoscaling.DescribeAdjustmentTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeAutoScalingGroupsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeAutoScalingGroupsResult) Get(ctx workflow.Context) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
    var output autoscaling.DescribeAutoScalingGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeAutoScalingInstancesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeAutoScalingInstancesResult) Get(ctx workflow.Context) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
    var output autoscaling.DescribeAutoScalingInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeAutoScalingNotificationTypesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeAutoScalingNotificationTypesResult) Get(ctx workflow.Context) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
    var output autoscaling.DescribeAutoScalingNotificationTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeInstanceRefreshesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeInstanceRefreshesResult) Get(ctx workflow.Context) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
    var output autoscaling.DescribeInstanceRefreshesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeLaunchConfigurationsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeLaunchConfigurationsResult) Get(ctx workflow.Context) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
    var output autoscaling.DescribeLaunchConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeLifecycleHookTypesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeLifecycleHookTypesResult) Get(ctx workflow.Context) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
    var output autoscaling.DescribeLifecycleHookTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeLifecycleHooksResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeLifecycleHooksResult) Get(ctx workflow.Context) (*autoscaling.DescribeLifecycleHooksOutput, error) {
    var output autoscaling.DescribeLifecycleHooksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeLoadBalancerTargetGroupsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeLoadBalancerTargetGroupsResult) Get(ctx workflow.Context) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
    var output autoscaling.DescribeLoadBalancerTargetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeLoadBalancersResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeLoadBalancersResult) Get(ctx workflow.Context) (*autoscaling.DescribeLoadBalancersOutput, error) {
    var output autoscaling.DescribeLoadBalancersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeMetricCollectionTypesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeMetricCollectionTypesResult) Get(ctx workflow.Context) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
    var output autoscaling.DescribeMetricCollectionTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeNotificationConfigurationsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeNotificationConfigurationsResult) Get(ctx workflow.Context) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
    var output autoscaling.DescribeNotificationConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribePoliciesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribePoliciesResult) Get(ctx workflow.Context) (*autoscaling.DescribePoliciesOutput, error) {
    var output autoscaling.DescribePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeScalingActivitiesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeScalingActivitiesResult) Get(ctx workflow.Context) (*autoscaling.DescribeScalingActivitiesOutput, error) {
    var output autoscaling.DescribeScalingActivitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeScalingProcessTypesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeScalingProcessTypesResult) Get(ctx workflow.Context) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
    var output autoscaling.DescribeScalingProcessTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeScheduledActionsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeScheduledActionsResult) Get(ctx workflow.Context) (*autoscaling.DescribeScheduledActionsOutput, error) {
    var output autoscaling.DescribeScheduledActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeTagsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeTagsResult) Get(ctx workflow.Context) (*autoscaling.DescribeTagsOutput, error) {
    var output autoscaling.DescribeTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDescribeTerminationPolicyTypesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDescribeTerminationPolicyTypesResult) Get(ctx workflow.Context) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
    var output autoscaling.DescribeTerminationPolicyTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDetachInstancesResult struct {
	Result workflow.Future
}

func (r *AutoscalingDetachInstancesResult) Get(ctx workflow.Context) (*autoscaling.DetachInstancesOutput, error) {
    var output autoscaling.DetachInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDetachLoadBalancerTargetGroupsResult struct {
	Result workflow.Future
}

func (r *AutoscalingDetachLoadBalancerTargetGroupsResult) Get(ctx workflow.Context) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error) {
    var output autoscaling.DetachLoadBalancerTargetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDetachLoadBalancersResult struct {
	Result workflow.Future
}

func (r *AutoscalingDetachLoadBalancersResult) Get(ctx workflow.Context) (*autoscaling.DetachLoadBalancersOutput, error) {
    var output autoscaling.DetachLoadBalancersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingDisableMetricsCollectionResult struct {
	Result workflow.Future
}

func (r *AutoscalingDisableMetricsCollectionResult) Get(ctx workflow.Context) (*autoscaling.DisableMetricsCollectionOutput, error) {
    var output autoscaling.DisableMetricsCollectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingEnableMetricsCollectionResult struct {
	Result workflow.Future
}

func (r *AutoscalingEnableMetricsCollectionResult) Get(ctx workflow.Context) (*autoscaling.EnableMetricsCollectionOutput, error) {
    var output autoscaling.EnableMetricsCollectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingEnterStandbyResult struct {
	Result workflow.Future
}

func (r *AutoscalingEnterStandbyResult) Get(ctx workflow.Context) (*autoscaling.EnterStandbyOutput, error) {
    var output autoscaling.EnterStandbyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingExecutePolicyResult struct {
	Result workflow.Future
}

func (r *AutoscalingExecutePolicyResult) Get(ctx workflow.Context) (*autoscaling.ExecutePolicyOutput, error) {
    var output autoscaling.ExecutePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingExitStandbyResult struct {
	Result workflow.Future
}

func (r *AutoscalingExitStandbyResult) Get(ctx workflow.Context) (*autoscaling.ExitStandbyOutput, error) {
    var output autoscaling.ExitStandbyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingPutLifecycleHookResult struct {
	Result workflow.Future
}

func (r *AutoscalingPutLifecycleHookResult) Get(ctx workflow.Context) (*autoscaling.PutLifecycleHookOutput, error) {
    var output autoscaling.PutLifecycleHookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingPutNotificationConfigurationResult struct {
	Result workflow.Future
}

func (r *AutoscalingPutNotificationConfigurationResult) Get(ctx workflow.Context) (*autoscaling.PutNotificationConfigurationOutput, error) {
    var output autoscaling.PutNotificationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingPutScalingPolicyResult struct {
	Result workflow.Future
}

func (r *AutoscalingPutScalingPolicyResult) Get(ctx workflow.Context) (*autoscaling.PutScalingPolicyOutput, error) {
    var output autoscaling.PutScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingPutScheduledUpdateGroupActionResult struct {
	Result workflow.Future
}

func (r *AutoscalingPutScheduledUpdateGroupActionResult) Get(ctx workflow.Context) (*autoscaling.PutScheduledUpdateGroupActionOutput, error) {
    var output autoscaling.PutScheduledUpdateGroupActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingRecordLifecycleActionHeartbeatResult struct {
	Result workflow.Future
}

func (r *AutoscalingRecordLifecycleActionHeartbeatResult) Get(ctx workflow.Context) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error) {
    var output autoscaling.RecordLifecycleActionHeartbeatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingResumeProcessesResult struct {
	Result workflow.Future
}

func (r *AutoscalingResumeProcessesResult) Get(ctx workflow.Context) (*autoscaling.ResumeProcessesOutput, error) {
    var output autoscaling.ResumeProcessesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingSetDesiredCapacityResult struct {
	Result workflow.Future
}

func (r *AutoscalingSetDesiredCapacityResult) Get(ctx workflow.Context) (*autoscaling.SetDesiredCapacityOutput, error) {
    var output autoscaling.SetDesiredCapacityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingSetInstanceHealthResult struct {
	Result workflow.Future
}

func (r *AutoscalingSetInstanceHealthResult) Get(ctx workflow.Context) (*autoscaling.SetInstanceHealthOutput, error) {
    var output autoscaling.SetInstanceHealthOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingSetInstanceProtectionResult struct {
	Result workflow.Future
}

func (r *AutoscalingSetInstanceProtectionResult) Get(ctx workflow.Context) (*autoscaling.SetInstanceProtectionOutput, error) {
    var output autoscaling.SetInstanceProtectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingStartInstanceRefreshResult struct {
	Result workflow.Future
}

func (r *AutoscalingStartInstanceRefreshResult) Get(ctx workflow.Context) (*autoscaling.StartInstanceRefreshOutput, error) {
    var output autoscaling.StartInstanceRefreshOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingSuspendProcessesResult struct {
	Result workflow.Future
}

func (r *AutoscalingSuspendProcessesResult) Get(ctx workflow.Context) (*autoscaling.SuspendProcessesOutput, error) {
    var output autoscaling.SuspendProcessesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingTerminateInstanceInAutoScalingGroupResult struct {
	Result workflow.Future
}

func (r *AutoscalingTerminateInstanceInAutoScalingGroupResult) Get(ctx workflow.Context) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error) {
    var output autoscaling.TerminateInstanceInAutoScalingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingUpdateAutoScalingGroupResult struct {
	Result workflow.Future
}

func (r *AutoscalingUpdateAutoScalingGroupResult) Get(ctx workflow.Context) (*autoscaling.UpdateAutoScalingGroupOutput, error) {
    var output autoscaling.UpdateAutoScalingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type AutoScalingStub struct {
    activities awsactivities.AutoScalingActivities
}

func NewAutoScalingStub() AutoScalingClient {
    return &AutoScalingStub{}
}
func (a *AutoScalingStub) AttachInstances(ctx workflow.Context, input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error) {
    var output autoscaling.AttachInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) AttachInstancesAsync(ctx workflow.Context, input *autoscaling.AttachInstancesInput) *AutoscalingAttachInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachInstances, input)
    return &AutoscalingAttachInstancesResult{Result: future}
}
func (a *AutoScalingStub) AttachLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error) {
    var output autoscaling.AttachLoadBalancerTargetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachLoadBalancerTargetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) AttachLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) *AutoscalingAttachLoadBalancerTargetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachLoadBalancerTargetGroups, input)
    return &AutoscalingAttachLoadBalancerTargetGroupsResult{Result: future}
}
func (a *AutoScalingStub) AttachLoadBalancers(ctx workflow.Context, input *autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error) {
    var output autoscaling.AttachLoadBalancersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachLoadBalancers, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) AttachLoadBalancersAsync(ctx workflow.Context, input *autoscaling.AttachLoadBalancersInput) *AutoscalingAttachLoadBalancersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachLoadBalancers, input)
    return &AutoscalingAttachLoadBalancersResult{Result: future}
}
func (a *AutoScalingStub) BatchDeleteScheduledAction(ctx workflow.Context, input *autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error) {
    var output autoscaling.BatchDeleteScheduledActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteScheduledAction, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) BatchDeleteScheduledActionAsync(ctx workflow.Context, input *autoscaling.BatchDeleteScheduledActionInput) *AutoscalingBatchDeleteScheduledActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteScheduledAction, input)
    return &AutoscalingBatchDeleteScheduledActionResult{Result: future}
}
func (a *AutoScalingStub) BatchPutScheduledUpdateGroupAction(ctx workflow.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error) {
    var output autoscaling.BatchPutScheduledUpdateGroupActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchPutScheduledUpdateGroupAction, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) BatchPutScheduledUpdateGroupActionAsync(ctx workflow.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) *AutoscalingBatchPutScheduledUpdateGroupActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchPutScheduledUpdateGroupAction, input)
    return &AutoscalingBatchPutScheduledUpdateGroupActionResult{Result: future}
}
func (a *AutoScalingStub) CancelInstanceRefresh(ctx workflow.Context, input *autoscaling.CancelInstanceRefreshInput) (*autoscaling.CancelInstanceRefreshOutput, error) {
    var output autoscaling.CancelInstanceRefreshOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelInstanceRefresh, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) CancelInstanceRefreshAsync(ctx workflow.Context, input *autoscaling.CancelInstanceRefreshInput) *AutoscalingCancelInstanceRefreshResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelInstanceRefresh, input)
    return &AutoscalingCancelInstanceRefreshResult{Result: future}
}
func (a *AutoScalingStub) CompleteLifecycleAction(ctx workflow.Context, input *autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error) {
    var output autoscaling.CompleteLifecycleActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CompleteLifecycleAction, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) CompleteLifecycleActionAsync(ctx workflow.Context, input *autoscaling.CompleteLifecycleActionInput) *AutoscalingCompleteLifecycleActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CompleteLifecycleAction, input)
    return &AutoscalingCompleteLifecycleActionResult{Result: future}
}
func (a *AutoScalingStub) CreateAutoScalingGroup(ctx workflow.Context, input *autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error) {
    var output autoscaling.CreateAutoScalingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAutoScalingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) CreateAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.CreateAutoScalingGroupInput) *AutoscalingCreateAutoScalingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAutoScalingGroup, input)
    return &AutoscalingCreateAutoScalingGroupResult{Result: future}
}
func (a *AutoScalingStub) CreateLaunchConfiguration(ctx workflow.Context, input *autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error) {
    var output autoscaling.CreateLaunchConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLaunchConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) CreateLaunchConfigurationAsync(ctx workflow.Context, input *autoscaling.CreateLaunchConfigurationInput) *AutoscalingCreateLaunchConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLaunchConfiguration, input)
    return &AutoscalingCreateLaunchConfigurationResult{Result: future}
}
func (a *AutoScalingStub) CreateOrUpdateTags(ctx workflow.Context, input *autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error) {
    var output autoscaling.CreateOrUpdateTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOrUpdateTags, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) CreateOrUpdateTagsAsync(ctx workflow.Context, input *autoscaling.CreateOrUpdateTagsInput) *AutoscalingCreateOrUpdateTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateOrUpdateTags, input)
    return &AutoscalingCreateOrUpdateTagsResult{Result: future}
}
func (a *AutoScalingStub) DeleteAutoScalingGroup(ctx workflow.Context, input *autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error) {
    var output autoscaling.DeleteAutoScalingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAutoScalingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DeleteAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.DeleteAutoScalingGroupInput) *AutoscalingDeleteAutoScalingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAutoScalingGroup, input)
    return &AutoscalingDeleteAutoScalingGroupResult{Result: future}
}
func (a *AutoScalingStub) DeleteLaunchConfiguration(ctx workflow.Context, input *autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error) {
    var output autoscaling.DeleteLaunchConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLaunchConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DeleteLaunchConfigurationAsync(ctx workflow.Context, input *autoscaling.DeleteLaunchConfigurationInput) *AutoscalingDeleteLaunchConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLaunchConfiguration, input)
    return &AutoscalingDeleteLaunchConfigurationResult{Result: future}
}
func (a *AutoScalingStub) DeleteLifecycleHook(ctx workflow.Context, input *autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error) {
    var output autoscaling.DeleteLifecycleHookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLifecycleHook, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DeleteLifecycleHookAsync(ctx workflow.Context, input *autoscaling.DeleteLifecycleHookInput) *AutoscalingDeleteLifecycleHookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLifecycleHook, input)
    return &AutoscalingDeleteLifecycleHookResult{Result: future}
}
func (a *AutoScalingStub) DeleteNotificationConfiguration(ctx workflow.Context, input *autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error) {
    var output autoscaling.DeleteNotificationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNotificationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DeleteNotificationConfigurationAsync(ctx workflow.Context, input *autoscaling.DeleteNotificationConfigurationInput) *AutoscalingDeleteNotificationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNotificationConfiguration, input)
    return &AutoscalingDeleteNotificationConfigurationResult{Result: future}
}
func (a *AutoScalingStub) DeletePolicy(ctx workflow.Context, input *autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error) {
    var output autoscaling.DeletePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DeletePolicyAsync(ctx workflow.Context, input *autoscaling.DeletePolicyInput) *AutoscalingDeletePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input)
    return &AutoscalingDeletePolicyResult{Result: future}
}
func (a *AutoScalingStub) DeleteScheduledAction(ctx workflow.Context, input *autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error) {
    var output autoscaling.DeleteScheduledActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteScheduledAction, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DeleteScheduledActionAsync(ctx workflow.Context, input *autoscaling.DeleteScheduledActionInput) *AutoscalingDeleteScheduledActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteScheduledAction, input)
    return &AutoscalingDeleteScheduledActionResult{Result: future}
}
func (a *AutoScalingStub) DeleteTags(ctx workflow.Context, input *autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error) {
    var output autoscaling.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DeleteTagsAsync(ctx workflow.Context, input *autoscaling.DeleteTagsInput) *AutoscalingDeleteTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input)
    return &AutoscalingDeleteTagsResult{Result: future}
}
func (a *AutoScalingStub) DescribeAccountLimits(ctx workflow.Context, input *autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error) {
    var output autoscaling.DescribeAccountLimitsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountLimits, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeAccountLimitsAsync(ctx workflow.Context, input *autoscaling.DescribeAccountLimitsInput) *AutoscalingDescribeAccountLimitsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountLimits, input)
    return &AutoscalingDescribeAccountLimitsResult{Result: future}
}
func (a *AutoScalingStub) DescribeAdjustmentTypes(ctx workflow.Context, input *autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
    var output autoscaling.DescribeAdjustmentTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAdjustmentTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeAdjustmentTypesAsync(ctx workflow.Context, input *autoscaling.DescribeAdjustmentTypesInput) *AutoscalingDescribeAdjustmentTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAdjustmentTypes, input)
    return &AutoscalingDescribeAdjustmentTypesResult{Result: future}
}
func (a *AutoScalingStub) DescribeAutoScalingGroups(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
    var output autoscaling.DescribeAutoScalingGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAutoScalingGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeAutoScalingGroupsAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) *AutoscalingDescribeAutoScalingGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAutoScalingGroups, input)
    return &AutoscalingDescribeAutoScalingGroupsResult{Result: future}
}
func (a *AutoScalingStub) DescribeAutoScalingInstances(ctx workflow.Context, input *autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
    var output autoscaling.DescribeAutoScalingInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAutoScalingInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeAutoScalingInstancesAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingInstancesInput) *AutoscalingDescribeAutoScalingInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAutoScalingInstances, input)
    return &AutoscalingDescribeAutoScalingInstancesResult{Result: future}
}
func (a *AutoScalingStub) DescribeAutoScalingNotificationTypes(ctx workflow.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
    var output autoscaling.DescribeAutoScalingNotificationTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAutoScalingNotificationTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeAutoScalingNotificationTypesAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) *AutoscalingDescribeAutoScalingNotificationTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAutoScalingNotificationTypes, input)
    return &AutoscalingDescribeAutoScalingNotificationTypesResult{Result: future}
}
func (a *AutoScalingStub) DescribeInstanceRefreshes(ctx workflow.Context, input *autoscaling.DescribeInstanceRefreshesInput) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
    var output autoscaling.DescribeInstanceRefreshesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceRefreshes, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeInstanceRefreshesAsync(ctx workflow.Context, input *autoscaling.DescribeInstanceRefreshesInput) *AutoscalingDescribeInstanceRefreshesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceRefreshes, input)
    return &AutoscalingDescribeInstanceRefreshesResult{Result: future}
}
func (a *AutoScalingStub) DescribeLaunchConfigurations(ctx workflow.Context, input *autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
    var output autoscaling.DescribeLaunchConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLaunchConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeLaunchConfigurationsAsync(ctx workflow.Context, input *autoscaling.DescribeLaunchConfigurationsInput) *AutoscalingDescribeLaunchConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLaunchConfigurations, input)
    return &AutoscalingDescribeLaunchConfigurationsResult{Result: future}
}
func (a *AutoScalingStub) DescribeLifecycleHookTypes(ctx workflow.Context, input *autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
    var output autoscaling.DescribeLifecycleHookTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLifecycleHookTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeLifecycleHookTypesAsync(ctx workflow.Context, input *autoscaling.DescribeLifecycleHookTypesInput) *AutoscalingDescribeLifecycleHookTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLifecycleHookTypes, input)
    return &AutoscalingDescribeLifecycleHookTypesResult{Result: future}
}
func (a *AutoScalingStub) DescribeLifecycleHooks(ctx workflow.Context, input *autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error) {
    var output autoscaling.DescribeLifecycleHooksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLifecycleHooks, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeLifecycleHooksAsync(ctx workflow.Context, input *autoscaling.DescribeLifecycleHooksInput) *AutoscalingDescribeLifecycleHooksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLifecycleHooks, input)
    return &AutoscalingDescribeLifecycleHooksResult{Result: future}
}
func (a *AutoScalingStub) DescribeLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
    var output autoscaling.DescribeLoadBalancerTargetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLoadBalancerTargetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) *AutoscalingDescribeLoadBalancerTargetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLoadBalancerTargetGroups, input)
    return &AutoscalingDescribeLoadBalancerTargetGroupsResult{Result: future}
}
func (a *AutoScalingStub) DescribeLoadBalancers(ctx workflow.Context, input *autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error) {
    var output autoscaling.DescribeLoadBalancersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLoadBalancers, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeLoadBalancersAsync(ctx workflow.Context, input *autoscaling.DescribeLoadBalancersInput) *AutoscalingDescribeLoadBalancersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLoadBalancers, input)
    return &AutoscalingDescribeLoadBalancersResult{Result: future}
}
func (a *AutoScalingStub) DescribeMetricCollectionTypes(ctx workflow.Context, input *autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
    var output autoscaling.DescribeMetricCollectionTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMetricCollectionTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeMetricCollectionTypesAsync(ctx workflow.Context, input *autoscaling.DescribeMetricCollectionTypesInput) *AutoscalingDescribeMetricCollectionTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMetricCollectionTypes, input)
    return &AutoscalingDescribeMetricCollectionTypesResult{Result: future}
}
func (a *AutoScalingStub) DescribeNotificationConfigurations(ctx workflow.Context, input *autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
    var output autoscaling.DescribeNotificationConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNotificationConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeNotificationConfigurationsAsync(ctx workflow.Context, input *autoscaling.DescribeNotificationConfigurationsInput) *AutoscalingDescribeNotificationConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNotificationConfigurations, input)
    return &AutoscalingDescribeNotificationConfigurationsResult{Result: future}
}
func (a *AutoScalingStub) DescribePolicies(ctx workflow.Context, input *autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error) {
    var output autoscaling.DescribePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribePoliciesAsync(ctx workflow.Context, input *autoscaling.DescribePoliciesInput) *AutoscalingDescribePoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePolicies, input)
    return &AutoscalingDescribePoliciesResult{Result: future}
}
func (a *AutoScalingStub) DescribeScalingActivities(ctx workflow.Context, input *autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error) {
    var output autoscaling.DescribeScalingActivitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingActivities, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeScalingActivitiesAsync(ctx workflow.Context, input *autoscaling.DescribeScalingActivitiesInput) *AutoscalingDescribeScalingActivitiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingActivities, input)
    return &AutoscalingDescribeScalingActivitiesResult{Result: future}
}
func (a *AutoScalingStub) DescribeScalingProcessTypes(ctx workflow.Context, input *autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
    var output autoscaling.DescribeScalingProcessTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingProcessTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeScalingProcessTypesAsync(ctx workflow.Context, input *autoscaling.DescribeScalingProcessTypesInput) *AutoscalingDescribeScalingProcessTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingProcessTypes, input)
    return &AutoscalingDescribeScalingProcessTypesResult{Result: future}
}
func (a *AutoScalingStub) DescribeScheduledActions(ctx workflow.Context, input *autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error) {
    var output autoscaling.DescribeScheduledActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledActions, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeScheduledActionsAsync(ctx workflow.Context, input *autoscaling.DescribeScheduledActionsInput) *AutoscalingDescribeScheduledActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledActions, input)
    return &AutoscalingDescribeScheduledActionsResult{Result: future}
}
func (a *AutoScalingStub) DescribeTags(ctx workflow.Context, input *autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error) {
    var output autoscaling.DescribeTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeTagsAsync(ctx workflow.Context, input *autoscaling.DescribeTagsInput) *AutoscalingDescribeTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input)
    return &AutoscalingDescribeTagsResult{Result: future}
}
func (a *AutoScalingStub) DescribeTerminationPolicyTypes(ctx workflow.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
    var output autoscaling.DescribeTerminationPolicyTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTerminationPolicyTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DescribeTerminationPolicyTypesAsync(ctx workflow.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) *AutoscalingDescribeTerminationPolicyTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTerminationPolicyTypes, input)
    return &AutoscalingDescribeTerminationPolicyTypesResult{Result: future}
}
func (a *AutoScalingStub) DetachInstances(ctx workflow.Context, input *autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error) {
    var output autoscaling.DetachInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DetachInstancesAsync(ctx workflow.Context, input *autoscaling.DetachInstancesInput) *AutoscalingDetachInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachInstances, input)
    return &AutoscalingDetachInstancesResult{Result: future}
}
func (a *AutoScalingStub) DetachLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error) {
    var output autoscaling.DetachLoadBalancerTargetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachLoadBalancerTargetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DetachLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) *AutoscalingDetachLoadBalancerTargetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachLoadBalancerTargetGroups, input)
    return &AutoscalingDetachLoadBalancerTargetGroupsResult{Result: future}
}
func (a *AutoScalingStub) DetachLoadBalancers(ctx workflow.Context, input *autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error) {
    var output autoscaling.DetachLoadBalancersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachLoadBalancers, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DetachLoadBalancersAsync(ctx workflow.Context, input *autoscaling.DetachLoadBalancersInput) *AutoscalingDetachLoadBalancersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachLoadBalancers, input)
    return &AutoscalingDetachLoadBalancersResult{Result: future}
}
func (a *AutoScalingStub) DisableMetricsCollection(ctx workflow.Context, input *autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error) {
    var output autoscaling.DisableMetricsCollectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableMetricsCollection, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) DisableMetricsCollectionAsync(ctx workflow.Context, input *autoscaling.DisableMetricsCollectionInput) *AutoscalingDisableMetricsCollectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableMetricsCollection, input)
    return &AutoscalingDisableMetricsCollectionResult{Result: future}
}
func (a *AutoScalingStub) EnableMetricsCollection(ctx workflow.Context, input *autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error) {
    var output autoscaling.EnableMetricsCollectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableMetricsCollection, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) EnableMetricsCollectionAsync(ctx workflow.Context, input *autoscaling.EnableMetricsCollectionInput) *AutoscalingEnableMetricsCollectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableMetricsCollection, input)
    return &AutoscalingEnableMetricsCollectionResult{Result: future}
}
func (a *AutoScalingStub) EnterStandby(ctx workflow.Context, input *autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error) {
    var output autoscaling.EnterStandbyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnterStandby, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) EnterStandbyAsync(ctx workflow.Context, input *autoscaling.EnterStandbyInput) *AutoscalingEnterStandbyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnterStandby, input)
    return &AutoscalingEnterStandbyResult{Result: future}
}
func (a *AutoScalingStub) ExecutePolicy(ctx workflow.Context, input *autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error) {
    var output autoscaling.ExecutePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExecutePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) ExecutePolicyAsync(ctx workflow.Context, input *autoscaling.ExecutePolicyInput) *AutoscalingExecutePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExecutePolicy, input)
    return &AutoscalingExecutePolicyResult{Result: future}
}
func (a *AutoScalingStub) ExitStandby(ctx workflow.Context, input *autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error) {
    var output autoscaling.ExitStandbyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExitStandby, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) ExitStandbyAsync(ctx workflow.Context, input *autoscaling.ExitStandbyInput) *AutoscalingExitStandbyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExitStandby, input)
    return &AutoscalingExitStandbyResult{Result: future}
}
func (a *AutoScalingStub) PutLifecycleHook(ctx workflow.Context, input *autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error) {
    var output autoscaling.PutLifecycleHookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLifecycleHook, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) PutLifecycleHookAsync(ctx workflow.Context, input *autoscaling.PutLifecycleHookInput) *AutoscalingPutLifecycleHookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutLifecycleHook, input)
    return &AutoscalingPutLifecycleHookResult{Result: future}
}
func (a *AutoScalingStub) PutNotificationConfiguration(ctx workflow.Context, input *autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error) {
    var output autoscaling.PutNotificationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutNotificationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) PutNotificationConfigurationAsync(ctx workflow.Context, input *autoscaling.PutNotificationConfigurationInput) *AutoscalingPutNotificationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutNotificationConfiguration, input)
    return &AutoscalingPutNotificationConfigurationResult{Result: future}
}
func (a *AutoScalingStub) PutScalingPolicy(ctx workflow.Context, input *autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error) {
    var output autoscaling.PutScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) PutScalingPolicyAsync(ctx workflow.Context, input *autoscaling.PutScalingPolicyInput) *AutoscalingPutScalingPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutScalingPolicy, input)
    return &AutoscalingPutScalingPolicyResult{Result: future}
}
func (a *AutoScalingStub) PutScheduledUpdateGroupAction(ctx workflow.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error) {
    var output autoscaling.PutScheduledUpdateGroupActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutScheduledUpdateGroupAction, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) PutScheduledUpdateGroupActionAsync(ctx workflow.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) *AutoscalingPutScheduledUpdateGroupActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutScheduledUpdateGroupAction, input)
    return &AutoscalingPutScheduledUpdateGroupActionResult{Result: future}
}
func (a *AutoScalingStub) RecordLifecycleActionHeartbeat(ctx workflow.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error) {
    var output autoscaling.RecordLifecycleActionHeartbeatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RecordLifecycleActionHeartbeat, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) RecordLifecycleActionHeartbeatAsync(ctx workflow.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) *AutoscalingRecordLifecycleActionHeartbeatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RecordLifecycleActionHeartbeat, input)
    return &AutoscalingRecordLifecycleActionHeartbeatResult{Result: future}
}
func (a *AutoScalingStub) ResumeProcesses(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error) {
    var output autoscaling.ResumeProcessesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResumeProcesses, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) ResumeProcessesAsync(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) *AutoscalingResumeProcessesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResumeProcesses, input)
    return &AutoscalingResumeProcessesResult{Result: future}
}
func (a *AutoScalingStub) SetDesiredCapacity(ctx workflow.Context, input *autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error) {
    var output autoscaling.SetDesiredCapacityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetDesiredCapacity, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) SetDesiredCapacityAsync(ctx workflow.Context, input *autoscaling.SetDesiredCapacityInput) *AutoscalingSetDesiredCapacityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetDesiredCapacity, input)
    return &AutoscalingSetDesiredCapacityResult{Result: future}
}
func (a *AutoScalingStub) SetInstanceHealth(ctx workflow.Context, input *autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error) {
    var output autoscaling.SetInstanceHealthOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetInstanceHealth, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) SetInstanceHealthAsync(ctx workflow.Context, input *autoscaling.SetInstanceHealthInput) *AutoscalingSetInstanceHealthResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetInstanceHealth, input)
    return &AutoscalingSetInstanceHealthResult{Result: future}
}
func (a *AutoScalingStub) SetInstanceProtection(ctx workflow.Context, input *autoscaling.SetInstanceProtectionInput) (*autoscaling.SetInstanceProtectionOutput, error) {
    var output autoscaling.SetInstanceProtectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetInstanceProtection, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) SetInstanceProtectionAsync(ctx workflow.Context, input *autoscaling.SetInstanceProtectionInput) *AutoscalingSetInstanceProtectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetInstanceProtection, input)
    return &AutoscalingSetInstanceProtectionResult{Result: future}
}
func (a *AutoScalingStub) StartInstanceRefresh(ctx workflow.Context, input *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error) {
    var output autoscaling.StartInstanceRefreshOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartInstanceRefresh, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) StartInstanceRefreshAsync(ctx workflow.Context, input *autoscaling.StartInstanceRefreshInput) *AutoscalingStartInstanceRefreshResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartInstanceRefresh, input)
    return &AutoscalingStartInstanceRefreshResult{Result: future}
}
func (a *AutoScalingStub) SuspendProcesses(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error) {
    var output autoscaling.SuspendProcessesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SuspendProcesses, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) SuspendProcessesAsync(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) *AutoscalingSuspendProcessesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SuspendProcesses, input)
    return &AutoscalingSuspendProcessesResult{Result: future}
}
func (a *AutoScalingStub) TerminateInstanceInAutoScalingGroup(ctx workflow.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error) {
    var output autoscaling.TerminateInstanceInAutoScalingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateInstanceInAutoScalingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) TerminateInstanceInAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) *AutoscalingTerminateInstanceInAutoScalingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateInstanceInAutoScalingGroup, input)
    return &AutoscalingTerminateInstanceInAutoScalingGroupResult{Result: future}
}
func (a *AutoScalingStub) UpdateAutoScalingGroup(ctx workflow.Context, input *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error) {
    var output autoscaling.UpdateAutoScalingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAutoScalingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingStub) UpdateAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.UpdateAutoScalingGroupInput) *AutoscalingUpdateAutoScalingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAutoScalingGroup, input)
    return &AutoscalingUpdateAutoScalingGroupResult{Result: future}
}

func (a *AutoScalingStub) WaitUntilGroupExists(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilGroupExists, input).Get(ctx, nil)
}

func (a *AutoScalingStub) WaitUntilGroupExistsAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilGroupExists, input)
}

func (a *AutoScalingStub) WaitUntilGroupInService(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilGroupInService, input).Get(ctx, nil)
}

func (a *AutoScalingStub) WaitUntilGroupInServiceAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilGroupInService, input)
}

func (a *AutoScalingStub) WaitUntilGroupNotExists(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilGroupNotExists, input).Get(ctx, nil)
}

func (a *AutoScalingStub) WaitUntilGroupNotExistsAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilGroupNotExists, input)
}