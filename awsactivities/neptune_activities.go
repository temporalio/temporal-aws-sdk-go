package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/neptune/neptuneiface"
)

type NeptuneActivities struct {
	client neptuneiface.NeptuneAPI
}

func NewNeptuneActivities(client neptuneiface.NeptuneAPI) *NeptuneActivities {
    return &NeptuneActivities{client: client}
}

func (a *NeptuneActivities) AddRoleToDBCluster(input *neptune.AddRoleToDBClusterInput) (*neptune.AddRoleToDBClusterOutput, error) {
    return a.client.AddRoleToDBCluster(input)
}

func (a *NeptuneActivities) AddSourceIdentifierToSubscription(input *neptune.AddSourceIdentifierToSubscriptionInput) (*neptune.AddSourceIdentifierToSubscriptionOutput, error) {
    return a.client.AddSourceIdentifierToSubscription(input)
}

func (a *NeptuneActivities) AddTagsToResource(input *neptune.AddTagsToResourceInput) (*neptune.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResource(input)
}

func (a *NeptuneActivities) ApplyPendingMaintenanceAction(input *neptune.ApplyPendingMaintenanceActionInput) (*neptune.ApplyPendingMaintenanceActionOutput, error) {
    return a.client.ApplyPendingMaintenanceAction(input)
}

func (a *NeptuneActivities) CopyDBClusterParameterGroup(input *neptune.CopyDBClusterParameterGroupInput) (*neptune.CopyDBClusterParameterGroupOutput, error) {
    return a.client.CopyDBClusterParameterGroup(input)
}

func (a *NeptuneActivities) CopyDBClusterSnapshot(input *neptune.CopyDBClusterSnapshotInput) (*neptune.CopyDBClusterSnapshotOutput, error) {
    return a.client.CopyDBClusterSnapshot(input)
}

func (a *NeptuneActivities) CopyDBParameterGroup(input *neptune.CopyDBParameterGroupInput) (*neptune.CopyDBParameterGroupOutput, error) {
    return a.client.CopyDBParameterGroup(input)
}

func (a *NeptuneActivities) CreateDBCluster(input *neptune.CreateDBClusterInput) (*neptune.CreateDBClusterOutput, error) {
    return a.client.CreateDBCluster(input)
}

func (a *NeptuneActivities) CreateDBClusterParameterGroup(input *neptune.CreateDBClusterParameterGroupInput) (*neptune.CreateDBClusterParameterGroupOutput, error) {
    return a.client.CreateDBClusterParameterGroup(input)
}

func (a *NeptuneActivities) CreateDBClusterSnapshot(input *neptune.CreateDBClusterSnapshotInput) (*neptune.CreateDBClusterSnapshotOutput, error) {
    return a.client.CreateDBClusterSnapshot(input)
}

func (a *NeptuneActivities) CreateDBInstance(input *neptune.CreateDBInstanceInput) (*neptune.CreateDBInstanceOutput, error) {
    return a.client.CreateDBInstance(input)
}

func (a *NeptuneActivities) CreateDBParameterGroup(input *neptune.CreateDBParameterGroupInput) (*neptune.CreateDBParameterGroupOutput, error) {
    return a.client.CreateDBParameterGroup(input)
}

func (a *NeptuneActivities) CreateDBSubnetGroup(input *neptune.CreateDBSubnetGroupInput) (*neptune.CreateDBSubnetGroupOutput, error) {
    return a.client.CreateDBSubnetGroup(input)
}

func (a *NeptuneActivities) CreateEventSubscription(input *neptune.CreateEventSubscriptionInput) (*neptune.CreateEventSubscriptionOutput, error) {
    return a.client.CreateEventSubscription(input)
}

func (a *NeptuneActivities) DeleteDBCluster(input *neptune.DeleteDBClusterInput) (*neptune.DeleteDBClusterOutput, error) {
    return a.client.DeleteDBCluster(input)
}

func (a *NeptuneActivities) DeleteDBClusterParameterGroup(input *neptune.DeleteDBClusterParameterGroupInput) (*neptune.DeleteDBClusterParameterGroupOutput, error) {
    return a.client.DeleteDBClusterParameterGroup(input)
}

func (a *NeptuneActivities) DeleteDBClusterSnapshot(input *neptune.DeleteDBClusterSnapshotInput) (*neptune.DeleteDBClusterSnapshotOutput, error) {
    return a.client.DeleteDBClusterSnapshot(input)
}

func (a *NeptuneActivities) DeleteDBInstance(input *neptune.DeleteDBInstanceInput) (*neptune.DeleteDBInstanceOutput, error) {
    return a.client.DeleteDBInstance(input)
}

func (a *NeptuneActivities) DeleteDBParameterGroup(input *neptune.DeleteDBParameterGroupInput) (*neptune.DeleteDBParameterGroupOutput, error) {
    return a.client.DeleteDBParameterGroup(input)
}

