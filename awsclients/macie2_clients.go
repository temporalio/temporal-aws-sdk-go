package awsclients

import (
	"github.com/aws/aws-sdk-go/service/macie2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type Macie2Client interface {
    AcceptInvitation(ctx workflow.Context, input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error)
    AcceptInvitationAsync(ctx workflow.Context, input *macie2.AcceptInvitationInput) *Macie2AcceptInvitationResult

    BatchGetCustomDataIdentifiers(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error)
    BatchGetCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) *Macie2BatchGetCustomDataIdentifiersResult

    CreateClassificationJob(ctx workflow.Context, input *macie2.CreateClassificationJobInput) (*macie2.CreateClassificationJobOutput, error)
    CreateClassificationJobAsync(ctx workflow.Context, input *macie2.CreateClassificationJobInput) *Macie2CreateClassificationJobResult

    CreateCustomDataIdentifier(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) (*macie2.CreateCustomDataIdentifierOutput, error)
    CreateCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) *Macie2CreateCustomDataIdentifierResult

    CreateFindingsFilter(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) (*macie2.CreateFindingsFilterOutput, error)
    CreateFindingsFilterAsync(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) *Macie2CreateFindingsFilterResult

    CreateInvitations(ctx workflow.Context, input *macie2.CreateInvitationsInput) (*macie2.CreateInvitationsOutput, error)
    CreateInvitationsAsync(ctx workflow.Context, input *macie2.CreateInvitationsInput) *Macie2CreateInvitationsResult

    CreateMember(ctx workflow.Context, input *macie2.CreateMemberInput) (*macie2.CreateMemberOutput, error)
    CreateMemberAsync(ctx workflow.Context, input *macie2.CreateMemberInput) *Macie2CreateMemberResult

    CreateSampleFindings(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) (*macie2.CreateSampleFindingsOutput, error)
    CreateSampleFindingsAsync(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) *Macie2CreateSampleFindingsResult

    DeclineInvitations(ctx workflow.Context, input *macie2.DeclineInvitationsInput) (*macie2.DeclineInvitationsOutput, error)
    DeclineInvitationsAsync(ctx workflow.Context, input *macie2.DeclineInvitationsInput) *Macie2DeclineInvitationsResult

    DeleteCustomDataIdentifier(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) (*macie2.DeleteCustomDataIdentifierOutput, error)
    DeleteCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) *Macie2DeleteCustomDataIdentifierResult

    DeleteFindingsFilter(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) (*macie2.DeleteFindingsFilterOutput, error)
    DeleteFindingsFilterAsync(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) *Macie2DeleteFindingsFilterResult

    DeleteInvitations(ctx workflow.Context, input *macie2.DeleteInvitationsInput) (*macie2.DeleteInvitationsOutput, error)
    DeleteInvitationsAsync(ctx workflow.Context, input *macie2.DeleteInvitationsInput) *Macie2DeleteInvitationsResult

    DeleteMember(ctx workflow.Context, input *macie2.DeleteMemberInput) (*macie2.DeleteMemberOutput, error)
    DeleteMemberAsync(ctx workflow.Context, input *macie2.DeleteMemberInput) *Macie2DeleteMemberResult

    DescribeBuckets(ctx workflow.Context, input *macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error)
    DescribeBucketsAsync(ctx workflow.Context, input *macie2.DescribeBucketsInput) *Macie2DescribeBucketsResult

    DescribeClassificationJob(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error)
    DescribeClassificationJobAsync(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) *Macie2DescribeClassificationJobResult

    DescribeOrganizationConfiguration(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error)
    DescribeOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) *Macie2DescribeOrganizationConfigurationResult

    DisableMacie(ctx workflow.Context, input *macie2.DisableMacieInput) (*macie2.DisableMacieOutput, error)
    DisableMacieAsync(ctx workflow.Context, input *macie2.DisableMacieInput) *Macie2DisableMacieResult

    DisableOrganizationAdminAccount(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) (*macie2.DisableOrganizationAdminAccountOutput, error)
    DisableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) *Macie2DisableOrganizationAdminAccountResult

    DisassociateFromMasterAccount(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) (*macie2.DisassociateFromMasterAccountOutput, error)
    DisassociateFromMasterAccountAsync(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) *Macie2DisassociateFromMasterAccountResult

    DisassociateMember(ctx workflow.Context, input *macie2.DisassociateMemberInput) (*macie2.DisassociateMemberOutput, error)
    DisassociateMemberAsync(ctx workflow.Context, input *macie2.DisassociateMemberInput) *Macie2DisassociateMemberResult

    EnableMacie(ctx workflow.Context, input *macie2.EnableMacieInput) (*macie2.EnableMacieOutput, error)
    EnableMacieAsync(ctx workflow.Context, input *macie2.EnableMacieInput) *Macie2EnableMacieResult

    EnableOrganizationAdminAccount(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) (*macie2.EnableOrganizationAdminAccountOutput, error)
    EnableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) *Macie2EnableOrganizationAdminAccountResult

    GetBucketStatistics(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error)
    GetBucketStatisticsAsync(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) *Macie2GetBucketStatisticsResult

    GetClassificationExportConfiguration(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error)
    GetClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) *Macie2GetClassificationExportConfigurationResult

    GetCustomDataIdentifier(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error)
    GetCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) *Macie2GetCustomDataIdentifierResult

    GetFindingStatistics(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error)
    GetFindingStatisticsAsync(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) *Macie2GetFindingStatisticsResult

    GetFindings(ctx workflow.Context, input *macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error)
    GetFindingsAsync(ctx workflow.Context, input *macie2.GetFindingsInput) *Macie2GetFindingsResult

    GetFindingsFilter(ctx workflow.Context, input *macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error)
    GetFindingsFilterAsync(ctx workflow.Context, input *macie2.GetFindingsFilterInput) *Macie2GetFindingsFilterResult

    GetInvitationsCount(ctx workflow.Context, input *macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error)
    GetInvitationsCountAsync(ctx workflow.Context, input *macie2.GetInvitationsCountInput) *Macie2GetInvitationsCountResult

    GetMacieSession(ctx workflow.Context, input *macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error)
    GetMacieSessionAsync(ctx workflow.Context, input *macie2.GetMacieSessionInput) *Macie2GetMacieSessionResult

    GetMasterAccount(ctx workflow.Context, input *macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error)
    GetMasterAccountAsync(ctx workflow.Context, input *macie2.GetMasterAccountInput) *Macie2GetMasterAccountResult

    GetMember(ctx workflow.Context, input *macie2.GetMemberInput) (*macie2.GetMemberOutput, error)
    GetMemberAsync(ctx workflow.Context, input *macie2.GetMemberInput) *Macie2GetMemberResult

    GetUsageStatistics(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error)
    GetUsageStatisticsAsync(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) *Macie2GetUsageStatisticsResult

    GetUsageTotals(ctx workflow.Context, input *macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error)
    GetUsageTotalsAsync(ctx workflow.Context, input *macie2.GetUsageTotalsInput) *Macie2GetUsageTotalsResult

    ListClassificationJobs(ctx workflow.Context, input *macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error)
    ListClassificationJobsAsync(ctx workflow.Context, input *macie2.ListClassificationJobsInput) *Macie2ListClassificationJobsResult

    ListCustomDataIdentifiers(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error)
    ListCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) *Macie2ListCustomDataIdentifiersResult

    ListFindings(ctx workflow.Context, input *macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error)
    ListFindingsAsync(ctx workflow.Context, input *macie2.ListFindingsInput) *Macie2ListFindingsResult

    ListFindingsFilters(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error)
    ListFindingsFiltersAsync(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) *Macie2ListFindingsFiltersResult

    ListInvitations(ctx workflow.Context, input *macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error)
    ListInvitationsAsync(ctx workflow.Context, input *macie2.ListInvitationsInput) *Macie2ListInvitationsResult

    ListMembers(ctx workflow.Context, input *macie2.ListMembersInput) (*macie2.ListMembersOutput, error)
    ListMembersAsync(ctx workflow.Context, input *macie2.ListMembersInput) *Macie2ListMembersResult

    ListOrganizationAdminAccounts(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error)
    ListOrganizationAdminAccountsAsync(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) *Macie2ListOrganizationAdminAccountsResult

    ListTagsForResource(ctx workflow.Context, input *macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *macie2.ListTagsForResourceInput) *Macie2ListTagsForResourceResult

    PutClassificationExportConfiguration(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) (*macie2.PutClassificationExportConfigurationOutput, error)
    PutClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) *Macie2PutClassificationExportConfigurationResult

    TagResource(ctx workflow.Context, input *macie2.TagResourceInput) (*macie2.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *macie2.TagResourceInput) *Macie2TagResourceResult

    TestCustomDataIdentifier(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) (*macie2.TestCustomDataIdentifierOutput, error)
    TestCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) *Macie2TestCustomDataIdentifierResult

    UntagResource(ctx workflow.Context, input *macie2.UntagResourceInput) (*macie2.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *macie2.UntagResourceInput) *Macie2UntagResourceResult

    UpdateClassificationJob(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) (*macie2.UpdateClassificationJobOutput, error)
    UpdateClassificationJobAsync(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) *Macie2UpdateClassificationJobResult

    UpdateFindingsFilter(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) (*macie2.UpdateFindingsFilterOutput, error)
    UpdateFindingsFilterAsync(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) *Macie2UpdateFindingsFilterResult

    UpdateMacieSession(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) (*macie2.UpdateMacieSessionOutput, error)
    UpdateMacieSessionAsync(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) *Macie2UpdateMacieSessionResult

    UpdateMemberSession(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) (*macie2.UpdateMemberSessionOutput, error)
    UpdateMemberSessionAsync(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) *Macie2UpdateMemberSessionResult

    UpdateOrganizationConfiguration(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) (*macie2.UpdateOrganizationConfigurationOutput, error)
    UpdateOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) *Macie2UpdateOrganizationConfigurationResult
}

