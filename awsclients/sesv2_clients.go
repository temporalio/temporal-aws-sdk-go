package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sesv2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SESV2Client interface {
    CreateConfigurationSet(ctx workflow.Context, input *sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error)
    CreateConfigurationSetAsync(ctx workflow.Context, input *sesv2.CreateConfigurationSetInput) *Sesv2CreateConfigurationSetResult

    CreateConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) (*sesv2.CreateConfigurationSetEventDestinationOutput, error)
    CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) *Sesv2CreateConfigurationSetEventDestinationResult

    CreateCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error)
    CreateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) *Sesv2CreateCustomVerificationEmailTemplateResult

    CreateDedicatedIpPool(ctx workflow.Context, input *sesv2.CreateDedicatedIpPoolInput) (*sesv2.CreateDedicatedIpPoolOutput, error)
    CreateDedicatedIpPoolAsync(ctx workflow.Context, input *sesv2.CreateDedicatedIpPoolInput) *Sesv2CreateDedicatedIpPoolResult

    CreateDeliverabilityTestReport(ctx workflow.Context, input *sesv2.CreateDeliverabilityTestReportInput) (*sesv2.CreateDeliverabilityTestReportOutput, error)
    CreateDeliverabilityTestReportAsync(ctx workflow.Context, input *sesv2.CreateDeliverabilityTestReportInput) *Sesv2CreateDeliverabilityTestReportResult

    CreateEmailIdentity(ctx workflow.Context, input *sesv2.CreateEmailIdentityInput) (*sesv2.CreateEmailIdentityOutput, error)
    CreateEmailIdentityAsync(ctx workflow.Context, input *sesv2.CreateEmailIdentityInput) *Sesv2CreateEmailIdentityResult

    CreateEmailIdentityPolicy(ctx workflow.Context, input *sesv2.CreateEmailIdentityPolicyInput) (*sesv2.CreateEmailIdentityPolicyOutput, error)
    CreateEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.CreateEmailIdentityPolicyInput) *Sesv2CreateEmailIdentityPolicyResult

    CreateEmailTemplate(ctx workflow.Context, input *sesv2.CreateEmailTemplateInput) (*sesv2.CreateEmailTemplateOutput, error)
    CreateEmailTemplateAsync(ctx workflow.Context, input *sesv2.CreateEmailTemplateInput) *Sesv2CreateEmailTemplateResult

    CreateImportJob(ctx workflow.Context, input *sesv2.CreateImportJobInput) (*sesv2.CreateImportJobOutput, error)
    CreateImportJobAsync(ctx workflow.Context, input *sesv2.CreateImportJobInput) *Sesv2CreateImportJobResult

    DeleteConfigurationSet(ctx workflow.Context, input *sesv2.DeleteConfigurationSetInput) (*sesv2.DeleteConfigurationSetOutput, error)
    DeleteConfigurationSetAsync(ctx workflow.Context, input *sesv2.DeleteConfigurationSetInput) *Sesv2DeleteConfigurationSetResult

    DeleteConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error)
    DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) *Sesv2DeleteConfigurationSetEventDestinationResult

    DeleteCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error)
    DeleteCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) *Sesv2DeleteCustomVerificationEmailTemplateResult

    DeleteDedicatedIpPool(ctx workflow.Context, input *sesv2.DeleteDedicatedIpPoolInput) (*sesv2.DeleteDedicatedIpPoolOutput, error)
    DeleteDedicatedIpPoolAsync(ctx workflow.Context, input *sesv2.DeleteDedicatedIpPoolInput) *Sesv2DeleteDedicatedIpPoolResult

    DeleteEmailIdentity(ctx workflow.Context, input *sesv2.DeleteEmailIdentityInput) (*sesv2.DeleteEmailIdentityOutput, error)
    DeleteEmailIdentityAsync(ctx workflow.Context, input *sesv2.DeleteEmailIdentityInput) *Sesv2DeleteEmailIdentityResult

    DeleteEmailIdentityPolicy(ctx workflow.Context, input *sesv2.DeleteEmailIdentityPolicyInput) (*sesv2.DeleteEmailIdentityPolicyOutput, error)
    DeleteEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.DeleteEmailIdentityPolicyInput) *Sesv2DeleteEmailIdentityPolicyResult

    DeleteEmailTemplate(ctx workflow.Context, input *sesv2.DeleteEmailTemplateInput) (*sesv2.DeleteEmailTemplateOutput, error)
    DeleteEmailTemplateAsync(ctx workflow.Context, input *sesv2.DeleteEmailTemplateInput) *Sesv2DeleteEmailTemplateResult

    DeleteSuppressedDestination(ctx workflow.Context, input *sesv2.DeleteSuppressedDestinationInput) (*sesv2.DeleteSuppressedDestinationOutput, error)
    DeleteSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.DeleteSuppressedDestinationInput) *Sesv2DeleteSuppressedDestinationResult

    GetAccount(ctx workflow.Context, input *sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error)
    GetAccountAsync(ctx workflow.Context, input *sesv2.GetAccountInput) *Sesv2GetAccountResult

    GetBlacklistReports(ctx workflow.Context, input *sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error)
    GetBlacklistReportsAsync(ctx workflow.Context, input *sesv2.GetBlacklistReportsInput) *Sesv2GetBlacklistReportsResult

    GetConfigurationSet(ctx workflow.Context, input *sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error)
    GetConfigurationSetAsync(ctx workflow.Context, input *sesv2.GetConfigurationSetInput) *Sesv2GetConfigurationSetResult

    GetConfigurationSetEventDestinations(ctx workflow.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error)
    GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) *Sesv2GetConfigurationSetEventDestinationsResult

    GetCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error)
    GetCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) *Sesv2GetCustomVerificationEmailTemplateResult

    GetDedicatedIp(ctx workflow.Context, input *sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error)
    GetDedicatedIpAsync(ctx workflow.Context, input *sesv2.GetDedicatedIpInput) *Sesv2GetDedicatedIpResult

    GetDedicatedIps(ctx workflow.Context, input *sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error)
    GetDedicatedIpsAsync(ctx workflow.Context, input *sesv2.GetDedicatedIpsInput) *Sesv2GetDedicatedIpsResult

    GetDeliverabilityDashboardOptions(ctx workflow.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error)
    GetDeliverabilityDashboardOptionsAsync(ctx workflow.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) *Sesv2GetDeliverabilityDashboardOptionsResult

    GetDeliverabilityTestReport(ctx workflow.Context, input *sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error)
    GetDeliverabilityTestReportAsync(ctx workflow.Context, input *sesv2.GetDeliverabilityTestReportInput) *Sesv2GetDeliverabilityTestReportResult

    GetDomainDeliverabilityCampaign(ctx workflow.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error)
    GetDomainDeliverabilityCampaignAsync(ctx workflow.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) *Sesv2GetDomainDeliverabilityCampaignResult

    GetDomainStatisticsReport(ctx workflow.Context, input *sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error)
    GetDomainStatisticsReportAsync(ctx workflow.Context, input *sesv2.GetDomainStatisticsReportInput) *Sesv2GetDomainStatisticsReportResult

    GetEmailIdentity(ctx workflow.Context, input *sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error)
    GetEmailIdentityAsync(ctx workflow.Context, input *sesv2.GetEmailIdentityInput) *Sesv2GetEmailIdentityResult

    GetEmailIdentityPolicies(ctx workflow.Context, input *sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error)
    GetEmailIdentityPoliciesAsync(ctx workflow.Context, input *sesv2.GetEmailIdentityPoliciesInput) *Sesv2GetEmailIdentityPoliciesResult

    GetEmailTemplate(ctx workflow.Context, input *sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error)
    GetEmailTemplateAsync(ctx workflow.Context, input *sesv2.GetEmailTemplateInput) *Sesv2GetEmailTemplateResult

    GetImportJob(ctx workflow.Context, input *sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error)
    GetImportJobAsync(ctx workflow.Context, input *sesv2.GetImportJobInput) *Sesv2GetImportJobResult

    GetSuppressedDestination(ctx workflow.Context, input *sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error)
    GetSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.GetSuppressedDestinationInput) *Sesv2GetSuppressedDestinationResult

    ListConfigurationSets(ctx workflow.Context, input *sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error)
    ListConfigurationSetsAsync(ctx workflow.Context, input *sesv2.ListConfigurationSetsInput) *Sesv2ListConfigurationSetsResult

    ListCustomVerificationEmailTemplates(ctx workflow.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error)
    ListCustomVerificationEmailTemplatesAsync(ctx workflow.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) *Sesv2ListCustomVerificationEmailTemplatesResult

    ListDedicatedIpPools(ctx workflow.Context, input *sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error)
    ListDedicatedIpPoolsAsync(ctx workflow.Context, input *sesv2.ListDedicatedIpPoolsInput) *Sesv2ListDedicatedIpPoolsResult

    ListDeliverabilityTestReports(ctx workflow.Context, input *sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error)
    ListDeliverabilityTestReportsAsync(ctx workflow.Context, input *sesv2.ListDeliverabilityTestReportsInput) *Sesv2ListDeliverabilityTestReportsResult

    ListDomainDeliverabilityCampaigns(ctx workflow.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error)
    ListDomainDeliverabilityCampaignsAsync(ctx workflow.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) *Sesv2ListDomainDeliverabilityCampaignsResult

    ListEmailIdentities(ctx workflow.Context, input *sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error)
    ListEmailIdentitiesAsync(ctx workflow.Context, input *sesv2.ListEmailIdentitiesInput) *Sesv2ListEmailIdentitiesResult

    ListEmailTemplates(ctx workflow.Context, input *sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error)
    ListEmailTemplatesAsync(ctx workflow.Context, input *sesv2.ListEmailTemplatesInput) *Sesv2ListEmailTemplatesResult

    ListImportJobs(ctx workflow.Context, input *sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error)
    ListImportJobsAsync(ctx workflow.Context, input *sesv2.ListImportJobsInput) *Sesv2ListImportJobsResult

    ListSuppressedDestinations(ctx workflow.Context, input *sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error)
    ListSuppressedDestinationsAsync(ctx workflow.Context, input *sesv2.ListSuppressedDestinationsInput) *Sesv2ListSuppressedDestinationsResult

    ListTagsForResource(ctx workflow.Context, input *sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *sesv2.ListTagsForResourceInput) *Sesv2ListTagsForResourceResult

    PutAccountDedicatedIpWarmupAttributes(ctx workflow.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error)
    PutAccountDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) *Sesv2PutAccountDedicatedIpWarmupAttributesResult

    PutAccountDetails(ctx workflow.Context, input *sesv2.PutAccountDetailsInput) (*sesv2.PutAccountDetailsOutput, error)
    PutAccountDetailsAsync(ctx workflow.Context, input *sesv2.PutAccountDetailsInput) *Sesv2PutAccountDetailsResult

    PutAccountSendingAttributes(ctx workflow.Context, input *sesv2.PutAccountSendingAttributesInput) (*sesv2.PutAccountSendingAttributesOutput, error)
    PutAccountSendingAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountSendingAttributesInput) *Sesv2PutAccountSendingAttributesResult

    PutAccountSuppressionAttributes(ctx workflow.Context, input *sesv2.PutAccountSuppressionAttributesInput) (*sesv2.PutAccountSuppressionAttributesOutput, error)
    PutAccountSuppressionAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountSuppressionAttributesInput) *Sesv2PutAccountSuppressionAttributesResult

    PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error)
    PutConfigurationSetDeliveryOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) *Sesv2PutConfigurationSetDeliveryOptionsResult

    PutConfigurationSetReputationOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) (*sesv2.PutConfigurationSetReputationOptionsOutput, error)
    PutConfigurationSetReputationOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) *Sesv2PutConfigurationSetReputationOptionsResult

    PutConfigurationSetSendingOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) (*sesv2.PutConfigurationSetSendingOptionsOutput, error)
    PutConfigurationSetSendingOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) *Sesv2PutConfigurationSetSendingOptionsResult

    PutConfigurationSetSuppressionOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error)
    PutConfigurationSetSuppressionOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) *Sesv2PutConfigurationSetSuppressionOptionsResult

    PutConfigurationSetTrackingOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error)
    PutConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) *Sesv2PutConfigurationSetTrackingOptionsResult

    PutDedicatedIpInPool(ctx workflow.Context, input *sesv2.PutDedicatedIpInPoolInput) (*sesv2.PutDedicatedIpInPoolOutput, error)
    PutDedicatedIpInPoolAsync(ctx workflow.Context, input *sesv2.PutDedicatedIpInPoolInput) *Sesv2PutDedicatedIpInPoolResult

    PutDedicatedIpWarmupAttributes(ctx workflow.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error)
    PutDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) *Sesv2PutDedicatedIpWarmupAttributesResult

    PutDeliverabilityDashboardOption(ctx workflow.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) (*sesv2.PutDeliverabilityDashboardOptionOutput, error)
    PutDeliverabilityDashboardOptionAsync(ctx workflow.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) *Sesv2PutDeliverabilityDashboardOptionResult

    PutEmailIdentityDkimAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) (*sesv2.PutEmailIdentityDkimAttributesOutput, error)
    PutEmailIdentityDkimAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) *Sesv2PutEmailIdentityDkimAttributesResult

    PutEmailIdentityDkimSigningAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error)
    PutEmailIdentityDkimSigningAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) *Sesv2PutEmailIdentityDkimSigningAttributesResult

    PutEmailIdentityFeedbackAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error)
    PutEmailIdentityFeedbackAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) *Sesv2PutEmailIdentityFeedbackAttributesResult

    PutEmailIdentityMailFromAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error)
    PutEmailIdentityMailFromAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) *Sesv2PutEmailIdentityMailFromAttributesResult

    PutSuppressedDestination(ctx workflow.Context, input *sesv2.PutSuppressedDestinationInput) (*sesv2.PutSuppressedDestinationOutput, error)
    PutSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.PutSuppressedDestinationInput) *Sesv2PutSuppressedDestinationResult

    SendBulkEmail(ctx workflow.Context, input *sesv2.SendBulkEmailInput) (*sesv2.SendBulkEmailOutput, error)
    SendBulkEmailAsync(ctx workflow.Context, input *sesv2.SendBulkEmailInput) *Sesv2SendBulkEmailResult

    SendCustomVerificationEmail(ctx workflow.Context, input *sesv2.SendCustomVerificationEmailInput) (*sesv2.SendCustomVerificationEmailOutput, error)
    SendCustomVerificationEmailAsync(ctx workflow.Context, input *sesv2.SendCustomVerificationEmailInput) *Sesv2SendCustomVerificationEmailResult

    SendEmail(ctx workflow.Context, input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error)
    SendEmailAsync(ctx workflow.Context, input *sesv2.SendEmailInput) *Sesv2SendEmailResult

    TagResource(ctx workflow.Context, input *sesv2.TagResourceInput) (*sesv2.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *sesv2.TagResourceInput) *Sesv2TagResourceResult

    TestRenderEmailTemplate(ctx workflow.Context, input *sesv2.TestRenderEmailTemplateInput) (*sesv2.TestRenderEmailTemplateOutput, error)
    TestRenderEmailTemplateAsync(ctx workflow.Context, input *sesv2.TestRenderEmailTemplateInput) *Sesv2TestRenderEmailTemplateResult

    UntagResource(ctx workflow.Context, input *sesv2.UntagResourceInput) (*sesv2.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *sesv2.UntagResourceInput) *Sesv2UntagResourceResult

    UpdateConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error)
    UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) *Sesv2UpdateConfigurationSetEventDestinationResult

    UpdateCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error)
    UpdateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) *Sesv2UpdateCustomVerificationEmailTemplateResult

    UpdateEmailIdentityPolicy(ctx workflow.Context, input *sesv2.UpdateEmailIdentityPolicyInput) (*sesv2.UpdateEmailIdentityPolicyOutput, error)
    UpdateEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.UpdateEmailIdentityPolicyInput) *Sesv2UpdateEmailIdentityPolicyResult

    UpdateEmailTemplate(ctx workflow.Context, input *sesv2.UpdateEmailTemplateInput) (*sesv2.UpdateEmailTemplateOutput, error)
    UpdateEmailTemplateAsync(ctx workflow.Context, input *sesv2.UpdateEmailTemplateInput) *Sesv2UpdateEmailTemplateResult
}

