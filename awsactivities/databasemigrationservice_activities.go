package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
)

type DatabaseMigrationServiceActivities struct {
	client databasemigrationserviceiface.DatabaseMigrationServiceAPI
}

func NewDatabaseMigrationServiceActivities(client databasemigrationserviceiface.DatabaseMigrationServiceAPI) *DatabaseMigrationServiceActivities {
    return &DatabaseMigrationServiceActivities{client: client}
}


func (a *DatabaseMigrationServiceActivities) AddTagsToResource(input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResource(input)
}



func (a *DatabaseMigrationServiceActivities) ApplyPendingMaintenanceAction(input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error) {
    return a.client.ApplyPendingMaintenanceAction(input)
}



func (a *DatabaseMigrationServiceActivities) CancelReplicationTaskAssessmentRun(input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error) {
    return a.client.CancelReplicationTaskAssessmentRun(input)
}



func (a *DatabaseMigrationServiceActivities) CreateEndpoint(input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error) {
    return a.client.CreateEndpoint(input)
}



func (a *DatabaseMigrationServiceActivities) CreateEventSubscription(input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error) {
    return a.client.CreateEventSubscription(input)
}



func (a *DatabaseMigrationServiceActivities) CreateReplicationInstance(input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error) {
    return a.client.CreateReplicationInstance(input)
}



func (a *DatabaseMigrationServiceActivities) CreateReplicationSubnetGroup(input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error) {
    return a.client.CreateReplicationSubnetGroup(input)
}



func (a *DatabaseMigrationServiceActivities) CreateReplicationTask(input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error) {
    return a.client.CreateReplicationTask(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteCertificate(input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error) {
    return a.client.DeleteCertificate(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteConnection(input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error) {
    return a.client.DeleteConnection(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteEndpoint(input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error) {
    return a.client.DeleteEndpoint(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteEventSubscription(input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error) {
    return a.client.DeleteEventSubscription(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteReplicationInstance(input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error) {
    return a.client.DeleteReplicationInstance(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteReplicationSubnetGroup(input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error) {
    return a.client.DeleteReplicationSubnetGroup(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteReplicationTask(input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error) {
    return a.client.DeleteReplicationTask(input)
}



func (a *DatabaseMigrationServiceActivities) DeleteReplicationTaskAssessmentRun(input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error) {
    return a.client.DeleteReplicationTaskAssessmentRun(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeAccountAttributes(input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributes(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeApplicableIndividualAssessments(input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
    return a.client.DescribeApplicableIndividualAssessments(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeCertificates(input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error) {
    return a.client.DescribeCertificates(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeConnections(input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error) {
    return a.client.DescribeConnections(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeEndpointTypes(input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
    return a.client.DescribeEndpointTypes(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeEndpoints(input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error) {
    return a.client.DescribeEndpoints(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeEventCategories(input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
    return a.client.DescribeEventCategories(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeEventSubscriptions(input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
    return a.client.DescribeEventSubscriptions(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeEvents(input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeOrderableReplicationInstances(input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
    return a.client.DescribeOrderableReplicationInstances(input)
}



func (a *DatabaseMigrationServiceActivities) DescribePendingMaintenanceActions(input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
    return a.client.DescribePendingMaintenanceActions(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeRefreshSchemasStatus(input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
    return a.client.DescribeRefreshSchemasStatus(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeReplicationInstanceTaskLogs(input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
    return a.client.DescribeReplicationInstanceTaskLogs(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeReplicationInstances(input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
    return a.client.DescribeReplicationInstances(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeReplicationSubnetGroups(input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
    return a.client.DescribeReplicationSubnetGroups(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskAssessmentResults(input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
    return a.client.DescribeReplicationTaskAssessmentResults(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskAssessmentRuns(input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
    return a.client.DescribeReplicationTaskAssessmentRuns(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskIndividualAssessments(input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
    return a.client.DescribeReplicationTaskIndividualAssessments(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeReplicationTasks(input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
    return a.client.DescribeReplicationTasks(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeSchemas(input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error) {
    return a.client.DescribeSchemas(input)
}



func (a *DatabaseMigrationServiceActivities) DescribeTableStatistics(input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
    return a.client.DescribeTableStatistics(input)
}



func (a *DatabaseMigrationServiceActivities) ImportCertificate(input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error) {
    return a.client.ImportCertificate(input)
}



func (a *DatabaseMigrationServiceActivities) ListTagsForResource(input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *DatabaseMigrationServiceActivities) ModifyEndpoint(input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error) {
    return a.client.ModifyEndpoint(input)
}



func (a *DatabaseMigrationServiceActivities) ModifyEventSubscription(input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error) {
    return a.client.ModifyEventSubscription(input)
}



func (a *DatabaseMigrationServiceActivities) ModifyReplicationInstance(input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error) {
    return a.client.ModifyReplicationInstance(input)
}



func (a *DatabaseMigrationServiceActivities) ModifyReplicationSubnetGroup(input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error) {
    return a.client.ModifyReplicationSubnetGroup(input)
}



func (a *DatabaseMigrationServiceActivities) ModifyReplicationTask(input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error) {
    return a.client.ModifyReplicationTask(input)
}



func (a *DatabaseMigrationServiceActivities) RebootReplicationInstance(input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error) {
    return a.client.RebootReplicationInstance(input)
}



func (a *DatabaseMigrationServiceActivities) RefreshSchemas(input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error) {
    return a.client.RefreshSchemas(input)
}



func (a *DatabaseMigrationServiceActivities) ReloadTables(input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error) {
    return a.client.ReloadTables(input)
}



func (a *DatabaseMigrationServiceActivities) RemoveTagsFromResource(input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResource(input)
}



func (a *DatabaseMigrationServiceActivities) StartReplicationTask(input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error) {
    return a.client.StartReplicationTask(input)
}



func (a *DatabaseMigrationServiceActivities) StartReplicationTaskAssessment(input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error) {
    return a.client.StartReplicationTaskAssessment(input)
}



func (a *DatabaseMigrationServiceActivities) StartReplicationTaskAssessmentRun(input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error) {
    return a.client.StartReplicationTaskAssessmentRun(input)
}



func (a *DatabaseMigrationServiceActivities) StopReplicationTask(input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error) {
    return a.client.StopReplicationTask(input)
}



func (a *DatabaseMigrationServiceActivities) TestConnection(input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error) {
    return a.client.TestConnection(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilEndpointDeleted(input *databasemigrationservice.DescribeEndpointsInput) error {
    return a.client.WaitUntilEndpointDeleted(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationInstanceAvailable(input *databasemigrationservice.DescribeReplicationInstancesInput) error {
    return a.client.WaitUntilReplicationInstanceAvailable(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationInstanceDeleted(input *databasemigrationservice.DescribeReplicationInstancesInput) error {
    return a.client.WaitUntilReplicationInstanceDeleted(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskDeleted(input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return a.client.WaitUntilReplicationTaskDeleted(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskReady(input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return a.client.WaitUntilReplicationTaskReady(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskRunning(input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return a.client.WaitUntilReplicationTaskRunning(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskStopped(input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return a.client.WaitUntilReplicationTaskStopped(input)
}



func (a *DatabaseMigrationServiceActivities) WaitUntilTestConnectionSucceeds(input *databasemigrationservice.DescribeConnectionsInput) error {
    return a.client.WaitUntilTestConnectionSucceeds(input)
}

