package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/aws/aws-sdk-go/service/securityhub/securityhubiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SecurityHubActivities struct {
	client securityhubiface.SecurityHubAPI
}

func NewSecurityHubActivities(session *session.Session, config ...*aws.Config) *SecurityHubActivities {
	client := securityhub.New(session, config...)
	return &SecurityHubActivities{client: client}
}

func (a *SecurityHubActivities) AcceptInvitation(ctx context.Context, input *securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error) {
	return a.client.AcceptInvitationWithContext(ctx, input)
}

func (a *SecurityHubActivities) BatchDisableStandards(ctx context.Context, input *securityhub.BatchDisableStandardsInput) (*securityhub.BatchDisableStandardsOutput, error) {
	return a.client.BatchDisableStandardsWithContext(ctx, input)
}

func (a *SecurityHubActivities) BatchEnableStandards(ctx context.Context, input *securityhub.BatchEnableStandardsInput) (*securityhub.BatchEnableStandardsOutput, error) {
	return a.client.BatchEnableStandardsWithContext(ctx, input)
}

func (a *SecurityHubActivities) BatchImportFindings(ctx context.Context, input *securityhub.BatchImportFindingsInput) (*securityhub.BatchImportFindingsOutput, error) {
	return a.client.BatchImportFindingsWithContext(ctx, input)
}

func (a *SecurityHubActivities) BatchUpdateFindings(ctx context.Context, input *securityhub.BatchUpdateFindingsInput) (*securityhub.BatchUpdateFindingsOutput, error) {
	return a.client.BatchUpdateFindingsWithContext(ctx, input)
}

func (a *SecurityHubActivities) CreateActionTarget(ctx context.Context, input *securityhub.CreateActionTargetInput) (*securityhub.CreateActionTargetOutput, error) {
	return a.client.CreateActionTargetWithContext(ctx, input)
}

func (a *SecurityHubActivities) CreateInsight(ctx context.Context, input *securityhub.CreateInsightInput) (*securityhub.CreateInsightOutput, error) {
	return a.client.CreateInsightWithContext(ctx, input)
}

func (a *SecurityHubActivities) CreateMembers(ctx context.Context, input *securityhub.CreateMembersInput) (*securityhub.CreateMembersOutput, error) {
	return a.client.CreateMembersWithContext(ctx, input)
}

func (a *SecurityHubActivities) DeclineInvitations(ctx context.Context, input *securityhub.DeclineInvitationsInput) (*securityhub.DeclineInvitationsOutput, error) {
	return a.client.DeclineInvitationsWithContext(ctx, input)
}

func (a *SecurityHubActivities) DeleteActionTarget(ctx context.Context, input *securityhub.DeleteActionTargetInput) (*securityhub.DeleteActionTargetOutput, error) {
	return a.client.DeleteActionTargetWithContext(ctx, input)
}

func (a *SecurityHubActivities) DeleteInsight(ctx context.Context, input *securityhub.DeleteInsightInput) (*securityhub.DeleteInsightOutput, error) {
	return a.client.DeleteInsightWithContext(ctx, input)
}

func (a *SecurityHubActivities) DeleteInvitations(ctx context.Context, input *securityhub.DeleteInvitationsInput) (*securityhub.DeleteInvitationsOutput, error) {
	return a.client.DeleteInvitationsWithContext(ctx, input)
}

func (a *SecurityHubActivities) DeleteMembers(ctx context.Context, input *securityhub.DeleteMembersInput) (*securityhub.DeleteMembersOutput, error) {
	return a.client.DeleteMembersWithContext(ctx, input)
}

func (a *SecurityHubActivities) DescribeActionTargets(ctx context.Context, input *securityhub.DescribeActionTargetsInput) (*securityhub.DescribeActionTargetsOutput, error) {
	return a.client.DescribeActionTargetsWithContext(ctx, input)
}

func (a *SecurityHubActivities) DescribeHub(ctx context.Context, input *securityhub.DescribeHubInput) (*securityhub.DescribeHubOutput, error) {
	return a.client.DescribeHubWithContext(ctx, input)
}

func (a *SecurityHubActivities) DescribeProducts(ctx context.Context, input *securityhub.DescribeProductsInput) (*securityhub.DescribeProductsOutput, error) {
	return a.client.DescribeProductsWithContext(ctx, input)
}

func (a *SecurityHubActivities) DescribeStandards(ctx context.Context, input *securityhub.DescribeStandardsInput) (*securityhub.DescribeStandardsOutput, error) {
	return a.client.DescribeStandardsWithContext(ctx, input)
}

func (a *SecurityHubActivities) DescribeStandardsControls(ctx context.Context, input *securityhub.DescribeStandardsControlsInput) (*securityhub.DescribeStandardsControlsOutput, error) {
	return a.client.DescribeStandardsControlsWithContext(ctx, input)
}

func (a *SecurityHubActivities) DisableImportFindingsForProduct(ctx context.Context, input *securityhub.DisableImportFindingsForProductInput) (*securityhub.DisableImportFindingsForProductOutput, error) {
	return a.client.DisableImportFindingsForProductWithContext(ctx, input)
}

func (a *SecurityHubActivities) DisableSecurityHub(ctx context.Context, input *securityhub.DisableSecurityHubInput) (*securityhub.DisableSecurityHubOutput, error) {
	return a.client.DisableSecurityHubWithContext(ctx, input)
}

