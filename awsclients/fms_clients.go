package awsclients

import (
	"github.com/aws/aws-sdk-go/service/fms"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type FMSClient interface {
       AssociateAdminAccount(ctx workflow.Context, input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error)
       AssociateAdminAccountAsync(ctx workflow.Context, input *fms.AssociateAdminAccountInput) *FmsAssociateAdminAccountResult

       DeleteAppsList(ctx workflow.Context, input *fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error)
       DeleteAppsListAsync(ctx workflow.Context, input *fms.DeleteAppsListInput) *FmsDeleteAppsListResult

       DeleteNotificationChannel(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error)
       DeleteNotificationChannelAsync(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) *FmsDeleteNotificationChannelResult

       DeletePolicy(ctx workflow.Context, input *fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error)
       DeletePolicyAsync(ctx workflow.Context, input *fms.DeletePolicyInput) *FmsDeletePolicyResult

       DeleteProtocolsList(ctx workflow.Context, input *fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error)
       DeleteProtocolsListAsync(ctx workflow.Context, input *fms.DeleteProtocolsListInput) *FmsDeleteProtocolsListResult

       DisassociateAdminAccount(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error)
       DisassociateAdminAccountAsync(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) *FmsDisassociateAdminAccountResult

       GetAdminAccount(ctx workflow.Context, input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error)
       GetAdminAccountAsync(ctx workflow.Context, input *fms.GetAdminAccountInput) *FmsGetAdminAccountResult

       GetAppsList(ctx workflow.Context, input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error)
       GetAppsListAsync(ctx workflow.Context, input *fms.GetAppsListInput) *FmsGetAppsListResult

       GetComplianceDetail(ctx workflow.Context, input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error)
       GetComplianceDetailAsync(ctx workflow.Context, input *fms.GetComplianceDetailInput) *FmsGetComplianceDetailResult

       GetNotificationChannel(ctx workflow.Context, input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error)
       GetNotificationChannelAsync(ctx workflow.Context, input *fms.GetNotificationChannelInput) *FmsGetNotificationChannelResult

       GetPolicy(ctx workflow.Context, input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error)
       GetPolicyAsync(ctx workflow.Context, input *fms.GetPolicyInput) *FmsGetPolicyResult

       GetProtectionStatus(ctx workflow.Context, input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error)
       GetProtectionStatusAsync(ctx workflow.Context, input *fms.GetProtectionStatusInput) *FmsGetProtectionStatusResult

       GetProtocolsList(ctx workflow.Context, input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error)
       GetProtocolsListAsync(ctx workflow.Context, input *fms.GetProtocolsListInput) *FmsGetProtocolsListResult

       GetViolationDetails(ctx workflow.Context, input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error)
       GetViolationDetailsAsync(ctx workflow.Context, input *fms.GetViolationDetailsInput) *FmsGetViolationDetailsResult

       ListAppsLists(ctx workflow.Context, input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error)
       ListAppsListsAsync(ctx workflow.Context, input *fms.ListAppsListsInput) *FmsListAppsListsResult

       ListComplianceStatus(ctx workflow.Context, input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error)
       ListComplianceStatusAsync(ctx workflow.Context, input *fms.ListComplianceStatusInput) *FmsListComplianceStatusResult

       ListMemberAccounts(ctx workflow.Context, input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error)
       ListMemberAccountsAsync(ctx workflow.Context, input *fms.ListMemberAccountsInput) *FmsListMemberAccountsResult

       ListPolicies(ctx workflow.Context, input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error)
       ListPoliciesAsync(ctx workflow.Context, input *fms.ListPoliciesInput) *FmsListPoliciesResult

       ListProtocolsLists(ctx workflow.Context, input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error)
       ListProtocolsListsAsync(ctx workflow.Context, input *fms.ListProtocolsListsInput) *FmsListProtocolsListsResult

       ListTagsForResource(ctx workflow.Context, input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *fms.ListTagsForResourceInput) *FmsListTagsForResourceResult

       PutAppsList(ctx workflow.Context, input *fms.PutAppsListInput) (*fms.PutAppsListOutput, error)
       PutAppsListAsync(ctx workflow.Context, input *fms.PutAppsListInput) *FmsPutAppsListResult

       PutNotificationChannel(ctx workflow.Context, input *fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error)
       PutNotificationChannelAsync(ctx workflow.Context, input *fms.PutNotificationChannelInput) *FmsPutNotificationChannelResult

       PutPolicy(ctx workflow.Context, input *fms.PutPolicyInput) (*fms.PutPolicyOutput, error)
       PutPolicyAsync(ctx workflow.Context, input *fms.PutPolicyInput) *FmsPutPolicyResult

       PutProtocolsList(ctx workflow.Context, input *fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error)
       PutProtocolsListAsync(ctx workflow.Context, input *fms.PutProtocolsListInput) *FmsPutProtocolsListResult

       TagResource(ctx workflow.Context, input *fms.TagResourceInput) (*fms.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *fms.TagResourceInput) *FmsTagResourceResult

       UntagResource(ctx workflow.Context, input *fms.UntagResourceInput) (*fms.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *fms.UntagResourceInput) *FmsUntagResourceResult
}

type FmsAssociateAdminAccountResult struct {
	Result workflow.Future
}

func (r *FmsAssociateAdminAccountResult) Get(ctx workflow.Context) (*fms.AssociateAdminAccountOutput, error) {
    var output fms.AssociateAdminAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsDeleteAppsListResult struct {
	Result workflow.Future
}

func (r *FmsDeleteAppsListResult) Get(ctx workflow.Context) (*fms.DeleteAppsListOutput, error) {
    var output fms.DeleteAppsListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsDeleteNotificationChannelResult struct {
	Result workflow.Future
}

func (r *FmsDeleteNotificationChannelResult) Get(ctx workflow.Context) (*fms.DeleteNotificationChannelOutput, error) {
    var output fms.DeleteNotificationChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsDeletePolicyResult struct {
	Result workflow.Future
}

func (r *FmsDeletePolicyResult) Get(ctx workflow.Context) (*fms.DeletePolicyOutput, error) {
    var output fms.DeletePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsDeleteProtocolsListResult struct {
	Result workflow.Future
}

func (r *FmsDeleteProtocolsListResult) Get(ctx workflow.Context) (*fms.DeleteProtocolsListOutput, error) {
    var output fms.DeleteProtocolsListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsDisassociateAdminAccountResult struct {
	Result workflow.Future
}

func (r *FmsDisassociateAdminAccountResult) Get(ctx workflow.Context) (*fms.DisassociateAdminAccountOutput, error) {
    var output fms.DisassociateAdminAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetAdminAccountResult struct {
	Result workflow.Future
}

func (r *FmsGetAdminAccountResult) Get(ctx workflow.Context) (*fms.GetAdminAccountOutput, error) {
    var output fms.GetAdminAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetAppsListResult struct {
	Result workflow.Future
}

func (r *FmsGetAppsListResult) Get(ctx workflow.Context) (*fms.GetAppsListOutput, error) {
    var output fms.GetAppsListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetComplianceDetailResult struct {
	Result workflow.Future
}

func (r *FmsGetComplianceDetailResult) Get(ctx workflow.Context) (*fms.GetComplianceDetailOutput, error) {
    var output fms.GetComplianceDetailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetNotificationChannelResult struct {
	Result workflow.Future
}

func (r *FmsGetNotificationChannelResult) Get(ctx workflow.Context) (*fms.GetNotificationChannelOutput, error) {
    var output fms.GetNotificationChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetPolicyResult struct {
	Result workflow.Future
}

func (r *FmsGetPolicyResult) Get(ctx workflow.Context) (*fms.GetPolicyOutput, error) {
    var output fms.GetPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetProtectionStatusResult struct {
	Result workflow.Future
}

func (r *FmsGetProtectionStatusResult) Get(ctx workflow.Context) (*fms.GetProtectionStatusOutput, error) {
    var output fms.GetProtectionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetProtocolsListResult struct {
	Result workflow.Future
}

func (r *FmsGetProtocolsListResult) Get(ctx workflow.Context) (*fms.GetProtocolsListOutput, error) {
    var output fms.GetProtocolsListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsGetViolationDetailsResult struct {
	Result workflow.Future
}

func (r *FmsGetViolationDetailsResult) Get(ctx workflow.Context) (*fms.GetViolationDetailsOutput, error) {
    var output fms.GetViolationDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsListAppsListsResult struct {
	Result workflow.Future
}

func (r *FmsListAppsListsResult) Get(ctx workflow.Context) (*fms.ListAppsListsOutput, error) {
    var output fms.ListAppsListsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsListComplianceStatusResult struct {
	Result workflow.Future
}

func (r *FmsListComplianceStatusResult) Get(ctx workflow.Context) (*fms.ListComplianceStatusOutput, error) {
    var output fms.ListComplianceStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsListMemberAccountsResult struct {
	Result workflow.Future
}

func (r *FmsListMemberAccountsResult) Get(ctx workflow.Context) (*fms.ListMemberAccountsOutput, error) {
    var output fms.ListMemberAccountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsListPoliciesResult struct {
	Result workflow.Future
}

func (r *FmsListPoliciesResult) Get(ctx workflow.Context) (*fms.ListPoliciesOutput, error) {
    var output fms.ListPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsListProtocolsListsResult struct {
	Result workflow.Future
}

func (r *FmsListProtocolsListsResult) Get(ctx workflow.Context) (*fms.ListProtocolsListsOutput, error) {
    var output fms.ListProtocolsListsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *FmsListTagsForResourceResult) Get(ctx workflow.Context) (*fms.ListTagsForResourceOutput, error) {
    var output fms.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsPutAppsListResult struct {
	Result workflow.Future
}

func (r *FmsPutAppsListResult) Get(ctx workflow.Context) (*fms.PutAppsListOutput, error) {
    var output fms.PutAppsListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsPutNotificationChannelResult struct {
	Result workflow.Future
}

func (r *FmsPutNotificationChannelResult) Get(ctx workflow.Context) (*fms.PutNotificationChannelOutput, error) {
    var output fms.PutNotificationChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsPutPolicyResult struct {
	Result workflow.Future
}

func (r *FmsPutPolicyResult) Get(ctx workflow.Context) (*fms.PutPolicyOutput, error) {
    var output fms.PutPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsPutProtocolsListResult struct {
	Result workflow.Future
}

func (r *FmsPutProtocolsListResult) Get(ctx workflow.Context) (*fms.PutProtocolsListOutput, error) {
    var output fms.PutProtocolsListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsTagResourceResult struct {
	Result workflow.Future
}

func (r *FmsTagResourceResult) Get(ctx workflow.Context) (*fms.TagResourceOutput, error) {
    var output fms.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FmsUntagResourceResult struct {
	Result workflow.Future
}

func (r *FmsUntagResourceResult) Get(ctx workflow.Context) (*fms.UntagResourceOutput, error) {
    var output fms.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FMSStub struct {
    activities awsactivities.FMSActivities
}

func NewFMSStub() FMSClient {
    return &FMSStub{}
}

func (a *FMSStub) AssociateAdminAccount(ctx workflow.Context, input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
    var output fms.AssociateAdminAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateAdminAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) AssociateAdminAccountAsync(ctx workflow.Context, input *fms.AssociateAdminAccountInput) *FmsAssociateAdminAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateAdminAccount, input)
    return &FmsAssociateAdminAccountResult{Result: future}
}

func (a *FMSStub) DeleteAppsList(ctx workflow.Context, input *fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error) {
    var output fms.DeleteAppsListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAppsList, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) DeleteAppsListAsync(ctx workflow.Context, input *fms.DeleteAppsListInput) *FmsDeleteAppsListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAppsList, input)
    return &FmsDeleteAppsListResult{Result: future}
}

func (a *FMSStub) DeleteNotificationChannel(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error) {
    var output fms.DeleteNotificationChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNotificationChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) DeleteNotificationChannelAsync(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) *FmsDeleteNotificationChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNotificationChannel, input)
    return &FmsDeleteNotificationChannelResult{Result: future}
}

func (a *FMSStub) DeletePolicy(ctx workflow.Context, input *fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error) {
    var output fms.DeletePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) DeletePolicyAsync(ctx workflow.Context, input *fms.DeletePolicyInput) *FmsDeletePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input)
    return &FmsDeletePolicyResult{Result: future}
}

func (a *FMSStub) DeleteProtocolsList(ctx workflow.Context, input *fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error) {
    var output fms.DeleteProtocolsListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProtocolsList, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) DeleteProtocolsListAsync(ctx workflow.Context, input *fms.DeleteProtocolsListInput) *FmsDeleteProtocolsListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProtocolsList, input)
    return &FmsDeleteProtocolsListResult{Result: future}
}

func (a *FMSStub) DisassociateAdminAccount(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error) {
    var output fms.DisassociateAdminAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateAdminAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) DisassociateAdminAccountAsync(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) *FmsDisassociateAdminAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateAdminAccount, input)
    return &FmsDisassociateAdminAccountResult{Result: future}
}

func (a *FMSStub) GetAdminAccount(ctx workflow.Context, input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error) {
    var output fms.GetAdminAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAdminAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetAdminAccountAsync(ctx workflow.Context, input *fms.GetAdminAccountInput) *FmsGetAdminAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAdminAccount, input)
    return &FmsGetAdminAccountResult{Result: future}
}

func (a *FMSStub) GetAppsList(ctx workflow.Context, input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error) {
    var output fms.GetAppsListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAppsList, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetAppsListAsync(ctx workflow.Context, input *fms.GetAppsListInput) *FmsGetAppsListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAppsList, input)
    return &FmsGetAppsListResult{Result: future}
}

func (a *FMSStub) GetComplianceDetail(ctx workflow.Context, input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error) {
    var output fms.GetComplianceDetailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComplianceDetail, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetComplianceDetailAsync(ctx workflow.Context, input *fms.GetComplianceDetailInput) *FmsGetComplianceDetailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetComplianceDetail, input)
    return &FmsGetComplianceDetailResult{Result: future}
}

func (a *FMSStub) GetNotificationChannel(ctx workflow.Context, input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error) {
    var output fms.GetNotificationChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNotificationChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetNotificationChannelAsync(ctx workflow.Context, input *fms.GetNotificationChannelInput) *FmsGetNotificationChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetNotificationChannel, input)
    return &FmsGetNotificationChannelResult{Result: future}
}

func (a *FMSStub) GetPolicy(ctx workflow.Context, input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error) {
    var output fms.GetPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetPolicyAsync(ctx workflow.Context, input *fms.GetPolicyInput) *FmsGetPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input)
    return &FmsGetPolicyResult{Result: future}
}

func (a *FMSStub) GetProtectionStatus(ctx workflow.Context, input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error) {
    var output fms.GetProtectionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProtectionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetProtectionStatusAsync(ctx workflow.Context, input *fms.GetProtectionStatusInput) *FmsGetProtectionStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetProtectionStatus, input)
    return &FmsGetProtectionStatusResult{Result: future}
}

func (a *FMSStub) GetProtocolsList(ctx workflow.Context, input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error) {
    var output fms.GetProtocolsListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProtocolsList, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetProtocolsListAsync(ctx workflow.Context, input *fms.GetProtocolsListInput) *FmsGetProtocolsListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetProtocolsList, input)
    return &FmsGetProtocolsListResult{Result: future}
}

func (a *FMSStub) GetViolationDetails(ctx workflow.Context, input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error) {
    var output fms.GetViolationDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetViolationDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) GetViolationDetailsAsync(ctx workflow.Context, input *fms.GetViolationDetailsInput) *FmsGetViolationDetailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetViolationDetails, input)
    return &FmsGetViolationDetailsResult{Result: future}
}

func (a *FMSStub) ListAppsLists(ctx workflow.Context, input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error) {
    var output fms.ListAppsListsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAppsLists, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) ListAppsListsAsync(ctx workflow.Context, input *fms.ListAppsListsInput) *FmsListAppsListsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAppsLists, input)
    return &FmsListAppsListsResult{Result: future}
}

func (a *FMSStub) ListComplianceStatus(ctx workflow.Context, input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error) {
    var output fms.ListComplianceStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListComplianceStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) ListComplianceStatusAsync(ctx workflow.Context, input *fms.ListComplianceStatusInput) *FmsListComplianceStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListComplianceStatus, input)
    return &FmsListComplianceStatusResult{Result: future}
}

func (a *FMSStub) ListMemberAccounts(ctx workflow.Context, input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error) {
    var output fms.ListMemberAccountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMemberAccounts, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) ListMemberAccountsAsync(ctx workflow.Context, input *fms.ListMemberAccountsInput) *FmsListMemberAccountsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMemberAccounts, input)
    return &FmsListMemberAccountsResult{Result: future}
}

