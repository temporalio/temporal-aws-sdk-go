
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface"
)

type KinesisVideoMediaActivities struct {
	client kinesisvideomediaiface.KinesisVideoMediaAPI
}

func NewKinesisVideoMediaActivities(client kinesisvideomediaiface.KinesisVideoMediaAPI) *KinesisVideoMediaActivities {
    return &KinesisVideoMediaActivities{client: client}
}

func (a *KinesisVideoMediaActivities) GetMedia(input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
    return a.client.GetMedia(input)
}
