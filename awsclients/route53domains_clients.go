package awsclients

import (
	"github.com/aws/aws-sdk-go/service/route53domains"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type Route53DomainsClient interface {
       AcceptDomainTransferFromAnotherAwsAccount(ctx workflow.Context, input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error)
       AcceptDomainTransferFromAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) *Route53domainsAcceptDomainTransferFromAnotherAwsAccountResult

       CancelDomainTransferToAnotherAwsAccount(ctx workflow.Context, input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error)
       CancelDomainTransferToAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) *Route53domainsCancelDomainTransferToAnotherAwsAccountResult

       CheckDomainAvailability(ctx workflow.Context, input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error)
       CheckDomainAvailabilityAsync(ctx workflow.Context, input *route53domains.CheckDomainAvailabilityInput) *Route53domainsCheckDomainAvailabilityResult

       CheckDomainTransferability(ctx workflow.Context, input *route53domains.CheckDomainTransferabilityInput) (*route53domains.CheckDomainTransferabilityOutput, error)
       CheckDomainTransferabilityAsync(ctx workflow.Context, input *route53domains.CheckDomainTransferabilityInput) *Route53domainsCheckDomainTransferabilityResult

       DeleteTagsForDomain(ctx workflow.Context, input *route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error)
       DeleteTagsForDomainAsync(ctx workflow.Context, input *route53domains.DeleteTagsForDomainInput) *Route53domainsDeleteTagsForDomainResult

       DisableDomainAutoRenew(ctx workflow.Context, input *route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error)
       DisableDomainAutoRenewAsync(ctx workflow.Context, input *route53domains.DisableDomainAutoRenewInput) *Route53domainsDisableDomainAutoRenewResult

       DisableDomainTransferLock(ctx workflow.Context, input *route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error)
       DisableDomainTransferLockAsync(ctx workflow.Context, input *route53domains.DisableDomainTransferLockInput) *Route53domainsDisableDomainTransferLockResult

       EnableDomainAutoRenew(ctx workflow.Context, input *route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error)
       EnableDomainAutoRenewAsync(ctx workflow.Context, input *route53domains.EnableDomainAutoRenewInput) *Route53domainsEnableDomainAutoRenewResult

       EnableDomainTransferLock(ctx workflow.Context, input *route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error)
       EnableDomainTransferLockAsync(ctx workflow.Context, input *route53domains.EnableDomainTransferLockInput) *Route53domainsEnableDomainTransferLockResult

       GetContactReachabilityStatus(ctx workflow.Context, input *route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error)
       GetContactReachabilityStatusAsync(ctx workflow.Context, input *route53domains.GetContactReachabilityStatusInput) *Route53domainsGetContactReachabilityStatusResult

       GetDomainDetail(ctx workflow.Context, input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error)
       GetDomainDetailAsync(ctx workflow.Context, input *route53domains.GetDomainDetailInput) *Route53domainsGetDomainDetailResult

       GetDomainSuggestions(ctx workflow.Context, input *route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error)
       GetDomainSuggestionsAsync(ctx workflow.Context, input *route53domains.GetDomainSuggestionsInput) *Route53domainsGetDomainSuggestionsResult

       GetOperationDetail(ctx workflow.Context, input *route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error)
       GetOperationDetailAsync(ctx workflow.Context, input *route53domains.GetOperationDetailInput) *Route53domainsGetOperationDetailResult

       ListDomains(ctx workflow.Context, input *route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error)
       ListDomainsAsync(ctx workflow.Context, input *route53domains.ListDomainsInput) *Route53domainsListDomainsResult

       ListOperations(ctx workflow.Context, input *route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error)
       ListOperationsAsync(ctx workflow.Context, input *route53domains.ListOperationsInput) *Route53domainsListOperationsResult

       ListTagsForDomain(ctx workflow.Context, input *route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error)
       ListTagsForDomainAsync(ctx workflow.Context, input *route53domains.ListTagsForDomainInput) *Route53domainsListTagsForDomainResult

       RegisterDomain(ctx workflow.Context, input *route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error)
       RegisterDomainAsync(ctx workflow.Context, input *route53domains.RegisterDomainInput) *Route53domainsRegisterDomainResult

       RejectDomainTransferFromAnotherAwsAccount(ctx workflow.Context, input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error)
       RejectDomainTransferFromAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) *Route53domainsRejectDomainTransferFromAnotherAwsAccountResult

       RenewDomain(ctx workflow.Context, input *route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error)
       RenewDomainAsync(ctx workflow.Context, input *route53domains.RenewDomainInput) *Route53domainsRenewDomainResult

       ResendContactReachabilityEmail(ctx workflow.Context, input *route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error)
       ResendContactReachabilityEmailAsync(ctx workflow.Context, input *route53domains.ResendContactReachabilityEmailInput) *Route53domainsResendContactReachabilityEmailResult

       RetrieveDomainAuthCode(ctx workflow.Context, input *route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error)
       RetrieveDomainAuthCodeAsync(ctx workflow.Context, input *route53domains.RetrieveDomainAuthCodeInput) *Route53domainsRetrieveDomainAuthCodeResult

       TransferDomain(ctx workflow.Context, input *route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error)
       TransferDomainAsync(ctx workflow.Context, input *route53domains.TransferDomainInput) *Route53domainsTransferDomainResult

       TransferDomainToAnotherAwsAccount(ctx workflow.Context, input *route53domains.TransferDomainToAnotherAwsAccountInput) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error)
       TransferDomainToAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.TransferDomainToAnotherAwsAccountInput) *Route53domainsTransferDomainToAnotherAwsAccountResult

       UpdateDomainContact(ctx workflow.Context, input *route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error)
       UpdateDomainContactAsync(ctx workflow.Context, input *route53domains.UpdateDomainContactInput) *Route53domainsUpdateDomainContactResult

       UpdateDomainContactPrivacy(ctx workflow.Context, input *route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error)
       UpdateDomainContactPrivacyAsync(ctx workflow.Context, input *route53domains.UpdateDomainContactPrivacyInput) *Route53domainsUpdateDomainContactPrivacyResult

       UpdateDomainNameservers(ctx workflow.Context, input *route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error)
       UpdateDomainNameserversAsync(ctx workflow.Context, input *route53domains.UpdateDomainNameserversInput) *Route53domainsUpdateDomainNameserversResult

       UpdateTagsForDomain(ctx workflow.Context, input *route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error)
       UpdateTagsForDomainAsync(ctx workflow.Context, input *route53domains.UpdateTagsForDomainInput) *Route53domainsUpdateTagsForDomainResult

       ViewBilling(ctx workflow.Context, input *route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error)
       ViewBillingAsync(ctx workflow.Context, input *route53domains.ViewBillingInput) *Route53domainsViewBillingResult
}

