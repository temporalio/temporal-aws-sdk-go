
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/aws/aws-sdk-go/service/sesv2/sesv2iface"
)

type SESV2Activities struct {
    client sesv2iface.SESV2API
}

func NewSESV2Activities(session *session.Session, config ...*aws.Config) *SESV2Activities {
    client := sesv2.New(session, config...)
    return &SESV2Activities{client: client}
}

func (a *SESV2Activities) CreateConfigurationSet(input *sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error) {
    return a.client.CreateConfigurationSet(input)
}

func (a *SESV2Activities) CreateConfigurationSetEventDestination(input *sesv2.CreateConfigurationSetEventDestinationInput) (*sesv2.CreateConfigurationSetEventDestinationOutput, error) {
    return a.client.CreateConfigurationSetEventDestination(input)
}

func (a *SESV2Activities) CreateCustomVerificationEmailTemplate(input *sesv2.CreateCustomVerificationEmailTemplateInput) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error) {
    return a.client.CreateCustomVerificationEmailTemplate(input)
}

func (a *SESV2Activities) CreateDedicatedIpPool(input *sesv2.CreateDedicatedIpPoolInput) (*sesv2.CreateDedicatedIpPoolOutput, error) {
    return a.client.CreateDedicatedIpPool(input)
}

func (a *SESV2Activities) CreateDeliverabilityTestReport(input *sesv2.CreateDeliverabilityTestReportInput) (*sesv2.CreateDeliverabilityTestReportOutput, error) {
    return a.client.CreateDeliverabilityTestReport(input)
}

func (a *SESV2Activities) CreateEmailIdentity(input *sesv2.CreateEmailIdentityInput) (*sesv2.CreateEmailIdentityOutput, error) {
    return a.client.CreateEmailIdentity(input)
}

func (a *SESV2Activities) CreateEmailIdentityPolicy(input *sesv2.CreateEmailIdentityPolicyInput) (*sesv2.CreateEmailIdentityPolicyOutput, error) {
    return a.client.CreateEmailIdentityPolicy(input)
}

func (a *SESV2Activities) CreateEmailTemplate(input *sesv2.CreateEmailTemplateInput) (*sesv2.CreateEmailTemplateOutput, error) {
    return a.client.CreateEmailTemplate(input)
}

func (a *SESV2Activities) CreateImportJob(input *sesv2.CreateImportJobInput) (*sesv2.CreateImportJobOutput, error) {
    return a.client.CreateImportJob(input)
}

func (a *SESV2Activities) DeleteConfigurationSet(input *sesv2.DeleteConfigurationSetInput) (*sesv2.DeleteConfigurationSetOutput, error) {
    return a.client.DeleteConfigurationSet(input)
}

func (a *SESV2Activities) DeleteConfigurationSetEventDestination(input *sesv2.DeleteConfigurationSetEventDestinationInput) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error) {
    return a.client.DeleteConfigurationSetEventDestination(input)
}

func (a *SESV2Activities) DeleteCustomVerificationEmailTemplate(input *sesv2.DeleteCustomVerificationEmailTemplateInput) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error) {
    return a.client.DeleteCustomVerificationEmailTemplate(input)
}

func (a *SESV2Activities) DeleteDedicatedIpPool(input *sesv2.DeleteDedicatedIpPoolInput) (*sesv2.DeleteDedicatedIpPoolOutput, error) {
    return a.client.DeleteDedicatedIpPool(input)
}

func (a *SESV2Activities) DeleteEmailIdentity(input *sesv2.DeleteEmailIdentityInput) (*sesv2.DeleteEmailIdentityOutput, error) {
    return a.client.DeleteEmailIdentity(input)
}

func (a *SESV2Activities) DeleteEmailIdentityPolicy(input *sesv2.DeleteEmailIdentityPolicyInput) (*sesv2.DeleteEmailIdentityPolicyOutput, error) {
    return a.client.DeleteEmailIdentityPolicy(input)
}

func (a *SESV2Activities) DeleteEmailTemplate(input *sesv2.DeleteEmailTemplateInput) (*sesv2.DeleteEmailTemplateOutput, error) {
    return a.client.DeleteEmailTemplate(input)
}

func (a *SESV2Activities) DeleteSuppressedDestination(input *sesv2.DeleteSuppressedDestinationInput) (*sesv2.DeleteSuppressedDestinationOutput, error) {
    return a.client.DeleteSuppressedDestination(input)
}

func (a *SESV2Activities) GetAccount(input *sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error) {
    return a.client.GetAccount(input)
}

func (a *SESV2Activities) GetBlacklistReports(input *sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error) {
    return a.client.GetBlacklistReports(input)
}

