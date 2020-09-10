package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ApplicationDiscoveryServiceActivities struct {
	client applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI
}

func NewApplicationDiscoveryServiceActivities(session *session.Session, config ...*aws.Config) *ApplicationDiscoveryServiceActivities {
	client := applicationdiscoveryservice.New(session, config...)
	return &ApplicationDiscoveryServiceActivities{client: client}
}

func (a *ApplicationDiscoveryServiceActivities) AssociateConfigurationItemsToApplication(ctx context.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
	return a.client.AssociateConfigurationItemsToApplicationWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) BatchDeleteImportData(ctx context.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
	return a.client.BatchDeleteImportDataWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) CreateApplication(ctx context.Context, input *applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error) {
	return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) CreateTags(ctx context.Context, input *applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error) {
	return a.client.CreateTagsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DeleteApplications(ctx context.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error) {
	return a.client.DeleteApplicationsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DeleteTags(ctx context.Context, input *applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error) {
	return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeAgents(ctx context.Context, input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
	return a.client.DescribeAgentsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeConfigurations(ctx context.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
	return a.client.DescribeConfigurationsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeContinuousExports(ctx context.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
	return a.client.DescribeContinuousExportsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeExportConfigurations(ctx context.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
	return a.client.DescribeExportConfigurationsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeExportTasks(ctx context.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
	return a.client.DescribeExportTasksWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeImportTasks(ctx context.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
	return a.client.DescribeImportTasksWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DescribeTags(ctx context.Context, input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
	return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) DisassociateConfigurationItemsFromApplication(ctx context.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error) {
	return a.client.DisassociateConfigurationItemsFromApplicationWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) ExportConfigurations(ctx context.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error) {
	return a.client.ExportConfigurationsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) GetDiscoverySummary(ctx context.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
	return a.client.GetDiscoverySummaryWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) ListConfigurations(ctx context.Context, input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
	return a.client.ListConfigurationsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) ListServerNeighbors(ctx context.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
	return a.client.ListServerNeighborsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) StartContinuousExport(ctx context.Context, input *applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error) {
	return a.client.StartContinuousExportWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) StartDataCollectionByAgentIds(ctx context.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error) {
	return a.client.StartDataCollectionByAgentIdsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) StartExportTask(ctx context.Context, input *applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error) {
	return a.client.StartExportTaskWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) StartImportTask(ctx context.Context, input *applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error) {
	return a.client.StartImportTaskWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) StopContinuousExport(ctx context.Context, input *applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error) {
	return a.client.StopContinuousExportWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) StopDataCollectionByAgentIds(ctx context.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error) {
	return a.client.StopDataCollectionByAgentIdsWithContext(ctx, input)
}

func (a *ApplicationDiscoveryServiceActivities) UpdateApplication(ctx context.Context, input *applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error) {
	return a.client.UpdateApplicationWithContext(ctx, input)
}