type Sesv2CreateConfigurationSetResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateConfigurationSetResult) Get(ctx workflow.Context) (*sesv2.CreateConfigurationSetOutput, error) {
    var output sesv2.CreateConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*sesv2.CreateConfigurationSetEventDestinationOutput, error) {
    var output sesv2.CreateConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.CreateCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateDedicatedIpPoolResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateDedicatedIpPoolResult) Get(ctx workflow.Context) (*sesv2.CreateDedicatedIpPoolOutput, error) {
    var output sesv2.CreateDedicatedIpPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateDeliverabilityTestReportResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateDeliverabilityTestReportResult) Get(ctx workflow.Context) (*sesv2.CreateDeliverabilityTestReportOutput, error) {
    var output sesv2.CreateDeliverabilityTestReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateEmailIdentityResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateEmailIdentityResult) Get(ctx workflow.Context) (*sesv2.CreateEmailIdentityOutput, error) {
    var output sesv2.CreateEmailIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateEmailIdentityPolicyResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateEmailIdentityPolicyResult) Get(ctx workflow.Context) (*sesv2.CreateEmailIdentityPolicyOutput, error) {
    var output sesv2.CreateEmailIdentityPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.CreateEmailTemplateOutput, error) {
    var output sesv2.CreateEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2CreateImportJobResult struct {
	Result workflow.Future
}

func (r *Sesv2CreateImportJobResult) Get(ctx workflow.Context) (*sesv2.CreateImportJobOutput, error) {
    var output sesv2.CreateImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteConfigurationSetResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteConfigurationSetResult) Get(ctx workflow.Context) (*sesv2.DeleteConfigurationSetOutput, error) {
    var output sesv2.DeleteConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error) {
    var output sesv2.DeleteConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.DeleteCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteDedicatedIpPoolResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteDedicatedIpPoolResult) Get(ctx workflow.Context) (*sesv2.DeleteDedicatedIpPoolOutput, error) {
    var output sesv2.DeleteDedicatedIpPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteEmailIdentityResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteEmailIdentityResult) Get(ctx workflow.Context) (*sesv2.DeleteEmailIdentityOutput, error) {
    var output sesv2.DeleteEmailIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteEmailIdentityPolicyResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteEmailIdentityPolicyResult) Get(ctx workflow.Context) (*sesv2.DeleteEmailIdentityPolicyOutput, error) {
    var output sesv2.DeleteEmailIdentityPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.DeleteEmailTemplateOutput, error) {
    var output sesv2.DeleteEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2DeleteSuppressedDestinationResult struct {
	Result workflow.Future
}

func (r *Sesv2DeleteSuppressedDestinationResult) Get(ctx workflow.Context) (*sesv2.DeleteSuppressedDestinationOutput, error) {
    var output sesv2.DeleteSuppressedDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetAccountResult struct {
	Result workflow.Future
}

func (r *Sesv2GetAccountResult) Get(ctx workflow.Context) (*sesv2.GetAccountOutput, error) {
    var output sesv2.GetAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetBlacklistReportsResult struct {
	Result workflow.Future
}

func (r *Sesv2GetBlacklistReportsResult) Get(ctx workflow.Context) (*sesv2.GetBlacklistReportsOutput, error) {
    var output sesv2.GetBlacklistReportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetConfigurationSetResult struct {
	Result workflow.Future
}

func (r *Sesv2GetConfigurationSetResult) Get(ctx workflow.Context) (*sesv2.GetConfigurationSetOutput, error) {
    var output sesv2.GetConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetConfigurationSetEventDestinationsResult struct {
	Result workflow.Future
}

func (r *Sesv2GetConfigurationSetEventDestinationsResult) Get(ctx workflow.Context) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
    var output sesv2.GetConfigurationSetEventDestinationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2GetCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.GetCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetDedicatedIpResult struct {
	Result workflow.Future
}

func (r *Sesv2GetDedicatedIpResult) Get(ctx workflow.Context) (*sesv2.GetDedicatedIpOutput, error) {
    var output sesv2.GetDedicatedIpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetDedicatedIpsResult struct {
	Result workflow.Future
}

func (r *Sesv2GetDedicatedIpsResult) Get(ctx workflow.Context) (*sesv2.GetDedicatedIpsOutput, error) {
    var output sesv2.GetDedicatedIpsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetDeliverabilityDashboardOptionsResult struct {
	Result workflow.Future
}

func (r *Sesv2GetDeliverabilityDashboardOptionsResult) Get(ctx workflow.Context) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
    var output sesv2.GetDeliverabilityDashboardOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetDeliverabilityTestReportResult struct {
	Result workflow.Future
}

func (r *Sesv2GetDeliverabilityTestReportResult) Get(ctx workflow.Context) (*sesv2.GetDeliverabilityTestReportOutput, error) {
    var output sesv2.GetDeliverabilityTestReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetDomainDeliverabilityCampaignResult struct {
	Result workflow.Future
}

func (r *Sesv2GetDomainDeliverabilityCampaignResult) Get(ctx workflow.Context) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
    var output sesv2.GetDomainDeliverabilityCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetDomainStatisticsReportResult struct {
	Result workflow.Future
}

func (r *Sesv2GetDomainStatisticsReportResult) Get(ctx workflow.Context) (*sesv2.GetDomainStatisticsReportOutput, error) {
    var output sesv2.GetDomainStatisticsReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetEmailIdentityResult struct {
	Result workflow.Future
}

func (r *Sesv2GetEmailIdentityResult) Get(ctx workflow.Context) (*sesv2.GetEmailIdentityOutput, error) {
    var output sesv2.GetEmailIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetEmailIdentityPoliciesResult struct {
	Result workflow.Future
}

func (r *Sesv2GetEmailIdentityPoliciesResult) Get(ctx workflow.Context) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
    var output sesv2.GetEmailIdentityPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2GetEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.GetEmailTemplateOutput, error) {
    var output sesv2.GetEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetImportJobResult struct {
	Result workflow.Future
}

func (r *Sesv2GetImportJobResult) Get(ctx workflow.Context) (*sesv2.GetImportJobOutput, error) {
    var output sesv2.GetImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2GetSuppressedDestinationResult struct {
	Result workflow.Future
}

func (r *Sesv2GetSuppressedDestinationResult) Get(ctx workflow.Context) (*sesv2.GetSuppressedDestinationOutput, error) {
    var output sesv2.GetSuppressedDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListConfigurationSetsResult struct {
	Result workflow.Future
}

func (r *Sesv2ListConfigurationSetsResult) Get(ctx workflow.Context) (*sesv2.ListConfigurationSetsOutput, error) {
    var output sesv2.ListConfigurationSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListCustomVerificationEmailTemplatesResult struct {
	Result workflow.Future
}

func (r *Sesv2ListCustomVerificationEmailTemplatesResult) Get(ctx workflow.Context) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
    var output sesv2.ListCustomVerificationEmailTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListDedicatedIpPoolsResult struct {
	Result workflow.Future
}

func (r *Sesv2ListDedicatedIpPoolsResult) Get(ctx workflow.Context) (*sesv2.ListDedicatedIpPoolsOutput, error) {
    var output sesv2.ListDedicatedIpPoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListDeliverabilityTestReportsResult struct {
	Result workflow.Future
}

func (r *Sesv2ListDeliverabilityTestReportsResult) Get(ctx workflow.Context) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
    var output sesv2.ListDeliverabilityTestReportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListDomainDeliverabilityCampaignsResult struct {
	Result workflow.Future
}

func (r *Sesv2ListDomainDeliverabilityCampaignsResult) Get(ctx workflow.Context) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
    var output sesv2.ListDomainDeliverabilityCampaignsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListEmailIdentitiesResult struct {
	Result workflow.Future
}

func (r *Sesv2ListEmailIdentitiesResult) Get(ctx workflow.Context) (*sesv2.ListEmailIdentitiesOutput, error) {
    var output sesv2.ListEmailIdentitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListEmailTemplatesResult struct {
	Result workflow.Future
}

func (r *Sesv2ListEmailTemplatesResult) Get(ctx workflow.Context) (*sesv2.ListEmailTemplatesOutput, error) {
    var output sesv2.ListEmailTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListImportJobsResult struct {
	Result workflow.Future
}

func (r *Sesv2ListImportJobsResult) Get(ctx workflow.Context) (*sesv2.ListImportJobsOutput, error) {
    var output sesv2.ListImportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListSuppressedDestinationsResult struct {
	Result workflow.Future
}

func (r *Sesv2ListSuppressedDestinationsResult) Get(ctx workflow.Context) (*sesv2.ListSuppressedDestinationsOutput, error) {
    var output sesv2.ListSuppressedDestinationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2ListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Sesv2ListTagsForResourceResult) Get(ctx workflow.Context) (*sesv2.ListTagsForResourceOutput, error) {
    var output sesv2.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutAccountDedicatedIpWarmupAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutAccountDedicatedIpWarmupAttributesResult) Get(ctx workflow.Context) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    var output sesv2.PutAccountDedicatedIpWarmupAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutAccountDetailsResult struct {
	Result workflow.Future
}

func (r *Sesv2PutAccountDetailsResult) Get(ctx workflow.Context) (*sesv2.PutAccountDetailsOutput, error) {
    var output sesv2.PutAccountDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutAccountSendingAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutAccountSendingAttributesResult) Get(ctx workflow.Context) (*sesv2.PutAccountSendingAttributesOutput, error) {
    var output sesv2.PutAccountSendingAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutAccountSuppressionAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutAccountSuppressionAttributesResult) Get(ctx workflow.Context) (*sesv2.PutAccountSuppressionAttributesOutput, error) {
    var output sesv2.PutAccountSuppressionAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutConfigurationSetDeliveryOptionsResult struct {
	Result workflow.Future
}

func (r *Sesv2PutConfigurationSetDeliveryOptionsResult) Get(ctx workflow.Context) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error) {
    var output sesv2.PutConfigurationSetDeliveryOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutConfigurationSetReputationOptionsResult struct {
	Result workflow.Future
}

func (r *Sesv2PutConfigurationSetReputationOptionsResult) Get(ctx workflow.Context) (*sesv2.PutConfigurationSetReputationOptionsOutput, error) {
    var output sesv2.PutConfigurationSetReputationOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutConfigurationSetSendingOptionsResult struct {
	Result workflow.Future
}

func (r *Sesv2PutConfigurationSetSendingOptionsResult) Get(ctx workflow.Context) (*sesv2.PutConfigurationSetSendingOptionsOutput, error) {
    var output sesv2.PutConfigurationSetSendingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutConfigurationSetSuppressionOptionsResult struct {
	Result workflow.Future
}

func (r *Sesv2PutConfigurationSetSuppressionOptionsResult) Get(ctx workflow.Context) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error) {
    var output sesv2.PutConfigurationSetSuppressionOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutConfigurationSetTrackingOptionsResult struct {
	Result workflow.Future
}

func (r *Sesv2PutConfigurationSetTrackingOptionsResult) Get(ctx workflow.Context) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error) {
    var output sesv2.PutConfigurationSetTrackingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutDedicatedIpInPoolResult struct {
	Result workflow.Future
}

func (r *Sesv2PutDedicatedIpInPoolResult) Get(ctx workflow.Context) (*sesv2.PutDedicatedIpInPoolOutput, error) {
    var output sesv2.PutDedicatedIpInPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutDedicatedIpWarmupAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutDedicatedIpWarmupAttributesResult) Get(ctx workflow.Context) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error) {
    var output sesv2.PutDedicatedIpWarmupAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutDeliverabilityDashboardOptionResult struct {
	Result workflow.Future
}

func (r *Sesv2PutDeliverabilityDashboardOptionResult) Get(ctx workflow.Context) (*sesv2.PutDeliverabilityDashboardOptionOutput, error) {
    var output sesv2.PutDeliverabilityDashboardOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutEmailIdentityDkimAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutEmailIdentityDkimAttributesResult) Get(ctx workflow.Context) (*sesv2.PutEmailIdentityDkimAttributesOutput, error) {
    var output sesv2.PutEmailIdentityDkimAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutEmailIdentityDkimSigningAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutEmailIdentityDkimSigningAttributesResult) Get(ctx workflow.Context) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error) {
    var output sesv2.PutEmailIdentityDkimSigningAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutEmailIdentityFeedbackAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutEmailIdentityFeedbackAttributesResult) Get(ctx workflow.Context) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error) {
    var output sesv2.PutEmailIdentityFeedbackAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutEmailIdentityMailFromAttributesResult struct {
	Result workflow.Future
}

func (r *Sesv2PutEmailIdentityMailFromAttributesResult) Get(ctx workflow.Context) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error) {
    var output sesv2.PutEmailIdentityMailFromAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2PutSuppressedDestinationResult struct {
	Result workflow.Future
}

func (r *Sesv2PutSuppressedDestinationResult) Get(ctx workflow.Context) (*sesv2.PutSuppressedDestinationOutput, error) {
    var output sesv2.PutSuppressedDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2SendBulkEmailResult struct {
	Result workflow.Future
}

func (r *Sesv2SendBulkEmailResult) Get(ctx workflow.Context) (*sesv2.SendBulkEmailOutput, error) {
    var output sesv2.SendBulkEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2SendCustomVerificationEmailResult struct {
	Result workflow.Future
}

func (r *Sesv2SendCustomVerificationEmailResult) Get(ctx workflow.Context) (*sesv2.SendCustomVerificationEmailOutput, error) {
    var output sesv2.SendCustomVerificationEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2SendEmailResult struct {
	Result workflow.Future
}

func (r *Sesv2SendEmailResult) Get(ctx workflow.Context) (*sesv2.SendEmailOutput, error) {
    var output sesv2.SendEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2TagResourceResult struct {
	Result workflow.Future
}

func (r *Sesv2TagResourceResult) Get(ctx workflow.Context) (*sesv2.TagResourceOutput, error) {
    var output sesv2.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2TestRenderEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2TestRenderEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.TestRenderEmailTemplateOutput, error) {
    var output sesv2.TestRenderEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2UntagResourceResult struct {
	Result workflow.Future
}

func (r *Sesv2UntagResourceResult) Get(ctx workflow.Context) (*sesv2.UntagResourceOutput, error) {
    var output sesv2.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2UpdateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *Sesv2UpdateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error) {
    var output sesv2.UpdateConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2UpdateCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2UpdateCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.UpdateCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2UpdateEmailIdentityPolicyResult struct {
	Result workflow.Future
}

func (r *Sesv2UpdateEmailIdentityPolicyResult) Get(ctx workflow.Context) (*sesv2.UpdateEmailIdentityPolicyOutput, error) {
    var output sesv2.UpdateEmailIdentityPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Sesv2UpdateEmailTemplateResult struct {
	Result workflow.Future
}

func (r *Sesv2UpdateEmailTemplateResult) Get(ctx workflow.Context) (*sesv2.UpdateEmailTemplateOutput, error) {
    var output sesv2.UpdateEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SESV2Stub struct {
    activities awsactivities.SESV2Activities
}

func NewSESV2Stub() SESV2Client {
    return &SESV2Stub{}
}

func (a *SESV2Stub) CreateConfigurationSet(ctx workflow.Context, input *sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error) {
    var output sesv2.CreateConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateConfigurationSetAsync(ctx workflow.Context, input *sesv2.CreateConfigurationSetInput) *Sesv2CreateConfigurationSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSet, input)
    return &Sesv2CreateConfigurationSetResult{Result: future}
}

func (a *SESV2Stub) CreateConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) (*sesv2.CreateConfigurationSetEventDestinationOutput, error) {
    var output sesv2.CreateConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) *Sesv2CreateConfigurationSetEventDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSetEventDestination, input)
    return &Sesv2CreateConfigurationSetEventDestinationResult{Result: future}
}

func (a *SESV2Stub) CreateCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.CreateCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) *Sesv2CreateCustomVerificationEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCustomVerificationEmailTemplate, input)
    return &Sesv2CreateCustomVerificationEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) CreateDedicatedIpPool(ctx workflow.Context, input *sesv2.CreateDedicatedIpPoolInput) (*sesv2.CreateDedicatedIpPoolOutput, error) {
    var output sesv2.CreateDedicatedIpPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDedicatedIpPool, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateDedicatedIpPoolAsync(ctx workflow.Context, input *sesv2.CreateDedicatedIpPoolInput) *Sesv2CreateDedicatedIpPoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDedicatedIpPool, input)
    return &Sesv2CreateDedicatedIpPoolResult{Result: future}
}

func (a *SESV2Stub) CreateDeliverabilityTestReport(ctx workflow.Context, input *sesv2.CreateDeliverabilityTestReportInput) (*sesv2.CreateDeliverabilityTestReportOutput, error) {
    var output sesv2.CreateDeliverabilityTestReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeliverabilityTestReport, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateDeliverabilityTestReportAsync(ctx workflow.Context, input *sesv2.CreateDeliverabilityTestReportInput) *Sesv2CreateDeliverabilityTestReportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDeliverabilityTestReport, input)
    return &Sesv2CreateDeliverabilityTestReportResult{Result: future}
}

func (a *SESV2Stub) CreateEmailIdentity(ctx workflow.Context, input *sesv2.CreateEmailIdentityInput) (*sesv2.CreateEmailIdentityOutput, error) {
    var output sesv2.CreateEmailIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEmailIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateEmailIdentityAsync(ctx workflow.Context, input *sesv2.CreateEmailIdentityInput) *Sesv2CreateEmailIdentityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEmailIdentity, input)
    return &Sesv2CreateEmailIdentityResult{Result: future}
}

func (a *SESV2Stub) CreateEmailIdentityPolicy(ctx workflow.Context, input *sesv2.CreateEmailIdentityPolicyInput) (*sesv2.CreateEmailIdentityPolicyOutput, error) {
    var output sesv2.CreateEmailIdentityPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEmailIdentityPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.CreateEmailIdentityPolicyInput) *Sesv2CreateEmailIdentityPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEmailIdentityPolicy, input)
    return &Sesv2CreateEmailIdentityPolicyResult{Result: future}
}

func (a *SESV2Stub) CreateEmailTemplate(ctx workflow.Context, input *sesv2.CreateEmailTemplateInput) (*sesv2.CreateEmailTemplateOutput, error) {
    var output sesv2.CreateEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateEmailTemplateAsync(ctx workflow.Context, input *sesv2.CreateEmailTemplateInput) *Sesv2CreateEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEmailTemplate, input)
    return &Sesv2CreateEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) CreateImportJob(ctx workflow.Context, input *sesv2.CreateImportJobInput) (*sesv2.CreateImportJobOutput, error) {
    var output sesv2.CreateImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) CreateImportJobAsync(ctx workflow.Context, input *sesv2.CreateImportJobInput) *Sesv2CreateImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateImportJob, input)
    return &Sesv2CreateImportJobResult{Result: future}
}

func (a *SESV2Stub) DeleteConfigurationSet(ctx workflow.Context, input *sesv2.DeleteConfigurationSetInput) (*sesv2.DeleteConfigurationSetOutput, error) {
    var output sesv2.DeleteConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteConfigurationSetAsync(ctx workflow.Context, input *sesv2.DeleteConfigurationSetInput) *Sesv2DeleteConfigurationSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSet, input)
    return &Sesv2DeleteConfigurationSetResult{Result: future}
}

func (a *SESV2Stub) DeleteConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error) {
    var output sesv2.DeleteConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) *Sesv2DeleteConfigurationSetEventDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSetEventDestination, input)
    return &Sesv2DeleteConfigurationSetEventDestinationResult{Result: future}
}

func (a *SESV2Stub) DeleteCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.DeleteCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) *Sesv2DeleteCustomVerificationEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomVerificationEmailTemplate, input)
    return &Sesv2DeleteCustomVerificationEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) DeleteDedicatedIpPool(ctx workflow.Context, input *sesv2.DeleteDedicatedIpPoolInput) (*sesv2.DeleteDedicatedIpPoolOutput, error) {
    var output sesv2.DeleteDedicatedIpPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDedicatedIpPool, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteDedicatedIpPoolAsync(ctx workflow.Context, input *sesv2.DeleteDedicatedIpPoolInput) *Sesv2DeleteDedicatedIpPoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDedicatedIpPool, input)
    return &Sesv2DeleteDedicatedIpPoolResult{Result: future}
}

func (a *SESV2Stub) DeleteEmailIdentity(ctx workflow.Context, input *sesv2.DeleteEmailIdentityInput) (*sesv2.DeleteEmailIdentityOutput, error) {
    var output sesv2.DeleteEmailIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteEmailIdentityAsync(ctx workflow.Context, input *sesv2.DeleteEmailIdentityInput) *Sesv2DeleteEmailIdentityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailIdentity, input)
    return &Sesv2DeleteEmailIdentityResult{Result: future}
}

func (a *SESV2Stub) DeleteEmailIdentityPolicy(ctx workflow.Context, input *sesv2.DeleteEmailIdentityPolicyInput) (*sesv2.DeleteEmailIdentityPolicyOutput, error) {
    var output sesv2.DeleteEmailIdentityPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailIdentityPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.DeleteEmailIdentityPolicyInput) *Sesv2DeleteEmailIdentityPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailIdentityPolicy, input)
    return &Sesv2DeleteEmailIdentityPolicyResult{Result: future}
}

func (a *SESV2Stub) DeleteEmailTemplate(ctx workflow.Context, input *sesv2.DeleteEmailTemplateInput) (*sesv2.DeleteEmailTemplateOutput, error) {
    var output sesv2.DeleteEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteEmailTemplateAsync(ctx workflow.Context, input *sesv2.DeleteEmailTemplateInput) *Sesv2DeleteEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailTemplate, input)
    return &Sesv2DeleteEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) DeleteSuppressedDestination(ctx workflow.Context, input *sesv2.DeleteSuppressedDestinationInput) (*sesv2.DeleteSuppressedDestinationOutput, error) {
    var output sesv2.DeleteSuppressedDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSuppressedDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) DeleteSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.DeleteSuppressedDestinationInput) *Sesv2DeleteSuppressedDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSuppressedDestination, input)
    return &Sesv2DeleteSuppressedDestinationResult{Result: future}
}

