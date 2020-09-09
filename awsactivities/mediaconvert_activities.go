package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/aws/aws-sdk-go/service/mediaconvert/mediaconvertiface"
)

type MediaConvertActivities struct {
	client mediaconvertiface.MediaConvertAPI
}

func NewMediaConvertActivities(session *session.Session, config ...*aws.Config) *MediaConvertActivities {
	client := mediaconvert.New(session, config...)
	return &MediaConvertActivities{client: client}
}

func (a *MediaConvertActivities) AssociateCertificate(input *mediaconvert.AssociateCertificateInput) (*mediaconvert.AssociateCertificateOutput, error) {
	return a.client.AssociateCertificate(input)
}

func (a *MediaConvertActivities) CancelJob(input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error) {
	return a.client.CancelJob(input)
}

func (a *MediaConvertActivities) CreateJob(input *mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error) {
	return a.client.CreateJob(input)
}

func (a *MediaConvertActivities) CreateJobTemplate(input *mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error) {
	return a.client.CreateJobTemplate(input)
}

func (a *MediaConvertActivities) CreatePreset(input *mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error) {
	return a.client.CreatePreset(input)
}

func (a *MediaConvertActivities) CreateQueue(input *mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error) {
	return a.client.CreateQueue(input)
}

func (a *MediaConvertActivities) DeleteJobTemplate(input *mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error) {
	return a.client.DeleteJobTemplate(input)
}

func (a *MediaConvertActivities) DeletePreset(input *mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error) {
	return a.client.DeletePreset(input)
}

func (a *MediaConvertActivities) DeleteQueue(input *mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error) {
	return a.client.DeleteQueue(input)
}

func (a *MediaConvertActivities) DescribeEndpoints(input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error) {
	return a.client.DescribeEndpoints(input)
}

func (a *MediaConvertActivities) DisassociateCertificate(input *mediaconvert.DisassociateCertificateInput) (*mediaconvert.DisassociateCertificateOutput, error) {
	return a.client.DisassociateCertificate(input)
}

func (a *MediaConvertActivities) GetJob(input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error) {
	return a.client.GetJob(input)
}

func (a *MediaConvertActivities) GetJobTemplate(input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error) {
	return a.client.GetJobTemplate(input)
}

func (a *MediaConvertActivities) GetPreset(input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error) {
	return a.client.GetPreset(input)
}

func (a *MediaConvertActivities) GetQueue(input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error) {
	return a.client.GetQueue(input)
}

func (a *MediaConvertActivities) ListJobTemplates(input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error) {
	return a.client.ListJobTemplates(input)
}

func (a *MediaConvertActivities) ListJobs(input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error) {
	return a.client.ListJobs(input)
}

func (a *MediaConvertActivities) ListPresets(input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error) {
	return a.client.ListPresets(input)
}

func (a *MediaConvertActivities) ListQueues(input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error) {
	return a.client.ListQueues(input)
}

func (a *MediaConvertActivities) ListTagsForResource(input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *MediaConvertActivities) TagResource(input *mediaconvert.TagResourceInput) (*mediaconvert.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *MediaConvertActivities) UntagResource(input *mediaconvert.UntagResourceInput) (*mediaconvert.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *MediaConvertActivities) UpdateJobTemplate(input *mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error) {
	return a.client.UpdateJobTemplate(input)
}

func (a *MediaConvertActivities) UpdatePreset(input *mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error) {
	return a.client.UpdatePreset(input)
}

func (a *MediaConvertActivities) UpdateQueue(input *mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error) {
	return a.client.UpdateQueue(input)
}
