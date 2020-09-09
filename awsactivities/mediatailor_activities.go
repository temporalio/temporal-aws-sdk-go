
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface"
)

type MediaTailorActivities struct {
	client mediatailoriface.MediaTailorAPI
}

func NewMediaTailorActivities(session *session.Session, config... *aws.Config) *MediaTailorActivities {
    client := mediatailor.New(session, config...)
    return &MediaTailorActivities{client: client}
}

func (a *MediaTailorActivities) DeletePlaybackConfiguration(input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
    return a.client.DeletePlaybackConfiguration(input)
}

func (a *MediaTailorActivities) GetPlaybackConfiguration(input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
    return a.client.GetPlaybackConfiguration(input)
}

func (a *MediaTailorActivities) ListPlaybackConfigurations(input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
    return a.client.ListPlaybackConfigurations(input)
}

func (a *MediaTailorActivities) ListTagsForResource(input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *MediaTailorActivities) PutPlaybackConfiguration(input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error) {
    return a.client.PutPlaybackConfiguration(input)
}

func (a *MediaTailorActivities) TagResource(input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *MediaTailorActivities) UntagResource(input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
