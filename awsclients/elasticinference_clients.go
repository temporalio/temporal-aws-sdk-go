package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"go.temporal.io/sdk/workflow"
)

type ElasticInferenceClient interface {
       DescribeAcceleratorOfferings(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error)
       DescribeAcceleratorOfferingsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) *ElasticinferenceDescribeAcceleratorOfferingsResult

       DescribeAcceleratorTypes(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error)
       DescribeAcceleratorTypesAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) *ElasticinferenceDescribeAcceleratorTypesResult

       DescribeAccelerators(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error)
       DescribeAcceleratorsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) *ElasticinferenceDescribeAcceleratorsResult

       ListTagsForResource(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) *ElasticinferenceListTagsForResourceResult

       TagResource(ctx workflow.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *elasticinference.TagResourceInput) *ElasticinferenceTagResourceResult

       UntagResource(ctx workflow.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *elasticinference.UntagResourceInput) *ElasticinferenceUntagResourceResult
}

type ElasticInferenceStub struct {
}

func NewElasticInferenceStub() ElasticInferenceClient {
    return &ElasticInferenceStub{}
}

type ElasticinferenceDescribeAcceleratorOfferingsResult struct {
	Result workflow.Future
}

func (r *ElasticinferenceDescribeAcceleratorOfferingsResult) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
    var output elasticinference.DescribeAcceleratorOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticinferenceDescribeAcceleratorTypesResult struct {
	Result workflow.Future
}

func (r *ElasticinferenceDescribeAcceleratorTypesResult) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
    var output elasticinference.DescribeAcceleratorTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticinferenceDescribeAcceleratorsResult struct {
	Result workflow.Future
}

func (r *ElasticinferenceDescribeAcceleratorsResult) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorsOutput, error) {
    var output elasticinference.DescribeAcceleratorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticinferenceListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ElasticinferenceListTagsForResourceResult) Get(ctx workflow.Context) (*elasticinference.ListTagsForResourceOutput, error) {
    var output elasticinference.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticinferenceTagResourceResult struct {
	Result workflow.Future
}

func (r *ElasticinferenceTagResourceResult) Get(ctx workflow.Context) (*elasticinference.TagResourceOutput, error) {
    var output elasticinference.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ElasticinferenceUntagResourceResult struct {
	Result workflow.Future
}

func (r *ElasticinferenceUntagResourceResult) Get(ctx workflow.Context) (*elasticinference.UntagResourceOutput, error) {
    var output elasticinference.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorOfferings(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
    var output elasticinference.DescribeAcceleratorOfferingsOutput
    err := workflow.ExecuteActivity(ctx, "ElasticInference.DescribeAcceleratorOfferings", input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorOfferingsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) *ElasticinferenceDescribeAcceleratorOfferingsResult {
    future := workflow.ExecuteActivity(ctx, "ElasticInference.DescribeAcceleratorOfferings", input)
    return &ElasticinferenceDescribeAcceleratorOfferingsResult{Result: future}
}

func (a *ElasticInferenceStub) DescribeAcceleratorTypes(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
    var output elasticinference.DescribeAcceleratorTypesOutput
    err := workflow.ExecuteActivity(ctx, "ElasticInference.DescribeAcceleratorTypes", input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorTypesAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) *ElasticinferenceDescribeAcceleratorTypesResult {
    future := workflow.ExecuteActivity(ctx, "ElasticInference.DescribeAcceleratorTypes", input)
    return &ElasticinferenceDescribeAcceleratorTypesResult{Result: future}
}

func (a *ElasticInferenceStub) DescribeAccelerators(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error) {
    var output elasticinference.DescribeAcceleratorsOutput
    err := workflow.ExecuteActivity(ctx, "ElasticInference.DescribeAccelerators", input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) *ElasticinferenceDescribeAcceleratorsResult {
    future := workflow.ExecuteActivity(ctx, "ElasticInference.DescribeAccelerators", input)
    return &ElasticinferenceDescribeAcceleratorsResult{Result: future}
}

func (a *ElasticInferenceStub) ListTagsForResource(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error) {
    var output elasticinference.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "ElasticInference.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticInferenceStub) ListTagsForResourceAsync(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) *ElasticinferenceListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "ElasticInference.ListTagsForResource", input)
    return &ElasticinferenceListTagsForResourceResult{Result: future}
}

func (a *ElasticInferenceStub) TagResource(ctx workflow.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error) {
    var output elasticinference.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "ElasticInference.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticInferenceStub) TagResourceAsync(ctx workflow.Context, input *elasticinference.TagResourceInput) *ElasticinferenceTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "ElasticInference.TagResource", input)
    return &ElasticinferenceTagResourceResult{Result: future}
}

func (a *ElasticInferenceStub) UntagResource(ctx workflow.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error) {
    var output elasticinference.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "ElasticInference.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticInferenceStub) UntagResourceAsync(ctx workflow.Context, input *elasticinference.UntagResourceInput) *ElasticinferenceUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "ElasticInference.UntagResource", input)
    return &ElasticinferenceUntagResourceResult{Result: future}
}
