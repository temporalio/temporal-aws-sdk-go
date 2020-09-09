
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
)

type Route53DomainsActivities struct {
	client route53domainsiface.Route53DomainsAPI
}

func NewRoute53DomainsActivities(session *session.Session, config... *aws.Config) *Route53DomainsActivities {
    client := route53domains.New(session, config...)
    return &Route53DomainsActivities{client: client}
}

func (a *Route53DomainsActivities) AcceptDomainTransferFromAnotherAwsAccount(input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error) {
    return a.client.AcceptDomainTransferFromAnotherAwsAccount(input)
}

func (a *Route53DomainsActivities) CancelDomainTransferToAnotherAwsAccount(input *route53domains.CancelDomainTransferToAnotherAwsAccountInput) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error) {
    return a.client.CancelDomainTransferToAnotherAwsAccount(input)
}

func (a *Route53DomainsActivities) CheckDomainAvailability(input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error) {
    return a.client.CheckDomainAvailability(input)
}

func (a *Route53DomainsActivities) CheckDomainTransferability(input *route53domains.CheckDomainTransferabilityInput) (*route53domains.CheckDomainTransferabilityOutput, error) {
    return a.client.CheckDomainTransferability(input)
}

func (a *Route53DomainsActivities) DeleteTagsForDomain(input *route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error) {
    return a.client.DeleteTagsForDomain(input)
}

func (a *Route53DomainsActivities) DisableDomainAutoRenew(input *route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error) {
    return a.client.DisableDomainAutoRenew(input)
}

func (a *Route53DomainsActivities) DisableDomainTransferLock(input *route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error) {
    return a.client.DisableDomainTransferLock(input)
}

func (a *Route53DomainsActivities) EnableDomainAutoRenew(input *route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error) {
    return a.client.EnableDomainAutoRenew(input)
}

func (a *Route53DomainsActivities) EnableDomainTransferLock(input *route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error) {
    return a.client.EnableDomainTransferLock(input)
}

func (a *Route53DomainsActivities) GetContactReachabilityStatus(input *route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error) {
    return a.client.GetContactReachabilityStatus(input)
}

func (a *Route53DomainsActivities) GetDomainDetail(input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error) {
    return a.client.GetDomainDetail(input)
}

func (a *Route53DomainsActivities) GetDomainSuggestions(input *route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error) {
    return a.client.GetDomainSuggestions(input)
}

func (a *Route53DomainsActivities) GetOperationDetail(input *route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error) {
    return a.client.GetOperationDetail(input)
}

func (a *Route53DomainsActivities) ListDomains(input *route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error) {
    return a.client.ListDomains(input)
}

func (a *Route53DomainsActivities) ListOperations(input *route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error) {
    return a.client.ListOperations(input)
}

func (a *Route53DomainsActivities) ListTagsForDomain(input *route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error) {
    return a.client.ListTagsForDomain(input)
}

func (a *Route53DomainsActivities) RegisterDomain(input *route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error) {
    return a.client.RegisterDomain(input)
}

func (a *Route53DomainsActivities) RejectDomainTransferFromAnotherAwsAccount(input *route53domains.RejectDomainTransferFromAnotherAwsAccountInput) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error) {
    return a.client.RejectDomainTransferFromAnotherAwsAccount(input)
}

func (a *Route53DomainsActivities) RenewDomain(input *route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error) {
    return a.client.RenewDomain(input)
}

func (a *Route53DomainsActivities) ResendContactReachabilityEmail(input *route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error) {
    return a.client.ResendContactReachabilityEmail(input)
}

func (a *Route53DomainsActivities) RetrieveDomainAuthCode(input *route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error) {
    return a.client.RetrieveDomainAuthCode(input)
}

func (a *Route53DomainsActivities) TransferDomain(input *route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error) {
    return a.client.TransferDomain(input)
}

func (a *Route53DomainsActivities) TransferDomainToAnotherAwsAccount(input *route53domains.TransferDomainToAnotherAwsAccountInput) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error) {
    return a.client.TransferDomainToAnotherAwsAccount(input)
}

func (a *Route53DomainsActivities) UpdateDomainContact(input *route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error) {
    return a.client.UpdateDomainContact(input)
}

func (a *Route53DomainsActivities) UpdateDomainContactPrivacy(input *route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error) {
    return a.client.UpdateDomainContactPrivacy(input)
}

func (a *Route53DomainsActivities) UpdateDomainNameservers(input *route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error) {
    return a.client.UpdateDomainNameservers(input)
}

func (a *Route53DomainsActivities) UpdateTagsForDomain(input *route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error) {
    return a.client.UpdateTagsForDomain(input)
}

func (a *Route53DomainsActivities) ViewBilling(input *route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error) {
    return a.client.ViewBilling(input)
}
