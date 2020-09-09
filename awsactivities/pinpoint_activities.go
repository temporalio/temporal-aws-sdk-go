
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/pinpoint/pinpointiface"
)

type PinpointActivities struct {
	client pinpointiface.PinpointAPI
}

func NewPinpointActivities(session *session.Session, config... *aws.Config) *PinpointActivities {
    client := pinpoint.New(session, config...)
    return &PinpointActivities{client: client}
}

func (a *PinpointActivities) CreateApp(input *pinpoint.CreateAppInput) (*pinpoint.CreateAppOutput, error) {
    return a.client.CreateApp(input)
}

func (a *PinpointActivities) CreateCampaign(input *pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error) {
    return a.client.CreateCampaign(input)
}

func (a *PinpointActivities) CreateEmailTemplate(input *pinpoint.CreateEmailTemplateInput) (*pinpoint.CreateEmailTemplateOutput, error) {
    return a.client.CreateEmailTemplate(input)
}

func (a *PinpointActivities) CreateExportJob(input *pinpoint.CreateExportJobInput) (*pinpoint.CreateExportJobOutput, error) {
    return a.client.CreateExportJob(input)
}

func (a *PinpointActivities) CreateImportJob(input *pinpoint.CreateImportJobInput) (*pinpoint.CreateImportJobOutput, error) {
    return a.client.CreateImportJob(input)
}

func (a *PinpointActivities) CreateJourney(input *pinpoint.CreateJourneyInput) (*pinpoint.CreateJourneyOutput, error) {
    return a.client.CreateJourney(input)
}

func (a *PinpointActivities) CreatePushTemplate(input *pinpoint.CreatePushTemplateInput) (*pinpoint.CreatePushTemplateOutput, error) {
    return a.client.CreatePushTemplate(input)
}

func (a *PinpointActivities) CreateRecommenderConfiguration(input *pinpoint.CreateRecommenderConfigurationInput) (*pinpoint.CreateRecommenderConfigurationOutput, error) {
    return a.client.CreateRecommenderConfiguration(input)
}

func (a *PinpointActivities) CreateSegment(input *pinpoint.CreateSegmentInput) (*pinpoint.CreateSegmentOutput, error) {
    return a.client.CreateSegment(input)
}

func (a *PinpointActivities) CreateSmsTemplate(input *pinpoint.CreateSmsTemplateInput) (*pinpoint.CreateSmsTemplateOutput, error) {
    return a.client.CreateSmsTemplate(input)
}

func (a *PinpointActivities) CreateVoiceTemplate(input *pinpoint.CreateVoiceTemplateInput) (*pinpoint.CreateVoiceTemplateOutput, error) {
    return a.client.CreateVoiceTemplate(input)
}

func (a *PinpointActivities) DeleteAdmChannel(input *pinpoint.DeleteAdmChannelInput) (*pinpoint.DeleteAdmChannelOutput, error) {
    return a.client.DeleteAdmChannel(input)
}

func (a *PinpointActivities) DeleteApnsChannel(input *pinpoint.DeleteApnsChannelInput) (*pinpoint.DeleteApnsChannelOutput, error) {
    return a.client.DeleteApnsChannel(input)
}

func (a *PinpointActivities) DeleteApnsSandboxChannel(input *pinpoint.DeleteApnsSandboxChannelInput) (*pinpoint.DeleteApnsSandboxChannelOutput, error) {
    return a.client.DeleteApnsSandboxChannel(input)
}

func (a *PinpointActivities) DeleteApnsVoipChannel(input *pinpoint.DeleteApnsVoipChannelInput) (*pinpoint.DeleteApnsVoipChannelOutput, error) {
    return a.client.DeleteApnsVoipChannel(input)
}

func (a *PinpointActivities) DeleteApnsVoipSandboxChannel(input *pinpoint.DeleteApnsVoipSandboxChannelInput) (*pinpoint.DeleteApnsVoipSandboxChannelOutput, error) {
    return a.client.DeleteApnsVoipSandboxChannel(input)
}

