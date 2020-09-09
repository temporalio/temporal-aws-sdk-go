package awsclients

import (
	"github.com/aws/aws-sdk-go/service/macie"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MacieClient interface {
	AssociateMemberAccount(ctx workflow.Context, input *macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error)
	AssociateMemberAccountAsync(ctx workflow.Context, input *macie.AssociateMemberAccountInput) *MacieAssociateMemberAccountResult

	AssociateS3Resources(ctx workflow.Context, input *macie.AssociateS3ResourcesInput) (*macie.AssociateS3ResourcesOutput, error)
	AssociateS3ResourcesAsync(ctx workflow.Context, input *macie.AssociateS3ResourcesInput) *MacieAssociateS3ResourcesResult

	DisassociateMemberAccount(ctx workflow.Context, input *macie.DisassociateMemberAccountInput) (*macie.DisassociateMemberAccountOutput, error)
	DisassociateMemberAccountAsync(ctx workflow.Context, input *macie.DisassociateMemberAccountInput) *MacieDisassociateMemberAccountResult

	DisassociateS3Resources(ctx workflow.Context, input *macie.DisassociateS3ResourcesInput) (*macie.DisassociateS3ResourcesOutput, error)
	DisassociateS3ResourcesAsync(ctx workflow.Context, input *macie.DisassociateS3ResourcesInput) *MacieDisassociateS3ResourcesResult

	ListMemberAccounts(ctx workflow.Context, input *macie.ListMemberAccountsInput) (*macie.ListMemberAccountsOutput, error)
	ListMemberAccountsAsync(ctx workflow.Context, input *macie.ListMemberAccountsInput) *MacieListMemberAccountsResult

	ListS3Resources(ctx workflow.Context, input *macie.ListS3ResourcesInput) (*macie.ListS3ResourcesOutput, error)
	ListS3ResourcesAsync(ctx workflow.Context, input *macie.ListS3ResourcesInput) *MacieListS3ResourcesResult

	UpdateS3Resources(ctx workflow.Context, input *macie.UpdateS3ResourcesInput) (*macie.UpdateS3ResourcesOutput, error)
	UpdateS3ResourcesAsync(ctx workflow.Context, input *macie.UpdateS3ResourcesInput) *MacieUpdateS3ResourcesResult
}

type MacieAssociateMemberAccountResult struct {
	Result workflow.Future
}

func (r *MacieAssociateMemberAccountResult) Get(ctx workflow.Context) (*macie.AssociateMemberAccountOutput, error) {
	var output macie.AssociateMemberAccountOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MacieAssociateS3ResourcesResult struct {
	Result workflow.Future
}

func (r *MacieAssociateS3ResourcesResult) Get(ctx workflow.Context) (*macie.AssociateS3ResourcesOutput, error) {
	var output macie.AssociateS3ResourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MacieDisassociateMemberAccountResult struct {
	Result workflow.Future
}

func (r *MacieDisassociateMemberAccountResult) Get(ctx workflow.Context) (*macie.DisassociateMemberAccountOutput, error) {
	var output macie.DisassociateMemberAccountOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MacieDisassociateS3ResourcesResult struct {
	Result workflow.Future
}

func (r *MacieDisassociateS3ResourcesResult) Get(ctx workflow.Context) (*macie.DisassociateS3ResourcesOutput, error) {
	var output macie.DisassociateS3ResourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MacieListMemberAccountsResult struct {
	Result workflow.Future
}

func (r *MacieListMemberAccountsResult) Get(ctx workflow.Context) (*macie.ListMemberAccountsOutput, error) {
	var output macie.ListMemberAccountsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MacieListS3ResourcesResult struct {
	Result workflow.Future
}

func (r *MacieListS3ResourcesResult) Get(ctx workflow.Context) (*macie.ListS3ResourcesOutput, error) {
	var output macie.ListS3ResourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MacieUpdateS3ResourcesResult struct {
	Result workflow.Future
}

func (r *MacieUpdateS3ResourcesResult) Get(ctx workflow.Context) (*macie.UpdateS3ResourcesOutput, error) {
	var output macie.UpdateS3ResourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type MacieStub struct {
	activities awsactivities.MacieActivities
}

func NewMacieStub() MacieClient {
	return &MacieStub{}
}

func (a *MacieStub) AssociateMemberAccount(ctx workflow.Context, input *macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error) {
	var output macie.AssociateMemberAccountOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateMemberAccount, input).Get(ctx, &output)
	return &output, err
}

func (a *MacieStub) AssociateMemberAccountAsync(ctx workflow.Context, input *macie.AssociateMemberAccountInput) *MacieAssociateMemberAccountResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateMemberAccount, input)
	return &MacieAssociateMemberAccountResult{Result: future}
}

func (a *MacieStub) AssociateS3Resources(ctx workflow.Context, input *macie.AssociateS3ResourcesInput) (*macie.AssociateS3ResourcesOutput, error) {
	var output macie.AssociateS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateS3Resources, input).Get(ctx, &output)
	return &output, err
}

func (a *MacieStub) AssociateS3ResourcesAsync(ctx workflow.Context, input *macie.AssociateS3ResourcesInput) *MacieAssociateS3ResourcesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateS3Resources, input)
	return &MacieAssociateS3ResourcesResult{Result: future}
}

func (a *MacieStub) DisassociateMemberAccount(ctx workflow.Context, input *macie.DisassociateMemberAccountInput) (*macie.DisassociateMemberAccountOutput, error) {
	var output macie.DisassociateMemberAccountOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateMemberAccount, input).Get(ctx, &output)
	return &output, err
}

func (a *MacieStub) DisassociateMemberAccountAsync(ctx workflow.Context, input *macie.DisassociateMemberAccountInput) *MacieDisassociateMemberAccountResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateMemberAccount, input)
	return &MacieDisassociateMemberAccountResult{Result: future}
}

