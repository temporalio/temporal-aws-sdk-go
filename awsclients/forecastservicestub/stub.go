// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package forecastservicestub

import (
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type ForecastServiceCreateDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceCreateDatasetFuture) Get(ctx workflow.Context) (*forecastservice.CreateDatasetOutput, error) {
	var output forecastservice.CreateDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceCreateDatasetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceCreateDatasetGroupFuture) Get(ctx workflow.Context) (*forecastservice.CreateDatasetGroupOutput, error) {
	var output forecastservice.CreateDatasetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceCreateDatasetImportJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceCreateDatasetImportJobFuture) Get(ctx workflow.Context) (*forecastservice.CreateDatasetImportJobOutput, error) {
	var output forecastservice.CreateDatasetImportJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceCreateForecastFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceCreateForecastFuture) Get(ctx workflow.Context) (*forecastservice.CreateForecastOutput, error) {
	var output forecastservice.CreateForecastOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceCreateForecastExportJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceCreateForecastExportJobFuture) Get(ctx workflow.Context) (*forecastservice.CreateForecastExportJobOutput, error) {
	var output forecastservice.CreateForecastExportJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceCreatePredictorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceCreatePredictorFuture) Get(ctx workflow.Context) (*forecastservice.CreatePredictorOutput, error) {
	var output forecastservice.CreatePredictorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDeleteDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDeleteDatasetFuture) Get(ctx workflow.Context) (*forecastservice.DeleteDatasetOutput, error) {
	var output forecastservice.DeleteDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDeleteDatasetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDeleteDatasetGroupFuture) Get(ctx workflow.Context) (*forecastservice.DeleteDatasetGroupOutput, error) {
	var output forecastservice.DeleteDatasetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDeleteDatasetImportJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDeleteDatasetImportJobFuture) Get(ctx workflow.Context) (*forecastservice.DeleteDatasetImportJobOutput, error) {
	var output forecastservice.DeleteDatasetImportJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDeleteForecastFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDeleteForecastFuture) Get(ctx workflow.Context) (*forecastservice.DeleteForecastOutput, error) {
	var output forecastservice.DeleteForecastOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDeleteForecastExportJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDeleteForecastExportJobFuture) Get(ctx workflow.Context) (*forecastservice.DeleteForecastExportJobOutput, error) {
	var output forecastservice.DeleteForecastExportJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDeletePredictorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDeletePredictorFuture) Get(ctx workflow.Context) (*forecastservice.DeletePredictorOutput, error) {
	var output forecastservice.DeletePredictorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDescribeDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDescribeDatasetFuture) Get(ctx workflow.Context) (*forecastservice.DescribeDatasetOutput, error) {
	var output forecastservice.DescribeDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDescribeDatasetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDescribeDatasetGroupFuture) Get(ctx workflow.Context) (*forecastservice.DescribeDatasetGroupOutput, error) {
	var output forecastservice.DescribeDatasetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDescribeDatasetImportJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDescribeDatasetImportJobFuture) Get(ctx workflow.Context) (*forecastservice.DescribeDatasetImportJobOutput, error) {
	var output forecastservice.DescribeDatasetImportJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDescribeForecastFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDescribeForecastFuture) Get(ctx workflow.Context) (*forecastservice.DescribeForecastOutput, error) {
	var output forecastservice.DescribeForecastOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDescribeForecastExportJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDescribeForecastExportJobFuture) Get(ctx workflow.Context) (*forecastservice.DescribeForecastExportJobOutput, error) {
	var output forecastservice.DescribeForecastExportJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceDescribePredictorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceDescribePredictorFuture) Get(ctx workflow.Context) (*forecastservice.DescribePredictorOutput, error) {
	var output forecastservice.DescribePredictorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceGetAccuracyMetricsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceGetAccuracyMetricsFuture) Get(ctx workflow.Context) (*forecastservice.GetAccuracyMetricsOutput, error) {
	var output forecastservice.GetAccuracyMetricsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceListDatasetGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceListDatasetGroupsFuture) Get(ctx workflow.Context) (*forecastservice.ListDatasetGroupsOutput, error) {
	var output forecastservice.ListDatasetGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceListDatasetImportJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceListDatasetImportJobsFuture) Get(ctx workflow.Context) (*forecastservice.ListDatasetImportJobsOutput, error) {
	var output forecastservice.ListDatasetImportJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceListDatasetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceListDatasetsFuture) Get(ctx workflow.Context) (*forecastservice.ListDatasetsOutput, error) {
	var output forecastservice.ListDatasetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceListForecastExportJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceListForecastExportJobsFuture) Get(ctx workflow.Context) (*forecastservice.ListForecastExportJobsOutput, error) {
	var output forecastservice.ListForecastExportJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceListForecastsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceListForecastsFuture) Get(ctx workflow.Context) (*forecastservice.ListForecastsOutput, error) {
	var output forecastservice.ListForecastsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceListPredictorsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceListPredictorsFuture) Get(ctx workflow.Context) (*forecastservice.ListPredictorsOutput, error) {
	var output forecastservice.ListPredictorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceListTagsForResourceFuture) Get(ctx workflow.Context) (*forecastservice.ListTagsForResourceOutput, error) {
	var output forecastservice.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceTagResourceFuture) Get(ctx workflow.Context) (*forecastservice.TagResourceOutput, error) {
	var output forecastservice.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceUntagResourceFuture) Get(ctx workflow.Context) (*forecastservice.UntagResourceOutput, error) {
	var output forecastservice.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ForecastServiceUpdateDatasetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ForecastServiceUpdateDatasetGroupFuture) Get(ctx workflow.Context) (*forecastservice.UpdateDatasetGroupOutput, error) {
	var output forecastservice.UpdateDatasetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataset(ctx workflow.Context, input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error) {
	var output forecastservice.CreateDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatasetAsync(ctx workflow.Context, input *forecastservice.CreateDatasetInput) *ForecastServiceCreateDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateDataset", input)
	return &ForecastServiceCreateDatasetFuture{Future: future}
}

func (a *stub) CreateDatasetGroup(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error) {
	var output forecastservice.CreateDatasetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateDatasetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.CreateDatasetGroupInput) *ForecastServiceCreateDatasetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateDatasetGroup", input)
	return &ForecastServiceCreateDatasetGroupFuture{Future: future}
}