func (a *SESV2Stub) GetAccount(ctx workflow.Context, input *sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error) {
    var output sesv2.GetAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetAccountAsync(ctx workflow.Context, input *sesv2.GetAccountInput) *Sesv2GetAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccount, input)
    return &Sesv2GetAccountResult{Result: future}
}

func (a *SESV2Stub) GetBlacklistReports(ctx workflow.Context, input *sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error) {
    var output sesv2.GetBlacklistReportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBlacklistReports, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetBlacklistReportsAsync(ctx workflow.Context, input *sesv2.GetBlacklistReportsInput) *Sesv2GetBlacklistReportsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBlacklistReports, input)
    return &Sesv2GetBlacklistReportsResult{Result: future}
}

func (a *SESV2Stub) GetConfigurationSet(ctx workflow.Context, input *sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error) {
    var output sesv2.GetConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetConfigurationSetAsync(ctx workflow.Context, input *sesv2.GetConfigurationSetInput) *Sesv2GetConfigurationSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSet, input)
    return &Sesv2GetConfigurationSetResult{Result: future}
}

func (a *SESV2Stub) GetConfigurationSetEventDestinations(ctx workflow.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
    var output sesv2.GetConfigurationSetEventDestinationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSetEventDestinations, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) *Sesv2GetConfigurationSetEventDestinationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSetEventDestinations, input)
    return &Sesv2GetConfigurationSetEventDestinationsResult{Result: future}
}