func (a *MacieStub) DisassociateS3Resources(ctx workflow.Context, input *macie.DisassociateS3ResourcesInput) (*macie.DisassociateS3ResourcesOutput, error) {
	var output macie.DisassociateS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateS3Resources, input).Get(ctx, &output)
	return &output, err
}

func (a *MacieStub) DisassociateS3ResourcesAsync(ctx workflow.Context, input *macie.DisassociateS3ResourcesInput) *MacieDisassociateS3ResourcesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateS3Resources, input)
	return &MacieDisassociateS3ResourcesResult{Result: future}
}

func (a *MacieStub) ListMemberAccounts(ctx workflow.Context, input *macie.ListMemberAccountsInput) (*macie.ListMemberAccountsOutput, error) {
	var output macie.ListMemberAccountsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListMemberAccounts, input).Get(ctx, &output)
	return &output, err
}

func (a *MacieStub) ListMemberAccountsAsync(ctx workflow.Context, input *macie.ListMemberAccountsInput) *MacieListMemberAccountsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListMemberAccounts, input)
	return &MacieListMemberAccountsResult{Result: future}
}

func (a *MacieStub) ListS3Resources(ctx workflow.Context, input *macie.ListS3ResourcesInput) (*macie.ListS3ResourcesOutput, error) {
	var output macie.ListS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListS3Resources, input).Get(ctx, &output)
	return &output, err
}

func (a *MacieStub) ListS3ResourcesAsync(ctx workflow.Context, input *macie.ListS3ResourcesInput) *MacieListS3ResourcesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListS3Resources, input)
	return &MacieListS3ResourcesResult{Result: future}
}

func (a *MacieStub) UpdateS3Resources(ctx workflow.Context, input *macie.UpdateS3ResourcesInput) (*macie.UpdateS3ResourcesOutput, error) {
	var output macie.UpdateS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateS3Resources, input).Get(ctx, &output)
	return &output, err
}

func (a *MacieStub) UpdateS3ResourcesAsync(ctx workflow.Context, input *macie.UpdateS3ResourcesInput) *MacieUpdateS3ResourcesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateS3Resources, input)
	return &MacieUpdateS3ResourcesResult{Result: future}
}
