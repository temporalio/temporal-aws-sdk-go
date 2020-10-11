// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"go.temporal.io/sdk/workflow"
)

type CostExplorerClient interface {
	CreateAnomalyMonitor(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error)
	CreateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) *CostexplorerCreateAnomalyMonitorResult

	CreateAnomalySubscription(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error)
	CreateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) *CostexplorerCreateAnomalySubscriptionResult

	CreateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error)
	CreateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) *CostexplorerCreateCostCategoryDefinitionResult

	DeleteAnomalyMonitor(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error)
	DeleteAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) *CostexplorerDeleteAnomalyMonitorResult

	DeleteAnomalySubscription(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error)
	DeleteAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) *CostexplorerDeleteAnomalySubscriptionResult

	DeleteCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error)
	DeleteCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) *CostexplorerDeleteCostCategoryDefinitionResult

	DescribeCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	DescribeCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) *CostexplorerDescribeCostCategoryDefinitionResult

	GetAnomalies(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomaliesAsync(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) *CostexplorerGetAnomaliesResult

	GetAnomalyMonitors(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalyMonitorsAsync(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) *CostexplorerGetAnomalyMonitorsResult

	GetAnomalySubscriptions(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetAnomalySubscriptionsAsync(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) *CostexplorerGetAnomalySubscriptionsResult

	GetCostAndUsage(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) *CostexplorerGetCostAndUsageResult

	GetCostAndUsageWithResources(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostAndUsageWithResourcesAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) *CostexplorerGetCostAndUsageWithResourcesResult

	GetCostForecast(ctx workflow.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error)
	GetCostForecastAsync(ctx workflow.Context, input *costexplorer.GetCostForecastInput) *CostexplorerGetCostForecastResult

	GetDimensionValues(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error)
	GetDimensionValuesAsync(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) *CostexplorerGetDimensionValuesResult

	GetReservationCoverage(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationCoverageAsync(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) *CostexplorerGetReservationCoverageResult

	GetReservationPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) *CostexplorerGetReservationPurchaseRecommendationResult

	GetReservationUtilization(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error)
	GetReservationUtilizationAsync(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) *CostexplorerGetReservationUtilizationResult

	GetRightsizingRecommendation(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetRightsizingRecommendationAsync(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) *CostexplorerGetRightsizingRecommendationResult

	GetSavingsPlansCoverage(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansCoverageAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) *CostexplorerGetSavingsPlansCoverageResult

	GetSavingsPlansPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) *CostexplorerGetSavingsPlansPurchaseRecommendationResult

	GetSavingsPlansUtilization(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) *CostexplorerGetSavingsPlansUtilizationResult

	GetSavingsPlansUtilizationDetails(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetSavingsPlansUtilizationDetailsAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) *CostexplorerGetSavingsPlansUtilizationDetailsResult

	GetTags(ctx workflow.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *costexplorer.GetTagsInput) *CostexplorerGetTagsResult

	GetUsageForecast(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error)
	GetUsageForecastAsync(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) *CostexplorerGetUsageForecastResult

	ListCostCategoryDefinitions(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListCostCategoryDefinitionsAsync(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) *CostexplorerListCostCategoryDefinitionsResult

	ProvideAnomalyFeedback(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error)
	ProvideAnomalyFeedbackAsync(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) *CostexplorerProvideAnomalyFeedbackResult

	UpdateAnomalyMonitor(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error)
	UpdateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) *CostexplorerUpdateAnomalyMonitorResult

	UpdateAnomalySubscription(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error)
	UpdateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) *CostexplorerUpdateAnomalySubscriptionResult

	UpdateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error)
	UpdateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) *CostexplorerUpdateCostCategoryDefinitionResult
}

type CostExplorerStub struct{}

func NewCostExplorerStub() CostExplorerClient {
	return &CostExplorerStub{}
}

type CostexplorerCreateAnomalyMonitorResult struct {
	Result workflow.Future
}

