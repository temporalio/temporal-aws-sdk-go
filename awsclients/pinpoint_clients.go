package awsclients

import (
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type PinpointClient interface {
       CreateApp(ctx workflow.Context, input *pinpoint.CreateAppInput) (*pinpoint.CreateAppOutput, error)
       CreateAppAsync(ctx workflow.Context, input *pinpoint.CreateAppInput) *PinpointCreateAppResult

       CreateCampaign(ctx workflow.Context, input *pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error)
       CreateCampaignAsync(ctx workflow.Context, input *pinpoint.CreateCampaignInput) *PinpointCreateCampaignResult

       CreateEmailTemplate(ctx workflow.Context, input *pinpoint.CreateEmailTemplateInput) (*pinpoint.CreateEmailTemplateOutput, error)
       CreateEmailTemplateAsync(ctx workflow.Context, input *pinpoint.CreateEmailTemplateInput) *PinpointCreateEmailTemplateResult

       CreateExportJob(ctx workflow.Context, input *pinpoint.CreateExportJobInput) (*pinpoint.CreateExportJobOutput, error)
       CreateExportJobAsync(ctx workflow.Context, input *pinpoint.CreateExportJobInput) *PinpointCreateExportJobResult

       CreateImportJob(ctx workflow.Context, input *pinpoint.CreateImportJobInput) (*pinpoint.CreateImportJobOutput, error)
       CreateImportJobAsync(ctx workflow.Context, input *pinpoint.CreateImportJobInput) *PinpointCreateImportJobResult

       CreateJourney(ctx workflow.Context, input *pinpoint.CreateJourneyInput) (*pinpoint.CreateJourneyOutput, error)
       CreateJourneyAsync(ctx workflow.Context, input *pinpoint.CreateJourneyInput) *PinpointCreateJourneyResult

       CreatePushTemplate(ctx workflow.Context, input *pinpoint.CreatePushTemplateInput) (*pinpoint.CreatePushTemplateOutput, error)
       CreatePushTemplateAsync(ctx workflow.Context, input *pinpoint.CreatePushTemplateInput) *PinpointCreatePushTemplateResult

       CreateRecommenderConfiguration(ctx workflow.Context, input *pinpoint.CreateRecommenderConfigurationInput) (*pinpoint.CreateRecommenderConfigurationOutput, error)
       CreateRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.CreateRecommenderConfigurationInput) *PinpointCreateRecommenderConfigurationResult

       CreateSegment(ctx workflow.Context, input *pinpoint.CreateSegmentInput) (*pinpoint.CreateSegmentOutput, error)
       CreateSegmentAsync(ctx workflow.Context, input *pinpoint.CreateSegmentInput) *PinpointCreateSegmentResult

       CreateSmsTemplate(ctx workflow.Context, input *pinpoint.CreateSmsTemplateInput) (*pinpoint.CreateSmsTemplateOutput, error)
       CreateSmsTemplateAsync(ctx workflow.Context, input *pinpoint.CreateSmsTemplateInput) *PinpointCreateSmsTemplateResult

       CreateVoiceTemplate(ctx workflow.Context, input *pinpoint.CreateVoiceTemplateInput) (*pinpoint.CreateVoiceTemplateOutput, error)
       CreateVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.CreateVoiceTemplateInput) *PinpointCreateVoiceTemplateResult

       DeleteAdmChannel(ctx workflow.Context, input *pinpoint.DeleteAdmChannelInput) (*pinpoint.DeleteAdmChannelOutput, error)
       DeleteAdmChannelAsync(ctx workflow.Context, input *pinpoint.DeleteAdmChannelInput) *PinpointDeleteAdmChannelResult

       DeleteApnsChannel(ctx workflow.Context, input *pinpoint.DeleteApnsChannelInput) (*pinpoint.DeleteApnsChannelOutput, error)
       DeleteApnsChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsChannelInput) *PinpointDeleteApnsChannelResult

       DeleteApnsSandboxChannel(ctx workflow.Context, input *pinpoint.DeleteApnsSandboxChannelInput) (*pinpoint.DeleteApnsSandboxChannelOutput, error)
       DeleteApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsSandboxChannelInput) *PinpointDeleteApnsSandboxChannelResult

       DeleteApnsVoipChannel(ctx workflow.Context, input *pinpoint.DeleteApnsVoipChannelInput) (*pinpoint.DeleteApnsVoipChannelOutput, error)
       DeleteApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsVoipChannelInput) *PinpointDeleteApnsVoipChannelResult

       DeleteApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.DeleteApnsVoipSandboxChannelInput) (*pinpoint.DeleteApnsVoipSandboxChannelOutput, error)
       DeleteApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsVoipSandboxChannelInput) *PinpointDeleteApnsVoipSandboxChannelResult

       DeleteApp(ctx workflow.Context, input *pinpoint.DeleteAppInput) (*pinpoint.DeleteAppOutput, error)
       DeleteAppAsync(ctx workflow.Context, input *pinpoint.DeleteAppInput) *PinpointDeleteAppResult

       DeleteBaiduChannel(ctx workflow.Context, input *pinpoint.DeleteBaiduChannelInput) (*pinpoint.DeleteBaiduChannelOutput, error)
       DeleteBaiduChannelAsync(ctx workflow.Context, input *pinpoint.DeleteBaiduChannelInput) *PinpointDeleteBaiduChannelResult

       DeleteCampaign(ctx workflow.Context, input *pinpoint.DeleteCampaignInput) (*pinpoint.DeleteCampaignOutput, error)
       DeleteCampaignAsync(ctx workflow.Context, input *pinpoint.DeleteCampaignInput) *PinpointDeleteCampaignResult

       DeleteEmailChannel(ctx workflow.Context, input *pinpoint.DeleteEmailChannelInput) (*pinpoint.DeleteEmailChannelOutput, error)
       DeleteEmailChannelAsync(ctx workflow.Context, input *pinpoint.DeleteEmailChannelInput) *PinpointDeleteEmailChannelResult

       DeleteEmailTemplate(ctx workflow.Context, input *pinpoint.DeleteEmailTemplateInput) (*pinpoint.DeleteEmailTemplateOutput, error)
       DeleteEmailTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteEmailTemplateInput) *PinpointDeleteEmailTemplateResult

       DeleteEndpoint(ctx workflow.Context, input *pinpoint.DeleteEndpointInput) (*pinpoint.DeleteEndpointOutput, error)
       DeleteEndpointAsync(ctx workflow.Context, input *pinpoint.DeleteEndpointInput) *PinpointDeleteEndpointResult

       DeleteEventStream(ctx workflow.Context, input *pinpoint.DeleteEventStreamInput) (*pinpoint.DeleteEventStreamOutput, error)
       DeleteEventStreamAsync(ctx workflow.Context, input *pinpoint.DeleteEventStreamInput) *PinpointDeleteEventStreamResult

       DeleteGcmChannel(ctx workflow.Context, input *pinpoint.DeleteGcmChannelInput) (*pinpoint.DeleteGcmChannelOutput, error)
       DeleteGcmChannelAsync(ctx workflow.Context, input *pinpoint.DeleteGcmChannelInput) *PinpointDeleteGcmChannelResult

       DeleteJourney(ctx workflow.Context, input *pinpoint.DeleteJourneyInput) (*pinpoint.DeleteJourneyOutput, error)
       DeleteJourneyAsync(ctx workflow.Context, input *pinpoint.DeleteJourneyInput) *PinpointDeleteJourneyResult

       DeletePushTemplate(ctx workflow.Context, input *pinpoint.DeletePushTemplateInput) (*pinpoint.DeletePushTemplateOutput, error)
       DeletePushTemplateAsync(ctx workflow.Context, input *pinpoint.DeletePushTemplateInput) *PinpointDeletePushTemplateResult

       DeleteRecommenderConfiguration(ctx workflow.Context, input *pinpoint.DeleteRecommenderConfigurationInput) (*pinpoint.DeleteRecommenderConfigurationOutput, error)
       DeleteRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.DeleteRecommenderConfigurationInput) *PinpointDeleteRecommenderConfigurationResult

       DeleteSegment(ctx workflow.Context, input *pinpoint.DeleteSegmentInput) (*pinpoint.DeleteSegmentOutput, error)
       DeleteSegmentAsync(ctx workflow.Context, input *pinpoint.DeleteSegmentInput) *PinpointDeleteSegmentResult

       DeleteSmsChannel(ctx workflow.Context, input *pinpoint.DeleteSmsChannelInput) (*pinpoint.DeleteSmsChannelOutput, error)
       DeleteSmsChannelAsync(ctx workflow.Context, input *pinpoint.DeleteSmsChannelInput) *PinpointDeleteSmsChannelResult

       DeleteSmsTemplate(ctx workflow.Context, input *pinpoint.DeleteSmsTemplateInput) (*pinpoint.DeleteSmsTemplateOutput, error)
       DeleteSmsTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteSmsTemplateInput) *PinpointDeleteSmsTemplateResult

       DeleteUserEndpoints(ctx workflow.Context, input *pinpoint.DeleteUserEndpointsInput) (*pinpoint.DeleteUserEndpointsOutput, error)
       DeleteUserEndpointsAsync(ctx workflow.Context, input *pinpoint.DeleteUserEndpointsInput) *PinpointDeleteUserEndpointsResult

       DeleteVoiceChannel(ctx workflow.Context, input *pinpoint.DeleteVoiceChannelInput) (*pinpoint.DeleteVoiceChannelOutput, error)
       DeleteVoiceChannelAsync(ctx workflow.Context, input *pinpoint.DeleteVoiceChannelInput) *PinpointDeleteVoiceChannelResult

       DeleteVoiceTemplate(ctx workflow.Context, input *pinpoint.DeleteVoiceTemplateInput) (*pinpoint.DeleteVoiceTemplateOutput, error)
       DeleteVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteVoiceTemplateInput) *PinpointDeleteVoiceTemplateResult

       GetAdmChannel(ctx workflow.Context, input *pinpoint.GetAdmChannelInput) (*pinpoint.GetAdmChannelOutput, error)
       GetAdmChannelAsync(ctx workflow.Context, input *pinpoint.GetAdmChannelInput) *PinpointGetAdmChannelResult

       GetApnsChannel(ctx workflow.Context, input *pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error)
       GetApnsChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsChannelInput) *PinpointGetApnsChannelResult

       GetApnsSandboxChannel(ctx workflow.Context, input *pinpoint.GetApnsSandboxChannelInput) (*pinpoint.GetApnsSandboxChannelOutput, error)
       GetApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsSandboxChannelInput) *PinpointGetApnsSandboxChannelResult

       GetApnsVoipChannel(ctx workflow.Context, input *pinpoint.GetApnsVoipChannelInput) (*pinpoint.GetApnsVoipChannelOutput, error)
       GetApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsVoipChannelInput) *PinpointGetApnsVoipChannelResult

       GetApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.GetApnsVoipSandboxChannelInput) (*pinpoint.GetApnsVoipSandboxChannelOutput, error)
       GetApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsVoipSandboxChannelInput) *PinpointGetApnsVoipSandboxChannelResult

       GetApp(ctx workflow.Context, input *pinpoint.GetAppInput) (*pinpoint.GetAppOutput, error)
       GetAppAsync(ctx workflow.Context, input *pinpoint.GetAppInput) *PinpointGetAppResult

       GetApplicationDateRangeKpi(ctx workflow.Context, input *pinpoint.GetApplicationDateRangeKpiInput) (*pinpoint.GetApplicationDateRangeKpiOutput, error)
       GetApplicationDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetApplicationDateRangeKpiInput) *PinpointGetApplicationDateRangeKpiResult

       GetApplicationSettings(ctx workflow.Context, input *pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error)
       GetApplicationSettingsAsync(ctx workflow.Context, input *pinpoint.GetApplicationSettingsInput) *PinpointGetApplicationSettingsResult

       GetApps(ctx workflow.Context, input *pinpoint.GetAppsInput) (*pinpoint.GetAppsOutput, error)
       GetAppsAsync(ctx workflow.Context, input *pinpoint.GetAppsInput) *PinpointGetAppsResult

       GetBaiduChannel(ctx workflow.Context, input *pinpoint.GetBaiduChannelInput) (*pinpoint.GetBaiduChannelOutput, error)
       GetBaiduChannelAsync(ctx workflow.Context, input *pinpoint.GetBaiduChannelInput) *PinpointGetBaiduChannelResult

       GetCampaign(ctx workflow.Context, input *pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error)
       GetCampaignAsync(ctx workflow.Context, input *pinpoint.GetCampaignInput) *PinpointGetCampaignResult

       GetCampaignActivities(ctx workflow.Context, input *pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error)
       GetCampaignActivitiesAsync(ctx workflow.Context, input *pinpoint.GetCampaignActivitiesInput) *PinpointGetCampaignActivitiesResult

       GetCampaignDateRangeKpi(ctx workflow.Context, input *pinpoint.GetCampaignDateRangeKpiInput) (*pinpoint.GetCampaignDateRangeKpiOutput, error)
       GetCampaignDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetCampaignDateRangeKpiInput) *PinpointGetCampaignDateRangeKpiResult

       GetCampaignVersion(ctx workflow.Context, input *pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error)
       GetCampaignVersionAsync(ctx workflow.Context, input *pinpoint.GetCampaignVersionInput) *PinpointGetCampaignVersionResult

       GetCampaignVersions(ctx workflow.Context, input *pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error)
       GetCampaignVersionsAsync(ctx workflow.Context, input *pinpoint.GetCampaignVersionsInput) *PinpointGetCampaignVersionsResult

       GetCampaigns(ctx workflow.Context, input *pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error)
       GetCampaignsAsync(ctx workflow.Context, input *pinpoint.GetCampaignsInput) *PinpointGetCampaignsResult

       GetChannels(ctx workflow.Context, input *pinpoint.GetChannelsInput) (*pinpoint.GetChannelsOutput, error)
       GetChannelsAsync(ctx workflow.Context, input *pinpoint.GetChannelsInput) *PinpointGetChannelsResult

       GetEmailChannel(ctx workflow.Context, input *pinpoint.GetEmailChannelInput) (*pinpoint.GetEmailChannelOutput, error)
       GetEmailChannelAsync(ctx workflow.Context, input *pinpoint.GetEmailChannelInput) *PinpointGetEmailChannelResult

       GetEmailTemplate(ctx workflow.Context, input *pinpoint.GetEmailTemplateInput) (*pinpoint.GetEmailTemplateOutput, error)
       GetEmailTemplateAsync(ctx workflow.Context, input *pinpoint.GetEmailTemplateInput) *PinpointGetEmailTemplateResult

       GetEndpoint(ctx workflow.Context, input *pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error)
       GetEndpointAsync(ctx workflow.Context, input *pinpoint.GetEndpointInput) *PinpointGetEndpointResult

       GetEventStream(ctx workflow.Context, input *pinpoint.GetEventStreamInput) (*pinpoint.GetEventStreamOutput, error)
       GetEventStreamAsync(ctx workflow.Context, input *pinpoint.GetEventStreamInput) *PinpointGetEventStreamResult

       GetExportJob(ctx workflow.Context, input *pinpoint.GetExportJobInput) (*pinpoint.GetExportJobOutput, error)
       GetExportJobAsync(ctx workflow.Context, input *pinpoint.GetExportJobInput) *PinpointGetExportJobResult

       GetExportJobs(ctx workflow.Context, input *pinpoint.GetExportJobsInput) (*pinpoint.GetExportJobsOutput, error)
       GetExportJobsAsync(ctx workflow.Context, input *pinpoint.GetExportJobsInput) *PinpointGetExportJobsResult

       GetGcmChannel(ctx workflow.Context, input *pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error)
       GetGcmChannelAsync(ctx workflow.Context, input *pinpoint.GetGcmChannelInput) *PinpointGetGcmChannelResult

       GetImportJob(ctx workflow.Context, input *pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error)
       GetImportJobAsync(ctx workflow.Context, input *pinpoint.GetImportJobInput) *PinpointGetImportJobResult

       GetImportJobs(ctx workflow.Context, input *pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error)
       GetImportJobsAsync(ctx workflow.Context, input *pinpoint.GetImportJobsInput) *PinpointGetImportJobsResult

       GetJourney(ctx workflow.Context, input *pinpoint.GetJourneyInput) (*pinpoint.GetJourneyOutput, error)
       GetJourneyAsync(ctx workflow.Context, input *pinpoint.GetJourneyInput) *PinpointGetJourneyResult

       GetJourneyDateRangeKpi(ctx workflow.Context, input *pinpoint.GetJourneyDateRangeKpiInput) (*pinpoint.GetJourneyDateRangeKpiOutput, error)
       GetJourneyDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetJourneyDateRangeKpiInput) *PinpointGetJourneyDateRangeKpiResult

       GetJourneyExecutionActivityMetrics(ctx workflow.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error)
       GetJourneyExecutionActivityMetricsAsync(ctx workflow.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput) *PinpointGetJourneyExecutionActivityMetricsResult

       GetJourneyExecutionMetrics(ctx workflow.Context, input *pinpoint.GetJourneyExecutionMetricsInput) (*pinpoint.GetJourneyExecutionMetricsOutput, error)
       GetJourneyExecutionMetricsAsync(ctx workflow.Context, input *pinpoint.GetJourneyExecutionMetricsInput) *PinpointGetJourneyExecutionMetricsResult

       GetPushTemplate(ctx workflow.Context, input *pinpoint.GetPushTemplateInput) (*pinpoint.GetPushTemplateOutput, error)
       GetPushTemplateAsync(ctx workflow.Context, input *pinpoint.GetPushTemplateInput) *PinpointGetPushTemplateResult

       GetRecommenderConfiguration(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationInput) (*pinpoint.GetRecommenderConfigurationOutput, error)
       GetRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationInput) *PinpointGetRecommenderConfigurationResult

       GetRecommenderConfigurations(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationsInput) (*pinpoint.GetRecommenderConfigurationsOutput, error)
       GetRecommenderConfigurationsAsync(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationsInput) *PinpointGetRecommenderConfigurationsResult

       GetSegment(ctx workflow.Context, input *pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error)
       GetSegmentAsync(ctx workflow.Context, input *pinpoint.GetSegmentInput) *PinpointGetSegmentResult

       GetSegmentExportJobs(ctx workflow.Context, input *pinpoint.GetSegmentExportJobsInput) (*pinpoint.GetSegmentExportJobsOutput, error)
       GetSegmentExportJobsAsync(ctx workflow.Context, input *pinpoint.GetSegmentExportJobsInput) *PinpointGetSegmentExportJobsResult

       GetSegmentImportJobs(ctx workflow.Context, input *pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error)
       GetSegmentImportJobsAsync(ctx workflow.Context, input *pinpoint.GetSegmentImportJobsInput) *PinpointGetSegmentImportJobsResult

       GetSegmentVersion(ctx workflow.Context, input *pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error)
       GetSegmentVersionAsync(ctx workflow.Context, input *pinpoint.GetSegmentVersionInput) *PinpointGetSegmentVersionResult

       GetSegmentVersions(ctx workflow.Context, input *pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error)
       GetSegmentVersionsAsync(ctx workflow.Context, input *pinpoint.GetSegmentVersionsInput) *PinpointGetSegmentVersionsResult

       GetSegments(ctx workflow.Context, input *pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error)
       GetSegmentsAsync(ctx workflow.Context, input *pinpoint.GetSegmentsInput) *PinpointGetSegmentsResult

       GetSmsChannel(ctx workflow.Context, input *pinpoint.GetSmsChannelInput) (*pinpoint.GetSmsChannelOutput, error)
       GetSmsChannelAsync(ctx workflow.Context, input *pinpoint.GetSmsChannelInput) *PinpointGetSmsChannelResult

       GetSmsTemplate(ctx workflow.Context, input *pinpoint.GetSmsTemplateInput) (*pinpoint.GetSmsTemplateOutput, error)
       GetSmsTemplateAsync(ctx workflow.Context, input *pinpoint.GetSmsTemplateInput) *PinpointGetSmsTemplateResult

       GetUserEndpoints(ctx workflow.Context, input *pinpoint.GetUserEndpointsInput) (*pinpoint.GetUserEndpointsOutput, error)
       GetUserEndpointsAsync(ctx workflow.Context, input *pinpoint.GetUserEndpointsInput) *PinpointGetUserEndpointsResult

       GetVoiceChannel(ctx workflow.Context, input *pinpoint.GetVoiceChannelInput) (*pinpoint.GetVoiceChannelOutput, error)
       GetVoiceChannelAsync(ctx workflow.Context, input *pinpoint.GetVoiceChannelInput) *PinpointGetVoiceChannelResult

       GetVoiceTemplate(ctx workflow.Context, input *pinpoint.GetVoiceTemplateInput) (*pinpoint.GetVoiceTemplateOutput, error)
       GetVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.GetVoiceTemplateInput) *PinpointGetVoiceTemplateResult

       ListJourneys(ctx workflow.Context, input *pinpoint.ListJourneysInput) (*pinpoint.ListJourneysOutput, error)
       ListJourneysAsync(ctx workflow.Context, input *pinpoint.ListJourneysInput) *PinpointListJourneysResult

       ListTagsForResource(ctx workflow.Context, input *pinpoint.ListTagsForResourceInput) (*pinpoint.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *pinpoint.ListTagsForResourceInput) *PinpointListTagsForResourceResult

       ListTemplateVersions(ctx workflow.Context, input *pinpoint.ListTemplateVersionsInput) (*pinpoint.ListTemplateVersionsOutput, error)
       ListTemplateVersionsAsync(ctx workflow.Context, input *pinpoint.ListTemplateVersionsInput) *PinpointListTemplateVersionsResult

       ListTemplates(ctx workflow.Context, input *pinpoint.ListTemplatesInput) (*pinpoint.ListTemplatesOutput, error)
       ListTemplatesAsync(ctx workflow.Context, input *pinpoint.ListTemplatesInput) *PinpointListTemplatesResult

       PhoneNumberValidate(ctx workflow.Context, input *pinpoint.PhoneNumberValidateInput) (*pinpoint.PhoneNumberValidateOutput, error)
       PhoneNumberValidateAsync(ctx workflow.Context, input *pinpoint.PhoneNumberValidateInput) *PinpointPhoneNumberValidateResult

       PutEventStream(ctx workflow.Context, input *pinpoint.PutEventStreamInput) (*pinpoint.PutEventStreamOutput, error)
       PutEventStreamAsync(ctx workflow.Context, input *pinpoint.PutEventStreamInput) *PinpointPutEventStreamResult

       PutEvents(ctx workflow.Context, input *pinpoint.PutEventsInput) (*pinpoint.PutEventsOutput, error)
       PutEventsAsync(ctx workflow.Context, input *pinpoint.PutEventsInput) *PinpointPutEventsResult

       RemoveAttributes(ctx workflow.Context, input *pinpoint.RemoveAttributesInput) (*pinpoint.RemoveAttributesOutput, error)
       RemoveAttributesAsync(ctx workflow.Context, input *pinpoint.RemoveAttributesInput) *PinpointRemoveAttributesResult

       SendMessages(ctx workflow.Context, input *pinpoint.SendMessagesInput) (*pinpoint.SendMessagesOutput, error)
       SendMessagesAsync(ctx workflow.Context, input *pinpoint.SendMessagesInput) *PinpointSendMessagesResult

       SendUsersMessages(ctx workflow.Context, input *pinpoint.SendUsersMessagesInput) (*pinpoint.SendUsersMessagesOutput, error)
       SendUsersMessagesAsync(ctx workflow.Context, input *pinpoint.SendUsersMessagesInput) *PinpointSendUsersMessagesResult

       TagResource(ctx workflow.Context, input *pinpoint.TagResourceInput) (*pinpoint.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *pinpoint.TagResourceInput) *PinpointTagResourceResult

       UntagResource(ctx workflow.Context, input *pinpoint.UntagResourceInput) (*pinpoint.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *pinpoint.UntagResourceInput) *PinpointUntagResourceResult

       UpdateAdmChannel(ctx workflow.Context, input *pinpoint.UpdateAdmChannelInput) (*pinpoint.UpdateAdmChannelOutput, error)
       UpdateAdmChannelAsync(ctx workflow.Context, input *pinpoint.UpdateAdmChannelInput) *PinpointUpdateAdmChannelResult

       UpdateApnsChannel(ctx workflow.Context, input *pinpoint.UpdateApnsChannelInput) (*pinpoint.UpdateApnsChannelOutput, error)
       UpdateApnsChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsChannelInput) *PinpointUpdateApnsChannelResult

       UpdateApnsSandboxChannel(ctx workflow.Context, input *pinpoint.UpdateApnsSandboxChannelInput) (*pinpoint.UpdateApnsSandboxChannelOutput, error)
       UpdateApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsSandboxChannelInput) *PinpointUpdateApnsSandboxChannelResult

       UpdateApnsVoipChannel(ctx workflow.Context, input *pinpoint.UpdateApnsVoipChannelInput) (*pinpoint.UpdateApnsVoipChannelOutput, error)
       UpdateApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsVoipChannelInput) *PinpointUpdateApnsVoipChannelResult

       UpdateApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.UpdateApnsVoipSandboxChannelInput) (*pinpoint.UpdateApnsVoipSandboxChannelOutput, error)
       UpdateApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsVoipSandboxChannelInput) *PinpointUpdateApnsVoipSandboxChannelResult

       UpdateApplicationSettings(ctx workflow.Context, input *pinpoint.UpdateApplicationSettingsInput) (*pinpoint.UpdateApplicationSettingsOutput, error)
       UpdateApplicationSettingsAsync(ctx workflow.Context, input *pinpoint.UpdateApplicationSettingsInput) *PinpointUpdateApplicationSettingsResult

       UpdateBaiduChannel(ctx workflow.Context, input *pinpoint.UpdateBaiduChannelInput) (*pinpoint.UpdateBaiduChannelOutput, error)
       UpdateBaiduChannelAsync(ctx workflow.Context, input *pinpoint.UpdateBaiduChannelInput) *PinpointUpdateBaiduChannelResult

       UpdateCampaign(ctx workflow.Context, input *pinpoint.UpdateCampaignInput) (*pinpoint.UpdateCampaignOutput, error)
       UpdateCampaignAsync(ctx workflow.Context, input *pinpoint.UpdateCampaignInput) *PinpointUpdateCampaignResult

       UpdateEmailChannel(ctx workflow.Context, input *pinpoint.UpdateEmailChannelInput) (*pinpoint.UpdateEmailChannelOutput, error)
       UpdateEmailChannelAsync(ctx workflow.Context, input *pinpoint.UpdateEmailChannelInput) *PinpointUpdateEmailChannelResult

       UpdateEmailTemplate(ctx workflow.Context, input *pinpoint.UpdateEmailTemplateInput) (*pinpoint.UpdateEmailTemplateOutput, error)
       UpdateEmailTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateEmailTemplateInput) *PinpointUpdateEmailTemplateResult

       UpdateEndpoint(ctx workflow.Context, input *pinpoint.UpdateEndpointInput) (*pinpoint.UpdateEndpointOutput, error)
       UpdateEndpointAsync(ctx workflow.Context, input *pinpoint.UpdateEndpointInput) *PinpointUpdateEndpointResult

       UpdateEndpointsBatch(ctx workflow.Context, input *pinpoint.UpdateEndpointsBatchInput) (*pinpoint.UpdateEndpointsBatchOutput, error)
       UpdateEndpointsBatchAsync(ctx workflow.Context, input *pinpoint.UpdateEndpointsBatchInput) *PinpointUpdateEndpointsBatchResult

       UpdateGcmChannel(ctx workflow.Context, input *pinpoint.UpdateGcmChannelInput) (*pinpoint.UpdateGcmChannelOutput, error)
       UpdateGcmChannelAsync(ctx workflow.Context, input *pinpoint.UpdateGcmChannelInput) *PinpointUpdateGcmChannelResult

       UpdateJourney(ctx workflow.Context, input *pinpoint.UpdateJourneyInput) (*pinpoint.UpdateJourneyOutput, error)
       UpdateJourneyAsync(ctx workflow.Context, input *pinpoint.UpdateJourneyInput) *PinpointUpdateJourneyResult

       UpdateJourneyState(ctx workflow.Context, input *pinpoint.UpdateJourneyStateInput) (*pinpoint.UpdateJourneyStateOutput, error)
       UpdateJourneyStateAsync(ctx workflow.Context, input *pinpoint.UpdateJourneyStateInput) *PinpointUpdateJourneyStateResult

       UpdatePushTemplate(ctx workflow.Context, input *pinpoint.UpdatePushTemplateInput) (*pinpoint.UpdatePushTemplateOutput, error)
       UpdatePushTemplateAsync(ctx workflow.Context, input *pinpoint.UpdatePushTemplateInput) *PinpointUpdatePushTemplateResult

       UpdateRecommenderConfiguration(ctx workflow.Context, input *pinpoint.UpdateRecommenderConfigurationInput) (*pinpoint.UpdateRecommenderConfigurationOutput, error)
       UpdateRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.UpdateRecommenderConfigurationInput) *PinpointUpdateRecommenderConfigurationResult

       UpdateSegment(ctx workflow.Context, input *pinpoint.UpdateSegmentInput) (*pinpoint.UpdateSegmentOutput, error)
       UpdateSegmentAsync(ctx workflow.Context, input *pinpoint.UpdateSegmentInput) *PinpointUpdateSegmentResult

       UpdateSmsChannel(ctx workflow.Context, input *pinpoint.UpdateSmsChannelInput) (*pinpoint.UpdateSmsChannelOutput, error)
       UpdateSmsChannelAsync(ctx workflow.Context, input *pinpoint.UpdateSmsChannelInput) *PinpointUpdateSmsChannelResult

       UpdateSmsTemplate(ctx workflow.Context, input *pinpoint.UpdateSmsTemplateInput) (*pinpoint.UpdateSmsTemplateOutput, error)
       UpdateSmsTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateSmsTemplateInput) *PinpointUpdateSmsTemplateResult

       UpdateTemplateActiveVersion(ctx workflow.Context, input *pinpoint.UpdateTemplateActiveVersionInput) (*pinpoint.UpdateTemplateActiveVersionOutput, error)
       UpdateTemplateActiveVersionAsync(ctx workflow.Context, input *pinpoint.UpdateTemplateActiveVersionInput) *PinpointUpdateTemplateActiveVersionResult

       UpdateVoiceChannel(ctx workflow.Context, input *pinpoint.UpdateVoiceChannelInput) (*pinpoint.UpdateVoiceChannelOutput, error)
       UpdateVoiceChannelAsync(ctx workflow.Context, input *pinpoint.UpdateVoiceChannelInput) *PinpointUpdateVoiceChannelResult

       UpdateVoiceTemplate(ctx workflow.Context, input *pinpoint.UpdateVoiceTemplateInput) (*pinpoint.UpdateVoiceTemplateOutput, error)
       UpdateVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateVoiceTemplateInput) *PinpointUpdateVoiceTemplateResult
}

