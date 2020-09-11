
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/macie2"
	"github.com/aws/aws-sdk-go/service/macie2/macie2iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type Macie2Activities struct {
    client macie2iface.Macie2API
}

func NewMacie2Activities(session *session.Session, config ...*aws.Config) *Macie2Activities {
    client := macie2.New(session, config...)
    return &Macie2Activities{client: client}
}

func (a *Macie2Activities) AcceptInvitation(ctx context.Context, input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error) {
    return a.client.AcceptInvitationWithContext(ctx, input)
}

func (a *Macie2Activities) BatchGetCustomDataIdentifiers(ctx context.Context, input *macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error) {
    return a.client.BatchGetCustomDataIdentifiersWithContext(ctx, input)
}

func (a *Macie2Activities) CreateClassificationJob(ctx context.Context, input *macie2.CreateClassificationJobInput) (*macie2.CreateClassificationJobOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateClassificationJobWithContext(ctx, input)
}

func (a *Macie2Activities) CreateCustomDataIdentifier(ctx context.Context, input *macie2.CreateCustomDataIdentifierInput) (*macie2.CreateCustomDataIdentifierOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateCustomDataIdentifierWithContext(ctx, input)
}

func (a *Macie2Activities) CreateFindingsFilter(ctx context.Context, input *macie2.CreateFindingsFilterInput) (*macie2.CreateFindingsFilterOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateFindingsFilterWithContext(ctx, input)
}

func (a *Macie2Activities) CreateInvitations(ctx context.Context, input *macie2.CreateInvitationsInput) (*macie2.CreateInvitationsOutput, error) {
    return a.client.CreateInvitationsWithContext(ctx, input)
}

func (a *Macie2Activities) CreateMember(ctx context.Context, input *macie2.CreateMemberInput) (*macie2.CreateMemberOutput, error) {
    return a.client.CreateMemberWithContext(ctx, input)
}

func (a *Macie2Activities) CreateSampleFindings(ctx context.Context, input *macie2.CreateSampleFindingsInput) (*macie2.CreateSampleFindingsOutput, error) {
    return a.client.CreateSampleFindingsWithContext(ctx, input)
}

func (a *Macie2Activities) DeclineInvitations(ctx context.Context, input *macie2.DeclineInvitationsInput) (*macie2.DeclineInvitationsOutput, error) {
    return a.client.DeclineInvitationsWithContext(ctx, input)
}

func (a *Macie2Activities) DeleteCustomDataIdentifier(ctx context.Context, input *macie2.DeleteCustomDataIdentifierInput) (*macie2.DeleteCustomDataIdentifierOutput, error) {
    return a.client.DeleteCustomDataIdentifierWithContext(ctx, input)
}

func (a *Macie2Activities) DeleteFindingsFilter(ctx context.Context, input *macie2.DeleteFindingsFilterInput) (*macie2.DeleteFindingsFilterOutput, error) {
    return a.client.DeleteFindingsFilterWithContext(ctx, input)
}

func (a *Macie2Activities) DeleteInvitations(ctx context.Context, input *macie2.DeleteInvitationsInput) (*macie2.DeleteInvitationsOutput, error) {
    return a.client.DeleteInvitationsWithContext(ctx, input)
}

func (a *Macie2Activities) DeleteMember(ctx context.Context, input *macie2.DeleteMemberInput) (*macie2.DeleteMemberOutput, error) {
    return a.client.DeleteMemberWithContext(ctx, input)
}

func (a *Macie2Activities) DescribeBuckets(ctx context.Context, input *macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error) {
    return a.client.DescribeBucketsWithContext(ctx, input)
}

func (a *Macie2Activities) DescribeClassificationJob(ctx context.Context, input *macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error) {
    return a.client.DescribeClassificationJobWithContext(ctx, input)
}

func (a *Macie2Activities) DescribeOrganizationConfiguration(ctx context.Context, input *macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error) {
    return a.client.DescribeOrganizationConfigurationWithContext(ctx, input)
}

func (a *Macie2Activities) DisableMacie(ctx context.Context, input *macie2.DisableMacieInput) (*macie2.DisableMacieOutput, error) {
    return a.client.DisableMacieWithContext(ctx, input)
}

func (a *Macie2Activities) DisableOrganizationAdminAccount(ctx context.Context, input *macie2.DisableOrganizationAdminAccountInput) (*macie2.DisableOrganizationAdminAccountOutput, error) {
    return a.client.DisableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *Macie2Activities) DisassociateFromMasterAccount(ctx context.Context, input *macie2.DisassociateFromMasterAccountInput) (*macie2.DisassociateFromMasterAccountOutput, error) {
    return a.client.DisassociateFromMasterAccountWithContext(ctx, input)
}

func (a *Macie2Activities) DisassociateMember(ctx context.Context, input *macie2.DisassociateMemberInput) (*macie2.DisassociateMemberOutput, error) {
    return a.client.DisassociateMemberWithContext(ctx, input)
}

func (a *Macie2Activities) EnableMacie(ctx context.Context, input *macie2.EnableMacieInput) (*macie2.EnableMacieOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.EnableMacieWithContext(ctx, input)
}

func (a *Macie2Activities) EnableOrganizationAdminAccount(ctx context.Context, input *macie2.EnableOrganizationAdminAccountInput) (*macie2.EnableOrganizationAdminAccountOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.EnableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *Macie2Activities) GetBucketStatistics(ctx context.Context, input *macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error) {
    return a.client.GetBucketStatisticsWithContext(ctx, input)
}