func (a *FMSStub) ListPolicies(ctx workflow.Context, input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error) {
    var output fms.ListPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) ListPoliciesAsync(ctx workflow.Context, input *fms.ListPoliciesInput) *FmsListPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPolicies, input)
    return &FmsListPoliciesResult{Result: future}
}

func (a *FMSStub) ListProtocolsLists(ctx workflow.Context, input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error) {
    var output fms.ListProtocolsListsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProtocolsLists, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) ListProtocolsListsAsync(ctx workflow.Context, input *fms.ListProtocolsListsInput) *FmsListProtocolsListsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProtocolsLists, input)
    return &FmsListProtocolsListsResult{Result: future}
}

func (a *FMSStub) ListTagsForResource(ctx workflow.Context, input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error) {
    var output fms.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) ListTagsForResourceAsync(ctx workflow.Context, input *fms.ListTagsForResourceInput) *FmsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &FmsListTagsForResourceResult{Result: future}
}

func (a *FMSStub) PutAppsList(ctx workflow.Context, input *fms.PutAppsListInput) (*fms.PutAppsListOutput, error) {
    var output fms.PutAppsListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAppsList, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) PutAppsListAsync(ctx workflow.Context, input *fms.PutAppsListInput) *FmsPutAppsListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAppsList, input)
    return &FmsPutAppsListResult{Result: future}
}

