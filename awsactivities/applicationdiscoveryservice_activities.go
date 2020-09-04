
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface"
)

type ApplicationDiscoveryServiceActivities struct {
	client applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI
}

func NewApplicationDiscoveryServiceActivities(client applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI) *ApplicationDiscoveryServiceActivities {
    return &ApplicationDiscoveryServiceActivities{client: client}
}

func (a *ApplicationDiscoveryServiceActivities) AssociateConfigurationItemsToApplication(input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
    return a.client.AssociateConfigurationItemsToApplication(input)
}

func (a *ApplicationDiscoveryServiceActivities) BatchDeleteImportData(input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
    return a.client.BatchDeleteImportData(input)
}

func (a *ApplicationDiscoveryServiceActivities) CreateApplication(input *applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error) {
    return a.client.CreateApplication(input)
}

func (a *ApplicationDiscoveryServiceActivities) CreateTags(input *applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error) {
    return a.client.CreateTags(input)
}

func (a *ApplicationDiscoveryServiceActivities) DeleteApplications(input *applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error) {
    return a.client.DeleteApplications(input)
}

func (a *ApplicationDiscoveryServiceActivities) DeleteTags(input *applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeAgents(input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
    return a.client.DescribeAgents(input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeConfigurations(input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
    return a.client.DescribeConfigurations(input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeContinuousExports(input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
    return a.client.DescribeContinuousExports(input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeExportConfigurations(input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
    return a.client.DescribeExportConfigurations(input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeExportTasks(input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
    return a.client.DescribeExportTasks(input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeImportTasks(input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
    return a.client.DescribeImportTasks(input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeTags(input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}

func (a *ApplicationDiscoveryServiceActivities) DisassociateConfigurationItemsFromApplication(input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error) {
    return a.client.DisassociateConfigurationItemsFromApplication(input)
}

func (a *ApplicationDiscoveryServiceActivities) ExportConfigurations(input *applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error) {
    return a.client.ExportConfigurations(input)
}

func (a *ApplicationDiscoveryServiceActivities) GetDiscoverySummary(input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
    return a.client.GetDiscoverySummary(input)
}

func (a *ApplicationDiscoveryServiceActivities) ListConfigurations(input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
    return a.client.ListConfigurations(input)
}

func (a *ApplicationDiscoveryServiceActivities) ListServerNeighbors(input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
    return a.client.ListServerNeighbors(input)
}

func (a *ApplicationDiscoveryServiceActivities) StartContinuousExport(input *applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error) {
    return a.client.StartContinuousExport(input)
}

func (a *ApplicationDiscoveryServiceActivities) StartDataCollectionByAgentIds(input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error) {
    return a.client.StartDataCollectionByAgentIds(input)
}

func (a *ApplicationDiscoveryServiceActivities) StartExportTask(input *applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error) {
    return a.client.StartExportTask(input)
}

func (a *ApplicationDiscoveryServiceActivities) StartImportTask(input *applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error) {
    return a.client.StartImportTask(input)
}

func (a *ApplicationDiscoveryServiceActivities) StopContinuousExport(input *applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error) {
    return a.client.StopContinuousExport(input)
}

func (a *ApplicationDiscoveryServiceActivities) StopDataCollectionByAgentIds(input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error) {
    return a.client.StopDataCollectionByAgentIds(input)
}

func (a *ApplicationDiscoveryServiceActivities) UpdateApplication(input *applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error) {
    return a.client.UpdateApplication(input)
}