func (a *PinpointActivities) DeleteApp(input *pinpoint.DeleteAppInput) (*pinpoint.DeleteAppOutput, error) {
    return a.client.DeleteApp(input)
}

func (a *PinpointActivities) DeleteBaiduChannel(input *pinpoint.DeleteBaiduChannelInput) (*pinpoint.DeleteBaiduChannelOutput, error) {
    return a.client.DeleteBaiduChannel(input)
}

func (a *PinpointActivities) DeleteCampaign(input *pinpoint.DeleteCampaignInput) (*pinpoint.DeleteCampaignOutput, error) {
    return a.client.DeleteCampaign(input)
}

func (a *PinpointActivities) DeleteEmailChannel(input *pinpoint.DeleteEmailChannelInput) (*pinpoint.DeleteEmailChannelOutput, error) {
    return a.client.DeleteEmailChannel(input)
}

func (a *PinpointActivities) DeleteEmailTemplate(input *pinpoint.DeleteEmailTemplateInput) (*pinpoint.DeleteEmailTemplateOutput, error) {
    return a.client.DeleteEmailTemplate(input)
}

func (a *PinpointActivities) DeleteEndpoint(input *pinpoint.DeleteEndpointInput) (*pinpoint.DeleteEndpointOutput, error) {
    return a.client.DeleteEndpoint(input)
}

func (a *PinpointActivities) DeleteEventStream(input *pinpoint.DeleteEventStreamInput) (*pinpoint.DeleteEventStreamOutput, error) {
    return a.client.DeleteEventStream(input)
}

func (a *PinpointActivities) DeleteGcmChannel(input *pinpoint.DeleteGcmChannelInput) (*pinpoint.DeleteGcmChannelOutput, error) {
    return a.client.DeleteGcmChannel(input)
}

func (a *PinpointActivities) DeleteJourney(input *pinpoint.DeleteJourneyInput) (*pinpoint.DeleteJourneyOutput, error) {
    return a.client.DeleteJourney(input)
}

func (a *PinpointActivities) DeletePushTemplate(input *pinpoint.DeletePushTemplateInput) (*pinpoint.DeletePushTemplateOutput, error) {
    return a.client.DeletePushTemplate(input)
}

func (a *PinpointActivities) DeleteRecommenderConfiguration(input *pinpoint.DeleteRecommenderConfigurationInput) (*pinpoint.DeleteRecommenderConfigurationOutput, error) {
    return a.client.DeleteRecommenderConfiguration(input)
}

func (a *PinpointActivities) DeleteSegment(input *pinpoint.DeleteSegmentInput) (*pinpoint.DeleteSegmentOutput, error) {
    return a.client.DeleteSegment(input)
}

func (a *PinpointActivities) DeleteSmsChannel(input *pinpoint.DeleteSmsChannelInput) (*pinpoint.DeleteSmsChannelOutput, error) {
    return a.client.DeleteSmsChannel(input)
}

func (a *PinpointActivities) DeleteSmsTemplate(input *pinpoint.DeleteSmsTemplateInput) (*pinpoint.DeleteSmsTemplateOutput, error) {
    return a.client.DeleteSmsTemplate(input)
}

func (a *PinpointActivities) DeleteUserEndpoints(input *pinpoint.DeleteUserEndpointsInput) (*pinpoint.DeleteUserEndpointsOutput, error) {
    return a.client.DeleteUserEndpoints(input)
}

func (a *PinpointActivities) DeleteVoiceChannel(input *pinpoint.DeleteVoiceChannelInput) (*pinpoint.DeleteVoiceChannelOutput, error) {
    return a.client.DeleteVoiceChannel(input)
}

func (a *PinpointActivities) DeleteVoiceTemplate(input *pinpoint.DeleteVoiceTemplateInput) (*pinpoint.DeleteVoiceTemplateOutput, error) {
    return a.client.DeleteVoiceTemplate(input)
}