type Route53domainsAcceptDomainTransferFromAnotherAwsAccountResult struct {
	Result workflow.Future
}

func (r *Route53domainsAcceptDomainTransferFromAnotherAwsAccountResult) Get(ctx workflow.Context) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error) {
    var output route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsCancelDomainTransferToAnotherAwsAccountResult struct {
	Result workflow.Future
}

func (r *Route53domainsCancelDomainTransferToAnotherAwsAccountResult) Get(ctx workflow.Context) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error) {
    var output route53domains.CancelDomainTransferToAnotherAwsAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsCheckDomainAvailabilityResult struct {
	Result workflow.Future
}

func (r *Route53domainsCheckDomainAvailabilityResult) Get(ctx workflow.Context) (*route53domains.CheckDomainAvailabilityOutput, error) {
    var output route53domains.CheckDomainAvailabilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsCheckDomainTransferabilityResult struct {
	Result workflow.Future
}

func (r *Route53domainsCheckDomainTransferabilityResult) Get(ctx workflow.Context) (*route53domains.CheckDomainTransferabilityOutput, error) {
    var output route53domains.CheckDomainTransferabilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsDeleteTagsForDomainResult struct {
	Result workflow.Future
}

func (r *Route53domainsDeleteTagsForDomainResult) Get(ctx workflow.Context) (*route53domains.DeleteTagsForDomainOutput, error) {
    var output route53domains.DeleteTagsForDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsDisableDomainAutoRenewResult struct {
	Result workflow.Future
}

func (r *Route53domainsDisableDomainAutoRenewResult) Get(ctx workflow.Context) (*route53domains.DisableDomainAutoRenewOutput, error) {
    var output route53domains.DisableDomainAutoRenewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsDisableDomainTransferLockResult struct {
	Result workflow.Future
}

func (r *Route53domainsDisableDomainTransferLockResult) Get(ctx workflow.Context) (*route53domains.DisableDomainTransferLockOutput, error) {
    var output route53domains.DisableDomainTransferLockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsEnableDomainAutoRenewResult struct {
	Result workflow.Future
}

func (r *Route53domainsEnableDomainAutoRenewResult) Get(ctx workflow.Context) (*route53domains.EnableDomainAutoRenewOutput, error) {
    var output route53domains.EnableDomainAutoRenewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsEnableDomainTransferLockResult struct {
	Result workflow.Future
}

func (r *Route53domainsEnableDomainTransferLockResult) Get(ctx workflow.Context) (*route53domains.EnableDomainTransferLockOutput, error) {
    var output route53domains.EnableDomainTransferLockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsGetContactReachabilityStatusResult struct {
	Result workflow.Future
}

func (r *Route53domainsGetContactReachabilityStatusResult) Get(ctx workflow.Context) (*route53domains.GetContactReachabilityStatusOutput, error) {
    var output route53domains.GetContactReachabilityStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsGetDomainDetailResult struct {
	Result workflow.Future
}

func (r *Route53domainsGetDomainDetailResult) Get(ctx workflow.Context) (*route53domains.GetDomainDetailOutput, error) {
    var output route53domains.GetDomainDetailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsGetDomainSuggestionsResult struct {
	Result workflow.Future
}

func (r *Route53domainsGetDomainSuggestionsResult) Get(ctx workflow.Context) (*route53domains.GetDomainSuggestionsOutput, error) {
    var output route53domains.GetDomainSuggestionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsGetOperationDetailResult struct {
	Result workflow.Future
}

func (r *Route53domainsGetOperationDetailResult) Get(ctx workflow.Context) (*route53domains.GetOperationDetailOutput, error) {
    var output route53domains.GetOperationDetailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsListDomainsResult struct {
	Result workflow.Future
}

func (r *Route53domainsListDomainsResult) Get(ctx workflow.Context) (*route53domains.ListDomainsOutput, error) {
    var output route53domains.ListDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsListOperationsResult struct {
	Result workflow.Future
}

func (r *Route53domainsListOperationsResult) Get(ctx workflow.Context) (*route53domains.ListOperationsOutput, error) {
    var output route53domains.ListOperationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsListTagsForDomainResult struct {
	Result workflow.Future
}

func (r *Route53domainsListTagsForDomainResult) Get(ctx workflow.Context) (*route53domains.ListTagsForDomainOutput, error) {
    var output route53domains.ListTagsForDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsRegisterDomainResult struct {
	Result workflow.Future
}

func (r *Route53domainsRegisterDomainResult) Get(ctx workflow.Context) (*route53domains.RegisterDomainOutput, error) {
    var output route53domains.RegisterDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsRejectDomainTransferFromAnotherAwsAccountResult struct {
	Result workflow.Future
}

func (r *Route53domainsRejectDomainTransferFromAnotherAwsAccountResult) Get(ctx workflow.Context) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error) {
    var output route53domains.RejectDomainTransferFromAnotherAwsAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsRenewDomainResult struct {
	Result workflow.Future
}

func (r *Route53domainsRenewDomainResult) Get(ctx workflow.Context) (*route53domains.RenewDomainOutput, error) {
    var output route53domains.RenewDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsResendContactReachabilityEmailResult struct {
	Result workflow.Future
}

func (r *Route53domainsResendContactReachabilityEmailResult) Get(ctx workflow.Context) (*route53domains.ResendContactReachabilityEmailOutput, error) {
    var output route53domains.ResendContactReachabilityEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsRetrieveDomainAuthCodeResult struct {
	Result workflow.Future
}

func (r *Route53domainsRetrieveDomainAuthCodeResult) Get(ctx workflow.Context) (*route53domains.RetrieveDomainAuthCodeOutput, error) {
    var output route53domains.RetrieveDomainAuthCodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsTransferDomainResult struct {
	Result workflow.Future
}

func (r *Route53domainsTransferDomainResult) Get(ctx workflow.Context) (*route53domains.TransferDomainOutput, error) {
    var output route53domains.TransferDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsTransferDomainToAnotherAwsAccountResult struct {
	Result workflow.Future
}

func (r *Route53domainsTransferDomainToAnotherAwsAccountResult) Get(ctx workflow.Context) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error) {
    var output route53domains.TransferDomainToAnotherAwsAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsUpdateDomainContactResult struct {
	Result workflow.Future
}

func (r *Route53domainsUpdateDomainContactResult) Get(ctx workflow.Context) (*route53domains.UpdateDomainContactOutput, error) {
    var output route53domains.UpdateDomainContactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsUpdateDomainContactPrivacyResult struct {
	Result workflow.Future
}

func (r *Route53domainsUpdateDomainContactPrivacyResult) Get(ctx workflow.Context) (*route53domains.UpdateDomainContactPrivacyOutput, error) {
    var output route53domains.UpdateDomainContactPrivacyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsUpdateDomainNameserversResult struct {
	Result workflow.Future
}

func (r *Route53domainsUpdateDomainNameserversResult) Get(ctx workflow.Context) (*route53domains.UpdateDomainNameserversOutput, error) {
    var output route53domains.UpdateDomainNameserversOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsUpdateTagsForDomainResult struct {
	Result workflow.Future
}

func (r *Route53domainsUpdateTagsForDomainResult) Get(ctx workflow.Context) (*route53domains.UpdateTagsForDomainOutput, error) {
    var output route53domains.UpdateTagsForDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53domainsViewBillingResult struct {
	Result workflow.Future
}

func (r *Route53domainsViewBillingResult) Get(ctx workflow.Context) (*route53domains.ViewBillingOutput, error) {
    var output route53domains.ViewBillingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DomainsStub struct {
    activities awsactivities.Route53DomainsActivities
}

func NewRoute53DomainsStub() Route53DomainsClient {
    return &Route53DomainsStub{}
}

func (a *Route53DomainsStub) AcceptDomainTransferFromAnotherAwsAccount(ctx workflow.Context, input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error) {
    var output route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptDomainTransferFromAnotherAwsAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) AcceptDomainTransferFromAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) *Route53domainsAcceptDomainTransferFromAnotherAwsAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptDomainTransferFromAnotherAwsAccount, input)
    return &Route53domainsAcceptDomainTransferFromAnotherAwsAccountResult{Result: future}
}

func (a *Route53DomainsStub) CancelDomainTransferToAnotherAwsAccount(ctx workflow.Context, input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error) {
    var output route53domains.CancelDomainTransferToAnotherAwsAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelDomainTransferToAnotherAwsAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) CancelDomainTransferToAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) *Route53domainsCancelDomainTransferToAnotherAwsAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelDomainTransferToAnotherAwsAccount, input)
    return &Route53domainsCancelDomainTransferToAnotherAwsAccountResult{Result: future}
}

func (a *Route53DomainsStub) CheckDomainAvailability(ctx workflow.Context, input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error) {
    var output route53domains.CheckDomainAvailabilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CheckDomainAvailability, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) CheckDomainAvailabilityAsync(ctx workflow.Context, input *route53domains.CheckDomainAvailabilityInput) *Route53domainsCheckDomainAvailabilityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CheckDomainAvailability, input)
    return &Route53domainsCheckDomainAvailabilityResult{Result: future}
}

func (a *Route53DomainsStub) CheckDomainTransferability(ctx workflow.Context, input *route53domains.CheckDomainTransferabilityInput) (*route53domains.CheckDomainTransferabilityOutput, error) {
    var output route53domains.CheckDomainTransferabilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CheckDomainTransferability, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) CheckDomainTransferabilityAsync(ctx workflow.Context, input *route53domains.CheckDomainTransferabilityInput) *Route53domainsCheckDomainTransferabilityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CheckDomainTransferability, input)
    return &Route53domainsCheckDomainTransferabilityResult{Result: future}
}

func (a *Route53DomainsStub) DeleteTagsForDomain(ctx workflow.Context, input *route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error) {
    var output route53domains.DeleteTagsForDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTagsForDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) DeleteTagsForDomainAsync(ctx workflow.Context, input *route53domains.DeleteTagsForDomainInput) *Route53domainsDeleteTagsForDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTagsForDomain, input)
    return &Route53domainsDeleteTagsForDomainResult{Result: future}
}

func (a *Route53DomainsStub) DisableDomainAutoRenew(ctx workflow.Context, input *route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error) {
    var output route53domains.DisableDomainAutoRenewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableDomainAutoRenew, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) DisableDomainAutoRenewAsync(ctx workflow.Context, input *route53domains.DisableDomainAutoRenewInput) *Route53domainsDisableDomainAutoRenewResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableDomainAutoRenew, input)
    return &Route53domainsDisableDomainAutoRenewResult{Result: future}
}

func (a *Route53DomainsStub) DisableDomainTransferLock(ctx workflow.Context, input *route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error) {
    var output route53domains.DisableDomainTransferLockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableDomainTransferLock, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) DisableDomainTransferLockAsync(ctx workflow.Context, input *route53domains.DisableDomainTransferLockInput) *Route53domainsDisableDomainTransferLockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableDomainTransferLock, input)
    return &Route53domainsDisableDomainTransferLockResult{Result: future}
}

func (a *Route53DomainsStub) EnableDomainAutoRenew(ctx workflow.Context, input *route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error) {
    var output route53domains.EnableDomainAutoRenewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableDomainAutoRenew, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) EnableDomainAutoRenewAsync(ctx workflow.Context, input *route53domains.EnableDomainAutoRenewInput) *Route53domainsEnableDomainAutoRenewResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableDomainAutoRenew, input)
    return &Route53domainsEnableDomainAutoRenewResult{Result: future}
}

func (a *Route53DomainsStub) EnableDomainTransferLock(ctx workflow.Context, input *route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error) {
    var output route53domains.EnableDomainTransferLockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableDomainTransferLock, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) EnableDomainTransferLockAsync(ctx workflow.Context, input *route53domains.EnableDomainTransferLockInput) *Route53domainsEnableDomainTransferLockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableDomainTransferLock, input)
    return &Route53domainsEnableDomainTransferLockResult{Result: future}
}

func (a *Route53DomainsStub) GetContactReachabilityStatus(ctx workflow.Context, input *route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error) {
    var output route53domains.GetContactReachabilityStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetContactReachabilityStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) GetContactReachabilityStatusAsync(ctx workflow.Context, input *route53domains.GetContactReachabilityStatusInput) *Route53domainsGetContactReachabilityStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetContactReachabilityStatus, input)
    return &Route53domainsGetContactReachabilityStatusResult{Result: future}
}

func (a *Route53DomainsStub) GetDomainDetail(ctx workflow.Context, input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error) {
    var output route53domains.GetDomainDetailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainDetail, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) GetDomainDetailAsync(ctx workflow.Context, input *route53domains.GetDomainDetailInput) *Route53domainsGetDomainDetailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDomainDetail, input)
    return &Route53domainsGetDomainDetailResult{Result: future}
}

func (a *Route53DomainsStub) GetDomainSuggestions(ctx workflow.Context, input *route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error) {
    var output route53domains.GetDomainSuggestionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainSuggestions, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) GetDomainSuggestionsAsync(ctx workflow.Context, input *route53domains.GetDomainSuggestionsInput) *Route53domainsGetDomainSuggestionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDomainSuggestions, input)
    return &Route53domainsGetDomainSuggestionsResult{Result: future}
}

