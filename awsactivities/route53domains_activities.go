
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type Route53DomainsActivities struct {
    client route53domainsiface.Route53DomainsAPI
}

func NewRoute53DomainsActivities(session *session.Session, config ...*aws.Config) *Route53DomainsActivities {
    client := route53domains.New(session, config...)
    return &Route53DomainsActivities{client: client}
}

func (a *Route53DomainsActivities) AcceptDomainTransferFromAnotherAwsAccount(ctx context.Context, input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error) {
    return a.client.AcceptDomainTransferFromAnotherAwsAccountWithContext(ctx, input)
}

func (a *Route53DomainsActivities) CancelDomainTransferToAnotherAwsAccount(ctx context.Context, input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error) {
    return a.client.CancelDomainTransferToAnotherAwsAccountWithContext(ctx, input)
}

func (a *Route53DomainsActivities) CheckDomainAvailability(ctx context.Context, input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error) {
    return a.client.CheckDomainAvailabilityWithContext(ctx, input)
}

func (a *Route53DomainsActivities) CheckDomainTransferability(ctx context.Context, input *route53domains.CheckDomainTransferabilityInput) (*route53domains.CheckDomainTransferabilityOutput, error) {
    return a.client.CheckDomainTransferabilityWithContext(ctx, input)
}

func (a *Route53DomainsActivities) DeleteTagsForDomain(ctx context.Context, input *route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error) {
    return a.client.DeleteTagsForDomainWithContext(ctx, input)
}

func (a *Route53DomainsActivities) DisableDomainAutoRenew(ctx context.Context, input *route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error) {
    return a.client.DisableDomainAutoRenewWithContext(ctx, input)
}

func (a *Route53DomainsActivities) DisableDomainTransferLock(ctx context.Context, input *route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error) {
    return a.client.DisableDomainTransferLockWithContext(ctx, input)
}

func (a *Route53DomainsActivities) EnableDomainAutoRenew(ctx context.Context, input *route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error) {
    return a.client.EnableDomainAutoRenewWithContext(ctx, input)
}

func (a *Route53DomainsActivities) EnableDomainTransferLock(ctx context.Context, input *route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error) {
    return a.client.EnableDomainTransferLockWithContext(ctx, input)
}

func (a *Route53DomainsActivities) GetContactReachabilityStatus(ctx context.Context, input *route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error) {
    return a.client.GetContactReachabilityStatusWithContext(ctx, input)
}

func (a *Route53DomainsActivities) GetDomainDetail(ctx context.Context, input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error) {
    return a.client.GetDomainDetailWithContext(ctx, input)
}

func (a *Route53DomainsActivities) GetDomainSuggestions(ctx context.Context, input *route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error) {
    return a.client.GetDomainSuggestionsWithContext(ctx, input)
}

func (a *Route53DomainsActivities) GetOperationDetail(ctx context.Context, input *route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error) {
    return a.client.GetOperationDetailWithContext(ctx, input)
}

func (a *Route53DomainsActivities) ListDomains(ctx context.Context, input *route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error) {
    return a.client.ListDomainsWithContext(ctx, input)
}

func (a *Route53DomainsActivities) ListOperations(ctx context.Context, input *route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error) {
    return a.client.ListOperationsWithContext(ctx, input)
}

func (a *Route53DomainsActivities) ListTagsForDomain(ctx context.Context, input *route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error) {
    return a.client.ListTagsForDomainWithContext(ctx, input)
}

func (a *Route53DomainsActivities) RegisterDomain(ctx context.Context, input *route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error) {
    return a.client.RegisterDomainWithContext(ctx, input)
}

func (a *Route53DomainsActivities) RejectDomainTransferFromAnotherAwsAccount(ctx context.Context, input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error) {
    return a.client.RejectDomainTransferFromAnotherAwsAccountWithContext(ctx, input)
}

func (a *Route53DomainsActivities) RenewDomain(ctx context.Context, input *route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error) {
    return a.client.RenewDomainWithContext(ctx, input)
}

func (a *Route53DomainsActivities) ResendContactReachabilityEmail(ctx context.Context, input *route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error) {
    return a.client.ResendContactReachabilityEmailWithContext(ctx, input)
}

func (a *Route53DomainsActivities) RetrieveDomainAuthCode(ctx context.Context, input *route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error) {
    return a.client.RetrieveDomainAuthCodeWithContext(ctx, input)
}

func (a *Route53DomainsActivities) TransferDomain(ctx context.Context, input *route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error) {
    return a.client.TransferDomainWithContext(ctx, input)
}

func (a *Route53DomainsActivities) TransferDomainToAnotherAwsAccount(ctx context.Context, input *route53domains.TransferDomainToAnotherAwsAccountInput) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error) {
    return a.client.TransferDomainToAnotherAwsAccountWithContext(ctx, input)
}

func (a *Route53DomainsActivities) UpdateDomainContact(ctx context.Context, input *route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error) {
    return a.client.UpdateDomainContactWithContext(ctx, input)
}

func (a *Route53DomainsActivities) UpdateDomainContactPrivacy(ctx context.Context, input *route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error) {
    return a.client.UpdateDomainContactPrivacyWithContext(ctx, input)
}

func (a *Route53DomainsActivities) UpdateDomainNameservers(ctx context.Context, input *route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error) {
    return a.client.UpdateDomainNameserversWithContext(ctx, input)
}

func (a *Route53DomainsActivities) UpdateTagsForDomain(ctx context.Context, input *route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error) {
    return a.client.UpdateTagsForDomainWithContext(ctx, input)
}

func (a *Route53DomainsActivities) ViewBilling(ctx context.Context, input *route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error) {
    return a.client.ViewBillingWithContext(ctx, input)
}