func (a *SESV2Stub) GetCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.GetCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) *Sesv2GetCustomVerificationEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCustomVerificationEmailTemplate, input)
    return &Sesv2GetCustomVerificationEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) GetDedicatedIp(ctx workflow.Context, input *sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error) {
    var output sesv2.GetDedicatedIpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDedicatedIp, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetDedicatedIpAsync(ctx workflow.Context, input *sesv2.GetDedicatedIpInput) *Sesv2GetDedicatedIpResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDedicatedIp, input)
    return &Sesv2GetDedicatedIpResult{Result: future}
}

func (a *SESV2Stub) GetDedicatedIps(ctx workflow.Context, input *sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error) {
    var output sesv2.GetDedicatedIpsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDedicatedIps, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetDedicatedIpsAsync(ctx workflow.Context, input *sesv2.GetDedicatedIpsInput) *Sesv2GetDedicatedIpsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDedicatedIps, input)
    return &Sesv2GetDedicatedIpsResult{Result: future}
}

func (a *SESV2Stub) GetDeliverabilityDashboardOptions(ctx workflow.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
    var output sesv2.GetDeliverabilityDashboardOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeliverabilityDashboardOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetDeliverabilityDashboardOptionsAsync(ctx workflow.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) *Sesv2GetDeliverabilityDashboardOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeliverabilityDashboardOptions, input)
    return &Sesv2GetDeliverabilityDashboardOptionsResult{Result: future}
}

