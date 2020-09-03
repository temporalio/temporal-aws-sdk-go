package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"go.temporal.io/sdk/workflow"
)

type MarketplaceMeteringClient interface {
    BatchMeterUsage(ctx workflow.Context, input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error)
    BatchMeterUsageAsync(ctx workflow.Context, input *marketplacemetering.BatchMeterUsageInput) *MarketplacemeteringBatchMeterUsageResult

    MeterUsage(ctx workflow.Context, input *marketplacemetering.MeterUsageInput) (*marketplacemetering.MeterUsageOutput, error)
    MeterUsageAsync(ctx workflow.Context, input *marketplacemetering.MeterUsageInput) *MarketplacemeteringMeterUsageResult

    RegisterUsage(ctx workflow.Context, input *marketplacemetering.RegisterUsageInput) (*marketplacemetering.RegisterUsageOutput, error)
    RegisterUsageAsync(ctx workflow.Context, input *marketplacemetering.RegisterUsageInput) *MarketplacemeteringRegisterUsageResult

    ResolveCustomer(ctx workflow.Context, input *marketplacemetering.ResolveCustomerInput) (*marketplacemetering.ResolveCustomerOutput, error)
    ResolveCustomerAsync(ctx workflow.Context, input *marketplacemetering.ResolveCustomerInput) *MarketplacemeteringResolveCustomerResult
}
type MarketplacemeteringBatchMeterUsageResult struct {
	Result workflow.Future
}

func (r *MarketplacemeteringBatchMeterUsageResult) Get(ctx workflow.Context) (*marketplacemetering.BatchMeterUsageOutput, error) {
    var output marketplacemetering.BatchMeterUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacemeteringMeterUsageResult struct {
	Result workflow.Future
}

func (r *MarketplacemeteringMeterUsageResult) Get(ctx workflow.Context) (*marketplacemetering.MeterUsageOutput, error) {
    var output marketplacemetering.MeterUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacemeteringRegisterUsageResult struct {
	Result workflow.Future
}

func (r *MarketplacemeteringRegisterUsageResult) Get(ctx workflow.Context) (*marketplacemetering.RegisterUsageOutput, error) {
    var output marketplacemetering.RegisterUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacemeteringResolveCustomerResult struct {
	Result workflow.Future
}

func (r *MarketplacemeteringResolveCustomerResult) Get(ctx workflow.Context) (*marketplacemetering.ResolveCustomerOutput, error) {
    var output marketplacemetering.ResolveCustomerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MarketplaceMeteringStub struct {
    activities MarketplaceMeteringClient
}

func NewMarketplaceMeteringStub() MarketplaceMeteringClient {
    return &MarketplaceMeteringStub{}
}

func (a *MarketplaceMeteringStub) BatchMeterUsage(ctx workflow.Context, input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error) {
    var output marketplacemetering.BatchMeterUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchMeterUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceMeteringStub) MeterUsage(ctx workflow.Context, input *marketplacemetering.MeterUsageInput) (*marketplacemetering.MeterUsageOutput, error) {
    var output marketplacemetering.MeterUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.MeterUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceMeteringStub) RegisterUsage(ctx workflow.Context, input *marketplacemetering.RegisterUsageInput) (*marketplacemetering.RegisterUsageOutput, error) {
    var output marketplacemetering.RegisterUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceMeteringStub) ResolveCustomer(ctx workflow.Context, input *marketplacemetering.ResolveCustomerInput) (*marketplacemetering.ResolveCustomerOutput, error) {
    var output marketplacemetering.ResolveCustomerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResolveCustomer, input).Get(ctx, &output)
    return &output, err
}