func (a *Route53DomainsStub) GetOperationDetail(ctx workflow.Context, input *route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error) {
    var output route53domains.GetOperationDetailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOperationDetail, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) GetOperationDetailAsync(ctx workflow.Context, input *route53domains.GetOperationDetailInput) *Route53domainsGetOperationDetailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetOperationDetail, input)
    return &Route53domainsGetOperationDetailResult{Result: future}
}

func (a *Route53DomainsStub) ListDomains(ctx workflow.Context, input *route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error) {
    var output route53domains.ListDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) ListDomainsAsync(ctx workflow.Context, input *route53domains.ListDomainsInput) *Route53domainsListDomainsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDomains, input)
    return &Route53domainsListDomainsResult{Result: future}
}

func (a *Route53DomainsStub) ListOperations(ctx workflow.Context, input *route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error) {
    var output route53domains.ListOperationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOperations, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) ListOperationsAsync(ctx workflow.Context, input *route53domains.ListOperationsInput) *Route53domainsListOperationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOperations, input)
    return &Route53domainsListOperationsResult{Result: future}
}

func (a *Route53DomainsStub) ListTagsForDomain(ctx workflow.Context, input *route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error) {
    var output route53domains.ListTagsForDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) ListTagsForDomainAsync(ctx workflow.Context, input *route53domains.ListTagsForDomainInput) *Route53domainsListTagsForDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForDomain, input)
    return &Route53domainsListTagsForDomainResult{Result: future}
}