func (a *PinpointActivities) GetAdmChannel(input *pinpoint.GetAdmChannelInput) (*pinpoint.GetAdmChannelOutput, error) {
    return a.client.GetAdmChannel(input)
}

func (a *PinpointActivities) GetApnsChannel(input *pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error) {
    return a.client.GetApnsChannel(input)
}

func (a *PinpointActivities) GetApnsSandboxChannel(input *pinpoint.GetApnsSandboxChannelInput) (*pinpoint.GetApnsSandboxChannelOutput, error) {
    return a.client.GetApnsSandboxChannel(input)
}

func (a *PinpointActivities) GetApnsVoipChannel(input *pinpoint.GetApnsVoipChannelInput) (*pinpoint.GetApnsVoipChannelOutput, error) {
    return a.client.GetApnsVoipChannel(input)
}

func (a *PinpointActivities) GetApnsVoipSandboxChannel(input *pinpoint.GetApnsVoipSandboxChannelInput) (*pinpoint.GetApnsVoipSandboxChannelOutput, error) {
    return a.client.GetApnsVoipSandboxChannel(input)
}

func (a *PinpointActivities) GetApp(input *pinpoint.GetAppInput) (*pinpoint.GetAppOutput, error) {
    return a.client.GetApp(input)
}

func (a *PinpointActivities) GetApplicationDateRangeKpi(input *pinpoint.GetApplicationDateRangeKpiInput) (*pinpoint.GetApplicationDateRangeKpiOutput, error) {
    return a.client.GetApplicationDateRangeKpi(input)
}

func (a *PinpointActivities) GetApplicationSettings(input *pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error) {
    return a.client.GetApplicationSettings(input)
}

func (a *PinpointActivities) GetApps(input *pinpoint.GetAppsInput) (*pinpoint.GetAppsOutput, error) {
    return a.client.GetApps(input)
}

func (a *PinpointActivities) GetBaiduChannel(input *pinpoint.GetBaiduChannelInput) (*pinpoint.GetBaiduChannelOutput, error) {
    return a.client.GetBaiduChannel(input)
}

func (a *PinpointActivities) GetCampaign(input *pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error) {
    return a.client.GetCampaign(input)
}

func (a *PinpointActivities) GetCampaignActivities(input *pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error) {
    return a.client.GetCampaignActivities(input)
}

func (a *PinpointActivities) GetCampaignDateRangeKpi(input *pinpoint.GetCampaignDateRangeKpiInput) (*pinpoint.GetCampaignDateRangeKpiOutput, error) {
    return a.client.GetCampaignDateRangeKpi(input)
}

func (a *PinpointActivities) GetCampaignVersion(input *pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error) {
    return a.client.GetCampaignVersion(input)
}

func (a *PinpointActivities) GetCampaignVersions(input *pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error) {
    return a.client.GetCampaignVersions(input)
}

func (a *PinpointActivities) GetCampaigns(input *pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error) {
    return a.client.GetCampaigns(input)
}

func (a *PinpointActivities) GetChannels(input *pinpoint.GetChannelsInput) (*pinpoint.GetChannelsOutput, error) {
    return a.client.GetChannels(input)
}

func (a *PinpointActivities) GetEmailChannel(input *pinpoint.GetEmailChannelInput) (*pinpoint.GetEmailChannelOutput, error) {
    return a.client.GetEmailChannel(input)
}

func (a *PinpointActivities) GetEmailTemplate(input *pinpoint.GetEmailTemplateInput) (*pinpoint.GetEmailTemplateOutput, error) {
    return a.client.GetEmailTemplate(input)
}

func (a *PinpointActivities) GetEndpoint(input *pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error) {
    return a.client.GetEndpoint(input)
}

func (a *PinpointActivities) GetEventStream(input *pinpoint.GetEventStreamInput) (*pinpoint.GetEventStreamOutput, error) {
    return a.client.GetEventStream(input)
}

func (a *PinpointActivities) GetExportJob(input *pinpoint.GetExportJobInput) (*pinpoint.GetExportJobOutput, error) {
    return a.client.GetExportJob(input)
}

