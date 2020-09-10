package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/guardduty/guarddutyiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type GuardDutyActivities struct {
	client guarddutyiface.GuardDutyAPI
}

func NewGuardDutyActivities(session *session.Session, config ...*aws.Config) *GuardDutyActivities {
	client := guardduty.New(session, config...)
	return &GuardDutyActivities{client: client}
}

func (a *GuardDutyActivities) AcceptInvitation(ctx context.Context, input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error) {
	return a.client.AcceptInvitationWithContext(ctx, input)
}

func (a *GuardDutyActivities) ArchiveFindings(ctx context.Context, input *guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error) {
	return a.client.ArchiveFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateDetector(ctx context.Context, input *guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateFilter(ctx context.Context, input *guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateIPSet(ctx context.Context, input *guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateMembers(ctx context.Context, input *guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error) {
	return a.client.CreateMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreatePublishingDestination(ctx context.Context, input *guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreatePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateSampleFindings(ctx context.Context, input *guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error) {
	return a.client.CreateSampleFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) CreateThreatIntelSet(ctx context.Context, input *guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateThreatIntelSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeclineInvitations(ctx context.Context, input *guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error) {
	return a.client.DeclineInvitationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteDetector(ctx context.Context, input *guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error) {
	return a.client.DeleteDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteFilter(ctx context.Context, input *guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error) {
	return a.client.DeleteFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteIPSet(ctx context.Context, input *guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error) {
	return a.client.DeleteIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteInvitations(ctx context.Context, input *guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error) {
	return a.client.DeleteInvitationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteMembers(ctx context.Context, input *guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error) {
	return a.client.DeleteMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeletePublishingDestination(ctx context.Context, input *guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error) {
	return a.client.DeletePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) DeleteThreatIntelSet(ctx context.Context, input *guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error) {
	return a.client.DeleteThreatIntelSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) DescribeOrganizationConfiguration(ctx context.Context, input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
	return a.client.DescribeOrganizationConfigurationWithContext(ctx, input)
}

func (a *GuardDutyActivities) DescribePublishingDestination(ctx context.Context, input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error) {
	return a.client.DescribePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) DisableOrganizationAdminAccount(ctx context.Context, input *guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error) {
	return a.client.DisableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) DisassociateFromMasterAccount(ctx context.Context, input *guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error) {
	return a.client.DisassociateFromMasterAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) DisassociateMembers(ctx context.Context, input *guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error) {
	return a.client.DisassociateMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) EnableOrganizationAdminAccount(ctx context.Context, input *guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error) {
	return a.client.EnableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetDetector(ctx context.Context, input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error) {
	return a.client.GetDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetFilter(ctx context.Context, input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error) {
	return a.client.GetFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetFindings(ctx context.Context, input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error) {
	return a.client.GetFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetFindingsStatistics(ctx context.Context, input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error) {
	return a.client.GetFindingsStatisticsWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetIPSet(ctx context.Context, input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error) {
	return a.client.GetIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetInvitationsCount(ctx context.Context, input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error) {
	return a.client.GetInvitationsCountWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetMasterAccount(ctx context.Context, input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error) {
	return a.client.GetMasterAccountWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetMemberDetectors(ctx context.Context, input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error) {
	return a.client.GetMemberDetectorsWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetMembers(ctx context.Context, input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error) {
	return a.client.GetMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetThreatIntelSet(ctx context.Context, input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error) {
	return a.client.GetThreatIntelSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) GetUsageStatistics(ctx context.Context, input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error) {
	return a.client.GetUsageStatisticsWithContext(ctx, input)
}

func (a *GuardDutyActivities) InviteMembers(ctx context.Context, input *guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error) {
	return a.client.InviteMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListDetectors(ctx context.Context, input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error) {
	return a.client.ListDetectorsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListFilters(ctx context.Context, input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error) {
	return a.client.ListFiltersWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListFindings(ctx context.Context, input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error) {
	return a.client.ListFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListIPSets(ctx context.Context, input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error) {
	return a.client.ListIPSetsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListInvitations(ctx context.Context, input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error) {
	return a.client.ListInvitationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListMembers(ctx context.Context, input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error) {
	return a.client.ListMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListOrganizationAdminAccounts(ctx context.Context, input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
	return a.client.ListOrganizationAdminAccountsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListPublishingDestinations(ctx context.Context, input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error) {
	return a.client.ListPublishingDestinationsWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListTagsForResource(ctx context.Context, input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *GuardDutyActivities) ListThreatIntelSets(ctx context.Context, input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error) {
	return a.client.ListThreatIntelSetsWithContext(ctx, input)
}

func (a *GuardDutyActivities) StartMonitoringMembers(ctx context.Context, input *guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error) {
	return a.client.StartMonitoringMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) StopMonitoringMembers(ctx context.Context, input *guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error) {
	return a.client.StopMonitoringMembersWithContext(ctx, input)
}

func (a *GuardDutyActivities) TagResource(ctx context.Context, input *guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *GuardDutyActivities) UnarchiveFindings(ctx context.Context, input *guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error) {
	return a.client.UnarchiveFindingsWithContext(ctx, input)
}

func (a *GuardDutyActivities) UntagResource(ctx context.Context, input *guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateDetector(ctx context.Context, input *guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error) {
	return a.client.UpdateDetectorWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateFilter(ctx context.Context, input *guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error) {
	return a.client.UpdateFilterWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateFindingsFeedback(ctx context.Context, input *guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error) {
	return a.client.UpdateFindingsFeedbackWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateIPSet(ctx context.Context, input *guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error) {
	return a.client.UpdateIPSetWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateMemberDetectors(ctx context.Context, input *guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error) {
	return a.client.UpdateMemberDetectorsWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateOrganizationConfiguration(ctx context.Context, input *guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error) {
	return a.client.UpdateOrganizationConfigurationWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdatePublishingDestination(ctx context.Context, input *guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error) {
	return a.client.UpdatePublishingDestinationWithContext(ctx, input)
}

func (a *GuardDutyActivities) UpdateThreatIntelSet(ctx context.Context, input *guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error) {
	return a.client.UpdateThreatIntelSetWithContext(ctx, input)
}