func (a *NeptuneActivities) DeleteDBSubnetGroup(input *neptune.DeleteDBSubnetGroupInput) (*neptune.DeleteDBSubnetGroupOutput, error) {
    return a.client.DeleteDBSubnetGroup(input)
}

func (a *NeptuneActivities) DeleteEventSubscription(input *neptune.DeleteEventSubscriptionInput) (*neptune.DeleteEventSubscriptionOutput, error) {
    return a.client.DeleteEventSubscription(input)
}

func (a *NeptuneActivities) DescribeDBClusterParameterGroups(input *neptune.DescribeDBClusterParameterGroupsInput) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
    return a.client.DescribeDBClusterParameterGroups(input)
}

func (a *NeptuneActivities) DescribeDBClusterParameters(input *neptune.DescribeDBClusterParametersInput) (*neptune.DescribeDBClusterParametersOutput, error) {
    return a.client.DescribeDBClusterParameters(input)
}

func (a *NeptuneActivities) DescribeDBClusterSnapshotAttributes(input *neptune.DescribeDBClusterSnapshotAttributesInput) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
    return a.client.DescribeDBClusterSnapshotAttributes(input)
}

func (a *NeptuneActivities) DescribeDBClusterSnapshots(input *neptune.DescribeDBClusterSnapshotsInput) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
    return a.client.DescribeDBClusterSnapshots(input)
}

func (a *NeptuneActivities) DescribeDBClusters(input *neptune.DescribeDBClustersInput) (*neptune.DescribeDBClustersOutput, error) {
    return a.client.DescribeDBClusters(input)
}

func (a *NeptuneActivities) DescribeDBEngineVersions(input *neptune.DescribeDBEngineVersionsInput) (*neptune.DescribeDBEngineVersionsOutput, error) {
    return a.client.DescribeDBEngineVersions(input)
}

func (a *NeptuneActivities) DescribeDBInstances(input *neptune.DescribeDBInstancesInput) (*neptune.DescribeDBInstancesOutput, error) {
    return a.client.DescribeDBInstances(input)
}

func (a *NeptuneActivities) DescribeDBParameterGroups(input *neptune.DescribeDBParameterGroupsInput) (*neptune.DescribeDBParameterGroupsOutput, error) {
    return a.client.DescribeDBParameterGroups(input)
}

func (a *NeptuneActivities) DescribeDBParameters(input *neptune.DescribeDBParametersInput) (*neptune.DescribeDBParametersOutput, error) {
    return a.client.DescribeDBParameters(input)
}

func (a *NeptuneActivities) DescribeDBSubnetGroups(input *neptune.DescribeDBSubnetGroupsInput) (*neptune.DescribeDBSubnetGroupsOutput, error) {
    return a.client.DescribeDBSubnetGroups(input)
}

func (a *NeptuneActivities) DescribeEngineDefaultClusterParameters(input *neptune.DescribeEngineDefaultClusterParametersInput) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
    return a.client.DescribeEngineDefaultClusterParameters(input)
}

func (a *NeptuneActivities) DescribeEngineDefaultParameters(input *neptune.DescribeEngineDefaultParametersInput) (*neptune.DescribeEngineDefaultParametersOutput, error) {
    return a.client.DescribeEngineDefaultParameters(input)
}

func (a *NeptuneActivities) DescribeEventCategories(input *neptune.DescribeEventCategoriesInput) (*neptune.DescribeEventCategoriesOutput, error) {
    return a.client.DescribeEventCategories(input)
}

func (a *NeptuneActivities) DescribeEventSubscriptions(input *neptune.DescribeEventSubscriptionsInput) (*neptune.DescribeEventSubscriptionsOutput, error) {
    return a.client.DescribeEventSubscriptions(input)
}

func (a *NeptuneActivities) DescribeEvents(input *neptune.DescribeEventsInput) (*neptune.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}

func (a *NeptuneActivities) DescribeOrderableDBInstanceOptions(input *neptune.DescribeOrderableDBInstanceOptionsInput) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
    return a.client.DescribeOrderableDBInstanceOptions(input)
}

func (a *NeptuneActivities) DescribePendingMaintenanceActions(input *neptune.DescribePendingMaintenanceActionsInput) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
    return a.client.DescribePendingMaintenanceActions(input)
}

func (a *NeptuneActivities) DescribeValidDBInstanceModifications(input *neptune.DescribeValidDBInstanceModificationsInput) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
    return a.client.DescribeValidDBInstanceModifications(input)
}

func (a *NeptuneActivities) FailoverDBCluster(input *neptune.FailoverDBClusterInput) (*neptune.FailoverDBClusterOutput, error) {
    return a.client.FailoverDBCluster(input)
}

