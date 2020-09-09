
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

type AutoScalingActivities struct {
    client autoscalingiface.AutoScalingAPI
}

func NewAutoScalingActivities(session *session.Session, config ...*aws.Config) *AutoScalingActivities {
    client := autoscaling.New(session, config...)
    return &AutoScalingActivities{client: client}
}

func (a *AutoScalingActivities) AttachInstances(input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error) {
    return a.client.AttachInstances(input)
}

func (a *AutoScalingActivities) AttachLoadBalancerTargetGroups(input *autoscaling.AttachLoadBalancerTargetGroupsInput) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error) {
    return a.client.AttachLoadBalancerTargetGroups(input)
}

func (a *AutoScalingActivities) AttachLoadBalancers(input *autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error) {
    return a.client.AttachLoadBalancers(input)
}

func (a *AutoScalingActivities) BatchDeleteScheduledAction(input *autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error) {
    return a.client.BatchDeleteScheduledAction(input)
}

func (a *AutoScalingActivities) BatchPutScheduledUpdateGroupAction(input *autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error) {
    return a.client.BatchPutScheduledUpdateGroupAction(input)
}

func (a *AutoScalingActivities) CancelInstanceRefresh(input *autoscaling.CancelInstanceRefreshInput) (*autoscaling.CancelInstanceRefreshOutput, error) {
    return a.client.CancelInstanceRefresh(input)
}

func (a *AutoScalingActivities) CompleteLifecycleAction(input *autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error) {
    return a.client.CompleteLifecycleAction(input)
}

func (a *AutoScalingActivities) CreateAutoScalingGroup(input *autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error) {
    return a.client.CreateAutoScalingGroup(input)
}

func (a *AutoScalingActivities) CreateLaunchConfiguration(input *autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error) {
    return a.client.CreateLaunchConfiguration(input)
}

func (a *AutoScalingActivities) CreateOrUpdateTags(input *autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error) {
    return a.client.CreateOrUpdateTags(input)
}

func (a *AutoScalingActivities) DeleteAutoScalingGroup(input *autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error) {
    return a.client.DeleteAutoScalingGroup(input)
}

func (a *AutoScalingActivities) DeleteLaunchConfiguration(input *autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error) {
    return a.client.DeleteLaunchConfiguration(input)
}

func (a *AutoScalingActivities) DeleteLifecycleHook(input *autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error) {
    return a.client.DeleteLifecycleHook(input)
}

func (a *AutoScalingActivities) DeleteNotificationConfiguration(input *autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error) {
    return a.client.DeleteNotificationConfiguration(input)
}

func (a *AutoScalingActivities) DeletePolicy(input *autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error) {
    return a.client.DeletePolicy(input)
}

func (a *AutoScalingActivities) DeleteScheduledAction(input *autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error) {
    return a.client.DeleteScheduledAction(input)
}

func (a *AutoScalingActivities) DeleteTags(input *autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}

func (a *AutoScalingActivities) DescribeAccountLimits(input *autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error) {
    return a.client.DescribeAccountLimits(input)
}

func (a *AutoScalingActivities) DescribeAdjustmentTypes(input *autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
    return a.client.DescribeAdjustmentTypes(input)
}

func (a *AutoScalingActivities) DescribeAutoScalingGroups(input *autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
    return a.client.DescribeAutoScalingGroups(input)
}

func (a *AutoScalingActivities) DescribeAutoScalingInstances(input *autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
    return a.client.DescribeAutoScalingInstances(input)
}

func (a *AutoScalingActivities) DescribeAutoScalingNotificationTypes(input *autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
    return a.client.DescribeAutoScalingNotificationTypes(input)
}

func (a *AutoScalingActivities) DescribeInstanceRefreshes(input *autoscaling.DescribeInstanceRefreshesInput) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
    return a.client.DescribeInstanceRefreshes(input)
}

func (a *AutoScalingActivities) DescribeLaunchConfigurations(input *autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
    return a.client.DescribeLaunchConfigurations(input)
}

func (a *AutoScalingActivities) DescribeLifecycleHookTypes(input *autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
    return a.client.DescribeLifecycleHookTypes(input)
}

func (a *AutoScalingActivities) DescribeLifecycleHooks(input *autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error) {
    return a.client.DescribeLifecycleHooks(input)
}

func (a *AutoScalingActivities) DescribeLoadBalancerTargetGroups(input *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
    return a.client.DescribeLoadBalancerTargetGroups(input)
}

func (a *AutoScalingActivities) DescribeLoadBalancers(input *autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error) {
    return a.client.DescribeLoadBalancers(input)
}

func (a *AutoScalingActivities) DescribeMetricCollectionTypes(input *autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
    return a.client.DescribeMetricCollectionTypes(input)
}

