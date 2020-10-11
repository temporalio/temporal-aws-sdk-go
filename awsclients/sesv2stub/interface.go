// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package sesv2stub

import (
	"github.com/aws/aws-sdk-go/service/sesv2"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateConfigurationSet(ctx workflow.Context, input *sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error)
	CreateConfigurationSetAsync(ctx workflow.Context, input *sesv2.CreateConfigurationSetInput) *SESV2CreateConfigurationSetFuture

	CreateConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) (*sesv2.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) *SESV2CreateConfigurationSetEventDestinationFuture

	CreateCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error)
	CreateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) *SESV2CreateCustomVerificationEmailTemplateFuture

	CreateDedicatedIpPool(ctx workflow.Context, input *sesv2.CreateDedicatedIpPoolInput) (*sesv2.CreateDedicatedIpPoolOutput, error)
	CreateDedicatedIpPoolAsync(ctx workflow.Context, input *sesv2.CreateDedicatedIpPoolInput) *SESV2CreateDedicatedIpPoolFuture

	CreateDeliverabilityTestReport(ctx workflow.Context, input *sesv2.CreateDeliverabilityTestReportInput) (*sesv2.CreateDeliverabilityTestReportOutput, error)
	CreateDeliverabilityTestReportAsync(ctx workflow.Context, input *sesv2.CreateDeliverabilityTestReportInput) *SESV2CreateDeliverabilityTestReportFuture

	CreateEmailIdentity(ctx workflow.Context, input *sesv2.CreateEmailIdentityInput) (*sesv2.CreateEmailIdentityOutput, error)
	CreateEmailIdentityAsync(ctx workflow.Context, input *sesv2.CreateEmailIdentityInput) *SESV2CreateEmailIdentityFuture

	CreateEmailIdentityPolicy(ctx workflow.Context, input *sesv2.CreateEmailIdentityPolicyInput) (*sesv2.CreateEmailIdentityPolicyOutput, error)
	CreateEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.CreateEmailIdentityPolicyInput) *SESV2CreateEmailIdentityPolicyFuture

	CreateEmailTemplate(ctx workflow.Context, input *sesv2.CreateEmailTemplateInput) (*sesv2.CreateEmailTemplateOutput, error)
	CreateEmailTemplateAsync(ctx workflow.Context, input *sesv2.CreateEmailTemplateInput) *SESV2CreateEmailTemplateFuture

	CreateImportJob(ctx workflow.Context, input *sesv2.CreateImportJobInput) (*sesv2.CreateImportJobOutput, error)
	CreateImportJobAsync(ctx workflow.Context, input *sesv2.CreateImportJobInput) *SESV2CreateImportJobFuture

	DeleteConfigurationSet(ctx workflow.Context, input *sesv2.DeleteConfigurationSetInput) (*sesv2.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetAsync(ctx workflow.Context, input *sesv2.DeleteConfigurationSetInput) *SESV2DeleteConfigurationSetFuture

	DeleteConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) *SESV2DeleteConfigurationSetEventDestinationFuture

	DeleteCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) *SESV2DeleteCustomVerificationEmailTemplateFuture

	DeleteDedicatedIpPool(ctx workflow.Context, input *sesv2.DeleteDedicatedIpPoolInput) (*sesv2.DeleteDedicatedIpPoolOutput, error)
	DeleteDedicatedIpPoolAsync(ctx workflow.Context, input *sesv2.DeleteDedicatedIpPoolInput) *SESV2DeleteDedicatedIpPoolFuture

	DeleteEmailIdentity(ctx workflow.Context, input *sesv2.DeleteEmailIdentityInput) (*sesv2.DeleteEmailIdentityOutput, error)
	DeleteEmailIdentityAsync(ctx workflow.Context, input *sesv2.DeleteEmailIdentityInput) *SESV2DeleteEmailIdentityFuture

	DeleteEmailIdentityPolicy(ctx workflow.Context, input *sesv2.DeleteEmailIdentityPolicyInput) (*sesv2.DeleteEmailIdentityPolicyOutput, error)
	DeleteEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.DeleteEmailIdentityPolicyInput) *SESV2DeleteEmailIdentityPolicyFuture

	DeleteEmailTemplate(ctx workflow.Context, input *sesv2.DeleteEmailTemplateInput) (*sesv2.DeleteEmailTemplateOutput, error)
	DeleteEmailTemplateAsync(ctx workflow.Context, input *sesv2.DeleteEmailTemplateInput) *SESV2DeleteEmailTemplateFuture

	DeleteSuppressedDestination(ctx workflow.Context, input *sesv2.DeleteSuppressedDestinationInput) (*sesv2.DeleteSuppressedDestinationOutput, error)
	DeleteSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.DeleteSuppressedDestinationInput) *SESV2DeleteSuppressedDestinationFuture

	GetAccount(ctx workflow.Context, input *sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error)
	GetAccountAsync(ctx workflow.Context, input *sesv2.GetAccountInput) *SESV2GetAccountFuture

	GetBlacklistReports(ctx workflow.Context, input *sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error)
	GetBlacklistReportsAsync(ctx workflow.Context, input *sesv2.GetBlacklistReportsInput) *SESV2GetBlacklistReportsFuture

	GetConfigurationSet(ctx workflow.Context, input *sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error)
	GetConfigurationSetAsync(ctx workflow.Context, input *sesv2.GetConfigurationSetInput) *SESV2GetConfigurationSetFuture

	GetConfigurationSetEventDestinations(ctx workflow.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) *SESV2GetConfigurationSetEventDestinationsFuture

	GetCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error)
	GetCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) *SESV2GetCustomVerificationEmailTemplateFuture

	GetDedicatedIp(ctx workflow.Context, input *sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error)
	GetDedicatedIpAsync(ctx workflow.Context, input *sesv2.GetDedicatedIpInput) *SESV2GetDedicatedIpFuture

	GetDedicatedIps(ctx workflow.Context, input *sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error)
	GetDedicatedIpsAsync(ctx workflow.Context, input *sesv2.GetDedicatedIpsInput) *SESV2GetDedicatedIpsFuture

	GetDeliverabilityDashboardOptions(ctx workflow.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error)
	GetDeliverabilityDashboardOptionsAsync(ctx workflow.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) *SESV2GetDeliverabilityDashboardOptionsFuture

	GetDeliverabilityTestReport(ctx workflow.Context, input *sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error)
	GetDeliverabilityTestReportAsync(ctx workflow.Context, input *sesv2.GetDeliverabilityTestReportInput) *SESV2GetDeliverabilityTestReportFuture

	GetDomainDeliverabilityCampaign(ctx workflow.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error)
	GetDomainDeliverabilityCampaignAsync(ctx workflow.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) *SESV2GetDomainDeliverabilityCampaignFuture

	GetDomainStatisticsReport(ctx workflow.Context, input *sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error)
	GetDomainStatisticsReportAsync(ctx workflow.Context, input *sesv2.GetDomainStatisticsReportInput) *SESV2GetDomainStatisticsReportFuture

	GetEmailIdentity(ctx workflow.Context, input *sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error)
	GetEmailIdentityAsync(ctx workflow.Context, input *sesv2.GetEmailIdentityInput) *SESV2GetEmailIdentityFuture

	GetEmailIdentityPolicies(ctx workflow.Context, input *sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error)
	GetEmailIdentityPoliciesAsync(ctx workflow.Context, input *sesv2.GetEmailIdentityPoliciesInput) *SESV2GetEmailIdentityPoliciesFuture

	GetEmailTemplate(ctx workflow.Context, input *sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error)
	GetEmailTemplateAsync(ctx workflow.Context, input *sesv2.GetEmailTemplateInput) *SESV2GetEmailTemplateFuture

	GetImportJob(ctx workflow.Context, input *sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error)
	GetImportJobAsync(ctx workflow.Context, input *sesv2.GetImportJobInput) *SESV2GetImportJobFuture

	GetSuppressedDestination(ctx workflow.Context, input *sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error)
	GetSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.GetSuppressedDestinationInput) *SESV2GetSuppressedDestinationFuture

	ListConfigurationSets(ctx workflow.Context, input *sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error)
	ListConfigurationSetsAsync(ctx workflow.Context, input *sesv2.ListConfigurationSetsInput) *SESV2ListConfigurationSetsFuture

	ListCustomVerificationEmailTemplates(ctx workflow.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error)
	ListCustomVerificationEmailTemplatesAsync(ctx workflow.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) *SESV2ListCustomVerificationEmailTemplatesFuture

	ListDedicatedIpPools(ctx workflow.Context, input *sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error)
	ListDedicatedIpPoolsAsync(ctx workflow.Context, input *sesv2.ListDedicatedIpPoolsInput) *SESV2ListDedicatedIpPoolsFuture

	ListDeliverabilityTestReports(ctx workflow.Context, input *sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error)
	ListDeliverabilityTestReportsAsync(ctx workflow.Context, input *sesv2.ListDeliverabilityTestReportsInput) *SESV2ListDeliverabilityTestReportsFuture

	ListDomainDeliverabilityCampaigns(ctx workflow.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error)
	ListDomainDeliverabilityCampaignsAsync(ctx workflow.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) *SESV2ListDomainDeliverabilityCampaignsFuture

	ListEmailIdentities(ctx workflow.Context, input *sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error)
	ListEmailIdentitiesAsync(ctx workflow.Context, input *sesv2.ListEmailIdentitiesInput) *SESV2ListEmailIdentitiesFuture

	ListEmailTemplates(ctx workflow.Context, input *sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error)
	ListEmailTemplatesAsync(ctx workflow.Context, input *sesv2.ListEmailTemplatesInput) *SESV2ListEmailTemplatesFuture

	ListImportJobs(ctx workflow.Context, input *sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error)
	ListImportJobsAsync(ctx workflow.Context, input *sesv2.ListImportJobsInput) *SESV2ListImportJobsFuture

	ListSuppressedDestinations(ctx workflow.Context, input *sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error)
	ListSuppressedDestinationsAsync(ctx workflow.Context, input *sesv2.ListSuppressedDestinationsInput) *SESV2ListSuppressedDestinationsFuture

	ListTagsForResource(ctx workflow.Context, input *sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *sesv2.ListTagsForResourceInput) *SESV2ListTagsForResourceFuture

	PutAccountDedicatedIpWarmupAttributes(ctx workflow.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error)
	PutAccountDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) *SESV2PutAccountDedicatedIpWarmupAttributesFuture

	PutAccountDetails(ctx workflow.Context, input *sesv2.PutAccountDetailsInput) (*sesv2.PutAccountDetailsOutput, error)
	PutAccountDetailsAsync(ctx workflow.Context, input *sesv2.PutAccountDetailsInput) *SESV2PutAccountDetailsFuture

	PutAccountSendingAttributes(ctx workflow.Context, input *sesv2.PutAccountSendingAttributesInput) (*sesv2.PutAccountSendingAttributesOutput, error)
	PutAccountSendingAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountSendingAttributesInput) *SESV2PutAccountSendingAttributesFuture

	PutAccountSuppressionAttributes(ctx workflow.Context, input *sesv2.PutAccountSuppressionAttributesInput) (*sesv2.PutAccountSuppressionAttributesOutput, error)
	PutAccountSuppressionAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountSuppressionAttributesInput) *SESV2PutAccountSuppressionAttributesFuture

	PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error)
	PutConfigurationSetDeliveryOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) *SESV2PutConfigurationSetDeliveryOptionsFuture

	PutConfigurationSetReputationOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) (*sesv2.PutConfigurationSetReputationOptionsOutput, error)
	PutConfigurationSetReputationOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) *SESV2PutConfigurationSetReputationOptionsFuture

	PutConfigurationSetSendingOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) (*sesv2.PutConfigurationSetSendingOptionsOutput, error)
	PutConfigurationSetSendingOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) *SESV2PutConfigurationSetSendingOptionsFuture

	PutConfigurationSetSuppressionOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error)
	PutConfigurationSetSuppressionOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) *SESV2PutConfigurationSetSuppressionOptionsFuture

	PutConfigurationSetTrackingOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error)
	PutConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) *SESV2PutConfigurationSetTrackingOptionsFuture

	PutDedicatedIpInPool(ctx workflow.Context, input *sesv2.PutDedicatedIpInPoolInput) (*sesv2.PutDedicatedIpInPoolOutput, error)
	PutDedicatedIpInPoolAsync(ctx workflow.Context, input *sesv2.PutDedicatedIpInPoolInput) *SESV2PutDedicatedIpInPoolFuture

	PutDedicatedIpWarmupAttributes(ctx workflow.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error)
	PutDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) *SESV2PutDedicatedIpWarmupAttributesFuture

	PutDeliverabilityDashboardOption(ctx workflow.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) (*sesv2.PutDeliverabilityDashboardOptionOutput, error)
	PutDeliverabilityDashboardOptionAsync(ctx workflow.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) *SESV2PutDeliverabilityDashboardOptionFuture

	PutEmailIdentityDkimAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) (*sesv2.PutEmailIdentityDkimAttributesOutput, error)
	PutEmailIdentityDkimAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) *SESV2PutEmailIdentityDkimAttributesFuture

	PutEmailIdentityDkimSigningAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error)
	PutEmailIdentityDkimSigningAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) *SESV2PutEmailIdentityDkimSigningAttributesFuture

	PutEmailIdentityFeedbackAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error)
	PutEmailIdentityFeedbackAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) *SESV2PutEmailIdentityFeedbackAttributesFuture

	PutEmailIdentityMailFromAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error)
	PutEmailIdentityMailFromAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) *SESV2PutEmailIdentityMailFromAttributesFuture

	PutSuppressedDestination(ctx workflow.Context, input *sesv2.PutSuppressedDestinationInput) (*sesv2.PutSuppressedDestinationOutput, error)
	PutSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.PutSuppressedDestinationInput) *SESV2PutSuppressedDestinationFuture

	SendBulkEmail(ctx workflow.Context, input *sesv2.SendBulkEmailInput) (*sesv2.SendBulkEmailOutput, error)
	SendBulkEmailAsync(ctx workflow.Context, input *sesv2.SendBulkEmailInput) *SESV2SendBulkEmailFuture

	SendCustomVerificationEmail(ctx workflow.Context, input *sesv2.SendCustomVerificationEmailInput) (*sesv2.SendCustomVerificationEmailOutput, error)
	SendCustomVerificationEmailAsync(ctx workflow.Context, input *sesv2.SendCustomVerificationEmailInput) *SESV2SendCustomVerificationEmailFuture

	SendEmail(ctx workflow.Context, input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error)
	SendEmailAsync(ctx workflow.Context, input *sesv2.SendEmailInput) *SESV2SendEmailFuture

	TagResource(ctx workflow.Context, input *sesv2.TagResourceInput) (*sesv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *sesv2.TagResourceInput) *SESV2TagResourceFuture

	TestRenderEmailTemplate(ctx workflow.Context, input *sesv2.TestRenderEmailTemplateInput) (*sesv2.TestRenderEmailTemplateOutput, error)
	TestRenderEmailTemplateAsync(ctx workflow.Context, input *sesv2.TestRenderEmailTemplateInput) *SESV2TestRenderEmailTemplateFuture

	UntagResource(ctx workflow.Context, input *sesv2.UntagResourceInput) (*sesv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *sesv2.UntagResourceInput) *SESV2UntagResourceFuture

	UpdateConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) *SESV2UpdateConfigurationSetEventDestinationFuture

	UpdateCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) *SESV2UpdateCustomVerificationEmailTemplateFuture

	UpdateEmailIdentityPolicy(ctx workflow.Context, input *sesv2.UpdateEmailIdentityPolicyInput) (*sesv2.UpdateEmailIdentityPolicyOutput, error)
	UpdateEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.UpdateEmailIdentityPolicyInput) *SESV2UpdateEmailIdentityPolicyFuture

	UpdateEmailTemplate(ctx workflow.Context, input *sesv2.UpdateEmailTemplateInput) (*sesv2.UpdateEmailTemplateOutput, error)
	UpdateEmailTemplateAsync(ctx workflow.Context, input *sesv2.UpdateEmailTemplateInput) *SESV2UpdateEmailTemplateFuture
}

func NewClient() Client {
	return &stub{}
}
