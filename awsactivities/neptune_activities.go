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
}

func NewNeptuneActivities(session *session.Session, config ...*aws.Config) *NeptuneActivities {
	client := neptune.New(session, config...)
	return &NeptuneActivities{client: client}
}

func (a *NeptuneActivities) AddRoleToDBCluster(ctx context.Context, input *neptune.AddRoleToDBClusterInput) (*neptune.AddRoleToDBClusterOutput, error) {
	return a.client.AddRoleToDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) AddSourceIdentifierToSubscription(ctx context.Context, input *neptune.AddSourceIdentifierToSubscriptionInput) (*neptune.AddSourceIdentifierToSubscriptionOutput, error) {
	return a.client.AddSourceIdentifierToSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) AddTagsToResource(ctx context.Context, input *neptune.AddTagsToResourceInput) (*neptune.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *NeptuneActivities) ApplyPendingMaintenanceAction(ctx context.Context, input *neptune.ApplyPendingMaintenanceActionInput) (*neptune.ApplyPendingMaintenanceActionOutput, error) {
	return a.client.ApplyPendingMaintenanceActionWithContext(ctx, input)
}

func (a *NeptuneActivities) CopyDBClusterParameterGroup(ctx context.Context, input *neptune.CopyDBClusterParameterGroupInput) (*neptune.CopyDBClusterParameterGroupOutput, error) {
	return a.client.CopyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CopyDBClusterSnapshot(ctx context.Context, input *neptune.CopyDBClusterSnapshotInput) (*neptune.CopyDBClusterSnapshotOutput, error) {
	return a.client.CopyDBClusterSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) CopyDBParameterGroup(ctx context.Context, input *neptune.CopyDBParameterGroupInput) (*neptune.CopyDBParameterGroupOutput, error) {
	return a.client.CopyDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBCluster(ctx context.Context, input *neptune.CreateDBClusterInput) (*neptune.CreateDBClusterOutput, error) {
	return a.client.CreateDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBClusterParameterGroup(ctx context.Context, input *neptune.CreateDBClusterParameterGroupInput) (*neptune.CreateDBClusterParameterGroupOutput, error) {
	return a.client.CreateDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBClusterSnapshot(ctx context.Context, input *neptune.CreateDBClusterSnapshotInput) (*neptune.CreateDBClusterSnapshotOutput, error) {
	return a.client.CreateDBClusterSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBInstance(ctx context.Context, input *neptune.CreateDBInstanceInput) (*neptune.CreateDBInstanceOutput, error) {
	return a.client.CreateDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBParameterGroup(ctx context.Context, input *neptune.CreateDBParameterGroupInput) (*neptune.CreateDBParameterGroupOutput, error) {
	return a.client.CreateDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateDBSubnetGroup(ctx context.Context, input *neptune.CreateDBSubnetGroupInput) (*neptune.CreateDBSubnetGroupOutput, error) {
	return a.client.CreateDBSubnetGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) CreateEventSubscription(ctx context.Context, input *neptune.CreateEventSubscriptionInput) (*neptune.CreateEventSubscriptionOutput, error) {
	return a.client.CreateEventSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBCluster(ctx context.Context, input *neptune.DeleteDBClusterInput) (*neptune.DeleteDBClusterOutput, error) {
	return a.client.DeleteDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBClusterParameterGroup(ctx context.Context, input *neptune.DeleteDBClusterParameterGroupInput) (*neptune.DeleteDBClusterParameterGroupOutput, error) {
	return a.client.DeleteDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBClusterSnapshot(ctx context.Context, input *neptune.DeleteDBClusterSnapshotInput) (*neptune.DeleteDBClusterSnapshotOutput, error) {
	return a.client.DeleteDBClusterSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBInstance(ctx context.Context, input *neptune.DeleteDBInstanceInput) (*neptune.DeleteDBInstanceOutput, error) {
	return a.client.DeleteDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBParameterGroup(ctx context.Context, input *neptune.DeleteDBParameterGroupInput) (*neptune.DeleteDBParameterGroupOutput, error) {
	return a.client.DeleteDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteDBSubnetGroup(ctx context.Context, input *neptune.DeleteDBSubnetGroupInput) (*neptune.DeleteDBSubnetGroupOutput, error) {
	return a.client.DeleteDBSubnetGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) DeleteEventSubscription(ctx context.Context, input *neptune.DeleteEventSubscriptionInput) (*neptune.DeleteEventSubscriptionOutput, error) {
	return a.client.DeleteEventSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterParameterGroups(ctx context.Context, input *neptune.DescribeDBClusterParameterGroupsInput) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
	return a.client.DescribeDBClusterParameterGroupsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterParameters(ctx context.Context, input *neptune.DescribeDBClusterParametersInput) (*neptune.DescribeDBClusterParametersOutput, error) {
	return a.client.DescribeDBClusterParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterSnapshotAttributes(ctx context.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
	return a.client.DescribeDBClusterSnapshotAttributesWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusterSnapshots(ctx context.Context, input *neptune.DescribeDBClusterSnapshotsInput) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
	return a.client.DescribeDBClusterSnapshotsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBClusters(ctx context.Context, input *neptune.DescribeDBClustersInput) (*neptune.DescribeDBClustersOutput, error) {
	return a.client.DescribeDBClustersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBEngineVersions(ctx context.Context, input *neptune.DescribeDBEngineVersionsInput) (*neptune.DescribeDBEngineVersionsOutput, error) {
	return a.client.DescribeDBEngineVersionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBInstances(ctx context.Context, input *neptune.DescribeDBInstancesInput) (*neptune.DescribeDBInstancesOutput, error) {
	return a.client.DescribeDBInstancesWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBParameterGroups(ctx context.Context, input *neptune.DescribeDBParameterGroupsInput) (*neptune.DescribeDBParameterGroupsOutput, error) {
	return a.client.DescribeDBParameterGroupsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBParameters(ctx context.Context, input *neptune.DescribeDBParametersInput) (*neptune.DescribeDBParametersOutput, error) {
	return a.client.DescribeDBParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeDBSubnetGroups(ctx context.Context, input *neptune.DescribeDBSubnetGroupsInput) (*neptune.DescribeDBSubnetGroupsOutput, error) {
	return a.client.DescribeDBSubnetGroupsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEngineDefaultClusterParameters(ctx context.Context, input *neptune.DescribeEngineDefaultClusterParametersInput) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
	return a.client.DescribeEngineDefaultClusterParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEngineDefaultParameters(ctx context.Context, input *neptune.DescribeEngineDefaultParametersInput) (*neptune.DescribeEngineDefaultParametersOutput, error) {
	return a.client.DescribeEngineDefaultParametersWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEventCategories(ctx context.Context, input *neptune.DescribeEventCategoriesInput) (*neptune.DescribeEventCategoriesOutput, error) {
	return a.client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEventSubscriptions(ctx context.Context, input *neptune.DescribeEventSubscriptionsInput) (*neptune.DescribeEventSubscriptionsOutput, error) {
	return a.client.DescribeEventSubscriptionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeEvents(ctx context.Context, input *neptune.DescribeEventsInput) (*neptune.DescribeEventsOutput, error) {
	return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeOrderableDBInstanceOptions(ctx context.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
	return a.client.DescribeOrderableDBInstanceOptionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribePendingMaintenanceActions(ctx context.Context, input *neptune.DescribePendingMaintenanceActionsInput) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
	return a.client.DescribePendingMaintenanceActionsWithContext(ctx, input)
}

func (a *NeptuneActivities) DescribeValidDBInstanceModifications(ctx context.Context, input *neptune.DescribeValidDBInstanceModificationsInput) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
	return a.client.DescribeValidDBInstanceModificationsWithContext(ctx, input)
}

func (a *NeptuneActivities) FailoverDBCluster(ctx context.Context, input *neptune.FailoverDBClusterInput) (*neptune.FailoverDBClusterOutput, error) {
	return a.client.FailoverDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) ListTagsForResource(ctx context.Context, input *neptune.ListTagsForResourceInput) (*neptune.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBCluster(ctx context.Context, input *neptune.ModifyDBClusterInput) (*neptune.ModifyDBClusterOutput, error) {
	return a.client.ModifyDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBClusterParameterGroup(ctx context.Context, input *neptune.ModifyDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
	return a.client.ModifyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBClusterSnapshotAttribute(ctx context.Context, input *neptune.ModifyDBClusterSnapshotAttributeInput) (*neptune.ModifyDBClusterSnapshotAttributeOutput, error) {
	return a.client.ModifyDBClusterSnapshotAttributeWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBInstance(ctx context.Context, input *neptune.ModifyDBInstanceInput) (*neptune.ModifyDBInstanceOutput, error) {
	return a.client.ModifyDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBParameterGroup(ctx context.Context, input *neptune.ModifyDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
	return a.client.ModifyDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyDBSubnetGroup(ctx context.Context, input *neptune.ModifyDBSubnetGroupInput) (*neptune.ModifyDBSubnetGroupOutput, error) {
	return a.client.ModifyDBSubnetGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ModifyEventSubscription(ctx context.Context, input *neptune.ModifyEventSubscriptionInput) (*neptune.ModifyEventSubscriptionOutput, error) {
	return a.client.ModifyEventSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) PromoteReadReplicaDBCluster(ctx context.Context, input *neptune.PromoteReadReplicaDBClusterInput) (*neptune.PromoteReadReplicaDBClusterOutput, error) {
	return a.client.PromoteReadReplicaDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) RebootDBInstance(ctx context.Context, input *neptune.RebootDBInstanceInput) (*neptune.RebootDBInstanceOutput, error) {
	return a.client.RebootDBInstanceWithContext(ctx, input)
}

func (a *NeptuneActivities) RemoveRoleFromDBCluster(ctx context.Context, input *neptune.RemoveRoleFromDBClusterInput) (*neptune.RemoveRoleFromDBClusterOutput, error) {
	return a.client.RemoveRoleFromDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) RemoveSourceIdentifierFromSubscription(ctx context.Context, input *neptune.RemoveSourceIdentifierFromSubscriptionInput) (*neptune.RemoveSourceIdentifierFromSubscriptionOutput, error) {
	return a.client.RemoveSourceIdentifierFromSubscriptionWithContext(ctx, input)
}

func (a *NeptuneActivities) RemoveTagsFromResource(ctx context.Context, input *neptune.RemoveTagsFromResourceInput) (*neptune.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *NeptuneActivities) ResetDBClusterParameterGroup(ctx context.Context, input *neptune.ResetDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
	return a.client.ResetDBClusterParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) ResetDBParameterGroup(ctx context.Context, input *neptune.ResetDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
	return a.client.ResetDBParameterGroupWithContext(ctx, input)
}

func (a *NeptuneActivities) RestoreDBClusterFromSnapshot(ctx context.Context, input *neptune.RestoreDBClusterFromSnapshotInput) (*neptune.RestoreDBClusterFromSnapshotOutput, error) {
	return a.client.RestoreDBClusterFromSnapshotWithContext(ctx, input)
}

func (a *NeptuneActivities) RestoreDBClusterToPointInTime(ctx context.Context, input *neptune.RestoreDBClusterToPointInTimeInput) (*neptune.RestoreDBClusterToPointInTimeOutput, error) {
	return a.client.RestoreDBClusterToPointInTimeWithContext(ctx, input)
}

func (a *NeptuneActivities) StartDBCluster(ctx context.Context, input *neptune.StartDBClusterInput) (*neptune.StartDBClusterOutput, error) {
	return a.client.StartDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) StopDBCluster(ctx context.Context, input *neptune.StopDBClusterInput) (*neptune.StopDBClusterOutput, error) {
	return a.client.StopDBClusterWithContext(ctx, input)
}

func (a *NeptuneActivities) WaitUntilDBInstanceAvailable(ctx context.Context, input *neptune.DescribeDBInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBInstanceAvailableWithContext(ctx, input, options...)
	})
}

func (a *NeptuneActivities) WaitUntilDBInstanceDeleted(ctx context.Context, input *neptune.DescribeDBInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBInstanceDeletedWithContext(ctx, input, options...)
	})
}
