package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SESActivities struct {
	client sesiface.SESAPI
}

func NewSESActivities(session *session.Session, config ...*aws.Config) *SESActivities {
	client := ses.New(session, config...)
	return &SESActivities{client: client}
}

func (a *SESActivities) CloneReceiptRuleSet(ctx context.Context, input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error) {
	return a.client.CloneReceiptRuleSetWithContext(ctx, input)
}

func (a *SESActivities) CreateConfigurationSet(ctx context.Context, input *ses.CreateConfigurationSetInput) (*ses.CreateConfigurationSetOutput, error) {
	return a.client.CreateConfigurationSetWithContext(ctx, input)
}

func (a *SESActivities) CreateConfigurationSetEventDestination(ctx context.Context, input *ses.CreateConfigurationSetEventDestinationInput) (*ses.CreateConfigurationSetEventDestinationOutput, error) {
	return a.client.CreateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESActivities) CreateConfigurationSetTrackingOptions(ctx context.Context, input *ses.CreateConfigurationSetTrackingOptionsInput) (*ses.CreateConfigurationSetTrackingOptionsOutput, error) {
	return a.client.CreateConfigurationSetTrackingOptionsWithContext(ctx, input)
}

func (a *SESActivities) CreateCustomVerificationEmailTemplate(ctx context.Context, input *ses.CreateCustomVerificationEmailTemplateInput) (*ses.CreateCustomVerificationEmailTemplateOutput, error) {
	return a.client.CreateCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESActivities) CreateReceiptFilter(ctx context.Context, input *ses.CreateReceiptFilterInput) (*ses.CreateReceiptFilterOutput, error) {
	return a.client.CreateReceiptFilterWithContext(ctx, input)
}

func (a *SESActivities) CreateReceiptRule(ctx context.Context, input *ses.CreateReceiptRuleInput) (*ses.CreateReceiptRuleOutput, error) {
	return a.client.CreateReceiptRuleWithContext(ctx, input)
}

func (a *SESActivities) CreateReceiptRuleSet(ctx context.Context, input *ses.CreateReceiptRuleSetInput) (*ses.CreateReceiptRuleSetOutput, error) {
	return a.client.CreateReceiptRuleSetWithContext(ctx, input)
}

func (a *SESActivities) CreateTemplate(ctx context.Context, input *ses.CreateTemplateInput) (*ses.CreateTemplateOutput, error) {
	return a.client.CreateTemplateWithContext(ctx, input)
}

func (a *SESActivities) DeleteConfigurationSet(ctx context.Context, input *ses.DeleteConfigurationSetInput) (*ses.DeleteConfigurationSetOutput, error) {
	return a.client.DeleteConfigurationSetWithContext(ctx, input)
}

func (a *SESActivities) DeleteConfigurationSetEventDestination(ctx context.Context, input *ses.DeleteConfigurationSetEventDestinationInput) (*ses.DeleteConfigurationSetEventDestinationOutput, error) {
	return a.client.DeleteConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESActivities) DeleteConfigurationSetTrackingOptions(ctx context.Context, input *ses.DeleteConfigurationSetTrackingOptionsInput) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error) {
	return a.client.DeleteConfigurationSetTrackingOptionsWithContext(ctx, input)
}

func (a *SESActivities) DeleteCustomVerificationEmailTemplate(ctx context.Context, input *ses.DeleteCustomVerificationEmailTemplateInput) (*ses.DeleteCustomVerificationEmailTemplateOutput, error) {
	return a.client.DeleteCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESActivities) DeleteIdentity(ctx context.Context, input *ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error) {
	return a.client.DeleteIdentityWithContext(ctx, input)
}

func (a *SESActivities) DeleteIdentityPolicy(ctx context.Context, input *ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error) {
	return a.client.DeleteIdentityPolicyWithContext(ctx, input)
}

func (a *SESActivities) DeleteReceiptFilter(ctx context.Context, input *ses.DeleteReceiptFilterInput) (*ses.DeleteReceiptFilterOutput, error) {
	return a.client.DeleteReceiptFilterWithContext(ctx, input)
}

