
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/aws/aws-sdk-go/service/sesv2/sesv2iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SESV2Activities struct {
    client sesv2iface.SESV2API
}

func NewSESV2Activities(session *session.Session, config ...*aws.Config) *SESV2Activities {
    client := sesv2.New(session, config...)
    return &SESV2Activities{client: client}
}

func (a *SESV2Activities) CreateConfigurationSet(ctx context.Context, input *sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error) {
    return a.client.CreateConfigurationSetWithContext(ctx, input)
}

func (a *SESV2Activities) CreateConfigurationSetEventDestination(ctx context.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) (*sesv2.CreateConfigurationSetEventDestinationOutput, error) {
    return a.client.CreateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) CreateCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error) {
    return a.client.CreateCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) CreateDedicatedIpPool(ctx context.Context, input *sesv2.CreateDedicatedIpPoolInput) (*sesv2.CreateDedicatedIpPoolOutput, error) {
    return a.client.CreateDedicatedIpPoolWithContext(ctx, input)
}

func (a *SESV2Activities) CreateDeliverabilityTestReport(ctx context.Context, input *sesv2.CreateDeliverabilityTestReportInput) (*sesv2.CreateDeliverabilityTestReportOutput, error) {
    return a.client.CreateDeliverabilityTestReportWithContext(ctx, input)
}

func (a *SESV2Activities) CreateEmailIdentity(ctx context.Context, input *sesv2.CreateEmailIdentityInput) (*sesv2.CreateEmailIdentityOutput, error) {
    return a.client.CreateEmailIdentityWithContext(ctx, input)
}

func (a *SESV2Activities) CreateEmailIdentityPolicy(ctx context.Context, input *sesv2.CreateEmailIdentityPolicyInput) (*sesv2.CreateEmailIdentityPolicyOutput, error) {
    return a.client.CreateEmailIdentityPolicyWithContext(ctx, input)
}

