
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

type SESActivities struct {
	client sesiface.SESAPI
}

func NewSESActivities(session *session.Session, config... *aws.Config) *SESActivities {
    client := ses.New(session, config...)
    return &SESActivities{client: client}
}

func (a *SESActivities) CloneReceiptRuleSet(input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error) {
    return a.client.CloneReceiptRuleSet(input)
}

func (a *SESActivities) CreateConfigurationSet(input *ses.CreateConfigurationSetInput) (*ses.CreateConfigurationSetOutput, error) {
    return a.client.CreateConfigurationSet(input)
}

func (a *SESActivities) CreateConfigurationSetEventDestination(input *ses.CreateConfigurationSetEventDestinationInput) (*ses.CreateConfigurationSetEventDestinationOutput, error) {
    return a.client.CreateConfigurationSetEventDestination(input)
}

func (a *SESActivities) CreateConfigurationSetTrackingOptions(input *ses.CreateConfigurationSetTrackingOptionsInput) (*ses.CreateConfigurationSetTrackingOptionsOutput, error) {
    return a.client.CreateConfigurationSetTrackingOptions(input)
}

func (a *SESActivities) CreateCustomVerificationEmailTemplate(input *ses.CreateCustomVerificationEmailTemplateInput) (*ses.CreateCustomVerificationEmailTemplateOutput, error) {
    return a.client.CreateCustomVerificationEmailTemplate(input)
}

func (a *SESActivities) CreateReceiptFilter(input *ses.CreateReceiptFilterInput) (*ses.CreateReceiptFilterOutput, error) {
    return a.client.CreateReceiptFilter(input)
}

func (a *SESActivities) CreateReceiptRule(input *ses.CreateReceiptRuleInput) (*ses.CreateReceiptRuleOutput, error) {
    return a.client.CreateReceiptRule(input)
}

func (a *SESActivities) CreateReceiptRuleSet(input *ses.CreateReceiptRuleSetInput) (*ses.CreateReceiptRuleSetOutput, error) {
    return a.client.CreateReceiptRuleSet(input)
}

func (a *SESActivities) CreateTemplate(input *ses.CreateTemplateInput) (*ses.CreateTemplateOutput, error) {
    return a.client.CreateTemplate(input)
}

func (a *SESActivities) DeleteConfigurationSet(input *ses.DeleteConfigurationSetInput) (*ses.DeleteConfigurationSetOutput, error) {
    return a.client.DeleteConfigurationSet(input)
}

func (a *SESActivities) DeleteConfigurationSetEventDestination(input *ses.DeleteConfigurationSetEventDestinationInput) (*ses.DeleteConfigurationSetEventDestinationOutput, error) {
    return a.client.DeleteConfigurationSetEventDestination(input)
}

func (a *SESActivities) DeleteConfigurationSetTrackingOptions(input *ses.DeleteConfigurationSetTrackingOptionsInput) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error) {
    return a.client.DeleteConfigurationSetTrackingOptions(input)
}

func (a *SESActivities) DeleteCustomVerificationEmailTemplate(input *ses.DeleteCustomVerificationEmailTemplateInput) (*ses.DeleteCustomVerificationEmailTemplateOutput, error) {
    return a.client.DeleteCustomVerificationEmailTemplate(input)
}

func (a *SESActivities) DeleteIdentity(input *ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error) {
    return a.client.DeleteIdentity(input)
}

func (a *SESActivities) DeleteIdentityPolicy(input *ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error) {
    return a.client.DeleteIdentityPolicy(input)
}

func (a *SESActivities) DeleteReceiptFilter(input *ses.DeleteReceiptFilterInput) (*ses.DeleteReceiptFilterOutput, error) {
    return a.client.DeleteReceiptFilter(input)
}

func (a *SESActivities) DeleteReceiptRule(input *ses.DeleteReceiptRuleInput) (*ses.DeleteReceiptRuleOutput, error) {
    return a.client.DeleteReceiptRule(input)
}

func (a *SESActivities) DeleteReceiptRuleSet(input *ses.DeleteReceiptRuleSetInput) (*ses.DeleteReceiptRuleSetOutput, error) {
    return a.client.DeleteReceiptRuleSet(input)
}

func (a *SESActivities) DeleteTemplate(input *ses.DeleteTemplateInput) (*ses.DeleteTemplateOutput, error) {
    return a.client.DeleteTemplate(input)
}

func (a *SESActivities) DeleteVerifiedEmailAddress(input *ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error) {
    return a.client.DeleteVerifiedEmailAddress(input)
}

func (a *SESActivities) DescribeActiveReceiptRuleSet(input *ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error) {
    return a.client.DescribeActiveReceiptRuleSet(input)
}