func (a *SESActivities) DeleteReceiptRule(ctx context.Context, input *ses.DeleteReceiptRuleInput) (*ses.DeleteReceiptRuleOutput, error) {
	return a.client.DeleteReceiptRuleWithContext(ctx, input)
}

func (a *SESActivities) DeleteReceiptRuleSet(ctx context.Context, input *ses.DeleteReceiptRuleSetInput) (*ses.DeleteReceiptRuleSetOutput, error) {
	return a.client.DeleteReceiptRuleSetWithContext(ctx, input)
}

func (a *SESActivities) DeleteTemplate(ctx context.Context, input *ses.DeleteTemplateInput) (*ses.DeleteTemplateOutput, error) {
	return a.client.DeleteTemplateWithContext(ctx, input)
}

func (a *SESActivities) DeleteVerifiedEmailAddress(ctx context.Context, input *ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error) {
	return a.client.DeleteVerifiedEmailAddressWithContext(ctx, input)
}

func (a *SESActivities) DescribeActiveReceiptRuleSet(ctx context.Context, input *ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error) {
	return a.client.DescribeActiveReceiptRuleSetWithContext(ctx, input)
}

func (a *SESActivities) DescribeConfigurationSet(ctx context.Context, input *ses.DescribeConfigurationSetInput) (*ses.DescribeConfigurationSetOutput, error) {
	return a.client.DescribeConfigurationSetWithContext(ctx, input)
}

func (a *SESActivities) DescribeReceiptRule(ctx context.Context, input *ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error) {
	return a.client.DescribeReceiptRuleWithContext(ctx, input)
}

func (a *SESActivities) DescribeReceiptRuleSet(ctx context.Context, input *ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error) {
	return a.client.DescribeReceiptRuleSetWithContext(ctx, input)
}

func (a *SESActivities) GetAccountSendingEnabled(ctx context.Context, input *ses.GetAccountSendingEnabledInput) (*ses.GetAccountSendingEnabledOutput, error) {
	return a.client.GetAccountSendingEnabledWithContext(ctx, input)
}

func (a *SESActivities) GetCustomVerificationEmailTemplate(ctx context.Context, input *ses.GetCustomVerificationEmailTemplateInput) (*ses.GetCustomVerificationEmailTemplateOutput, error) {
	return a.client.GetCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESActivities) GetIdentityDkimAttributes(ctx context.Context, input *ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error) {
	return a.client.GetIdentityDkimAttributesWithContext(ctx, input)
}

func (a *SESActivities) GetIdentityMailFromDomainAttributes(ctx context.Context, input *ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error) {
	return a.client.GetIdentityMailFromDomainAttributesWithContext(ctx, input)
}

func (a *SESActivities) GetIdentityNotificationAttributes(ctx context.Context, input *ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error) {
	return a.client.GetIdentityNotificationAttributesWithContext(ctx, input)
}

func (a *SESActivities) GetIdentityPolicies(ctx context.Context, input *ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error) {
	return a.client.GetIdentityPoliciesWithContext(ctx, input)
}

func (a *SESActivities) GetIdentityVerificationAttributes(ctx context.Context, input *ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error) {
	return a.client.GetIdentityVerificationAttributesWithContext(ctx, input)
}

func (a *SESActivities) GetSendQuota(ctx context.Context, input *ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error) {
	return a.client.GetSendQuotaWithContext(ctx, input)
}

func (a *SESActivities) GetSendStatistics(ctx context.Context, input *ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error) {
	return a.client.GetSendStatisticsWithContext(ctx, input)
}

func (a *SESActivities) GetTemplate(ctx context.Context, input *ses.GetTemplateInput) (*ses.GetTemplateOutput, error) {
	return a.client.GetTemplateWithContext(ctx, input)
}

func (a *SESActivities) ListConfigurationSets(ctx context.Context, input *ses.ListConfigurationSetsInput) (*ses.ListConfigurationSetsOutput, error) {
	return a.client.ListConfigurationSetsWithContext(ctx, input)
}

func (a *SESActivities) ListCustomVerificationEmailTemplates(ctx context.Context, input *ses.ListCustomVerificationEmailTemplatesInput) (*ses.ListCustomVerificationEmailTemplatesOutput, error) {
	return a.client.ListCustomVerificationEmailTemplatesWithContext(ctx, input)
}

