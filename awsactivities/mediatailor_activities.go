package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type MediaTailorActivities struct {
	client mediatailoriface.MediaTailorAPI
}

func NewMediaTailorActivities(session *session.Session, config ...*aws.Config) *MediaTailorActivities {
	client := mediatailor.New(session, config...)
	return &MediaTailorActivities{client: client}
}

func (a *MediaTailorActivities) DeletePlaybackConfiguration(ctx context.Context, input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	return a.client.DeletePlaybackConfigurationWithContext(ctx, input)
}

func (a *MediaTailorActivities) GetPlaybackConfiguration(ctx context.Context, input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	return a.client.GetPlaybackConfigurationWithContext(ctx, input)
}

func (a *MediaTailorActivities) ListPlaybackConfigurations(ctx context.Context, input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	return a.client.ListPlaybackConfigurationsWithContext(ctx, input)
}

func (a *MediaTailorActivities) ListTagsForResource(ctx context.Context, input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaTailorActivities) PutPlaybackConfiguration(ctx context.Context, input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	return a.client.PutPlaybackConfigurationWithContext(ctx, input)
}

func (a *MediaTailorActivities) TagResource(ctx context.Context, input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *MediaTailorActivities) UntagResource(ctx context.Context, input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