func (a *AutoScalingActivities) DescribeNotificationConfigurations(input *autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
    return a.client.DescribeNotificationConfigurations(input)
}

func (a *AutoScalingActivities) DescribePolicies(input *autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error) {
    return a.client.DescribePolicies(input)
}

func (a *AutoScalingActivities) DescribeScalingActivities(input *autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error) {
    return a.client.DescribeScalingActivities(input)
}

func (a *AutoScalingActivities) DescribeScalingProcessTypes(input *autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
    return a.client.DescribeScalingProcessTypes(input)
}

func (a *AutoScalingActivities) DescribeScheduledActions(input *autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error) {
    return a.client.DescribeScheduledActions(input)
}

func (a *AutoScalingActivities) DescribeTags(input *autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}

func (a *AutoScalingActivities) DescribeTerminationPolicyTypes(input *autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
    return a.client.DescribeTerminationPolicyTypes(input)
}

func (a *AutoScalingActivities) DetachInstances(input *autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error) {
    return a.client.DetachInstances(input)
}

func (a *AutoScalingActivities) DetachLoadBalancerTargetGroups(input *autoscaling.DetachLoadBalancerTargetGroupsInput) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error) {
    return a.client.DetachLoadBalancerTargetGroups(input)
}

func (a *AutoScalingActivities) DetachLoadBalancers(input *autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error) {
    return a.client.DetachLoadBalancers(input)
}

func (a *AutoScalingActivities) DisableMetricsCollection(input *autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error) {
    return a.client.DisableMetricsCollection(input)
}

func (a *AutoScalingActivities) EnableMetricsCollection(input *autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error) {
    return a.client.EnableMetricsCollection(input)
}

func (a *AutoScalingActivities) EnterStandby(input *autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error) {
    return a.client.EnterStandby(input)
}

func (a *AutoScalingActivities) ExecutePolicy(input *autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error) {
    return a.client.ExecutePolicy(input)
}

func (a *AutoScalingActivities) ExitStandby(input *autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error) {
    return a.client.ExitStandby(input)
}

func (a *AutoScalingActivities) PutLifecycleHook(input *autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error) {
    return a.client.PutLifecycleHook(input)
}

func (a *AutoScalingActivities) PutNotificationConfiguration(input *autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error) {
    return a.client.PutNotificationConfiguration(input)
}

func (a *AutoScalingActivities) PutScalingPolicy(input *autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error) {
    return a.client.PutScalingPolicy(input)
}

func (a *AutoScalingActivities) PutScheduledUpdateGroupAction(input *autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error) {
    return a.client.PutScheduledUpdateGroupAction(input)
}

func (a *AutoScalingActivities) RecordLifecycleActionHeartbeat(input *autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error) {
    return a.client.RecordLifecycleActionHeartbeat(input)
}

func (a *AutoScalingActivities) ResumeProcesses(input *autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error) {
    return a.client.ResumeProcesses(input)
}

func (a *AutoScalingActivities) SetDesiredCapacity(input *autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error) {
    return a.client.SetDesiredCapacity(input)
}

func (a *AutoScalingActivities) SetInstanceHealth(input *autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error) {
    return a.client.SetInstanceHealth(input)
}

func (a *AutoScalingActivities) SetInstanceProtection(input *autoscaling.SetInstanceProtectionInput) (*autoscaling.SetInstanceProtectionOutput, error) {
    return a.client.SetInstanceProtection(input)
}

func (a *AutoScalingActivities) StartInstanceRefresh(input *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error) {
    return a.client.StartInstanceRefresh(input)
}

func (a *AutoScalingActivities) SuspendProcesses(input *autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error) {
    return a.client.SuspendProcesses(input)
}

func (a *AutoScalingActivities) TerminateInstanceInAutoScalingGroup(input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error) {
    return a.client.TerminateInstanceInAutoScalingGroup(input)
}

func (a *AutoScalingActivities) UpdateAutoScalingGroup(input *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error) {
    return a.client.UpdateAutoScalingGroup(input)
}

func (a *AutoScalingActivities) WaitUntilGroupExists(input *autoscaling.DescribeAutoScalingGroupsInput) error {
    return a.client.WaitUntilGroupExists(input)
}

func (a *AutoScalingActivities) WaitUntilGroupInService(input *autoscaling.DescribeAutoScalingGroupsInput) error {
    return a.client.WaitUntilGroupInService(input)
}

func (a *AutoScalingActivities) WaitUntilGroupNotExists(input *autoscaling.DescribeAutoScalingGroupsInput) error {
    return a.client.WaitUntilGroupNotExists(input)
}