type Macie2AcceptInvitationResult struct {
	Result workflow.Future
}

func (r *Macie2AcceptInvitationResult) Get(ctx workflow.Context) (*macie2.AcceptInvitationOutput, error) {
    var output macie2.AcceptInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2BatchGetCustomDataIdentifiersResult struct {
	Result workflow.Future
}

func (r *Macie2BatchGetCustomDataIdentifiersResult) Get(ctx workflow.Context) (*macie2.BatchGetCustomDataIdentifiersOutput, error) {
    var output macie2.BatchGetCustomDataIdentifiersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2CreateClassificationJobResult struct {
	Result workflow.Future
}

func (r *Macie2CreateClassificationJobResult) Get(ctx workflow.Context) (*macie2.CreateClassificationJobOutput, error) {
    var output macie2.CreateClassificationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2CreateCustomDataIdentifierResult struct {
	Result workflow.Future
}

func (r *Macie2CreateCustomDataIdentifierResult) Get(ctx workflow.Context) (*macie2.CreateCustomDataIdentifierOutput, error) {
    var output macie2.CreateCustomDataIdentifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2CreateFindingsFilterResult struct {
	Result workflow.Future
}

func (r *Macie2CreateFindingsFilterResult) Get(ctx workflow.Context) (*macie2.CreateFindingsFilterOutput, error) {
    var output macie2.CreateFindingsFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2CreateInvitationsResult struct {
	Result workflow.Future
}

func (r *Macie2CreateInvitationsResult) Get(ctx workflow.Context) (*macie2.CreateInvitationsOutput, error) {
    var output macie2.CreateInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2CreateMemberResult struct {
	Result workflow.Future
}

func (r *Macie2CreateMemberResult) Get(ctx workflow.Context) (*macie2.CreateMemberOutput, error) {
    var output macie2.CreateMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2CreateSampleFindingsResult struct {
	Result workflow.Future
}

func (r *Macie2CreateSampleFindingsResult) Get(ctx workflow.Context) (*macie2.CreateSampleFindingsOutput, error) {
    var output macie2.CreateSampleFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DeclineInvitationsResult struct {
	Result workflow.Future
}

func (r *Macie2DeclineInvitationsResult) Get(ctx workflow.Context) (*macie2.DeclineInvitationsOutput, error) {
    var output macie2.DeclineInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DeleteCustomDataIdentifierResult struct {
	Result workflow.Future
}

func (r *Macie2DeleteCustomDataIdentifierResult) Get(ctx workflow.Context) (*macie2.DeleteCustomDataIdentifierOutput, error) {
    var output macie2.DeleteCustomDataIdentifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DeleteFindingsFilterResult struct {
	Result workflow.Future
}

func (r *Macie2DeleteFindingsFilterResult) Get(ctx workflow.Context) (*macie2.DeleteFindingsFilterOutput, error) {
    var output macie2.DeleteFindingsFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DeleteInvitationsResult struct {
	Result workflow.Future
}

func (r *Macie2DeleteInvitationsResult) Get(ctx workflow.Context) (*macie2.DeleteInvitationsOutput, error) {
    var output macie2.DeleteInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DeleteMemberResult struct {
	Result workflow.Future
}

func (r *Macie2DeleteMemberResult) Get(ctx workflow.Context) (*macie2.DeleteMemberOutput, error) {
    var output macie2.DeleteMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DescribeBucketsResult struct {
	Result workflow.Future
}

func (r *Macie2DescribeBucketsResult) Get(ctx workflow.Context) (*macie2.DescribeBucketsOutput, error) {
    var output macie2.DescribeBucketsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DescribeClassificationJobResult struct {
	Result workflow.Future
}

func (r *Macie2DescribeClassificationJobResult) Get(ctx workflow.Context) (*macie2.DescribeClassificationJobOutput, error) {
    var output macie2.DescribeClassificationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DescribeOrganizationConfigurationResult struct {
	Result workflow.Future
}

func (r *Macie2DescribeOrganizationConfigurationResult) Get(ctx workflow.Context) (*macie2.DescribeOrganizationConfigurationOutput, error) {
    var output macie2.DescribeOrganizationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DisableMacieResult struct {
	Result workflow.Future
}

func (r *Macie2DisableMacieResult) Get(ctx workflow.Context) (*macie2.DisableMacieOutput, error) {
    var output macie2.DisableMacieOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DisableOrganizationAdminAccountResult struct {
	Result workflow.Future
}

func (r *Macie2DisableOrganizationAdminAccountResult) Get(ctx workflow.Context) (*macie2.DisableOrganizationAdminAccountOutput, error) {
    var output macie2.DisableOrganizationAdminAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DisassociateFromMasterAccountResult struct {
	Result workflow.Future
}

func (r *Macie2DisassociateFromMasterAccountResult) Get(ctx workflow.Context) (*macie2.DisassociateFromMasterAccountOutput, error) {
    var output macie2.DisassociateFromMasterAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2DisassociateMemberResult struct {
	Result workflow.Future
}

func (r *Macie2DisassociateMemberResult) Get(ctx workflow.Context) (*macie2.DisassociateMemberOutput, error) {
    var output macie2.DisassociateMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2EnableMacieResult struct {
	Result workflow.Future
}

func (r *Macie2EnableMacieResult) Get(ctx workflow.Context) (*macie2.EnableMacieOutput, error) {
    var output macie2.EnableMacieOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2EnableOrganizationAdminAccountResult struct {
	Result workflow.Future
}

func (r *Macie2EnableOrganizationAdminAccountResult) Get(ctx workflow.Context) (*macie2.EnableOrganizationAdminAccountOutput, error) {
    var output macie2.EnableOrganizationAdminAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetBucketStatisticsResult struct {
	Result workflow.Future
}

func (r *Macie2GetBucketStatisticsResult) Get(ctx workflow.Context) (*macie2.GetBucketStatisticsOutput, error) {
    var output macie2.GetBucketStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetClassificationExportConfigurationResult struct {
	Result workflow.Future
}

func (r *Macie2GetClassificationExportConfigurationResult) Get(ctx workflow.Context) (*macie2.GetClassificationExportConfigurationOutput, error) {
    var output macie2.GetClassificationExportConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetCustomDataIdentifierResult struct {
	Result workflow.Future
}

func (r *Macie2GetCustomDataIdentifierResult) Get(ctx workflow.Context) (*macie2.GetCustomDataIdentifierOutput, error) {
    var output macie2.GetCustomDataIdentifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetFindingStatisticsResult struct {
	Result workflow.Future
}

func (r *Macie2GetFindingStatisticsResult) Get(ctx workflow.Context) (*macie2.GetFindingStatisticsOutput, error) {
    var output macie2.GetFindingStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetFindingsResult struct {
	Result workflow.Future
}

func (r *Macie2GetFindingsResult) Get(ctx workflow.Context) (*macie2.GetFindingsOutput, error) {
    var output macie2.GetFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetFindingsFilterResult struct {
	Result workflow.Future
}

func (r *Macie2GetFindingsFilterResult) Get(ctx workflow.Context) (*macie2.GetFindingsFilterOutput, error) {
    var output macie2.GetFindingsFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetInvitationsCountResult struct {
	Result workflow.Future
}

func (r *Macie2GetInvitationsCountResult) Get(ctx workflow.Context) (*macie2.GetInvitationsCountOutput, error) {
    var output macie2.GetInvitationsCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetMacieSessionResult struct {
	Result workflow.Future
}

func (r *Macie2GetMacieSessionResult) Get(ctx workflow.Context) (*macie2.GetMacieSessionOutput, error) {
    var output macie2.GetMacieSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetMasterAccountResult struct {
	Result workflow.Future
}

func (r *Macie2GetMasterAccountResult) Get(ctx workflow.Context) (*macie2.GetMasterAccountOutput, error) {
    var output macie2.GetMasterAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetMemberResult struct {
	Result workflow.Future
}

func (r *Macie2GetMemberResult) Get(ctx workflow.Context) (*macie2.GetMemberOutput, error) {
    var output macie2.GetMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetUsageStatisticsResult struct {
	Result workflow.Future
}

func (r *Macie2GetUsageStatisticsResult) Get(ctx workflow.Context) (*macie2.GetUsageStatisticsOutput, error) {
    var output macie2.GetUsageStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2GetUsageTotalsResult struct {
	Result workflow.Future
}

func (r *Macie2GetUsageTotalsResult) Get(ctx workflow.Context) (*macie2.GetUsageTotalsOutput, error) {
    var output macie2.GetUsageTotalsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListClassificationJobsResult struct {
	Result workflow.Future
}

func (r *Macie2ListClassificationJobsResult) Get(ctx workflow.Context) (*macie2.ListClassificationJobsOutput, error) {
    var output macie2.ListClassificationJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListCustomDataIdentifiersResult struct {
	Result workflow.Future
}

func (r *Macie2ListCustomDataIdentifiersResult) Get(ctx workflow.Context) (*macie2.ListCustomDataIdentifiersOutput, error) {
    var output macie2.ListCustomDataIdentifiersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListFindingsResult struct {
	Result workflow.Future
}

func (r *Macie2ListFindingsResult) Get(ctx workflow.Context) (*macie2.ListFindingsOutput, error) {
    var output macie2.ListFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListFindingsFiltersResult struct {
	Result workflow.Future
}

func (r *Macie2ListFindingsFiltersResult) Get(ctx workflow.Context) (*macie2.ListFindingsFiltersOutput, error) {
    var output macie2.ListFindingsFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListInvitationsResult struct {
	Result workflow.Future
}

func (r *Macie2ListInvitationsResult) Get(ctx workflow.Context) (*macie2.ListInvitationsOutput, error) {
    var output macie2.ListInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListMembersResult struct {
	Result workflow.Future
}

func (r *Macie2ListMembersResult) Get(ctx workflow.Context) (*macie2.ListMembersOutput, error) {
    var output macie2.ListMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListOrganizationAdminAccountsResult struct {
	Result workflow.Future
}

func (r *Macie2ListOrganizationAdminAccountsResult) Get(ctx workflow.Context) (*macie2.ListOrganizationAdminAccountsOutput, error) {
    var output macie2.ListOrganizationAdminAccountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2ListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Macie2ListTagsForResourceResult) Get(ctx workflow.Context) (*macie2.ListTagsForResourceOutput, error) {
    var output macie2.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2PutClassificationExportConfigurationResult struct {
	Result workflow.Future
}

func (r *Macie2PutClassificationExportConfigurationResult) Get(ctx workflow.Context) (*macie2.PutClassificationExportConfigurationOutput, error) {
    var output macie2.PutClassificationExportConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2TagResourceResult struct {
	Result workflow.Future
}

func (r *Macie2TagResourceResult) Get(ctx workflow.Context) (*macie2.TagResourceOutput, error) {
    var output macie2.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2TestCustomDataIdentifierResult struct {
	Result workflow.Future
}

func (r *Macie2TestCustomDataIdentifierResult) Get(ctx workflow.Context) (*macie2.TestCustomDataIdentifierOutput, error) {
    var output macie2.TestCustomDataIdentifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2UntagResourceResult struct {
	Result workflow.Future
}

func (r *Macie2UntagResourceResult) Get(ctx workflow.Context) (*macie2.UntagResourceOutput, error) {
    var output macie2.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2UpdateClassificationJobResult struct {
	Result workflow.Future
}

func (r *Macie2UpdateClassificationJobResult) Get(ctx workflow.Context) (*macie2.UpdateClassificationJobOutput, error) {
    var output macie2.UpdateClassificationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2UpdateFindingsFilterResult struct {
	Result workflow.Future
}

func (r *Macie2UpdateFindingsFilterResult) Get(ctx workflow.Context) (*macie2.UpdateFindingsFilterOutput, error) {
    var output macie2.UpdateFindingsFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2UpdateMacieSessionResult struct {
	Result workflow.Future
}

func (r *Macie2UpdateMacieSessionResult) Get(ctx workflow.Context) (*macie2.UpdateMacieSessionOutput, error) {
    var output macie2.UpdateMacieSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2UpdateMemberSessionResult struct {
	Result workflow.Future
}

func (r *Macie2UpdateMemberSessionResult) Get(ctx workflow.Context) (*macie2.UpdateMemberSessionOutput, error) {
    var output macie2.UpdateMemberSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2UpdateOrganizationConfigurationResult struct {
	Result workflow.Future
}

func (r *Macie2UpdateOrganizationConfigurationResult) Get(ctx workflow.Context) (*macie2.UpdateOrganizationConfigurationOutput, error) {
    var output macie2.UpdateOrganizationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Macie2Stub struct {
    activities awsactivities.Macie2Activities
}

func NewMacie2Stub() Macie2Client {
    return &Macie2Stub{}
}

func (a *Macie2Stub) AcceptInvitation(ctx workflow.Context, input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error) {
    var output macie2.AcceptInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) AcceptInvitationAsync(ctx workflow.Context, input *macie2.AcceptInvitationInput) *Macie2AcceptInvitationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptInvitation, input)
    return &Macie2AcceptInvitationResult{Result: future}
}

func (a *Macie2Stub) BatchGetCustomDataIdentifiers(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error) {
    var output macie2.BatchGetCustomDataIdentifiersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetCustomDataIdentifiers, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) BatchGetCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) *Macie2BatchGetCustomDataIdentifiersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetCustomDataIdentifiers, input)
    return &Macie2BatchGetCustomDataIdentifiersResult{Result: future}
}

func (a *Macie2Stub) CreateClassificationJob(ctx workflow.Context, input *macie2.CreateClassificationJobInput) (*macie2.CreateClassificationJobOutput, error) {
    var output macie2.CreateClassificationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateClassificationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) CreateClassificationJobAsync(ctx workflow.Context, input *macie2.CreateClassificationJobInput) *Macie2CreateClassificationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateClassificationJob, input)
    return &Macie2CreateClassificationJobResult{Result: future}
}

func (a *Macie2Stub) CreateCustomDataIdentifier(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) (*macie2.CreateCustomDataIdentifierOutput, error) {
    var output macie2.CreateCustomDataIdentifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCustomDataIdentifier, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) CreateCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) *Macie2CreateCustomDataIdentifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCustomDataIdentifier, input)
    return &Macie2CreateCustomDataIdentifierResult{Result: future}
}

func (a *Macie2Stub) CreateFindingsFilter(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) (*macie2.CreateFindingsFilterOutput, error) {
    var output macie2.CreateFindingsFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFindingsFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) CreateFindingsFilterAsync(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) *Macie2CreateFindingsFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFindingsFilter, input)
    return &Macie2CreateFindingsFilterResult{Result: future}
}

func (a *Macie2Stub) CreateInvitations(ctx workflow.Context, input *macie2.CreateInvitationsInput) (*macie2.CreateInvitationsOutput, error) {
    var output macie2.CreateInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) CreateInvitationsAsync(ctx workflow.Context, input *macie2.CreateInvitationsInput) *Macie2CreateInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateInvitations, input)
    return &Macie2CreateInvitationsResult{Result: future}
}

func (a *Macie2Stub) CreateMember(ctx workflow.Context, input *macie2.CreateMemberInput) (*macie2.CreateMemberOutput, error) {
    var output macie2.CreateMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMember, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) CreateMemberAsync(ctx workflow.Context, input *macie2.CreateMemberInput) *Macie2CreateMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMember, input)
    return &Macie2CreateMemberResult{Result: future}
}

func (a *Macie2Stub) CreateSampleFindings(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) (*macie2.CreateSampleFindingsOutput, error) {
    var output macie2.CreateSampleFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSampleFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) CreateSampleFindingsAsync(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) *Macie2CreateSampleFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSampleFindings, input)
    return &Macie2CreateSampleFindingsResult{Result: future}
}

func (a *Macie2Stub) DeclineInvitations(ctx workflow.Context, input *macie2.DeclineInvitationsInput) (*macie2.DeclineInvitationsOutput, error) {
    var output macie2.DeclineInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeclineInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DeclineInvitationsAsync(ctx workflow.Context, input *macie2.DeclineInvitationsInput) *Macie2DeclineInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeclineInvitations, input)
    return &Macie2DeclineInvitationsResult{Result: future}
}

func (a *Macie2Stub) DeleteCustomDataIdentifier(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) (*macie2.DeleteCustomDataIdentifierOutput, error) {
    var output macie2.DeleteCustomDataIdentifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomDataIdentifier, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DeleteCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) *Macie2DeleteCustomDataIdentifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomDataIdentifier, input)
    return &Macie2DeleteCustomDataIdentifierResult{Result: future}
}

func (a *Macie2Stub) DeleteFindingsFilter(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) (*macie2.DeleteFindingsFilterOutput, error) {
    var output macie2.DeleteFindingsFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFindingsFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DeleteFindingsFilterAsync(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) *Macie2DeleteFindingsFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFindingsFilter, input)
    return &Macie2DeleteFindingsFilterResult{Result: future}
}

func (a *Macie2Stub) DeleteInvitations(ctx workflow.Context, input *macie2.DeleteInvitationsInput) (*macie2.DeleteInvitationsOutput, error) {
    var output macie2.DeleteInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DeleteInvitationsAsync(ctx workflow.Context, input *macie2.DeleteInvitationsInput) *Macie2DeleteInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteInvitations, input)
    return &Macie2DeleteInvitationsResult{Result: future}
}

func (a *Macie2Stub) DeleteMember(ctx workflow.Context, input *macie2.DeleteMemberInput) (*macie2.DeleteMemberOutput, error) {
    var output macie2.DeleteMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMember, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DeleteMemberAsync(ctx workflow.Context, input *macie2.DeleteMemberInput) *Macie2DeleteMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMember, input)
    return &Macie2DeleteMemberResult{Result: future}
}

func (a *Macie2Stub) DescribeBuckets(ctx workflow.Context, input *macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error) {
    var output macie2.DescribeBucketsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBuckets, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DescribeBucketsAsync(ctx workflow.Context, input *macie2.DescribeBucketsInput) *Macie2DescribeBucketsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBuckets, input)
    return &Macie2DescribeBucketsResult{Result: future}
}

func (a *Macie2Stub) DescribeClassificationJob(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error) {
    var output macie2.DescribeClassificationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClassificationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DescribeClassificationJobAsync(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) *Macie2DescribeClassificationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClassificationJob, input)
    return &Macie2DescribeClassificationJobResult{Result: future}
}

func (a *Macie2Stub) DescribeOrganizationConfiguration(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error) {
    var output macie2.DescribeOrganizationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DescribeOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) *Macie2DescribeOrganizationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationConfiguration, input)
    return &Macie2DescribeOrganizationConfigurationResult{Result: future}
}

func (a *Macie2Stub) DisableMacie(ctx workflow.Context, input *macie2.DisableMacieInput) (*macie2.DisableMacieOutput, error) {
    var output macie2.DisableMacieOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableMacie, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DisableMacieAsync(ctx workflow.Context, input *macie2.DisableMacieInput) *Macie2DisableMacieResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableMacie, input)
    return &Macie2DisableMacieResult{Result: future}
}

func (a *Macie2Stub) DisableOrganizationAdminAccount(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) (*macie2.DisableOrganizationAdminAccountOutput, error) {
    var output macie2.DisableOrganizationAdminAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableOrganizationAdminAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DisableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) *Macie2DisableOrganizationAdminAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableOrganizationAdminAccount, input)
    return &Macie2DisableOrganizationAdminAccountResult{Result: future}
}

func (a *Macie2Stub) DisassociateFromMasterAccount(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) (*macie2.DisassociateFromMasterAccountOutput, error) {
    var output macie2.DisassociateFromMasterAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateFromMasterAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DisassociateFromMasterAccountAsync(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) *Macie2DisassociateFromMasterAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateFromMasterAccount, input)
    return &Macie2DisassociateFromMasterAccountResult{Result: future}
}

func (a *Macie2Stub) DisassociateMember(ctx workflow.Context, input *macie2.DisassociateMemberInput) (*macie2.DisassociateMemberOutput, error) {
    var output macie2.DisassociateMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateMember, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) DisassociateMemberAsync(ctx workflow.Context, input *macie2.DisassociateMemberInput) *Macie2DisassociateMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateMember, input)
    return &Macie2DisassociateMemberResult{Result: future}
}

func (a *Macie2Stub) EnableMacie(ctx workflow.Context, input *macie2.EnableMacieInput) (*macie2.EnableMacieOutput, error) {
    var output macie2.EnableMacieOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableMacie, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) EnableMacieAsync(ctx workflow.Context, input *macie2.EnableMacieInput) *Macie2EnableMacieResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableMacie, input)
    return &Macie2EnableMacieResult{Result: future}
}

func (a *Macie2Stub) EnableOrganizationAdminAccount(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) (*macie2.EnableOrganizationAdminAccountOutput, error) {
    var output macie2.EnableOrganizationAdminAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableOrganizationAdminAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) EnableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) *Macie2EnableOrganizationAdminAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableOrganizationAdminAccount, input)
    return &Macie2EnableOrganizationAdminAccountResult{Result: future}
}

func (a *Macie2Stub) GetBucketStatistics(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error) {
    var output macie2.GetBucketStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBucketStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetBucketStatisticsAsync(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) *Macie2GetBucketStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBucketStatistics, input)
    return &Macie2GetBucketStatisticsResult{Result: future}
}

func (a *Macie2Stub) GetClassificationExportConfiguration(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error) {
    var output macie2.GetClassificationExportConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetClassificationExportConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) *Macie2GetClassificationExportConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetClassificationExportConfiguration, input)
    return &Macie2GetClassificationExportConfigurationResult{Result: future}
}

func (a *Macie2Stub) GetCustomDataIdentifier(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error) {
    var output macie2.GetCustomDataIdentifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCustomDataIdentifier, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) *Macie2GetCustomDataIdentifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCustomDataIdentifier, input)
    return &Macie2GetCustomDataIdentifierResult{Result: future}
}

func (a *Macie2Stub) GetFindingStatistics(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error) {
    var output macie2.GetFindingStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFindingStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetFindingStatisticsAsync(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) *Macie2GetFindingStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFindingStatistics, input)
    return &Macie2GetFindingStatisticsResult{Result: future}
}

func (a *Macie2Stub) GetFindings(ctx workflow.Context, input *macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error) {
    var output macie2.GetFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetFindingsAsync(ctx workflow.Context, input *macie2.GetFindingsInput) *Macie2GetFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFindings, input)
    return &Macie2GetFindingsResult{Result: future}
}

func (a *Macie2Stub) GetFindingsFilter(ctx workflow.Context, input *macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error) {
    var output macie2.GetFindingsFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFindingsFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetFindingsFilterAsync(ctx workflow.Context, input *macie2.GetFindingsFilterInput) *Macie2GetFindingsFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFindingsFilter, input)
    return &Macie2GetFindingsFilterResult{Result: future}
}

func (a *Macie2Stub) GetInvitationsCount(ctx workflow.Context, input *macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error) {
    var output macie2.GetInvitationsCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInvitationsCount, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetInvitationsCountAsync(ctx workflow.Context, input *macie2.GetInvitationsCountInput) *Macie2GetInvitationsCountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetInvitationsCount, input)
    return &Macie2GetInvitationsCountResult{Result: future}
}

func (a *Macie2Stub) GetMacieSession(ctx workflow.Context, input *macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error) {
    var output macie2.GetMacieSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMacieSession, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetMacieSessionAsync(ctx workflow.Context, input *macie2.GetMacieSessionInput) *Macie2GetMacieSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMacieSession, input)
    return &Macie2GetMacieSessionResult{Result: future}
}

func (a *Macie2Stub) GetMasterAccount(ctx workflow.Context, input *macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error) {
    var output macie2.GetMasterAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMasterAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetMasterAccountAsync(ctx workflow.Context, input *macie2.GetMasterAccountInput) *Macie2GetMasterAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMasterAccount, input)
    return &Macie2GetMasterAccountResult{Result: future}
}

