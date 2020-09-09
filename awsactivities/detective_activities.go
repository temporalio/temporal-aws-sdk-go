
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/detective"
	"github.com/aws/aws-sdk-go/service/detective/detectiveiface"
)

type DetectiveActivities struct {
	client detectiveiface.DetectiveAPI
}

func NewDetectiveActivities(session *session.Session, config... *aws.Config) *DetectiveActivities {
    client := detective.New(session, config...)
    return &DetectiveActivities{client: client}
}

func (a *DetectiveActivities) AcceptInvitation(input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error) {
    return a.client.AcceptInvitation(input)
}

func (a *DetectiveActivities) CreateGraph(input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error) {
    return a.client.CreateGraph(input)
}

func (a *DetectiveActivities) CreateMembers(input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error) {
    return a.client.CreateMembers(input)
}

func (a *DetectiveActivities) DeleteGraph(input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error) {
    return a.client.DeleteGraph(input)
}

func (a *DetectiveActivities) DeleteMembers(input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error) {
    return a.client.DeleteMembers(input)
}

func (a *DetectiveActivities) DisassociateMembership(input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error) {
    return a.client.DisassociateMembership(input)
}

func (a *DetectiveActivities) GetMembers(input *detective.GetMembersInput) (*detective.GetMembersOutput, error) {
    return a.client.GetMembers(input)
}

func (a *DetectiveActivities) ListGraphs(input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error) {
    return a.client.ListGraphs(input)
}

func (a *DetectiveActivities) ListInvitations(input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error) {
    return a.client.ListInvitations(input)
}

func (a *DetectiveActivities) ListMembers(input *detective.ListMembersInput) (*detective.ListMembersOutput, error) {
    return a.client.ListMembers(input)
}

func (a *DetectiveActivities) RejectInvitation(input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error) {
    return a.client.RejectInvitation(input)
}

func (a *DetectiveActivities) StartMonitoringMember(input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error) {
    return a.client.StartMonitoringMember(input)
}
