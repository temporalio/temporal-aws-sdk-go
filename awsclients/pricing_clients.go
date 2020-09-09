package awsclients

import (
	"github.com/aws/aws-sdk-go/service/pricing"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type PricingClient interface {
       DescribeServices(ctx workflow.Context, input *pricing.DescribeServicesInput) (*pricing.DescribeServicesOutput, error)
       DescribeServicesAsync(ctx workflow.Context, input *pricing.DescribeServicesInput) *PricingDescribeServicesResult

       GetAttributeValues(ctx workflow.Context, input *pricing.GetAttributeValuesInput) (*pricing.GetAttributeValuesOutput, error)
       GetAttributeValuesAsync(ctx workflow.Context, input *pricing.GetAttributeValuesInput) *PricingGetAttributeValuesResult

       GetProducts(ctx workflow.Context, input *pricing.GetProductsInput) (*pricing.GetProductsOutput, error)
       GetProductsAsync(ctx workflow.Context, input *pricing.GetProductsInput) *PricingGetProductsResult
}

type PricingDescribeServicesResult struct {
	Result workflow.Future
}

func (r *PricingDescribeServicesResult) Get(ctx workflow.Context) (*pricing.DescribeServicesOutput, error) {
    var output pricing.DescribeServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PricingGetAttributeValuesResult struct {
	Result workflow.Future
}

func (r *PricingGetAttributeValuesResult) Get(ctx workflow.Context) (*pricing.GetAttributeValuesOutput, error) {
    var output pricing.GetAttributeValuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PricingGetProductsResult struct {
	Result workflow.Future
}

func (r *PricingGetProductsResult) Get(ctx workflow.Context) (*pricing.GetProductsOutput, error) {
    var output pricing.GetProductsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PricingStub struct {
    activities awsactivities.PricingActivities
}

func NewPricingStub() PricingClient {
    return &PricingStub{}
}

func (a *PricingStub) DescribeServices(ctx workflow.Context, input *pricing.DescribeServicesInput) (*pricing.DescribeServicesOutput, error) {
    var output pricing.DescribeServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServices, input).Get(ctx, &output)
    return &output, err
}

func (a *PricingStub) DescribeServicesAsync(ctx workflow.Context, input *pricing.DescribeServicesInput) *PricingDescribeServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeServices, input)
    return &PricingDescribeServicesResult{Result: future}
}

func (a *PricingStub) GetAttributeValues(ctx workflow.Context, input *pricing.GetAttributeValuesInput) (*pricing.GetAttributeValuesOutput, error) {
    var output pricing.GetAttributeValuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAttributeValues, input).Get(ctx, &output)
    return &output, err
}

func (a *PricingStub) GetAttributeValuesAsync(ctx workflow.Context, input *pricing.GetAttributeValuesInput) *PricingGetAttributeValuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAttributeValues, input)
    return &PricingGetAttributeValuesResult{Result: future}
}

func (a *PricingStub) GetProducts(ctx workflow.Context, input *pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
    var output pricing.GetProductsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProducts, input).Get(ctx, &output)
    return &output, err
}

func (a *PricingStub) GetProductsAsync(ctx workflow.Context, input *pricing.GetProductsInput) *PricingGetProductsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetProducts, input)
    return &PricingGetProductsResult{Result: future}
}
