// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/neptune/neptuneiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type NeptuneActivities struct {
	client neptuneiface.NeptuneAPI

	sessionFactory SessionFactory
}

func NewNeptuneActivities(sess *session.Session, config ...*aws.Config) *NeptuneActivities {
	client := neptune.New(sess, config...)
	return &NeptuneActivities{client: client}
}

func NewNeptuneActivitiesWithSessionFactory(sessionFactory SessionFactory) *NeptuneActivities {
	return &NeptuneActivities{sessionFactory: sessionFactory}
}

func (a *NeptuneActivities) getClient(ctx context.Context) (neptuneiface.NeptuneAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return neptune.New(sess), nil
}

func (a *NeptuneActivities) AddRoleToDBCluster(ctx context.Context, input *neptune.AddRoleToDBClusterInput) (*neptune.AddRoleToDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddRoleToDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) AddSourceIdentifierToSubscription(ctx context.Context, input *neptune.AddSourceIdentifierToSubscriptionInput) (*neptune.AddSourceIdentifierToSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddSourceIdentifierToSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) AddTagsToResource(ctx context.Context, input *neptune.AddTagsToResourceInput) (*neptune.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToResourceWithContext(ctx, input)
}

func (a *NeptuneActivities) ApplyPendingMaintenanceAction(ctx context.Context, input *neptune.ApplyPendingMaintenanceActionInput) (*neptune.ApplyPendingMaintenanceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ApplyPendingMaintenanceActionWithContext(ctx, input)
}