func (a *SESActivities) DescribeConfigurationSet(input *ses.DescribeConfigurationSetInput) (*ses.DescribeConfigurationSetOutput, error) {
    return a.client.DescribeConfigurationSet(input)
}

func (a *SESActivities) DescribeReceiptRule(input *ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error) {
    return a.client.DescribeReceiptRule(input)
}

func (a *SESActivities) DescribeReceiptRuleSet(input *ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error) {
    return a.client.DescribeReceiptRuleSet(input)
}

func (a *SESActivities) GetAccountSendingEnabled(input *ses.GetAccountSendingEnabledInput) (*ses.GetAccountSendingEnabledOutput, error) {
    return a.client.GetAccountSendingEnabled(input)
}

func (a *SESActivities) GetCustomVerificationEmailTemplate(input *ses.GetCustomVerificationEmailTemplateInput) (*ses.GetCustomVerificationEmailTemplateOutput, error) {
    return a.client.GetCustomVerificationEmailTemplate(input)
}

func (a *SESActivities) GetIdentityDkimAttributes(input *ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error) {
    return a.client.GetIdentityDkimAttributes(input)
}

func (a *SESActivities) GetIdentityMailFromDomainAttributes(input *ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error) {
    return a.client.GetIdentityMailFromDomainAttributes(input)
}

func (a *SESActivities) GetIdentityNotificationAttributes(input *ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error) {
    return a.client.GetIdentityNotificationAttributes(input)
}

func (a *SESActivities) GetIdentityPolicies(input *ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error) {
    return a.client.GetIdentityPolicies(input)
}

func (a *SESActivities) GetIdentityVerificationAttributes(input *ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error) {
    return a.client.GetIdentityVerificationAttributes(input)
}

func (a *SESActivities) GetSendQuota(input *ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error) {
    return a.client.GetSendQuota(input)
}

func (a *SESActivities) GetSendStatistics(input *ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error) {
    return a.client.GetSendStatistics(input)
}

func (a *SESActivities) GetTemplate(input *ses.GetTemplateInput) (*ses.GetTemplateOutput, error) {
    return a.client.GetTemplate(input)
}

func (a *SESActivities) ListConfigurationSets(input *ses.ListConfigurationSetsInput) (*ses.ListConfigurationSetsOutput, error) {
    return a.client.ListConfigurationSets(input)
}

func (a *SESActivities) ListCustomVerificationEmailTemplates(input *ses.ListCustomVerificationEmailTemplatesInput) (*ses.ListCustomVerificationEmailTemplatesOutput, error) {
    return a.client.ListCustomVerificationEmailTemplates(input)
}

func (a *SESActivities) ListIdentities(input *ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error) {
    return a.client.ListIdentities(input)
}

func (a *SESActivities) ListIdentityPolicies(input *ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error) {
    return a.client.ListIdentityPolicies(input)
}

func (a *SESActivities) ListReceiptFilters(input *ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error) {
    return a.client.ListReceiptFilters(input)
}

func (a *SESActivities) ListReceiptRuleSets(input *ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error) {
    return a.client.ListReceiptRuleSets(input)
}

func (a *SESActivities) ListTemplates(input *ses.ListTemplatesInput) (*ses.ListTemplatesOutput, error) {
    return a.client.ListTemplates(input)
}

func (a *SESActivities) ListVerifiedEmailAddresses(input *ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error) {
    return a.client.ListVerifiedEmailAddresses(input)
}

func (a *SESActivities) PutConfigurationSetDeliveryOptions(input *ses.PutConfigurationSetDeliveryOptionsInput) (*ses.PutConfigurationSetDeliveryOptionsOutput, error) {
    return a.client.PutConfigurationSetDeliveryOptions(input)
}

func (a *SESActivities) PutIdentityPolicy(input *ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error) {
    return a.client.PutIdentityPolicy(input)
}

func (a *SESActivities) ReorderReceiptRuleSet(input *ses.ReorderReceiptRuleSetInput) (*ses.ReorderReceiptRuleSetOutput, error) {
    return a.client.ReorderReceiptRuleSet(input)
}

func (a *SESActivities) SendBounce(input *ses.SendBounceInput) (*ses.SendBounceOutput, error) {
    return a.client.SendBounce(input)
}

func (a *SESActivities) SendBulkTemplatedEmail(input *ses.SendBulkTemplatedEmailInput) (*ses.SendBulkTemplatedEmailOutput, error) {
    return a.client.SendBulkTemplatedEmail(input)
}

func (a *SESActivities) SendCustomVerificationEmail(input *ses.SendCustomVerificationEmailInput) (*ses.SendCustomVerificationEmailOutput, error) {
    return a.client.SendCustomVerificationEmail(input)
}

func (a *SESActivities) SendEmail(input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
    return a.client.SendEmail(input)
}