func (r *CostexplorerCreateAnomalyMonitorResult) Get(ctx workflow.Context) (*costexplorer.CreateAnomalyMonitorOutput, error) {
	var output costexplorer.CreateAnomalyMonitorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerCreateAnomalySubscriptionResult struct {
	Result workflow.Future
}

func (r *CostexplorerCreateAnomalySubscriptionResult) Get(ctx workflow.Context) (*costexplorer.CreateAnomalySubscriptionOutput, error) {
	var output costexplorer.CreateAnomalySubscriptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerCreateCostCategoryDefinitionResult struct {
	Result workflow.Future
}

func (r *CostexplorerCreateCostCategoryDefinitionResult) Get(ctx workflow.Context) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
	var output costexplorer.CreateCostCategoryDefinitionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerDeleteAnomalyMonitorResult struct {
	Result workflow.Future
}

func (r *CostexplorerDeleteAnomalyMonitorResult) Get(ctx workflow.Context) (*costexplorer.DeleteAnomalyMonitorOutput, error) {
	var output costexplorer.DeleteAnomalyMonitorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerDeleteAnomalySubscriptionResult struct {
	Result workflow.Future
}

func (r *CostexplorerDeleteAnomalySubscriptionResult) Get(ctx workflow.Context) (*costexplorer.DeleteAnomalySubscriptionOutput, error) {
	var output costexplorer.DeleteAnomalySubscriptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerDeleteCostCategoryDefinitionResult struct {
	Result workflow.Future
}

func (r *CostexplorerDeleteCostCategoryDefinitionResult) Get(ctx workflow.Context) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
	var output costexplorer.DeleteCostCategoryDefinitionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerDescribeCostCategoryDefinitionResult struct {
	Result workflow.Future
}

func (r *CostexplorerDescribeCostCategoryDefinitionResult) Get(ctx workflow.Context) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	var output costexplorer.DescribeCostCategoryDefinitionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetAnomaliesResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetAnomaliesResult) Get(ctx workflow.Context) (*costexplorer.GetAnomaliesOutput, error) {
	var output costexplorer.GetAnomaliesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetAnomalyMonitorsResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetAnomalyMonitorsResult) Get(ctx workflow.Context) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	var output costexplorer.GetAnomalyMonitorsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetAnomalySubscriptionsResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetAnomalySubscriptionsResult) Get(ctx workflow.Context) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	var output costexplorer.GetAnomalySubscriptionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetCostAndUsageResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetCostAndUsageResult) Get(ctx workflow.Context) (*costexplorer.GetCostAndUsageOutput, error) {
	var output costexplorer.GetCostAndUsageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetCostAndUsageWithResourcesResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetCostAndUsageWithResourcesResult) Get(ctx workflow.Context) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	var output costexplorer.GetCostAndUsageWithResourcesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetCostForecastResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetCostForecastResult) Get(ctx workflow.Context) (*costexplorer.GetCostForecastOutput, error) {
	var output costexplorer.GetCostForecastOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetDimensionValuesResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetDimensionValuesResult) Get(ctx workflow.Context) (*costexplorer.GetDimensionValuesOutput, error) {
	var output costexplorer.GetDimensionValuesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetReservationCoverageResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetReservationCoverageResult) Get(ctx workflow.Context) (*costexplorer.GetReservationCoverageOutput, error) {
	var output costexplorer.GetReservationCoverageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetReservationPurchaseRecommendationResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetReservationPurchaseRecommendationResult) Get(ctx workflow.Context) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	var output costexplorer.GetReservationPurchaseRecommendationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetReservationUtilizationResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetReservationUtilizationResult) Get(ctx workflow.Context) (*costexplorer.GetReservationUtilizationOutput, error) {
	var output costexplorer.GetReservationUtilizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetRightsizingRecommendationResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetRightsizingRecommendationResult) Get(ctx workflow.Context) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	var output costexplorer.GetRightsizingRecommendationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetSavingsPlansCoverageResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetSavingsPlansCoverageResult) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	var output costexplorer.GetSavingsPlansCoverageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetSavingsPlansPurchaseRecommendationResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetSavingsPlansPurchaseRecommendationResult) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	var output costexplorer.GetSavingsPlansPurchaseRecommendationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetSavingsPlansUtilizationResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetSavingsPlansUtilizationResult) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetSavingsPlansUtilizationDetailsResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetSavingsPlansUtilizationDetailsResult) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationDetailsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetTagsResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetTagsResult) Get(ctx workflow.Context) (*costexplorer.GetTagsOutput, error) {
	var output costexplorer.GetTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerGetUsageForecastResult struct {
	Result workflow.Future
}

func (r *CostexplorerGetUsageForecastResult) Get(ctx workflow.Context) (*costexplorer.GetUsageForecastOutput, error) {
	var output costexplorer.GetUsageForecastOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerListCostCategoryDefinitionsResult struct {
	Result workflow.Future
}

func (r *CostexplorerListCostCategoryDefinitionsResult) Get(ctx workflow.Context) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	var output costexplorer.ListCostCategoryDefinitionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerProvideAnomalyFeedbackResult struct {
	Result workflow.Future
}

func (r *CostexplorerProvideAnomalyFeedbackResult) Get(ctx workflow.Context) (*costexplorer.ProvideAnomalyFeedbackOutput, error) {
	var output costexplorer.ProvideAnomalyFeedbackOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerUpdateAnomalyMonitorResult struct {
	Result workflow.Future
}

func (r *CostexplorerUpdateAnomalyMonitorResult) Get(ctx workflow.Context) (*costexplorer.UpdateAnomalyMonitorOutput, error) {
	var output costexplorer.UpdateAnomalyMonitorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerUpdateAnomalySubscriptionResult struct {
	Result workflow.Future
}

func (r *CostexplorerUpdateAnomalySubscriptionResult) Get(ctx workflow.Context) (*costexplorer.UpdateAnomalySubscriptionOutput, error) {
	var output costexplorer.UpdateAnomalySubscriptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CostexplorerUpdateCostCategoryDefinitionResult struct {
	Result workflow.Future
}

func (r *CostexplorerUpdateCostCategoryDefinitionResult) Get(ctx workflow.Context) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
	var output costexplorer.UpdateCostCategoryDefinitionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateAnomalyMonitor(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error) {
	var output costexplorer.CreateAnomalyMonitorOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalyMonitor", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) *CostexplorerCreateAnomalyMonitorResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalyMonitor", input)
	return &CostexplorerCreateAnomalyMonitorResult{Result: future}
}

func (a *CostExplorerStub) CreateAnomalySubscription(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error) {
	var output costexplorer.CreateAnomalySubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalySubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) *CostexplorerCreateAnomalySubscriptionResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalySubscription", input)
	return &CostexplorerCreateAnomalySubscriptionResult{Result: future}
}

func (a *CostExplorerStub) CreateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
	var output costexplorer.CreateCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) *CostexplorerCreateCostCategoryDefinitionResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateCostCategoryDefinition", input)
	return &CostexplorerCreateCostCategoryDefinitionResult{Result: future}
}

func (a *CostExplorerStub) DeleteAnomalyMonitor(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error) {
	var output costexplorer.DeleteAnomalyMonitorOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalyMonitor", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DeleteAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) *CostexplorerDeleteAnomalyMonitorResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalyMonitor", input)
	return &CostexplorerDeleteAnomalyMonitorResult{Result: future}
}

