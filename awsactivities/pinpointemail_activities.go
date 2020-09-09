
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/pinpointemail/pinpointemailiface"
)

type PinpointEmailActivities struct {
    client pinpointemailiface.PinpointEmailAPI
}

func NewPinpointEmailActivities(session *session.Session, config ...*aws.Config) *PinpointEmailActivities {
    client := pinpointemail.New(session, config...)
    return &PinpointEmailActivities{client: client}
}

func (a *PinpointEmailActivities) CreateConfigurationSet(input *pinpointemail.CreateConfigurationSetInput) (*pinpointemail.CreateConfigurationSetOutput, error) {
    return a.client.CreateConfigurationSet(input)
}

func (a *PinpointEmailActivities) CreateConfigurationSetEventDestination(input *pinpointemail.CreateConfigurationSetEventDestinationInput) (*pinpointemail.CreateConfigurationSetEventDestinationOutput, error) {
    return a.client.CreateConfigurationSetEventDestination(input)
}

func (a *PinpointEmailActivities) CreateDedicatedIpPool(input *pinpointemail.CreateDedicatedIpPoolInput) (*pinpointemail.CreateDedicatedIpPoolOutput, error) {
    return a.client.CreateDedicatedIpPool(input)
}

func (a *PinpointEmailActivities) CreateDeliverabilityTestReport(input *pinpointemail.CreateDeliverabilityTestReportInput) (*pinpointemail.CreateDeliverabilityTestReportOutput, error) {
    return a.client.CreateDeliverabilityTestReport(input)
}

func (a *PinpointEmailActivities) CreateEmailIdentity(input *pinpointemail.CreateEmailIdentityInput) (*pinpointemail.CreateEmailIdentityOutput, error) {
    return a.client.CreateEmailIdentity(input)
}

func (a *PinpointEmailActivities) DeleteConfigurationSet(input *pinpointemail.DeleteConfigurationSetInput) (*pinpointemail.DeleteConfigurationSetOutput, error) {
    return a.client.DeleteConfigurationSet(input)
}

func (a *PinpointEmailActivities) DeleteConfigurationSetEventDestination(input *pinpointemail.DeleteConfigurationSetEventDestinationInput) (*pinpointemail.DeleteConfigurationSetEventDestinationOutput, error) {
    return a.client.DeleteConfigurationSetEventDestination(input)
}

func (a *PinpointEmailActivities) DeleteDedicatedIpPool(input *pinpointemail.DeleteDedicatedIpPoolInput) (*pinpointemail.DeleteDedicatedIpPoolOutput, error) {
    return a.client.DeleteDedicatedIpPool(input)
}

func (a *PinpointEmailActivities) DeleteEmailIdentity(input *pinpointemail.DeleteEmailIdentityInput) (*pinpointemail.DeleteEmailIdentityOutput, error) {
    return a.client.DeleteEmailIdentity(input)
}

func (a *PinpointEmailActivities) GetAccount(input *pinpointemail.GetAccountInput) (*pinpointemail.GetAccountOutput, error) {
    return a.client.GetAccount(input)
}

func (a *PinpointEmailActivities) GetBlacklistReports(input *pinpointemail.GetBlacklistReportsInput) (*pinpointemail.GetBlacklistReportsOutput, error) {
    return a.client.GetBlacklistReports(input)
}

func (a *PinpointEmailActivities) GetConfigurationSet(input *pinpointemail.GetConfigurationSetInput) (*pinpointemail.GetConfigurationSetOutput, error) {
    return a.client.GetConfigurationSet(input)
}

func (a *PinpointEmailActivities) GetConfigurationSetEventDestinations(input *pinpointemail.GetConfigurationSetEventDestinationsInput) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error) {
    return a.client.GetConfigurationSetEventDestinations(input)
}

func (a *PinpointEmailActivities) GetDedicatedIp(input *pinpointemail.GetDedicatedIpInput) (*pinpointemail.GetDedicatedIpOutput, error) {
    return a.client.GetDedicatedIp(input)
}

func (a *PinpointEmailActivities) GetDedicatedIps(input *pinpointemail.GetDedicatedIpsInput) (*pinpointemail.GetDedicatedIpsOutput, error) {
    return a.client.GetDedicatedIps(input)
}

func (a *PinpointEmailActivities) GetDeliverabilityDashboardOptions(input *pinpointemail.GetDeliverabilityDashboardOptionsInput) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error) {
    return a.client.GetDeliverabilityDashboardOptions(input)
}

func (a *PinpointEmailActivities) GetDeliverabilityTestReport(input *pinpointemail.GetDeliverabilityTestReportInput) (*pinpointemail.GetDeliverabilityTestReportOutput, error) {
    return a.client.GetDeliverabilityTestReport(input)
}

func (a *PinpointEmailActivities) GetDomainDeliverabilityCampaign(input *pinpointemail.GetDomainDeliverabilityCampaignInput) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error) {
    return a.client.GetDomainDeliverabilityCampaign(input)
}

func (a *PinpointEmailActivities) GetDomainStatisticsReport(input *pinpointemail.GetDomainStatisticsReportInput) (*pinpointemail.GetDomainStatisticsReportOutput, error) {
    return a.client.GetDomainStatisticsReport(input)
}

func (a *PinpointEmailActivities) GetEmailIdentity(input *pinpointemail.GetEmailIdentityInput) (*pinpointemail.GetEmailIdentityOutput, error) {
    return a.client.GetEmailIdentity(input)
}