type PinpointCreateAppResult struct {
	Result workflow.Future
}

func (r *PinpointCreateAppResult) Get(ctx workflow.Context) (*pinpoint.CreateAppOutput, error) {
    var output pinpoint.CreateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateCampaignResult struct {
	Result workflow.Future
}

func (r *PinpointCreateCampaignResult) Get(ctx workflow.Context) (*pinpoint.CreateCampaignOutput, error) {
    var output pinpoint.CreateCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateEmailTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointCreateEmailTemplateResult) Get(ctx workflow.Context) (*pinpoint.CreateEmailTemplateOutput, error) {
    var output pinpoint.CreateEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateExportJobResult struct {
	Result workflow.Future
}

func (r *PinpointCreateExportJobResult) Get(ctx workflow.Context) (*pinpoint.CreateExportJobOutput, error) {
    var output pinpoint.CreateExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateImportJobResult struct {
	Result workflow.Future
}

func (r *PinpointCreateImportJobResult) Get(ctx workflow.Context) (*pinpoint.CreateImportJobOutput, error) {
    var output pinpoint.CreateImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateJourneyResult struct {
	Result workflow.Future
}

func (r *PinpointCreateJourneyResult) Get(ctx workflow.Context) (*pinpoint.CreateJourneyOutput, error) {
    var output pinpoint.CreateJourneyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreatePushTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointCreatePushTemplateResult) Get(ctx workflow.Context) (*pinpoint.CreatePushTemplateOutput, error) {
    var output pinpoint.CreatePushTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateRecommenderConfigurationResult struct {
	Result workflow.Future
}

func (r *PinpointCreateRecommenderConfigurationResult) Get(ctx workflow.Context) (*pinpoint.CreateRecommenderConfigurationOutput, error) {
    var output pinpoint.CreateRecommenderConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateSegmentResult struct {
	Result workflow.Future
}

func (r *PinpointCreateSegmentResult) Get(ctx workflow.Context) (*pinpoint.CreateSegmentOutput, error) {
    var output pinpoint.CreateSegmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateSmsTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointCreateSmsTemplateResult) Get(ctx workflow.Context) (*pinpoint.CreateSmsTemplateOutput, error) {
    var output pinpoint.CreateSmsTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointCreateVoiceTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointCreateVoiceTemplateResult) Get(ctx workflow.Context) (*pinpoint.CreateVoiceTemplateOutput, error) {
    var output pinpoint.CreateVoiceTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteAdmChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteAdmChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteAdmChannelOutput, error) {
    var output pinpoint.DeleteAdmChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteApnsChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteApnsChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteApnsChannelOutput, error) {
    var output pinpoint.DeleteApnsChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteApnsSandboxChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteApnsSandboxChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteApnsSandboxChannelOutput, error) {
    var output pinpoint.DeleteApnsSandboxChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteApnsVoipChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteApnsVoipChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteApnsVoipChannelOutput, error) {
    var output pinpoint.DeleteApnsVoipChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteApnsVoipSandboxChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteApnsVoipSandboxChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteApnsVoipSandboxChannelOutput, error) {
    var output pinpoint.DeleteApnsVoipSandboxChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteAppResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteAppResult) Get(ctx workflow.Context) (*pinpoint.DeleteAppOutput, error) {
    var output pinpoint.DeleteAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteBaiduChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteBaiduChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteBaiduChannelOutput, error) {
    var output pinpoint.DeleteBaiduChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteCampaignResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteCampaignResult) Get(ctx workflow.Context) (*pinpoint.DeleteCampaignOutput, error) {
    var output pinpoint.DeleteCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteEmailChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteEmailChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteEmailChannelOutput, error) {
    var output pinpoint.DeleteEmailChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteEmailTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteEmailTemplateResult) Get(ctx workflow.Context) (*pinpoint.DeleteEmailTemplateOutput, error) {
    var output pinpoint.DeleteEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteEndpointResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteEndpointResult) Get(ctx workflow.Context) (*pinpoint.DeleteEndpointOutput, error) {
    var output pinpoint.DeleteEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteEventStreamResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteEventStreamResult) Get(ctx workflow.Context) (*pinpoint.DeleteEventStreamOutput, error) {
    var output pinpoint.DeleteEventStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteGcmChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteGcmChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteGcmChannelOutput, error) {
    var output pinpoint.DeleteGcmChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteJourneyResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteJourneyResult) Get(ctx workflow.Context) (*pinpoint.DeleteJourneyOutput, error) {
    var output pinpoint.DeleteJourneyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeletePushTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointDeletePushTemplateResult) Get(ctx workflow.Context) (*pinpoint.DeletePushTemplateOutput, error) {
    var output pinpoint.DeletePushTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteRecommenderConfigurationResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteRecommenderConfigurationResult) Get(ctx workflow.Context) (*pinpoint.DeleteRecommenderConfigurationOutput, error) {
    var output pinpoint.DeleteRecommenderConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteSegmentResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteSegmentResult) Get(ctx workflow.Context) (*pinpoint.DeleteSegmentOutput, error) {
    var output pinpoint.DeleteSegmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteSmsChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteSmsChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteSmsChannelOutput, error) {
    var output pinpoint.DeleteSmsChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteSmsTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteSmsTemplateResult) Get(ctx workflow.Context) (*pinpoint.DeleteSmsTemplateOutput, error) {
    var output pinpoint.DeleteSmsTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteUserEndpointsResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteUserEndpointsResult) Get(ctx workflow.Context) (*pinpoint.DeleteUserEndpointsOutput, error) {
    var output pinpoint.DeleteUserEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteVoiceChannelResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteVoiceChannelResult) Get(ctx workflow.Context) (*pinpoint.DeleteVoiceChannelOutput, error) {
    var output pinpoint.DeleteVoiceChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointDeleteVoiceTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointDeleteVoiceTemplateResult) Get(ctx workflow.Context) (*pinpoint.DeleteVoiceTemplateOutput, error) {
    var output pinpoint.DeleteVoiceTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetAdmChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetAdmChannelResult) Get(ctx workflow.Context) (*pinpoint.GetAdmChannelOutput, error) {
    var output pinpoint.GetAdmChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetApnsChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetApnsChannelResult) Get(ctx workflow.Context) (*pinpoint.GetApnsChannelOutput, error) {
    var output pinpoint.GetApnsChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetApnsSandboxChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetApnsSandboxChannelResult) Get(ctx workflow.Context) (*pinpoint.GetApnsSandboxChannelOutput, error) {
    var output pinpoint.GetApnsSandboxChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetApnsVoipChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetApnsVoipChannelResult) Get(ctx workflow.Context) (*pinpoint.GetApnsVoipChannelOutput, error) {
    var output pinpoint.GetApnsVoipChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetApnsVoipSandboxChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetApnsVoipSandboxChannelResult) Get(ctx workflow.Context) (*pinpoint.GetApnsVoipSandboxChannelOutput, error) {
    var output pinpoint.GetApnsVoipSandboxChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetAppResult struct {
	Result workflow.Future
}

func (r *PinpointGetAppResult) Get(ctx workflow.Context) (*pinpoint.GetAppOutput, error) {
    var output pinpoint.GetAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetApplicationDateRangeKpiResult struct {
	Result workflow.Future
}

func (r *PinpointGetApplicationDateRangeKpiResult) Get(ctx workflow.Context) (*pinpoint.GetApplicationDateRangeKpiOutput, error) {
    var output pinpoint.GetApplicationDateRangeKpiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetApplicationSettingsResult struct {
	Result workflow.Future
}

func (r *PinpointGetApplicationSettingsResult) Get(ctx workflow.Context) (*pinpoint.GetApplicationSettingsOutput, error) {
    var output pinpoint.GetApplicationSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetAppsResult struct {
	Result workflow.Future
}

func (r *PinpointGetAppsResult) Get(ctx workflow.Context) (*pinpoint.GetAppsOutput, error) {
    var output pinpoint.GetAppsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetBaiduChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetBaiduChannelResult) Get(ctx workflow.Context) (*pinpoint.GetBaiduChannelOutput, error) {
    var output pinpoint.GetBaiduChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetCampaignResult struct {
	Result workflow.Future
}

func (r *PinpointGetCampaignResult) Get(ctx workflow.Context) (*pinpoint.GetCampaignOutput, error) {
    var output pinpoint.GetCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetCampaignActivitiesResult struct {
	Result workflow.Future
}

func (r *PinpointGetCampaignActivitiesResult) Get(ctx workflow.Context) (*pinpoint.GetCampaignActivitiesOutput, error) {
    var output pinpoint.GetCampaignActivitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetCampaignDateRangeKpiResult struct {
	Result workflow.Future
}

func (r *PinpointGetCampaignDateRangeKpiResult) Get(ctx workflow.Context) (*pinpoint.GetCampaignDateRangeKpiOutput, error) {
    var output pinpoint.GetCampaignDateRangeKpiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetCampaignVersionResult struct {
	Result workflow.Future
}

func (r *PinpointGetCampaignVersionResult) Get(ctx workflow.Context) (*pinpoint.GetCampaignVersionOutput, error) {
    var output pinpoint.GetCampaignVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetCampaignVersionsResult struct {
	Result workflow.Future
}

func (r *PinpointGetCampaignVersionsResult) Get(ctx workflow.Context) (*pinpoint.GetCampaignVersionsOutput, error) {
    var output pinpoint.GetCampaignVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetCampaignsResult struct {
	Result workflow.Future
}

func (r *PinpointGetCampaignsResult) Get(ctx workflow.Context) (*pinpoint.GetCampaignsOutput, error) {
    var output pinpoint.GetCampaignsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetChannelsResult struct {
	Result workflow.Future
}

func (r *PinpointGetChannelsResult) Get(ctx workflow.Context) (*pinpoint.GetChannelsOutput, error) {
    var output pinpoint.GetChannelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetEmailChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetEmailChannelResult) Get(ctx workflow.Context) (*pinpoint.GetEmailChannelOutput, error) {
    var output pinpoint.GetEmailChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetEmailTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointGetEmailTemplateResult) Get(ctx workflow.Context) (*pinpoint.GetEmailTemplateOutput, error) {
    var output pinpoint.GetEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetEndpointResult struct {
	Result workflow.Future
}

func (r *PinpointGetEndpointResult) Get(ctx workflow.Context) (*pinpoint.GetEndpointOutput, error) {
    var output pinpoint.GetEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetEventStreamResult struct {
	Result workflow.Future
}

func (r *PinpointGetEventStreamResult) Get(ctx workflow.Context) (*pinpoint.GetEventStreamOutput, error) {
    var output pinpoint.GetEventStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetExportJobResult struct {
	Result workflow.Future
}

func (r *PinpointGetExportJobResult) Get(ctx workflow.Context) (*pinpoint.GetExportJobOutput, error) {
    var output pinpoint.GetExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetExportJobsResult struct {
	Result workflow.Future
}

func (r *PinpointGetExportJobsResult) Get(ctx workflow.Context) (*pinpoint.GetExportJobsOutput, error) {
    var output pinpoint.GetExportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetGcmChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetGcmChannelResult) Get(ctx workflow.Context) (*pinpoint.GetGcmChannelOutput, error) {
    var output pinpoint.GetGcmChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetImportJobResult struct {
	Result workflow.Future
}

func (r *PinpointGetImportJobResult) Get(ctx workflow.Context) (*pinpoint.GetImportJobOutput, error) {
    var output pinpoint.GetImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetImportJobsResult struct {
	Result workflow.Future
}

func (r *PinpointGetImportJobsResult) Get(ctx workflow.Context) (*pinpoint.GetImportJobsOutput, error) {
    var output pinpoint.GetImportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetJourneyResult struct {
	Result workflow.Future
}

func (r *PinpointGetJourneyResult) Get(ctx workflow.Context) (*pinpoint.GetJourneyOutput, error) {
    var output pinpoint.GetJourneyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetJourneyDateRangeKpiResult struct {
	Result workflow.Future
}

func (r *PinpointGetJourneyDateRangeKpiResult) Get(ctx workflow.Context) (*pinpoint.GetJourneyDateRangeKpiOutput, error) {
    var output pinpoint.GetJourneyDateRangeKpiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetJourneyExecutionActivityMetricsResult struct {
	Result workflow.Future
}

func (r *PinpointGetJourneyExecutionActivityMetricsResult) Get(ctx workflow.Context) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error) {
    var output pinpoint.GetJourneyExecutionActivityMetricsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetJourneyExecutionMetricsResult struct {
	Result workflow.Future
}

func (r *PinpointGetJourneyExecutionMetricsResult) Get(ctx workflow.Context) (*pinpoint.GetJourneyExecutionMetricsOutput, error) {
    var output pinpoint.GetJourneyExecutionMetricsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetPushTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointGetPushTemplateResult) Get(ctx workflow.Context) (*pinpoint.GetPushTemplateOutput, error) {
    var output pinpoint.GetPushTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetRecommenderConfigurationResult struct {
	Result workflow.Future
}

func (r *PinpointGetRecommenderConfigurationResult) Get(ctx workflow.Context) (*pinpoint.GetRecommenderConfigurationOutput, error) {
    var output pinpoint.GetRecommenderConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetRecommenderConfigurationsResult struct {
	Result workflow.Future
}

func (r *PinpointGetRecommenderConfigurationsResult) Get(ctx workflow.Context) (*pinpoint.GetRecommenderConfigurationsOutput, error) {
    var output pinpoint.GetRecommenderConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSegmentResult struct {
	Result workflow.Future
}

func (r *PinpointGetSegmentResult) Get(ctx workflow.Context) (*pinpoint.GetSegmentOutput, error) {
    var output pinpoint.GetSegmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSegmentExportJobsResult struct {
	Result workflow.Future
}

func (r *PinpointGetSegmentExportJobsResult) Get(ctx workflow.Context) (*pinpoint.GetSegmentExportJobsOutput, error) {
    var output pinpoint.GetSegmentExportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSegmentImportJobsResult struct {
	Result workflow.Future
}

func (r *PinpointGetSegmentImportJobsResult) Get(ctx workflow.Context) (*pinpoint.GetSegmentImportJobsOutput, error) {
    var output pinpoint.GetSegmentImportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSegmentVersionResult struct {
	Result workflow.Future
}

func (r *PinpointGetSegmentVersionResult) Get(ctx workflow.Context) (*pinpoint.GetSegmentVersionOutput, error) {
    var output pinpoint.GetSegmentVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSegmentVersionsResult struct {
	Result workflow.Future
}

func (r *PinpointGetSegmentVersionsResult) Get(ctx workflow.Context) (*pinpoint.GetSegmentVersionsOutput, error) {
    var output pinpoint.GetSegmentVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSegmentsResult struct {
	Result workflow.Future
}

func (r *PinpointGetSegmentsResult) Get(ctx workflow.Context) (*pinpoint.GetSegmentsOutput, error) {
    var output pinpoint.GetSegmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSmsChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetSmsChannelResult) Get(ctx workflow.Context) (*pinpoint.GetSmsChannelOutput, error) {
    var output pinpoint.GetSmsChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetSmsTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointGetSmsTemplateResult) Get(ctx workflow.Context) (*pinpoint.GetSmsTemplateOutput, error) {
    var output pinpoint.GetSmsTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetUserEndpointsResult struct {
	Result workflow.Future
}

func (r *PinpointGetUserEndpointsResult) Get(ctx workflow.Context) (*pinpoint.GetUserEndpointsOutput, error) {
    var output pinpoint.GetUserEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetVoiceChannelResult struct {
	Result workflow.Future
}

func (r *PinpointGetVoiceChannelResult) Get(ctx workflow.Context) (*pinpoint.GetVoiceChannelOutput, error) {
    var output pinpoint.GetVoiceChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointGetVoiceTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointGetVoiceTemplateResult) Get(ctx workflow.Context) (*pinpoint.GetVoiceTemplateOutput, error) {
    var output pinpoint.GetVoiceTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointListJourneysResult struct {
	Result workflow.Future
}

func (r *PinpointListJourneysResult) Get(ctx workflow.Context) (*pinpoint.ListJourneysOutput, error) {
    var output pinpoint.ListJourneysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *PinpointListTagsForResourceResult) Get(ctx workflow.Context) (*pinpoint.ListTagsForResourceOutput, error) {
    var output pinpoint.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointListTemplateVersionsResult struct {
	Result workflow.Future
}

func (r *PinpointListTemplateVersionsResult) Get(ctx workflow.Context) (*pinpoint.ListTemplateVersionsOutput, error) {
    var output pinpoint.ListTemplateVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointListTemplatesResult struct {
	Result workflow.Future
}

func (r *PinpointListTemplatesResult) Get(ctx workflow.Context) (*pinpoint.ListTemplatesOutput, error) {
    var output pinpoint.ListTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointPhoneNumberValidateResult struct {
	Result workflow.Future
}

func (r *PinpointPhoneNumberValidateResult) Get(ctx workflow.Context) (*pinpoint.PhoneNumberValidateOutput, error) {
    var output pinpoint.PhoneNumberValidateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointPutEventStreamResult struct {
	Result workflow.Future
}

func (r *PinpointPutEventStreamResult) Get(ctx workflow.Context) (*pinpoint.PutEventStreamOutput, error) {
    var output pinpoint.PutEventStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointPutEventsResult struct {
	Result workflow.Future
}

func (r *PinpointPutEventsResult) Get(ctx workflow.Context) (*pinpoint.PutEventsOutput, error) {
    var output pinpoint.PutEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointRemoveAttributesResult struct {
	Result workflow.Future
}

func (r *PinpointRemoveAttributesResult) Get(ctx workflow.Context) (*pinpoint.RemoveAttributesOutput, error) {
    var output pinpoint.RemoveAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointSendMessagesResult struct {
	Result workflow.Future
}

func (r *PinpointSendMessagesResult) Get(ctx workflow.Context) (*pinpoint.SendMessagesOutput, error) {
    var output pinpoint.SendMessagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointSendUsersMessagesResult struct {
	Result workflow.Future
}

func (r *PinpointSendUsersMessagesResult) Get(ctx workflow.Context) (*pinpoint.SendUsersMessagesOutput, error) {
    var output pinpoint.SendUsersMessagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointTagResourceResult struct {
	Result workflow.Future
}

func (r *PinpointTagResourceResult) Get(ctx workflow.Context) (*pinpoint.TagResourceOutput, error) {
    var output pinpoint.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUntagResourceResult struct {
	Result workflow.Future
}

func (r *PinpointUntagResourceResult) Get(ctx workflow.Context) (*pinpoint.UntagResourceOutput, error) {
    var output pinpoint.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateAdmChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateAdmChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateAdmChannelOutput, error) {
    var output pinpoint.UpdateAdmChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateApnsChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateApnsChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateApnsChannelOutput, error) {
    var output pinpoint.UpdateApnsChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateApnsSandboxChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateApnsSandboxChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateApnsSandboxChannelOutput, error) {
    var output pinpoint.UpdateApnsSandboxChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateApnsVoipChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateApnsVoipChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateApnsVoipChannelOutput, error) {
    var output pinpoint.UpdateApnsVoipChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateApnsVoipSandboxChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateApnsVoipSandboxChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateApnsVoipSandboxChannelOutput, error) {
    var output pinpoint.UpdateApnsVoipSandboxChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateApplicationSettingsResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateApplicationSettingsResult) Get(ctx workflow.Context) (*pinpoint.UpdateApplicationSettingsOutput, error) {
    var output pinpoint.UpdateApplicationSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateBaiduChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateBaiduChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateBaiduChannelOutput, error) {
    var output pinpoint.UpdateBaiduChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateCampaignResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateCampaignResult) Get(ctx workflow.Context) (*pinpoint.UpdateCampaignOutput, error) {
    var output pinpoint.UpdateCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateEmailChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateEmailChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateEmailChannelOutput, error) {
    var output pinpoint.UpdateEmailChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateEmailTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateEmailTemplateResult) Get(ctx workflow.Context) (*pinpoint.UpdateEmailTemplateOutput, error) {
    var output pinpoint.UpdateEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateEndpointResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateEndpointResult) Get(ctx workflow.Context) (*pinpoint.UpdateEndpointOutput, error) {
    var output pinpoint.UpdateEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateEndpointsBatchResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateEndpointsBatchResult) Get(ctx workflow.Context) (*pinpoint.UpdateEndpointsBatchOutput, error) {
    var output pinpoint.UpdateEndpointsBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateGcmChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateGcmChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateGcmChannelOutput, error) {
    var output pinpoint.UpdateGcmChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateJourneyResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateJourneyResult) Get(ctx workflow.Context) (*pinpoint.UpdateJourneyOutput, error) {
    var output pinpoint.UpdateJourneyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateJourneyStateResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateJourneyStateResult) Get(ctx workflow.Context) (*pinpoint.UpdateJourneyStateOutput, error) {
    var output pinpoint.UpdateJourneyStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdatePushTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointUpdatePushTemplateResult) Get(ctx workflow.Context) (*pinpoint.UpdatePushTemplateOutput, error) {
    var output pinpoint.UpdatePushTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateRecommenderConfigurationResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateRecommenderConfigurationResult) Get(ctx workflow.Context) (*pinpoint.UpdateRecommenderConfigurationOutput, error) {
    var output pinpoint.UpdateRecommenderConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateSegmentResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateSegmentResult) Get(ctx workflow.Context) (*pinpoint.UpdateSegmentOutput, error) {
    var output pinpoint.UpdateSegmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateSmsChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateSmsChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateSmsChannelOutput, error) {
    var output pinpoint.UpdateSmsChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateSmsTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateSmsTemplateResult) Get(ctx workflow.Context) (*pinpoint.UpdateSmsTemplateOutput, error) {
    var output pinpoint.UpdateSmsTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateTemplateActiveVersionResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateTemplateActiveVersionResult) Get(ctx workflow.Context) (*pinpoint.UpdateTemplateActiveVersionOutput, error) {
    var output pinpoint.UpdateTemplateActiveVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateVoiceChannelResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateVoiceChannelResult) Get(ctx workflow.Context) (*pinpoint.UpdateVoiceChannelOutput, error) {
    var output pinpoint.UpdateVoiceChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointUpdateVoiceTemplateResult struct {
	Result workflow.Future
}

func (r *PinpointUpdateVoiceTemplateResult) Get(ctx workflow.Context) (*pinpoint.UpdateVoiceTemplateOutput, error) {
    var output pinpoint.UpdateVoiceTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PinpointStub struct {
    activities awsactivities.PinpointActivities
}

func NewPinpointStub() PinpointClient {
    return &PinpointStub{}
}

func (a *PinpointStub) CreateApp(ctx workflow.Context, input *pinpoint.CreateAppInput) (*pinpoint.CreateAppOutput, error) {
    var output pinpoint.CreateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateAppAsync(ctx workflow.Context, input *pinpoint.CreateAppInput) *PinpointCreateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input)
    return &PinpointCreateAppResult{Result: future}
}

func (a *PinpointStub) CreateCampaign(ctx workflow.Context, input *pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error) {
    var output pinpoint.CreateCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateCampaignAsync(ctx workflow.Context, input *pinpoint.CreateCampaignInput) *PinpointCreateCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCampaign, input)
    return &PinpointCreateCampaignResult{Result: future}
}

func (a *PinpointStub) CreateEmailTemplate(ctx workflow.Context, input *pinpoint.CreateEmailTemplateInput) (*pinpoint.CreateEmailTemplateOutput, error) {
    var output pinpoint.CreateEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateEmailTemplateAsync(ctx workflow.Context, input *pinpoint.CreateEmailTemplateInput) *PinpointCreateEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEmailTemplate, input)
    return &PinpointCreateEmailTemplateResult{Result: future}
}

func (a *PinpointStub) CreateExportJob(ctx workflow.Context, input *pinpoint.CreateExportJobInput) (*pinpoint.CreateExportJobOutput, error) {
    var output pinpoint.CreateExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateExportJobAsync(ctx workflow.Context, input *pinpoint.CreateExportJobInput) *PinpointCreateExportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateExportJob, input)
    return &PinpointCreateExportJobResult{Result: future}
}

func (a *PinpointStub) CreateImportJob(ctx workflow.Context, input *pinpoint.CreateImportJobInput) (*pinpoint.CreateImportJobOutput, error) {
    var output pinpoint.CreateImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateImportJobAsync(ctx workflow.Context, input *pinpoint.CreateImportJobInput) *PinpointCreateImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateImportJob, input)
    return &PinpointCreateImportJobResult{Result: future}
}

func (a *PinpointStub) CreateJourney(ctx workflow.Context, input *pinpoint.CreateJourneyInput) (*pinpoint.CreateJourneyOutput, error) {
    var output pinpoint.CreateJourneyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJourney, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateJourneyAsync(ctx workflow.Context, input *pinpoint.CreateJourneyInput) *PinpointCreateJourneyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJourney, input)
    return &PinpointCreateJourneyResult{Result: future}
}

func (a *PinpointStub) CreatePushTemplate(ctx workflow.Context, input *pinpoint.CreatePushTemplateInput) (*pinpoint.CreatePushTemplateOutput, error) {
    var output pinpoint.CreatePushTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePushTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreatePushTemplateAsync(ctx workflow.Context, input *pinpoint.CreatePushTemplateInput) *PinpointCreatePushTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePushTemplate, input)
    return &PinpointCreatePushTemplateResult{Result: future}
}

func (a *PinpointStub) CreateRecommenderConfiguration(ctx workflow.Context, input *pinpoint.CreateRecommenderConfigurationInput) (*pinpoint.CreateRecommenderConfigurationOutput, error) {
    var output pinpoint.CreateRecommenderConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRecommenderConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.CreateRecommenderConfigurationInput) *PinpointCreateRecommenderConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRecommenderConfiguration, input)
    return &PinpointCreateRecommenderConfigurationResult{Result: future}
}

func (a *PinpointStub) CreateSegment(ctx workflow.Context, input *pinpoint.CreateSegmentInput) (*pinpoint.CreateSegmentOutput, error) {
    var output pinpoint.CreateSegmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSegment, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateSegmentAsync(ctx workflow.Context, input *pinpoint.CreateSegmentInput) *PinpointCreateSegmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSegment, input)
    return &PinpointCreateSegmentResult{Result: future}
}

func (a *PinpointStub) CreateSmsTemplate(ctx workflow.Context, input *pinpoint.CreateSmsTemplateInput) (*pinpoint.CreateSmsTemplateOutput, error) {
    var output pinpoint.CreateSmsTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSmsTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateSmsTemplateAsync(ctx workflow.Context, input *pinpoint.CreateSmsTemplateInput) *PinpointCreateSmsTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSmsTemplate, input)
    return &PinpointCreateSmsTemplateResult{Result: future}
}

func (a *PinpointStub) CreateVoiceTemplate(ctx workflow.Context, input *pinpoint.CreateVoiceTemplateInput) (*pinpoint.CreateVoiceTemplateOutput, error) {
    var output pinpoint.CreateVoiceTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVoiceTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) CreateVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.CreateVoiceTemplateInput) *PinpointCreateVoiceTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVoiceTemplate, input)
    return &PinpointCreateVoiceTemplateResult{Result: future}
}

func (a *PinpointStub) DeleteAdmChannel(ctx workflow.Context, input *pinpoint.DeleteAdmChannelInput) (*pinpoint.DeleteAdmChannelOutput, error) {
    var output pinpoint.DeleteAdmChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAdmChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteAdmChannelAsync(ctx workflow.Context, input *pinpoint.DeleteAdmChannelInput) *PinpointDeleteAdmChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAdmChannel, input)
    return &PinpointDeleteAdmChannelResult{Result: future}
}

func (a *PinpointStub) DeleteApnsChannel(ctx workflow.Context, input *pinpoint.DeleteApnsChannelInput) (*pinpoint.DeleteApnsChannelOutput, error) {
    var output pinpoint.DeleteApnsChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteApnsChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsChannelInput) *PinpointDeleteApnsChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsChannel, input)
    return &PinpointDeleteApnsChannelResult{Result: future}
}

func (a *PinpointStub) DeleteApnsSandboxChannel(ctx workflow.Context, input *pinpoint.DeleteApnsSandboxChannelInput) (*pinpoint.DeleteApnsSandboxChannelOutput, error) {
    var output pinpoint.DeleteApnsSandboxChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsSandboxChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsSandboxChannelInput) *PinpointDeleteApnsSandboxChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsSandboxChannel, input)
    return &PinpointDeleteApnsSandboxChannelResult{Result: future}
}

func (a *PinpointStub) DeleteApnsVoipChannel(ctx workflow.Context, input *pinpoint.DeleteApnsVoipChannelInput) (*pinpoint.DeleteApnsVoipChannelOutput, error) {
    var output pinpoint.DeleteApnsVoipChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsVoipChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsVoipChannelInput) *PinpointDeleteApnsVoipChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsVoipChannel, input)
    return &PinpointDeleteApnsVoipChannelResult{Result: future}
}

func (a *PinpointStub) DeleteApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.DeleteApnsVoipSandboxChannelInput) (*pinpoint.DeleteApnsVoipSandboxChannelOutput, error) {
    var output pinpoint.DeleteApnsVoipSandboxChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsVoipSandboxChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsVoipSandboxChannelInput) *PinpointDeleteApnsVoipSandboxChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApnsVoipSandboxChannel, input)
    return &PinpointDeleteApnsVoipSandboxChannelResult{Result: future}
}

func (a *PinpointStub) DeleteApp(ctx workflow.Context, input *pinpoint.DeleteAppInput) (*pinpoint.DeleteAppOutput, error) {
    var output pinpoint.DeleteAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteAppAsync(ctx workflow.Context, input *pinpoint.DeleteAppInput) *PinpointDeleteAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input)
    return &PinpointDeleteAppResult{Result: future}
}

func (a *PinpointStub) DeleteBaiduChannel(ctx workflow.Context, input *pinpoint.DeleteBaiduChannelInput) (*pinpoint.DeleteBaiduChannelOutput, error) {
    var output pinpoint.DeleteBaiduChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBaiduChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteBaiduChannelAsync(ctx workflow.Context, input *pinpoint.DeleteBaiduChannelInput) *PinpointDeleteBaiduChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBaiduChannel, input)
    return &PinpointDeleteBaiduChannelResult{Result: future}
}

func (a *PinpointStub) DeleteCampaign(ctx workflow.Context, input *pinpoint.DeleteCampaignInput) (*pinpoint.DeleteCampaignOutput, error) {
    var output pinpoint.DeleteCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteCampaignAsync(ctx workflow.Context, input *pinpoint.DeleteCampaignInput) *PinpointDeleteCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCampaign, input)
    return &PinpointDeleteCampaignResult{Result: future}
}

func (a *PinpointStub) DeleteEmailChannel(ctx workflow.Context, input *pinpoint.DeleteEmailChannelInput) (*pinpoint.DeleteEmailChannelOutput, error) {
    var output pinpoint.DeleteEmailChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteEmailChannelAsync(ctx workflow.Context, input *pinpoint.DeleteEmailChannelInput) *PinpointDeleteEmailChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailChannel, input)
    return &PinpointDeleteEmailChannelResult{Result: future}
}

func (a *PinpointStub) DeleteEmailTemplate(ctx workflow.Context, input *pinpoint.DeleteEmailTemplateInput) (*pinpoint.DeleteEmailTemplateOutput, error) {
    var output pinpoint.DeleteEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteEmailTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteEmailTemplateInput) *PinpointDeleteEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEmailTemplate, input)
    return &PinpointDeleteEmailTemplateResult{Result: future}
}

func (a *PinpointStub) DeleteEndpoint(ctx workflow.Context, input *pinpoint.DeleteEndpointInput) (*pinpoint.DeleteEndpointOutput, error) {
    var output pinpoint.DeleteEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteEndpointAsync(ctx workflow.Context, input *pinpoint.DeleteEndpointInput) *PinpointDeleteEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input)
    return &PinpointDeleteEndpointResult{Result: future}
}

func (a *PinpointStub) DeleteEventStream(ctx workflow.Context, input *pinpoint.DeleteEventStreamInput) (*pinpoint.DeleteEventStreamOutput, error) {
    var output pinpoint.DeleteEventStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventStream, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteEventStreamAsync(ctx workflow.Context, input *pinpoint.DeleteEventStreamInput) *PinpointDeleteEventStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEventStream, input)
    return &PinpointDeleteEventStreamResult{Result: future}
}

func (a *PinpointStub) DeleteGcmChannel(ctx workflow.Context, input *pinpoint.DeleteGcmChannelInput) (*pinpoint.DeleteGcmChannelOutput, error) {
    var output pinpoint.DeleteGcmChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGcmChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteGcmChannelAsync(ctx workflow.Context, input *pinpoint.DeleteGcmChannelInput) *PinpointDeleteGcmChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGcmChannel, input)
    return &PinpointDeleteGcmChannelResult{Result: future}
}

func (a *PinpointStub) DeleteJourney(ctx workflow.Context, input *pinpoint.DeleteJourneyInput) (*pinpoint.DeleteJourneyOutput, error) {
    var output pinpoint.DeleteJourneyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJourney, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteJourneyAsync(ctx workflow.Context, input *pinpoint.DeleteJourneyInput) *PinpointDeleteJourneyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJourney, input)
    return &PinpointDeleteJourneyResult{Result: future}
}

func (a *PinpointStub) DeletePushTemplate(ctx workflow.Context, input *pinpoint.DeletePushTemplateInput) (*pinpoint.DeletePushTemplateOutput, error) {
    var output pinpoint.DeletePushTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePushTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeletePushTemplateAsync(ctx workflow.Context, input *pinpoint.DeletePushTemplateInput) *PinpointDeletePushTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePushTemplate, input)
    return &PinpointDeletePushTemplateResult{Result: future}
}

func (a *PinpointStub) DeleteRecommenderConfiguration(ctx workflow.Context, input *pinpoint.DeleteRecommenderConfigurationInput) (*pinpoint.DeleteRecommenderConfigurationOutput, error) {
    var output pinpoint.DeleteRecommenderConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRecommenderConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.DeleteRecommenderConfigurationInput) *PinpointDeleteRecommenderConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRecommenderConfiguration, input)
    return &PinpointDeleteRecommenderConfigurationResult{Result: future}
}

func (a *PinpointStub) DeleteSegment(ctx workflow.Context, input *pinpoint.DeleteSegmentInput) (*pinpoint.DeleteSegmentOutput, error) {
    var output pinpoint.DeleteSegmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSegment, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteSegmentAsync(ctx workflow.Context, input *pinpoint.DeleteSegmentInput) *PinpointDeleteSegmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSegment, input)
    return &PinpointDeleteSegmentResult{Result: future}
}

func (a *PinpointStub) DeleteSmsChannel(ctx workflow.Context, input *pinpoint.DeleteSmsChannelInput) (*pinpoint.DeleteSmsChannelOutput, error) {
    var output pinpoint.DeleteSmsChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSmsChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteSmsChannelAsync(ctx workflow.Context, input *pinpoint.DeleteSmsChannelInput) *PinpointDeleteSmsChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSmsChannel, input)
    return &PinpointDeleteSmsChannelResult{Result: future}
}

func (a *PinpointStub) DeleteSmsTemplate(ctx workflow.Context, input *pinpoint.DeleteSmsTemplateInput) (*pinpoint.DeleteSmsTemplateOutput, error) {
    var output pinpoint.DeleteSmsTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSmsTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteSmsTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteSmsTemplateInput) *PinpointDeleteSmsTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSmsTemplate, input)
    return &PinpointDeleteSmsTemplateResult{Result: future}
}

func (a *PinpointStub) DeleteUserEndpoints(ctx workflow.Context, input *pinpoint.DeleteUserEndpointsInput) (*pinpoint.DeleteUserEndpointsOutput, error) {
    var output pinpoint.DeleteUserEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteUserEndpointsAsync(ctx workflow.Context, input *pinpoint.DeleteUserEndpointsInput) *PinpointDeleteUserEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUserEndpoints, input)
    return &PinpointDeleteUserEndpointsResult{Result: future}
}

func (a *PinpointStub) DeleteVoiceChannel(ctx workflow.Context, input *pinpoint.DeleteVoiceChannelInput) (*pinpoint.DeleteVoiceChannelOutput, error) {
    var output pinpoint.DeleteVoiceChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteVoiceChannelAsync(ctx workflow.Context, input *pinpoint.DeleteVoiceChannelInput) *PinpointDeleteVoiceChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceChannel, input)
    return &PinpointDeleteVoiceChannelResult{Result: future}
}

func (a *PinpointStub) DeleteVoiceTemplate(ctx workflow.Context, input *pinpoint.DeleteVoiceTemplateInput) (*pinpoint.DeleteVoiceTemplateOutput, error) {
    var output pinpoint.DeleteVoiceTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) DeleteVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteVoiceTemplateInput) *PinpointDeleteVoiceTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceTemplate, input)
    return &PinpointDeleteVoiceTemplateResult{Result: future}
}

func (a *PinpointStub) GetAdmChannel(ctx workflow.Context, input *pinpoint.GetAdmChannelInput) (*pinpoint.GetAdmChannelOutput, error) {
    var output pinpoint.GetAdmChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAdmChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetAdmChannelAsync(ctx workflow.Context, input *pinpoint.GetAdmChannelInput) *PinpointGetAdmChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAdmChannel, input)
    return &PinpointGetAdmChannelResult{Result: future}
}

func (a *PinpointStub) GetApnsChannel(ctx workflow.Context, input *pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error) {
    var output pinpoint.GetApnsChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApnsChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetApnsChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsChannelInput) *PinpointGetApnsChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApnsChannel, input)
    return &PinpointGetApnsChannelResult{Result: future}
}