func (a *NeptuneActivities) CopyDBClusterParameterGroup(ctx context.Context, input *neptune.CopyDBClusterParameterGroupInput) (*neptune.CopyDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CopyDBClusterSnapshot(ctx context.Context, input *neptune.CopyDBClusterSnapshotInput) (*neptune.CopyDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyDBClusterSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) CopyDBParameterGroup(ctx context.Context, input *neptune.CopyDBParameterGroupInput) (*neptune.CopyDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBCluster(ctx context.Context, input *neptune.CreateDBClusterInput) (*neptune.CreateDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBClusterParameterGroup(ctx context.Context, input *neptune.CreateDBClusterParameterGroupInput) (*neptune.CreateDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBClusterSnapshot(ctx context.Context, input *neptune.CreateDBClusterSnapshotInput) (*neptune.CreateDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBClusterSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBInstance(ctx context.Context, input *neptune.CreateDBInstanceInput) (*neptune.CreateDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBParameterGroup(ctx context.Context, input *neptune.CreateDBParameterGroupInput) (*neptune.CreateDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBSubnetGroup(ctx context.Context, input *neptune.CreateDBSubnetGroupInput) (*neptune.CreateDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBSubnetGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateEventSubscription(ctx context.Context, input *neptune.CreateEventSubscriptionInput) (*neptune.CreateEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEventSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBCluster(ctx context.Context, input *neptune.DeleteDBClusterInput) (*neptune.DeleteDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBClusterParameterGroup(ctx context.Context, input *neptune.DeleteDBClusterParameterGroupInput) (*neptune.DeleteDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBClusterSnapshot(ctx context.Context, input *neptune.DeleteDBClusterSnapshotInput) (*neptune.DeleteDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBClusterSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBInstance(ctx context.Context, input *neptune.DeleteDBInstanceInput) (*neptune.DeleteDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBParameterGroup(ctx context.Context, input *neptune.DeleteDBParameterGroupInput) (*neptune.DeleteDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBSubnetGroup(ctx context.Context, input *neptune.DeleteDBSubnetGroupInput) (*neptune.DeleteDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBSubnetGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteEventSubscription(ctx context.Context, input *neptune.DeleteEventSubscriptionInput) (*neptune.DeleteEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEventSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterParameterGroups(ctx context.Context, input *neptune.DescribeDBClusterParameterGroupsInput) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterParameterGroupsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterParameters(ctx context.Context, input *neptune.DescribeDBClusterParametersInput) (*neptune.DescribeDBClusterParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterSnapshotAttributes(ctx context.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterSnapshotAttributesWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterSnapshots(ctx context.Context, input *neptune.DescribeDBClusterSnapshotsInput) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterSnapshotsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusters(ctx context.Context, input *neptune.DescribeDBClustersInput) (*neptune.DescribeDBClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClustersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBEngineVersions(ctx context.Context, input *neptune.DescribeDBEngineVersionsInput) (*neptune.DescribeDBEngineVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBEngineVersionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBInstances(ctx context.Context, input *neptune.DescribeDBInstancesInput) (*neptune.DescribeDBInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBInstancesWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBParameterGroups(ctx context.Context, input *neptune.DescribeDBParameterGroupsInput) (*neptune.DescribeDBParameterGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBParameterGroupsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBParameters(ctx context.Context, input *neptune.DescribeDBParametersInput) (*neptune.DescribeDBParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBSubnetGroups(ctx context.Context, input *neptune.DescribeDBSubnetGroupsInput) (*neptune.DescribeDBSubnetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBSubnetGroupsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEngineDefaultClusterParameters(ctx context.Context, input *neptune.DescribeEngineDefaultClusterParametersInput) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEngineDefaultClusterParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEngineDefaultParameters(ctx context.Context, input *neptune.DescribeEngineDefaultParametersInput) (*neptune.DescribeEngineDefaultParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEngineDefaultParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEventCategories(ctx context.Context, input *neptune.DescribeEventCategoriesInput) (*neptune.DescribeEventCategoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEventSubscriptions(ctx context.Context, input *neptune.DescribeEventSubscriptionsInput) (*neptune.DescribeEventSubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventSubscriptionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEvents(ctx context.Context, input *neptune.DescribeEventsInput) (*neptune.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeOrderableDBInstanceOptions(ctx context.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeOrderableDBInstanceOptionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribePendingMaintenanceActions(ctx context.Context, input *neptune.DescribePendingMaintenanceActionsInput) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePendingMaintenanceActionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeValidDBInstanceModifications(ctx context.Context, input *neptune.DescribeValidDBInstanceModificationsInput) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeValidDBInstanceModificationsWithContext(ctx, input)
}

func (a *NeptuneActivities) FailoverDBCluster(ctx context.Context, input *neptune.FailoverDBClusterInput) (*neptune.FailoverDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.FailoverDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) ListTagsForResource(ctx context.Context, input *neptune.ListTagsForResourceInput) (*neptune.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBCluster(ctx context.Context, input *neptune.ModifyDBClusterInput) (*neptune.ModifyDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBClusterParameterGroup(ctx context.Context, input *neptune.ModifyDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBClusterSnapshotAttribute(ctx context.Context, input *neptune.ModifyDBClusterSnapshotAttributeInput) (*neptune.ModifyDBClusterSnapshotAttributeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBClusterSnapshotAttributeWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBInstance(ctx context.Context, input *neptune.ModifyDBInstanceInput) (*neptune.ModifyDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBParameterGroup(ctx context.Context, input *neptune.ModifyDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBSubnetGroup(ctx context.Context, input *neptune.ModifyDBSubnetGroupInput) (*neptune.ModifyDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBSubnetGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyEventSubscription(ctx context.Context, input *neptune.ModifyEventSubscriptionInput) (*neptune.ModifyEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyEventSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) PromoteReadReplicaDBCluster(ctx context.Context, input *neptune.PromoteReadReplicaDBClusterInput) (*neptune.PromoteReadReplicaDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PromoteReadReplicaDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) RebootDBInstance(ctx context.Context, input *neptune.RebootDBInstanceInput) (*neptune.RebootDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) RemoveRoleFromDBCluster(ctx context.Context, input *neptune.RemoveRoleFromDBClusterInput) (*neptune.RemoveRoleFromDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveRoleFromDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) RemoveSourceIdentifierFromSubscription(ctx context.Context, input *neptune.RemoveSourceIdentifierFromSubscriptionInput) (*neptune.RemoveSourceIdentifierFromSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveSourceIdentifierFromSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) RemoveTagsFromResource(ctx context.Context, input *neptune.RemoveTagsFromResourceInput) (*neptune.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *NeptuneActivities) ResetDBClusterParameterGroup(ctx context.Context, input *neptune.ResetDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ResetDBParameterGroup(ctx context.Context, input *neptune.ResetDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) RestoreDBClusterFromSnapshot(ctx context.Context, input *neptune.RestoreDBClusterFromSnapshotInput) (*neptune.RestoreDBClusterFromSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreDBClusterFromSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) RestoreDBClusterToPointInTime(ctx context.Context, input *neptune.RestoreDBClusterToPointInTimeInput) (*neptune.RestoreDBClusterToPointInTimeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreDBClusterToPointInTimeWithContext(ctx, input)
}

func (a *NeptuneActivities) StartDBCluster(ctx context.Context, input *neptune.StartDBClusterInput) (*neptune.StartDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) StopDBCluster(ctx context.Context, input *neptune.StopDBClusterInput) (*neptune.StopDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) WaitUntilDBInstanceAvailable(ctx context.Context, input *neptune.DescribeDBInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilDBInstanceAvailableWithContext(ctx, input, options...)
	})
}

func (a *NeptuneActivities) WaitUntilDBInstanceDeleted(ctx context.Context, input *neptune.DescribeDBInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilDBInstanceDeletedWithContext(ctx, input, options...)
	})
}
