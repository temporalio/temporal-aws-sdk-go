package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/pinpoint/pinpointiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type PinpointActivities struct {
	client pinpointiface.PinpointAPI
}

func NewPinpointActivities(session *session.Session, config ...*aws.Config) *PinpointActivities {
	client := pinpoint.New(session, config...)
	return &PinpointActivities{client: client}
}

func (a *PinpointActivities) CreateApp(ctx context.Context, input *pinpoint.CreateAppInput) (*pinpoint.CreateAppOutput, error) {
	return a.client.CreateAppWithContext(ctx, input)
}

func (a *PinpointActivities) CreateCampaign(ctx context.Context, input *pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error) {
	return a.client.CreateCampaignWithContext(ctx, input)
}

func (a *PinpointActivities) CreateEmailTemplate(ctx context.Context, input *pinpoint.CreateEmailTemplateInput) (*pinpoint.CreateEmailTemplateOutput, error) {
	return a.client.CreateEmailTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) CreateExportJob(ctx context.Context, input *pinpoint.CreateExportJobInput) (*pinpoint.CreateExportJobOutput, error) {
	return a.client.CreateExportJobWithContext(ctx, input)
}

func (a *PinpointActivities) CreateImportJob(ctx context.Context, input *pinpoint.CreateImportJobInput) (*pinpoint.CreateImportJobOutput, error) {
	return a.client.CreateImportJobWithContext(ctx, input)
}

func (a *PinpointActivities) CreateJourney(ctx context.Context, input *pinpoint.CreateJourneyInput) (*pinpoint.CreateJourneyOutput, error) {
	return a.client.CreateJourneyWithContext(ctx, input)
}

func (a *PinpointActivities) CreatePushTemplate(ctx context.Context, input *pinpoint.CreatePushTemplateInput) (*pinpoint.CreatePushTemplateOutput, error) {
	return a.client.CreatePushTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) CreateRecommenderConfiguration(ctx context.Context, input *pinpoint.CreateRecommenderConfigurationInput) (*pinpoint.CreateRecommenderConfigurationOutput, error) {
	return a.client.CreateRecommenderConfigurationWithContext(ctx, input)
}

func (a *PinpointActivities) CreateSegment(ctx context.Context, input *pinpoint.CreateSegmentInput) (*pinpoint.CreateSegmentOutput, error) {
	return a.client.CreateSegmentWithContext(ctx, input)
}