func (a *PinpointStub) GetApnsSandboxChannel(ctx workflow.Context, input *pinpoint.GetApnsSandboxChannelInput) (*pinpoint.GetApnsSandboxChannelOutput, error) {
    var output pinpoint.GetApnsSandboxChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApnsSandboxChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsSandboxChannelInput) *PinpointGetApnsSandboxChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApnsSandboxChannel, input)
    return &PinpointGetApnsSandboxChannelResult{Result: future}
}

func (a *PinpointStub) GetApnsVoipChannel(ctx workflow.Context, input *pinpoint.GetApnsVoipChannelInput) (*pinpoint.GetApnsVoipChannelOutput, error) {
    var output pinpoint.GetApnsVoipChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApnsVoipChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsVoipChannelInput) *PinpointGetApnsVoipChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApnsVoipChannel, input)
    return &PinpointGetApnsVoipChannelResult{Result: future}
}

func (a *PinpointStub) GetApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.GetApnsVoipSandboxChannelInput) (*pinpoint.GetApnsVoipSandboxChannelOutput, error) {
    var output pinpoint.GetApnsVoipSandboxChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApnsVoipSandboxChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsVoipSandboxChannelInput) *PinpointGetApnsVoipSandboxChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApnsVoipSandboxChannel, input)
    return &PinpointGetApnsVoipSandboxChannelResult{Result: future}
}

