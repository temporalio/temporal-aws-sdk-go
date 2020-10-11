// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package costexplorerstub

import (
	"github.com/aws/aws-sdk-go/service/costexplorer"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	CreateAnomalyMonitor(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error)
	CreateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) *CostExplorerCreateAnomalyMonitorFuture

	CreateAnomalySubscription(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error)
	CreateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) *CostExplorerCreateAnomalySubscriptionFuture

	CreateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error)
	CreateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) *CostExplorerCreateCostCategoryDefinitionFuture

	DeleteAnomalyMonitor(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error)
	DeleteAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) *CostExplorerDeleteAnomalyMonitorFuture

	DeleteAnomalySubscription(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error)
	DeleteAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) *CostExplorerDeleteAnomalySubscriptionFuture

	DeleteCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error)
	DeleteCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) *CostExplorerDeleteCostCategoryDefinitionFuture

	DescribeCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	DescribeCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) *CostExplorerDescribeCostCategoryDefinitionFuture

	GetAnomalies(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomaliesAsync(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) *CostExplorerGetAnomaliesFuture

	GetAnomalyMonitors(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalyMonitorsAsync(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) *CostExplorerGetAnomalyMonitorsFuture

	GetAnomalySubscriptions(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetAnomalySubscriptionsAsync(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) *CostExplorerGetAnomalySubscriptionsFuture

	GetCostAndUsage(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) *CostExplorerGetCostAndUsageFuture

	GetCostAndUsageWithResources(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostAndUsageWithResourcesAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) *CostExplorerGetCostAndUsageWithResourcesFuture

	GetCostForecast(ctx workflow.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error)
	GetCostForecastAsync(ctx workflow.Context, input *costexplorer.GetCostForecastInput) *CostExplorerGetCostForecastFuture

	GetDimensionValues(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error)
	GetDimensionValuesAsync(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) *CostExplorerGetDimensionValuesFuture

	GetReservationCoverage(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationCoverageAsync(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) *CostExplorerGetReservationCoverageFuture

	GetReservationPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) *CostExplorerGetReservationPurchaseRecommendationFuture

	GetReservationUtilization(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error)
	GetReservationUtilizationAsync(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) *CostExplorerGetReservationUtilizationFuture

	GetRightsizingRecommendation(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetRightsizingRecommendationAsync(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) *CostExplorerGetRightsizingRecommendationFuture

	GetSavingsPlansCoverage(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansCoverageAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) *CostExplorerGetSavingsPlansCoverageFuture

	GetSavingsPlansPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) *CostExplorerGetSavingsPlansPurchaseRecommendationFuture

	GetSavingsPlansUtilization(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) *CostExplorerGetSavingsPlansUtilizationFuture

	GetSavingsPlansUtilizationDetails(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetSavingsPlansUtilizationDetailsAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) *CostExplorerGetSavingsPlansUtilizationDetailsFuture

	GetTags(ctx workflow.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *costexplorer.GetTagsInput) *CostExplorerGetTagsFuture

	GetUsageForecast(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error)
	GetUsageForecastAsync(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) *CostExplorerGetUsageForecastFuture

	ListCostCategoryDefinitions(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListCostCategoryDefinitionsAsync(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) *CostExplorerListCostCategoryDefinitionsFuture

	ProvideAnomalyFeedback(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error)
	ProvideAnomalyFeedbackAsync(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) *CostExplorerProvideAnomalyFeedbackFuture

	UpdateAnomalyMonitor(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error)
	UpdateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) *CostExplorerUpdateAnomalyMonitorFuture

	UpdateAnomalySubscription(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error)
	UpdateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) *CostExplorerUpdateAnomalySubscriptionFuture

	UpdateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error)
	UpdateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) *CostExplorerUpdateCostCategoryDefinitionFuture
}

func NewClient() Client {
	return &stub{}
}

