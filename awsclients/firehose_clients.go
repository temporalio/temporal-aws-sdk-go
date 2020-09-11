package awsclients

import (
	"github.com/aws/aws-sdk-go/service/firehose"
	"go.temporal.io/sdk/workflow"
)

type FirehoseClient interface {
       CreateDeliveryStream(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error)
       CreateDeliveryStreamAsync(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) *FirehoseCreateDeliveryStreamResult

       DeleteDeliveryStream(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error)
       DeleteDeliveryStreamAsync(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) *FirehoseDeleteDeliveryStreamResult

       DescribeDeliveryStream(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error)
       DescribeDeliveryStreamAsync(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) *FirehoseDescribeDeliveryStreamResult

       ListDeliveryStreams(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error)
       ListDeliveryStreamsAsync(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) *FirehoseListDeliveryStreamsResult

       ListTagsForDeliveryStream(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error)
       ListTagsForDeliveryStreamAsync(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) *FirehoseListTagsForDeliveryStreamResult

       PutRecord(ctx workflow.Context, input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error)
       PutRecordAsync(ctx workflow.Context, input *firehose.PutRecordInput) *FirehosePutRecordResult

       PutRecordBatch(ctx workflow.Context, input *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error)
       PutRecordBatchAsync(ctx workflow.Context, input *firehose.PutRecordBatchInput) *FirehosePutRecordBatchResult

       StartDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error)
       StartDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) *FirehoseStartDeliveryStreamEncryptionResult

       StopDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error)
       StopDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) *FirehoseStopDeliveryStreamEncryptionResult

       TagDeliveryStream(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error)
       TagDeliveryStreamAsync(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) *FirehoseTagDeliveryStreamResult

       UntagDeliveryStream(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error)
       UntagDeliveryStreamAsync(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) *FirehoseUntagDeliveryStreamResult

       UpdateDestination(ctx workflow.Context, input *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error)
       UpdateDestinationAsync(ctx workflow.Context, input *firehose.UpdateDestinationInput) *FirehoseUpdateDestinationResult
}

type FirehoseStub struct {
}

func NewFirehoseStub() FirehoseClient {
    return &FirehoseStub{}
}

type FirehoseCreateDeliveryStreamResult struct {
	Result workflow.Future
}