func (a *PinpointStub) GetApp(ctx workflow.Context, input *pinpoint.GetAppInput) (*pinpoint.GetAppOutput, error) {
    var output pinpoint.GetAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApp, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetAppAsync(ctx workflow.Context, input *pinpoint.GetAppInput) *PinpointGetAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApp, input)
    return &PinpointGetAppResult{Result: future}
}

func (a *PinpointStub) GetApplicationDateRangeKpi(ctx workflow.Context, input *pinpoint.GetApplicationDateRangeKpiInput) (*pinpoint.GetApplicationDateRangeKpiOutput, error) {
    var output pinpoint.GetApplicationDateRangeKpiOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApplicationDateRangeKpi, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetApplicationDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetApplicationDateRangeKpiInput) *PinpointGetApplicationDateRangeKpiResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApplicationDateRangeKpi, input)
    return &PinpointGetApplicationDateRangeKpiResult{Result: future}
}

func (a *PinpointStub) GetApplicationSettings(ctx workflow.Context, input *pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error) {
    var output pinpoint.GetApplicationSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApplicationSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetApplicationSettingsAsync(ctx workflow.Context, input *pinpoint.GetApplicationSettingsInput) *PinpointGetApplicationSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApplicationSettings, input)
    return &PinpointGetApplicationSettingsResult{Result: future}
}

