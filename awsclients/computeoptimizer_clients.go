// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"go.temporal.io/sdk/workflow"
)

type ComputeOptimizerClient interface {
	DescribeRecommendationExportJobs(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
	DescribeRecommendationExportJobsAsync(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) *ComputeOptimizerDescribeRecommendationExportJobsFuture

	ExportAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
	ExportAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) *ComputeOptimizerExportAutoScalingGroupRecommendationsFuture

	ExportEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
	ExportEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) *ComputeOptimizerExportEC2InstanceRecommendationsFuture

	GetAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
	GetAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) *ComputeOptimizerGetAutoScalingGroupRecommendationsFuture

	GetEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
	GetEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) *ComputeOptimizerGetEC2InstanceRecommendationsFuture

	GetEC2RecommendationProjectedMetrics(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
	GetEC2RecommendationProjectedMetricsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) *ComputeOptimizerGetEC2RecommendationProjectedMetricsFuture

	GetEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error)
	GetEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) *ComputeOptimizerGetEnrollmentStatusFuture

	GetRecommendationSummaries(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error)
	GetRecommendationSummariesAsync(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) *ComputeOptimizerGetRecommendationSummariesFuture

	UpdateEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
	UpdateEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) *ComputeOptimizerUpdateEnrollmentStatusFuture
}

type ComputeOptimizerStub struct{}

func NewComputeOptimizerStub() ComputeOptimizerClient {
	return &ComputeOptimizerStub{}
}

type ComputeOptimizerDescribeRecommendationExportJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerDescribeRecommendationExportJobsFuture) Get(ctx workflow.Context) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
	var output computeoptimizer.DescribeRecommendationExportJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerExportAutoScalingGroupRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerExportAutoScalingGroupRecommendationsFuture) Get(ctx workflow.Context) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error) {
	var output computeoptimizer.ExportAutoScalingGroupRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerExportEC2InstanceRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerExportEC2InstanceRecommendationsFuture) Get(ctx workflow.Context) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error) {
	var output computeoptimizer.ExportEC2InstanceRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerGetAutoScalingGroupRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerGetAutoScalingGroupRecommendationsFuture) Get(ctx workflow.Context) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	var output computeoptimizer.GetAutoScalingGroupRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerGetEC2InstanceRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerGetEC2InstanceRecommendationsFuture) Get(ctx workflow.Context) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	var output computeoptimizer.GetEC2InstanceRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerGetEC2RecommendationProjectedMetricsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerGetEC2RecommendationProjectedMetricsFuture) Get(ctx workflow.Context) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	var output computeoptimizer.GetEC2RecommendationProjectedMetricsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerGetEnrollmentStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerGetEnrollmentStatusFuture) Get(ctx workflow.Context) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	var output computeoptimizer.GetEnrollmentStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerGetRecommendationSummariesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerGetRecommendationSummariesFuture) Get(ctx workflow.Context) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	var output computeoptimizer.GetRecommendationSummariesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComputeOptimizerUpdateEnrollmentStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ComputeOptimizerUpdateEnrollmentStatusFuture) Get(ctx workflow.Context) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
	var output computeoptimizer.UpdateEnrollmentStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) DescribeRecommendationExportJobs(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
	var output computeoptimizer.DescribeRecommendationExportJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.DescribeRecommendationExportJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) DescribeRecommendationExportJobsAsync(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) *ComputeOptimizerDescribeRecommendationExportJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.DescribeRecommendationExportJobs", input)
	return &ComputeOptimizerDescribeRecommendationExportJobsFuture{Future: future}
}

func (a *ComputeOptimizerStub) ExportAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error) {
	var output computeoptimizer.ExportAutoScalingGroupRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.ExportAutoScalingGroupRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) ExportAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) *ComputeOptimizerExportAutoScalingGroupRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.ExportAutoScalingGroupRecommendations", input)
	return &ComputeOptimizerExportAutoScalingGroupRecommendationsFuture{Future: future}
}

func (a *ComputeOptimizerStub) ExportEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error) {
	var output computeoptimizer.ExportEC2InstanceRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.ExportEC2InstanceRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) ExportEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) *ComputeOptimizerExportEC2InstanceRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.ExportEC2InstanceRecommendations", input)
	return &ComputeOptimizerExportEC2InstanceRecommendationsFuture{Future: future}
}

func (a *ComputeOptimizerStub) GetAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	var output computeoptimizer.GetAutoScalingGroupRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetAutoScalingGroupRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) GetAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) *ComputeOptimizerGetAutoScalingGroupRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetAutoScalingGroupRecommendations", input)
	return &ComputeOptimizerGetAutoScalingGroupRecommendationsFuture{Future: future}
}

func (a *ComputeOptimizerStub) GetEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	var output computeoptimizer.GetEC2InstanceRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetEC2InstanceRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) GetEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) *ComputeOptimizerGetEC2InstanceRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetEC2InstanceRecommendations", input)
	return &ComputeOptimizerGetEC2InstanceRecommendationsFuture{Future: future}
}

func (a *ComputeOptimizerStub) GetEC2RecommendationProjectedMetrics(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	var output computeoptimizer.GetEC2RecommendationProjectedMetricsOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetEC2RecommendationProjectedMetrics", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) GetEC2RecommendationProjectedMetricsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) *ComputeOptimizerGetEC2RecommendationProjectedMetricsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetEC2RecommendationProjectedMetrics", input)
	return &ComputeOptimizerGetEC2RecommendationProjectedMetricsFuture{Future: future}
}

func (a *ComputeOptimizerStub) GetEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	var output computeoptimizer.GetEnrollmentStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetEnrollmentStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) GetEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) *ComputeOptimizerGetEnrollmentStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetEnrollmentStatus", input)
	return &ComputeOptimizerGetEnrollmentStatusFuture{Future: future}
}

func (a *ComputeOptimizerStub) GetRecommendationSummaries(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	var output computeoptimizer.GetRecommendationSummariesOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetRecommendationSummaries", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) GetRecommendationSummariesAsync(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) *ComputeOptimizerGetRecommendationSummariesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.GetRecommendationSummaries", input)
	return &ComputeOptimizerGetRecommendationSummariesFuture{Future: future}
}

func (a *ComputeOptimizerStub) UpdateEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error) {
	var output computeoptimizer.UpdateEnrollmentStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.UpdateEnrollmentStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *ComputeOptimizerStub) UpdateEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) *ComputeOptimizerUpdateEnrollmentStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.computeoptimizer.UpdateEnrollmentStatus", input)
	return &ComputeOptimizerUpdateEnrollmentStatusFuture{Future: future}
}