func (a *CostExplorerStub) DeleteAnomalySubscription(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error) {
	var output costexplorer.DeleteAnomalySubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalySubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DeleteAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) *CostexplorerDeleteAnomalySubscriptionResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalySubscription", input)
	return &CostexplorerDeleteAnomalySubscriptionResult{Result: future}
}

func (a *CostExplorerStub) DeleteCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
	var output costexplorer.DeleteCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DeleteCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) *CostexplorerDeleteCostCategoryDefinitionResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteCostCategoryDefinition", input)
	return &CostexplorerDeleteCostCategoryDefinitionResult{Result: future}
}

func (a *CostExplorerStub) DescribeCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	var output costexplorer.DescribeCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DescribeCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DescribeCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) *CostexplorerDescribeCostCategoryDefinitionResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DescribeCostCategoryDefinition", input)
	return &CostexplorerDescribeCostCategoryDefinitionResult{Result: future}
}

func (a *CostExplorerStub) GetAnomalies(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error) {
	var output costexplorer.GetAnomaliesOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalies", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetAnomaliesAsync(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) *CostexplorerGetAnomaliesResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalies", input)
	return &CostexplorerGetAnomaliesResult{Result: future}
}

func (a *CostExplorerStub) GetAnomalyMonitors(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	var output costexplorer.GetAnomalyMonitorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalyMonitors", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetAnomalyMonitorsAsync(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) *CostexplorerGetAnomalyMonitorsResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalyMonitors", input)
	return &CostexplorerGetAnomalyMonitorsResult{Result: future}
}