func (a *PinpointStub) GetApps(ctx workflow.Context, input *pinpoint.GetAppsInput) (*pinpoint.GetAppsOutput, error) {
    var output pinpoint.GetAppsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApps, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetAppsAsync(ctx workflow.Context, input *pinpoint.GetAppsInput) *PinpointGetAppsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApps, input)
    return &PinpointGetAppsResult{Result: future}
}

func (a *PinpointStub) GetBaiduChannel(ctx workflow.Context, input *pinpoint.GetBaiduChannelInput) (*pinpoint.GetBaiduChannelOutput, error) {
    var output pinpoint.GetBaiduChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBaiduChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetBaiduChannelAsync(ctx workflow.Context, input *pinpoint.GetBaiduChannelInput) *PinpointGetBaiduChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBaiduChannel, input)
    return &PinpointGetBaiduChannelResult{Result: future}
}

func (a *PinpointStub) GetCampaign(ctx workflow.Context, input *pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error) {
    var output pinpoint.GetCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetCampaignAsync(ctx workflow.Context, input *pinpoint.GetCampaignInput) *PinpointGetCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCampaign, input)
    return &PinpointGetCampaignResult{Result: future}
}

func (a *PinpointStub) GetCampaignActivities(ctx workflow.Context, input *pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error) {
    var output pinpoint.GetCampaignActivitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCampaignActivities, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetCampaignActivitiesAsync(ctx workflow.Context, input *pinpoint.GetCampaignActivitiesInput) *PinpointGetCampaignActivitiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCampaignActivities, input)
    return &PinpointGetCampaignActivitiesResult{Result: future}
}

