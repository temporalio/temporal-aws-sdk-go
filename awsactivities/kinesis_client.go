package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"go.temporal.io/sdk/workflow"
)

type KinesisClient interface {
    AddTagsToStream(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error)
    AddTagsToStreamAsync(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) *KinesisAddTagsToStreamResult

    CreateStream(ctx workflow.Context, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error)
    CreateStreamAsync(ctx workflow.Context, input *kinesis.CreateStreamInput) *KinesisCreateStreamResult

    DecreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error)
    DecreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) *KinesisDecreaseStreamRetentionPeriodResult

    DeleteStream(ctx workflow.Context, input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error)
    DeleteStreamAsync(ctx workflow.Context, input *kinesis.DeleteStreamInput) *KinesisDeleteStreamResult

    DeregisterStreamConsumer(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error)
    DeregisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) *KinesisDeregisterStreamConsumerResult

    DescribeLimits(ctx workflow.Context, input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error)
    DescribeLimitsAsync(ctx workflow.Context, input *kinesis.DescribeLimitsInput) *KinesisDescribeLimitsResult

    DescribeStream(ctx workflow.Context, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error)
    DescribeStreamAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *KinesisDescribeStreamResult

    DescribeStreamConsumer(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error)
    DescribeStreamConsumerAsync(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) *KinesisDescribeStreamConsumerResult

    DescribeStreamSummary(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error)
    DescribeStreamSummaryAsync(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) *KinesisDescribeStreamSummaryResult

    DisableEnhancedMonitoring(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)
    DisableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) *KinesisDisableEnhancedMonitoringResult

    EnableEnhancedMonitoring(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)
    EnableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) *KinesisEnableEnhancedMonitoringResult

    GetRecords(ctx workflow.Context, input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error)
    GetRecordsAsync(ctx workflow.Context, input *kinesis.GetRecordsInput) *KinesisGetRecordsResult

    GetShardIterator(ctx workflow.Context, input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error)
    GetShardIteratorAsync(ctx workflow.Context, input *kinesis.GetShardIteratorInput) *KinesisGetShardIteratorResult

    IncreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error)
    IncreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) *KinesisIncreaseStreamRetentionPeriodResult

    ListShards(ctx workflow.Context, input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error)
    ListShardsAsync(ctx workflow.Context, input *kinesis.ListShardsInput) *KinesisListShardsResult

    ListStreamConsumers(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error)
    ListStreamConsumersAsync(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) *KinesisListStreamConsumersResult

    ListStreams(ctx workflow.Context, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error)
    ListStreamsAsync(ctx workflow.Context, input *kinesis.ListStreamsInput) *KinesisListStreamsResult

    ListTagsForStream(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error)
    ListTagsForStreamAsync(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) *KinesisListTagsForStreamResult

    MergeShards(ctx workflow.Context, input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error)
    MergeShardsAsync(ctx workflow.Context, input *kinesis.MergeShardsInput) *KinesisMergeShardsResult

    PutRecord(ctx workflow.Context, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error)
    PutRecordAsync(ctx workflow.Context, input *kinesis.PutRecordInput) *KinesisPutRecordResult

    PutRecords(ctx workflow.Context, input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error)
    PutRecordsAsync(ctx workflow.Context, input *kinesis.PutRecordsInput) *KinesisPutRecordsResult

    RegisterStreamConsumer(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error)
    RegisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) *KinesisRegisterStreamConsumerResult

    RemoveTagsFromStream(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error)
    RemoveTagsFromStreamAsync(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) *KinesisRemoveTagsFromStreamResult

    SplitShard(ctx workflow.Context, input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error)
    SplitShardAsync(ctx workflow.Context, input *kinesis.SplitShardInput) *KinesisSplitShardResult

    StartStreamEncryption(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error)
    StartStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) *KinesisStartStreamEncryptionResult

    StopStreamEncryption(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error)
    StopStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) *KinesisStopStreamEncryptionResult

    SubscribeToShard(ctx workflow.Context, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error)
    SubscribeToShardAsync(ctx workflow.Context, input *kinesis.SubscribeToShardInput) *KinesisSubscribeToShardResult

    UpdateShardCount(ctx workflow.Context, input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error)
    UpdateShardCountAsync(ctx workflow.Context, input *kinesis.UpdateShardCountInput) *KinesisUpdateShardCountResult

    WaitUntilStreamExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error
    WaitUntilStreamNotExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error}
type KinesisAddTagsToStreamResult struct {
	Result workflow.Future
}

