// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package databasemigrationservicestub

import (
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	AddTagsToResource(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) *DatabaseMigrationServiceAddTagsToResourceFuture

	ApplyPendingMaintenanceAction(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error)
	ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) *DatabaseMigrationServiceApplyPendingMaintenanceActionFuture

	CancelReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error)
	CancelReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) *DatabaseMigrationServiceCancelReplicationTaskAssessmentRunFuture

	CreateEndpoint(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) *DatabaseMigrationServiceCreateEndpointFuture

	CreateEventSubscription(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) *DatabaseMigrationServiceCreateEventSubscriptionFuture

	CreateReplicationInstance(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error)
	CreateReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) *DatabaseMigrationServiceCreateReplicationInstanceFuture

	CreateReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error)
	CreateReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) *DatabaseMigrationServiceCreateReplicationSubnetGroupFuture

	CreateReplicationTask(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error)
	CreateReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) *DatabaseMigrationServiceCreateReplicationTaskFuture

	DeleteCertificate(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error)
	DeleteCertificateAsync(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) *DatabaseMigrationServiceDeleteCertificateFuture

	DeleteConnection(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) *DatabaseMigrationServiceDeleteConnectionFuture

	DeleteEndpoint(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) *DatabaseMigrationServiceDeleteEndpointFuture

	DeleteEventSubscription(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) *DatabaseMigrationServiceDeleteEventSubscriptionFuture

	DeleteReplicationInstance(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error)
	DeleteReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) *DatabaseMigrationServiceDeleteReplicationInstanceFuture

	DeleteReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error)
	DeleteReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) *DatabaseMigrationServiceDeleteReplicationSubnetGroupFuture

	DeleteReplicationTask(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error)
	DeleteReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) *DatabaseMigrationServiceDeleteReplicationTaskFuture

	DeleteReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error)
	DeleteReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) *DatabaseMigrationServiceDeleteReplicationTaskAssessmentRunFuture

	DescribeAccountAttributes(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) *DatabaseMigrationServiceDescribeAccountAttributesFuture

	DescribeApplicableIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error)
	DescribeApplicableIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) *DatabaseMigrationServiceDescribeApplicableIndividualAssessmentsFuture

	DescribeCertificates(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error)
	DescribeCertificatesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) *DatabaseMigrationServiceDescribeCertificatesFuture

	DescribeConnections(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error)
	DescribeConnectionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) *DatabaseMigrationServiceDescribeConnectionsFuture

	DescribeEndpointTypes(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error)
	DescribeEndpointTypesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) *DatabaseMigrationServiceDescribeEndpointTypesFuture

	DescribeEndpoints(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error)
	DescribeEndpointsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) *DatabaseMigrationServiceDescribeEndpointsFuture

	DescribeEventCategories(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) *DatabaseMigrationServiceDescribeEventCategoriesFuture

	DescribeEventSubscriptions(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) *DatabaseMigrationServiceDescribeEventSubscriptionsFuture

	DescribeEvents(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) *DatabaseMigrationServiceDescribeEventsFuture

	DescribeOrderableReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)
	DescribeOrderableReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) *DatabaseMigrationServiceDescribeOrderableReplicationInstancesFuture

	DescribePendingMaintenanceActions(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error)
	DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) *DatabaseMigrationServiceDescribePendingMaintenanceActionsFuture

	DescribeRefreshSchemasStatus(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)
	DescribeRefreshSchemasStatusAsync(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) *DatabaseMigrationServiceDescribeRefreshSchemasStatusFuture

	DescribeReplicationInstanceTaskLogs(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error)
	DescribeReplicationInstanceTaskLogsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) *DatabaseMigrationServiceDescribeReplicationInstanceTaskLogsFuture

	DescribeReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	DescribeReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *DatabaseMigrationServiceDescribeReplicationInstancesFuture

	DescribeReplicationSubnetGroups(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)
	DescribeReplicationSubnetGroupsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) *DatabaseMigrationServiceDescribeReplicationSubnetGroupsFuture

	DescribeReplicationTaskAssessmentResults(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error)
	DescribeReplicationTaskAssessmentResultsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) *DatabaseMigrationServiceDescribeReplicationTaskAssessmentResultsFuture

	DescribeReplicationTaskAssessmentRuns(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error)
	DescribeReplicationTaskAssessmentRunsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) *DatabaseMigrationServiceDescribeReplicationTaskAssessmentRunsFuture

	DescribeReplicationTaskIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error)
	DescribeReplicationTaskIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) *DatabaseMigrationServiceDescribeReplicationTaskIndividualAssessmentsFuture

	DescribeReplicationTasks(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error)
	DescribeReplicationTasksAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *DatabaseMigrationServiceDescribeReplicationTasksFuture

	DescribeSchemas(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error)
	DescribeSchemasAsync(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) *DatabaseMigrationServiceDescribeSchemasFuture

	DescribeTableStatistics(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error)
	DescribeTableStatisticsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) *DatabaseMigrationServiceDescribeTableStatisticsFuture

	ImportCertificate(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error)
	ImportCertificateAsync(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) *DatabaseMigrationServiceImportCertificateFuture

	ListTagsForResource(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) *DatabaseMigrationServiceListTagsForResourceFuture

	ModifyEndpoint(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error)
	ModifyEndpointAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) *DatabaseMigrationServiceModifyEndpointFuture

	ModifyEventSubscription(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) *DatabaseMigrationServiceModifyEventSubscriptionFuture

	ModifyReplicationInstance(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error)
	ModifyReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) *DatabaseMigrationServiceModifyReplicationInstanceFuture

	ModifyReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error)
	ModifyReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) *DatabaseMigrationServiceModifyReplicationSubnetGroupFuture

	ModifyReplicationTask(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error)
	ModifyReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) *DatabaseMigrationServiceModifyReplicationTaskFuture

	RebootReplicationInstance(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error)
	RebootReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) *DatabaseMigrationServiceRebootReplicationInstanceFuture

	RefreshSchemas(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error)
	RefreshSchemasAsync(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) *DatabaseMigrationServiceRefreshSchemasFuture

	ReloadTables(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error)
	ReloadTablesAsync(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) *DatabaseMigrationServiceReloadTablesFuture

	RemoveTagsFromResource(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) *DatabaseMigrationServiceRemoveTagsFromResourceFuture

	StartReplicationTask(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error)
	StartReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) *DatabaseMigrationServiceStartReplicationTaskFuture

	StartReplicationTaskAssessment(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error)
	StartReplicationTaskAssessmentAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) *DatabaseMigrationServiceStartReplicationTaskAssessmentFuture

	StartReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error)
	StartReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) *DatabaseMigrationServiceStartReplicationTaskAssessmentRunFuture

	StopReplicationTask(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error)
	StopReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) *DatabaseMigrationServiceStopReplicationTaskFuture

	TestConnection(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error)
	TestConnectionAsync(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) *DatabaseMigrationServiceTestConnectionFuture

	WaitUntilEndpointDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) error
	WaitUntilEndpointDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) *awsclients.VoidFuture

	WaitUntilReplicationInstanceAvailable(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error
	WaitUntilReplicationInstanceAvailableAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *awsclients.VoidFuture

	WaitUntilReplicationInstanceDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error
	WaitUntilReplicationInstanceDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *awsclients.VoidFuture

	WaitUntilReplicationTaskDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *awsclients.VoidFuture

	WaitUntilReplicationTaskReady(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskReadyAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *awsclients.VoidFuture

	WaitUntilReplicationTaskRunning(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskRunningAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *awsclients.VoidFuture

	WaitUntilReplicationTaskStopped(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
	WaitUntilReplicationTaskStoppedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *awsclients.VoidFuture

	WaitUntilTestConnectionSucceeds(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) error
	WaitUntilTestConnectionSucceedsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}

