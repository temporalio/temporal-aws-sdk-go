// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client databasemigrationserviceiface.DatabaseMigrationServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := databasemigrationservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (databasemigrationserviceiface.DatabaseMigrationServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return databasemigrationservice.New(sess), nil
}

func (a *Activities) AddTagsToResource(ctx context.Context, input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsToResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ApplyPendingMaintenanceAction(ctx context.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ApplyPendingMaintenanceActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CancelReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelReplicationTaskAssessmentRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateEndpoint(ctx context.Context, input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateEventSubscription(ctx context.Context, input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateEventSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateReplicationInstance(ctx context.Context, input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateReplicationInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateReplicationSubnetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateReplicationTask(ctx context.Context, input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateReplicationTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCertificate(ctx context.Context, input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConnection(ctx context.Context, input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEndpoint(ctx context.Context, input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEventSubscription(ctx context.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEventSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReplicationInstance(ctx context.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReplicationInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReplicationSubnetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReplicationTask(ctx context.Context, input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReplicationTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReplicationTaskAssessmentRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAccountAttributes(ctx context.Context, input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAccountAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeApplicableIndividualAssessments(ctx context.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeApplicableIndividualAssessmentsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCertificates(ctx context.Context, input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCertificatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeConnections(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeConnectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEndpointTypes(ctx context.Context, input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEndpointTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEndpoints(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEndpointsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEventCategories(ctx context.Context, input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventCategoriesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEventSubscriptions(ctx context.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventSubscriptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEvents(ctx context.Context, input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeOrderableReplicationInstances(ctx context.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeOrderableReplicationInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePendingMaintenanceActions(ctx context.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePendingMaintenanceActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRefreshSchemasStatus(ctx context.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRefreshSchemasStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReplicationInstanceTaskLogs(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReplicationInstanceTaskLogsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReplicationInstances(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReplicationInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReplicationSubnetGroups(ctx context.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReplicationSubnetGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReplicationTaskAssessmentResults(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReplicationTaskAssessmentResultsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReplicationTaskAssessmentRuns(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReplicationTaskAssessmentRunsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReplicationTaskIndividualAssessments(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReplicationTaskIndividualAssessmentsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReplicationTasks(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReplicationTasksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSchemas(ctx context.Context, input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSchemasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTableStatistics(ctx context.Context, input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTableStatisticsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportCertificate(ctx context.Context, input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyEndpoint(ctx context.Context, input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyEventSubscription(ctx context.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyEventSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyReplicationInstance(ctx context.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyReplicationInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyReplicationSubnetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyReplicationTask(ctx context.Context, input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyReplicationTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RebootReplicationInstance(ctx context.Context, input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RebootReplicationInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RefreshSchemas(ctx context.Context, input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RefreshSchemasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ReloadTables(ctx context.Context, input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ReloadTablesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTagsFromResource(ctx context.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsFromResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartReplicationTask(ctx context.Context, input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartReplicationTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartReplicationTaskAssessment(ctx context.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartReplicationTaskAssessmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartReplicationTaskAssessmentRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopReplicationTask(ctx context.Context, input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopReplicationTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TestConnection(ctx context.Context, input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TestConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilEndpointDeleted(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilEndpointDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilReplicationInstanceAvailable(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilReplicationInstanceAvailableWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilReplicationInstanceDeleted(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilReplicationInstanceDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilReplicationTaskDeleted(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilReplicationTaskDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilReplicationTaskReady(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilReplicationTaskReadyWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilReplicationTaskRunning(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilReplicationTaskRunningWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilReplicationTaskStopped(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilReplicationTaskStoppedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilTestConnectionSucceeds(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilTestConnectionSucceedsWithContext(ctx, input, options...))
	})
}