func (a *SESActivities) ListIdentities(ctx context.Context, input *ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error) {
	return a.client.ListIdentitiesWithContext(ctx, input)
}

func (a *SESActivities) ListIdentityPolicies(ctx context.Context, input *ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error) {
	return a.client.ListIdentityPoliciesWithContext(ctx, input)
}

func (a *SESActivities) ListReceiptFilters(ctx context.Context, input *ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error) {
	return a.client.ListReceiptFiltersWithContext(ctx, input)
}

func (a *SESActivities) ListReceiptRuleSets(ctx context.Context, input *ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error) {
	return a.client.ListReceiptRuleSetsWithContext(ctx, input)
}

func (a *SESActivities) ListTemplates(ctx context.Context, input *ses.ListTemplatesInput) (*ses.ListTemplatesOutput, error) {
	return a.client.ListTemplatesWithContext(ctx, input)
}

func (a *SESActivities) ListVerifiedEmailAddresses(ctx context.Context, input *ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error) {
	return a.client.ListVerifiedEmailAddressesWithContext(ctx, input)
}

func (a *SESActivities) PutConfigurationSetDeliveryOptions(ctx context.Context, input *ses.PutConfigurationSetDeliveryOptionsInput) (*ses.PutConfigurationSetDeliveryOptionsOutput, error) {
	return a.client.PutConfigurationSetDeliveryOptionsWithContext(ctx, input)
}

func (a *SESActivities) PutIdentityPolicy(ctx context.Context, input *ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error) {
	return a.client.PutIdentityPolicyWithContext(ctx, input)
}

func (a *SESActivities) ReorderReceiptRuleSet(ctx context.Context, input *ses.ReorderReceiptRuleSetInput) (*ses.ReorderReceiptRuleSetOutput, error) {
	return a.client.ReorderReceiptRuleSetWithContext(ctx, input)
}

func (a *SESActivities) SendBounce(ctx context.Context, input *ses.SendBounceInput) (*ses.SendBounceOutput, error) {
	return a.client.SendBounceWithContext(ctx, input)
}

func (a *SESActivities) SendBulkTemplatedEmail(ctx context.Context, input *ses.SendBulkTemplatedEmailInput) (*ses.SendBulkTemplatedEmailOutput, error) {
	return a.client.SendBulkTemplatedEmailWithContext(ctx, input)
}

func (a *SESActivities) SendCustomVerificationEmail(ctx context.Context, input *ses.SendCustomVerificationEmailInput) (*ses.SendCustomVerificationEmailOutput, error) {
	return a.client.SendCustomVerificationEmailWithContext(ctx, input)
}

func (a *SESActivities) SendEmail(ctx context.Context, input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	return a.client.SendEmailWithContext(ctx, input)
}

func (a *SESActivities) SendRawEmail(ctx context.Context, input *ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error) {
	return a.client.SendRawEmailWithContext(ctx, input)
}

func (a *SESActivities) SendTemplatedEmail(ctx context.Context, input *ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error) {
	return a.client.SendTemplatedEmailWithContext(ctx, input)
}

func (a *SESActivities) SetActiveReceiptRuleSet(ctx context.Context, input *ses.SetActiveReceiptRuleSetInput) (*ses.SetActiveReceiptRuleSetOutput, error) {
	return a.client.SetActiveReceiptRuleSetWithContext(ctx, input)
}

func (a *SESActivities) SetIdentityDkimEnabled(ctx context.Context, input *ses.SetIdentityDkimEnabledInput) (*ses.SetIdentityDkimEnabledOutput, error) {
	return a.client.SetIdentityDkimEnabledWithContext(ctx, input)
}

func (a *SESActivities) SetIdentityFeedbackForwardingEnabled(ctx context.Context, input *ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error) {
	return a.client.SetIdentityFeedbackForwardingEnabledWithContext(ctx, input)
}

func (a *SESActivities) SetIdentityHeadersInNotificationsEnabled(ctx context.Context, input *ses.SetIdentityHeadersInNotificationsEnabledInput) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error) {
	return a.client.SetIdentityHeadersInNotificationsEnabledWithContext(ctx, input)
}

