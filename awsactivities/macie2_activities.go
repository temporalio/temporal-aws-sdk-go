
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/macie2"
	"github.com/aws/aws-sdk-go/service/macie2/macie2iface"
)

type Macie2Activities struct {
    client macie2iface.Macie2API
}

func NewMacie2Activities(session *session.Session, config ...*aws.Config) *Macie2Activities {
    client := macie2.New(session, config...)
    return &Macie2Activities{client: client}
}

func (a *Macie2Activities) AcceptInvitation(input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error) {
    return a.client.AcceptInvitation(input)
}

func (a *Macie2Activities) BatchGetCustomDataIdentifiers(input *macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error) {
    return a.client.BatchGetCustomDataIdentifiers(input)
}

func (a *Macie2Activities) CreateClassificationJob(input *macie2.CreateClassificationJobInput) (*macie2.CreateClassificationJobOutput, error) {
    return a.client.CreateClassificationJob(input)
}

func (a *Macie2Activities) CreateCustomDataIdentifier(input *macie2.CreateCustomDataIdentifierInput) (*macie2.CreateCustomDataIdentifierOutput, error) {
    return a.client.CreateCustomDataIdentifier(input)
}

func (a *Macie2Activities) CreateFindingsFilter(input *macie2.CreateFindingsFilterInput) (*macie2.CreateFindingsFilterOutput, error) {
    return a.client.CreateFindingsFilter(input)
}

func (a *Macie2Activities) CreateInvitations(input *macie2.CreateInvitationsInput) (*macie2.CreateInvitationsOutput, error) {
    return a.client.CreateInvitations(input)
}

func (a *Macie2Activities) CreateMember(input *macie2.CreateMemberInput) (*macie2.CreateMemberOutput, error) {
    return a.client.CreateMember(input)
}

func (a *Macie2Activities) CreateSampleFindings(input *macie2.CreateSampleFindingsInput) (*macie2.CreateSampleFindingsOutput, error) {
    return a.client.CreateSampleFindings(input)
}

func (a *Macie2Activities) DeclineInvitations(input *macie2.DeclineInvitationsInput) (*macie2.DeclineInvitationsOutput, error) {
    return a.client.DeclineInvitations(input)
}

func (a *Macie2Activities) DeleteCustomDataIdentifier(input *macie2.DeleteCustomDataIdentifierInput) (*macie2.DeleteCustomDataIdentifierOutput, error) {
    return a.client.DeleteCustomDataIdentifier(input)
}

func (a *Macie2Activities) DeleteFindingsFilter(input *macie2.DeleteFindingsFilterInput) (*macie2.DeleteFindingsFilterOutput, error) {
    return a.client.DeleteFindingsFilter(input)
}

func (a *Macie2Activities) DeleteInvitations(input *macie2.DeleteInvitationsInput) (*macie2.DeleteInvitationsOutput, error) {
    return a.client.DeleteInvitations(input)
}

func (a *Macie2Activities) DeleteMember(input *macie2.DeleteMemberInput) (*macie2.DeleteMemberOutput, error) {
    return a.client.DeleteMember(input)
}

func (a *Macie2Activities) DescribeBuckets(input *macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error) {
    return a.client.DescribeBuckets(input)
}

func (a *Macie2Activities) DescribeClassificationJob(input *macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error) {
    return a.client.DescribeClassificationJob(input)
}

func (a *Macie2Activities) DescribeOrganizationConfiguration(input *macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error) {
    return a.client.DescribeOrganizationConfiguration(input)
}

func (a *Macie2Activities) DisableMacie(input *macie2.DisableMacieInput) (*macie2.DisableMacieOutput, error) {
    return a.client.DisableMacie(input)
}

func (a *Macie2Activities) DisableOrganizationAdminAccount(input *macie2.DisableOrganizationAdminAccountInput) (*macie2.DisableOrganizationAdminAccountOutput, error) {
    return a.client.DisableOrganizationAdminAccount(input)
}

func (a *Macie2Activities) DisassociateFromMasterAccount(input *macie2.DisassociateFromMasterAccountInput) (*macie2.DisassociateFromMasterAccountOutput, error) {
    return a.client.DisassociateFromMasterAccount(input)
}

func (a *Macie2Activities) DisassociateMember(input *macie2.DisassociateMemberInput) (*macie2.DisassociateMemberOutput, error) {
    return a.client.DisassociateMember(input)
}

func (a *Macie2Activities) EnableMacie(input *macie2.EnableMacieInput) (*macie2.EnableMacieOutput, error) {
    return a.client.EnableMacie(input)
}

