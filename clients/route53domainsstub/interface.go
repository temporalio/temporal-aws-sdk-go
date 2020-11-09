// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package route53domainsstub

import (
	"github.com/aws/aws-sdk-go/service/route53domains"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptDomainTransferFromAnotherAwsAccount(ctx workflow.Context, input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error)
	AcceptDomainTransferFromAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) *AcceptDomainTransferFromAnotherAwsAccountFuture

	CancelDomainTransferToAnotherAwsAccount(ctx workflow.Context, input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error)
	CancelDomainTransferToAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) *CancelDomainTransferToAnotherAwsAccountFuture

	CheckDomainAvailability(ctx workflow.Context, input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error)
	CheckDomainAvailabilityAsync(ctx workflow.Context, input *route53domains.CheckDomainAvailabilityInput) *CheckDomainAvailabilityFuture

	CheckDomainTransferability(ctx workflow.Context, input *route53domains.CheckDomainTransferabilityInput) (*route53domains.CheckDomainTransferabilityOutput, error)
	CheckDomainTransferabilityAsync(ctx workflow.Context, input *route53domains.CheckDomainTransferabilityInput) *CheckDomainTransferabilityFuture

	DeleteTagsForDomain(ctx workflow.Context, input *route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error)
	DeleteTagsForDomainAsync(ctx workflow.Context, input *route53domains.DeleteTagsForDomainInput) *DeleteTagsForDomainFuture

	DisableDomainAutoRenew(ctx workflow.Context, input *route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error)
	DisableDomainAutoRenewAsync(ctx workflow.Context, input *route53domains.DisableDomainAutoRenewInput) *DisableDomainAutoRenewFuture

	DisableDomainTransferLock(ctx workflow.Context, input *route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error)
	DisableDomainTransferLockAsync(ctx workflow.Context, input *route53domains.DisableDomainTransferLockInput) *DisableDomainTransferLockFuture

	EnableDomainAutoRenew(ctx workflow.Context, input *route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error)
	EnableDomainAutoRenewAsync(ctx workflow.Context, input *route53domains.EnableDomainAutoRenewInput) *EnableDomainAutoRenewFuture

	EnableDomainTransferLock(ctx workflow.Context, input *route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error)
	EnableDomainTransferLockAsync(ctx workflow.Context, input *route53domains.EnableDomainTransferLockInput) *EnableDomainTransferLockFuture

	GetContactReachabilityStatus(ctx workflow.Context, input *route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error)
	GetContactReachabilityStatusAsync(ctx workflow.Context, input *route53domains.GetContactReachabilityStatusInput) *GetContactReachabilityStatusFuture

	GetDomainDetail(ctx workflow.Context, input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error)
	GetDomainDetailAsync(ctx workflow.Context, input *route53domains.GetDomainDetailInput) *GetDomainDetailFuture

	GetDomainSuggestions(ctx workflow.Context, input *route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error)
	GetDomainSuggestionsAsync(ctx workflow.Context, input *route53domains.GetDomainSuggestionsInput) *GetDomainSuggestionsFuture

	GetOperationDetail(ctx workflow.Context, input *route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error)
	GetOperationDetailAsync(ctx workflow.Context, input *route53domains.GetOperationDetailInput) *GetOperationDetailFuture

	ListDomains(ctx workflow.Context, input *route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error)
	ListDomainsAsync(ctx workflow.Context, input *route53domains.ListDomainsInput) *ListDomainsFuture

	ListOperations(ctx workflow.Context, input *route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error)
	ListOperationsAsync(ctx workflow.Context, input *route53domains.ListOperationsInput) *ListOperationsFuture

	ListTagsForDomain(ctx workflow.Context, input *route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error)
	ListTagsForDomainAsync(ctx workflow.Context, input *route53domains.ListTagsForDomainInput) *ListTagsForDomainFuture

	RegisterDomain(ctx workflow.Context, input *route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error)
	RegisterDomainAsync(ctx workflow.Context, input *route53domains.RegisterDomainInput) *RegisterDomainFuture

	RejectDomainTransferFromAnotherAwsAccount(ctx workflow.Context, input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error)
	RejectDomainTransferFromAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) *RejectDomainTransferFromAnotherAwsAccountFuture

	RenewDomain(ctx workflow.Context, input *route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error)
	RenewDomainAsync(ctx workflow.Context, input *route53domains.RenewDomainInput) *RenewDomainFuture

	ResendContactReachabilityEmail(ctx workflow.Context, input *route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error)
	ResendContactReachabilityEmailAsync(ctx workflow.Context, input *route53domains.ResendContactReachabilityEmailInput) *ResendContactReachabilityEmailFuture

	RetrieveDomainAuthCode(ctx workflow.Context, input *route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error)
	RetrieveDomainAuthCodeAsync(ctx workflow.Context, input *route53domains.RetrieveDomainAuthCodeInput) *RetrieveDomainAuthCodeFuture

	TransferDomain(ctx workflow.Context, input *route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error)
	TransferDomainAsync(ctx workflow.Context, input *route53domains.TransferDomainInput) *TransferDomainFuture

	TransferDomainToAnotherAwsAccount(ctx workflow.Context, input *route53domains.TransferDomainToAnotherAwsAccountInput) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error)
	TransferDomainToAnotherAwsAccountAsync(ctx workflow.Context, input *route53domains.TransferDomainToAnotherAwsAccountInput) *TransferDomainToAnotherAwsAccountFuture

	UpdateDomainContact(ctx workflow.Context, input *route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error)
	UpdateDomainContactAsync(ctx workflow.Context, input *route53domains.UpdateDomainContactInput) *UpdateDomainContactFuture

	UpdateDomainContactPrivacy(ctx workflow.Context, input *route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error)
	UpdateDomainContactPrivacyAsync(ctx workflow.Context, input *route53domains.UpdateDomainContactPrivacyInput) *UpdateDomainContactPrivacyFuture

	UpdateDomainNameservers(ctx workflow.Context, input *route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error)
	UpdateDomainNameserversAsync(ctx workflow.Context, input *route53domains.UpdateDomainNameserversInput) *UpdateDomainNameserversFuture

	UpdateTagsForDomain(ctx workflow.Context, input *route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error)
	UpdateTagsForDomainAsync(ctx workflow.Context, input *route53domains.UpdateTagsForDomainInput) *UpdateTagsForDomainFuture

	ViewBilling(ctx workflow.Context, input *route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error)
	ViewBillingAsync(ctx workflow.Context, input *route53domains.ViewBillingInput) *ViewBillingFuture
}

func NewClient() Client {
	return &stub{}
}
