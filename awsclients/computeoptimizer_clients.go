package awsclients

import (
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"go.temporal.io/sdk/workflow"
)

type ComputeOptimizerClient interface {
       DescribeRecommendationExportJobs(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
       DescribeRecommendationExportJobsAsync(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) *ComputeoptimizerDescribeRecommendationExportJobsResult

       ExportAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
       ExportAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) *ComputeoptimizerExportAutoScalingGroupRecommendationsResult

       ExportEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
       ExportEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) *ComputeoptimizerExportEC2InstanceRecommendationsResult

       GetAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
       GetAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) *ComputeoptimizerGetAutoScalingGroupRecommendationsResult

       GetEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
       GetEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) *ComputeoptimizerGetEC2InstanceRecommendationsResult

       GetEC2RecommendationProjectedMetrics(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
       GetEC2RecommendationProjectedMetricsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) *ComputeoptimizerGetEC2RecommendationProjectedMetricsResult

       GetEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error)
       GetEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) *ComputeoptimizerGetEnrollmentStatusResult

       GetRecommendationSummaries(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error)
       GetRecommendationSummariesAsync(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) *ComputeoptimizerGetRecommendationSummariesResult

       UpdateEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
       UpdateEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) *ComputeoptimizerUpdateEnrollmentStatusResult
}

type ComputeOptimizerStub struct {
}

func NewComputeOptimizerStub() ComputeOptimizerClient {
    return &ComputeOptimizerStub{}
}

type ComputeoptimizerDescribeRecommendationExportJobsResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerDescribeRecommendationExportJobsResult) Get(ctx workflow.Context) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
    var output computeoptimizer.DescribeRecommendationExportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerExportAutoScalingGroupRecommendationsResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerExportAutoScalingGroupRecommendationsResult) Get(ctx workflow.Context) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error) {
    var output computeoptimizer.ExportAutoScalingGroupRecommendationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerExportEC2InstanceRecommendationsResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerExportEC2InstanceRecommendationsResult) Get(ctx workflow.Context) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error) {
    var output computeoptimizer.ExportEC2InstanceRecommendationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerGetAutoScalingGroupRecommendationsResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerGetAutoScalingGroupRecommendationsResult) Get(ctx workflow.Context) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
    var output computeoptimizer.GetAutoScalingGroupRecommendationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerGetEC2InstanceRecommendationsResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerGetEC2InstanceRecommendationsResult) Get(ctx workflow.Context) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
    var output computeoptimizer.GetEC2InstanceRecommendationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerGetEC2RecommendationProjectedMetricsResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerGetEC2RecommendationProjectedMetricsResult) Get(ctx workflow.Context) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
    var output computeoptimizer.GetEC2RecommendationProjectedMetricsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerGetEnrollmentStatusResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerGetEnrollmentStatusResult) Get(ctx workflow.Context) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
    var output computeoptimizer.GetEnrollmentStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerGetRecommendationSummariesResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerGetRecommendationSummariesResult) Get(ctx workflow.Context) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
    var output computeoptimizer.GetRecommendationSummariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ComputeoptimizerUpdateEnrollmentStatusResult struct {
	Result workflow.Future
}

func (r *ComputeoptimizerUpdateEnrollmentStatusResult) Get(ctx workflow.Context) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
    var output computeoptimizer.UpdateEnrollmentStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) DescribeRecommendationExportJobs(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
    var output computeoptimizer.DescribeRecommendationExportJobsOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.DescribeRecommendationExportJobs", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) DescribeRecommendationExportJobsAsync(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) *ComputeoptimizerDescribeRecommendationExportJobsResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.DescribeRecommendationExportJobs", input)
    return &ComputeoptimizerDescribeRecommendationExportJobsResult{Result: future}
}

func (a *ComputeOptimizerStub) ExportAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error) {
    var output computeoptimizer.ExportAutoScalingGroupRecommendationsOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.ExportAutoScalingGroupRecommendations", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) ExportAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) *ComputeoptimizerExportAutoScalingGroupRecommendationsResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.ExportAutoScalingGroupRecommendations", input)
    return &ComputeoptimizerExportAutoScalingGroupRecommendationsResult{Result: future}
}

func (a *ComputeOptimizerStub) ExportEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error) {
    var output computeoptimizer.ExportEC2InstanceRecommendationsOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.ExportEC2InstanceRecommendations", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) ExportEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) *ComputeoptimizerExportEC2InstanceRecommendationsResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.ExportEC2InstanceRecommendations", input)
    return &ComputeoptimizerExportEC2InstanceRecommendationsResult{Result: future}
}

func (a *ComputeOptimizerStub) GetAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
    var output computeoptimizer.GetAutoScalingGroupRecommendationsOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetAutoScalingGroupRecommendations", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) GetAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) *ComputeoptimizerGetAutoScalingGroupRecommendationsResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetAutoScalingGroupRecommendations", input)
    return &ComputeoptimizerGetAutoScalingGroupRecommendationsResult{Result: future}
}

func (a *ComputeOptimizerStub) GetEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
    var output computeoptimizer.GetEC2InstanceRecommendationsOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetEC2InstanceRecommendations", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) GetEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) *ComputeoptimizerGetEC2InstanceRecommendationsResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetEC2InstanceRecommendations", input)
    return &ComputeoptimizerGetEC2InstanceRecommendationsResult{Result: future}
}

func (a *ComputeOptimizerStub) GetEC2RecommendationProjectedMetrics(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
    var output computeoptimizer.GetEC2RecommendationProjectedMetricsOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetEC2RecommendationProjectedMetrics", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) GetEC2RecommendationProjectedMetricsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) *ComputeoptimizerGetEC2RecommendationProjectedMetricsResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetEC2RecommendationProjectedMetrics", input)
    return &ComputeoptimizerGetEC2RecommendationProjectedMetricsResult{Result: future}
}

func (a *ComputeOptimizerStub) GetEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
    var output computeoptimizer.GetEnrollmentStatusOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetEnrollmentStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) GetEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) *ComputeoptimizerGetEnrollmentStatusResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetEnrollmentStatus", input)
    return &ComputeoptimizerGetEnrollmentStatusResult{Result: future}
}

func (a *ComputeOptimizerStub) GetRecommendationSummaries(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
    var output computeoptimizer.GetRecommendationSummariesOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetRecommendationSummaries", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) GetRecommendationSummariesAsync(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) *ComputeoptimizerGetRecommendationSummariesResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.GetRecommendationSummaries", input)
    return &ComputeoptimizerGetRecommendationSummariesResult{Result: future}
}

func (a *ComputeOptimizerStub) UpdateEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
    var output computeoptimizer.UpdateEnrollmentStatusOutput
    err := workflow.ExecuteActivity(ctx, "ComputeOptimizer.UpdateEnrollmentStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *ComputeOptimizerStub) UpdateEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) *ComputeoptimizerUpdateEnrollmentStatusResult {
    future := workflow.ExecuteActivity(ctx, "ComputeOptimizer.UpdateEnrollmentStatus", input)
    return &ComputeoptimizerUpdateEnrollmentStatusResult{Result: future}
}
