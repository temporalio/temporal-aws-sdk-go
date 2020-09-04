package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

type KinesisActivities struct {
	client kinesisiface.KinesisAPI
}

func NewKinesisActivities(client kinesisiface.KinesisAPI) *KinesisActivities {
    return &KinesisActivities{client: client}
}


func (a *KinesisActivities) AddTagsToStream(input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
    return a.client.AddTagsToStream(input)
}



func (a *KinesisActivities) CreateStream(input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {
    return a.client.CreateStream(input)
}



func (a *KinesisActivities) DecreaseStreamRetentionPeriod(input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
    return a.client.DecreaseStreamRetentionPeriod(input)
}



func (a *KinesisActivities) DeleteStream(input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
    return a.client.DeleteStream(input)
}



func (a *KinesisActivities) DeregisterStreamConsumer(input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error) {
    return a.client.DeregisterStreamConsumer(input)
}



func (a *KinesisActivities) DescribeLimits(input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
    return a.client.DescribeLimits(input)
}



func (a *KinesisActivities) DescribeStream(input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
    return a.client.DescribeStream(input)
}



func (a *KinesisActivities) DescribeStreamConsumer(input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
    return a.client.DescribeStreamConsumer(input)
}



func (a *KinesisActivities) DescribeStreamSummary(input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
    return a.client.DescribeStreamSummary(input)
}



func (a *KinesisActivities) DisableEnhancedMonitoring(input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
    return a.client.DisableEnhancedMonitoring(input)
}



func (a *KinesisActivities) EnableEnhancedMonitoring(input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
    return a.client.EnableEnhancedMonitoring(input)
}



func (a *KinesisActivities) GetRecords(input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
    return a.client.GetRecords(input)
}



func (a *KinesisActivities) GetShardIterator(input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error) {
    return a.client.GetShardIterator(input)
}



func (a *KinesisActivities) IncreaseStreamRetentionPeriod(input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
    return a.client.IncreaseStreamRetentionPeriod(input)
}



func (a *KinesisActivities) ListShards(input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
    return a.client.ListShards(input)
}



func (a *KinesisActivities) ListStreamConsumers(input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
    return a.client.ListStreamConsumers(input)
}



func (a *KinesisActivities) ListStreams(input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
    return a.client.ListStreams(input)
}



func (a *KinesisActivities) ListTagsForStream(input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
    return a.client.ListTagsForStream(input)
}



func (a *KinesisActivities) MergeShards(input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error) {
    return a.client.MergeShards(input)
}



func (a *KinesisActivities) PutRecord(input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
    return a.client.PutRecord(input)
}



func (a *KinesisActivities) PutRecords(input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error) {
    return a.client.PutRecords(input)
}



func (a *KinesisActivities) RegisterStreamConsumer(input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error) {
    return a.client.RegisterStreamConsumer(input)
}



func (a *KinesisActivities) RemoveTagsFromStream(input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error) {
    return a.client.RemoveTagsFromStream(input)
}



func (a *KinesisActivities) SplitShard(input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error) {
    return a.client.SplitShard(input)
}



func (a *KinesisActivities) StartStreamEncryption(input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error) {
    return a.client.StartStreamEncryption(input)
}



func (a *KinesisActivities) StopStreamEncryption(input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error) {
    return a.client.StopStreamEncryption(input)
}



func (a *KinesisActivities) SubscribeToShard(input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
    return a.client.SubscribeToShard(input)
}



func (a *KinesisActivities) UpdateShardCount(input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error) {
    return a.client.UpdateShardCount(input)
}



func (a *KinesisActivities) WaitUntilStreamExists(input *kinesis.DescribeStreamInput) error {
    return a.client.WaitUntilStreamExists(input)
}



func (a *KinesisActivities) WaitUntilStreamNotExists(input *kinesis.DescribeStreamInput) error {
    return a.client.WaitUntilStreamNotExists(input)
}