func (a *Route53DomainsStub) RegisterDomain(ctx workflow.Context, input *route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error) {
    var output route53domains.RegisterDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) RegisterDomainAsync(ctx workflow.Context, input *route53domains.RegisterDomainInput) *Route53domainsRegisterDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterDomain, input)
    return &Route53domainsRegisterDomainResult{Result: future}
}

func (a *Route53DomainsStub) RejectDomainTransferFromAnotherAwsAccount(ctx workflow.Context, input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error) {
    var output route53domains.RejectDomainTransferFromAnotherAwsAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectDomainTransferFromAnotherAwsAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) RejectDomainTransferFromAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) *Route53domainsRejectDomainTransferFromAnotherAwsAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectDomainTransferFromAnotherAwsAccount, input)
    return &Route53domainsRejectDomainTransferFromAnotherAwsAccountResult{Result: future}
}

func (a *Route53DomainsStub) RenewDomain(ctx workflow.Context, input *route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error) {
    var output route53domains.RenewDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RenewDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) RenewDomainAsync(ctx workflow.Context, input *route53domains.RenewDomainInput) *Route53domainsRenewDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RenewDomain, input)
    return &Route53domainsRenewDomainResult{Result: future}
}

func (a *Route53DomainsStub) ResendContactReachabilityEmail(ctx workflow.Context, input *route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error) {
    var output route53domains.ResendContactReachabilityEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResendContactReachabilityEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) ResendContactReachabilityEmailAsync(ctx workflow.Context, input *route53domains.ResendContactReachabilityEmailInput) *Route53domainsResendContactReachabilityEmailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResendContactReachabilityEmail, input)
    return &Route53domainsResendContactReachabilityEmailResult{Result: future}
}