func (a *PinpointStub) GetCampaignDateRangeKpi(ctx workflow.Context, input *pinpoint.GetCampaignDateRangeKpiInput) (*pinpoint.GetCampaignDateRangeKpiOutput, error) {
    var output pinpoint.GetCampaignDateRangeKpiOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCampaignDateRangeKpi, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetCampaignDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetCampaignDateRangeKpiInput) *PinpointGetCampaignDateRangeKpiResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCampaignDateRangeKpi, input)
    return &PinpointGetCampaignDateRangeKpiResult{Result: future}
}

func (a *PinpointStub) GetCampaignVersion(ctx workflow.Context, input *pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error) {
    var output pinpoint.GetCampaignVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCampaignVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetCampaignVersionAsync(ctx workflow.Context, input *pinpoint.GetCampaignVersionInput) *PinpointGetCampaignVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCampaignVersion, input)
    return &PinpointGetCampaignVersionResult{Result: future}
}

func (a *PinpointStub) GetCampaignVersions(ctx workflow.Context, input *pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error) {
    var output pinpoint.GetCampaignVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCampaignVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetCampaignVersionsAsync(ctx workflow.Context, input *pinpoint.GetCampaignVersionsInput) *PinpointGetCampaignVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCampaignVersions, input)
    return &PinpointGetCampaignVersionsResult{Result: future}
}

func (a *PinpointStub) GetCampaigns(ctx workflow.Context, input *pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error) {
    var output pinpoint.GetCampaignsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCampaigns, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetCampaignsAsync(ctx workflow.Context, input *pinpoint.GetCampaignsInput) *PinpointGetCampaignsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCampaigns, input)
    return &PinpointGetCampaignsResult{Result: future}
}

func (a *PinpointStub) GetChannels(ctx workflow.Context, input *pinpoint.GetChannelsInput) (*pinpoint.GetChannelsOutput, error) {
    var output pinpoint.GetChannelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetChannels, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetChannelsAsync(ctx workflow.Context, input *pinpoint.GetChannelsInput) *PinpointGetChannelsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetChannels, input)
    return &PinpointGetChannelsResult{Result: future}
}

func (a *PinpointStub) GetEmailChannel(ctx workflow.Context, input *pinpoint.GetEmailChannelInput) (*pinpoint.GetEmailChannelOutput, error) {
    var output pinpoint.GetEmailChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEmailChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetEmailChannelAsync(ctx workflow.Context, input *pinpoint.GetEmailChannelInput) *PinpointGetEmailChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEmailChannel, input)
    return &PinpointGetEmailChannelResult{Result: future}
}

func (a *PinpointStub) GetEmailTemplate(ctx workflow.Context, input *pinpoint.GetEmailTemplateInput) (*pinpoint.GetEmailTemplateOutput, error) {
    var output pinpoint.GetEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetEmailTemplateAsync(ctx workflow.Context, input *pinpoint.GetEmailTemplateInput) *PinpointGetEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEmailTemplate, input)
    return &PinpointGetEmailTemplateResult{Result: future}
}

func (a *PinpointStub) GetEndpoint(ctx workflow.Context, input *pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error) {
    var output pinpoint.GetEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetEndpointAsync(ctx workflow.Context, input *pinpoint.GetEndpointInput) *PinpointGetEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEndpoint, input)
    return &PinpointGetEndpointResult{Result: future}
}

