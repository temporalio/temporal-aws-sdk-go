package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/securityhub"
	"go.temporal.io/sdk/workflow"
)

type SecurityHubClient interface {
    AcceptInvitation(ctx workflow.Context, input *securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error)
    AcceptInvitationAsync(ctx workflow.Context, input *securityhub.AcceptInvitationInput) *SecurityhubAcceptInvitationResult

    BatchDisableStandards(ctx workflow.Context, input *securityhub.BatchDisableStandardsInput) (*securityhub.BatchDisableStandardsOutput, error)
    BatchDisableStandardsAsync(ctx workflow.Context, input *securityhub.BatchDisableStandardsInput) *SecurityhubBatchDisableStandardsResult

    BatchEnableStandards(ctx workflow.Context, input *securityhub.BatchEnableStandardsInput) (*securityhub.BatchEnableStandardsOutput, error)
    BatchEnableStandardsAsync(ctx workflow.Context, input *securityhub.BatchEnableStandardsInput) *SecurityhubBatchEnableStandardsResult

    BatchImportFindings(ctx workflow.Context, input *securityhub.BatchImportFindingsInput) (*securityhub.BatchImportFindingsOutput, error)
    BatchImportFindingsAsync(ctx workflow.Context, input *securityhub.BatchImportFindingsInput) *SecurityhubBatchImportFindingsResult

    BatchUpdateFindings(ctx workflow.Context, input *securityhub.BatchUpdateFindingsInput) (*securityhub.BatchUpdateFindingsOutput, error)
    BatchUpdateFindingsAsync(ctx workflow.Context, input *securityhub.BatchUpdateFindingsInput) *SecurityhubBatchUpdateFindingsResult

    CreateActionTarget(ctx workflow.Context, input *securityhub.CreateActionTargetInput) (*securityhub.CreateActionTargetOutput, error)
    CreateActionTargetAsync(ctx workflow.Context, input *securityhub.CreateActionTargetInput) *SecurityhubCreateActionTargetResult

    CreateInsight(ctx workflow.Context, input *securityhub.CreateInsightInput) (*securityhub.CreateInsightOutput, error)
    CreateInsightAsync(ctx workflow.Context, input *securityhub.CreateInsightInput) *SecurityhubCreateInsightResult

    CreateMembers(ctx workflow.Context, input *securityhub.CreateMembersInput) (*securityhub.CreateMembersOutput, error)
    CreateMembersAsync(ctx workflow.Context, input *securityhub.CreateMembersInput) *SecurityhubCreateMembersResult

    DeclineInvitations(ctx workflow.Context, input *securityhub.DeclineInvitationsInput) (*securityhub.DeclineInvitationsOutput, error)
    DeclineInvitationsAsync(ctx workflow.Context, input *securityhub.DeclineInvitationsInput) *SecurityhubDeclineInvitationsResult

    DeleteActionTarget(ctx workflow.Context, input *securityhub.DeleteActionTargetInput) (*securityhub.DeleteActionTargetOutput, error)
    DeleteActionTargetAsync(ctx workflow.Context, input *securityhub.DeleteActionTargetInput) *SecurityhubDeleteActionTargetResult

    DeleteInsight(ctx workflow.Context, input *securityhub.DeleteInsightInput) (*securityhub.DeleteInsightOutput, error)
    DeleteInsightAsync(ctx workflow.Context, input *securityhub.DeleteInsightInput) *SecurityhubDeleteInsightResult

    DeleteInvitations(ctx workflow.Context, input *securityhub.DeleteInvitationsInput) (*securityhub.DeleteInvitationsOutput, error)
    DeleteInvitationsAsync(ctx workflow.Context, input *securityhub.DeleteInvitationsInput) *SecurityhubDeleteInvitationsResult

    DeleteMembers(ctx workflow.Context, input *securityhub.DeleteMembersInput) (*securityhub.DeleteMembersOutput, error)
    DeleteMembersAsync(ctx workflow.Context, input *securityhub.DeleteMembersInput) *SecurityhubDeleteMembersResult

    DescribeActionTargets(ctx workflow.Context, input *securityhub.DescribeActionTargetsInput) (*securityhub.DescribeActionTargetsOutput, error)
    DescribeActionTargetsAsync(ctx workflow.Context, input *securityhub.DescribeActionTargetsInput) *SecurityhubDescribeActionTargetsResult

    DescribeHub(ctx workflow.Context, input *securityhub.DescribeHubInput) (*securityhub.DescribeHubOutput, error)
    DescribeHubAsync(ctx workflow.Context, input *securityhub.DescribeHubInput) *SecurityhubDescribeHubResult

    DescribeProducts(ctx workflow.Context, input *securityhub.DescribeProductsInput) (*securityhub.DescribeProductsOutput, error)
    DescribeProductsAsync(ctx workflow.Context, input *securityhub.DescribeProductsInput) *SecurityhubDescribeProductsResult

    DescribeStandards(ctx workflow.Context, input *securityhub.DescribeStandardsInput) (*securityhub.DescribeStandardsOutput, error)
    DescribeStandardsAsync(ctx workflow.Context, input *securityhub.DescribeStandardsInput) *SecurityhubDescribeStandardsResult

    DescribeStandardsControls(ctx workflow.Context, input *securityhub.DescribeStandardsControlsInput) (*securityhub.DescribeStandardsControlsOutput, error)
    DescribeStandardsControlsAsync(ctx workflow.Context, input *securityhub.DescribeStandardsControlsInput) *SecurityhubDescribeStandardsControlsResult

    DisableImportFindingsForProduct(ctx workflow.Context, input *securityhub.DisableImportFindingsForProductInput) (*securityhub.DisableImportFindingsForProductOutput, error)
    DisableImportFindingsForProductAsync(ctx workflow.Context, input *securityhub.DisableImportFindingsForProductInput) *SecurityhubDisableImportFindingsForProductResult

    DisableSecurityHub(ctx workflow.Context, input *securityhub.DisableSecurityHubInput) (*securityhub.DisableSecurityHubOutput, error)
    DisableSecurityHubAsync(ctx workflow.Context, input *securityhub.DisableSecurityHubInput) *SecurityhubDisableSecurityHubResult

    DisassociateFromMasterAccount(ctx workflow.Context, input *securityhub.DisassociateFromMasterAccountInput) (*securityhub.DisassociateFromMasterAccountOutput, error)
    DisassociateFromMasterAccountAsync(ctx workflow.Context, input *securityhub.DisassociateFromMasterAccountInput) *SecurityhubDisassociateFromMasterAccountResult

    DisassociateMembers(ctx workflow.Context, input *securityhub.DisassociateMembersInput) (*securityhub.DisassociateMembersOutput, error)
    DisassociateMembersAsync(ctx workflow.Context, input *securityhub.DisassociateMembersInput) *SecurityhubDisassociateMembersResult

    EnableImportFindingsForProduct(ctx workflow.Context, input *securityhub.EnableImportFindingsForProductInput) (*securityhub.EnableImportFindingsForProductOutput, error)
    EnableImportFindingsForProductAsync(ctx workflow.Context, input *securityhub.EnableImportFindingsForProductInput) *SecurityhubEnableImportFindingsForProductResult

    EnableSecurityHub(ctx workflow.Context, input *securityhub.EnableSecurityHubInput) (*securityhub.EnableSecurityHubOutput, error)
    EnableSecurityHubAsync(ctx workflow.Context, input *securityhub.EnableSecurityHubInput) *SecurityhubEnableSecurityHubResult

    GetEnabledStandards(ctx workflow.Context, input *securityhub.GetEnabledStandardsInput) (*securityhub.GetEnabledStandardsOutput, error)
    GetEnabledStandardsAsync(ctx workflow.Context, input *securityhub.GetEnabledStandardsInput) *SecurityhubGetEnabledStandardsResult

    GetFindings(ctx workflow.Context, input *securityhub.GetFindingsInput) (*securityhub.GetFindingsOutput, error)
    GetFindingsAsync(ctx workflow.Context, input *securityhub.GetFindingsInput) *SecurityhubGetFindingsResult

    GetInsightResults(ctx workflow.Context, input *securityhub.GetInsightResultsInput) (*securityhub.GetInsightResultsOutput, error)
    GetInsightResultsAsync(ctx workflow.Context, input *securityhub.GetInsightResultsInput) *SecurityhubGetInsightResultsResult

    GetInsights(ctx workflow.Context, input *securityhub.GetInsightsInput) (*securityhub.GetInsightsOutput, error)
    GetInsightsAsync(ctx workflow.Context, input *securityhub.GetInsightsInput) *SecurityhubGetInsightsResult

    GetInvitationsCount(ctx workflow.Context, input *securityhub.GetInvitationsCountInput) (*securityhub.GetInvitationsCountOutput, error)
    GetInvitationsCountAsync(ctx workflow.Context, input *securityhub.GetInvitationsCountInput) *SecurityhubGetInvitationsCountResult

    GetMasterAccount(ctx workflow.Context, input *securityhub.GetMasterAccountInput) (*securityhub.GetMasterAccountOutput, error)
    GetMasterAccountAsync(ctx workflow.Context, input *securityhub.GetMasterAccountInput) *SecurityhubGetMasterAccountResult

    GetMembers(ctx workflow.Context, input *securityhub.GetMembersInput) (*securityhub.GetMembersOutput, error)
    GetMembersAsync(ctx workflow.Context, input *securityhub.GetMembersInput) *SecurityhubGetMembersResult

    InviteMembers(ctx workflow.Context, input *securityhub.InviteMembersInput) (*securityhub.InviteMembersOutput, error)
    InviteMembersAsync(ctx workflow.Context, input *securityhub.InviteMembersInput) *SecurityhubInviteMembersResult

    ListEnabledProductsForImport(ctx workflow.Context, input *securityhub.ListEnabledProductsForImportInput) (*securityhub.ListEnabledProductsForImportOutput, error)
    ListEnabledProductsForImportAsync(ctx workflow.Context, input *securityhub.ListEnabledProductsForImportInput) *SecurityhubListEnabledProductsForImportResult

    ListInvitations(ctx workflow.Context, input *securityhub.ListInvitationsInput) (*securityhub.ListInvitationsOutput, error)
    ListInvitationsAsync(ctx workflow.Context, input *securityhub.ListInvitationsInput) *SecurityhubListInvitationsResult

    ListMembers(ctx workflow.Context, input *securityhub.ListMembersInput) (*securityhub.ListMembersOutput, error)
    ListMembersAsync(ctx workflow.Context, input *securityhub.ListMembersInput) *SecurityhubListMembersResult

    ListTagsForResource(ctx workflow.Context, input *securityhub.ListTagsForResourceInput) (*securityhub.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *securityhub.ListTagsForResourceInput) *SecurityhubListTagsForResourceResult

    TagResource(ctx workflow.Context, input *securityhub.TagResourceInput) (*securityhub.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *securityhub.TagResourceInput) *SecurityhubTagResourceResult

    UntagResource(ctx workflow.Context, input *securityhub.UntagResourceInput) (*securityhub.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *securityhub.UntagResourceInput) *SecurityhubUntagResourceResult

    UpdateActionTarget(ctx workflow.Context, input *securityhub.UpdateActionTargetInput) (*securityhub.UpdateActionTargetOutput, error)
    UpdateActionTargetAsync(ctx workflow.Context, input *securityhub.UpdateActionTargetInput) *SecurityhubUpdateActionTargetResult

    UpdateFindings(ctx workflow.Context, input *securityhub.UpdateFindingsInput) (*securityhub.UpdateFindingsOutput, error)
    UpdateFindingsAsync(ctx workflow.Context, input *securityhub.UpdateFindingsInput) *SecurityhubUpdateFindingsResult

    UpdateInsight(ctx workflow.Context, input *securityhub.UpdateInsightInput) (*securityhub.UpdateInsightOutput, error)
    UpdateInsightAsync(ctx workflow.Context, input *securityhub.UpdateInsightInput) *SecurityhubUpdateInsightResult

    UpdateSecurityHubConfiguration(ctx workflow.Context, input *securityhub.UpdateSecurityHubConfigurationInput) (*securityhub.UpdateSecurityHubConfigurationOutput, error)
    UpdateSecurityHubConfigurationAsync(ctx workflow.Context, input *securityhub.UpdateSecurityHubConfigurationInput) *SecurityhubUpdateSecurityHubConfigurationResult

    UpdateStandardsControl(ctx workflow.Context, input *securityhub.UpdateStandardsControlInput) (*securityhub.UpdateStandardsControlOutput, error)
    UpdateStandardsControlAsync(ctx workflow.Context, input *securityhub.UpdateStandardsControlInput) *SecurityhubUpdateStandardsControlResult
}
type SecurityhubAcceptInvitationResult struct {
	Result workflow.Future
}

