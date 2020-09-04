
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
)

type FirehoseActivities struct {
	client firehoseiface.FirehoseAPI
}

func NewFirehoseActivities(client firehoseiface.FirehoseAPI) *FirehoseActivities {
    return &FirehoseActivities{client: client}
}

func (a *FirehoseActivities) CreateDeliveryStream(input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
    return a.client.CreateDeliveryStream(input)
}

func (a *FirehoseActivities) DeleteDeliveryStream(input *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error) {
    return a.client.DeleteDeliveryStream(input)
}

func (a *FirehoseActivities) DescribeDeliveryStream(input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error) {
    return a.client.DescribeDeliveryStream(input)
}

func (a *FirehoseActivities) ListDeliveryStreams(input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error) {
    return a.client.ListDeliveryStreams(input)
}

func (a *FirehoseActivities) ListTagsForDeliveryStream(input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error) {
    return a.client.ListTagsForDeliveryStream(input)
}

func (a *FirehoseActivities) PutRecord(input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error) {
    return a.client.PutRecord(input)
}

func (a *FirehoseActivities) PutRecordBatch(input *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error) {
    return a.client.PutRecordBatch(input)
}

func (a *FirehoseActivities) StartDeliveryStreamEncryption(input *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
    return a.client.StartDeliveryStreamEncryption(input)
}

func (a *FirehoseActivities) StopDeliveryStreamEncryption(input *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
    return a.client.StopDeliveryStreamEncryption(input)
}

func (a *FirehoseActivities) TagDeliveryStream(input *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error) {
    return a.client.TagDeliveryStream(input)
}

func (a *FirehoseActivities) UntagDeliveryStream(input *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error) {
    return a.client.UntagDeliveryStream(input)
}

func (a *FirehoseActivities) UpdateDestination(input *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error) {
    return a.client.UpdateDestination(input)
}
