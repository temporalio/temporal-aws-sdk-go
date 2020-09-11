
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/pinpointemail/pinpointemailiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type PinpointEmailActivities struct {
    client pinpointemailiface.PinpointEmailAPI
}

func NewPinpointEmailActivities(session *session.Session, config ...*aws.Config) *PinpointEmailActivities {
    client := pinpointemail.New(session, config...)
    return &PinpointEmailActivities{client: client}
}

func (a *PinpointEmailActivities) CreateConfigurationSet(ctx context.Context, input *pinpointemail.CreateConfigurationSetInput) (*pinpointemail.CreateConfigurationSetOutput, error) {
    return a.client.CreateConfigurationSetWithContext(ctx, input)
}

func (a *PinpointEmailActivities) CreateConfigurationSetEventDestination(ctx context.Context, input *pinpointemail.CreateConfigurationSetEventDestinationInput) (*pinpointemail.CreateConfigurationSetEventDestinationOutput, error) {
    return a.client.CreateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *PinpointEmailActivities) CreateDedicatedIpPool(ctx context.Context, input *pinpointemail.CreateDedicatedIpPoolInput) (*pinpointemail.CreateDedicatedIpPoolOutput, error) {
    return a.client.CreateDedicatedIpPoolWithContext(ctx, input)
}

func (a *PinpointEmailActivities) CreateDeliverabilityTestReport(ctx context.Context, input *pinpointemail.CreateDeliverabilityTestReportInput) (*pinpointemail.CreateDeliverabilityTestReportOutput, error) {
    return a.client.CreateDeliverabilityTestReportWithContext(ctx, input)
}

func (a *PinpointEmailActivities) CreateEmailIdentity(ctx context.Context, input *pinpointemail.CreateEmailIdentityInput) (*pinpointemail.CreateEmailIdentityOutput, error) {
    return a.client.CreateEmailIdentityWithContext(ctx, input)
}

func (a *PinpointEmailActivities) DeleteConfigurationSet(ctx context.Context, input *pinpointemail.DeleteConfigurationSetInput) (*pinpointemail.DeleteConfigurationSetOutput, error) {
    return a.client.DeleteConfigurationSetWithContext(ctx, input)
}

func (a *PinpointEmailActivities) DeleteConfigurationSetEventDestination(ctx context.Context, input *pinpointemail.DeleteConfigurationSetEventDestinationInput) (*pinpointemail.DeleteConfigurationSetEventDestinationOutput, error) {
    return a.client.DeleteConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *PinpointEmailActivities) DeleteDedicatedIpPool(ctx context.Context, input *pinpointemail.DeleteDedicatedIpPoolInput) (*pinpointemail.DeleteDedicatedIpPoolOutput, error) {
    return a.client.DeleteDedicatedIpPoolWithContext(ctx, input)
}

