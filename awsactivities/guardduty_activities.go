package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/guardduty/guarddutyiface"
)

type GuardDutyActivities struct {
	client guarddutyiface.GuardDutyAPI
}

func NewGuardDutyActivities(client guarddutyiface.GuardDutyAPI) *GuardDutyActivities {
    return &GuardDutyActivities{client: client}
}

func (a *GuardDutyActivities) AcceptInvitation(input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error) {
    return a.client.AcceptInvitation(input)
}

func (a *GuardDutyActivities) ArchiveFindings(input *guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error) {
    return a.client.ArchiveFindings(input)
}

func (a *GuardDutyActivities) CreateDetector(input *guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error) {
    return a.client.CreateDetector(input)
}

func (a *GuardDutyActivities) CreateFilter(input *guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error) {
    return a.client.CreateFilter(input)
}

func (a *GuardDutyActivities) CreateIPSet(input *guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error) {
    return a.client.CreateIPSet(input)
}

func (a *GuardDutyActivities) CreateMembers(input *guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error) {
    return a.client.CreateMembers(input)
}

func (a *GuardDutyActivities) CreatePublishingDestination(input *guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error) {
    return a.client.CreatePublishingDestination(input)
}

func (a *GuardDutyActivities) CreateSampleFindings(input *guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error) {
    return a.client.CreateSampleFindings(input)
}

func (a *GuardDutyActivities) CreateThreatIntelSet(input *guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error) {
    return a.client.CreateThreatIntelSet(input)
}

func (a *GuardDutyActivities) DeclineInvitations(input *guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error) {
    return a.client.DeclineInvitations(input)
}

func (a *GuardDutyActivities) DeleteDetector(input *guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error) {
    return a.client.DeleteDetector(input)
}

func (a *GuardDutyActivities) DeleteFilter(input *guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error) {
    return a.client.DeleteFilter(input)
}

func (a *GuardDutyActivities) DeleteIPSet(input *guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error) {
    return a.client.DeleteIPSet(input)
}

func (a *GuardDutyActivities) DeleteInvitations(input *guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error) {
    return a.client.DeleteInvitations(input)
}

func (a *GuardDutyActivities) DeleteMembers(input *guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error) {
    return a.client.DeleteMembers(input)
}

func (a *GuardDutyActivities) DeletePublishingDestination(input *guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error) {
    return a.client.DeletePublishingDestination(input)
}

func (a *GuardDutyActivities) DeleteThreatIntelSet(input *guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error) {
    return a.client.DeleteThreatIntelSet(input)
}

func (a *GuardDutyActivities) DescribeOrganizationConfiguration(input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
    return a.client.DescribeOrganizationConfiguration(input)
}

func (a *GuardDutyActivities) DescribePublishingDestination(input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error) {
    return a.client.DescribePublishingDestination(input)
}

func (a *GuardDutyActivities) DisableOrganizationAdminAccount(input *guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error) {
    return a.client.DisableOrganizationAdminAccount(input)
}

func (a *GuardDutyActivities) DisassociateFromMasterAccount(input *guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error) {
    return a.client.DisassociateFromMasterAccount(input)
}

func (a *GuardDutyActivities) DisassociateMembers(input *guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error) {
    return a.client.DisassociateMembers(input)
}

func (a *GuardDutyActivities) EnableOrganizationAdminAccount(input *guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error) {
    return a.client.EnableOrganizationAdminAccount(input)
}

func (a *GuardDutyActivities) GetDetector(input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error) {
    return a.client.GetDetector(input)
}

func (a *GuardDutyActivities) GetFilter(input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error) {
    return a.client.GetFilter(input)
}

func (a *GuardDutyActivities) GetFindings(input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error) {
    return a.client.GetFindings(input)
}

func (a *GuardDutyActivities) GetFindingsStatistics(input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error) {
    return a.client.GetFindingsStatistics(input)
}

func (a *GuardDutyActivities) GetIPSet(input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error) {
    return a.client.GetIPSet(input)
}

func (a *GuardDutyActivities) GetInvitationsCount(input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error) {
    return a.client.GetInvitationsCount(input)
}

func (a *GuardDutyActivities) GetMasterAccount(input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error) {
    return a.client.GetMasterAccount(input)
}

func (a *GuardDutyActivities) GetMemberDetectors(input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error) {
    return a.client.GetMemberDetectors(input)
}

func (a *GuardDutyActivities) GetMembers(input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error) {
    return a.client.GetMembers(input)
}

func (a *GuardDutyActivities) GetThreatIntelSet(input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error) {
    return a.client.GetThreatIntelSet(input)
}

func (a *GuardDutyActivities) GetUsageStatistics(input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error) {
    return a.client.GetUsageStatistics(input)
}

func (a *GuardDutyActivities) InviteMembers(input *guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error) {
    return a.client.InviteMembers(input)
}

func (a *GuardDutyActivities) ListDetectors(input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error) {
    return a.client.ListDetectors(input)
}

func (a *GuardDutyActivities) ListFilters(input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error) {
    return a.client.ListFilters(input)
}

func (a *GuardDutyActivities) ListFindings(input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error) {
    return a.client.ListFindings(input)
}

func (a *GuardDutyActivities) ListIPSets(input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error) {
    return a.client.ListIPSets(input)
}

func (a *GuardDutyActivities) ListInvitations(input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error) {
    return a.client.ListInvitations(input)
}

func (a *GuardDutyActivities) ListMembers(input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error) {
    return a.client.ListMembers(input)
}

func (a *GuardDutyActivities) ListOrganizationAdminAccounts(input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
    return a.client.ListOrganizationAdminAccounts(input)
}

func (a *GuardDutyActivities) ListPublishingDestinations(input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error) {
    return a.client.ListPublishingDestinations(input)
}

func (a *GuardDutyActivities) ListTagsForResource(input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *GuardDutyActivities) ListThreatIntelSets(input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error) {
    return a.client.ListThreatIntelSets(input)
}

func (a *GuardDutyActivities) StartMonitoringMembers(input *guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error) {
    return a.client.StartMonitoringMembers(input)
}

func (a *GuardDutyActivities) StopMonitoringMembers(input *guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error) {
    return a.client.StopMonitoringMembers(input)
}

func (a *GuardDutyActivities) TagResource(input *guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *GuardDutyActivities) UnarchiveFindings(input *guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error) {
    return a.client.UnarchiveFindings(input)
}

func (a *GuardDutyActivities) UntagResource(input *guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *GuardDutyActivities) UpdateDetector(input *guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error) {
    return a.client.UpdateDetector(input)
}

func (a *GuardDutyActivities) UpdateFilter(input *guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error) {
    return a.client.UpdateFilter(input)
}

func (a *GuardDutyActivities) UpdateFindingsFeedback(input *guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error) {
    return a.client.UpdateFindingsFeedback(input)
}

func (a *GuardDutyActivities) UpdateIPSet(input *guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error) {
    return a.client.UpdateIPSet(input)
}

func (a *GuardDutyActivities) UpdateMemberDetectors(input *guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error) {
    return a.client.UpdateMemberDetectors(input)
}

func (a *GuardDutyActivities) UpdateOrganizationConfiguration(input *guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error) {
    return a.client.UpdateOrganizationConfiguration(input)
}

func (a *GuardDutyActivities) UpdatePublishingDestination(input *guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error) {
    return a.client.UpdatePublishingDestination(input)
}

func (a *GuardDutyActivities) UpdateThreatIntelSet(input *guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error) {
    return a.client.UpdateThreatIntelSet(input)
}
