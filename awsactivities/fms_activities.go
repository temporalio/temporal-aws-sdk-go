
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/fms/fmsiface"
)

type FMSActivities struct {
	client fmsiface.FMSAPI
}

func NewFMSActivities(client fmsiface.FMSAPI) *FMSActivities {
    return &FMSActivities{client: client}
}

func (a *FMSActivities) AssociateAdminAccount(input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
    return a.client.AssociateAdminAccount(input)
}

func (a *FMSActivities) DeleteAppsList(input *fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error) {
    return a.client.DeleteAppsList(input)
}

func (a *FMSActivities) DeleteNotificationChannel(input *fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error) {
    return a.client.DeleteNotificationChannel(input)
}

func (a *FMSActivities) DeletePolicy(input *fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error) {
    return a.client.DeletePolicy(input)
}

func (a *FMSActivities) DeleteProtocolsList(input *fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error) {
    return a.client.DeleteProtocolsList(input)
}

func (a *FMSActivities) DisassociateAdminAccount(input *fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error) {
    return a.client.DisassociateAdminAccount(input)
}

func (a *FMSActivities) GetAdminAccount(input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error) {
    return a.client.GetAdminAccount(input)
}

func (a *FMSActivities) GetAppsList(input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error) {
    return a.client.GetAppsList(input)
}

func (a *FMSActivities) GetComplianceDetail(input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error) {
    return a.client.GetComplianceDetail(input)
}

func (a *FMSActivities) GetNotificationChannel(input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error) {
    return a.client.GetNotificationChannel(input)
}

func (a *FMSActivities) GetPolicy(input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error) {
    return a.client.GetPolicy(input)
}

func (a *FMSActivities) GetProtectionStatus(input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error) {
    return a.client.GetProtectionStatus(input)
}

func (a *FMSActivities) GetProtocolsList(input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error) {
    return a.client.GetProtocolsList(input)
}

func (a *FMSActivities) GetViolationDetails(input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error) {
    return a.client.GetViolationDetails(input)
}

func (a *FMSActivities) ListAppsLists(input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error) {
    return a.client.ListAppsLists(input)
}

func (a *FMSActivities) ListComplianceStatus(input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error) {
    return a.client.ListComplianceStatus(input)
}

func (a *FMSActivities) ListMemberAccounts(input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error) {
    return a.client.ListMemberAccounts(input)
}

func (a *FMSActivities) ListPolicies(input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error) {
    return a.client.ListPolicies(input)
}

func (a *FMSActivities) ListProtocolsLists(input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error) {
    return a.client.ListProtocolsLists(input)
}

func (a *FMSActivities) ListTagsForResource(input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *FMSActivities) PutAppsList(input *fms.PutAppsListInput) (*fms.PutAppsListOutput, error) {
    return a.client.PutAppsList(input)
}

func (a *FMSActivities) PutNotificationChannel(input *fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error) {
    return a.client.PutNotificationChannel(input)
}

func (a *FMSActivities) PutPolicy(input *fms.PutPolicyInput) (*fms.PutPolicyOutput, error) {
    return a.client.PutPolicy(input)
}

func (a *FMSActivities) PutProtocolsList(input *fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error) {
    return a.client.PutProtocolsList(input)
}

func (a *FMSActivities) TagResource(input *fms.TagResourceInput) (*fms.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *FMSActivities) UntagResource(input *fms.UntagResourceInput) (*fms.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