func (a *PinpointStub) GetEventStream(ctx workflow.Context, input *pinpoint.GetEventStreamInput) (*pinpoint.GetEventStreamOutput, error) {
    var output pinpoint.GetEventStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEventStream, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetEventStreamAsync(ctx workflow.Context, input *pinpoint.GetEventStreamInput) *PinpointGetEventStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEventStream, input)
    return &PinpointGetEventStreamResult{Result: future}
}

func (a *PinpointStub) GetExportJob(ctx workflow.Context, input *pinpoint.GetExportJobInput) (*pinpoint.GetExportJobOutput, error) {
    var output pinpoint.GetExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetExportJobAsync(ctx workflow.Context, input *pinpoint.GetExportJobInput) *PinpointGetExportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetExportJob, input)
    return &PinpointGetExportJobResult{Result: future}
}

func (a *PinpointStub) GetExportJobs(ctx workflow.Context, input *pinpoint.GetExportJobsInput) (*pinpoint.GetExportJobsOutput, error) {
    var output pinpoint.GetExportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetExportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetExportJobsAsync(ctx workflow.Context, input *pinpoint.GetExportJobsInput) *PinpointGetExportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetExportJobs, input)
    return &PinpointGetExportJobsResult{Result: future}
}

func (a *PinpointStub) GetGcmChannel(ctx workflow.Context, input *pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error) {
    var output pinpoint.GetGcmChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGcmChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetGcmChannelAsync(ctx workflow.Context, input *pinpoint.GetGcmChannelInput) *PinpointGetGcmChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGcmChannel, input)
    return &PinpointGetGcmChannelResult{Result: future}
}

func (a *PinpointStub) GetImportJob(ctx workflow.Context, input *pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error) {
    var output pinpoint.GetImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetImportJobAsync(ctx workflow.Context, input *pinpoint.GetImportJobInput) *PinpointGetImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetImportJob, input)
    return &PinpointGetImportJobResult{Result: future}
}

func (a *PinpointStub) GetImportJobs(ctx workflow.Context, input *pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error) {
    var output pinpoint.GetImportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetImportJobsAsync(ctx workflow.Context, input *pinpoint.GetImportJobsInput) *PinpointGetImportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetImportJobs, input)
    return &PinpointGetImportJobsResult{Result: future}
}

func (a *PinpointStub) GetJourney(ctx workflow.Context, input *pinpoint.GetJourneyInput) (*pinpoint.GetJourneyOutput, error) {
    var output pinpoint.GetJourneyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJourney, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetJourneyAsync(ctx workflow.Context, input *pinpoint.GetJourneyInput) *PinpointGetJourneyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJourney, input)
    return &PinpointGetJourneyResult{Result: future}
}

func (a *PinpointStub) GetJourneyDateRangeKpi(ctx workflow.Context, input *pinpoint.GetJourneyDateRangeKpiInput) (*pinpoint.GetJourneyDateRangeKpiOutput, error) {
    var output pinpoint.GetJourneyDateRangeKpiOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJourneyDateRangeKpi, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetJourneyDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetJourneyDateRangeKpiInput) *PinpointGetJourneyDateRangeKpiResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJourneyDateRangeKpi, input)
    return &PinpointGetJourneyDateRangeKpiResult{Result: future}
}

func (a *PinpointStub) GetJourneyExecutionActivityMetrics(ctx workflow.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error) {
    var output pinpoint.GetJourneyExecutionActivityMetricsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJourneyExecutionActivityMetrics, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetJourneyExecutionActivityMetricsAsync(ctx workflow.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput) *PinpointGetJourneyExecutionActivityMetricsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJourneyExecutionActivityMetrics, input)
    return &PinpointGetJourneyExecutionActivityMetricsResult{Result: future}
}

func (a *PinpointStub) GetJourneyExecutionMetrics(ctx workflow.Context, input *pinpoint.GetJourneyExecutionMetricsInput) (*pinpoint.GetJourneyExecutionMetricsOutput, error) {
    var output pinpoint.GetJourneyExecutionMetricsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJourneyExecutionMetrics, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetJourneyExecutionMetricsAsync(ctx workflow.Context, input *pinpoint.GetJourneyExecutionMetricsInput) *PinpointGetJourneyExecutionMetricsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJourneyExecutionMetrics, input)
    return &PinpointGetJourneyExecutionMetricsResult{Result: future}
}

func (a *PinpointStub) GetPushTemplate(ctx workflow.Context, input *pinpoint.GetPushTemplateInput) (*pinpoint.GetPushTemplateOutput, error) {
    var output pinpoint.GetPushTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPushTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetPushTemplateAsync(ctx workflow.Context, input *pinpoint.GetPushTemplateInput) *PinpointGetPushTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPushTemplate, input)
    return &PinpointGetPushTemplateResult{Result: future}
}

func (a *PinpointStub) GetRecommenderConfiguration(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationInput) (*pinpoint.GetRecommenderConfigurationOutput, error) {
    var output pinpoint.GetRecommenderConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRecommenderConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationInput) *PinpointGetRecommenderConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRecommenderConfiguration, input)
    return &PinpointGetRecommenderConfigurationResult{Result: future}
}

func (a *PinpointStub) GetRecommenderConfigurations(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationsInput) (*pinpoint.GetRecommenderConfigurationsOutput, error) {
    var output pinpoint.GetRecommenderConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRecommenderConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetRecommenderConfigurationsAsync(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationsInput) *PinpointGetRecommenderConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRecommenderConfigurations, input)
    return &PinpointGetRecommenderConfigurationsResult{Result: future}
}

func (a *PinpointStub) GetSegment(ctx workflow.Context, input *pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error) {
    var output pinpoint.GetSegmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSegment, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSegmentAsync(ctx workflow.Context, input *pinpoint.GetSegmentInput) *PinpointGetSegmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSegment, input)
    return &PinpointGetSegmentResult{Result: future}
}

func (a *PinpointStub) GetSegmentExportJobs(ctx workflow.Context, input *pinpoint.GetSegmentExportJobsInput) (*pinpoint.GetSegmentExportJobsOutput, error) {
    var output pinpoint.GetSegmentExportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSegmentExportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSegmentExportJobsAsync(ctx workflow.Context, input *pinpoint.GetSegmentExportJobsInput) *PinpointGetSegmentExportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSegmentExportJobs, input)
    return &PinpointGetSegmentExportJobsResult{Result: future}
}

func (a *PinpointStub) GetSegmentImportJobs(ctx workflow.Context, input *pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error) {
    var output pinpoint.GetSegmentImportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSegmentImportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSegmentImportJobsAsync(ctx workflow.Context, input *pinpoint.GetSegmentImportJobsInput) *PinpointGetSegmentImportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSegmentImportJobs, input)
    return &PinpointGetSegmentImportJobsResult{Result: future}
}

func (a *PinpointStub) GetSegmentVersion(ctx workflow.Context, input *pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error) {
    var output pinpoint.GetSegmentVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSegmentVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSegmentVersionAsync(ctx workflow.Context, input *pinpoint.GetSegmentVersionInput) *PinpointGetSegmentVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSegmentVersion, input)
    return &PinpointGetSegmentVersionResult{Result: future}
}

func (a *PinpointStub) GetSegmentVersions(ctx workflow.Context, input *pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error) {
    var output pinpoint.GetSegmentVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSegmentVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSegmentVersionsAsync(ctx workflow.Context, input *pinpoint.GetSegmentVersionsInput) *PinpointGetSegmentVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSegmentVersions, input)
    return &PinpointGetSegmentVersionsResult{Result: future}
}

func (a *PinpointStub) GetSegments(ctx workflow.Context, input *pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error) {
    var output pinpoint.GetSegmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSegments, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSegmentsAsync(ctx workflow.Context, input *pinpoint.GetSegmentsInput) *PinpointGetSegmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSegments, input)
    return &PinpointGetSegmentsResult{Result: future}
}

func (a *PinpointStub) GetSmsChannel(ctx workflow.Context, input *pinpoint.GetSmsChannelInput) (*pinpoint.GetSmsChannelOutput, error) {
    var output pinpoint.GetSmsChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSmsChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSmsChannelAsync(ctx workflow.Context, input *pinpoint.GetSmsChannelInput) *PinpointGetSmsChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSmsChannel, input)
    return &PinpointGetSmsChannelResult{Result: future}
}

func (a *PinpointStub) GetSmsTemplate(ctx workflow.Context, input *pinpoint.GetSmsTemplateInput) (*pinpoint.GetSmsTemplateOutput, error) {
    var output pinpoint.GetSmsTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSmsTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetSmsTemplateAsync(ctx workflow.Context, input *pinpoint.GetSmsTemplateInput) *PinpointGetSmsTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSmsTemplate, input)
    return &PinpointGetSmsTemplateResult{Result: future}
}

func (a *PinpointStub) GetUserEndpoints(ctx workflow.Context, input *pinpoint.GetUserEndpointsInput) (*pinpoint.GetUserEndpointsOutput, error) {
    var output pinpoint.GetUserEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUserEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetUserEndpointsAsync(ctx workflow.Context, input *pinpoint.GetUserEndpointsInput) *PinpointGetUserEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUserEndpoints, input)
    return &PinpointGetUserEndpointsResult{Result: future}
}

func (a *PinpointStub) GetVoiceChannel(ctx workflow.Context, input *pinpoint.GetVoiceChannelInput) (*pinpoint.GetVoiceChannelOutput, error) {
    var output pinpoint.GetVoiceChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetVoiceChannelAsync(ctx workflow.Context, input *pinpoint.GetVoiceChannelInput) *PinpointGetVoiceChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVoiceChannel, input)
    return &PinpointGetVoiceChannelResult{Result: future}
}

func (a *PinpointStub) GetVoiceTemplate(ctx workflow.Context, input *pinpoint.GetVoiceTemplateInput) (*pinpoint.GetVoiceTemplateOutput, error) {
    var output pinpoint.GetVoiceTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) GetVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.GetVoiceTemplateInput) *PinpointGetVoiceTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVoiceTemplate, input)
    return &PinpointGetVoiceTemplateResult{Result: future}
}

func (a *PinpointStub) ListJourneys(ctx workflow.Context, input *pinpoint.ListJourneysInput) (*pinpoint.ListJourneysOutput, error) {
    var output pinpoint.ListJourneysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJourneys, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) ListJourneysAsync(ctx workflow.Context, input *pinpoint.ListJourneysInput) *PinpointListJourneysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJourneys, input)
    return &PinpointListJourneysResult{Result: future}
}

func (a *PinpointStub) ListTagsForResource(ctx workflow.Context, input *pinpoint.ListTagsForResourceInput) (*pinpoint.ListTagsForResourceOutput, error) {
    var output pinpoint.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) ListTagsForResourceAsync(ctx workflow.Context, input *pinpoint.ListTagsForResourceInput) *PinpointListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &PinpointListTagsForResourceResult{Result: future}
}

func (a *PinpointStub) ListTemplateVersions(ctx workflow.Context, input *pinpoint.ListTemplateVersionsInput) (*pinpoint.ListTemplateVersionsOutput, error) {
    var output pinpoint.ListTemplateVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTemplateVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) ListTemplateVersionsAsync(ctx workflow.Context, input *pinpoint.ListTemplateVersionsInput) *PinpointListTemplateVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTemplateVersions, input)
    return &PinpointListTemplateVersionsResult{Result: future}
}

func (a *PinpointStub) ListTemplates(ctx workflow.Context, input *pinpoint.ListTemplatesInput) (*pinpoint.ListTemplatesOutput, error) {
    var output pinpoint.ListTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) ListTemplatesAsync(ctx workflow.Context, input *pinpoint.ListTemplatesInput) *PinpointListTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTemplates, input)
    return &PinpointListTemplatesResult{Result: future}
}

func (a *PinpointStub) PhoneNumberValidate(ctx workflow.Context, input *pinpoint.PhoneNumberValidateInput) (*pinpoint.PhoneNumberValidateOutput, error) {
    var output pinpoint.PhoneNumberValidateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PhoneNumberValidate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) PhoneNumberValidateAsync(ctx workflow.Context, input *pinpoint.PhoneNumberValidateInput) *PinpointPhoneNumberValidateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PhoneNumberValidate, input)
    return &PinpointPhoneNumberValidateResult{Result: future}
}

func (a *PinpointStub) PutEventStream(ctx workflow.Context, input *pinpoint.PutEventStreamInput) (*pinpoint.PutEventStreamOutput, error) {
    var output pinpoint.PutEventStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEventStream, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) PutEventStreamAsync(ctx workflow.Context, input *pinpoint.PutEventStreamInput) *PinpointPutEventStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEventStream, input)
    return &PinpointPutEventStreamResult{Result: future}
}

func (a *PinpointStub) PutEvents(ctx workflow.Context, input *pinpoint.PutEventsInput) (*pinpoint.PutEventsOutput, error) {
    var output pinpoint.PutEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) PutEventsAsync(ctx workflow.Context, input *pinpoint.PutEventsInput) *PinpointPutEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEvents, input)
    return &PinpointPutEventsResult{Result: future}
}

