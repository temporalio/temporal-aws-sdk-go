package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"go.temporal.io/sdk/workflow"
)

type AutoScalingPlansClient interface {
    CreateScalingPlan(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error)
    CreateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) *AutoscalingplansCreateScalingPlanResult

    DeleteScalingPlan(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error)
    DeleteScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) *AutoscalingplansDeleteScalingPlanResult

    DescribeScalingPlanResources(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error)
    DescribeScalingPlanResourcesAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) *AutoscalingplansDescribeScalingPlanResourcesResult

    DescribeScalingPlans(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error)
    DescribeScalingPlansAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) *AutoscalingplansDescribeScalingPlansResult

    GetScalingPlanResourceForecastData(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error)
    GetScalingPlanResourceForecastDataAsync(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) *AutoscalingplansGetScalingPlanResourceForecastDataResult

    UpdateScalingPlan(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error)
    UpdateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) *AutoscalingplansUpdateScalingPlanResult
}
type AutoscalingplansCreateScalingPlanResult struct {
	Result workflow.Future
}

func (r *AutoscalingplansCreateScalingPlanResult) Get(ctx workflow.Context) (*autoscalingplans.CreateScalingPlanOutput, error) {
    var output autoscalingplans.CreateScalingPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingplansDeleteScalingPlanResult struct {
	Result workflow.Future
}

func (r *AutoscalingplansDeleteScalingPlanResult) Get(ctx workflow.Context) (*autoscalingplans.DeleteScalingPlanOutput, error) {
    var output autoscalingplans.DeleteScalingPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingplansDescribeScalingPlanResourcesResult struct {
	Result workflow.Future
}

func (r *AutoscalingplansDescribeScalingPlanResourcesResult) Get(ctx workflow.Context) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
    var output autoscalingplans.DescribeScalingPlanResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingplansDescribeScalingPlansResult struct {
	Result workflow.Future
}

func (r *AutoscalingplansDescribeScalingPlansResult) Get(ctx workflow.Context) (*autoscalingplans.DescribeScalingPlansOutput, error) {
    var output autoscalingplans.DescribeScalingPlansOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingplansGetScalingPlanResourceForecastDataResult struct {
	Result workflow.Future
}

func (r *AutoscalingplansGetScalingPlanResourceForecastDataResult) Get(ctx workflow.Context) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
    var output autoscalingplans.GetScalingPlanResourceForecastDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AutoscalingplansUpdateScalingPlanResult struct {
	Result workflow.Future
}

func (r *AutoscalingplansUpdateScalingPlanResult) Get(ctx workflow.Context) (*autoscalingplans.UpdateScalingPlanOutput, error) {
    var output autoscalingplans.UpdateScalingPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type AutoScalingPlansStub struct {
    activities AutoScalingPlansClient
}

func NewAutoScalingPlansStub() AutoScalingPlansClient {
    return &AutoScalingPlansStub{}
}

func (a *AutoScalingPlansStub) CreateScalingPlan(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
    var output autoscalingplans.CreateScalingPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateScalingPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingPlansStub) DeleteScalingPlan(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error) {
    var output autoscalingplans.DeleteScalingPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteScalingPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingPlansStub) DescribeScalingPlanResources(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
    var output autoscalingplans.DescribeScalingPlanResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingPlanResources, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingPlansStub) DescribeScalingPlans(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
    var output autoscalingplans.DescribeScalingPlansOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingPlans, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingPlansStub) GetScalingPlanResourceForecastData(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
    var output autoscalingplans.GetScalingPlanResourceForecastDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetScalingPlanResourceForecastData, input).Get(ctx, &output)
    return &output, err
}

func (a *AutoScalingPlansStub) UpdateScalingPlan(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error) {
    var output autoscalingplans.UpdateScalingPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateScalingPlan, input).Get(ctx, &output)
    return &output, err
}