func (a *SESV2Stub) GetDeliverabilityTestReport(ctx workflow.Context, input *sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error) {
    var output sesv2.GetDeliverabilityTestReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeliverabilityTestReport, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetDeliverabilityTestReportAsync(ctx workflow.Context, input *sesv2.GetDeliverabilityTestReportInput) *Sesv2GetDeliverabilityTestReportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeliverabilityTestReport, input)
    return &Sesv2GetDeliverabilityTestReportResult{Result: future}
}

func (a *SESV2Stub) GetDomainDeliverabilityCampaign(ctx workflow.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
    var output sesv2.GetDomainDeliverabilityCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainDeliverabilityCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetDomainDeliverabilityCampaignAsync(ctx workflow.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) *Sesv2GetDomainDeliverabilityCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDomainDeliverabilityCampaign, input)
    return &Sesv2GetDomainDeliverabilityCampaignResult{Result: future}
}

func (a *SESV2Stub) GetDomainStatisticsReport(ctx workflow.Context, input *sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error) {
    var output sesv2.GetDomainStatisticsReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainStatisticsReport, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetDomainStatisticsReportAsync(ctx workflow.Context, input *sesv2.GetDomainStatisticsReportInput) *Sesv2GetDomainStatisticsReportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDomainStatisticsReport, input)
    return &Sesv2GetDomainStatisticsReportResult{Result: future}
}