func (a *Macie2Activities) GetClassificationExportConfiguration(ctx context.Context, input *macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error) {
    return a.client.GetClassificationExportConfigurationWithContext(ctx, input)
}

func (a *Macie2Activities) GetCustomDataIdentifier(ctx context.Context, input *macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error) {
    return a.client.GetCustomDataIdentifierWithContext(ctx, input)
}

func (a *Macie2Activities) GetFindingStatistics(ctx context.Context, input *macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error) {
    return a.client.GetFindingStatisticsWithContext(ctx, input)
}

func (a *Macie2Activities) GetFindings(ctx context.Context, input *macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error) {
    return a.client.GetFindingsWithContext(ctx, input)
}

func (a *Macie2Activities) GetFindingsFilter(ctx context.Context, input *macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error) {
    return a.client.GetFindingsFilterWithContext(ctx, input)
}

func (a *Macie2Activities) GetInvitationsCount(ctx context.Context, input *macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error) {
    return a.client.GetInvitationsCountWithContext(ctx, input)
}

func (a *Macie2Activities) GetMacieSession(ctx context.Context, input *macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error) {
    return a.client.GetMacieSessionWithContext(ctx, input)
}

func (a *Macie2Activities) GetMasterAccount(ctx context.Context, input *macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error) {
    return a.client.GetMasterAccountWithContext(ctx, input)
}

func (a *Macie2Activities) GetMember(ctx context.Context, input *macie2.GetMemberInput) (*macie2.GetMemberOutput, error) {
    return a.client.GetMemberWithContext(ctx, input)
}

func (a *Macie2Activities) GetUsageStatistics(ctx context.Context, input *macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error) {
    return a.client.GetUsageStatisticsWithContext(ctx, input)
}

func (a *Macie2Activities) GetUsageTotals(ctx context.Context, input *macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error) {
    return a.client.GetUsageTotalsWithContext(ctx, input)
}

func (a *Macie2Activities) ListClassificationJobs(ctx context.Context, input *macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error) {
    return a.client.ListClassificationJobsWithContext(ctx, input)
}

func (a *Macie2Activities) ListCustomDataIdentifiers(ctx context.Context, input *macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error) {
    return a.client.ListCustomDataIdentifiersWithContext(ctx, input)
}

func (a *Macie2Activities) ListFindings(ctx context.Context, input *macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error) {
    return a.client.ListFindingsWithContext(ctx, input)
}

func (a *Macie2Activities) ListFindingsFilters(ctx context.Context, input *macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error) {
    return a.client.ListFindingsFiltersWithContext(ctx, input)
}

func (a *Macie2Activities) ListInvitations(ctx context.Context, input *macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error) {
    return a.client.ListInvitationsWithContext(ctx, input)
}

func (a *Macie2Activities) ListMembers(ctx context.Context, input *macie2.ListMembersInput) (*macie2.ListMembersOutput, error) {
    return a.client.ListMembersWithContext(ctx, input)
}

func (a *Macie2Activities) ListOrganizationAdminAccounts(ctx context.Context, input *macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error) {
    return a.client.ListOrganizationAdminAccountsWithContext(ctx, input)
}

func (a *Macie2Activities) ListTagsForResource(ctx context.Context, input *macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Macie2Activities) PutClassificationExportConfiguration(ctx context.Context, input *macie2.PutClassificationExportConfigurationInput) (*macie2.PutClassificationExportConfigurationOutput, error) {
    return a.client.PutClassificationExportConfigurationWithContext(ctx, input)
}

func (a *Macie2Activities) TagResource(ctx context.Context, input *macie2.TagResourceInput) (*macie2.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *Macie2Activities) TestCustomDataIdentifier(ctx context.Context, input *macie2.TestCustomDataIdentifierInput) (*macie2.TestCustomDataIdentifierOutput, error) {
    return a.client.TestCustomDataIdentifierWithContext(ctx, input)
}

func (a *Macie2Activities) UntagResource(ctx context.Context, input *macie2.UntagResourceInput) (*macie2.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *Macie2Activities) UpdateClassificationJob(ctx context.Context, input *macie2.UpdateClassificationJobInput) (*macie2.UpdateClassificationJobOutput, error) {
    return a.client.UpdateClassificationJobWithContext(ctx, input)
}

func (a *Macie2Activities) UpdateFindingsFilter(ctx context.Context, input *macie2.UpdateFindingsFilterInput) (*macie2.UpdateFindingsFilterOutput, error) {
    return a.client.UpdateFindingsFilterWithContext(ctx, input)
}

func (a *Macie2Activities) UpdateMacieSession(ctx context.Context, input *macie2.UpdateMacieSessionInput) (*macie2.UpdateMacieSessionOutput, error) {
    return a.client.UpdateMacieSessionWithContext(ctx, input)
}

func (a *Macie2Activities) UpdateMemberSession(ctx context.Context, input *macie2.UpdateMemberSessionInput) (*macie2.UpdateMemberSessionOutput, error) {
    return a.client.UpdateMemberSessionWithContext(ctx, input)
}

func (a *Macie2Activities) UpdateOrganizationConfiguration(ctx context.Context, input *macie2.UpdateOrganizationConfigurationInput) (*macie2.UpdateOrganizationConfigurationOutput, error) {
    return a.client.UpdateOrganizationConfigurationWithContext(ctx, input)
}