func (a *SESActivities) SendRawEmail(input *ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error) {
    return a.client.SendRawEmail(input)
}

func (a *SESActivities) SendTemplatedEmail(input *ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error) {
    return a.client.SendTemplatedEmail(input)
}

func (a *SESActivities) SetActiveReceiptRuleSet(input *ses.SetActiveReceiptRuleSetInput) (*ses.SetActiveReceiptRuleSetOutput, error) {
    return a.client.SetActiveReceiptRuleSet(input)
}

func (a *SESActivities) SetIdentityDkimEnabled(input *ses.SetIdentityDkimEnabledInput) (*ses.SetIdentityDkimEnabledOutput, error) {
    return a.client.SetIdentityDkimEnabled(input)
}

func (a *SESActivities) SetIdentityFeedbackForwardingEnabled(input *ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error) {
    return a.client.SetIdentityFeedbackForwardingEnabled(input)
}

func (a *SESActivities) SetIdentityHeadersInNotificationsEnabled(input *ses.SetIdentityHeadersInNotificationsEnabledInput) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error) {
    return a.client.SetIdentityHeadersInNotificationsEnabled(input)
}

func (a *SESActivities) SetIdentityMailFromDomain(input *ses.SetIdentityMailFromDomainInput) (*ses.SetIdentityMailFromDomainOutput, error) {
    return a.client.SetIdentityMailFromDomain(input)
}

func (a *SESActivities) SetIdentityNotificationTopic(input *ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error) {
    return a.client.SetIdentityNotificationTopic(input)
}

func (a *SESActivities) SetReceiptRulePosition(input *ses.SetReceiptRulePositionInput) (*ses.SetReceiptRulePositionOutput, error) {
    return a.client.SetReceiptRulePosition(input)
}

func (a *SESActivities) TestRenderTemplate(input *ses.TestRenderTemplateInput) (*ses.TestRenderTemplateOutput, error) {
    return a.client.TestRenderTemplate(input)
}

func (a *SESActivities) UpdateAccountSendingEnabled(input *ses.UpdateAccountSendingEnabledInput) (*ses.UpdateAccountSendingEnabledOutput, error) {
    return a.client.UpdateAccountSendingEnabled(input)
}

func (a *SESActivities) UpdateConfigurationSetEventDestination(input *ses.UpdateConfigurationSetEventDestinationInput) (*ses.UpdateConfigurationSetEventDestinationOutput, error) {
    return a.client.UpdateConfigurationSetEventDestination(input)
}

func (a *SESActivities) UpdateConfigurationSetReputationMetricsEnabled(input *ses.UpdateConfigurationSetReputationMetricsEnabledInput) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error) {
    return a.client.UpdateConfigurationSetReputationMetricsEnabled(input)
}

func (a *SESActivities) UpdateConfigurationSetSendingEnabled(input *ses.UpdateConfigurationSetSendingEnabledInput) (*ses.UpdateConfigurationSetSendingEnabledOutput, error) {
    return a.client.UpdateConfigurationSetSendingEnabled(input)
}

func (a *SESActivities) UpdateConfigurationSetTrackingOptions(input *ses.UpdateConfigurationSetTrackingOptionsInput) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error) {
    return a.client.UpdateConfigurationSetTrackingOptions(input)
}

func (a *SESActivities) UpdateCustomVerificationEmailTemplate(input *ses.UpdateCustomVerificationEmailTemplateInput) (*ses.UpdateCustomVerificationEmailTemplateOutput, error) {
    return a.client.UpdateCustomVerificationEmailTemplate(input)
}

func (a *SESActivities) UpdateReceiptRule(input *ses.UpdateReceiptRuleInput) (*ses.UpdateReceiptRuleOutput, error) {
    return a.client.UpdateReceiptRule(input)
}

func (a *SESActivities) UpdateTemplate(input *ses.UpdateTemplateInput) (*ses.UpdateTemplateOutput, error) {
    return a.client.UpdateTemplate(input)
}

func (a *SESActivities) VerifyDomainDkim(input *ses.VerifyDomainDkimInput) (*ses.VerifyDomainDkimOutput, error) {
    return a.client.VerifyDomainDkim(input)
}

func (a *SESActivities) VerifyDomainIdentity(input *ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error) {
    return a.client.VerifyDomainIdentity(input)
}

func (a *SESActivities) VerifyEmailAddress(input *ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error) {
    return a.client.VerifyEmailAddress(input)
}

func (a *SESActivities) VerifyEmailIdentity(input *ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error) {
    return a.client.VerifyEmailIdentity(input)
}

func (a *SESActivities) WaitUntilIdentityExists(input *ses.GetIdentityVerificationAttributesInput) error {
    return a.client.WaitUntilIdentityExists(input)
}