func (r *KinesisAddTagsToStreamResult) Get(ctx workflow.Context) (*kinesis.AddTagsToStreamOutput, error) {
    var output kinesis.AddTagsToStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisCreateStreamResult struct {
	Result workflow.Future
}

func (r *KinesisCreateStreamResult) Get(ctx workflow.Context) (*kinesis.CreateStreamOutput, error) {
    var output kinesis.CreateStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDecreaseStreamRetentionPeriodResult struct {
	Result workflow.Future
}

func (r *KinesisDecreaseStreamRetentionPeriodResult) Get(ctx workflow.Context) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
    var output kinesis.DecreaseStreamRetentionPeriodOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDeleteStreamResult struct {
	Result workflow.Future
}

func (r *KinesisDeleteStreamResult) Get(ctx workflow.Context) (*kinesis.DeleteStreamOutput, error) {
    var output kinesis.DeleteStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDeregisterStreamConsumerResult struct {
	Result workflow.Future
}

func (r *KinesisDeregisterStreamConsumerResult) Get(ctx workflow.Context) (*kinesis.DeregisterStreamConsumerOutput, error) {
    var output kinesis.DeregisterStreamConsumerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDescribeLimitsResult struct {
	Result workflow.Future
}

func (r *KinesisDescribeLimitsResult) Get(ctx workflow.Context) (*kinesis.DescribeLimitsOutput, error) {
    var output kinesis.DescribeLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDescribeStreamResult struct {
	Result workflow.Future
}

func (r *KinesisDescribeStreamResult) Get(ctx workflow.Context) (*kinesis.DescribeStreamOutput, error) {
    var output kinesis.DescribeStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDescribeStreamConsumerResult struct {
	Result workflow.Future
}

func (r *KinesisDescribeStreamConsumerResult) Get(ctx workflow.Context) (*kinesis.DescribeStreamConsumerOutput, error) {
    var output kinesis.DescribeStreamConsumerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDescribeStreamSummaryResult struct {
	Result workflow.Future
}

func (r *KinesisDescribeStreamSummaryResult) Get(ctx workflow.Context) (*kinesis.DescribeStreamSummaryOutput, error) {
    var output kinesis.DescribeStreamSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisDisableEnhancedMonitoringResult struct {
	Result workflow.Future
}

func (r *KinesisDisableEnhancedMonitoringResult) Get(ctx workflow.Context) (*kinesis.EnhancedMonitoringOutput, error) {
    var output kinesis.EnhancedMonitoringOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisEnableEnhancedMonitoringResult struct {
	Result workflow.Future
}

func (r *KinesisEnableEnhancedMonitoringResult) Get(ctx workflow.Context) (*kinesis.EnhancedMonitoringOutput, error) {
    var output kinesis.EnhancedMonitoringOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisGetRecordsResult struct {
	Result workflow.Future
}

func (r *KinesisGetRecordsResult) Get(ctx workflow.Context) (*kinesis.GetRecordsOutput, error) {
    var output kinesis.GetRecordsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisGetShardIteratorResult struct {
	Result workflow.Future
}

func (r *KinesisGetShardIteratorResult) Get(ctx workflow.Context) (*kinesis.GetShardIteratorOutput, error) {
    var output kinesis.GetShardIteratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisIncreaseStreamRetentionPeriodResult struct {
	Result workflow.Future
}

func (r *KinesisIncreaseStreamRetentionPeriodResult) Get(ctx workflow.Context) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
    var output kinesis.IncreaseStreamRetentionPeriodOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisListShardsResult struct {
	Result workflow.Future
}

func (r *KinesisListShardsResult) Get(ctx workflow.Context) (*kinesis.ListShardsOutput, error) {
    var output kinesis.ListShardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisListStreamConsumersResult struct {
	Result workflow.Future
}

func (r *KinesisListStreamConsumersResult) Get(ctx workflow.Context) (*kinesis.ListStreamConsumersOutput, error) {
    var output kinesis.ListStreamConsumersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisListStreamsResult struct {
	Result workflow.Future
}

func (r *KinesisListStreamsResult) Get(ctx workflow.Context) (*kinesis.ListStreamsOutput, error) {
    var output kinesis.ListStreamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisListTagsForStreamResult struct {
	Result workflow.Future
}

func (r *KinesisListTagsForStreamResult) Get(ctx workflow.Context) (*kinesis.ListTagsForStreamOutput, error) {
    var output kinesis.ListTagsForStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisMergeShardsResult struct {
	Result workflow.Future
}

func (r *KinesisMergeShardsResult) Get(ctx workflow.Context) (*kinesis.MergeShardsOutput, error) {
    var output kinesis.MergeShardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisPutRecordResult struct {
	Result workflow.Future
}

func (r *KinesisPutRecordResult) Get(ctx workflow.Context) (*kinesis.PutRecordOutput, error) {
    var output kinesis.PutRecordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisPutRecordsResult struct {
	Result workflow.Future
}

func (r *KinesisPutRecordsResult) Get(ctx workflow.Context) (*kinesis.PutRecordsOutput, error) {
    var output kinesis.PutRecordsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisRegisterStreamConsumerResult struct {
	Result workflow.Future
}

func (r *KinesisRegisterStreamConsumerResult) Get(ctx workflow.Context) (*kinesis.RegisterStreamConsumerOutput, error) {
    var output kinesis.RegisterStreamConsumerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisRemoveTagsFromStreamResult struct {
	Result workflow.Future
}

func (r *KinesisRemoveTagsFromStreamResult) Get(ctx workflow.Context) (*kinesis.RemoveTagsFromStreamOutput, error) {
    var output kinesis.RemoveTagsFromStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisSplitShardResult struct {
	Result workflow.Future
}

func (r *KinesisSplitShardResult) Get(ctx workflow.Context) (*kinesis.SplitShardOutput, error) {
    var output kinesis.SplitShardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisStartStreamEncryptionResult struct {
	Result workflow.Future
}

func (r *KinesisStartStreamEncryptionResult) Get(ctx workflow.Context) (*kinesis.StartStreamEncryptionOutput, error) {
    var output kinesis.StartStreamEncryptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisStopStreamEncryptionResult struct {
	Result workflow.Future
}

func (r *KinesisStopStreamEncryptionResult) Get(ctx workflow.Context) (*kinesis.StopStreamEncryptionOutput, error) {
    var output kinesis.StopStreamEncryptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisSubscribeToShardResult struct {
	Result workflow.Future
}

func (r *KinesisSubscribeToShardResult) Get(ctx workflow.Context) (*kinesis.SubscribeToShardOutput, error) {
    var output kinesis.SubscribeToShardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisUpdateShardCountResult struct {
	Result workflow.Future
}

func (r *KinesisUpdateShardCountResult) Get(ctx workflow.Context) (*kinesis.UpdateShardCountOutput, error) {
    var output kinesis.UpdateShardCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type KinesisStub struct {
    activities KinesisClient
}

func NewKinesisStub() KinesisClient {
    return &KinesisStub{}
}

func (a *KinesisStub) AddTagsToStream(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
    var output kinesis.AddTagsToStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToStream, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) CreateStream(ctx workflow.Context, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {
    var output kinesis.CreateStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStream, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DecreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
    var output kinesis.DecreaseStreamRetentionPeriodOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DecreaseStreamRetentionPeriod, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DeleteStream(ctx workflow.Context, input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
    var output kinesis.DeleteStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStream, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DeregisterStreamConsumer(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error) {
    var output kinesis.DeregisterStreamConsumerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterStreamConsumer, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DescribeLimits(ctx workflow.Context, input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
    var output kinesis.DescribeLimitsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLimits, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DescribeStream(ctx workflow.Context, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
    var output kinesis.DescribeStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStream, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DescribeStreamConsumer(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
    var output kinesis.DescribeStreamConsumerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStreamConsumer, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DescribeStreamSummary(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
    var output kinesis.DescribeStreamSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStreamSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) DisableEnhancedMonitoring(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
    var output kinesis.EnhancedMonitoringOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableEnhancedMonitoring, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) EnableEnhancedMonitoring(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
    var output kinesis.EnhancedMonitoringOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableEnhancedMonitoring, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) GetRecords(ctx workflow.Context, input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
    var output kinesis.GetRecordsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRecords, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) GetShardIterator(ctx workflow.Context, input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error) {
    var output kinesis.GetShardIteratorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetShardIterator, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) IncreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
    var output kinesis.IncreaseStreamRetentionPeriodOutput
    err := workflow.ExecuteActivity(ctx, a.activities.IncreaseStreamRetentionPeriod, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) ListShards(ctx workflow.Context, input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
    var output kinesis.ListShardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListShards, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) ListStreamConsumers(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
    var output kinesis.ListStreamConsumersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreamConsumers, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) ListStreams(ctx workflow.Context, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
    var output kinesis.ListStreamsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) ListTagsForStream(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
    var output kinesis.ListTagsForStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForStream, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) MergeShards(ctx workflow.Context, input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error) {
    var output kinesis.MergeShardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.MergeShards, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) PutRecord(ctx workflow.Context, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
    var output kinesis.PutRecordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRecord, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) PutRecords(ctx workflow.Context, input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error) {
    var output kinesis.PutRecordsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRecords, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) RegisterStreamConsumer(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error) {
    var output kinesis.RegisterStreamConsumerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterStreamConsumer, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) RemoveTagsFromStream(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error) {
    var output kinesis.RemoveTagsFromStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromStream, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) SplitShard(ctx workflow.Context, input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error) {
    var output kinesis.SplitShardOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SplitShard, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) StartStreamEncryption(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error) {
    var output kinesis.StartStreamEncryptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartStreamEncryption, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) StopStreamEncryption(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error) {
    var output kinesis.StopStreamEncryptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopStreamEncryption, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) SubscribeToShard(ctx workflow.Context, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
    var output kinesis.SubscribeToShardOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubscribeToShard, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisStub) UpdateShardCount(ctx workflow.Context, input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error) {
    var output kinesis.UpdateShardCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateShardCount, input).Get(ctx, &output)
    return &output, err
}


func (a *KinesisStub) WaitUntilStreamExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error {
    return a.client.WaitUntilStreamExists(input)
}


func (a *KinesisStub) WaitUntilStreamNotExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error {
    return a.client.WaitUntilStreamNotExists(input)
}