func (a *SESV2Stub) GetEmailIdentity(ctx workflow.Context, input *sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error) {
    var output sesv2.GetEmailIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEmailIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetEmailIdentityAsync(ctx workflow.Context, input *sesv2.GetEmailIdentityInput) *Sesv2GetEmailIdentityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEmailIdentity, input)
    return &Sesv2GetEmailIdentityResult{Result: future}
}

func (a *SESV2Stub) GetEmailIdentityPolicies(ctx workflow.Context, input *sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
    var output sesv2.GetEmailIdentityPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEmailIdentityPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetEmailIdentityPoliciesAsync(ctx workflow.Context, input *sesv2.GetEmailIdentityPoliciesInput) *Sesv2GetEmailIdentityPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEmailIdentityPolicies, input)
    return &Sesv2GetEmailIdentityPoliciesResult{Result: future}
}

func (a *SESV2Stub) GetEmailTemplate(ctx workflow.Context, input *sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error) {
    var output sesv2.GetEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetEmailTemplateAsync(ctx workflow.Context, input *sesv2.GetEmailTemplateInput) *Sesv2GetEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEmailTemplate, input)
    return &Sesv2GetEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) GetImportJob(ctx workflow.Context, input *sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error) {
    var output sesv2.GetImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetImportJobAsync(ctx workflow.Context, input *sesv2.GetImportJobInput) *Sesv2GetImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetImportJob, input)
    return &Sesv2GetImportJobResult{Result: future}
}

func (a *SESV2Stub) GetSuppressedDestination(ctx workflow.Context, input *sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error) {
    var output sesv2.GetSuppressedDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSuppressedDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) GetSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.GetSuppressedDestinationInput) *Sesv2GetSuppressedDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSuppressedDestination, input)
    return &Sesv2GetSuppressedDestinationResult{Result: future}
}

func (a *SESV2Stub) ListConfigurationSets(ctx workflow.Context, input *sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error) {
    var output sesv2.ListConfigurationSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationSets, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListConfigurationSetsAsync(ctx workflow.Context, input *sesv2.ListConfigurationSetsInput) *Sesv2ListConfigurationSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationSets, input)
    return &Sesv2ListConfigurationSetsResult{Result: future}
}

func (a *SESV2Stub) ListCustomVerificationEmailTemplates(ctx workflow.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
    var output sesv2.ListCustomVerificationEmailTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCustomVerificationEmailTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListCustomVerificationEmailTemplatesAsync(ctx workflow.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) *Sesv2ListCustomVerificationEmailTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCustomVerificationEmailTemplates, input)
    return &Sesv2ListCustomVerificationEmailTemplatesResult{Result: future}
}

