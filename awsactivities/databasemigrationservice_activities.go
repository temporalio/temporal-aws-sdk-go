// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DatabaseMigrationServiceActivities struct {
	client databasemigrationserviceiface.DatabaseMigrationServiceAPI

	sessionFactory SessionFactory
}

func NewDatabaseMigrationServiceActivities(sess *session.Session, config ...*aws.Config) *DatabaseMigrationServiceActivities {
	client := databasemigrationservice.New(sess, config...)
	return &DatabaseMigrationServiceActivities{client: client}
}

func NewDatabaseMigrationServiceActivitiesWithSessionFactory(sessionFactory SessionFactory) *DatabaseMigrationServiceActivities {
	return &DatabaseMigrationServiceActivities{sessionFactory: sessionFactory}
}

func (a *DatabaseMigrationServiceActivities) getClient(ctx context.Context) (databasemigrationserviceiface.DatabaseMigrationServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return databasemigrationservice.New(sess), nil
}

func (a *DatabaseMigrationServiceActivities) AddTagsToResource(ctx context.Context, input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToResourceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ApplyPendingMaintenanceAction(ctx context.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ApplyPendingMaintenanceActionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CancelReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelReplicationTaskAssessmentRunWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateEndpoint(ctx context.Context, input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEndpointWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateEventSubscription(ctx context.Context, input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEventSubscriptionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateReplicationInstance(ctx context.Context, input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateReplicationSubnetGroupWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateReplicationTask(ctx context.Context, input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteCertificate(ctx context.Context, input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCertificateWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteConnection(ctx context.Context, input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConnectionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteEndpoint(ctx context.Context, input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEndpointWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteEventSubscription(ctx context.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEventSubscriptionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationInstance(ctx context.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReplicationSubnetGroupWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationTask(ctx context.Context, input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReplicationTaskAssessmentRunWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeAccountAttributes(ctx context.Context, input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAccountAttributesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeApplicableIndividualAssessments(ctx context.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeApplicableIndividualAssessmentsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeCertificates(ctx context.Context, input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCertificatesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeConnections(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConnectionsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEndpointTypes(ctx context.Context, input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEndpointTypesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEndpoints(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEndpointsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEventCategories(ctx context.Context, input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEventSubscriptions(ctx context.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventSubscriptionsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEvents(ctx context.Context, input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeOrderableReplicationInstances(ctx context.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeOrderableReplicationInstancesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribePendingMaintenanceActions(ctx context.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePendingMaintenanceActionsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeRefreshSchemasStatus(ctx context.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRefreshSchemasStatusWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationInstanceTaskLogs(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationInstanceTaskLogsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationInstances(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationInstancesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationSubnetGroups(ctx context.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationSubnetGroupsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskAssessmentResults(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationTaskAssessmentResultsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskAssessmentRuns(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationTaskAssessmentRunsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskIndividualAssessments(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationTaskIndividualAssessmentsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTasks(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationTasksWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeSchemas(ctx context.Context, input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSchemasWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeTableStatistics(ctx context.Context, input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTableStatisticsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ImportCertificate(ctx context.Context, input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ImportCertificateWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ListTagsForResource(ctx context.Context, input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyEndpoint(ctx context.Context, input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyEndpointWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyEventSubscription(ctx context.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyEventSubscriptionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyReplicationInstance(ctx context.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyReplicationSubnetGroupWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyReplicationTask(ctx context.Context, input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) RebootReplicationInstance(ctx context.Context, input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) RefreshSchemas(ctx context.Context, input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RefreshSchemasWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ReloadTables(ctx context.Context, input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ReloadTablesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) RemoveTagsFromResource(ctx context.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StartReplicationTask(ctx context.Context, input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StartReplicationTaskAssessment(ctx context.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartReplicationTaskAssessmentWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StartReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartReplicationTaskAssessmentRunWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StopReplicationTask(ctx context.Context, input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) TestConnection(ctx context.Context, input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestConnectionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) WaitUntilEndpointDeleted(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilEndpointDeletedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationInstanceAvailable(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationInstanceAvailableWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationInstanceDeleted(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationInstanceDeletedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskDeleted(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationTaskDeletedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskReady(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationTaskReadyWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskRunning(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationTaskRunningWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskStopped(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationTaskStoppedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilTestConnectionSucceeds(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilTestConnectionSucceedsWithContext(ctx, input, options...)
	})
}