func (r *SecurityhubAcceptInvitationResult) Get(ctx workflow.Context) (*securityhub.AcceptInvitationOutput, error) {
    var output securityhub.AcceptInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubBatchDisableStandardsResult struct {
	Result workflow.Future
}

func (r *SecurityhubBatchDisableStandardsResult) Get(ctx workflow.Context) (*securityhub.BatchDisableStandardsOutput, error) {
    var output securityhub.BatchDisableStandardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubBatchEnableStandardsResult struct {
	Result workflow.Future
}

func (r *SecurityhubBatchEnableStandardsResult) Get(ctx workflow.Context) (*securityhub.BatchEnableStandardsOutput, error) {
    var output securityhub.BatchEnableStandardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubBatchImportFindingsResult struct {
	Result workflow.Future
}

func (r *SecurityhubBatchImportFindingsResult) Get(ctx workflow.Context) (*securityhub.BatchImportFindingsOutput, error) {
    var output securityhub.BatchImportFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubBatchUpdateFindingsResult struct {
	Result workflow.Future
}

func (r *SecurityhubBatchUpdateFindingsResult) Get(ctx workflow.Context) (*securityhub.BatchUpdateFindingsOutput, error) {
    var output securityhub.BatchUpdateFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubCreateActionTargetResult struct {
	Result workflow.Future
}

func (r *SecurityhubCreateActionTargetResult) Get(ctx workflow.Context) (*securityhub.CreateActionTargetOutput, error) {
    var output securityhub.CreateActionTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubCreateInsightResult struct {
	Result workflow.Future
}

func (r *SecurityhubCreateInsightResult) Get(ctx workflow.Context) (*securityhub.CreateInsightOutput, error) {
    var output securityhub.CreateInsightOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubCreateMembersResult struct {
	Result workflow.Future
}

func (r *SecurityhubCreateMembersResult) Get(ctx workflow.Context) (*securityhub.CreateMembersOutput, error) {
    var output securityhub.CreateMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDeclineInvitationsResult struct {
	Result workflow.Future
}

func (r *SecurityhubDeclineInvitationsResult) Get(ctx workflow.Context) (*securityhub.DeclineInvitationsOutput, error) {
    var output securityhub.DeclineInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDeleteActionTargetResult struct {
	Result workflow.Future
}

func (r *SecurityhubDeleteActionTargetResult) Get(ctx workflow.Context) (*securityhub.DeleteActionTargetOutput, error) {
    var output securityhub.DeleteActionTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDeleteInsightResult struct {
	Result workflow.Future
}

func (r *SecurityhubDeleteInsightResult) Get(ctx workflow.Context) (*securityhub.DeleteInsightOutput, error) {
    var output securityhub.DeleteInsightOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDeleteInvitationsResult struct {
	Result workflow.Future
}

func (r *SecurityhubDeleteInvitationsResult) Get(ctx workflow.Context) (*securityhub.DeleteInvitationsOutput, error) {
    var output securityhub.DeleteInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDeleteMembersResult struct {
	Result workflow.Future
}

func (r *SecurityhubDeleteMembersResult) Get(ctx workflow.Context) (*securityhub.DeleteMembersOutput, error) {
    var output securityhub.DeleteMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDescribeActionTargetsResult struct {
	Result workflow.Future
}

func (r *SecurityhubDescribeActionTargetsResult) Get(ctx workflow.Context) (*securityhub.DescribeActionTargetsOutput, error) {
    var output securityhub.DescribeActionTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDescribeHubResult struct {
	Result workflow.Future
}

func (r *SecurityhubDescribeHubResult) Get(ctx workflow.Context) (*securityhub.DescribeHubOutput, error) {
    var output securityhub.DescribeHubOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDescribeProductsResult struct {
	Result workflow.Future
}

func (r *SecurityhubDescribeProductsResult) Get(ctx workflow.Context) (*securityhub.DescribeProductsOutput, error) {
    var output securityhub.DescribeProductsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDescribeStandardsResult struct {
	Result workflow.Future
}

func (r *SecurityhubDescribeStandardsResult) Get(ctx workflow.Context) (*securityhub.DescribeStandardsOutput, error) {
    var output securityhub.DescribeStandardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDescribeStandardsControlsResult struct {
	Result workflow.Future
}

func (r *SecurityhubDescribeStandardsControlsResult) Get(ctx workflow.Context) (*securityhub.DescribeStandardsControlsOutput, error) {
    var output securityhub.DescribeStandardsControlsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDisableImportFindingsForProductResult struct {
	Result workflow.Future
}

func (r *SecurityhubDisableImportFindingsForProductResult) Get(ctx workflow.Context) (*securityhub.DisableImportFindingsForProductOutput, error) {
    var output securityhub.DisableImportFindingsForProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDisableSecurityHubResult struct {
	Result workflow.Future
}

func (r *SecurityhubDisableSecurityHubResult) Get(ctx workflow.Context) (*securityhub.DisableSecurityHubOutput, error) {
    var output securityhub.DisableSecurityHubOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDisassociateFromMasterAccountResult struct {
	Result workflow.Future
}

func (r *SecurityhubDisassociateFromMasterAccountResult) Get(ctx workflow.Context) (*securityhub.DisassociateFromMasterAccountOutput, error) {
    var output securityhub.DisassociateFromMasterAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubDisassociateMembersResult struct {
	Result workflow.Future
}

func (r *SecurityhubDisassociateMembersResult) Get(ctx workflow.Context) (*securityhub.DisassociateMembersOutput, error) {
    var output securityhub.DisassociateMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubEnableImportFindingsForProductResult struct {
	Result workflow.Future
}

func (r *SecurityhubEnableImportFindingsForProductResult) Get(ctx workflow.Context) (*securityhub.EnableImportFindingsForProductOutput, error) {
    var output securityhub.EnableImportFindingsForProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubEnableSecurityHubResult struct {
	Result workflow.Future
}

func (r *SecurityhubEnableSecurityHubResult) Get(ctx workflow.Context) (*securityhub.EnableSecurityHubOutput, error) {
    var output securityhub.EnableSecurityHubOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubGetEnabledStandardsResult struct {
	Result workflow.Future
}

func (r *SecurityhubGetEnabledStandardsResult) Get(ctx workflow.Context) (*securityhub.GetEnabledStandardsOutput, error) {
    var output securityhub.GetEnabledStandardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubGetFindingsResult struct {
	Result workflow.Future
}

func (r *SecurityhubGetFindingsResult) Get(ctx workflow.Context) (*securityhub.GetFindingsOutput, error) {
    var output securityhub.GetFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubGetInsightResultsResult struct {
	Result workflow.Future
}

func (r *SecurityhubGetInsightResultsResult) Get(ctx workflow.Context) (*securityhub.GetInsightResultsOutput, error) {
    var output securityhub.GetInsightResultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubGetInsightsResult struct {
	Result workflow.Future
}

func (r *SecurityhubGetInsightsResult) Get(ctx workflow.Context) (*securityhub.GetInsightsOutput, error) {
    var output securityhub.GetInsightsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubGetInvitationsCountResult struct {
	Result workflow.Future
}

func (r *SecurityhubGetInvitationsCountResult) Get(ctx workflow.Context) (*securityhub.GetInvitationsCountOutput, error) {
    var output securityhub.GetInvitationsCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubGetMasterAccountResult struct {
	Result workflow.Future
}

func (r *SecurityhubGetMasterAccountResult) Get(ctx workflow.Context) (*securityhub.GetMasterAccountOutput, error) {
    var output securityhub.GetMasterAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubGetMembersResult struct {
	Result workflow.Future
}

func (r *SecurityhubGetMembersResult) Get(ctx workflow.Context) (*securityhub.GetMembersOutput, error) {
    var output securityhub.GetMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubInviteMembersResult struct {
	Result workflow.Future
}

func (r *SecurityhubInviteMembersResult) Get(ctx workflow.Context) (*securityhub.InviteMembersOutput, error) {
    var output securityhub.InviteMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubListEnabledProductsForImportResult struct {
	Result workflow.Future
}

func (r *SecurityhubListEnabledProductsForImportResult) Get(ctx workflow.Context) (*securityhub.ListEnabledProductsForImportOutput, error) {
    var output securityhub.ListEnabledProductsForImportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubListInvitationsResult struct {
	Result workflow.Future
}

func (r *SecurityhubListInvitationsResult) Get(ctx workflow.Context) (*securityhub.ListInvitationsOutput, error) {
    var output securityhub.ListInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubListMembersResult struct {
	Result workflow.Future
}

func (r *SecurityhubListMembersResult) Get(ctx workflow.Context) (*securityhub.ListMembersOutput, error) {
    var output securityhub.ListMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SecurityhubListTagsForResourceResult) Get(ctx workflow.Context) (*securityhub.ListTagsForResourceOutput, error) {
    var output securityhub.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubTagResourceResult struct {
	Result workflow.Future
}

func (r *SecurityhubTagResourceResult) Get(ctx workflow.Context) (*securityhub.TagResourceOutput, error) {
    var output securityhub.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubUntagResourceResult struct {
	Result workflow.Future
}

func (r *SecurityhubUntagResourceResult) Get(ctx workflow.Context) (*securityhub.UntagResourceOutput, error) {
    var output securityhub.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubUpdateActionTargetResult struct {
	Result workflow.Future
}

func (r *SecurityhubUpdateActionTargetResult) Get(ctx workflow.Context) (*securityhub.UpdateActionTargetOutput, error) {
    var output securityhub.UpdateActionTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubUpdateFindingsResult struct {
	Result workflow.Future
}

func (r *SecurityhubUpdateFindingsResult) Get(ctx workflow.Context) (*securityhub.UpdateFindingsOutput, error) {
    var output securityhub.UpdateFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubUpdateInsightResult struct {
	Result workflow.Future
}

func (r *SecurityhubUpdateInsightResult) Get(ctx workflow.Context) (*securityhub.UpdateInsightOutput, error) {
    var output securityhub.UpdateInsightOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubUpdateSecurityHubConfigurationResult struct {
	Result workflow.Future
}

func (r *SecurityhubUpdateSecurityHubConfigurationResult) Get(ctx workflow.Context) (*securityhub.UpdateSecurityHubConfigurationOutput, error) {
    var output securityhub.UpdateSecurityHubConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecurityhubUpdateStandardsControlResult struct {
	Result workflow.Future
}

func (r *SecurityhubUpdateStandardsControlResult) Get(ctx workflow.Context) (*securityhub.UpdateStandardsControlOutput, error) {
    var output securityhub.UpdateStandardsControlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SecurityHubStub struct {
    activities SecurityHubClient
}

func NewSecurityHubStub() SecurityHubClient {
    return &SecurityHubStub{}
}

func (a *SecurityHubStub) AcceptInvitation(ctx workflow.Context, input *securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error) {
    var output securityhub.AcceptInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) BatchDisableStandards(ctx workflow.Context, input *securityhub.BatchDisableStandardsInput) (*securityhub.BatchDisableStandardsOutput, error) {
    var output securityhub.BatchDisableStandardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDisableStandards, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) BatchEnableStandards(ctx workflow.Context, input *securityhub.BatchEnableStandardsInput) (*securityhub.BatchEnableStandardsOutput, error) {
    var output securityhub.BatchEnableStandardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchEnableStandards, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) BatchImportFindings(ctx workflow.Context, input *securityhub.BatchImportFindingsInput) (*securityhub.BatchImportFindingsOutput, error) {
    var output securityhub.BatchImportFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchImportFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) BatchUpdateFindings(ctx workflow.Context, input *securityhub.BatchUpdateFindingsInput) (*securityhub.BatchUpdateFindingsOutput, error) {
    var output securityhub.BatchUpdateFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchUpdateFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) CreateActionTarget(ctx workflow.Context, input *securityhub.CreateActionTargetInput) (*securityhub.CreateActionTargetOutput, error) {
    var output securityhub.CreateActionTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateActionTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) CreateInsight(ctx workflow.Context, input *securityhub.CreateInsightInput) (*securityhub.CreateInsightOutput, error) {
    var output securityhub.CreateInsightOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInsight, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) CreateMembers(ctx workflow.Context, input *securityhub.CreateMembersInput) (*securityhub.CreateMembersOutput, error) {
    var output securityhub.CreateMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DeclineInvitations(ctx workflow.Context, input *securityhub.DeclineInvitationsInput) (*securityhub.DeclineInvitationsOutput, error) {
    var output securityhub.DeclineInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeclineInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DeleteActionTarget(ctx workflow.Context, input *securityhub.DeleteActionTargetInput) (*securityhub.DeleteActionTargetOutput, error) {
    var output securityhub.DeleteActionTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteActionTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DeleteInsight(ctx workflow.Context, input *securityhub.DeleteInsightInput) (*securityhub.DeleteInsightOutput, error) {
    var output securityhub.DeleteInsightOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInsight, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DeleteInvitations(ctx workflow.Context, input *securityhub.DeleteInvitationsInput) (*securityhub.DeleteInvitationsOutput, error) {
    var output securityhub.DeleteInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DeleteMembers(ctx workflow.Context, input *securityhub.DeleteMembersInput) (*securityhub.DeleteMembersOutput, error) {
    var output securityhub.DeleteMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DescribeActionTargets(ctx workflow.Context, input *securityhub.DescribeActionTargetsInput) (*securityhub.DescribeActionTargetsOutput, error) {
    var output securityhub.DescribeActionTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeActionTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DescribeHub(ctx workflow.Context, input *securityhub.DescribeHubInput) (*securityhub.DescribeHubOutput, error) {
    var output securityhub.DescribeHubOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHub, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DescribeProducts(ctx workflow.Context, input *securityhub.DescribeProductsInput) (*securityhub.DescribeProductsOutput, error) {
    var output securityhub.DescribeProductsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProducts, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DescribeStandards(ctx workflow.Context, input *securityhub.DescribeStandardsInput) (*securityhub.DescribeStandardsOutput, error) {
    var output securityhub.DescribeStandardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStandards, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DescribeStandardsControls(ctx workflow.Context, input *securityhub.DescribeStandardsControlsInput) (*securityhub.DescribeStandardsControlsOutput, error) {
    var output securityhub.DescribeStandardsControlsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStandardsControls, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DisableImportFindingsForProduct(ctx workflow.Context, input *securityhub.DisableImportFindingsForProductInput) (*securityhub.DisableImportFindingsForProductOutput, error) {
    var output securityhub.DisableImportFindingsForProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableImportFindingsForProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DisableSecurityHub(ctx workflow.Context, input *securityhub.DisableSecurityHubInput) (*securityhub.DisableSecurityHubOutput, error) {
    var output securityhub.DisableSecurityHubOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableSecurityHub, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DisassociateFromMasterAccount(ctx workflow.Context, input *securityhub.DisassociateFromMasterAccountInput) (*securityhub.DisassociateFromMasterAccountOutput, error) {
    var output securityhub.DisassociateFromMasterAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateFromMasterAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) DisassociateMembers(ctx workflow.Context, input *securityhub.DisassociateMembersInput) (*securityhub.DisassociateMembersOutput, error) {
    var output securityhub.DisassociateMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) EnableImportFindingsForProduct(ctx workflow.Context, input *securityhub.EnableImportFindingsForProductInput) (*securityhub.EnableImportFindingsForProductOutput, error) {
    var output securityhub.EnableImportFindingsForProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableImportFindingsForProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) EnableSecurityHub(ctx workflow.Context, input *securityhub.EnableSecurityHubInput) (*securityhub.EnableSecurityHubOutput, error) {
    var output securityhub.EnableSecurityHubOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableSecurityHub, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) GetEnabledStandards(ctx workflow.Context, input *securityhub.GetEnabledStandardsInput) (*securityhub.GetEnabledStandardsOutput, error) {
    var output securityhub.GetEnabledStandardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEnabledStandards, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) GetFindings(ctx workflow.Context, input *securityhub.GetFindingsInput) (*securityhub.GetFindingsOutput, error) {
    var output securityhub.GetFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) GetInsightResults(ctx workflow.Context, input *securityhub.GetInsightResultsInput) (*securityhub.GetInsightResultsOutput, error) {
    var output securityhub.GetInsightResultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInsightResults, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) GetInsights(ctx workflow.Context, input *securityhub.GetInsightsInput) (*securityhub.GetInsightsOutput, error) {
    var output securityhub.GetInsightsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInsights, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) GetInvitationsCount(ctx workflow.Context, input *securityhub.GetInvitationsCountInput) (*securityhub.GetInvitationsCountOutput, error) {
    var output securityhub.GetInvitationsCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInvitationsCount, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) GetMasterAccount(ctx workflow.Context, input *securityhub.GetMasterAccountInput) (*securityhub.GetMasterAccountOutput, error) {
    var output securityhub.GetMasterAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMasterAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) GetMembers(ctx workflow.Context, input *securityhub.GetMembersInput) (*securityhub.GetMembersOutput, error) {
    var output securityhub.GetMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) InviteMembers(ctx workflow.Context, input *securityhub.InviteMembersInput) (*securityhub.InviteMembersOutput, error) {
    var output securityhub.InviteMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InviteMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) ListEnabledProductsForImport(ctx workflow.Context, input *securityhub.ListEnabledProductsForImportInput) (*securityhub.ListEnabledProductsForImportOutput, error) {
    var output securityhub.ListEnabledProductsForImportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEnabledProductsForImport, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) ListInvitations(ctx workflow.Context, input *securityhub.ListInvitationsInput) (*securityhub.ListInvitationsOutput, error) {
    var output securityhub.ListInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) ListMembers(ctx workflow.Context, input *securityhub.ListMembersInput) (*securityhub.ListMembersOutput, error) {
    var output securityhub.ListMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) ListTagsForResource(ctx workflow.Context, input *securityhub.ListTagsForResourceInput) (*securityhub.ListTagsForResourceOutput, error) {
    var output securityhub.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) TagResource(ctx workflow.Context, input *securityhub.TagResourceInput) (*securityhub.TagResourceOutput, error) {
    var output securityhub.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) UntagResource(ctx workflow.Context, input *securityhub.UntagResourceInput) (*securityhub.UntagResourceOutput, error) {
    var output securityhub.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) UpdateActionTarget(ctx workflow.Context, input *securityhub.UpdateActionTargetInput) (*securityhub.UpdateActionTargetOutput, error) {
    var output securityhub.UpdateActionTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateActionTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) UpdateFindings(ctx workflow.Context, input *securityhub.UpdateFindingsInput) (*securityhub.UpdateFindingsOutput, error) {
    var output securityhub.UpdateFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) UpdateInsight(ctx workflow.Context, input *securityhub.UpdateInsightInput) (*securityhub.UpdateInsightOutput, error) {
    var output securityhub.UpdateInsightOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateInsight, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) UpdateSecurityHubConfiguration(ctx workflow.Context, input *securityhub.UpdateSecurityHubConfigurationInput) (*securityhub.UpdateSecurityHubConfigurationOutput, error) {
    var output securityhub.UpdateSecurityHubConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSecurityHubConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *SecurityHubStub) UpdateStandardsControl(ctx workflow.Context, input *securityhub.UpdateStandardsControlInput) (*securityhub.UpdateStandardsControlOutput, error) {
    var output securityhub.UpdateStandardsControlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStandardsControl, input).Get(ctx, &output)
    return &output, err
}