func (a *PinpointActivities) GetExportJobs(input *pinpoint.GetExportJobsInput) (*pinpoint.GetExportJobsOutput, error) {
    return a.client.GetExportJobs(input)
}

func (a *PinpointActivities) GetGcmChannel(input *pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error) {
    return a.client.GetGcmChannel(input)
}

func (a *PinpointActivities) GetImportJob(input *pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error) {
    return a.client.GetImportJob(input)
}

func (a *PinpointActivities) GetImportJobs(input *pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error) {
    return a.client.GetImportJobs(input)
}

func (a *PinpointActivities) GetJourney(input *pinpoint.GetJourneyInput) (*pinpoint.GetJourneyOutput, error) {
    return a.client.GetJourney(input)
}

func (a *PinpointActivities) GetJourneyDateRangeKpi(input *pinpoint.GetJourneyDateRangeKpiInput) (*pinpoint.GetJourneyDateRangeKpiOutput, error) {
    return a.client.GetJourneyDateRangeKpi(input)
}

func (a *PinpointActivities) GetJourneyExecutionActivityMetrics(input *pinpoint.GetJourneyExecutionActivityMetricsInput) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error) {
    return a.client.GetJourneyExecutionActivityMetrics(input)
}

func (a *PinpointActivities) GetJourneyExecutionMetrics(input *pinpoint.GetJourneyExecutionMetricsInput) (*pinpoint.GetJourneyExecutionMetricsOutput, error) {
    return a.client.GetJourneyExecutionMetrics(input)
}

func (a *PinpointActivities) GetPushTemplate(input *pinpoint.GetPushTemplateInput) (*pinpoint.GetPushTemplateOutput, error) {
    return a.client.GetPushTemplate(input)
}

func (a *PinpointActivities) GetRecommenderConfiguration(input *pinpoint.GetRecommenderConfigurationInput) (*pinpoint.GetRecommenderConfigurationOutput, error) {
    return a.client.GetRecommenderConfiguration(input)
}

func (a *PinpointActivities) GetRecommenderConfigurations(input *pinpoint.GetRecommenderConfigurationsInput) (*pinpoint.GetRecommenderConfigurationsOutput, error) {
    return a.client.GetRecommenderConfigurations(input)
}

func (a *PinpointActivities) GetSegment(input *pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error) {
    return a.client.GetSegment(input)
}

func (a *PinpointActivities) GetSegmentExportJobs(input *pinpoint.GetSegmentExportJobsInput) (*pinpoint.GetSegmentExportJobsOutput, error) {
    return a.client.GetSegmentExportJobs(input)
}

func (a *PinpointActivities) GetSegmentImportJobs(input *pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error) {
    return a.client.GetSegmentImportJobs(input)
}

func (a *PinpointActivities) GetSegmentVersion(input *pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error) {
    return a.client.GetSegmentVersion(input)
}

func (a *PinpointActivities) GetSegmentVersions(input *pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error) {
    return a.client.GetSegmentVersions(input)
}

func (a *PinpointActivities) GetSegments(input *pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error) {
    return a.client.GetSegments(input)
}

func (a *PinpointActivities) GetSmsChannel(input *pinpoint.GetSmsChannelInput) (*pinpoint.GetSmsChannelOutput, error) {
    return a.client.GetSmsChannel(input)
}

func (a *PinpointActivities) GetSmsTemplate(input *pinpoint.GetSmsTemplateInput) (*pinpoint.GetSmsTemplateOutput, error) {
    return a.client.GetSmsTemplate(input)
}

func (a *PinpointActivities) GetUserEndpoints(input *pinpoint.GetUserEndpointsInput) (*pinpoint.GetUserEndpointsOutput, error) {
    return a.client.GetUserEndpoints(input)
}

func (a *PinpointActivities) GetVoiceChannel(input *pinpoint.GetVoiceChannelInput) (*pinpoint.GetVoiceChannelOutput, error) {
    return a.client.GetVoiceChannel(input)
}

