// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package macie2stub

import (
	"github.com/aws/aws-sdk-go/service/macie2"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptInvitation(ctx workflow.Context, input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error)
	AcceptInvitationAsync(ctx workflow.Context, input *macie2.AcceptInvitationInput) *AcceptInvitationFuture

	BatchGetCustomDataIdentifiers(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error)
	BatchGetCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.BatchGetCustomDataIdentifiersInput) *BatchGetCustomDataIdentifiersFuture

	CreateClassificationJob(ctx workflow.Context, input *macie2.CreateClassificationJobInput) (*macie2.CreateClassificationJobOutput, error)
	CreateClassificationJobAsync(ctx workflow.Context, input *macie2.CreateClassificationJobInput) *CreateClassificationJobFuture

	CreateCustomDataIdentifier(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) (*macie2.CreateCustomDataIdentifierOutput, error)
	CreateCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.CreateCustomDataIdentifierInput) *CreateCustomDataIdentifierFuture

	CreateFindingsFilter(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) (*macie2.CreateFindingsFilterOutput, error)
	CreateFindingsFilterAsync(ctx workflow.Context, input *macie2.CreateFindingsFilterInput) *CreateFindingsFilterFuture

	CreateInvitations(ctx workflow.Context, input *macie2.CreateInvitationsInput) (*macie2.CreateInvitationsOutput, error)
	CreateInvitationsAsync(ctx workflow.Context, input *macie2.CreateInvitationsInput) *CreateInvitationsFuture

	CreateMember(ctx workflow.Context, input *macie2.CreateMemberInput) (*macie2.CreateMemberOutput, error)
	CreateMemberAsync(ctx workflow.Context, input *macie2.CreateMemberInput) *CreateMemberFuture

	CreateSampleFindings(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) (*macie2.CreateSampleFindingsOutput, error)
	CreateSampleFindingsAsync(ctx workflow.Context, input *macie2.CreateSampleFindingsInput) *CreateSampleFindingsFuture

	DeclineInvitations(ctx workflow.Context, input *macie2.DeclineInvitationsInput) (*macie2.DeclineInvitationsOutput, error)
	DeclineInvitationsAsync(ctx workflow.Context, input *macie2.DeclineInvitationsInput) *DeclineInvitationsFuture

	DeleteCustomDataIdentifier(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) (*macie2.DeleteCustomDataIdentifierOutput, error)
	DeleteCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.DeleteCustomDataIdentifierInput) *DeleteCustomDataIdentifierFuture

	DeleteFindingsFilter(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) (*macie2.DeleteFindingsFilterOutput, error)
	DeleteFindingsFilterAsync(ctx workflow.Context, input *macie2.DeleteFindingsFilterInput) *DeleteFindingsFilterFuture

	DeleteInvitations(ctx workflow.Context, input *macie2.DeleteInvitationsInput) (*macie2.DeleteInvitationsOutput, error)
	DeleteInvitationsAsync(ctx workflow.Context, input *macie2.DeleteInvitationsInput) *DeleteInvitationsFuture

	DeleteMember(ctx workflow.Context, input *macie2.DeleteMemberInput) (*macie2.DeleteMemberOutput, error)
	DeleteMemberAsync(ctx workflow.Context, input *macie2.DeleteMemberInput) *DeleteMemberFuture

	DescribeBuckets(ctx workflow.Context, input *macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error)
	DescribeBucketsAsync(ctx workflow.Context, input *macie2.DescribeBucketsInput) *DescribeBucketsFuture

	DescribeClassificationJob(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error)
	DescribeClassificationJobAsync(ctx workflow.Context, input *macie2.DescribeClassificationJobInput) *DescribeClassificationJobFuture

	DescribeOrganizationConfiguration(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.DescribeOrganizationConfigurationInput) *DescribeOrganizationConfigurationFuture

	DisableMacie(ctx workflow.Context, input *macie2.DisableMacieInput) (*macie2.DisableMacieOutput, error)
	DisableMacieAsync(ctx workflow.Context, input *macie2.DisableMacieInput) *DisableMacieFuture

	DisableOrganizationAdminAccount(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) (*macie2.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.DisableOrganizationAdminAccountInput) *DisableOrganizationAdminAccountFuture

	DisassociateFromMasterAccount(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) (*macie2.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountAsync(ctx workflow.Context, input *macie2.DisassociateFromMasterAccountInput) *DisassociateFromMasterAccountFuture

	DisassociateMember(ctx workflow.Context, input *macie2.DisassociateMemberInput) (*macie2.DisassociateMemberOutput, error)
	DisassociateMemberAsync(ctx workflow.Context, input *macie2.DisassociateMemberInput) *DisassociateMemberFuture

	EnableMacie(ctx workflow.Context, input *macie2.EnableMacieInput) (*macie2.EnableMacieOutput, error)
	EnableMacieAsync(ctx workflow.Context, input *macie2.EnableMacieInput) *EnableMacieFuture

	EnableOrganizationAdminAccount(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) (*macie2.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountAsync(ctx workflow.Context, input *macie2.EnableOrganizationAdminAccountInput) *EnableOrganizationAdminAccountFuture

	GetBucketStatistics(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error)
	GetBucketStatisticsAsync(ctx workflow.Context, input *macie2.GetBucketStatisticsInput) *GetBucketStatisticsFuture

	GetClassificationExportConfiguration(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error)
	GetClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.GetClassificationExportConfigurationInput) *GetClassificationExportConfigurationFuture

	GetCustomDataIdentifier(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error)
	GetCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.GetCustomDataIdentifierInput) *GetCustomDataIdentifierFuture

	GetFindingStatistics(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error)
	GetFindingStatisticsAsync(ctx workflow.Context, input *macie2.GetFindingStatisticsInput) *GetFindingStatisticsFuture

	GetFindings(ctx workflow.Context, input *macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error)
	GetFindingsAsync(ctx workflow.Context, input *macie2.GetFindingsInput) *GetFindingsFuture

	GetFindingsFilter(ctx workflow.Context, input *macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error)
	GetFindingsFilterAsync(ctx workflow.Context, input *macie2.GetFindingsFilterInput) *GetFindingsFilterFuture

	GetInvitationsCount(ctx workflow.Context, input *macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error)
	GetInvitationsCountAsync(ctx workflow.Context, input *macie2.GetInvitationsCountInput) *GetInvitationsCountFuture

	GetMacieSession(ctx workflow.Context, input *macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error)
	GetMacieSessionAsync(ctx workflow.Context, input *macie2.GetMacieSessionInput) *GetMacieSessionFuture

	GetMasterAccount(ctx workflow.Context, input *macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error)
	GetMasterAccountAsync(ctx workflow.Context, input *macie2.GetMasterAccountInput) *GetMasterAccountFuture

	GetMember(ctx workflow.Context, input *macie2.GetMemberInput) (*macie2.GetMemberOutput, error)
	GetMemberAsync(ctx workflow.Context, input *macie2.GetMemberInput) *GetMemberFuture

	GetUsageStatistics(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error)
	GetUsageStatisticsAsync(ctx workflow.Context, input *macie2.GetUsageStatisticsInput) *GetUsageStatisticsFuture

	GetUsageTotals(ctx workflow.Context, input *macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error)
	GetUsageTotalsAsync(ctx workflow.Context, input *macie2.GetUsageTotalsInput) *GetUsageTotalsFuture

	ListClassificationJobs(ctx workflow.Context, input *macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error)
	ListClassificationJobsAsync(ctx workflow.Context, input *macie2.ListClassificationJobsInput) *ListClassificationJobsFuture

	ListCustomDataIdentifiers(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error)
	ListCustomDataIdentifiersAsync(ctx workflow.Context, input *macie2.ListCustomDataIdentifiersInput) *ListCustomDataIdentifiersFuture

	ListFindings(ctx workflow.Context, input *macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error)
	ListFindingsAsync(ctx workflow.Context, input *macie2.ListFindingsInput) *ListFindingsFuture

	ListFindingsFilters(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error)
	ListFindingsFiltersAsync(ctx workflow.Context, input *macie2.ListFindingsFiltersInput) *ListFindingsFiltersFuture

	ListInvitations(ctx workflow.Context, input *macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error)
	ListInvitationsAsync(ctx workflow.Context, input *macie2.ListInvitationsInput) *ListInvitationsFuture

	ListMembers(ctx workflow.Context, input *macie2.ListMembersInput) (*macie2.ListMembersOutput, error)
	ListMembersAsync(ctx workflow.Context, input *macie2.ListMembersInput) *ListMembersFuture

	ListOrganizationAdminAccounts(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsAsync(ctx workflow.Context, input *macie2.ListOrganizationAdminAccountsInput) *ListOrganizationAdminAccountsFuture

	ListTagsForResource(ctx workflow.Context, input *macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *macie2.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutClassificationExportConfiguration(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) (*macie2.PutClassificationExportConfigurationOutput, error)
	PutClassificationExportConfigurationAsync(ctx workflow.Context, input *macie2.PutClassificationExportConfigurationInput) *PutClassificationExportConfigurationFuture

	TagResource(ctx workflow.Context, input *macie2.TagResourceInput) (*macie2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *macie2.TagResourceInput) *TagResourceFuture

	TestCustomDataIdentifier(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) (*macie2.TestCustomDataIdentifierOutput, error)
	TestCustomDataIdentifierAsync(ctx workflow.Context, input *macie2.TestCustomDataIdentifierInput) *TestCustomDataIdentifierFuture

	UntagResource(ctx workflow.Context, input *macie2.UntagResourceInput) (*macie2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *macie2.UntagResourceInput) *UntagResourceFuture

	UpdateClassificationJob(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) (*macie2.UpdateClassificationJobOutput, error)
	UpdateClassificationJobAsync(ctx workflow.Context, input *macie2.UpdateClassificationJobInput) *UpdateClassificationJobFuture

	UpdateFindingsFilter(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) (*macie2.UpdateFindingsFilterOutput, error)
	UpdateFindingsFilterAsync(ctx workflow.Context, input *macie2.UpdateFindingsFilterInput) *UpdateFindingsFilterFuture

	UpdateMacieSession(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) (*macie2.UpdateMacieSessionOutput, error)
	UpdateMacieSessionAsync(ctx workflow.Context, input *macie2.UpdateMacieSessionInput) *UpdateMacieSessionFuture

	UpdateMemberSession(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) (*macie2.UpdateMemberSessionOutput, error)
	UpdateMemberSessionAsync(ctx workflow.Context, input *macie2.UpdateMemberSessionInput) *UpdateMemberSessionFuture

	UpdateOrganizationConfiguration(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) (*macie2.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationAsync(ctx workflow.Context, input *macie2.UpdateOrganizationConfigurationInput) *UpdateOrganizationConfigurationFuture
}

func NewClient() Client {
	return &stub{}
}