func (a *SESV2Stub) ListDedicatedIpPools(ctx workflow.Context, input *sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error) {
    var output sesv2.ListDedicatedIpPoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDedicatedIpPools, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListDedicatedIpPoolsAsync(ctx workflow.Context, input *sesv2.ListDedicatedIpPoolsInput) *Sesv2ListDedicatedIpPoolsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDedicatedIpPools, input)
    return &Sesv2ListDedicatedIpPoolsResult{Result: future}
}

func (a *SESV2Stub) ListDeliverabilityTestReports(ctx workflow.Context, input *sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
    var output sesv2.ListDeliverabilityTestReportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeliverabilityTestReports, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListDeliverabilityTestReportsAsync(ctx workflow.Context, input *sesv2.ListDeliverabilityTestReportsInput) *Sesv2ListDeliverabilityTestReportsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeliverabilityTestReports, input)
    return &Sesv2ListDeliverabilityTestReportsResult{Result: future}
}

func (a *SESV2Stub) ListDomainDeliverabilityCampaigns(ctx workflow.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
    var output sesv2.ListDomainDeliverabilityCampaignsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomainDeliverabilityCampaigns, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListDomainDeliverabilityCampaignsAsync(ctx workflow.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) *Sesv2ListDomainDeliverabilityCampaignsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDomainDeliverabilityCampaigns, input)
    return &Sesv2ListDomainDeliverabilityCampaignsResult{Result: future}
}

func (a *SESV2Stub) ListEmailIdentities(ctx workflow.Context, input *sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error) {
    var output sesv2.ListEmailIdentitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEmailIdentities, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListEmailIdentitiesAsync(ctx workflow.Context, input *sesv2.ListEmailIdentitiesInput) *Sesv2ListEmailIdentitiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEmailIdentities, input)
    return &Sesv2ListEmailIdentitiesResult{Result: future}
}

func (a *SESV2Stub) ListEmailTemplates(ctx workflow.Context, input *sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error) {
    var output sesv2.ListEmailTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEmailTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListEmailTemplatesAsync(ctx workflow.Context, input *sesv2.ListEmailTemplatesInput) *Sesv2ListEmailTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEmailTemplates, input)
    return &Sesv2ListEmailTemplatesResult{Result: future}
}

func (a *SESV2Stub) ListImportJobs(ctx workflow.Context, input *sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error) {
    var output sesv2.ListImportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListImportJobsAsync(ctx workflow.Context, input *sesv2.ListImportJobsInput) *Sesv2ListImportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListImportJobs, input)
    return &Sesv2ListImportJobsResult{Result: future}
}

func (a *SESV2Stub) ListSuppressedDestinations(ctx workflow.Context, input *sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error) {
    var output sesv2.ListSuppressedDestinationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSuppressedDestinations, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListSuppressedDestinationsAsync(ctx workflow.Context, input *sesv2.ListSuppressedDestinationsInput) *Sesv2ListSuppressedDestinationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSuppressedDestinations, input)
    return &Sesv2ListSuppressedDestinationsResult{Result: future}
}

func (a *SESV2Stub) ListTagsForResource(ctx workflow.Context, input *sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error) {
    var output sesv2.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) ListTagsForResourceAsync(ctx workflow.Context, input *sesv2.ListTagsForResourceInput) *Sesv2ListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Sesv2ListTagsForResourceResult{Result: future}
}

func (a *SESV2Stub) PutAccountDedicatedIpWarmupAttributes(ctx workflow.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    var output sesv2.PutAccountDedicatedIpWarmupAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountDedicatedIpWarmupAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutAccountDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) *Sesv2PutAccountDedicatedIpWarmupAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAccountDedicatedIpWarmupAttributes, input)
    return &Sesv2PutAccountDedicatedIpWarmupAttributesResult{Result: future}
}

func (a *SESV2Stub) PutAccountDetails(ctx workflow.Context, input *sesv2.PutAccountDetailsInput) (*sesv2.PutAccountDetailsOutput, error) {
    var output sesv2.PutAccountDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutAccountDetailsAsync(ctx workflow.Context, input *sesv2.PutAccountDetailsInput) *Sesv2PutAccountDetailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAccountDetails, input)
    return &Sesv2PutAccountDetailsResult{Result: future}
}

func (a *SESV2Stub) PutAccountSendingAttributes(ctx workflow.Context, input *sesv2.PutAccountSendingAttributesInput) (*sesv2.PutAccountSendingAttributesOutput, error) {
    var output sesv2.PutAccountSendingAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountSendingAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutAccountSendingAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountSendingAttributesInput) *Sesv2PutAccountSendingAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAccountSendingAttributes, input)
    return &Sesv2PutAccountSendingAttributesResult{Result: future}
}

func (a *SESV2Stub) PutAccountSuppressionAttributes(ctx workflow.Context, input *sesv2.PutAccountSuppressionAttributesInput) (*sesv2.PutAccountSuppressionAttributesOutput, error) {
    var output sesv2.PutAccountSuppressionAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountSuppressionAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutAccountSuppressionAttributesAsync(ctx workflow.Context, input *sesv2.PutAccountSuppressionAttributesInput) *Sesv2PutAccountSuppressionAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAccountSuppressionAttributes, input)
    return &Sesv2PutAccountSuppressionAttributesResult{Result: future}
}

func (a *SESV2Stub) PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error) {
    var output sesv2.PutConfigurationSetDeliveryOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetDeliveryOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutConfigurationSetDeliveryOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) *Sesv2PutConfigurationSetDeliveryOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetDeliveryOptions, input)
    return &Sesv2PutConfigurationSetDeliveryOptionsResult{Result: future}
}

func (a *SESV2Stub) PutConfigurationSetReputationOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) (*sesv2.PutConfigurationSetReputationOptionsOutput, error) {
    var output sesv2.PutConfigurationSetReputationOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetReputationOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutConfigurationSetReputationOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) *Sesv2PutConfigurationSetReputationOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetReputationOptions, input)
    return &Sesv2PutConfigurationSetReputationOptionsResult{Result: future}
}

func (a *SESV2Stub) PutConfigurationSetSendingOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) (*sesv2.PutConfigurationSetSendingOptionsOutput, error) {
    var output sesv2.PutConfigurationSetSendingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetSendingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutConfigurationSetSendingOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) *Sesv2PutConfigurationSetSendingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetSendingOptions, input)
    return &Sesv2PutConfigurationSetSendingOptionsResult{Result: future}
}

func (a *SESV2Stub) PutConfigurationSetSuppressionOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error) {
    var output sesv2.PutConfigurationSetSuppressionOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetSuppressionOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutConfigurationSetSuppressionOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) *Sesv2PutConfigurationSetSuppressionOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetSuppressionOptions, input)
    return &Sesv2PutConfigurationSetSuppressionOptionsResult{Result: future}
}

func (a *SESV2Stub) PutConfigurationSetTrackingOptions(ctx workflow.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error) {
    var output sesv2.PutConfigurationSetTrackingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetTrackingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) *Sesv2PutConfigurationSetTrackingOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetTrackingOptions, input)
    return &Sesv2PutConfigurationSetTrackingOptionsResult{Result: future}
}