func (a *PinpointActivities) CreateSmsTemplate(ctx context.Context, input *pinpoint.CreateSmsTemplateInput) (*pinpoint.CreateSmsTemplateOutput, error) {
	return a.client.CreateSmsTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) CreateVoiceTemplate(ctx context.Context, input *pinpoint.CreateVoiceTemplateInput) (*pinpoint.CreateVoiceTemplateOutput, error) {
	return a.client.CreateVoiceTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteAdmChannel(ctx context.Context, input *pinpoint.DeleteAdmChannelInput) (*pinpoint.DeleteAdmChannelOutput, error) {
	return a.client.DeleteAdmChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteApnsChannel(ctx context.Context, input *pinpoint.DeleteApnsChannelInput) (*pinpoint.DeleteApnsChannelOutput, error) {
	return a.client.DeleteApnsChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteApnsSandboxChannel(ctx context.Context, input *pinpoint.DeleteApnsSandboxChannelInput) (*pinpoint.DeleteApnsSandboxChannelOutput, error) {
	return a.client.DeleteApnsSandboxChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteApnsVoipChannel(ctx context.Context, input *pinpoint.DeleteApnsVoipChannelInput) (*pinpoint.DeleteApnsVoipChannelOutput, error) {
	return a.client.DeleteApnsVoipChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteApnsVoipSandboxChannel(ctx context.Context, input *pinpoint.DeleteApnsVoipSandboxChannelInput) (*pinpoint.DeleteApnsVoipSandboxChannelOutput, error) {
	return a.client.DeleteApnsVoipSandboxChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteApp(ctx context.Context, input *pinpoint.DeleteAppInput) (*pinpoint.DeleteAppOutput, error) {
	return a.client.DeleteAppWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteBaiduChannel(ctx context.Context, input *pinpoint.DeleteBaiduChannelInput) (*pinpoint.DeleteBaiduChannelOutput, error) {
	return a.client.DeleteBaiduChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteCampaign(ctx context.Context, input *pinpoint.DeleteCampaignInput) (*pinpoint.DeleteCampaignOutput, error) {
	return a.client.DeleteCampaignWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteEmailChannel(ctx context.Context, input *pinpoint.DeleteEmailChannelInput) (*pinpoint.DeleteEmailChannelOutput, error) {
	return a.client.DeleteEmailChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteEmailTemplate(ctx context.Context, input *pinpoint.DeleteEmailTemplateInput) (*pinpoint.DeleteEmailTemplateOutput, error) {
	return a.client.DeleteEmailTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteEndpoint(ctx context.Context, input *pinpoint.DeleteEndpointInput) (*pinpoint.DeleteEndpointOutput, error) {
	return a.client.DeleteEndpointWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteEventStream(ctx context.Context, input *pinpoint.DeleteEventStreamInput) (*pinpoint.DeleteEventStreamOutput, error) {
	return a.client.DeleteEventStreamWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteGcmChannel(ctx context.Context, input *pinpoint.DeleteGcmChannelInput) (*pinpoint.DeleteGcmChannelOutput, error) {
	return a.client.DeleteGcmChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteJourney(ctx context.Context, input *pinpoint.DeleteJourneyInput) (*pinpoint.DeleteJourneyOutput, error) {
	return a.client.DeleteJourneyWithContext(ctx, input)
}

func (a *PinpointActivities) DeletePushTemplate(ctx context.Context, input *pinpoint.DeletePushTemplateInput) (*pinpoint.DeletePushTemplateOutput, error) {
	return a.client.DeletePushTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteRecommenderConfiguration(ctx context.Context, input *pinpoint.DeleteRecommenderConfigurationInput) (*pinpoint.DeleteRecommenderConfigurationOutput, error) {
	return a.client.DeleteRecommenderConfigurationWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteSegment(ctx context.Context, input *pinpoint.DeleteSegmentInput) (*pinpoint.DeleteSegmentOutput, error) {
	return a.client.DeleteSegmentWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteSmsChannel(ctx context.Context, input *pinpoint.DeleteSmsChannelInput) (*pinpoint.DeleteSmsChannelOutput, error) {
	return a.client.DeleteSmsChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteSmsTemplate(ctx context.Context, input *pinpoint.DeleteSmsTemplateInput) (*pinpoint.DeleteSmsTemplateOutput, error) {
	return a.client.DeleteSmsTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteUserEndpoints(ctx context.Context, input *pinpoint.DeleteUserEndpointsInput) (*pinpoint.DeleteUserEndpointsOutput, error) {
	return a.client.DeleteUserEndpointsWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteVoiceChannel(ctx context.Context, input *pinpoint.DeleteVoiceChannelInput) (*pinpoint.DeleteVoiceChannelOutput, error) {
	return a.client.DeleteVoiceChannelWithContext(ctx, input)
}

func (a *PinpointActivities) DeleteVoiceTemplate(ctx context.Context, input *pinpoint.DeleteVoiceTemplateInput) (*pinpoint.DeleteVoiceTemplateOutput, error) {
	return a.client.DeleteVoiceTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) GetAdmChannel(ctx context.Context, input *pinpoint.GetAdmChannelInput) (*pinpoint.GetAdmChannelOutput, error) {
	return a.client.GetAdmChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetApnsChannel(ctx context.Context, input *pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error) {
	return a.client.GetApnsChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetApnsSandboxChannel(ctx context.Context, input *pinpoint.GetApnsSandboxChannelInput) (*pinpoint.GetApnsSandboxChannelOutput, error) {
	return a.client.GetApnsSandboxChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetApnsVoipChannel(ctx context.Context, input *pinpoint.GetApnsVoipChannelInput) (*pinpoint.GetApnsVoipChannelOutput, error) {
	return a.client.GetApnsVoipChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetApnsVoipSandboxChannel(ctx context.Context, input *pinpoint.GetApnsVoipSandboxChannelInput) (*pinpoint.GetApnsVoipSandboxChannelOutput, error) {
	return a.client.GetApnsVoipSandboxChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetApp(ctx context.Context, input *pinpoint.GetAppInput) (*pinpoint.GetAppOutput, error) {
	return a.client.GetAppWithContext(ctx, input)
}

func (a *PinpointActivities) GetApplicationDateRangeKpi(ctx context.Context, input *pinpoint.GetApplicationDateRangeKpiInput) (*pinpoint.GetApplicationDateRangeKpiOutput, error) {
	return a.client.GetApplicationDateRangeKpiWithContext(ctx, input)
}

func (a *PinpointActivities) GetApplicationSettings(ctx context.Context, input *pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error) {
	return a.client.GetApplicationSettingsWithContext(ctx, input)
}

func (a *PinpointActivities) GetApps(ctx context.Context, input *pinpoint.GetAppsInput) (*pinpoint.GetAppsOutput, error) {
	return a.client.GetAppsWithContext(ctx, input)
}

func (a *PinpointActivities) GetBaiduChannel(ctx context.Context, input *pinpoint.GetBaiduChannelInput) (*pinpoint.GetBaiduChannelOutput, error) {
	return a.client.GetBaiduChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetCampaign(ctx context.Context, input *pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error) {
	return a.client.GetCampaignWithContext(ctx, input)
}

func (a *PinpointActivities) GetCampaignActivities(ctx context.Context, input *pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error) {
	return a.client.GetCampaignActivitiesWithContext(ctx, input)
}

func (a *PinpointActivities) GetCampaignDateRangeKpi(ctx context.Context, input *pinpoint.GetCampaignDateRangeKpiInput) (*pinpoint.GetCampaignDateRangeKpiOutput, error) {
	return a.client.GetCampaignDateRangeKpiWithContext(ctx, input)
}

func (a *PinpointActivities) GetCampaignVersion(ctx context.Context, input *pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error) {
	return a.client.GetCampaignVersionWithContext(ctx, input)
}

func (a *PinpointActivities) GetCampaignVersions(ctx context.Context, input *pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error) {
	return a.client.GetCampaignVersionsWithContext(ctx, input)
}

func (a *PinpointActivities) GetCampaigns(ctx context.Context, input *pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error) {
	return a.client.GetCampaignsWithContext(ctx, input)
}

func (a *PinpointActivities) GetChannels(ctx context.Context, input *pinpoint.GetChannelsInput) (*pinpoint.GetChannelsOutput, error) {
	return a.client.GetChannelsWithContext(ctx, input)
}

func (a *PinpointActivities) GetEmailChannel(ctx context.Context, input *pinpoint.GetEmailChannelInput) (*pinpoint.GetEmailChannelOutput, error) {
	return a.client.GetEmailChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetEmailTemplate(ctx context.Context, input *pinpoint.GetEmailTemplateInput) (*pinpoint.GetEmailTemplateOutput, error) {
	return a.client.GetEmailTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) GetEndpoint(ctx context.Context, input *pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error) {
	return a.client.GetEndpointWithContext(ctx, input)
}

func (a *PinpointActivities) GetEventStream(ctx context.Context, input *pinpoint.GetEventStreamInput) (*pinpoint.GetEventStreamOutput, error) {
	return a.client.GetEventStreamWithContext(ctx, input)
}

func (a *PinpointActivities) GetExportJob(ctx context.Context, input *pinpoint.GetExportJobInput) (*pinpoint.GetExportJobOutput, error) {
	return a.client.GetExportJobWithContext(ctx, input)
}

func (a *PinpointActivities) GetExportJobs(ctx context.Context, input *pinpoint.GetExportJobsInput) (*pinpoint.GetExportJobsOutput, error) {
	return a.client.GetExportJobsWithContext(ctx, input)
}

func (a *PinpointActivities) GetGcmChannel(ctx context.Context, input *pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error) {
	return a.client.GetGcmChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetImportJob(ctx context.Context, input *pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error) {
	return a.client.GetImportJobWithContext(ctx, input)
}

func (a *PinpointActivities) GetImportJobs(ctx context.Context, input *pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error) {
	return a.client.GetImportJobsWithContext(ctx, input)
}

func (a *PinpointActivities) GetJourney(ctx context.Context, input *pinpoint.GetJourneyInput) (*pinpoint.GetJourneyOutput, error) {
	return a.client.GetJourneyWithContext(ctx, input)
}

func (a *PinpointActivities) GetJourneyDateRangeKpi(ctx context.Context, input *pinpoint.GetJourneyDateRangeKpiInput) (*pinpoint.GetJourneyDateRangeKpiOutput, error) {
	return a.client.GetJourneyDateRangeKpiWithContext(ctx, input)
}

func (a *PinpointActivities) GetJourneyExecutionActivityMetrics(ctx context.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error) {
	return a.client.GetJourneyExecutionActivityMetricsWithContext(ctx, input)
}

func (a *PinpointActivities) GetJourneyExecutionMetrics(ctx context.Context, input *pinpoint.GetJourneyExecutionMetricsInput) (*pinpoint.GetJourneyExecutionMetricsOutput, error) {
	return a.client.GetJourneyExecutionMetricsWithContext(ctx, input)
}

func (a *PinpointActivities) GetPushTemplate(ctx context.Context, input *pinpoint.GetPushTemplateInput) (*pinpoint.GetPushTemplateOutput, error) {
	return a.client.GetPushTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) GetRecommenderConfiguration(ctx context.Context, input *pinpoint.GetRecommenderConfigurationInput) (*pinpoint.GetRecommenderConfigurationOutput, error) {
	return a.client.GetRecommenderConfigurationWithContext(ctx, input)
}

func (a *PinpointActivities) GetRecommenderConfigurations(ctx context.Context, input *pinpoint.GetRecommenderConfigurationsInput) (*pinpoint.GetRecommenderConfigurationsOutput, error) {
	return a.client.GetRecommenderConfigurationsWithContext(ctx, input)
}

func (a *PinpointActivities) GetSegment(ctx context.Context, input *pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error) {
	return a.client.GetSegmentWithContext(ctx, input)
}

func (a *PinpointActivities) GetSegmentExportJobs(ctx context.Context, input *pinpoint.GetSegmentExportJobsInput) (*pinpoint.GetSegmentExportJobsOutput, error) {
	return a.client.GetSegmentExportJobsWithContext(ctx, input)
}

func (a *PinpointActivities) GetSegmentImportJobs(ctx context.Context, input *pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error) {
	return a.client.GetSegmentImportJobsWithContext(ctx, input)
}

func (a *PinpointActivities) GetSegmentVersion(ctx context.Context, input *pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error) {
	return a.client.GetSegmentVersionWithContext(ctx, input)
}

func (a *PinpointActivities) GetSegmentVersions(ctx context.Context, input *pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error) {
	return a.client.GetSegmentVersionsWithContext(ctx, input)
}

func (a *PinpointActivities) GetSegments(ctx context.Context, input *pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error) {
	return a.client.GetSegmentsWithContext(ctx, input)
}

func (a *PinpointActivities) GetSmsChannel(ctx context.Context, input *pinpoint.GetSmsChannelInput) (*pinpoint.GetSmsChannelOutput, error) {
	return a.client.GetSmsChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetSmsTemplate(ctx context.Context, input *pinpoint.GetSmsTemplateInput) (*pinpoint.GetSmsTemplateOutput, error) {
	return a.client.GetSmsTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) GetUserEndpoints(ctx context.Context, input *pinpoint.GetUserEndpointsInput) (*pinpoint.GetUserEndpointsOutput, error) {
	return a.client.GetUserEndpointsWithContext(ctx, input)
}

func (a *PinpointActivities) GetVoiceChannel(ctx context.Context, input *pinpoint.GetVoiceChannelInput) (*pinpoint.GetVoiceChannelOutput, error) {
	return a.client.GetVoiceChannelWithContext(ctx, input)
}

func (a *PinpointActivities) GetVoiceTemplate(ctx context.Context, input *pinpoint.GetVoiceTemplateInput) (*pinpoint.GetVoiceTemplateOutput, error) {
	return a.client.GetVoiceTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) ListJourneys(ctx context.Context, input *pinpoint.ListJourneysInput) (*pinpoint.ListJourneysOutput, error) {
	return a.client.ListJourneysWithContext(ctx, input)
}

func (a *PinpointActivities) ListTagsForResource(ctx context.Context, input *pinpoint.ListTagsForResourceInput) (*pinpoint.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *PinpointActivities) ListTemplateVersions(ctx context.Context, input *pinpoint.ListTemplateVersionsInput) (*pinpoint.ListTemplateVersionsOutput, error) {
	return a.client.ListTemplateVersionsWithContext(ctx, input)
}

func (a *PinpointActivities) ListTemplates(ctx context.Context, input *pinpoint.ListTemplatesInput) (*pinpoint.ListTemplatesOutput, error) {
	return a.client.ListTemplatesWithContext(ctx, input)
}

func (a *PinpointActivities) PhoneNumberValidate(ctx context.Context, input *pinpoint.PhoneNumberValidateInput) (*pinpoint.PhoneNumberValidateOutput, error) {
	return a.client.PhoneNumberValidateWithContext(ctx, input)
}

func (a *PinpointActivities) PutEventStream(ctx context.Context, input *pinpoint.PutEventStreamInput) (*pinpoint.PutEventStreamOutput, error) {
	return a.client.PutEventStreamWithContext(ctx, input)
}

func (a *PinpointActivities) PutEvents(ctx context.Context, input *pinpoint.PutEventsInput) (*pinpoint.PutEventsOutput, error) {
	return a.client.PutEventsWithContext(ctx, input)
}

func (a *PinpointActivities) RemoveAttributes(ctx context.Context, input *pinpoint.RemoveAttributesInput) (*pinpoint.RemoveAttributesOutput, error) {
	return a.client.RemoveAttributesWithContext(ctx, input)
}

func (a *PinpointActivities) SendMessages(ctx context.Context, input *pinpoint.SendMessagesInput) (*pinpoint.SendMessagesOutput, error) {
	return a.client.SendMessagesWithContext(ctx, input)
}

func (a *PinpointActivities) SendUsersMessages(ctx context.Context, input *pinpoint.SendUsersMessagesInput) (*pinpoint.SendUsersMessagesOutput, error) {
	return a.client.SendUsersMessagesWithContext(ctx, input)
}

func (a *PinpointActivities) TagResource(ctx context.Context, input *pinpoint.TagResourceInput) (*pinpoint.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *PinpointActivities) UntagResource(ctx context.Context, input *pinpoint.UntagResourceInput) (*pinpoint.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateAdmChannel(ctx context.Context, input *pinpoint.UpdateAdmChannelInput) (*pinpoint.UpdateAdmChannelOutput, error) {
	return a.client.UpdateAdmChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateApnsChannel(ctx context.Context, input *pinpoint.UpdateApnsChannelInput) (*pinpoint.UpdateApnsChannelOutput, error) {
	return a.client.UpdateApnsChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateApnsSandboxChannel(ctx context.Context, input *pinpoint.UpdateApnsSandboxChannelInput) (*pinpoint.UpdateApnsSandboxChannelOutput, error) {
	return a.client.UpdateApnsSandboxChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateApnsVoipChannel(ctx context.Context, input *pinpoint.UpdateApnsVoipChannelInput) (*pinpoint.UpdateApnsVoipChannelOutput, error) {
	return a.client.UpdateApnsVoipChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateApnsVoipSandboxChannel(ctx context.Context, input *pinpoint.UpdateApnsVoipSandboxChannelInput) (*pinpoint.UpdateApnsVoipSandboxChannelOutput, error) {
	return a.client.UpdateApnsVoipSandboxChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateApplicationSettings(ctx context.Context, input *pinpoint.UpdateApplicationSettingsInput) (*pinpoint.UpdateApplicationSettingsOutput, error) {
	return a.client.UpdateApplicationSettingsWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateBaiduChannel(ctx context.Context, input *pinpoint.UpdateBaiduChannelInput) (*pinpoint.UpdateBaiduChannelOutput, error) {
	return a.client.UpdateBaiduChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateCampaign(ctx context.Context, input *pinpoint.UpdateCampaignInput) (*pinpoint.UpdateCampaignOutput, error) {
	return a.client.UpdateCampaignWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateEmailChannel(ctx context.Context, input *pinpoint.UpdateEmailChannelInput) (*pinpoint.UpdateEmailChannelOutput, error) {
	return a.client.UpdateEmailChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateEmailTemplate(ctx context.Context, input *pinpoint.UpdateEmailTemplateInput) (*pinpoint.UpdateEmailTemplateOutput, error) {
	return a.client.UpdateEmailTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateEndpoint(ctx context.Context, input *pinpoint.UpdateEndpointInput) (*pinpoint.UpdateEndpointOutput, error) {
	return a.client.UpdateEndpointWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateEndpointsBatch(ctx context.Context, input *pinpoint.UpdateEndpointsBatchInput) (*pinpoint.UpdateEndpointsBatchOutput, error) {
	return a.client.UpdateEndpointsBatchWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateGcmChannel(ctx context.Context, input *pinpoint.UpdateGcmChannelInput) (*pinpoint.UpdateGcmChannelOutput, error) {
	return a.client.UpdateGcmChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateJourney(ctx context.Context, input *pinpoint.UpdateJourneyInput) (*pinpoint.UpdateJourneyOutput, error) {
	return a.client.UpdateJourneyWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateJourneyState(ctx context.Context, input *pinpoint.UpdateJourneyStateInput) (*pinpoint.UpdateJourneyStateOutput, error) {
	return a.client.UpdateJourneyStateWithContext(ctx, input)
}

func (a *PinpointActivities) UpdatePushTemplate(ctx context.Context, input *pinpoint.UpdatePushTemplateInput) (*pinpoint.UpdatePushTemplateOutput, error) {
	return a.client.UpdatePushTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateRecommenderConfiguration(ctx context.Context, input *pinpoint.UpdateRecommenderConfigurationInput) (*pinpoint.UpdateRecommenderConfigurationOutput, error) {
	return a.client.UpdateRecommenderConfigurationWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateSegment(ctx context.Context, input *pinpoint.UpdateSegmentInput) (*pinpoint.UpdateSegmentOutput, error) {
	return a.client.UpdateSegmentWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateSmsChannel(ctx context.Context, input *pinpoint.UpdateSmsChannelInput) (*pinpoint.UpdateSmsChannelOutput, error) {
	return a.client.UpdateSmsChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateSmsTemplate(ctx context.Context, input *pinpoint.UpdateSmsTemplateInput) (*pinpoint.UpdateSmsTemplateOutput, error) {
	return a.client.UpdateSmsTemplateWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateTemplateActiveVersion(ctx context.Context, input *pinpoint.UpdateTemplateActiveVersionInput) (*pinpoint.UpdateTemplateActiveVersionOutput, error) {
	return a.client.UpdateTemplateActiveVersionWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateVoiceChannel(ctx context.Context, input *pinpoint.UpdateVoiceChannelInput) (*pinpoint.UpdateVoiceChannelOutput, error) {
	return a.client.UpdateVoiceChannelWithContext(ctx, input)
}

func (a *PinpointActivities) UpdateVoiceTemplate(ctx context.Context, input *pinpoint.UpdateVoiceTemplateInput) (*pinpoint.UpdateVoiceTemplateOutput, error) {
	return a.client.UpdateVoiceTemplateWithContext(ctx, input)
}