func (a *Macie2Stub) GetMember(ctx workflow.Context, input *macie2.GetMemberInput) (*macie2.GetMemberOutput, error) {
    var output macie2.GetMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMember, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetMemberAsync(ctx workflow.Context, input *macie2.GetMemberInput) *Macie2GetMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMember, input)
    return &Macie2GetMemberResult{Result: future}
}

func (a *Macie2Stub) GetUsageStatistics(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error) {
    var output macie2.GetUsageStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUsageStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetUsageStatisticsAsync(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) *Macie2GetUsageStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUsageStatistics, input)
    return &Macie2GetUsageStatisticsResult{Result: future}
}

func (a *Macie2Stub) GetUsageTotals(ctx workflow.Context, input *macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error) {
    var output macie2.GetUsageTotalsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUsageTotals, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) GetUsageTotalsAsync(ctx workflow.Context, input *macie2.GetUsageTotalsInput) *Macie2GetUsageTotalsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUsageTotals, input)
    return &Macie2GetUsageTotalsResult{Result: future}
}

func (a *Macie2Stub) ListClassificationJobs(ctx workflow.Context, input *macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error) {
    var output macie2.ListClassificationJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListClassificationJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListClassificationJobsAsync(ctx workflow.Context, input *macie2.ListClassificationJobsInput) *Macie2ListClassificationJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListClassificationJobs, input)
    return &Macie2ListClassificationJobsResult{Result: future}
}

