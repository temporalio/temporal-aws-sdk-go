
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/aws/aws-sdk-go/service/securityhub/securityhubiface"
)

type SecurityHubActivities struct {
	client securityhubiface.SecurityHubAPI
}

func NewSecurityHubActivities(session *session.Session, config... *aws.Config) *SecurityHubActivities {
    client := securityhub.New(session, config...)
    return &SecurityHubActivities{client: client}
}

func (a *SecurityHubActivities) AcceptInvitation(input *securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error) {
    return a.client.AcceptInvitation(input)
}

func (a *SecurityHubActivities) BatchDisableStandards(input *securityhub.BatchDisableStandardsInput) (*securityhub.BatchDisableStandardsOutput, error) {
    return a.client.BatchDisableStandards(input)
}

func (a *SecurityHubActivities) BatchEnableStandards(input *securityhub.BatchEnableStandardsInput) (*securityhub.BatchEnableStandardsOutput, error) {
    return a.client.BatchEnableStandards(input)
}

func (a *SecurityHubActivities) BatchImportFindings(input *securityhub.BatchImportFindingsInput) (*securityhub.BatchImportFindingsOutput, error) {
    return a.client.BatchImportFindings(input)
}

func (a *SecurityHubActivities) BatchUpdateFindings(input *securityhub.BatchUpdateFindingsInput) (*securityhub.BatchUpdateFindingsOutput, error) {
    return a.client.BatchUpdateFindings(input)
}

func (a *SecurityHubActivities) CreateActionTarget(input *securityhub.CreateActionTargetInput) (*securityhub.CreateActionTargetOutput, error) {
    return a.client.CreateActionTarget(input)
}

func (a *SecurityHubActivities) CreateInsight(input *securityhub.CreateInsightInput) (*securityhub.CreateInsightOutput, error) {
    return a.client.CreateInsight(input)
}

func (a *SecurityHubActivities) CreateMembers(input *securityhub.CreateMembersInput) (*securityhub.CreateMembersOutput, error) {
    return a.client.CreateMembers(input)
}

func (a *SecurityHubActivities) DeclineInvitations(input *securityhub.DeclineInvitationsInput) (*securityhub.DeclineInvitationsOutput, error) {
    return a.client.DeclineInvitations(input)
}

func (a *SecurityHubActivities) DeleteActionTarget(input *securityhub.DeleteActionTargetInput) (*securityhub.DeleteActionTargetOutput, error) {
    return a.client.DeleteActionTarget(input)
}

func (a *SecurityHubActivities) DeleteInsight(input *securityhub.DeleteInsightInput) (*securityhub.DeleteInsightOutput, error) {
    return a.client.DeleteInsight(input)
}

func (a *SecurityHubActivities) DeleteInvitations(input *securityhub.DeleteInvitationsInput) (*securityhub.DeleteInvitationsOutput, error) {
    return a.client.DeleteInvitations(input)
}

func (a *SecurityHubActivities) DeleteMembers(input *securityhub.DeleteMembersInput) (*securityhub.DeleteMembersOutput, error) {
    return a.client.DeleteMembers(input)
}

func (a *SecurityHubActivities) DescribeActionTargets(input *securityhub.DescribeActionTargetsInput) (*securityhub.DescribeActionTargetsOutput, error) {
    return a.client.DescribeActionTargets(input)
}

func (a *SecurityHubActivities) DescribeHub(input *securityhub.DescribeHubInput) (*securityhub.DescribeHubOutput, error) {
    return a.client.DescribeHub(input)
}

func (a *SecurityHubActivities) DescribeProducts(input *securityhub.DescribeProductsInput) (*securityhub.DescribeProductsOutput, error) {
    return a.client.DescribeProducts(input)
}

func (a *SecurityHubActivities) DescribeStandards(input *securityhub.DescribeStandardsInput) (*securityhub.DescribeStandardsOutput, error) {
    return a.client.DescribeStandards(input)
}

func (a *SecurityHubActivities) DescribeStandardsControls(input *securityhub.DescribeStandardsControlsInput) (*securityhub.DescribeStandardsControlsOutput, error) {
    return a.client.DescribeStandardsControls(input)
}

func (a *SecurityHubActivities) DisableImportFindingsForProduct(input *securityhub.DisableImportFindingsForProductInput) (*securityhub.DisableImportFindingsForProductOutput, error) {
    return a.client.DisableImportFindingsForProduct(input)
}

func (a *SecurityHubActivities) DisableSecurityHub(input *securityhub.DisableSecurityHubInput) (*securityhub.DisableSecurityHubOutput, error) {
    return a.client.DisableSecurityHub(input)
}

