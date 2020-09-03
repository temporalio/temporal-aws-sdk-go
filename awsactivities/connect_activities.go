package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"
)

type ConnectActivities struct {
	client connectiface.ConnectAPI
}

func NewConnectActivities(client connectiface.ConnectAPI) *ConnectActivities {
    return &ConnectActivities{client: client}
}

func (a *ConnectActivities) CreateUser(input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *ConnectActivities) DeleteUser(input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *ConnectActivities) DescribeUser(input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error) {
    return a.client.DescribeUser(input)
}

func (a *ConnectActivities) DescribeUserHierarchyGroup(input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error) {
    return a.client.DescribeUserHierarchyGroup(input)
}

func (a *ConnectActivities) DescribeUserHierarchyStructure(input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error) {
    return a.client.DescribeUserHierarchyStructure(input)
}

func (a *ConnectActivities) GetContactAttributes(input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error) {
    return a.client.GetContactAttributes(input)
}

func (a *ConnectActivities) GetCurrentMetricData(input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error) {
    return a.client.GetCurrentMetricData(input)
}

func (a *ConnectActivities) GetFederationToken(input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error) {
    return a.client.GetFederationToken(input)
}

func (a *ConnectActivities) GetMetricData(input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error) {
    return a.client.GetMetricData(input)
}

func (a *ConnectActivities) ListContactFlows(input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error) {
    return a.client.ListContactFlows(input)
}

func (a *ConnectActivities) ListHoursOfOperations(input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error) {
    return a.client.ListHoursOfOperations(input)
}

func (a *ConnectActivities) ListPhoneNumbers(input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error) {
    return a.client.ListPhoneNumbers(input)
}

func (a *ConnectActivities) ListQueues(input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error) {
    return a.client.ListQueues(input)
}

func (a *ConnectActivities) ListRoutingProfiles(input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error) {
    return a.client.ListRoutingProfiles(input)
}

func (a *ConnectActivities) ListSecurityProfiles(input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error) {
    return a.client.ListSecurityProfiles(input)
}

func (a *ConnectActivities) ListTagsForResource(input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ConnectActivities) ListUserHierarchyGroups(input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error) {
    return a.client.ListUserHierarchyGroups(input)
}

func (a *ConnectActivities) ListUsers(input *connect.ListUsersInput) (*connect.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *ConnectActivities) ResumeContactRecording(input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error) {
    return a.client.ResumeContactRecording(input)
}

func (a *ConnectActivities) StartChatContact(input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error) {
    return a.client.StartChatContact(input)
}

func (a *ConnectActivities) StartContactRecording(input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error) {
    return a.client.StartContactRecording(input)
}

func (a *ConnectActivities) StartOutboundVoiceContact(input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error) {
    return a.client.StartOutboundVoiceContact(input)
}

func (a *ConnectActivities) StopContact(input *connect.StopContactInput) (*connect.StopContactOutput, error) {
    return a.client.StopContact(input)
}

func (a *ConnectActivities) StopContactRecording(input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error) {
    return a.client.StopContactRecording(input)
}

func (a *ConnectActivities) SuspendContactRecording(input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error) {
    return a.client.SuspendContactRecording(input)
}

func (a *ConnectActivities) TagResource(input *connect.TagResourceInput) (*connect.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ConnectActivities) UntagResource(input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ConnectActivities) UpdateContactAttributes(input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error) {
    return a.client.UpdateContactAttributes(input)
}

func (a *ConnectActivities) UpdateUserHierarchy(input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error) {
    return a.client.UpdateUserHierarchy(input)
}

func (a *ConnectActivities) UpdateUserIdentityInfo(input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error) {
    return a.client.UpdateUserIdentityInfo(input)
}

func (a *ConnectActivities) UpdateUserPhoneConfig(input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error) {
    return a.client.UpdateUserPhoneConfig(input)
}

func (a *ConnectActivities) UpdateUserRoutingProfile(input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error) {
    return a.client.UpdateUserRoutingProfile(input)
}

func (a *ConnectActivities) UpdateUserSecurityProfiles(input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error) {
    return a.client.UpdateUserSecurityProfiles(input)
}
