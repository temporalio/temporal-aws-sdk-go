package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface"
)

type ElasticTranscoderActivities struct {
	client elastictranscoderiface.ElasticTranscoderAPI
}

func NewElasticTranscoderActivities(client elastictranscoderiface.ElasticTranscoderAPI) *ElasticTranscoderActivities {
    return &ElasticTranscoderActivities{client: client}
}


func (a *ElasticTranscoderActivities) CancelJob(input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error) {
    return a.client.CancelJob(input)
}



func (a *ElasticTranscoderActivities) CreateJob(input *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error) {
    return a.client.CreateJob(input)
}



func (a *ElasticTranscoderActivities) CreatePipeline(input *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error) {
    return a.client.CreatePipeline(input)
}



func (a *ElasticTranscoderActivities) CreatePreset(input *elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error) {
    return a.client.CreatePreset(input)
}



func (a *ElasticTranscoderActivities) DeletePipeline(input *elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error) {
    return a.client.DeletePipeline(input)
}



func (a *ElasticTranscoderActivities) DeletePreset(input *elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error) {
    return a.client.DeletePreset(input)
}



func (a *ElasticTranscoderActivities) ListJobsByPipeline(input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error) {
    return a.client.ListJobsByPipeline(input)
}



func (a *ElasticTranscoderActivities) ListJobsByStatus(input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error) {
    return a.client.ListJobsByStatus(input)
}



func (a *ElasticTranscoderActivities) ListPipelines(input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error) {
    return a.client.ListPipelines(input)
}



func (a *ElasticTranscoderActivities) ListPresets(input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error) {
    return a.client.ListPresets(input)
}



func (a *ElasticTranscoderActivities) ReadJob(input *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error) {
    return a.client.ReadJob(input)
}



func (a *ElasticTranscoderActivities) ReadPipeline(input *elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error) {
    return a.client.ReadPipeline(input)
}



func (a *ElasticTranscoderActivities) ReadPreset(input *elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error) {
    return a.client.ReadPreset(input)
}



func (a *ElasticTranscoderActivities) TestRole(input *elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error) {
    return a.client.TestRole(input)
}



func (a *ElasticTranscoderActivities) UpdatePipeline(input *elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error) {
    return a.client.UpdatePipeline(input)
}



func (a *ElasticTranscoderActivities) UpdatePipelineNotifications(input *elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
    return a.client.UpdatePipelineNotifications(input)
}



func (a *ElasticTranscoderActivities) UpdatePipelineStatus(input *elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
    return a.client.UpdatePipelineStatus(input)
}



func (a *ElasticTranscoderActivities) WaitUntilJobComplete(input *elastictranscoder.ReadJobInput) error {
    return a.client.WaitUntilJobComplete(input)
}

