
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"github.com/aws/aws-sdk-go/service/computeoptimizer/computeoptimizeriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ComputeOptimizerActivities struct {
    client computeoptimizeriface.ComputeOptimizerAPI
}

func NewComputeOptimizerActivities(session *session.Session, config ...*aws.Config) *ComputeOptimizerActivities {
    client := computeoptimizer.New(session, config...)
    return &ComputeOptimizerActivities{client: client}
}

func (a *ComputeOptimizerActivities) DescribeRecommendationExportJobs(ctx context.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
    return a.client.DescribeRecommendationExportJobsWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) ExportAutoScalingGroupRecommendations(ctx context.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error) {
    return a.client.ExportAutoScalingGroupRecommendationsWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) ExportEC2InstanceRecommendations(ctx context.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error) {
    return a.client.ExportEC2InstanceRecommendationsWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) GetAutoScalingGroupRecommendations(ctx context.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
    return a.client.GetAutoScalingGroupRecommendationsWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) GetEC2InstanceRecommendations(ctx context.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
    return a.client.GetEC2InstanceRecommendationsWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) GetEC2RecommendationProjectedMetrics(ctx context.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
    return a.client.GetEC2RecommendationProjectedMetricsWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) GetEnrollmentStatus(ctx context.Context, input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
    return a.client.GetEnrollmentStatusWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) GetRecommendationSummaries(ctx context.Context, input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
    return a.client.GetRecommendationSummariesWithContext(ctx, input)
}

func (a *ComputeOptimizerActivities) UpdateEnrollmentStatus(ctx context.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
    return a.client.UpdateEnrollmentStatusWithContext(ctx, input)
}