func (a *PinpointEmailActivities) ListConfigurationSets(input *pinpointemail.ListConfigurationSetsInput) (*pinpointemail.ListConfigurationSetsOutput, error) {
    return a.client.ListConfigurationSets(input)
}

func (a *PinpointEmailActivities) ListDedicatedIpPools(input *pinpointemail.ListDedicatedIpPoolsInput) (*pinpointemail.ListDedicatedIpPoolsOutput, error) {
    return a.client.ListDedicatedIpPools(input)
}

func (a *PinpointEmailActivities) ListDeliverabilityTestReports(input *pinpointemail.ListDeliverabilityTestReportsInput) (*pinpointemail.ListDeliverabilityTestReportsOutput, error) {
    return a.client.ListDeliverabilityTestReports(input)
}

func (a *PinpointEmailActivities) ListDomainDeliverabilityCampaigns(input *pinpointemail.ListDomainDeliverabilityCampaignsInput) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error) {
    return a.client.ListDomainDeliverabilityCampaigns(input)
}

func (a *PinpointEmailActivities) ListEmailIdentities(input *pinpointemail.ListEmailIdentitiesInput) (*pinpointemail.ListEmailIdentitiesOutput, error) {
    return a.client.ListEmailIdentities(input)
}

func (a *PinpointEmailActivities) ListTagsForResource(input *pinpointemail.ListTagsForResourceInput) (*pinpointemail.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *PinpointEmailActivities) PutAccountDedicatedIpWarmupAttributes(input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) (*pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutAccountDedicatedIpWarmupAttributes(input)
}

func (a *PinpointEmailActivities) PutAccountSendingAttributes(input *pinpointemail.PutAccountSendingAttributesInput) (*pinpointemail.PutAccountSendingAttributesOutput, error) {
    return a.client.PutAccountSendingAttributes(input)
}

func (a *PinpointEmailActivities) PutConfigurationSetDeliveryOptions(input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) (*pinpointemail.PutConfigurationSetDeliveryOptionsOutput, error) {
    return a.client.PutConfigurationSetDeliveryOptions(input)
}

func (a *PinpointEmailActivities) PutConfigurationSetReputationOptions(input *pinpointemail.PutConfigurationSetReputationOptionsInput) (*pinpointemail.PutConfigurationSetReputationOptionsOutput, error) {
    return a.client.PutConfigurationSetReputationOptions(input)
}

func (a *PinpointEmailActivities) PutConfigurationSetSendingOptions(input *pinpointemail.PutConfigurationSetSendingOptionsInput) (*pinpointemail.PutConfigurationSetSendingOptionsOutput, error) {
    return a.client.PutConfigurationSetSendingOptions(input)
}

func (a *PinpointEmailActivities) PutConfigurationSetTrackingOptions(input *pinpointemail.PutConfigurationSetTrackingOptionsInput) (*pinpointemail.PutConfigurationSetTrackingOptionsOutput, error) {
    return a.client.PutConfigurationSetTrackingOptions(input)
}

func (a *PinpointEmailActivities) PutDedicatedIpInPool(input *pinpointemail.PutDedicatedIpInPoolInput) (*pinpointemail.PutDedicatedIpInPoolOutput, error) {
    return a.client.PutDedicatedIpInPool(input)
}

func (a *PinpointEmailActivities) PutDedicatedIpWarmupAttributes(input *pinpointemail.PutDedicatedIpWarmupAttributesInput) (*pinpointemail.PutDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutDedicatedIpWarmupAttributes(input)
}

func (a *PinpointEmailActivities) PutDeliverabilityDashboardOption(input *pinpointemail.PutDeliverabilityDashboardOptionInput) (*pinpointemail.PutDeliverabilityDashboardOptionOutput, error) {
    return a.client.PutDeliverabilityDashboardOption(input)
}

func (a *PinpointEmailActivities) PutEmailIdentityDkimAttributes(input *pinpointemail.PutEmailIdentityDkimAttributesInput) (*pinpointemail.PutEmailIdentityDkimAttributesOutput, error) {
    return a.client.PutEmailIdentityDkimAttributes(input)
}

func (a *PinpointEmailActivities) PutEmailIdentityFeedbackAttributes(input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) (*pinpointemail.PutEmailIdentityFeedbackAttributesOutput, error) {
    return a.client.PutEmailIdentityFeedbackAttributes(input)
}

func (a *PinpointEmailActivities) PutEmailIdentityMailFromAttributes(input *pinpointemail.PutEmailIdentityMailFromAttributesInput) (*pinpointemail.PutEmailIdentityMailFromAttributesOutput, error) {
    return a.client.PutEmailIdentityMailFromAttributes(input)
}

func (a *PinpointEmailActivities) SendEmail(input *pinpointemail.SendEmailInput) (*pinpointemail.SendEmailOutput, error) {
    return a.client.SendEmail(input)
}

func (a *PinpointEmailActivities) TagResource(input *pinpointemail.TagResourceInput) (*pinpointemail.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *PinpointEmailActivities) UntagResource(input *pinpointemail.UntagResourceInput) (*pinpointemail.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *PinpointEmailActivities) UpdateConfigurationSetEventDestination(input *pinpointemail.UpdateConfigurationSetEventDestinationInput) (*pinpointemail.UpdateConfigurationSetEventDestinationOutput, error) {
    return a.client.UpdateConfigurationSetEventDestination(input)
}