func (a *Route53DomainsStub) RetrieveDomainAuthCode(ctx workflow.Context, input *route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error) {
    var output route53domains.RetrieveDomainAuthCodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RetrieveDomainAuthCode, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) RetrieveDomainAuthCodeAsync(ctx workflow.Context, input *route53domains.RetrieveDomainAuthCodeInput) *Route53domainsRetrieveDomainAuthCodeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RetrieveDomainAuthCode, input)
    return &Route53domainsRetrieveDomainAuthCodeResult{Result: future}
}

func (a *Route53DomainsStub) TransferDomain(ctx workflow.Context, input *route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error) {
    var output route53domains.TransferDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TransferDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) TransferDomainAsync(ctx workflow.Context, input *route53domains.TransferDomainInput) *Route53domainsTransferDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TransferDomain, input)
    return &Route53domainsTransferDomainResult{Result: future}
}

func (a *Route53DomainsStub) TransferDomainToAnotherAwsAccount(ctx workflow.Context, input *route53domains.TransferDomainToAnotherAwsAccountInput) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error) {
    var output route53domains.TransferDomainToAnotherAwsAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TransferDomainToAnotherAwsAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) TransferDomainToAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.TransferDomainToAnotherAwsAccountInput) *Route53domainsTransferDomainToAnotherAwsAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TransferDomainToAnotherAwsAccount, input)
    return &Route53domainsTransferDomainToAnotherAwsAccountResult{Result: future}
}