func (a *SESV2Stub) PutDedicatedIpInPool(ctx workflow.Context, input *sesv2.PutDedicatedIpInPoolInput) (*sesv2.PutDedicatedIpInPoolOutput, error) {
    var output sesv2.PutDedicatedIpInPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDedicatedIpInPool, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutDedicatedIpInPoolAsync(ctx workflow.Context, input *sesv2.PutDedicatedIpInPoolInput) *Sesv2PutDedicatedIpInPoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutDedicatedIpInPool, input)
    return &Sesv2PutDedicatedIpInPoolResult{Result: future}
}

func (a *SESV2Stub) PutDedicatedIpWarmupAttributes(ctx workflow.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error) {
    var output sesv2.PutDedicatedIpWarmupAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDedicatedIpWarmupAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) *Sesv2PutDedicatedIpWarmupAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutDedicatedIpWarmupAttributes, input)
    return &Sesv2PutDedicatedIpWarmupAttributesResult{Result: future}
}

func (a *SESV2Stub) PutDeliverabilityDashboardOption(ctx workflow.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) (*sesv2.PutDeliverabilityDashboardOptionOutput, error) {
    var output sesv2.PutDeliverabilityDashboardOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDeliverabilityDashboardOption, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutDeliverabilityDashboardOptionAsync(ctx workflow.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) *Sesv2PutDeliverabilityDashboardOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutDeliverabilityDashboardOption, input)
    return &Sesv2PutDeliverabilityDashboardOptionResult{Result: future}
}

func (a *SESV2Stub) PutEmailIdentityDkimAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) (*sesv2.PutEmailIdentityDkimAttributesOutput, error) {
    var output sesv2.PutEmailIdentityDkimAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityDkimAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutEmailIdentityDkimAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) *Sesv2PutEmailIdentityDkimAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityDkimAttributes, input)
    return &Sesv2PutEmailIdentityDkimAttributesResult{Result: future}
}

func (a *SESV2Stub) PutEmailIdentityDkimSigningAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error) {
    var output sesv2.PutEmailIdentityDkimSigningAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityDkimSigningAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutEmailIdentityDkimSigningAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) *Sesv2PutEmailIdentityDkimSigningAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityDkimSigningAttributes, input)
    return &Sesv2PutEmailIdentityDkimSigningAttributesResult{Result: future}
}

func (a *SESV2Stub) PutEmailIdentityFeedbackAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error) {
    var output sesv2.PutEmailIdentityFeedbackAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityFeedbackAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutEmailIdentityFeedbackAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) *Sesv2PutEmailIdentityFeedbackAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityFeedbackAttributes, input)
    return &Sesv2PutEmailIdentityFeedbackAttributesResult{Result: future}
}

func (a *SESV2Stub) PutEmailIdentityMailFromAttributes(ctx workflow.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error) {
    var output sesv2.PutEmailIdentityMailFromAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityMailFromAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutEmailIdentityMailFromAttributesAsync(ctx workflow.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) *Sesv2PutEmailIdentityMailFromAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityMailFromAttributes, input)
    return &Sesv2PutEmailIdentityMailFromAttributesResult{Result: future}
}

func (a *SESV2Stub) PutSuppressedDestination(ctx workflow.Context, input *sesv2.PutSuppressedDestinationInput) (*sesv2.PutSuppressedDestinationOutput, error) {
    var output sesv2.PutSuppressedDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSuppressedDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) PutSuppressedDestinationAsync(ctx workflow.Context, input *sesv2.PutSuppressedDestinationInput) *Sesv2PutSuppressedDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutSuppressedDestination, input)
    return &Sesv2PutSuppressedDestinationResult{Result: future}
}

func (a *SESV2Stub) SendBulkEmail(ctx workflow.Context, input *sesv2.SendBulkEmailInput) (*sesv2.SendBulkEmailOutput, error) {
    var output sesv2.SendBulkEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendBulkEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) SendBulkEmailAsync(ctx workflow.Context, input *sesv2.SendBulkEmailInput) *Sesv2SendBulkEmailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendBulkEmail, input)
    return &Sesv2SendBulkEmailResult{Result: future}
}

func (a *SESV2Stub) SendCustomVerificationEmail(ctx workflow.Context, input *sesv2.SendCustomVerificationEmailInput) (*sesv2.SendCustomVerificationEmailOutput, error) {
    var output sesv2.SendCustomVerificationEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendCustomVerificationEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) SendCustomVerificationEmailAsync(ctx workflow.Context, input *sesv2.SendCustomVerificationEmailInput) *Sesv2SendCustomVerificationEmailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendCustomVerificationEmail, input)
    return &Sesv2SendCustomVerificationEmailResult{Result: future}
}

func (a *SESV2Stub) SendEmail(ctx workflow.Context, input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error) {
    var output sesv2.SendEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) SendEmailAsync(ctx workflow.Context, input *sesv2.SendEmailInput) *Sesv2SendEmailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendEmail, input)
    return &Sesv2SendEmailResult{Result: future}
}

func (a *SESV2Stub) TagResource(ctx workflow.Context, input *sesv2.TagResourceInput) (*sesv2.TagResourceOutput, error) {
    var output sesv2.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) TagResourceAsync(ctx workflow.Context, input *sesv2.TagResourceInput) *Sesv2TagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Sesv2TagResourceResult{Result: future}
}

func (a *SESV2Stub) TestRenderEmailTemplate(ctx workflow.Context, input *sesv2.TestRenderEmailTemplateInput) (*sesv2.TestRenderEmailTemplateOutput, error) {
    var output sesv2.TestRenderEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestRenderEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) TestRenderEmailTemplateAsync(ctx workflow.Context, input *sesv2.TestRenderEmailTemplateInput) *Sesv2TestRenderEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestRenderEmailTemplate, input)
    return &Sesv2TestRenderEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) UntagResource(ctx workflow.Context, input *sesv2.UntagResourceInput) (*sesv2.UntagResourceOutput, error) {
    var output sesv2.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) UntagResourceAsync(ctx workflow.Context, input *sesv2.UntagResourceInput) *Sesv2UntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Sesv2UntagResourceResult{Result: future}
}

func (a *SESV2Stub) UpdateConfigurationSetEventDestination(ctx workflow.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error) {
    var output sesv2.UpdateConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) *Sesv2UpdateConfigurationSetEventDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetEventDestination, input)
    return &Sesv2UpdateConfigurationSetEventDestinationResult{Result: future}
}

func (a *SESV2Stub) UpdateCustomVerificationEmailTemplate(ctx workflow.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error) {
    var output sesv2.UpdateCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) UpdateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) *Sesv2UpdateCustomVerificationEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCustomVerificationEmailTemplate, input)
    return &Sesv2UpdateCustomVerificationEmailTemplateResult{Result: future}
}

func (a *SESV2Stub) UpdateEmailIdentityPolicy(ctx workflow.Context, input *sesv2.UpdateEmailIdentityPolicyInput) (*sesv2.UpdateEmailIdentityPolicyOutput, error) {
    var output sesv2.UpdateEmailIdentityPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailIdentityPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) UpdateEmailIdentityPolicyAsync(ctx workflow.Context, input *sesv2.UpdateEmailIdentityPolicyInput) *Sesv2UpdateEmailIdentityPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailIdentityPolicy, input)
    return &Sesv2UpdateEmailIdentityPolicyResult{Result: future}
}

func (a *SESV2Stub) UpdateEmailTemplate(ctx workflow.Context, input *sesv2.UpdateEmailTemplateInput) (*sesv2.UpdateEmailTemplateOutput, error) {
    var output sesv2.UpdateEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESV2Stub) UpdateEmailTemplateAsync(ctx workflow.Context, input *sesv2.UpdateEmailTemplateInput) *Sesv2UpdateEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailTemplate, input)
    return &Sesv2UpdateEmailTemplateResult{Result: future}
}