func (a *PinpointActivities) GetVoiceTemplate(input *pinpoint.GetVoiceTemplateInput) (*pinpoint.GetVoiceTemplateOutput, error) {
    return a.client.GetVoiceTemplate(input)
}

func (a *PinpointActivities) ListJourneys(input *pinpoint.ListJourneysInput) (*pinpoint.ListJourneysOutput, error) {
    return a.client.ListJourneys(input)
}

func (a *PinpointActivities) ListTagsForResource(input *pinpoint.ListTagsForResourceInput) (*pinpoint.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *PinpointActivities) ListTemplateVersions(input *pinpoint.ListTemplateVersionsInput) (*pinpoint.ListTemplateVersionsOutput, error) {
    return a.client.ListTemplateVersions(input)
}

func (a *PinpointActivities) ListTemplates(input *pinpoint.ListTemplatesInput) (*pinpoint.ListTemplatesOutput, error) {
    return a.client.ListTemplates(input)
}

func (a *PinpointActivities) PhoneNumberValidate(input *pinpoint.PhoneNumberValidateInput) (*pinpoint.PhoneNumberValidateOutput, error) {
    return a.client.PhoneNumberValidate(input)
}

func (a *PinpointActivities) PutEventStream(input *pinpoint.PutEventStreamInput) (*pinpoint.PutEventStreamOutput, error) {
    return a.client.PutEventStream(input)
}

func (a *PinpointActivities) PutEvents(input *pinpoint.PutEventsInput) (*pinpoint.PutEventsOutput, error) {
    return a.client.PutEvents(input)
}

func (a *PinpointActivities) RemoveAttributes(input *pinpoint.RemoveAttributesInput) (*pinpoint.RemoveAttributesOutput, error) {
    return a.client.RemoveAttributes(input)
}

func (a *PinpointActivities) SendMessages(input *pinpoint.SendMessagesInput) (*pinpoint.SendMessagesOutput, error) {
    return a.client.SendMessages(input)
}

func (a *PinpointActivities) SendUsersMessages(input *pinpoint.SendUsersMessagesInput) (*pinpoint.SendUsersMessagesOutput, error) {
    return a.client.SendUsersMessages(input)
}

func (a *PinpointActivities) TagResource(input *pinpoint.TagResourceInput) (*pinpoint.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *PinpointActivities) UntagResource(input *pinpoint.UntagResourceInput) (*pinpoint.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *PinpointActivities) UpdateAdmChannel(input *pinpoint.UpdateAdmChannelInput) (*pinpoint.UpdateAdmChannelOutput, error) {
    return a.client.UpdateAdmChannel(input)
}

func (a *PinpointActivities) UpdateApnsChannel(input *pinpoint.UpdateApnsChannelInput) (*pinpoint.UpdateApnsChannelOutput, error) {
    return a.client.UpdateApnsChannel(input)
}

func (a *PinpointActivities) UpdateApnsSandboxChannel(input *pinpoint.UpdateApnsSandboxChannelInput) (*pinpoint.UpdateApnsSandboxChannelOutput, error) {
    return a.client.UpdateApnsSandboxChannel(input)
}

func (a *PinpointActivities) UpdateApnsVoipChannel(input *pinpoint.UpdateApnsVoipChannelInput) (*pinpoint.UpdateApnsVoipChannelOutput, error) {
    return a.client.UpdateApnsVoipChannel(input)
}

func (a *PinpointActivities) UpdateApnsVoipSandboxChannel(input *pinpoint.UpdateApnsVoipSandboxChannelInput) (*pinpoint.UpdateApnsVoipSandboxChannelOutput, error) {
    return a.client.UpdateApnsVoipSandboxChannel(input)
}

func (a *PinpointActivities) UpdateApplicationSettings(input *pinpoint.UpdateApplicationSettingsInput) (*pinpoint.UpdateApplicationSettingsOutput, error) {
    return a.client.UpdateApplicationSettings(input)
}