func (a *Route53DomainsStub) UpdateDomainContact(ctx workflow.Context, input *route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error) {
    var output route53domains.UpdateDomainContactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainContact, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) UpdateDomainContactAsync(ctx workflow.Context, input *route53domains.UpdateDomainContactInput) *Route53domainsUpdateDomainContactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainContact, input)
    return &Route53domainsUpdateDomainContactResult{Result: future}
}

func (a *Route53DomainsStub) UpdateDomainContactPrivacy(ctx workflow.Context, input *route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error) {
    var output route53domains.UpdateDomainContactPrivacyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainContactPrivacy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) UpdateDomainContactPrivacyAsync(ctx workflow.Context, input *route53domains.UpdateDomainContactPrivacyInput) *Route53domainsUpdateDomainContactPrivacyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainContactPrivacy, input)
    return &Route53domainsUpdateDomainContactPrivacyResult{Result: future}
}

func (a *Route53DomainsStub) UpdateDomainNameservers(ctx workflow.Context, input *route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error) {
    var output route53domains.UpdateDomainNameserversOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainNameservers, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) UpdateDomainNameserversAsync(ctx workflow.Context, input *route53domains.UpdateDomainNameserversInput) *Route53domainsUpdateDomainNameserversResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainNameservers, input)
    return &Route53domainsUpdateDomainNameserversResult{Result: future}
}

func (a *Route53DomainsStub) UpdateTagsForDomain(ctx workflow.Context, input *route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error) {
    var output route53domains.UpdateTagsForDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTagsForDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) UpdateTagsForDomainAsync(ctx workflow.Context, input *route53domains.UpdateTagsForDomainInput) *Route53domainsUpdateTagsForDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTagsForDomain, input)
    return &Route53domainsUpdateTagsForDomainResult{Result: future}
}

func (a *Route53DomainsStub) ViewBilling(ctx workflow.Context, input *route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error) {
    var output route53domains.ViewBillingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ViewBilling, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53DomainsStub) ViewBillingAsync(ctx workflow.Context, input *route53domains.ViewBillingInput) *Route53domainsViewBillingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ViewBilling, input)
    return &Route53domainsViewBillingResult{Result: future}
}