func (a *SecurityHubActivities) DisassociateFromMasterAccount(input *securityhub.DisassociateFromMasterAccountInput) (*securityhub.DisassociateFromMasterAccountOutput, error) {
    return a.client.DisassociateFromMasterAccount(input)
}

func (a *SecurityHubActivities) DisassociateMembers(input *securityhub.DisassociateMembersInput) (*securityhub.DisassociateMembersOutput, error) {
    return a.client.DisassociateMembers(input)
}

func (a *SecurityHubActivities) EnableImportFindingsForProduct(input *securityhub.EnableImportFindingsForProductInput) (*securityhub.EnableImportFindingsForProductOutput, error) {
    return a.client.EnableImportFindingsForProduct(input)
}

func (a *SecurityHubActivities) EnableSecurityHub(input *securityhub.EnableSecurityHubInput) (*securityhub.EnableSecurityHubOutput, error) {
    return a.client.EnableSecurityHub(input)
}

func (a *SecurityHubActivities) GetEnabledStandards(input *securityhub.GetEnabledStandardsInput) (*securityhub.GetEnabledStandardsOutput, error) {
    return a.client.GetEnabledStandards(input)
}

func (a *SecurityHubActivities) GetFindings(input *securityhub.GetFindingsInput) (*securityhub.GetFindingsOutput, error) {
    return a.client.GetFindings(input)
}

func (a *SecurityHubActivities) GetInsightResults(input *securityhub.GetInsightResultsInput) (*securityhub.GetInsightResultsOutput, error) {
    return a.client.GetInsightResults(input)
}

func (a *SecurityHubActivities) GetInsights(input *securityhub.GetInsightsInput) (*securityhub.GetInsightsOutput, error) {
    return a.client.GetInsights(input)
}

func (a *SecurityHubActivities) GetInvitationsCount(input *securityhub.GetInvitationsCountInput) (*securityhub.GetInvitationsCountOutput, error) {
    return a.client.GetInvitationsCount(input)
}

func (a *SecurityHubActivities) GetMasterAccount(input *securityhub.GetMasterAccountInput) (*securityhub.GetMasterAccountOutput, error) {
    return a.client.GetMasterAccount(input)
}

func (a *SecurityHubActivities) GetMembers(input *securityhub.GetMembersInput) (*securityhub.GetMembersOutput, error) {
    return a.client.GetMembers(input)
}

func (a *SecurityHubActivities) InviteMembers(input *securityhub.InviteMembersInput) (*securityhub.InviteMembersOutput, error) {
    return a.client.InviteMembers(input)
}

func (a *SecurityHubActivities) ListEnabledProductsForImport(input *securityhub.ListEnabledProductsForImportInput) (*securityhub.ListEnabledProductsForImportOutput, error) {
    return a.client.ListEnabledProductsForImport(input)
}

func (a *SecurityHubActivities) ListInvitations(input *securityhub.ListInvitationsInput) (*securityhub.ListInvitationsOutput, error) {
    return a.client.ListInvitations(input)
}

func (a *SecurityHubActivities) ListMembers(input *securityhub.ListMembersInput) (*securityhub.ListMembersOutput, error) {
    return a.client.ListMembers(input)
}

func (a *SecurityHubActivities) ListTagsForResource(input *securityhub.ListTagsForResourceInput) (*securityhub.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SecurityHubActivities) TagResource(input *securityhub.TagResourceInput) (*securityhub.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SecurityHubActivities) UntagResource(input *securityhub.UntagResourceInput) (*securityhub.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *SecurityHubActivities) UpdateActionTarget(input *securityhub.UpdateActionTargetInput) (*securityhub.UpdateActionTargetOutput, error) {
    return a.client.UpdateActionTarget(input)
}

func (a *SecurityHubActivities) UpdateFindings(input *securityhub.UpdateFindingsInput) (*securityhub.UpdateFindingsOutput, error) {
    return a.client.UpdateFindings(input)
}

func (a *SecurityHubActivities) UpdateInsight(input *securityhub.UpdateInsightInput) (*securityhub.UpdateInsightOutput, error) {
    return a.client.UpdateInsight(input)
}

func (a *SecurityHubActivities) UpdateSecurityHubConfiguration(input *securityhub.UpdateSecurityHubConfigurationInput) (*securityhub.UpdateSecurityHubConfigurationOutput, error) {
    return a.client.UpdateSecurityHubConfiguration(input)
}

func (a *SecurityHubActivities) UpdateStandardsControl(input *securityhub.UpdateStandardsControlInput) (*securityhub.UpdateStandardsControlOutput, error) {
    return a.client.UpdateStandardsControl(input)
}
