package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"go.temporal.io/sdk/workflow"
)

type PinpointEmailClient interface {
    CreateConfigurationSet(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetInput) (*pinpointemail.CreateConfigurationSetOutput, error)
    CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetInput) *PinpointemailCreateConfigurationSetResult

    CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetEventDestinationInput) (*pinpointemail.CreateConfigurationSetEventDestinationOutput, error)
    CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetEventDestinationInput) *PinpointemailCreateConfigurationSetEventDestinationResult

    CreateDedicatedIpPool(ctx workflow.Context, input *pinpointemail.CreateDedicatedIpPoolInput) (*pinpointemail.CreateDedicatedIpPoolOutput, error)
    CreateDedicatedIpPoolAsync(ctx workflow.Context, input *pinpointemail.CreateDedicatedIpPoolInput) *PinpointemailCreateDedicatedIpPoolResult

    CreateDeliverabilityTestReport(ctx workflow.Context, input *pinpointemail.CreateDeliverabilityTestReportInput) (*pinpointemail.CreateDeliverabilityTestReportOutput, error)
    CreateDeliverabilityTestReportAsync(ctx workflow.Context, input *pinpointemail.CreateDeliverabilityTestReportInput) *PinpointemailCreateDeliverabilityTestReportResult

    CreateEmailIdentity(ctx workflow.Context, input *pinpointemail.CreateEmailIdentityInput) (*pinpointemail.CreateEmailIdentityOutput, error)
    CreateEmailIdentityAsync(ctx workflow.Context, input *pinpointemail.CreateEmailIdentityInput) *PinpointemailCreateEmailIdentityResult

    DeleteConfigurationSet(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetInput) (*pinpointemail.DeleteConfigurationSetOutput, error)
    DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetInput) *PinpointemailDeleteConfigurationSetResult

    DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetEventDestinationInput) (*pinpointemail.DeleteConfigurationSetEventDestinationOutput, error)
    DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetEventDestinationInput) *PinpointemailDeleteConfigurationSetEventDestinationResult

    DeleteDedicatedIpPool(ctx workflow.Context, input *pinpointemail.DeleteDedicatedIpPoolInput) (*pinpointemail.DeleteDedicatedIpPoolOutput, error)
    DeleteDedicatedIpPoolAsync(ctx workflow.Context, input *pinpointemail.DeleteDedicatedIpPoolInput) *PinpointemailDeleteDedicatedIpPoolResult

    DeleteEmailIdentity(ctx workflow.Context, input *pinpointemail.DeleteEmailIdentityInput) (*pinpointemail.DeleteEmailIdentityOutput, error)
    DeleteEmailIdentityAsync(ctx workflow.Context, input *pinpointemail.DeleteEmailIdentityInput) *PinpointemailDeleteEmailIdentityResult

    GetAccount(ctx workflow.Context, input *pinpointemail.GetAccountInput) (*pinpointemail.GetAccountOutput, error)
    GetAccountAsync(ctx workflow.Context, input *pinpointemail.GetAccountInput) *PinpointemailGetAccountResult

    GetBlacklistReports(ctx workflow.Context, input *pinpointemail.GetBlacklistReportsInput) (*pinpointemail.GetBlacklistReportsOutput, error)
    GetBlacklistReportsAsync(ctx workflow.Context, input *pinpointemail.GetBlacklistReportsInput) *PinpointemailGetBlacklistReportsResult

    GetConfigurationSet(ctx workflow.Context, input *pinpointemail.GetConfigurationSetInput) (*pinpointemail.GetConfigurationSetOutput, error)
    GetConfigurationSetAsync(ctx workflow.Context, input *pinpointemail.GetConfigurationSetInput) *PinpointemailGetConfigurationSetResult

    GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error)
    GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput) *PinpointemailGetConfigurationSetEventDestinationsResult

    GetDedicatedIp(ctx workflow.Context, input *pinpointemail.GetDedicatedIpInput) (*pinpointemail.GetDedicatedIpOutput, error)
    GetDedicatedIpAsync(ctx workflow.Context, input *pinpointemail.GetDedicatedIpInput) *PinpointemailGetDedicatedIpResult

    GetDedicatedIps(ctx workflow.Context, input *pinpointemail.GetDedicatedIpsInput) (*pinpointemail.GetDedicatedIpsOutput, error)
    GetDedicatedIpsAsync(ctx workflow.Context, input *pinpointemail.GetDedicatedIpsInput) *PinpointemailGetDedicatedIpsResult

    GetDeliverabilityDashboardOptions(ctx workflow.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error)
    GetDeliverabilityDashboardOptionsAsync(ctx workflow.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput) *PinpointemailGetDeliverabilityDashboardOptionsResult

    GetDeliverabilityTestReport(ctx workflow.Context, input *pinpointemail.GetDeliverabilityTestReportInput) (*pinpointemail.GetDeliverabilityTestReportOutput, error)
    GetDeliverabilityTestReportAsync(ctx workflow.Context, input *pinpointemail.GetDeliverabilityTestReportInput) *PinpointemailGetDeliverabilityTestReportResult

    GetDomainDeliverabilityCampaign(ctx workflow.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error)
    GetDomainDeliverabilityCampaignAsync(ctx workflow.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput) *PinpointemailGetDomainDeliverabilityCampaignResult

    GetDomainStatisticsReport(ctx workflow.Context, input *pinpointemail.GetDomainStatisticsReportInput) (*pinpointemail.GetDomainStatisticsReportOutput, error)
    GetDomainStatisticsReportAsync(ctx workflow.Context, input *pinpointemail.GetDomainStatisticsReportInput) *PinpointemailGetDomainStatisticsReportResult

    GetEmailIdentity(ctx workflow.Context, input *pinpointemail.GetEmailIdentityInput) (*pinpointemail.GetEmailIdentityOutput, error)
    GetEmailIdentityAsync(ctx workflow.Context, input *pinpointemail.GetEmailIdentityInput) *PinpointemailGetEmailIdentityResult

    ListConfigurationSets(ctx workflow.Context, input *pinpointemail.ListConfigurationSetsInput) (*pinpointemail.ListConfigurationSetsOutput, error)
    ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointemail.ListConfigurationSetsInput) *PinpointemailListConfigurationSetsResult

    ListDedicatedIpPools(ctx workflow.Context, input *pinpointemail.ListDedicatedIpPoolsInput) (*pinpointemail.ListDedicatedIpPoolsOutput, error)
    ListDedicatedIpPoolsAsync(ctx workflow.Context, input *pinpointemail.ListDedicatedIpPoolsInput) *PinpointemailListDedicatedIpPoolsResult

    ListDeliverabilityTestReports(ctx workflow.Context, input *pinpointemail.ListDeliverabilityTestReportsInput) (*pinpointemail.ListDeliverabilityTestReportsOutput, error)
    ListDeliverabilityTestReportsAsync(ctx workflow.Context, input *pinpointemail.ListDeliverabilityTestReportsInput) *PinpointemailListDeliverabilityTestReportsResult

    ListDomainDeliverabilityCampaigns(ctx workflow.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error)
    ListDomainDeliverabilityCampaignsAsync(ctx workflow.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput) *PinpointemailListDomainDeliverabilityCampaignsResult

    ListEmailIdentities(ctx workflow.Context, input *pinpointemail.ListEmailIdentitiesInput) (*pinpointemail.ListEmailIdentitiesOutput, error)
    ListEmailIdentitiesAsync(ctx workflow.Context, input *pinpointemail.ListEmailIdentitiesInput) *PinpointemailListEmailIdentitiesResult

    ListTagsForResource(ctx workflow.Context, input *pinpointemail.ListTagsForResourceInput) (*pinpointemail.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *pinpointemail.ListTagsForResourceInput) *PinpointemailListTagsForResourceResult

    PutAccountDedicatedIpWarmupAttributes(ctx workflow.Context, input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) (*pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput, error)
    PutAccountDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) *PinpointemailPutAccountDedicatedIpWarmupAttributesResult

    PutAccountSendingAttributes(ctx workflow.Context, input *pinpointemail.PutAccountSendingAttributesInput) (*pinpointemail.PutAccountSendingAttributesOutput, error)
    PutAccountSendingAttributesAsync(ctx workflow.Context, input *pinpointemail.PutAccountSendingAttributesInput) *PinpointemailPutAccountSendingAttributesResult

    PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) (*pinpointemail.PutConfigurationSetDeliveryOptionsOutput, error)
    PutConfigurationSetDeliveryOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) *PinpointemailPutConfigurationSetDeliveryOptionsResult

    PutConfigurationSetReputationOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetReputationOptionsInput) (*pinpointemail.PutConfigurationSetReputationOptionsOutput, error)
    PutConfigurationSetReputationOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetReputationOptionsInput) *PinpointemailPutConfigurationSetReputationOptionsResult

    PutConfigurationSetSendingOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetSendingOptionsInput) (*pinpointemail.PutConfigurationSetSendingOptionsOutput, error)
    PutConfigurationSetSendingOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetSendingOptionsInput) *PinpointemailPutConfigurationSetSendingOptionsResult

    PutConfigurationSetTrackingOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetTrackingOptionsInput) (*pinpointemail.PutConfigurationSetTrackingOptionsOutput, error)
    PutConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetTrackingOptionsInput) *PinpointemailPutConfigurationSetTrackingOptionsResult

    PutDedicatedIpInPool(ctx workflow.Context, input *pinpointemail.PutDedicatedIpInPoolInput) (*pinpointemail.PutDedicatedIpInPoolOutput, error)
    PutDedicatedIpInPoolAsync(ctx workflow.Context, input *pinpointemail.PutDedicatedIpInPoolInput) *PinpointemailPutDedicatedIpInPoolResult

    PutDedicatedIpWarmupAttributes(ctx workflow.Context, input *pinpointemail.PutDedicatedIpWarmupAttributesInput) (*pinpointemail.PutDedicatedIpWarmupAttributesOutput, error)
    PutDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *pinpointemail.PutDedicatedIpWarmupAttributesInput) *PinpointemailPutDedicatedIpWarmupAttributesResult

    PutDeliverabilityDashboardOption(ctx workflow.Context, input *pinpointemail.PutDeliverabilityDashboardOptionInput) (*pinpointemail.PutDeliverabilityDashboardOptionOutput, error)
    PutDeliverabilityDashboardOptionAsync(ctx workflow.Context, input *pinpointemail.PutDeliverabilityDashboardOptionInput) *PinpointemailPutDeliverabilityDashboardOptionResult

    PutEmailIdentityDkimAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityDkimAttributesInput) (*pinpointemail.PutEmailIdentityDkimAttributesOutput, error)
    PutEmailIdentityDkimAttributesAsync(ctx workflow.Context, input *pinpointemail.PutEmailIdentityDkimAttributesInput) *PinpointemailPutEmailIdentityDkimAttributesResult

    PutEmailIdentityFeedbackAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) (*pinpointemail.PutEmailIdentityFeedbackAttributesOutput, error)
    PutEmailIdentityFeedbackAttributesAsync(ctx workflow.Context, input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) *PinpointemailPutEmailIdentityFeedbackAttributesResult

    PutEmailIdentityMailFromAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityMailFromAttributesInput) (*pinpointemail.PutEmailIdentityMailFromAttributesOutput, error)
    PutEmailIdentityMailFromAttributesAsync(ctx workflow.Context, input *pinpointemail.PutEmailIdentityMailFromAttributesInput) *PinpointemailPutEmailIdentityMailFromAttributesResult

    SendEmail(ctx workflow.Context, input *pinpointemail.SendEmailInput) (*pinpointemail.SendEmailOutput, error)
    SendEmailAsync(ctx workflow.Context, input *pinpointemail.SendEmailInput) *PinpointemailSendEmailResult

    TagResource(ctx workflow.Context, input *pinpointemail.TagResourceInput) (*pinpointemail.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *pinpointemail.TagResourceInput) *PinpointemailTagResourceResult

    UntagResource(ctx workflow.Context, input *pinpointemail.UntagResourceInput) (*pinpointemail.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *pinpointemail.UntagResourceInput) *PinpointemailUntagResourceResult

    UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.UpdateConfigurationSetEventDestinationInput) (*pinpointemail.UpdateConfigurationSetEventDestinationOutput, error)
    UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointemail.UpdateConfigurationSetEventDestinationInput) *PinpointemailUpdateConfigurationSetEventDestinationResult
}
type PinpointemailCreateConfigurationSetResult struct {
	Result workflow.Future
}