func (a *SESV2Activities) GetConfigurationSet(input *sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error) {
    return a.client.GetConfigurationSet(input)
}

func (a *SESV2Activities) GetConfigurationSetEventDestinations(input *sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
    return a.client.GetConfigurationSetEventDestinations(input)
}

func (a *SESV2Activities) GetCustomVerificationEmailTemplate(input *sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
    return a.client.GetCustomVerificationEmailTemplate(input)
}

func (a *SESV2Activities) GetDedicatedIp(input *sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error) {
    return a.client.GetDedicatedIp(input)
}

func (a *SESV2Activities) GetDedicatedIps(input *sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error) {
    return a.client.GetDedicatedIps(input)
}

func (a *SESV2Activities) GetDeliverabilityDashboardOptions(input *sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
    return a.client.GetDeliverabilityDashboardOptions(input)
}

func (a *SESV2Activities) GetDeliverabilityTestReport(input *sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error) {
    return a.client.GetDeliverabilityTestReport(input)
}

func (a *SESV2Activities) GetDomainDeliverabilityCampaign(input *sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
    return a.client.GetDomainDeliverabilityCampaign(input)
}

func (a *SESV2Activities) GetDomainStatisticsReport(input *sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error) {
    return a.client.GetDomainStatisticsReport(input)
}

func (a *SESV2Activities) GetEmailIdentity(input *sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error) {
    return a.client.GetEmailIdentity(input)
}

func (a *SESV2Activities) GetEmailIdentityPolicies(input *sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
    return a.client.GetEmailIdentityPolicies(input)
}

func (a *SESV2Activities) GetEmailTemplate(input *sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error) {
    return a.client.GetEmailTemplate(input)
}

func (a *SESV2Activities) GetImportJob(input *sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error) {
    return a.client.GetImportJob(input)
}

func (a *SESV2Activities) GetSuppressedDestination(input *sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error) {
    return a.client.GetSuppressedDestination(input)
}

func (a *SESV2Activities) ListConfigurationSets(input *sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error) {
    return a.client.ListConfigurationSets(input)
}

func (a *SESV2Activities) ListCustomVerificationEmailTemplates(input *sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
    return a.client.ListCustomVerificationEmailTemplates(input)
}

func (a *SESV2Activities) ListDedicatedIpPools(input *sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error) {
    return a.client.ListDedicatedIpPools(input)
}

func (a *SESV2Activities) ListDeliverabilityTestReports(input *sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
    return a.client.ListDeliverabilityTestReports(input)
}

func (a *SESV2Activities) ListDomainDeliverabilityCampaigns(input *sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
    return a.client.ListDomainDeliverabilityCampaigns(input)
}

func (a *SESV2Activities) ListEmailIdentities(input *sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error) {
    return a.client.ListEmailIdentities(input)
}

func (a *SESV2Activities) ListEmailTemplates(input *sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error) {
    return a.client.ListEmailTemplates(input)
}

func (a *SESV2Activities) ListImportJobs(input *sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error) {
    return a.client.ListImportJobs(input)
}

func (a *SESV2Activities) ListSuppressedDestinations(input *sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error) {
    return a.client.ListSuppressedDestinations(input)
}

func (a *SESV2Activities) ListTagsForResource(input *sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SESV2Activities) PutAccountDedicatedIpWarmupAttributes(input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutAccountDedicatedIpWarmupAttributes(input)
}

func (a *SESV2Activities) PutAccountDetails(input *sesv2.PutAccountDetailsInput) (*sesv2.PutAccountDetailsOutput, error) {
    return a.client.PutAccountDetails(input)
}

func (a *SESV2Activities) PutAccountSendingAttributes(input *sesv2.PutAccountSendingAttributesInput) (*sesv2.PutAccountSendingAttributesOutput, error) {
    return a.client.PutAccountSendingAttributes(input)
}

func (a *SESV2Activities) PutAccountSuppressionAttributes(input *sesv2.PutAccountSuppressionAttributesInput) (*sesv2.PutAccountSuppressionAttributesOutput, error) {
    return a.client.PutAccountSuppressionAttributes(input)
}

func (a *SESV2Activities) PutConfigurationSetDeliveryOptions(input *sesv2.PutConfigurationSetDeliveryOptionsInput) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error) {
    return a.client.PutConfigurationSetDeliveryOptions(input)
}

func (a *SESV2Activities) PutConfigurationSetReputationOptions(input *sesv2.PutConfigurationSetReputationOptionsInput) (*sesv2.PutConfigurationSetReputationOptionsOutput, error) {
    return a.client.PutConfigurationSetReputationOptions(input)
}