func (r *FirehoseCreateDeliveryStreamResult) Get(ctx workflow.Context) (*firehose.CreateDeliveryStreamOutput, error) {
    var output firehose.CreateDeliveryStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseDeleteDeliveryStreamResult struct {
	Result workflow.Future
}

func (r *FirehoseDeleteDeliveryStreamResult) Get(ctx workflow.Context) (*firehose.DeleteDeliveryStreamOutput, error) {
    var output firehose.DeleteDeliveryStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseDescribeDeliveryStreamResult struct {
	Result workflow.Future
}

func (r *FirehoseDescribeDeliveryStreamResult) Get(ctx workflow.Context) (*firehose.DescribeDeliveryStreamOutput, error) {
    var output firehose.DescribeDeliveryStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseListDeliveryStreamsResult struct {
	Result workflow.Future
}

func (r *FirehoseListDeliveryStreamsResult) Get(ctx workflow.Context) (*firehose.ListDeliveryStreamsOutput, error) {
    var output firehose.ListDeliveryStreamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseListTagsForDeliveryStreamResult struct {
	Result workflow.Future
}

func (r *FirehoseListTagsForDeliveryStreamResult) Get(ctx workflow.Context) (*firehose.ListTagsForDeliveryStreamOutput, error) {
    var output firehose.ListTagsForDeliveryStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehosePutRecordResult struct {
	Result workflow.Future
}

func (r *FirehosePutRecordResult) Get(ctx workflow.Context) (*firehose.PutRecordOutput, error) {
    var output firehose.PutRecordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehosePutRecordBatchResult struct {
	Result workflow.Future
}

func (r *FirehosePutRecordBatchResult) Get(ctx workflow.Context) (*firehose.PutRecordBatchOutput, error) {
    var output firehose.PutRecordBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseStartDeliveryStreamEncryptionResult struct {
	Result workflow.Future
}

func (r *FirehoseStartDeliveryStreamEncryptionResult) Get(ctx workflow.Context) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
    var output firehose.StartDeliveryStreamEncryptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseStopDeliveryStreamEncryptionResult struct {
	Result workflow.Future
}

func (r *FirehoseStopDeliveryStreamEncryptionResult) Get(ctx workflow.Context) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
    var output firehose.StopDeliveryStreamEncryptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseTagDeliveryStreamResult struct {
	Result workflow.Future
}

func (r *FirehoseTagDeliveryStreamResult) Get(ctx workflow.Context) (*firehose.TagDeliveryStreamOutput, error) {
    var output firehose.TagDeliveryStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseUntagDeliveryStreamResult struct {
	Result workflow.Future
}

func (r *FirehoseUntagDeliveryStreamResult) Get(ctx workflow.Context) (*firehose.UntagDeliveryStreamOutput, error) {
    var output firehose.UntagDeliveryStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type FirehoseUpdateDestinationResult struct {
	Result workflow.Future
}

func (r *FirehoseUpdateDestinationResult) Get(ctx workflow.Context) (*firehose.UpdateDestinationOutput, error) {
    var output firehose.UpdateDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) CreateDeliveryStream(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
    var output firehose.CreateDeliveryStreamOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.CreateDeliveryStream", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) CreateDeliveryStreamAsync(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) *FirehoseCreateDeliveryStreamResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.CreateDeliveryStream", input)
    return &FirehoseCreateDeliveryStreamResult{Result: future}
}

func (a *FirehoseStub) DeleteDeliveryStream(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error) {
    var output firehose.DeleteDeliveryStreamOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.DeleteDeliveryStream", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) DeleteDeliveryStreamAsync(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) *FirehoseDeleteDeliveryStreamResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.DeleteDeliveryStream", input)
    return &FirehoseDeleteDeliveryStreamResult{Result: future}
}

func (a *FirehoseStub) DescribeDeliveryStream(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error) {
    var output firehose.DescribeDeliveryStreamOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.DescribeDeliveryStream", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) DescribeDeliveryStreamAsync(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) *FirehoseDescribeDeliveryStreamResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.DescribeDeliveryStream", input)
    return &FirehoseDescribeDeliveryStreamResult{Result: future}
}

func (a *FirehoseStub) ListDeliveryStreams(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error) {
    var output firehose.ListDeliveryStreamsOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.ListDeliveryStreams", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) ListDeliveryStreamsAsync(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) *FirehoseListDeliveryStreamsResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.ListDeliveryStreams", input)
    return &FirehoseListDeliveryStreamsResult{Result: future}
}

func (a *FirehoseStub) ListTagsForDeliveryStream(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error) {
    var output firehose.ListTagsForDeliveryStreamOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.ListTagsForDeliveryStream", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) ListTagsForDeliveryStreamAsync(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) *FirehoseListTagsForDeliveryStreamResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.ListTagsForDeliveryStream", input)
    return &FirehoseListTagsForDeliveryStreamResult{Result: future}
}

func (a *FirehoseStub) PutRecord(ctx workflow.Context, input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error) {
    var output firehose.PutRecordOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.PutRecord", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) PutRecordAsync(ctx workflow.Context, input *firehose.PutRecordInput) *FirehosePutRecordResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.PutRecord", input)
    return &FirehosePutRecordResult{Result: future}
}

func (a *FirehoseStub) PutRecordBatch(ctx workflow.Context, input *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error) {
    var output firehose.PutRecordBatchOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.PutRecordBatch", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) PutRecordBatchAsync(ctx workflow.Context, input *firehose.PutRecordBatchInput) *FirehosePutRecordBatchResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.PutRecordBatch", input)
    return &FirehosePutRecordBatchResult{Result: future}
}

func (a *FirehoseStub) StartDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
    var output firehose.StartDeliveryStreamEncryptionOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.StartDeliveryStreamEncryption", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) StartDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) *FirehoseStartDeliveryStreamEncryptionResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.StartDeliveryStreamEncryption", input)
    return &FirehoseStartDeliveryStreamEncryptionResult{Result: future}
}

func (a *FirehoseStub) StopDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
    var output firehose.StopDeliveryStreamEncryptionOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.StopDeliveryStreamEncryption", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) StopDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) *FirehoseStopDeliveryStreamEncryptionResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.StopDeliveryStreamEncryption", input)
    return &FirehoseStopDeliveryStreamEncryptionResult{Result: future}
}

func (a *FirehoseStub) TagDeliveryStream(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error) {
    var output firehose.TagDeliveryStreamOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.TagDeliveryStream", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) TagDeliveryStreamAsync(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) *FirehoseTagDeliveryStreamResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.TagDeliveryStream", input)
    return &FirehoseTagDeliveryStreamResult{Result: future}
}

func (a *FirehoseStub) UntagDeliveryStream(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error) {
    var output firehose.UntagDeliveryStreamOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.UntagDeliveryStream", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) UntagDeliveryStreamAsync(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) *FirehoseUntagDeliveryStreamResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.UntagDeliveryStream", input)
    return &FirehoseUntagDeliveryStreamResult{Result: future}
}

func (a *FirehoseStub) UpdateDestination(ctx workflow.Context, input *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error) {
    var output firehose.UpdateDestinationOutput
    err := workflow.ExecuteActivity(ctx, "Firehose.UpdateDestination", input).Get(ctx, &output)
    return &output, err
}

func (a *FirehoseStub) UpdateDestinationAsync(ctx workflow.Context, input *firehose.UpdateDestinationInput) *FirehoseUpdateDestinationResult {
    future := workflow.ExecuteActivity(ctx, "Firehose.UpdateDestination", input)
    return &FirehoseUpdateDestinationResult{Result: future}
}
