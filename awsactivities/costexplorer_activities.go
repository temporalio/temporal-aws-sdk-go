package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
)

type CostExplorerActivities struct {
	client costexploreriface.CostExplorerAPI
}

func NewCostExplorerActivities(client costexploreriface.CostExplorerAPI) *CostExplorerActivities {
    return &CostExplorerActivities{client: client}
}


func (a *CostExplorerActivities) CreateCostCategoryDefinition(input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
    return a.client.CreateCostCategoryDefinition(input)
}



func (a *CostExplorerActivities) DeleteCostCategoryDefinition(input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
    return a.client.DeleteCostCategoryDefinition(input)
}



func (a *CostExplorerActivities) DescribeCostCategoryDefinition(input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
    return a.client.DescribeCostCategoryDefinition(input)
}



func (a *CostExplorerActivities) GetCostAndUsage(input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
    return a.client.GetCostAndUsage(input)
}



func (a *CostExplorerActivities) GetCostAndUsageWithResources(input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
    return a.client.GetCostAndUsageWithResources(input)
}



func (a *CostExplorerActivities) GetCostForecast(input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error) {
    return a.client.GetCostForecast(input)
}



func (a *CostExplorerActivities) GetDimensionValues(input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error) {
    return a.client.GetDimensionValues(input)
}



func (a *CostExplorerActivities) GetReservationCoverage(input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error) {
    return a.client.GetReservationCoverage(input)
}



func (a *CostExplorerActivities) GetReservationPurchaseRecommendation(input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
    return a.client.GetReservationPurchaseRecommendation(input)
}



func (a *CostExplorerActivities) GetReservationUtilization(input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error) {
    return a.client.GetReservationUtilization(input)
}



func (a *CostExplorerActivities) GetRightsizingRecommendation(input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error) {
    return a.client.GetRightsizingRecommendation(input)
}



func (a *CostExplorerActivities) GetSavingsPlansCoverage(input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
    return a.client.GetSavingsPlansCoverage(input)
}



func (a *CostExplorerActivities) GetSavingsPlansPurchaseRecommendation(input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
    return a.client.GetSavingsPlansPurchaseRecommendation(input)
}



func (a *CostExplorerActivities) GetSavingsPlansUtilization(input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
    return a.client.GetSavingsPlansUtilization(input)
}



func (a *CostExplorerActivities) GetSavingsPlansUtilizationDetails(input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
    return a.client.GetSavingsPlansUtilizationDetails(input)
}



func (a *CostExplorerActivities) GetTags(input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error) {
    return a.client.GetTags(input)
}



func (a *CostExplorerActivities) GetUsageForecast(input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error) {
    return a.client.GetUsageForecast(input)
}



func (a *CostExplorerActivities) ListCostCategoryDefinitions(input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
    return a.client.ListCostCategoryDefinitions(input)
}



func (a *CostExplorerActivities) UpdateCostCategoryDefinition(input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
    return a.client.UpdateCostCategoryDefinition(input)
}