func (a *Macie2Stub) ListCustomDataIdentifiers(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error) {
    var output macie2.ListCustomDataIdentifiersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCustomDataIdentifiers, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) *Macie2ListCustomDataIdentifiersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCustomDataIdentifiers, input)
    return &Macie2ListCustomDataIdentifiersResult{Result: future}
}

func (a *Macie2Stub) ListFindings(ctx workflow.Context, input *macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error) {
    var output macie2.ListFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListFindingsAsync(ctx workflow.Context, input *macie2.ListFindingsInput) *Macie2ListFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFindings, input)
    return &Macie2ListFindingsResult{Result: future}
}

func (a *Macie2Stub) ListFindingsFilters(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error) {
    var output macie2.ListFindingsFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFindingsFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListFindingsFiltersAsync(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) *Macie2ListFindingsFiltersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFindingsFilters, input)
    return &Macie2ListFindingsFiltersResult{Result: future}
}

func (a *Macie2Stub) ListInvitations(ctx workflow.Context, input *macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error) {
    var output macie2.ListInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListInvitationsAsync(ctx workflow.Context, input *macie2.ListInvitationsInput) *Macie2ListInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input)
    return &Macie2ListInvitationsResult{Result: future}
}

func (a *Macie2Stub) ListMembers(ctx workflow.Context, input *macie2.ListMembersInput) (*macie2.ListMembersOutput, error) {
    var output macie2.ListMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListMembersAsync(ctx workflow.Context, input *macie2.ListMembersInput) *Macie2ListMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input)
    return &Macie2ListMembersResult{Result: future}
}