func (a *NeptuneActivities) ListTagsForResource(input *neptune.ListTagsForResourceInput) (*neptune.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *NeptuneActivities) ModifyDBCluster(input *neptune.ModifyDBClusterInput) (*neptune.ModifyDBClusterOutput, error) {
    return a.client.ModifyDBCluster(input)
}

func (a *NeptuneActivities) ModifyDBClusterParameterGroup(input *neptune.ModifyDBClusterParameterGroupInput) (*neptune.ModifyDBClusterParameterGroupOutput, error) {
    return a.client.ModifyDBClusterParameterGroup(input)
}

func (a *NeptuneActivities) ModifyDBClusterSnapshotAttribute(input *neptune.ModifyDBClusterSnapshotAttributeInput) (*neptune.ModifyDBClusterSnapshotAttributeOutput, error) {
    return a.client.ModifyDBClusterSnapshotAttribute(input)
}

func (a *NeptuneActivities) ModifyDBInstance(input *neptune.ModifyDBInstanceInput) (*neptune.ModifyDBInstanceOutput, error) {
    return a.client.ModifyDBInstance(input)
}

func (a *NeptuneActivities) ModifyDBParameterGroup(input *neptune.ModifyDBParameterGroupInput) (*neptune.ModifyDBParameterGroupOutput, error) {
    return a.client.ModifyDBParameterGroup(input)
}

func (a *NeptuneActivities) ModifyDBSubnetGroup(input *neptune.ModifyDBSubnetGroupInput) (*neptune.ModifyDBSubnetGroupOutput, error) {
    return a.client.ModifyDBSubnetGroup(input)
}

func (a *NeptuneActivities) ModifyEventSubscription(input *neptune.ModifyEventSubscriptionInput) (*neptune.ModifyEventSubscriptionOutput, error) {
    return a.client.ModifyEventSubscription(input)
}

func (a *NeptuneActivities) PromoteReadReplicaDBCluster(input *neptune.PromoteReadReplicaDBClusterInput) (*neptune.PromoteReadReplicaDBClusterOutput, error) {
    return a.client.PromoteReadReplicaDBCluster(input)
}

func (a *NeptuneActivities) RebootDBInstance(input *neptune.RebootDBInstanceInput) (*neptune.RebootDBInstanceOutput, error) {
    return a.client.RebootDBInstance(input)
}

func (a *NeptuneActivities) RemoveRoleFromDBCluster(input *neptune.RemoveRoleFromDBClusterInput) (*neptune.RemoveRoleFromDBClusterOutput, error) {
    return a.client.RemoveRoleFromDBCluster(input)
}

func (a *NeptuneActivities) RemoveSourceIdentifierFromSubscription(input *neptune.RemoveSourceIdentifierFromSubscriptionInput) (*neptune.RemoveSourceIdentifierFromSubscriptionOutput, error) {
    return a.client.RemoveSourceIdentifierFromSubscription(input)
}

func (a *NeptuneActivities) RemoveTagsFromResource(input *neptune.RemoveTagsFromResourceInput) (*neptune.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResource(input)
}

func (a *NeptuneActivities) ResetDBClusterParameterGroup(input *neptune.ResetDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
    return a.client.ResetDBClusterParameterGroup(input)
}

func (a *NeptuneActivities) ResetDBParameterGroup(input *neptune.ResetDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
    return a.client.ResetDBParameterGroup(input)
}

func (a *NeptuneActivities) RestoreDBClusterFromSnapshot(input *neptune.RestoreDBClusterFromSnapshotInput) (*neptune.RestoreDBClusterFromSnapshotOutput, error) {
    return a.client.RestoreDBClusterFromSnapshot(input)
}

func (a *NeptuneActivities) RestoreDBClusterToPointInTime(input *neptune.RestoreDBClusterToPointInTimeInput) (*neptune.RestoreDBClusterToPointInTimeOutput, error) {
    return a.client.RestoreDBClusterToPointInTime(input)
}

func (a *NeptuneActivities) StartDBCluster(input *neptune.StartDBClusterInput) (*neptune.StartDBClusterOutput, error) {
    return a.client.StartDBCluster(input)
}

func (a *NeptuneActivities) StopDBCluster(input *neptune.StopDBClusterInput) (*neptune.StopDBClusterOutput, error) {
    return a.client.StopDBCluster(input)
}

func (a *NeptuneActivities) WaitUntilDBInstanceAvailable(input *neptune.WaitUntilDBInstanceAvailableInput) (*neptune.WaitUntilDBInstanceAvailableOutput, error) {
    return a.client.WaitUntilDBInstanceAvailable(input)
}

func (a *NeptuneActivities) WaitUntilDBInstanceDeleted(input *neptune.WaitUntilDBInstanceDeletedInput) (*neptune.WaitUntilDBInstanceDeletedOutput, error) {
    return a.client.WaitUntilDBInstanceDeleted(input)
}