func (a *PinpointStub) RemoveAttributes(ctx workflow.Context, input *pinpoint.RemoveAttributesInput) (*pinpoint.RemoveAttributesOutput, error) {
    var output pinpoint.RemoveAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) RemoveAttributesAsync(ctx workflow.Context, input *pinpoint.RemoveAttributesInput) *PinpointRemoveAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveAttributes, input)
    return &PinpointRemoveAttributesResult{Result: future}
}

func (a *PinpointStub) SendMessages(ctx workflow.Context, input *pinpoint.SendMessagesInput) (*pinpoint.SendMessagesOutput, error) {
    var output pinpoint.SendMessagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendMessages, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) SendMessagesAsync(ctx workflow.Context, input *pinpoint.SendMessagesInput) *PinpointSendMessagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendMessages, input)
    return &PinpointSendMessagesResult{Result: future}
}

func (a *PinpointStub) SendUsersMessages(ctx workflow.Context, input *pinpoint.SendUsersMessagesInput) (*pinpoint.SendUsersMessagesOutput, error) {
    var output pinpoint.SendUsersMessagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendUsersMessages, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) SendUsersMessagesAsync(ctx workflow.Context, input *pinpoint.SendUsersMessagesInput) *PinpointSendUsersMessagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendUsersMessages, input)
    return &PinpointSendUsersMessagesResult{Result: future}
}

func (a *PinpointStub) TagResource(ctx workflow.Context, input *pinpoint.TagResourceInput) (*pinpoint.TagResourceOutput, error) {
    var output pinpoint.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) TagResourceAsync(ctx workflow.Context, input *pinpoint.TagResourceInput) *PinpointTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &PinpointTagResourceResult{Result: future}
}

func (a *PinpointStub) UntagResource(ctx workflow.Context, input *pinpoint.UntagResourceInput) (*pinpoint.UntagResourceOutput, error) {
    var output pinpoint.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UntagResourceAsync(ctx workflow.Context, input *pinpoint.UntagResourceInput) *PinpointUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &PinpointUntagResourceResult{Result: future}
}

func (a *PinpointStub) UpdateAdmChannel(ctx workflow.Context, input *pinpoint.UpdateAdmChannelInput) (*pinpoint.UpdateAdmChannelOutput, error) {
    var output pinpoint.UpdateAdmChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAdmChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateAdmChannelAsync(ctx workflow.Context, input *pinpoint.UpdateAdmChannelInput) *PinpointUpdateAdmChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAdmChannel, input)
    return &PinpointUpdateAdmChannelResult{Result: future}
}

func (a *PinpointStub) UpdateApnsChannel(ctx workflow.Context, input *pinpoint.UpdateApnsChannelInput) (*pinpoint.UpdateApnsChannelOutput, error) {
    var output pinpoint.UpdateApnsChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateApnsChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsChannelInput) *PinpointUpdateApnsChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsChannel, input)
    return &PinpointUpdateApnsChannelResult{Result: future}
}

func (a *PinpointStub) UpdateApnsSandboxChannel(ctx workflow.Context, input *pinpoint.UpdateApnsSandboxChannelInput) (*pinpoint.UpdateApnsSandboxChannelOutput, error) {
    var output pinpoint.UpdateApnsSandboxChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsSandboxChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsSandboxChannelInput) *PinpointUpdateApnsSandboxChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsSandboxChannel, input)
    return &PinpointUpdateApnsSandboxChannelResult{Result: future}
}

func (a *PinpointStub) UpdateApnsVoipChannel(ctx workflow.Context, input *pinpoint.UpdateApnsVoipChannelInput) (*pinpoint.UpdateApnsVoipChannelOutput, error) {
    var output pinpoint.UpdateApnsVoipChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsVoipChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsVoipChannelInput) *PinpointUpdateApnsVoipChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsVoipChannel, input)
    return &PinpointUpdateApnsVoipChannelResult{Result: future}
}

func (a *PinpointStub) UpdateApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.UpdateApnsVoipSandboxChannelInput) (*pinpoint.UpdateApnsVoipSandboxChannelOutput, error) {
    var output pinpoint.UpdateApnsVoipSandboxChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsVoipSandboxChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsVoipSandboxChannelInput) *PinpointUpdateApnsVoipSandboxChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApnsVoipSandboxChannel, input)
    return &PinpointUpdateApnsVoipSandboxChannelResult{Result: future}
}

func (a *PinpointStub) UpdateApplicationSettings(ctx workflow.Context, input *pinpoint.UpdateApplicationSettingsInput) (*pinpoint.UpdateApplicationSettingsOutput, error) {
    var output pinpoint.UpdateApplicationSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplicationSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateApplicationSettingsAsync(ctx workflow.Context, input *pinpoint.UpdateApplicationSettingsInput) *PinpointUpdateApplicationSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApplicationSettings, input)
    return &PinpointUpdateApplicationSettingsResult{Result: future}
}

func (a *PinpointStub) UpdateBaiduChannel(ctx workflow.Context, input *pinpoint.UpdateBaiduChannelInput) (*pinpoint.UpdateBaiduChannelOutput, error) {
    var output pinpoint.UpdateBaiduChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBaiduChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateBaiduChannelAsync(ctx workflow.Context, input *pinpoint.UpdateBaiduChannelInput) *PinpointUpdateBaiduChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBaiduChannel, input)
    return &PinpointUpdateBaiduChannelResult{Result: future}
}

func (a *PinpointStub) UpdateCampaign(ctx workflow.Context, input *pinpoint.UpdateCampaignInput) (*pinpoint.UpdateCampaignOutput, error) {
    var output pinpoint.UpdateCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateCampaignAsync(ctx workflow.Context, input *pinpoint.UpdateCampaignInput) *PinpointUpdateCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCampaign, input)
    return &PinpointUpdateCampaignResult{Result: future}
}

func (a *PinpointStub) UpdateEmailChannel(ctx workflow.Context, input *pinpoint.UpdateEmailChannelInput) (*pinpoint.UpdateEmailChannelOutput, error) {
    var output pinpoint.UpdateEmailChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateEmailChannelAsync(ctx workflow.Context, input *pinpoint.UpdateEmailChannelInput) *PinpointUpdateEmailChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailChannel, input)
    return &PinpointUpdateEmailChannelResult{Result: future}
}

func (a *PinpointStub) UpdateEmailTemplate(ctx workflow.Context, input *pinpoint.UpdateEmailTemplateInput) (*pinpoint.UpdateEmailTemplateOutput, error) {
    var output pinpoint.UpdateEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateEmailTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateEmailTemplateInput) *PinpointUpdateEmailTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEmailTemplate, input)
    return &PinpointUpdateEmailTemplateResult{Result: future}
}

func (a *PinpointStub) UpdateEndpoint(ctx workflow.Context, input *pinpoint.UpdateEndpointInput) (*pinpoint.UpdateEndpointOutput, error) {
    var output pinpoint.UpdateEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateEndpointAsync(ctx workflow.Context, input *pinpoint.UpdateEndpointInput) *PinpointUpdateEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpoint, input)
    return &PinpointUpdateEndpointResult{Result: future}
}

func (a *PinpointStub) UpdateEndpointsBatch(ctx workflow.Context, input *pinpoint.UpdateEndpointsBatchInput) (*pinpoint.UpdateEndpointsBatchOutput, error) {
    var output pinpoint.UpdateEndpointsBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpointsBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateEndpointsBatchAsync(ctx workflow.Context, input *pinpoint.UpdateEndpointsBatchInput) *PinpointUpdateEndpointsBatchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpointsBatch, input)
    return &PinpointUpdateEndpointsBatchResult{Result: future}
}

func (a *PinpointStub) UpdateGcmChannel(ctx workflow.Context, input *pinpoint.UpdateGcmChannelInput) (*pinpoint.UpdateGcmChannelOutput, error) {
    var output pinpoint.UpdateGcmChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGcmChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateGcmChannelAsync(ctx workflow.Context, input *pinpoint.UpdateGcmChannelInput) *PinpointUpdateGcmChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGcmChannel, input)
    return &PinpointUpdateGcmChannelResult{Result: future}
}

func (a *PinpointStub) UpdateJourney(ctx workflow.Context, input *pinpoint.UpdateJourneyInput) (*pinpoint.UpdateJourneyOutput, error) {
    var output pinpoint.UpdateJourneyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJourney, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateJourneyAsync(ctx workflow.Context, input *pinpoint.UpdateJourneyInput) *PinpointUpdateJourneyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJourney, input)
    return &PinpointUpdateJourneyResult{Result: future}
}

func (a *PinpointStub) UpdateJourneyState(ctx workflow.Context, input *pinpoint.UpdateJourneyStateInput) (*pinpoint.UpdateJourneyStateOutput, error) {
    var output pinpoint.UpdateJourneyStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJourneyState, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateJourneyStateAsync(ctx workflow.Context, input *pinpoint.UpdateJourneyStateInput) *PinpointUpdateJourneyStateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJourneyState, input)
    return &PinpointUpdateJourneyStateResult{Result: future}
}

func (a *PinpointStub) UpdatePushTemplate(ctx workflow.Context, input *pinpoint.UpdatePushTemplateInput) (*pinpoint.UpdatePushTemplateOutput, error) {
    var output pinpoint.UpdatePushTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePushTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdatePushTemplateAsync(ctx workflow.Context, input *pinpoint.UpdatePushTemplateInput) *PinpointUpdatePushTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePushTemplate, input)
    return &PinpointUpdatePushTemplateResult{Result: future}
}

func (a *PinpointStub) UpdateRecommenderConfiguration(ctx workflow.Context, input *pinpoint.UpdateRecommenderConfigurationInput) (*pinpoint.UpdateRecommenderConfigurationOutput, error) {
    var output pinpoint.UpdateRecommenderConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRecommenderConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.UpdateRecommenderConfigurationInput) *PinpointUpdateRecommenderConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRecommenderConfiguration, input)
    return &PinpointUpdateRecommenderConfigurationResult{Result: future}
}

func (a *PinpointStub) UpdateSegment(ctx workflow.Context, input *pinpoint.UpdateSegmentInput) (*pinpoint.UpdateSegmentOutput, error) {
    var output pinpoint.UpdateSegmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSegment, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateSegmentAsync(ctx workflow.Context, input *pinpoint.UpdateSegmentInput) *PinpointUpdateSegmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSegment, input)
    return &PinpointUpdateSegmentResult{Result: future}
}

func (a *PinpointStub) UpdateSmsChannel(ctx workflow.Context, input *pinpoint.UpdateSmsChannelInput) (*pinpoint.UpdateSmsChannelOutput, error) {
    var output pinpoint.UpdateSmsChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSmsChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateSmsChannelAsync(ctx workflow.Context, input *pinpoint.UpdateSmsChannelInput) *PinpointUpdateSmsChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSmsChannel, input)
    return &PinpointUpdateSmsChannelResult{Result: future}
}

func (a *PinpointStub) UpdateSmsTemplate(ctx workflow.Context, input *pinpoint.UpdateSmsTemplateInput) (*pinpoint.UpdateSmsTemplateOutput, error) {
    var output pinpoint.UpdateSmsTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSmsTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateSmsTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateSmsTemplateInput) *PinpointUpdateSmsTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSmsTemplate, input)
    return &PinpointUpdateSmsTemplateResult{Result: future}
}

func (a *PinpointStub) UpdateTemplateActiveVersion(ctx workflow.Context, input *pinpoint.UpdateTemplateActiveVersionInput) (*pinpoint.UpdateTemplateActiveVersionOutput, error) {
    var output pinpoint.UpdateTemplateActiveVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTemplateActiveVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateTemplateActiveVersionAsync(ctx workflow.Context, input *pinpoint.UpdateTemplateActiveVersionInput) *PinpointUpdateTemplateActiveVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTemplateActiveVersion, input)
    return &PinpointUpdateTemplateActiveVersionResult{Result: future}
}

func (a *PinpointStub) UpdateVoiceChannel(ctx workflow.Context, input *pinpoint.UpdateVoiceChannelInput) (*pinpoint.UpdateVoiceChannelOutput, error) {
    var output pinpoint.UpdateVoiceChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVoiceChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateVoiceChannelAsync(ctx workflow.Context, input *pinpoint.UpdateVoiceChannelInput) *PinpointUpdateVoiceChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVoiceChannel, input)
    return &PinpointUpdateVoiceChannelResult{Result: future}
}

func (a *PinpointStub) UpdateVoiceTemplate(ctx workflow.Context, input *pinpoint.UpdateVoiceTemplateInput) (*pinpoint.UpdateVoiceTemplateOutput, error) {
    var output pinpoint.UpdateVoiceTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVoiceTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *PinpointStub) UpdateVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateVoiceTemplateInput) *PinpointUpdateVoiceTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVoiceTemplate, input)
    return &PinpointUpdateVoiceTemplateResult{Result: future}
}