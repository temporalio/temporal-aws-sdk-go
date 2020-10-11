// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ConnectActivities struct {
	client connectiface.ConnectAPI

	sessionFactory SessionFactory
}

func NewConnectActivities(sess *session.Session, config ...*aws.Config) *ConnectActivities {
	client := connect.New(sess, config...)
	return &ConnectActivities{client: client}
}

func NewConnectActivitiesWithSessionFactory(sessionFactory SessionFactory) *ConnectActivities {
	return &ConnectActivities{sessionFactory: sessionFactory}
}

func (a *ConnectActivities) getClient(ctx context.Context) (connectiface.ConnectAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return connect.New(sess), nil
}

func (a *ConnectActivities) CreateUser(ctx context.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserWithContext(ctx, input)
}

func (a *ConnectActivities) DeleteUser(ctx context.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserWithContext(ctx, input)
}

func (a *ConnectActivities) DescribeUser(ctx context.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserWithContext(ctx, input)
}

func (a *ConnectActivities) DescribeUserHierarchyGroup(ctx context.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserHierarchyGroupWithContext(ctx, input)
}

func (a *ConnectActivities) DescribeUserHierarchyStructure(ctx context.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserHierarchyStructureWithContext(ctx, input)
}

func (a *ConnectActivities) GetContactAttributes(ctx context.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetContactAttributesWithContext(ctx, input)
}

func (a *ConnectActivities) GetCurrentMetricData(ctx context.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCurrentMetricDataWithContext(ctx, input)
}

func (a *ConnectActivities) GetFederationToken(ctx context.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFederationTokenWithContext(ctx, input)
}

func (a *ConnectActivities) GetMetricData(ctx context.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMetricDataWithContext(ctx, input)
}

func (a *ConnectActivities) ListContactFlows(ctx context.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListContactFlowsWithContext(ctx, input)
}

func (a *ConnectActivities) ListHoursOfOperations(ctx context.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListHoursOfOperationsWithContext(ctx, input)
}

func (a *ConnectActivities) ListPhoneNumbers(ctx context.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPhoneNumbersWithContext(ctx, input)
}

func (a *ConnectActivities) ListQueues(ctx context.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListQueuesWithContext(ctx, input)
}

func (a *ConnectActivities) ListRoutingProfiles(ctx context.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRoutingProfilesWithContext(ctx, input)
}

func (a *ConnectActivities) ListSecurityProfiles(ctx context.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSecurityProfilesWithContext(ctx, input)
}

func (a *ConnectActivities) ListTagsForResource(ctx context.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ConnectActivities) ListUserHierarchyGroups(ctx context.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUserHierarchyGroupsWithContext(ctx, input)
}

func (a *ConnectActivities) ListUsers(ctx context.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUsersWithContext(ctx, input)
}

func (a *ConnectActivities) ResumeContactRecording(ctx context.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResumeContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) StartChatContact(ctx context.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.StartChatContactWithContext(ctx, input)
}

func (a *ConnectActivities) StartContactRecording(ctx context.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) StartOutboundVoiceContact(ctx context.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.StartOutboundVoiceContactWithContext(ctx, input)
}

func (a *ConnectActivities) StopContact(ctx context.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopContactWithContext(ctx, input)
}

func (a *ConnectActivities) StopContactRecording(ctx context.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) SuspendContactRecording(ctx context.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SuspendContactRecordingWithContext(ctx, input)
}

func (a *ConnectActivities) TagResource(ctx context.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *ConnectActivities) UntagResource(ctx context.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateContactAttributes(ctx context.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateContactAttributesWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserHierarchy(ctx context.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserHierarchyWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserIdentityInfo(ctx context.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserIdentityInfoWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserPhoneConfig(ctx context.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserPhoneConfigWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserRoutingProfile(ctx context.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserRoutingProfileWithContext(ctx, input)
}

func (a *ConnectActivities) UpdateUserSecurityProfiles(ctx context.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserSecurityProfilesWithContext(ctx, input)
}