func (r *PinpointemailCreateConfigurationSetResult) Get(ctx workflow.Context) (*pinpointemail.CreateConfigurationSetOutput, error) {
    var output pinpointemail.CreateConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailCreateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *PinpointemailCreateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*pinpointemail.CreateConfigurationSetEventDestinationOutput, error) {
    var output pinpointemail.CreateConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailCreateDedicatedIpPoolResult struct {
	Result workflow.Future
}

func (r *PinpointemailCreateDedicatedIpPoolResult) Get(ctx workflow.Context) (*pinpointemail.CreateDedicatedIpPoolOutput, error) {
    var output pinpointemail.CreateDedicatedIpPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailCreateDeliverabilityTestReportResult struct {
	Result workflow.Future
}

func (r *PinpointemailCreateDeliverabilityTestReportResult) Get(ctx workflow.Context) (*pinpointemail.CreateDeliverabilityTestReportOutput, error) {
    var output pinpointemail.CreateDeliverabilityTestReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailCreateEmailIdentityResult struct {
	Result workflow.Future
}

func (r *PinpointemailCreateEmailIdentityResult) Get(ctx workflow.Context) (*pinpointemail.CreateEmailIdentityOutput, error) {
    var output pinpointemail.CreateEmailIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailDeleteConfigurationSetResult struct {
	Result workflow.Future
}

func (r *PinpointemailDeleteConfigurationSetResult) Get(ctx workflow.Context) (*pinpointemail.DeleteConfigurationSetOutput, error) {
    var output pinpointemail.DeleteConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailDeleteConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *PinpointemailDeleteConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*pinpointemail.DeleteConfigurationSetEventDestinationOutput, error) {
    var output pinpointemail.DeleteConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailDeleteDedicatedIpPoolResult struct {
	Result workflow.Future
}

func (r *PinpointemailDeleteDedicatedIpPoolResult) Get(ctx workflow.Context) (*pinpointemail.DeleteDedicatedIpPoolOutput, error) {
    var output pinpointemail.DeleteDedicatedIpPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailDeleteEmailIdentityResult struct {
	Result workflow.Future
}

func (r *PinpointemailDeleteEmailIdentityResult) Get(ctx workflow.Context) (*pinpointemail.DeleteEmailIdentityOutput, error) {
    var output pinpointemail.DeleteEmailIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetAccountResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetAccountResult) Get(ctx workflow.Context) (*pinpointemail.GetAccountOutput, error) {
    var output pinpointemail.GetAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetBlacklistReportsResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetBlacklistReportsResult) Get(ctx workflow.Context) (*pinpointemail.GetBlacklistReportsOutput, error) {
    var output pinpointemail.GetBlacklistReportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetConfigurationSetResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetConfigurationSetResult) Get(ctx workflow.Context) (*pinpointemail.GetConfigurationSetOutput, error) {
    var output pinpointemail.GetConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetConfigurationSetEventDestinationsResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetConfigurationSetEventDestinationsResult) Get(ctx workflow.Context) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error) {
    var output pinpointemail.GetConfigurationSetEventDestinationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetDedicatedIpResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetDedicatedIpResult) Get(ctx workflow.Context) (*pinpointemail.GetDedicatedIpOutput, error) {
    var output pinpointemail.GetDedicatedIpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetDedicatedIpsResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetDedicatedIpsResult) Get(ctx workflow.Context) (*pinpointemail.GetDedicatedIpsOutput, error) {
    var output pinpointemail.GetDedicatedIpsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetDeliverabilityDashboardOptionsResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetDeliverabilityDashboardOptionsResult) Get(ctx workflow.Context) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error) {
    var output pinpointemail.GetDeliverabilityDashboardOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetDeliverabilityTestReportResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetDeliverabilityTestReportResult) Get(ctx workflow.Context) (*pinpointemail.GetDeliverabilityTestReportOutput, error) {
    var output pinpointemail.GetDeliverabilityTestReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetDomainDeliverabilityCampaignResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetDomainDeliverabilityCampaignResult) Get(ctx workflow.Context) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error) {
    var output pinpointemail.GetDomainDeliverabilityCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetDomainStatisticsReportResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetDomainStatisticsReportResult) Get(ctx workflow.Context) (*pinpointemail.GetDomainStatisticsReportOutput, error) {
    var output pinpointemail.GetDomainStatisticsReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailGetEmailIdentityResult struct {
	Result workflow.Future
}

func (r *PinpointemailGetEmailIdentityResult) Get(ctx workflow.Context) (*pinpointemail.GetEmailIdentityOutput, error) {
    var output pinpointemail.GetEmailIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailListConfigurationSetsResult struct {
	Result workflow.Future
}

func (r *PinpointemailListConfigurationSetsResult) Get(ctx workflow.Context) (*pinpointemail.ListConfigurationSetsOutput, error) {
    var output pinpointemail.ListConfigurationSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailListDedicatedIpPoolsResult struct {
	Result workflow.Future
}

func (r *PinpointemailListDedicatedIpPoolsResult) Get(ctx workflow.Context) (*pinpointemail.ListDedicatedIpPoolsOutput, error) {
    var output pinpointemail.ListDedicatedIpPoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailListDeliverabilityTestReportsResult struct {
	Result workflow.Future
}

func (r *PinpointemailListDeliverabilityTestReportsResult) Get(ctx workflow.Context) (*pinpointemail.ListDeliverabilityTestReportsOutput, error) {
    var output pinpointemail.ListDeliverabilityTestReportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailListDomainDeliverabilityCampaignsResult struct {
	Result workflow.Future
}

func (r *PinpointemailListDomainDeliverabilityCampaignsResult) Get(ctx workflow.Context) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error) {
    var output pinpointemail.ListDomainDeliverabilityCampaignsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailListEmailIdentitiesResult struct {
	Result workflow.Future
}

func (r *PinpointemailListEmailIdentitiesResult) Get(ctx workflow.Context) (*pinpointemail.ListEmailIdentitiesOutput, error) {
    var output pinpointemail.ListEmailIdentitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *PinpointemailListTagsForResourceResult) Get(ctx workflow.Context) (*pinpointemail.ListTagsForResourceOutput, error) {
    var output pinpointemail.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutAccountDedicatedIpWarmupAttributesResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutAccountDedicatedIpWarmupAttributesResult) Get(ctx workflow.Context) (*pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    var output pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutAccountSendingAttributesResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutAccountSendingAttributesResult) Get(ctx workflow.Context) (*pinpointemail.PutAccountSendingAttributesOutput, error) {
    var output pinpointemail.PutAccountSendingAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutConfigurationSetDeliveryOptionsResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutConfigurationSetDeliveryOptionsResult) Get(ctx workflow.Context) (*pinpointemail.PutConfigurationSetDeliveryOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetDeliveryOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutConfigurationSetReputationOptionsResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutConfigurationSetReputationOptionsResult) Get(ctx workflow.Context) (*pinpointemail.PutConfigurationSetReputationOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetReputationOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutConfigurationSetSendingOptionsResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutConfigurationSetSendingOptionsResult) Get(ctx workflow.Context) (*pinpointemail.PutConfigurationSetSendingOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetSendingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutConfigurationSetTrackingOptionsResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutConfigurationSetTrackingOptionsResult) Get(ctx workflow.Context) (*pinpointemail.PutConfigurationSetTrackingOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetTrackingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutDedicatedIpInPoolResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutDedicatedIpInPoolResult) Get(ctx workflow.Context) (*pinpointemail.PutDedicatedIpInPoolOutput, error) {
    var output pinpointemail.PutDedicatedIpInPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutDedicatedIpWarmupAttributesResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutDedicatedIpWarmupAttributesResult) Get(ctx workflow.Context) (*pinpointemail.PutDedicatedIpWarmupAttributesOutput, error) {
    var output pinpointemail.PutDedicatedIpWarmupAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutDeliverabilityDashboardOptionResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutDeliverabilityDashboardOptionResult) Get(ctx workflow.Context) (*pinpointemail.PutDeliverabilityDashboardOptionOutput, error) {
    var output pinpointemail.PutDeliverabilityDashboardOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutEmailIdentityDkimAttributesResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutEmailIdentityDkimAttributesResult) Get(ctx workflow.Context) (*pinpointemail.PutEmailIdentityDkimAttributesOutput, error) {
    var output pinpointemail.PutEmailIdentityDkimAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutEmailIdentityFeedbackAttributesResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutEmailIdentityFeedbackAttributesResult) Get(ctx workflow.Context) (*pinpointemail.PutEmailIdentityFeedbackAttributesOutput, error) {
    var output pinpointemail.PutEmailIdentityFeedbackAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailPutEmailIdentityMailFromAttributesResult struct {
	Result workflow.Future
}

func (r *PinpointemailPutEmailIdentityMailFromAttributesResult) Get(ctx workflow.Context) (*pinpointemail.PutEmailIdentityMailFromAttributesOutput, error) {
    var output pinpointemail.PutEmailIdentityMailFromAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailSendEmailResult struct {
	Result workflow.Future
}

func (r *PinpointemailSendEmailResult) Get(ctx workflow.Context) (*pinpointemail.SendEmailOutput, error) {
    var output pinpointemail.SendEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailTagResourceResult struct {
	Result workflow.Future
}

func (r *PinpointemailTagResourceResult) Get(ctx workflow.Context) (*pinpointemail.TagResourceOutput, error) {
    var output pinpointemail.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailUntagResourceResult struct {
	Result workflow.Future
}

func (r *PinpointemailUntagResourceResult) Get(ctx workflow.Context) (*pinpointemail.UntagResourceOutput, error) {
    var output pinpointemail.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointemailUpdateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *PinpointemailUpdateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*pinpointemail.UpdateConfigurationSetEventDestinationOutput, error) {
    var output pinpointemail.UpdateConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type PinpointEmailStub struct {
    activities PinpointEmailClient
}

func NewPinpointEmailStub() PinpointEmailClient {
    return &PinpointEmailStub{}
}

func (a *PinpointEmailStub) CreateConfigurationSet(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetInput) (*pinpointemail.CreateConfigurationSetOutput, error) {
    var output pinpointemail.CreateConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetEventDestinationInput) (*pinpointemail.CreateConfigurationSetEventDestinationOutput, error) {
    var output pinpointemail.CreateConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) CreateDedicatedIpPool(ctx workflow.Context, input *pinpointemail.CreateDedicatedIpPoolInput) (*pinpointemail.CreateDedicatedIpPoolOutput, error) {
    var output pinpointemail.CreateDedicatedIpPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDedicatedIpPool, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) CreateDeliverabilityTestReport(ctx workflow.Context, input *pinpointemail.CreateDeliverabilityTestReportInput) (*pinpointemail.CreateDeliverabilityTestReportOutput, error) {
    var output pinpointemail.CreateDeliverabilityTestReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeliverabilityTestReport, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) CreateEmailIdentity(ctx workflow.Context, input *pinpointemail.CreateEmailIdentityInput) (*pinpointemail.CreateEmailIdentityOutput, error) {
    var output pinpointemail.CreateEmailIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEmailIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) DeleteConfigurationSet(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetInput) (*pinpointemail.DeleteConfigurationSetOutput, error) {
    var output pinpointemail.DeleteConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetEventDestinationInput) (*pinpointemail.DeleteConfigurationSetEventDestinationOutput, error) {
    var output pinpointemail.DeleteConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) DeleteDedicatedIpPool(ctx workflow.Context, input *pinpointemail.DeleteDedicatedIpPoolInput) (*pinpointemail.DeleteDedicatedIpPoolOutput, error) {
    var output pinpointemail.DeleteDedicatedIpPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDedicatedIpPool, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) DeleteEmailIdentity(ctx workflow.Context, input *pinpointemail.DeleteEmailIdentityInput) (*pinpointemail.DeleteEmailIdentityOutput, error) {
    var output pinpointemail.DeleteEmailIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetAccount(ctx workflow.Context, input *pinpointemail.GetAccountInput) (*pinpointemail.GetAccountOutput, error) {
    var output pinpointemail.GetAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetBlacklistReports(ctx workflow.Context, input *pinpointemail.GetBlacklistReportsInput) (*pinpointemail.GetBlacklistReportsOutput, error) {
    var output pinpointemail.GetBlacklistReportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBlacklistReports, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetConfigurationSet(ctx workflow.Context, input *pinpointemail.GetConfigurationSetInput) (*pinpointemail.GetConfigurationSetOutput, error) {
    var output pinpointemail.GetConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error) {
    var output pinpointemail.GetConfigurationSetEventDestinationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSetEventDestinations, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetDedicatedIp(ctx workflow.Context, input *pinpointemail.GetDedicatedIpInput) (*pinpointemail.GetDedicatedIpOutput, error) {
    var output pinpointemail.GetDedicatedIpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDedicatedIp, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetDedicatedIps(ctx workflow.Context, input *pinpointemail.GetDedicatedIpsInput) (*pinpointemail.GetDedicatedIpsOutput, error) {
    var output pinpointemail.GetDedicatedIpsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDedicatedIps, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetDeliverabilityDashboardOptions(ctx workflow.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error) {
    var output pinpointemail.GetDeliverabilityDashboardOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeliverabilityDashboardOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetDeliverabilityTestReport(ctx workflow.Context, input *pinpointemail.GetDeliverabilityTestReportInput) (*pinpointemail.GetDeliverabilityTestReportOutput, error) {
    var output pinpointemail.GetDeliverabilityTestReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeliverabilityTestReport, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetDomainDeliverabilityCampaign(ctx workflow.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error) {
    var output pinpointemail.GetDomainDeliverabilityCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainDeliverabilityCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetDomainStatisticsReport(ctx workflow.Context, input *pinpointemail.GetDomainStatisticsReportInput) (*pinpointemail.GetDomainStatisticsReportOutput, error) {
    var output pinpointemail.GetDomainStatisticsReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainStatisticsReport, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) GetEmailIdentity(ctx workflow.Context, input *pinpointemail.GetEmailIdentityInput) (*pinpointemail.GetEmailIdentityOutput, error) {
    var output pinpointemail.GetEmailIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEmailIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) ListConfigurationSets(ctx workflow.Context, input *pinpointemail.ListConfigurationSetsInput) (*pinpointemail.ListConfigurationSetsOutput, error) {
    var output pinpointemail.ListConfigurationSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationSets, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) ListDedicatedIpPools(ctx workflow.Context, input *pinpointemail.ListDedicatedIpPoolsInput) (*pinpointemail.ListDedicatedIpPoolsOutput, error) {
    var output pinpointemail.ListDedicatedIpPoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDedicatedIpPools, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) ListDeliverabilityTestReports(ctx workflow.Context, input *pinpointemail.ListDeliverabilityTestReportsInput) (*pinpointemail.ListDeliverabilityTestReportsOutput, error) {
    var output pinpointemail.ListDeliverabilityTestReportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeliverabilityTestReports, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) ListDomainDeliverabilityCampaigns(ctx workflow.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error) {
    var output pinpointemail.ListDomainDeliverabilityCampaignsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomainDeliverabilityCampaigns, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) ListEmailIdentities(ctx workflow.Context, input *pinpointemail.ListEmailIdentitiesInput) (*pinpointemail.ListEmailIdentitiesOutput, error) {
    var output pinpointemail.ListEmailIdentitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEmailIdentities, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) ListTagsForResource(ctx workflow.Context, input *pinpointemail.ListTagsForResourceInput) (*pinpointemail.ListTagsForResourceOutput, error) {
    var output pinpointemail.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutAccountDedicatedIpWarmupAttributes(ctx workflow.Context, input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) (*pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput, error) {
    var output pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountDedicatedIpWarmupAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutAccountSendingAttributes(ctx workflow.Context, input *pinpointemail.PutAccountSendingAttributesInput) (*pinpointemail.PutAccountSendingAttributesOutput, error) {
    var output pinpointemail.PutAccountSendingAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccountSendingAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) (*pinpointemail.PutConfigurationSetDeliveryOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetDeliveryOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetDeliveryOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutConfigurationSetReputationOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetReputationOptionsInput) (*pinpointemail.PutConfigurationSetReputationOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetReputationOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetReputationOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutConfigurationSetSendingOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetSendingOptionsInput) (*pinpointemail.PutConfigurationSetSendingOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetSendingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetSendingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutConfigurationSetTrackingOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetTrackingOptionsInput) (*pinpointemail.PutConfigurationSetTrackingOptionsOutput, error) {
    var output pinpointemail.PutConfigurationSetTrackingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetTrackingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutDedicatedIpInPool(ctx workflow.Context, input *pinpointemail.PutDedicatedIpInPoolInput) (*pinpointemail.PutDedicatedIpInPoolOutput, error) {
    var output pinpointemail.PutDedicatedIpInPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDedicatedIpInPool, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutDedicatedIpWarmupAttributes(ctx workflow.Context, input *pinpointemail.PutDedicatedIpWarmupAttributesInput) (*pinpointemail.PutDedicatedIpWarmupAttributesOutput, error) {
    var output pinpointemail.PutDedicatedIpWarmupAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDedicatedIpWarmupAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutDeliverabilityDashboardOption(ctx workflow.Context, input *pinpointemail.PutDeliverabilityDashboardOptionInput) (*pinpointemail.PutDeliverabilityDashboardOptionOutput, error) {
    var output pinpointemail.PutDeliverabilityDashboardOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDeliverabilityDashboardOption, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutEmailIdentityDkimAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityDkimAttributesInput) (*pinpointemail.PutEmailIdentityDkimAttributesOutput, error) {
    var output pinpointemail.PutEmailIdentityDkimAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityDkimAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutEmailIdentityFeedbackAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) (*pinpointemail.PutEmailIdentityFeedbackAttributesOutput, error) {
    var output pinpointemail.PutEmailIdentityFeedbackAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityFeedbackAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) PutEmailIdentityMailFromAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityMailFromAttributesInput) (*pinpointemail.PutEmailIdentityMailFromAttributesOutput, error) {
    var output pinpointemail.PutEmailIdentityMailFromAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEmailIdentityMailFromAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) SendEmail(ctx workflow.Context, input *pinpointemail.SendEmailInput) (*pinpointemail.SendEmailOutput, error) {
    var output pinpointemail.SendEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) TagResource(ctx workflow.Context, input *pinpointemail.TagResourceInput) (*pinpointemail.TagResourceOutput, error) {
    var output pinpointemail.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) UntagResource(ctx workflow.Context, input *pinpointemail.UntagResourceInput) (*pinpointemail.UntagResourceOutput, error) {
    var output pinpointemail.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointEmailStub) UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.UpdateConfigurationSetEventDestinationInput) (*pinpointemail.UpdateConfigurationSetEventDestinationOutput, error) {
    var output pinpointemail.UpdateConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}
