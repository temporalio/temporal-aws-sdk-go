package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"github.com/aws/aws-sdk-go/service/computeoptimizer/computeoptimizeriface"
)

type ComputeOptimizerActivities struct {
	client computeoptimizeriface.ComputeOptimizerAPI
}

func NewComputeOptimizerActivities(session *session.Session, config ...*aws.Config) *ComputeOptimizerActivities {
	client := computeoptimizer.New(session, config...)
	return &ComputeOptimizerActivities{client: client}
}

func (a *ComputeOptimizerActivities) DescribeRecommendationExportJobs(input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
	return a.client.DescribeRecommendationExportJobs(input)
}

func (a *ComputeOptimizerActivities) ExportAutoScalingGroupRecommendations(input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error) {
	return a.client.ExportAutoScalingGroupRecommendations(input)
}

func (a *ComputeOptimizerActivities) ExportEC2InstanceRecommendations(input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error) {
	return a.client.ExportEC2InstanceRecommendations(input)
}

func (a *ComputeOptimizerActivities) GetAutoScalingGroupRecommendations(input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	return a.client.GetAutoScalingGroupRecommendations(input)
}

func (a *ComputeOptimizerActivities) GetEC2InstanceRecommendations(input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	return a.client.GetEC2InstanceRecommendations(input)
}

func (a *ComputeOptimizerActivities) GetEC2RecommendationProjectedMetrics(input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	return a.client.GetEC2RecommendationProjectedMetrics(input)
}

func (a *ComputeOptimizerActivities) GetEnrollmentStatus(input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	return a.client.GetEnrollmentStatus(input)
}

func (a *ComputeOptimizerActivities) GetRecommendationSummaries(input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	return a.client.GetRecommendationSummaries(input)
}

func (a *ComputeOptimizerActivities) UpdateEnrollmentStatus(input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
	return a.client.UpdateEnrollmentStatus(input)
}