func (a *FMSStub) PutNotificationChannel(ctx workflow.Context, input *fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error) {
    var output fms.PutNotificationChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutNotificationChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) PutNotificationChannelAsync(ctx workflow.Context, input *fms.PutNotificationChannelInput) *FmsPutNotificationChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutNotificationChannel, input)
    return &FmsPutNotificationChannelResult{Result: future}
}

func (a *FMSStub) PutPolicy(ctx workflow.Context, input *fms.PutPolicyInput) (*fms.PutPolicyOutput, error) {
    var output fms.PutPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) PutPolicyAsync(ctx workflow.Context, input *fms.PutPolicyInput) *FmsPutPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutPolicy, input)
    return &FmsPutPolicyResult{Result: future}
}

func (a *FMSStub) PutProtocolsList(ctx workflow.Context, input *fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error) {
    var output fms.PutProtocolsListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutProtocolsList, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) PutProtocolsListAsync(ctx workflow.Context, input *fms.PutProtocolsListInput) *FmsPutProtocolsListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutProtocolsList, input)
    return &FmsPutProtocolsListResult{Result: future}
}

func (a *FMSStub) TagResource(ctx workflow.Context, input *fms.TagResourceInput) (*fms.TagResourceOutput, error) {
    var output fms.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) TagResourceAsync(ctx workflow.Context, input *fms.TagResourceInput) *FmsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &FmsTagResourceResult{Result: future}
}

func (a *FMSStub) UntagResource(ctx workflow.Context, input *fms.UntagResourceInput) (*fms.UntagResourceOutput, error) {
    var output fms.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *FMSStub) UntagResourceAsync(ctx workflow.Context, input *fms.UntagResourceInput) *FmsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &FmsUntagResourceResult{Result: future}
}