func (a *Macie2Stub) ListOrganizationAdminAccounts(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error) {
    var output macie2.ListOrganizationAdminAccountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationAdminAccounts, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListOrganizationAdminAccountsAsync(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) *Macie2ListOrganizationAdminAccountsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationAdminAccounts, input)
    return &Macie2ListOrganizationAdminAccountsResult{Result: future}
}

func (a *Macie2Stub) ListTagsForResource(ctx workflow.Context, input *macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error) {
    var output macie2.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) ListTagsForResourceAsync(ctx workflow.Context, input *macie2.ListTagsForResourceInput) *Macie2ListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Macie2ListTagsForResourceResult{Result: future}
}

func (a *Macie2Stub) PutClassificationExportConfiguration(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) (*macie2.PutClassificationExportConfigurationOutput, error) {
    var output macie2.PutClassificationExportConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutClassificationExportConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) PutClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) *Macie2PutClassificationExportConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutClassificationExportConfiguration, input)
    return &Macie2PutClassificationExportConfigurationResult{Result: future}
}

func (a *Macie2Stub) TagResource(ctx workflow.Context, input *macie2.TagResourceInput) (*macie2.TagResourceOutput, error) {
    var output macie2.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) TagResourceAsync(ctx workflow.Context, input *macie2.TagResourceInput) *Macie2TagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Macie2TagResourceResult{Result: future}
}