func (a *stub) CreateDatasetImportJob(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error) {
	var output forecastservice.CreateDatasetImportJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateDatasetImportJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.CreateDatasetImportJobInput) *ForecastServiceCreateDatasetImportJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateDatasetImportJob", input)
	return &ForecastServiceCreateDatasetImportJobFuture{Future: future}
}

func (a *stub) CreateForecast(ctx workflow.Context, input *forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error) {
	var output forecastservice.CreateForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateForecastAsync(ctx workflow.Context, input *forecastservice.CreateForecastInput) *ForecastServiceCreateForecastFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateForecast", input)
	return &ForecastServiceCreateForecastFuture{Future: future}
}

func (a *stub) CreateForecastExportJob(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error) {
	var output forecastservice.CreateForecastExportJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateForecastExportJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateForecastExportJobAsync(ctx workflow.Context, input *forecastservice.CreateForecastExportJobInput) *ForecastServiceCreateForecastExportJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreateForecastExportJob", input)
	return &ForecastServiceCreateForecastExportJobFuture{Future: future}
}

func (a *stub) CreatePredictor(ctx workflow.Context, input *forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error) {
	var output forecastservice.CreatePredictorOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreatePredictor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePredictorAsync(ctx workflow.Context, input *forecastservice.CreatePredictorInput) *ForecastServiceCreatePredictorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.CreatePredictor", input)
	return &ForecastServiceCreatePredictorFuture{Future: future}
}

func (a *stub) DeleteDataset(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error) {
	var output forecastservice.DeleteDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatasetAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetInput) *ForecastServiceDeleteDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteDataset", input)
	return &ForecastServiceDeleteDatasetFuture{Future: future}
}

func (a *stub) DeleteDatasetGroup(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error) {
	var output forecastservice.DeleteDatasetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteDatasetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetGroupInput) *ForecastServiceDeleteDatasetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteDatasetGroup", input)
	return &ForecastServiceDeleteDatasetGroupFuture{Future: future}
}

func (a *stub) DeleteDatasetImportJob(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error) {
	var output forecastservice.DeleteDatasetImportJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteDatasetImportJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DeleteDatasetImportJobInput) *ForecastServiceDeleteDatasetImportJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteDatasetImportJob", input)
	return &ForecastServiceDeleteDatasetImportJobFuture{Future: future}
}

func (a *stub) DeleteForecast(ctx workflow.Context, input *forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error) {
	var output forecastservice.DeleteForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteForecastAsync(ctx workflow.Context, input *forecastservice.DeleteForecastInput) *ForecastServiceDeleteForecastFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteForecast", input)
	return &ForecastServiceDeleteForecastFuture{Future: future}
}

func (a *stub) DeleteForecastExportJob(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error) {
	var output forecastservice.DeleteForecastExportJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteForecastExportJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DeleteForecastExportJobInput) *ForecastServiceDeleteForecastExportJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeleteForecastExportJob", input)
	return &ForecastServiceDeleteForecastExportJobFuture{Future: future}
}

func (a *stub) DeletePredictor(ctx workflow.Context, input *forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error) {
	var output forecastservice.DeletePredictorOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeletePredictor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePredictorAsync(ctx workflow.Context, input *forecastservice.DeletePredictorInput) *ForecastServiceDeletePredictorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DeletePredictor", input)
	return &ForecastServiceDeletePredictorFuture{Future: future}
}

func (a *stub) DescribeDataset(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error) {
	var output forecastservice.DescribeDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatasetAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetInput) *ForecastServiceDescribeDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeDataset", input)
	return &ForecastServiceDescribeDatasetFuture{Future: future}
}

func (a *stub) DescribeDatasetGroup(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error) {
	var output forecastservice.DescribeDatasetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeDatasetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatasetGroupAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetGroupInput) *ForecastServiceDescribeDatasetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeDatasetGroup", input)
	return &ForecastServiceDescribeDatasetGroupFuture{Future: future}
}