func (a *SESV2Activities) CreateEmailTemplate(ctx context.Context, input *sesv2.CreateEmailTemplateInput) (*sesv2.CreateEmailTemplateOutput, error) {
    return a.client.CreateEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) CreateImportJob(ctx context.Context, input *sesv2.CreateImportJobInput) (*sesv2.CreateImportJobOutput, error) {
    return a.client.CreateImportJobWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteConfigurationSet(ctx context.Context, input *sesv2.DeleteConfigurationSetInput) (*sesv2.DeleteConfigurationSetOutput, error) {
    return a.client.DeleteConfigurationSetWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteConfigurationSetEventDestination(ctx context.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error) {
    return a.client.DeleteConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error) {
    return a.client.DeleteCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteDedicatedIpPool(ctx context.Context, input *sesv2.DeleteDedicatedIpPoolInput) (*sesv2.DeleteDedicatedIpPoolOutput, error) {
    return a.client.DeleteDedicatedIpPoolWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteEmailIdentity(ctx context.Context, input *sesv2.DeleteEmailIdentityInput) (*sesv2.DeleteEmailIdentityOutput, error) {
    return a.client.DeleteEmailIdentityWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteEmailIdentityPolicy(ctx context.Context, input *sesv2.DeleteEmailIdentityPolicyInput) (*sesv2.DeleteEmailIdentityPolicyOutput, error) {
    return a.client.DeleteEmailIdentityPolicyWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteEmailTemplate(ctx context.Context, input *sesv2.DeleteEmailTemplateInput) (*sesv2.DeleteEmailTemplateOutput, error) {
    return a.client.DeleteEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteSuppressedDestination(ctx context.Context, input *sesv2.DeleteSuppressedDestinationInput) (*sesv2.DeleteSuppressedDestinationOutput, error) {
    return a.client.DeleteSuppressedDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) GetAccount(ctx context.Context, input *sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error) {
    return a.client.GetAccountWithContext(ctx, input)
}

func (a *SESV2Activities) GetBlacklistReports(ctx context.Context, input *sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error) {
    return a.client.GetBlacklistReportsWithContext(ctx, input)
}

func (a *SESV2Activities) GetConfigurationSet(ctx context.Context, input *sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error) {
    return a.client.GetConfigurationSetWithContext(ctx, input)
}

func (a *SESV2Activities) GetConfigurationSetEventDestinations(ctx context.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
    return a.client.GetConfigurationSetEventDestinationsWithContext(ctx, input)
}

func (a *SESV2Activities) GetCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
    return a.client.GetCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) GetDedicatedIp(ctx context.Context, input *sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error) {
    return a.client.GetDedicatedIpWithContext(ctx, input)
}

func (a *SESV2Activities) GetDedicatedIps(ctx context.Context, input *sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error) {
    return a.client.GetDedicatedIpsWithContext(ctx, input)
}

func (a *SESV2Activities) GetDeliverabilityDashboardOptions(ctx context.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
    return a.client.GetDeliverabilityDashboardOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) GetDeliverabilityTestReport(ctx context.Context, input *sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error) {
    return a.client.GetDeliverabilityTestReportWithContext(ctx, input)
}

func (a *SESV2Activities) GetDomainDeliverabilityCampaign(ctx context.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
    return a.client.GetDomainDeliverabilityCampaignWithContext(ctx, input)
}

func (a *SESV2Activities) GetDomainStatisticsReport(ctx context.Context, input *sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error) {
    return a.client.GetDomainStatisticsReportWithContext(ctx, input)
}

func (a *SESV2Activities) GetEmailIdentity(ctx context.Context, input *sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error) {
    return a.client.GetEmailIdentityWithContext(ctx, input)
}

func (a *SESV2Activities) GetEmailIdentityPolicies(ctx context.Context, input *sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
    return a.client.GetEmailIdentityPoliciesWithContext(ctx, input)
}

func (a *SESV2Activities) GetEmailTemplate(ctx context.Context, input *sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error) {
    return a.client.GetEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) GetImportJob(ctx context.Context, input *sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error) {
    return a.client.GetImportJobWithContext(ctx, input)
}

func (a *SESV2Activities) GetSuppressedDestination(ctx context.Context, input *sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error) {
    return a.client.GetSuppressedDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) ListConfigurationSets(ctx context.Context, input *sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error) {
    return a.client.ListConfigurationSetsWithContext(ctx, input)
}

func (a *SESV2Activities) ListCustomVerificationEmailTemplates(ctx context.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
    return a.client.ListCustomVerificationEmailTemplatesWithContext(ctx, input)
}

func (a *SESV2Activities) ListDedicatedIpPools(ctx context.Context, input *sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error) {
    return a.client.ListDedicatedIpPoolsWithContext(ctx, input)
}

func (a *SESV2Activities) ListDeliverabilityTestReports(ctx context.Context, input *sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
    return a.client.ListDeliverabilityTestReportsWithContext(ctx, input)
}

func (a *SESV2Activities) ListDomainDeliverabilityCampaigns(ctx context.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
    return a.client.ListDomainDeliverabilityCampaignsWithContext(ctx, input)
}

func (a *SESV2Activities) ListEmailIdentities(ctx context.Context, input *sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error) {
    return a.client.ListEmailIdentitiesWithContext(ctx, input)
}

func (a *SESV2Activities) ListEmailTemplates(ctx context.Context, input *sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error) {
    return a.client.ListEmailTemplatesWithContext(ctx, input)
}

func (a *SESV2Activities) ListImportJobs(ctx context.Context, input *sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error) {
    return a.client.ListImportJobsWithContext(ctx, input)
}

func (a *SESV2Activities) ListSuppressedDestinations(ctx context.Context, input *sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error) {
    return a.client.ListSuppressedDestinationsWithContext(ctx, input)
}

func (a *SESV2Activities) ListTagsForResource(ctx context.Context, input *sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountDedicatedIpWarmupAttributes(ctx context.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutAccountDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountDetails(ctx context.Context, input *sesv2.PutAccountDetailsInput) (*sesv2.PutAccountDetailsOutput, error) {
    return a.client.PutAccountDetailsWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountSendingAttributes(ctx context.Context, input *sesv2.PutAccountSendingAttributesInput) (*sesv2.PutAccountSendingAttributesOutput, error) {
    return a.client.PutAccountSendingAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountSuppressionAttributes(ctx context.Context, input *sesv2.PutAccountSuppressionAttributesInput) (*sesv2.PutAccountSuppressionAttributesOutput, error) {
    return a.client.PutAccountSuppressionAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetDeliveryOptions(ctx context.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error) {
    return a.client.PutConfigurationSetDeliveryOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetReputationOptions(ctx context.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) (*sesv2.PutConfigurationSetReputationOptionsOutput, error) {
    return a.client.PutConfigurationSetReputationOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetSendingOptions(ctx context.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) (*sesv2.PutConfigurationSetSendingOptionsOutput, error) {
    return a.client.PutConfigurationSetSendingOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetSuppressionOptions(ctx context.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error) {
    return a.client.PutConfigurationSetSuppressionOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetTrackingOptions(ctx context.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error) {
    return a.client.PutConfigurationSetTrackingOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutDedicatedIpInPool(ctx context.Context, input *sesv2.PutDedicatedIpInPoolInput) (*sesv2.PutDedicatedIpInPoolOutput, error) {
    return a.client.PutDedicatedIpInPoolWithContext(ctx, input)
}

func (a *SESV2Activities) PutDedicatedIpWarmupAttributes(ctx context.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutDeliverabilityDashboardOption(ctx context.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) (*sesv2.PutDeliverabilityDashboardOptionOutput, error) {
    return a.client.PutDeliverabilityDashboardOptionWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityDkimAttributes(ctx context.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) (*sesv2.PutEmailIdentityDkimAttributesOutput, error) {
    return a.client.PutEmailIdentityDkimAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityDkimSigningAttributes(ctx context.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error) {
    return a.client.PutEmailIdentityDkimSigningAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityFeedbackAttributes(ctx context.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error) {
    return a.client.PutEmailIdentityFeedbackAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityMailFromAttributes(ctx context.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error) {
    return a.client.PutEmailIdentityMailFromAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutSuppressedDestination(ctx context.Context, input *sesv2.PutSuppressedDestinationInput) (*sesv2.PutSuppressedDestinationOutput, error) {
    return a.client.PutSuppressedDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) SendBulkEmail(ctx context.Context, input *sesv2.SendBulkEmailInput) (*sesv2.SendBulkEmailOutput, error) {
    return a.client.SendBulkEmailWithContext(ctx, input)
}

func (a *SESV2Activities) SendCustomVerificationEmail(ctx context.Context, input *sesv2.SendCustomVerificationEmailInput) (*sesv2.SendCustomVerificationEmailOutput, error) {
    return a.client.SendCustomVerificationEmailWithContext(ctx, input)
}

func (a *SESV2Activities) SendEmail(ctx context.Context, input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error) {
    return a.client.SendEmailWithContext(ctx, input)
}

func (a *SESV2Activities) TagResource(ctx context.Context, input *sesv2.TagResourceInput) (*sesv2.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *SESV2Activities) TestRenderEmailTemplate(ctx context.Context, input *sesv2.TestRenderEmailTemplateInput) (*sesv2.TestRenderEmailTemplateOutput, error) {
    return a.client.TestRenderEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) UntagResource(ctx context.Context, input *sesv2.UntagResourceInput) (*sesv2.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateConfigurationSetEventDestination(ctx context.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error) {
    return a.client.UpdateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error) {
    return a.client.UpdateCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateEmailIdentityPolicy(ctx context.Context, input *sesv2.UpdateEmailIdentityPolicyInput) (*sesv2.UpdateEmailIdentityPolicyOutput, error) {
    return a.client.UpdateEmailIdentityPolicyWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateEmailTemplate(ctx context.Context, input *sesv2.UpdateEmailTemplateInput) (*sesv2.UpdateEmailTemplateOutput, error) {
    return a.client.UpdateEmailTemplateWithContext(ctx, input)
}