func (a *SecurityHubActivities) DisassociateFromMasterAccount(ctx context.Context, input *securityhub.DisassociateFromMasterAccountInput) (*securityhub.DisassociateFromMasterAccountOutput, error) {
	return a.client.DisassociateFromMasterAccountWithContext(ctx, input)
}

func (a *SecurityHubActivities) DisassociateMembers(ctx context.Context, input *securityhub.DisassociateMembersInput) (*securityhub.DisassociateMembersOutput, error) {
	return a.client.DisassociateMembersWithContext(ctx, input)
}

func (a *SecurityHubActivities) EnableImportFindingsForProduct(ctx context.Context, input *securityhub.EnableImportFindingsForProductInput) (*securityhub.EnableImportFindingsForProductOutput, error) {
	return a.client.EnableImportFindingsForProductWithContext(ctx, input)
}

func (a *SecurityHubActivities) EnableSecurityHub(ctx context.Context, input *securityhub.EnableSecurityHubInput) (*securityhub.EnableSecurityHubOutput, error) {
	return a.client.EnableSecurityHubWithContext(ctx, input)
}

func (a *SecurityHubActivities) GetEnabledStandards(ctx context.Context, input *securityhub.GetEnabledStandardsInput) (*securityhub.GetEnabledStandardsOutput, error) {
	return a.client.GetEnabledStandardsWithContext(ctx, input)
}

func (a *SecurityHubActivities) GetFindings(ctx context.Context, input *securityhub.GetFindingsInput) (*securityhub.GetFindingsOutput, error) {
	return a.client.GetFindingsWithContext(ctx, input)
}

func (a *SecurityHubActivities) GetInsightResults(ctx context.Context, input *securityhub.GetInsightResultsInput) (*securityhub.GetInsightResultsOutput, error) {
	return a.client.GetInsightResultsWithContext(ctx, input)
}

func (a *SecurityHubActivities) GetInsights(ctx context.Context, input *securityhub.GetInsightsInput) (*securityhub.GetInsightsOutput, error) {
	return a.client.GetInsightsWithContext(ctx, input)
}

func (a *SecurityHubActivities) GetInvitationsCount(ctx context.Context, input *securityhub.GetInvitationsCountInput) (*securityhub.GetInvitationsCountOutput, error) {
	return a.client.GetInvitationsCountWithContext(ctx, input)
}

func (a *SecurityHubActivities) GetMasterAccount(ctx context.Context, input *securityhub.GetMasterAccountInput) (*securityhub.GetMasterAccountOutput, error) {
	return a.client.GetMasterAccountWithContext(ctx, input)
}

func (a *SecurityHubActivities) GetMembers(ctx context.Context, input *securityhub.GetMembersInput) (*securityhub.GetMembersOutput, error) {
	return a.client.GetMembersWithContext(ctx, input)
}

func (a *SecurityHubActivities) InviteMembers(ctx context.Context, input *securityhub.InviteMembersInput) (*securityhub.InviteMembersOutput, error) {
	return a.client.InviteMembersWithContext(ctx, input)
}

func (a *SecurityHubActivities) ListEnabledProductsForImport(ctx context.Context, input *securityhub.ListEnabledProductsForImportInput) (*securityhub.ListEnabledProductsForImportOutput, error) {
	return a.client.ListEnabledProductsForImportWithContext(ctx, input)
}

func (a *SecurityHubActivities) ListInvitations(ctx context.Context, input *securityhub.ListInvitationsInput) (*securityhub.ListInvitationsOutput, error) {
	return a.client.ListInvitationsWithContext(ctx, input)
}

func (a *SecurityHubActivities) ListMembers(ctx context.Context, input *securityhub.ListMembersInput) (*securityhub.ListMembersOutput, error) {
	return a.client.ListMembersWithContext(ctx, input)
}

func (a *SecurityHubActivities) ListTagsForResource(ctx context.Context, input *securityhub.ListTagsForResourceInput) (*securityhub.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SecurityHubActivities) TagResource(ctx context.Context, input *securityhub.TagResourceInput) (*securityhub.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SecurityHubActivities) UntagResource(ctx context.Context, input *securityhub.UntagResourceInput) (*securityhub.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SecurityHubActivities) UpdateActionTarget(ctx context.Context, input *securityhub.UpdateActionTargetInput) (*securityhub.UpdateActionTargetOutput, error) {
	return a.client.UpdateActionTargetWithContext(ctx, input)
}

func (a *SecurityHubActivities) UpdateFindings(ctx context.Context, input *securityhub.UpdateFindingsInput) (*securityhub.UpdateFindingsOutput, error) {
	return a.client.UpdateFindingsWithContext(ctx, input)
}

func (a *SecurityHubActivities) UpdateInsight(ctx context.Context, input *securityhub.UpdateInsightInput) (*securityhub.UpdateInsightOutput, error) {
	return a.client.UpdateInsightWithContext(ctx, input)
}

func (a *SecurityHubActivities) UpdateSecurityHubConfiguration(ctx context.Context, input *securityhub.UpdateSecurityHubConfigurationInput) (*securityhub.UpdateSecurityHubConfigurationOutput, error) {
	return a.client.UpdateSecurityHubConfigurationWithContext(ctx, input)
}

func (a *SecurityHubActivities) UpdateStandardsControl(ctx context.Context, input *securityhub.UpdateStandardsControlInput) (*securityhub.UpdateStandardsControlOutput, error) {
	return a.client.UpdateStandardsControlWithContext(ctx, input)
}
