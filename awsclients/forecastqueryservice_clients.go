package awsclients

import (
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ForecastQueryServiceClient interface {
       QueryForecast(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error)
       QueryForecastAsync(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) *ForecastqueryserviceQueryForecastResult
}

type ForecastqueryserviceQueryForecastResult struct {
	Result workflow.Future
}

func (r *ForecastqueryserviceQueryForecastResult) Get(ctx workflow.Context) (*forecastqueryservice.QueryForecastOutput, error) {
    var output forecastqueryservice.QueryForecastOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ForecastQueryServiceStub struct {
    activities awsactivities.ForecastQueryServiceActivities
}

func NewForecastQueryServiceStub() ForecastQueryServiceClient {
    return &ForecastQueryServiceStub{}
}

func (a *ForecastQueryServiceStub) QueryForecast(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
    var output forecastqueryservice.QueryForecastOutput
    err := workflow.ExecuteActivity(ctx, a.activities.QueryForecast, input).Get(ctx, &output)
    return &output, err
}

func (a *ForecastQueryServiceStub) QueryForecastAsync(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) *ForecastqueryserviceQueryForecastResult {
    future := workflow.ExecuteActivity(ctx, a.activities.QueryForecast, input)
    return &ForecastqueryserviceQueryForecastResult{Result: future}
}
