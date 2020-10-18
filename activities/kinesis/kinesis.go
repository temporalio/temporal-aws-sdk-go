// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client kinesisiface.KinesisAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := kinesis.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (kinesisiface.KinesisAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return kinesis.New(sess), nil
}

func (a *Activities) AddTagsToStream(ctx context.Context, input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsToStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStream(ctx context.Context, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DecreaseStreamRetentionPeriod(ctx context.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DecreaseStreamRetentionPeriodWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteStream(ctx context.Context, input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterStreamConsumer(ctx context.Context, input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterStreamConsumerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLimits(ctx context.Context, input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLimitsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStream(ctx context.Context, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStreamConsumer(ctx context.Context, input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStreamConsumerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStreamSummary(ctx context.Context, input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStreamSummaryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisableEnhancedMonitoring(ctx context.Context, input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisableEnhancedMonitoringWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) EnableEnhancedMonitoring(ctx context.Context, input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.EnableEnhancedMonitoringWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRecords(ctx context.Context, input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRecordsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetShardIterator(ctx context.Context, input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetShardIteratorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) IncreaseStreamRetentionPeriod(ctx context.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.IncreaseStreamRetentionPeriodWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListShards(ctx context.Context, input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListShardsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListStreamConsumers(ctx context.Context, input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListStreamConsumersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListStreams(ctx context.Context, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListStreamsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForStream(ctx context.Context, input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) MergeShards(ctx context.Context, input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.MergeShardsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRecord(ctx context.Context, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRecordWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRecords(ctx context.Context, input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRecordsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterStreamConsumer(ctx context.Context, input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterStreamConsumerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTagsFromStream(ctx context.Context, input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsFromStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SplitShard(ctx context.Context, input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SplitShardWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartStreamEncryption(ctx context.Context, input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartStreamEncryptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopStreamEncryption(ctx context.Context, input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopStreamEncryptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SubscribeToShard(ctx context.Context, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SubscribeToShardWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateShardCount(ctx context.Context, input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateShardCountWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilStreamExists(ctx context.Context, input *kinesis.DescribeStreamInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilStreamExistsWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilStreamNotExists(ctx context.Context, input *kinesis.DescribeStreamInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilStreamNotExistsWithContext(ctx, input, options...))
	})
}
