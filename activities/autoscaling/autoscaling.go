// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client autoscalingiface.AutoScalingAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := autoscaling.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (autoscalingiface.AutoScalingAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return autoscaling.New(sess), nil
}

func (a *Activities) AttachInstances(ctx context.Context, input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AttachInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AttachLoadBalancerTargetGroups(ctx context.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AttachLoadBalancerTargetGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AttachLoadBalancers(ctx context.Context, input *autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AttachLoadBalancersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDeleteScheduledAction(ctx context.Context, input *autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDeleteScheduledActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchPutScheduledUpdateGroupAction(ctx context.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchPutScheduledUpdateGroupActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CancelInstanceRefresh(ctx context.Context, input *autoscaling.CancelInstanceRefreshInput) (*autoscaling.CancelInstanceRefreshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelInstanceRefreshWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CompleteLifecycleAction(ctx context.Context, input *autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CompleteLifecycleActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateAutoScalingGroup(ctx context.Context, input *autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateAutoScalingGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLaunchConfiguration(ctx context.Context, input *autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLaunchConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateOrUpdateTags(ctx context.Context, input *autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateOrUpdateTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAutoScalingGroup(ctx context.Context, input *autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAutoScalingGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLaunchConfiguration(ctx context.Context, input *autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLaunchConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLifecycleHook(ctx context.Context, input *autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLifecycleHookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteNotificationConfiguration(ctx context.Context, input *autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteNotificationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePolicy(ctx context.Context, input *autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteScheduledAction(ctx context.Context, input *autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteScheduledActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTags(ctx context.Context, input *autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAccountLimits(ctx context.Context, input *autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAccountLimitsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAdjustmentTypes(ctx context.Context, input *autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAdjustmentTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAutoScalingGroups(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAutoScalingGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAutoScalingInstances(ctx context.Context, input *autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAutoScalingInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAutoScalingNotificationTypes(ctx context.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAutoScalingNotificationTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInstanceRefreshes(ctx context.Context, input *autoscaling.DescribeInstanceRefreshesInput) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInstanceRefreshesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLaunchConfigurations(ctx context.Context, input *autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLaunchConfigurationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLifecycleHookTypes(ctx context.Context, input *autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLifecycleHookTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLifecycleHooks(ctx context.Context, input *autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLifecycleHooksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancerTargetGroups(ctx context.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancerTargetGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancers(ctx context.Context, input *autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMetricCollectionTypes(ctx context.Context, input *autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMetricCollectionTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeNotificationConfigurations(ctx context.Context, input *autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeNotificationConfigurationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePolicies(ctx context.Context, input *autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePoliciesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeScalingActivities(ctx context.Context, input *autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScalingActivitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeScalingProcessTypes(ctx context.Context, input *autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScalingProcessTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeScheduledActions(ctx context.Context, input *autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScheduledActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTags(ctx context.Context, input *autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTerminationPolicyTypes(ctx context.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTerminationPolicyTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetachInstances(ctx context.Context, input *autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetachInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetachLoadBalancerTargetGroups(ctx context.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetachLoadBalancerTargetGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetachLoadBalancers(ctx context.Context, input *autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetachLoadBalancersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisableMetricsCollection(ctx context.Context, input *autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisableMetricsCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) EnableMetricsCollection(ctx context.Context, input *autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.EnableMetricsCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) EnterStandby(ctx context.Context, input *autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.EnterStandbyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExecutePolicy(ctx context.Context, input *autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExecutePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExitStandby(ctx context.Context, input *autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExitStandbyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutLifecycleHook(ctx context.Context, input *autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutLifecycleHookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutNotificationConfiguration(ctx context.Context, input *autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutNotificationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutScalingPolicy(ctx context.Context, input *autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutScalingPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutScheduledUpdateGroupAction(ctx context.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutScheduledUpdateGroupActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RecordLifecycleActionHeartbeat(ctx context.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RecordLifecycleActionHeartbeatWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResumeProcesses(ctx context.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResumeProcessesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetDesiredCapacity(ctx context.Context, input *autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetDesiredCapacityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetInstanceHealth(ctx context.Context, input *autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetInstanceHealthWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetInstanceProtection(ctx context.Context, input *autoscaling.SetInstanceProtectionInput) (*autoscaling.SetInstanceProtectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetInstanceProtectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartInstanceRefresh(ctx context.Context, input *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartInstanceRefreshWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SuspendProcesses(ctx context.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SuspendProcessesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TerminateInstanceInAutoScalingGroup(ctx context.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TerminateInstanceInAutoScalingGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateAutoScalingGroup(ctx context.Context, input *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateAutoScalingGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilGroupExists(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilGroupExistsWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilGroupInService(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilGroupInServiceWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilGroupNotExists(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilGroupNotExistsWithContext(ctx, input, options...))
	})
}
