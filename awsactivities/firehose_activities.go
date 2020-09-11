package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type FirehoseActivities struct {
	client firehoseiface.FirehoseAPI
}

func NewFirehoseActivities(session *session.Session, config ...*aws.Config) *FirehoseActivities {
	client := firehose.New(session, config...)
	return &FirehoseActivities{client: client}
}

func (a *FirehoseActivities) CreateDeliveryStream(ctx context.Context, input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
	return a.client.CreateDeliveryStreamWithContext(ctx, input)
}

func (a *FirehoseActivities) DeleteDeliveryStream(ctx context.Context, input *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error) {
	return a.client.DeleteDeliveryStreamWithContext(ctx, input)
}

func (a *FirehoseActivities) DescribeDeliveryStream(ctx context.Context, input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error) {
	return a.client.DescribeDeliveryStreamWithContext(ctx, input)
}

func (a *FirehoseActivities) ListDeliveryStreams(ctx context.Context, input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error) {
	return a.client.ListDeliveryStreamsWithContext(ctx, input)
}

func (a *FirehoseActivities) ListTagsForDeliveryStream(ctx context.Context, input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	return a.client.ListTagsForDeliveryStreamWithContext(ctx, input)
}

func (a *FirehoseActivities) PutRecord(ctx context.Context, input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error) {
	return a.client.PutRecordWithContext(ctx, input)
}

func (a *FirehoseActivities) PutRecordBatch(ctx context.Context, input *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error) {
	return a.client.PutRecordBatchWithContext(ctx, input)
}

func (a *FirehoseActivities) StartDeliveryStreamEncryption(ctx context.Context, input *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
	return a.client.StartDeliveryStreamEncryptionWithContext(ctx, input)
}

func (a *FirehoseActivities) StopDeliveryStreamEncryption(ctx context.Context, input *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
	return a.client.StopDeliveryStreamEncryptionWithContext(ctx, input)
}

func (a *FirehoseActivities) TagDeliveryStream(ctx context.Context, input *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error) {
	return a.client.TagDeliveryStreamWithContext(ctx, input)
}

func (a *FirehoseActivities) UntagDeliveryStream(ctx context.Context, input *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error) {
	return a.client.UntagDeliveryStreamWithContext(ctx, input)
}

func (a *FirehoseActivities) UpdateDestination(ctx context.Context, input *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error) {
	return a.client.UpdateDestinationWithContext(ctx, input)
}
