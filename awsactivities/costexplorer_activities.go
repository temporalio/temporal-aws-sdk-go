package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CostExplorerActivities struct {
	client costexploreriface.CostExplorerAPI
}

func NewCostExplorerActivities(session *session.Session, config ...*aws.Config) *CostExplorerActivities {
	client := costexplorer.New(session, config...)
	return &CostExplorerActivities{client: client}
}

func (a *CostExplorerActivities) CreateCostCategoryDefinition(ctx context.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
	return a.client.CreateCostCategoryDefinitionWithContext(ctx, input)
}

func (a *CostExplorerActivities) DeleteCostCategoryDefinition(ctx context.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
	return a.client.DeleteCostCategoryDefinitionWithContext(ctx, input)
}

func (a *CostExplorerActivities) DescribeCostCategoryDefinition(ctx context.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	return a.client.DescribeCostCategoryDefinitionWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetCostAndUsage(ctx context.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
	return a.client.GetCostAndUsageWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetCostAndUsageWithResources(ctx context.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	return a.client.GetCostAndUsageWithResourcesWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetCostForecast(ctx context.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error) {
	return a.client.GetCostForecastWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetDimensionValues(ctx context.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error) {
	return a.client.GetDimensionValuesWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetReservationCoverage(ctx context.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error) {
	return a.client.GetReservationCoverageWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetReservationPurchaseRecommendation(ctx context.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	return a.client.GetReservationPurchaseRecommendationWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetReservationUtilization(ctx context.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error) {
	return a.client.GetReservationUtilizationWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetRightsizingRecommendation(ctx context.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	return a.client.GetRightsizingRecommendationWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetSavingsPlansCoverage(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	return a.client.GetSavingsPlansCoverageWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetSavingsPlansPurchaseRecommendation(ctx context.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	return a.client.GetSavingsPlansPurchaseRecommendationWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetSavingsPlansUtilization(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	return a.client.GetSavingsPlansUtilizationWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetSavingsPlansUtilizationDetails(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	return a.client.GetSavingsPlansUtilizationDetailsWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetTags(ctx context.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error) {
	return a.client.GetTagsWithContext(ctx, input)
}

func (a *CostExplorerActivities) GetUsageForecast(ctx context.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error) {
	return a.client.GetUsageForecastWithContext(ctx, input)
}

func (a *CostExplorerActivities) ListCostCategoryDefinitions(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	return a.client.ListCostCategoryDefinitionsWithContext(ctx, input)
}

func (a *CostExplorerActivities) UpdateCostCategoryDefinition(ctx context.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
	return a.client.UpdateCostCategoryDefinitionWithContext(ctx, input)
}