func (a *CostExplorerStub) GetAnomalySubscriptions(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	var output costexplorer.GetAnomalySubscriptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalySubscriptions", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetAnomalySubscriptionsAsync(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) *CostexplorerGetAnomalySubscriptionsResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalySubscriptions", input)
	return &CostexplorerGetAnomalySubscriptionsResult{Result: future}
}

func (a *CostExplorerStub) GetCostAndUsage(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
	var output costexplorer.GetCostAndUsageOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsage", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetCostAndUsageAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) *CostexplorerGetCostAndUsageResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsage", input)
	return &CostexplorerGetCostAndUsageResult{Result: future}
}

func (a *CostExplorerStub) GetCostAndUsageWithResources(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	var output costexplorer.GetCostAndUsageWithResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsageWithResources", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetCostAndUsageWithResourcesAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) *CostexplorerGetCostAndUsageWithResourcesResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsageWithResources", input)
	return &CostexplorerGetCostAndUsageWithResourcesResult{Result: future}
}

func (a *CostExplorerStub) GetCostForecast(ctx workflow.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error) {
	var output costexplorer.GetCostForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetCostForecastAsync(ctx workflow.Context, input *costexplorer.GetCostForecastInput) *CostexplorerGetCostForecastResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostForecast", input)
	return &CostexplorerGetCostForecastResult{Result: future}
}

func (a *CostExplorerStub) GetDimensionValues(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error) {
	var output costexplorer.GetDimensionValuesOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetDimensionValues", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetDimensionValuesAsync(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) *CostexplorerGetDimensionValuesResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetDimensionValues", input)
	return &CostexplorerGetDimensionValuesResult{Result: future}
}

func (a *CostExplorerStub) GetReservationCoverage(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error) {
	var output costexplorer.GetReservationCoverageOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationCoverage", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetReservationCoverageAsync(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) *CostexplorerGetReservationCoverageResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationCoverage", input)
	return &CostexplorerGetReservationCoverageResult{Result: future}
}

func (a *CostExplorerStub) GetReservationPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	var output costexplorer.GetReservationPurchaseRecommendationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationPurchaseRecommendation", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetReservationPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) *CostexplorerGetReservationPurchaseRecommendationResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationPurchaseRecommendation", input)
	return &CostexplorerGetReservationPurchaseRecommendationResult{Result: future}
}

func (a *CostExplorerStub) GetReservationUtilization(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error) {
	var output costexplorer.GetReservationUtilizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationUtilization", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetReservationUtilizationAsync(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) *CostexplorerGetReservationUtilizationResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationUtilization", input)
	return &CostexplorerGetReservationUtilizationResult{Result: future}
}

