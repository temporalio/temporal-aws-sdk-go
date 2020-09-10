package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ElasticTranscoderActivities struct {
	client elastictranscoderiface.ElasticTranscoderAPI
}

func NewElasticTranscoderActivities(session *session.Session, config ...*aws.Config) *ElasticTranscoderActivities {
	client := elastictranscoder.New(session, config...)
	return &ElasticTranscoderActivities{client: client}
}

func (a *ElasticTranscoderActivities) CancelJob(ctx context.Context, input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error) {
	return a.client.CancelJobWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) CreateJob(ctx context.Context, input *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error) {
	return a.client.CreateJobWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) CreatePipeline(ctx context.Context, input *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error) {
	return a.client.CreatePipelineWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) CreatePreset(ctx context.Context, input *elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error) {
	return a.client.CreatePresetWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) DeletePipeline(ctx context.Context, input *elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error) {
	return a.client.DeletePipelineWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) DeletePreset(ctx context.Context, input *elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error) {
	return a.client.DeletePresetWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) ListJobsByPipeline(ctx context.Context, input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	return a.client.ListJobsByPipelineWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) ListJobsByStatus(ctx context.Context, input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error) {
	return a.client.ListJobsByStatusWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) ListPipelines(ctx context.Context, input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error) {
	return a.client.ListPipelinesWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) ListPresets(ctx context.Context, input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error) {
	return a.client.ListPresetsWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) ReadJob(ctx context.Context, input *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error) {
	return a.client.ReadJobWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) ReadPipeline(ctx context.Context, input *elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error) {
	return a.client.ReadPipelineWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) ReadPreset(ctx context.Context, input *elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error) {
	return a.client.ReadPresetWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) TestRole(ctx context.Context, input *elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error) {
	return a.client.TestRoleWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) UpdatePipeline(ctx context.Context, input *elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error) {
	return a.client.UpdatePipelineWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) UpdatePipelineNotifications(ctx context.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
	return a.client.UpdatePipelineNotificationsWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) UpdatePipelineStatus(ctx context.Context, input *elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
	return a.client.UpdatePipelineStatusWithContext(ctx, input)
}

func (a *ElasticTranscoderActivities) WaitUntilJobComplete(ctx context.Context, input *elastictranscoder.ReadJobInput) error {
	return a.client.WaitUntilJobCompleteWithContext(ctx, input)

}