func (a *stub) DescribeDatasetImportJob(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error) {
	var output forecastservice.DescribeDatasetImportJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeDatasetImportJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatasetImportJobAsync(ctx workflow.Context, input *forecastservice.DescribeDatasetImportJobInput) *ForecastServiceDescribeDatasetImportJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeDatasetImportJob", input)
	return &ForecastServiceDescribeDatasetImportJobFuture{Future: future}
}

func (a *stub) DescribeForecast(ctx workflow.Context, input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error) {
	var output forecastservice.DescribeForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeForecastAsync(ctx workflow.Context, input *forecastservice.DescribeForecastInput) *ForecastServiceDescribeForecastFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeForecast", input)
	return &ForecastServiceDescribeForecastFuture{Future: future}
}

func (a *stub) DescribeForecastExportJob(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error) {
	var output forecastservice.DescribeForecastExportJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeForecastExportJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeForecastExportJobAsync(ctx workflow.Context, input *forecastservice.DescribeForecastExportJobInput) *ForecastServiceDescribeForecastExportJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribeForecastExportJob", input)
	return &ForecastServiceDescribeForecastExportJobFuture{Future: future}
}

func (a *stub) DescribePredictor(ctx workflow.Context, input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error) {
	var output forecastservice.DescribePredictorOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribePredictor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePredictorAsync(ctx workflow.Context, input *forecastservice.DescribePredictorInput) *ForecastServiceDescribePredictorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.DescribePredictor", input)
	return &ForecastServiceDescribePredictorFuture{Future: future}
}

func (a *stub) GetAccuracyMetrics(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error) {
	var output forecastservice.GetAccuracyMetricsOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.GetAccuracyMetrics", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAccuracyMetricsAsync(ctx workflow.Context, input *forecastservice.GetAccuracyMetricsInput) *ForecastServiceGetAccuracyMetricsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.GetAccuracyMetrics", input)
	return &ForecastServiceGetAccuracyMetricsFuture{Future: future}
}

func (a *stub) ListDatasetGroups(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error) {
	var output forecastservice.ListDatasetGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListDatasetGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatasetGroupsAsync(ctx workflow.Context, input *forecastservice.ListDatasetGroupsInput) *ForecastServiceListDatasetGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListDatasetGroups", input)
	return &ForecastServiceListDatasetGroupsFuture{Future: future}
}

func (a *stub) ListDatasetImportJobs(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error) {
	var output forecastservice.ListDatasetImportJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListDatasetImportJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatasetImportJobsAsync(ctx workflow.Context, input *forecastservice.ListDatasetImportJobsInput) *ForecastServiceListDatasetImportJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListDatasetImportJobs", input)
	return &ForecastServiceListDatasetImportJobsFuture{Future: future}
}

func (a *stub) ListDatasets(ctx workflow.Context, input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error) {
	var output forecastservice.ListDatasetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListDatasets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatasetsAsync(ctx workflow.Context, input *forecastservice.ListDatasetsInput) *ForecastServiceListDatasetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListDatasets", input)
	return &ForecastServiceListDatasetsFuture{Future: future}
}

func (a *stub) ListForecastExportJobs(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error) {
	var output forecastservice.ListForecastExportJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListForecastExportJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListForecastExportJobsAsync(ctx workflow.Context, input *forecastservice.ListForecastExportJobsInput) *ForecastServiceListForecastExportJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListForecastExportJobs", input)
	return &ForecastServiceListForecastExportJobsFuture{Future: future}
}

func (a *stub) ListForecasts(ctx workflow.Context, input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error) {
	var output forecastservice.ListForecastsOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListForecasts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListForecastsAsync(ctx workflow.Context, input *forecastservice.ListForecastsInput) *ForecastServiceListForecastsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListForecasts", input)
	return &ForecastServiceListForecastsFuture{Future: future}
}

func (a *stub) ListPredictors(ctx workflow.Context, input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error) {
	var output forecastservice.ListPredictorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListPredictors", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPredictorsAsync(ctx workflow.Context, input *forecastservice.ListPredictorsInput) *ForecastServiceListPredictorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListPredictors", input)
	return &ForecastServiceListPredictorsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error) {
	var output forecastservice.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *forecastservice.ListTagsForResourceInput) *ForecastServiceListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.ListTagsForResource", input)
	return &ForecastServiceListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error) {
	var output forecastservice.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *forecastservice.TagResourceInput) *ForecastServiceTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.TagResource", input)
	return &ForecastServiceTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error) {
	var output forecastservice.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *forecastservice.UntagResourceInput) *ForecastServiceUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.UntagResource", input)
	return &ForecastServiceUntagResourceFuture{Future: future}
}

func (a *stub) UpdateDatasetGroup(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error) {
	var output forecastservice.UpdateDatasetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.forecastservice.UpdateDatasetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDatasetGroupAsync(ctx workflow.Context, input *forecastservice.UpdateDatasetGroupInput) *ForecastServiceUpdateDatasetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.forecastservice.UpdateDatasetGroup", input)
	return &ForecastServiceUpdateDatasetGroupFuture{Future: future}
}