func (a *CostExplorerStub) GetRightsizingRecommendation(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	var output costexplorer.GetRightsizingRecommendationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetRightsizingRecommendation", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetRightsizingRecommendationAsync(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) *CostexplorerGetRightsizingRecommendationResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetRightsizingRecommendation", input)
	return &CostexplorerGetRightsizingRecommendationResult{Result: future}
}

func (a *CostExplorerStub) GetSavingsPlansCoverage(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	var output costexplorer.GetSavingsPlansCoverageOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansCoverage", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansCoverageAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) *CostexplorerGetSavingsPlansCoverageResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansCoverage", input)
	return &CostexplorerGetSavingsPlansCoverageResult{Result: future}
}

func (a *CostExplorerStub) GetSavingsPlansPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	var output costexplorer.GetSavingsPlansPurchaseRecommendationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansPurchaseRecommendation", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) *CostexplorerGetSavingsPlansPurchaseRecommendationResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansPurchaseRecommendation", input)
	return &CostexplorerGetSavingsPlansPurchaseRecommendationResult{Result: future}
}

func (a *CostExplorerStub) GetSavingsPlansUtilization(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilization", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansUtilizationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) *CostexplorerGetSavingsPlansUtilizationResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilization", input)
	return &CostexplorerGetSavingsPlansUtilizationResult{Result: future}
}

func (a *CostExplorerStub) GetSavingsPlansUtilizationDetails(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationDetailsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilizationDetails", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansUtilizationDetailsAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) *CostexplorerGetSavingsPlansUtilizationDetailsResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilizationDetails", input)
	return &CostexplorerGetSavingsPlansUtilizationDetailsResult{Result: future}
}

func (a *CostExplorerStub) GetTags(ctx workflow.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error) {
	var output costexplorer.GetTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetTags", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetTagsAsync(ctx workflow.Context, input *costexplorer.GetTagsInput) *CostexplorerGetTagsResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetTags", input)
	return &CostexplorerGetTagsResult{Result: future}
}

func (a *CostExplorerStub) GetUsageForecast(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error) {
	var output costexplorer.GetUsageForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetUsageForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetUsageForecastAsync(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) *CostexplorerGetUsageForecastResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetUsageForecast", input)
	return &CostexplorerGetUsageForecastResult{Result: future}
}

func (a *CostExplorerStub) ListCostCategoryDefinitions(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	var output costexplorer.ListCostCategoryDefinitionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.ListCostCategoryDefinitions", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) ListCostCategoryDefinitionsAsync(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) *CostexplorerListCostCategoryDefinitionsResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.ListCostCategoryDefinitions", input)
	return &CostexplorerListCostCategoryDefinitionsResult{Result: future}
}

func (a *CostExplorerStub) ProvideAnomalyFeedback(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error) {
	var output costexplorer.ProvideAnomalyFeedbackOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.ProvideAnomalyFeedback", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) ProvideAnomalyFeedbackAsync(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) *CostexplorerProvideAnomalyFeedbackResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.ProvideAnomalyFeedback", input)
	return &CostexplorerProvideAnomalyFeedbackResult{Result: future}
}

func (a *CostExplorerStub) UpdateAnomalyMonitor(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error) {
	var output costexplorer.UpdateAnomalyMonitorOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalyMonitor", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) UpdateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) *CostexplorerUpdateAnomalyMonitorResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalyMonitor", input)
	return &CostexplorerUpdateAnomalyMonitorResult{Result: future}
}

func (a *CostExplorerStub) UpdateAnomalySubscription(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error) {
	var output costexplorer.UpdateAnomalySubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalySubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) UpdateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) *CostexplorerUpdateAnomalySubscriptionResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalySubscription", input)
	return &CostexplorerUpdateAnomalySubscriptionResult{Result: future}
}

func (a *CostExplorerStub) UpdateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
	var output costexplorer.UpdateCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) UpdateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) *CostexplorerUpdateCostCategoryDefinitionResult {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateCostCategoryDefinition", input)
	return &CostexplorerUpdateCostCategoryDefinitionResult{Result: future}
}