func (a *Macie2Activities) EnableOrganizationAdminAccount(input *macie2.EnableOrganizationAdminAccountInput) (*macie2.EnableOrganizationAdminAccountOutput, error) {
    return a.client.EnableOrganizationAdminAccount(input)
}

func (a *Macie2Activities) GetBucketStatistics(input *macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error) {
    return a.client.GetBucketStatistics(input)
}

func (a *Macie2Activities) GetClassificationExportConfiguration(input *macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error) {
    return a.client.GetClassificationExportConfiguration(input)
}

func (a *Macie2Activities) GetCustomDataIdentifier(input *macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error) {
    return a.client.GetCustomDataIdentifier(input)
}

func (a *Macie2Activities) GetFindingStatistics(input *macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error) {
    return a.client.GetFindingStatistics(input)
}

func (a *Macie2Activities) GetFindings(input *macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error) {
    return a.client.GetFindings(input)
}

func (a *Macie2Activities) GetFindingsFilter(input *macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error) {
    return a.client.GetFindingsFilter(input)
}

func (a *Macie2Activities) GetInvitationsCount(input *macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error) {
    return a.client.GetInvitationsCount(input)
}

func (a *Macie2Activities) GetMacieSession(input *macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error) {
    return a.client.GetMacieSession(input)
}

func (a *Macie2Activities) GetMasterAccount(input *macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error) {
    return a.client.GetMasterAccount(input)
}

func (a *Macie2Activities) GetMember(input *macie2.GetMemberInput) (*macie2.GetMemberOutput, error) {
    return a.client.GetMember(input)
}

func (a *Macie2Activities) GetUsageStatistics(input *macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error) {
    return a.client.GetUsageStatistics(input)
}

func (a *Macie2Activities) GetUsageTotals(input *macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error) {
    return a.client.GetUsageTotals(input)
}

func (a *Macie2Activities) ListClassificationJobs(input *macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error) {
    return a.client.ListClassificationJobs(input)
}

func (a *Macie2Activities) ListCustomDataIdentifiers(input *macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error) {
    return a.client.ListCustomDataIdentifiers(input)
}

func (a *Macie2Activities) ListFindings(input *macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error) {
    return a.client.ListFindings(input)
}

func (a *Macie2Activities) ListFindingsFilters(input *macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error) {
    return a.client.ListFindingsFilters(input)
}

func (a *Macie2Activities) ListInvitations(input *macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error) {
    return a.client.ListInvitations(input)
}

func (a *Macie2Activities) ListMembers(input *macie2.ListMembersInput) (*macie2.ListMembersOutput, error) {
    return a.client.ListMembers(input)
}

func (a *Macie2Activities) ListOrganizationAdminAccounts(input *macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error) {
    return a.client.ListOrganizationAdminAccounts(input)
}

func (a *Macie2Activities) ListTagsForResource(input *macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *Macie2Activities) PutClassificationExportConfiguration(input *macie2.PutClassificationExportConfigurationInput) (*macie2.PutClassificationExportConfigurationOutput, error) {
    return a.client.PutClassificationExportConfiguration(input)
}

func (a *Macie2Activities) TagResource(input *macie2.TagResourceInput) (*macie2.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *Macie2Activities) TestCustomDataIdentifier(input *macie2.TestCustomDataIdentifierInput) (*macie2.TestCustomDataIdentifierOutput, error) {
    return a.client.TestCustomDataIdentifier(input)
}

func (a *Macie2Activities) UntagResource(input *macie2.UntagResourceInput) (*macie2.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *Macie2Activities) UpdateClassificationJob(input *macie2.UpdateClassificationJobInput) (*macie2.UpdateClassificationJobOutput, error) {
    return a.client.UpdateClassificationJob(input)
}

func (a *Macie2Activities) UpdateFindingsFilter(input *macie2.UpdateFindingsFilterInput) (*macie2.UpdateFindingsFilterOutput, error) {
    return a.client.UpdateFindingsFilter(input)
}

func (a *Macie2Activities) UpdateMacieSession(input *macie2.UpdateMacieSessionInput) (*macie2.UpdateMacieSessionOutput, error) {
    return a.client.UpdateMacieSession(input)
}

func (a *Macie2Activities) UpdateMemberSession(input *macie2.UpdateMemberSessionInput) (*macie2.UpdateMemberSessionOutput, error) {
    return a.client.UpdateMemberSession(input)
}

func (a *Macie2Activities) UpdateOrganizationConfiguration(input *macie2.UpdateOrganizationConfigurationInput) (*macie2.UpdateOrganizationConfigurationOutput, error) {
    return a.client.UpdateOrganizationConfiguration(input)
}