func (a *PinpointEmailActivities) DeleteEmailIdentity(ctx context.Context, input *pinpointemail.DeleteEmailIdentityInput) (*pinpointemail.DeleteEmailIdentityOutput, error) {
    return a.client.DeleteEmailIdentityWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetAccount(ctx context.Context, input *pinpointemail.GetAccountInput) (*pinpointemail.GetAccountOutput, error) {
    return a.client.GetAccountWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetBlacklistReports(ctx context.Context, input *pinpointemail.GetBlacklistReportsInput) (*pinpointemail.GetBlacklistReportsOutput, error) {
    return a.client.GetBlacklistReportsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetConfigurationSet(ctx context.Context, input *pinpointemail.GetConfigurationSetInput) (*pinpointemail.GetConfigurationSetOutput, error) {
    return a.client.GetConfigurationSetWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetConfigurationSetEventDestinations(ctx context.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error) {
    return a.client.GetConfigurationSetEventDestinationsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetDedicatedIp(ctx context.Context, input *pinpointemail.GetDedicatedIpInput) (*pinpointemail.GetDedicatedIpOutput, error) {
    return a.client.GetDedicatedIpWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetDedicatedIps(ctx context.Context, input *pinpointemail.GetDedicatedIpsInput) (*pinpointemail.GetDedicatedIpsOutput, error) {
    return a.client.GetDedicatedIpsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetDeliverabilityDashboardOptions(ctx context.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error) {
    return a.client.GetDeliverabilityDashboardOptionsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetDeliverabilityTestReport(ctx context.Context, input *pinpointemail.GetDeliverabilityTestReportInput) (*pinpointemail.GetDeliverabilityTestReportOutput, error) {
    return a.client.GetDeliverabilityTestReportWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetDomainDeliverabilityCampaign(ctx context.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error) {
    return a.client.GetDomainDeliverabilityCampaignWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetDomainStatisticsReport(ctx context.Context, input *pinpointemail.GetDomainStatisticsReportInput) (*pinpointemail.GetDomainStatisticsReportOutput, error) {
    return a.client.GetDomainStatisticsReportWithContext(ctx, input)
}

func (a *PinpointEmailActivities) GetEmailIdentity(ctx context.Context, input *pinpointemail.GetEmailIdentityInput) (*pinpointemail.GetEmailIdentityOutput, error) {
    return a.client.GetEmailIdentityWithContext(ctx, input)
}

func (a *PinpointEmailActivities) ListConfigurationSets(ctx context.Context, input *pinpointemail.ListConfigurationSetsInput) (*pinpointemail.ListConfigurationSetsOutput, error) {
    return a.client.ListConfigurationSetsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) ListDedicatedIpPools(ctx context.Context, input *pinpointemail.ListDedicatedIpPoolsInput) (*pinpointemail.ListDedicatedIpPoolsOutput, error) {
    return a.client.ListDedicatedIpPoolsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) ListDeliverabilityTestReports(ctx context.Context, input *pinpointemail.ListDeliverabilityTestReportsInput) (*pinpointemail.ListDeliverabilityTestReportsOutput, error) {
    return a.client.ListDeliverabilityTestReportsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) ListDomainDeliverabilityCampaigns(ctx context.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error) {
    return a.client.ListDomainDeliverabilityCampaignsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) ListEmailIdentities(ctx context.Context, input *pinpointemail.ListEmailIdentitiesInput) (*pinpointemail.ListEmailIdentitiesOutput, error) {
    return a.client.ListEmailIdentitiesWithContext(ctx, input)
}

func (a *PinpointEmailActivities) ListTagsForResource(ctx context.Context, input *pinpointemail.ListTagsForResourceInput) (*pinpointemail.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutAccountDedicatedIpWarmupAttributes(ctx context.Context, input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) (*pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutAccountDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutAccountSendingAttributes(ctx context.Context, input *pinpointemail.PutAccountSendingAttributesInput) (*pinpointemail.PutAccountSendingAttributesOutput, error) {
    return a.client.PutAccountSendingAttributesWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutConfigurationSetDeliveryOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) (*pinpointemail.PutConfigurationSetDeliveryOptionsOutput, error) {
    return a.client.PutConfigurationSetDeliveryOptionsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutConfigurationSetReputationOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetReputationOptionsInput) (*pinpointemail.PutConfigurationSetReputationOptionsOutput, error) {
    return a.client.PutConfigurationSetReputationOptionsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutConfigurationSetSendingOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetSendingOptionsInput) (*pinpointemail.PutConfigurationSetSendingOptionsOutput, error) {
    return a.client.PutConfigurationSetSendingOptionsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutConfigurationSetTrackingOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetTrackingOptionsInput) (*pinpointemail.PutConfigurationSetTrackingOptionsOutput, error) {
    return a.client.PutConfigurationSetTrackingOptionsWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutDedicatedIpInPool(ctx context.Context, input *pinpointemail.PutDedicatedIpInPoolInput) (*pinpointemail.PutDedicatedIpInPoolOutput, error) {
    return a.client.PutDedicatedIpInPoolWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutDedicatedIpWarmupAttributes(ctx context.Context, input *pinpointemail.PutDedicatedIpWarmupAttributesInput) (*pinpointemail.PutDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutDeliverabilityDashboardOption(ctx context.Context, input *pinpointemail.PutDeliverabilityDashboardOptionInput) (*pinpointemail.PutDeliverabilityDashboardOptionOutput, error) {
    return a.client.PutDeliverabilityDashboardOptionWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutEmailIdentityDkimAttributes(ctx context.Context, input *pinpointemail.PutEmailIdentityDkimAttributesInput) (*pinpointemail.PutEmailIdentityDkimAttributesOutput, error) {
    return a.client.PutEmailIdentityDkimAttributesWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutEmailIdentityFeedbackAttributes(ctx context.Context, input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) (*pinpointemail.PutEmailIdentityFeedbackAttributesOutput, error) {
    return a.client.PutEmailIdentityFeedbackAttributesWithContext(ctx, input)
}

func (a *PinpointEmailActivities) PutEmailIdentityMailFromAttributes(ctx context.Context, input *pinpointemail.PutEmailIdentityMailFromAttributesInput) (*pinpointemail.PutEmailIdentityMailFromAttributesOutput, error) {
    return a.client.PutEmailIdentityMailFromAttributesWithContext(ctx, input)
}

func (a *PinpointEmailActivities) SendEmail(ctx context.Context, input *pinpointemail.SendEmailInput) (*pinpointemail.SendEmailOutput, error) {
    return a.client.SendEmailWithContext(ctx, input)
}

func (a *PinpointEmailActivities) TagResource(ctx context.Context, input *pinpointemail.TagResourceInput) (*pinpointemail.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *PinpointEmailActivities) UntagResource(ctx context.Context, input *pinpointemail.UntagResourceInput) (*pinpointemail.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *PinpointEmailActivities) UpdateConfigurationSetEventDestination(ctx context.Context, input *pinpointemail.UpdateConfigurationSetEventDestinationInput) (*pinpointemail.UpdateConfigurationSetEventDestinationOutput, error) {
    return a.client.UpdateConfigurationSetEventDestinationWithContext(ctx, input)
}