func (a *SESV2Activities) PutConfigurationSetSendingOptions(input *sesv2.PutConfigurationSetSendingOptionsInput) (*sesv2.PutConfigurationSetSendingOptionsOutput, error) {
    return a.client.PutConfigurationSetSendingOptions(input)
}

func (a *SESV2Activities) PutConfigurationSetSuppressionOptions(input *sesv2.PutConfigurationSetSuppressionOptionsInput) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error) {
    return a.client.PutConfigurationSetSuppressionOptions(input)
}

func (a *SESV2Activities) PutConfigurationSetTrackingOptions(input *sesv2.PutConfigurationSetTrackingOptionsInput) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error) {
    return a.client.PutConfigurationSetTrackingOptions(input)
}

func (a *SESV2Activities) PutDedicatedIpInPool(input *sesv2.PutDedicatedIpInPoolInput) (*sesv2.PutDedicatedIpInPoolOutput, error) {
    return a.client.PutDedicatedIpInPool(input)
}

func (a *SESV2Activities) PutDedicatedIpWarmupAttributes(input *sesv2.PutDedicatedIpWarmupAttributesInput) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error) {
    return a.client.PutDedicatedIpWarmupAttributes(input)
}

func (a *SESV2Activities) PutDeliverabilityDashboardOption(input *sesv2.PutDeliverabilityDashboardOptionInput) (*sesv2.PutDeliverabilityDashboardOptionOutput, error) {
    return a.client.PutDeliverabilityDashboardOption(input)
}

func (a *SESV2Activities) PutEmailIdentityDkimAttributes(input *sesv2.PutEmailIdentityDkimAttributesInput) (*sesv2.PutEmailIdentityDkimAttributesOutput, error) {
    return a.client.PutEmailIdentityDkimAttributes(input)
}

func (a *SESV2Activities) PutEmailIdentityDkimSigningAttributes(input *sesv2.PutEmailIdentityDkimSigningAttributesInput) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error) {
    return a.client.PutEmailIdentityDkimSigningAttributes(input)
}

func (a *SESV2Activities) PutEmailIdentityFeedbackAttributes(input *sesv2.PutEmailIdentityFeedbackAttributesInput) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error) {
    return a.client.PutEmailIdentityFeedbackAttributes(input)
}

func (a *SESV2Activities) PutEmailIdentityMailFromAttributes(input *sesv2.PutEmailIdentityMailFromAttributesInput) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error) {
    return a.client.PutEmailIdentityMailFromAttributes(input)
}

func (a *SESV2Activities) PutSuppressedDestination(input *sesv2.PutSuppressedDestinationInput) (*sesv2.PutSuppressedDestinationOutput, error) {
    return a.client.PutSuppressedDestination(input)
}

func (a *SESV2Activities) SendBulkEmail(input *sesv2.SendBulkEmailInput) (*sesv2.SendBulkEmailOutput, error) {
    return a.client.SendBulkEmail(input)
}

func (a *SESV2Activities) SendCustomVerificationEmail(input *sesv2.SendCustomVerificationEmailInput) (*sesv2.SendCustomVerificationEmailOutput, error) {
    return a.client.SendCustomVerificationEmail(input)
}

func (a *SESV2Activities) SendEmail(input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error) {
    return a.client.SendEmail(input)
}

func (a *SESV2Activities) TagResource(input *sesv2.TagResourceInput) (*sesv2.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SESV2Activities) TestRenderEmailTemplate(input *sesv2.TestRenderEmailTemplateInput) (*sesv2.TestRenderEmailTemplateOutput, error) {
    return a.client.TestRenderEmailTemplate(input)
}

func (a *SESV2Activities) UntagResource(input *sesv2.UntagResourceInput) (*sesv2.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *SESV2Activities) UpdateConfigurationSetEventDestination(input *sesv2.UpdateConfigurationSetEventDestinationInput) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error) {
    return a.client.UpdateConfigurationSetEventDestination(input)
}

func (a *SESV2Activities) UpdateCustomVerificationEmailTemplate(input *sesv2.UpdateCustomVerificationEmailTemplateInput) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error) {
    return a.client.UpdateCustomVerificationEmailTemplate(input)
}

func (a *SESV2Activities) UpdateEmailIdentityPolicy(input *sesv2.UpdateEmailIdentityPolicyInput) (*sesv2.UpdateEmailIdentityPolicyOutput, error) {
    return a.client.UpdateEmailIdentityPolicy(input)
}

func (a *SESV2Activities) UpdateEmailTemplate(input *sesv2.UpdateEmailTemplateInput) (*sesv2.UpdateEmailTemplateOutput, error) {
    return a.client.UpdateEmailTemplate(input)
}