func (a *Macie2Stub) TestCustomDataIdentifier(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) (*macie2.TestCustomDataIdentifierOutput, error) {
    var output macie2.TestCustomDataIdentifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestCustomDataIdentifier, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) TestCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) *Macie2TestCustomDataIdentifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestCustomDataIdentifier, input)
    return &Macie2TestCustomDataIdentifierResult{Result: future}
}

func (a *Macie2Stub) UntagResource(ctx workflow.Context, input *macie2.UntagResourceInput) (*macie2.UntagResourceOutput, error) {
    var output macie2.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) UntagResourceAsync(ctx workflow.Context, input *macie2.UntagResourceInput) *Macie2UntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Macie2UntagResourceResult{Result: future}
}

func (a *Macie2Stub) UpdateClassificationJob(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) (*macie2.UpdateClassificationJobOutput, error) {
    var output macie2.UpdateClassificationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateClassificationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) UpdateClassificationJobAsync(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) *Macie2UpdateClassificationJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateClassificationJob, input)
    return &Macie2UpdateClassificationJobResult{Result: future}
}

func (a *Macie2Stub) UpdateFindingsFilter(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) (*macie2.UpdateFindingsFilterOutput, error) {
    var output macie2.UpdateFindingsFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFindingsFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) UpdateFindingsFilterAsync(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) *Macie2UpdateFindingsFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFindingsFilter, input)
    return &Macie2UpdateFindingsFilterResult{Result: future}
}

func (a *Macie2Stub) UpdateMacieSession(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) (*macie2.UpdateMacieSessionOutput, error) {
    var output macie2.UpdateMacieSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMacieSession, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) UpdateMacieSessionAsync(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) *Macie2UpdateMacieSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMacieSession, input)
    return &Macie2UpdateMacieSessionResult{Result: future}
}

func (a *Macie2Stub) UpdateMemberSession(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) (*macie2.UpdateMemberSessionOutput, error) {
    var output macie2.UpdateMemberSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMemberSession, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) UpdateMemberSessionAsync(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) *Macie2UpdateMemberSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMemberSession, input)
    return &Macie2UpdateMemberSessionResult{Result: future}
}

func (a *Macie2Stub) UpdateOrganizationConfiguration(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) (*macie2.UpdateOrganizationConfigurationOutput, error) {
    var output macie2.UpdateOrganizationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateOrganizationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *Macie2Stub) UpdateOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) *Macie2UpdateOrganizationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateOrganizationConfiguration, input)
    return &Macie2UpdateOrganizationConfigurationResult{Result: future}
}