func (a *SESActivities) SetIdentityMailFromDomain(ctx context.Context, input *ses.SetIdentityMailFromDomainInput) (*ses.SetIdentityMailFromDomainOutput, error) {
	return a.client.SetIdentityMailFromDomainWithContext(ctx, input)
}

func (a *SESActivities) SetIdentityNotificationTopic(ctx context.Context, input *ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error) {
	return a.client.SetIdentityNotificationTopicWithContext(ctx, input)
}

func (a *SESActivities) SetReceiptRulePosition(ctx context.Context, input *ses.SetReceiptRulePositionInput) (*ses.SetReceiptRulePositionOutput, error) {
	return a.client.SetReceiptRulePositionWithContext(ctx, input)
}

func (a *SESActivities) TestRenderTemplate(ctx context.Context, input *ses.TestRenderTemplateInput) (*ses.TestRenderTemplateOutput, error) {
	return a.client.TestRenderTemplateWithContext(ctx, input)
}

func (a *SESActivities) UpdateAccountSendingEnabled(ctx context.Context, input *ses.UpdateAccountSendingEnabledInput) (*ses.UpdateAccountSendingEnabledOutput, error) {
	return a.client.UpdateAccountSendingEnabledWithContext(ctx, input)
}

func (a *SESActivities) UpdateConfigurationSetEventDestination(ctx context.Context, input *ses.UpdateConfigurationSetEventDestinationInput) (*ses.UpdateConfigurationSetEventDestinationOutput, error) {
	return a.client.UpdateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESActivities) UpdateConfigurationSetReputationMetricsEnabled(ctx context.Context, input *ses.UpdateConfigurationSetReputationMetricsEnabledInput) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error) {
	return a.client.UpdateConfigurationSetReputationMetricsEnabledWithContext(ctx, input)
}

func (a *SESActivities) UpdateConfigurationSetSendingEnabled(ctx context.Context, input *ses.UpdateConfigurationSetSendingEnabledInput) (*ses.UpdateConfigurationSetSendingEnabledOutput, error) {
	return a.client.UpdateConfigurationSetSendingEnabledWithContext(ctx, input)
}

func (a *SESActivities) UpdateConfigurationSetTrackingOptions(ctx context.Context, input *ses.UpdateConfigurationSetTrackingOptionsInput) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error) {
	return a.client.UpdateConfigurationSetTrackingOptionsWithContext(ctx, input)
}

func (a *SESActivities) UpdateCustomVerificationEmailTemplate(ctx context.Context, input *ses.UpdateCustomVerificationEmailTemplateInput) (*ses.UpdateCustomVerificationEmailTemplateOutput, error) {
	return a.client.UpdateCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESActivities) UpdateReceiptRule(ctx context.Context, input *ses.UpdateReceiptRuleInput) (*ses.UpdateReceiptRuleOutput, error) {
	return a.client.UpdateReceiptRuleWithContext(ctx, input)
}

func (a *SESActivities) UpdateTemplate(ctx context.Context, input *ses.UpdateTemplateInput) (*ses.UpdateTemplateOutput, error) {
	return a.client.UpdateTemplateWithContext(ctx, input)
}

func (a *SESActivities) VerifyDomainDkim(ctx context.Context, input *ses.VerifyDomainDkimInput) (*ses.VerifyDomainDkimOutput, error) {
	return a.client.VerifyDomainDkimWithContext(ctx, input)
}

func (a *SESActivities) VerifyDomainIdentity(ctx context.Context, input *ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error) {
	return a.client.VerifyDomainIdentityWithContext(ctx, input)
}

func (a *SESActivities) VerifyEmailAddress(ctx context.Context, input *ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error) {
	return a.client.VerifyEmailAddressWithContext(ctx, input)
}

func (a *SESActivities) VerifyEmailIdentity(ctx context.Context, input *ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error) {
	return a.client.VerifyEmailIdentityWithContext(ctx, input)
}

func (a *SESActivities) WaitUntilIdentityExists(ctx context.Context, input *ses.GetIdentityVerificationAttributesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilIdentityExistsWithContext(ctx, input, options...)
	})
}