func (a *PinpointActivities) UpdateBaiduChannel(input *pinpoint.UpdateBaiduChannelInput) (*pinpoint.UpdateBaiduChannelOutput, error) {
    return a.client.UpdateBaiduChannel(input)
}

func (a *PinpointActivities) UpdateCampaign(input *pinpoint.UpdateCampaignInput) (*pinpoint.UpdateCampaignOutput, error) {
    return a.client.UpdateCampaign(input)
}

func (a *PinpointActivities) UpdateEmailChannel(input *pinpoint.UpdateEmailChannelInput) (*pinpoint.UpdateEmailChannelOutput, error) {
    return a.client.UpdateEmailChannel(input)
}

func (a *PinpointActivities) UpdateEmailTemplate(input *pinpoint.UpdateEmailTemplateInput) (*pinpoint.UpdateEmailTemplateOutput, error) {
    return a.client.UpdateEmailTemplate(input)
}

func (a *PinpointActivities) UpdateEndpoint(input *pinpoint.UpdateEndpointInput) (*pinpoint.UpdateEndpointOutput, error) {
    return a.client.UpdateEndpoint(input)
}

func (a *PinpointActivities) UpdateEndpointsBatch(input *pinpoint.UpdateEndpointsBatchInput) (*pinpoint.UpdateEndpointsBatchOutput, error) {
    return a.client.UpdateEndpointsBatch(input)
}

func (a *PinpointActivities) UpdateGcmChannel(input *pinpoint.UpdateGcmChannelInput) (*pinpoint.UpdateGcmChannelOutput, error) {
    return a.client.UpdateGcmChannel(input)
}

func (a *PinpointActivities) UpdateJourney(input *pinpoint.UpdateJourneyInput) (*pinpoint.UpdateJourneyOutput, error) {
    return a.client.UpdateJourney(input)
}

func (a *PinpointActivities) UpdateJourneyState(input *pinpoint.UpdateJourneyStateInput) (*pinpoint.UpdateJourneyStateOutput, error) {
    return a.client.UpdateJourneyState(input)
}

func (a *PinpointActivities) UpdatePushTemplate(input *pinpoint.UpdatePushTemplateInput) (*pinpoint.UpdatePushTemplateOutput, error) {
    return a.client.UpdatePushTemplate(input)
}

func (a *PinpointActivities) UpdateRecommenderConfiguration(input *pinpoint.UpdateRecommenderConfigurationInput) (*pinpoint.UpdateRecommenderConfigurationOutput, error) {
    return a.client.UpdateRecommenderConfiguration(input)
}

func (a *PinpointActivities) UpdateSegment(input *pinpoint.UpdateSegmentInput) (*pinpoint.UpdateSegmentOutput, error) {
    return a.client.UpdateSegment(input)
}

func (a *PinpointActivities) UpdateSmsChannel(input *pinpoint.UpdateSmsChannelInput) (*pinpoint.UpdateSmsChannelOutput, error) {
    return a.client.UpdateSmsChannel(input)
}

func (a *PinpointActivities) UpdateSmsTemplate(input *pinpoint.UpdateSmsTemplateInput) (*pinpoint.UpdateSmsTemplateOutput, error) {
    return a.client.UpdateSmsTemplate(input)
}

func (a *PinpointActivities) UpdateTemplateActiveVersion(input *pinpoint.UpdateTemplateActiveVersionInput) (*pinpoint.UpdateTemplateActiveVersionOutput, error) {
    return a.client.UpdateTemplateActiveVersion(input)
}

func (a *PinpointActivities) UpdateVoiceChannel(input *pinpoint.UpdateVoiceChannelInput) (*pinpoint.UpdateVoiceChannelOutput, error) {
    return a.client.UpdateVoiceChannel(input)
}

func (a *PinpointActivities) UpdateVoiceTemplate(input *pinpoint.UpdateVoiceTemplateInput) (*pinpoint.UpdateVoiceTemplateOutput, error) {
    return a.client.UpdateVoiceTemplate(input)
}
