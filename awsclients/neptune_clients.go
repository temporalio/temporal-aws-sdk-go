package awsclients

import (
	"github.com/aws/aws-sdk-go/service/neptune"
	"go.temporal.io/sdk/workflow"
)

type NeptuneClient interface {
       AddRoleToDBCluster(ctx workflow.Context, input *neptune.AddRoleToDBClusterInput) (*neptune.AddRoleToDBClusterOutput, error)
       AddRoleToDBClusterAsync(ctx workflow.Context, input *neptune.AddRoleToDBClusterInput) *NeptuneAddRoleToDBClusterResult

       AddSourceIdentifierToSubscription(ctx workflow.Context, input *neptune.AddSourceIdentifierToSubscriptionInput) (*neptune.AddSourceIdentifierToSubscriptionOutput, error)
       AddSourceIdentifierToSubscriptionAsync(ctx workflow.Context, input *neptune.AddSourceIdentifierToSubscriptionInput) *NeptuneAddSourceIdentifierToSubscriptionResult

       AddTagsToResource(ctx workflow.Context, input *neptune.AddTagsToResourceInput) (*neptune.AddTagsToResourceOutput, error)
       AddTagsToResourceAsync(ctx workflow.Context, input *neptune.AddTagsToResourceInput) *NeptuneAddTagsToResourceResult

       ApplyPendingMaintenanceAction(ctx workflow.Context, input *neptune.ApplyPendingMaintenanceActionInput) (*neptune.ApplyPendingMaintenanceActionOutput, error)
       ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *neptune.ApplyPendingMaintenanceActionInput) *NeptuneApplyPendingMaintenanceActionResult

       CopyDBClusterParameterGroup(ctx workflow.Context, input *neptune.CopyDBClusterParameterGroupInput) (*neptune.CopyDBClusterParameterGroupOutput, error)
       CopyDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.CopyDBClusterParameterGroupInput) *NeptuneCopyDBClusterParameterGroupResult

       CopyDBClusterSnapshot(ctx workflow.Context, input *neptune.CopyDBClusterSnapshotInput) (*neptune.CopyDBClusterSnapshotOutput, error)
       CopyDBClusterSnapshotAsync(ctx workflow.Context, input *neptune.CopyDBClusterSnapshotInput) *NeptuneCopyDBClusterSnapshotResult

       CopyDBParameterGroup(ctx workflow.Context, input *neptune.CopyDBParameterGroupInput) (*neptune.CopyDBParameterGroupOutput, error)
       CopyDBParameterGroupAsync(ctx workflow.Context, input *neptune.CopyDBParameterGroupInput) *NeptuneCopyDBParameterGroupResult

       CreateDBCluster(ctx workflow.Context, input *neptune.CreateDBClusterInput) (*neptune.CreateDBClusterOutput, error)
       CreateDBClusterAsync(ctx workflow.Context, input *neptune.CreateDBClusterInput) *NeptuneCreateDBClusterResult

       CreateDBClusterParameterGroup(ctx workflow.Context, input *neptune.CreateDBClusterParameterGroupInput) (*neptune.CreateDBClusterParameterGroupOutput, error)
       CreateDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.CreateDBClusterParameterGroupInput) *NeptuneCreateDBClusterParameterGroupResult

       CreateDBClusterSnapshot(ctx workflow.Context, input *neptune.CreateDBClusterSnapshotInput) (*neptune.CreateDBClusterSnapshotOutput, error)
       CreateDBClusterSnapshotAsync(ctx workflow.Context, input *neptune.CreateDBClusterSnapshotInput) *NeptuneCreateDBClusterSnapshotResult

       CreateDBInstance(ctx workflow.Context, input *neptune.CreateDBInstanceInput) (*neptune.CreateDBInstanceOutput, error)
       CreateDBInstanceAsync(ctx workflow.Context, input *neptune.CreateDBInstanceInput) *NeptuneCreateDBInstanceResult

       CreateDBParameterGroup(ctx workflow.Context, input *neptune.CreateDBParameterGroupInput) (*neptune.CreateDBParameterGroupOutput, error)
       CreateDBParameterGroupAsync(ctx workflow.Context, input *neptune.CreateDBParameterGroupInput) *NeptuneCreateDBParameterGroupResult

       CreateDBSubnetGroup(ctx workflow.Context, input *neptune.CreateDBSubnetGroupInput) (*neptune.CreateDBSubnetGroupOutput, error)
       CreateDBSubnetGroupAsync(ctx workflow.Context, input *neptune.CreateDBSubnetGroupInput) *NeptuneCreateDBSubnetGroupResult

       CreateEventSubscription(ctx workflow.Context, input *neptune.CreateEventSubscriptionInput) (*neptune.CreateEventSubscriptionOutput, error)
       CreateEventSubscriptionAsync(ctx workflow.Context, input *neptune.CreateEventSubscriptionInput) *NeptuneCreateEventSubscriptionResult

       DeleteDBCluster(ctx workflow.Context, input *neptune.DeleteDBClusterInput) (*neptune.DeleteDBClusterOutput, error)
       DeleteDBClusterAsync(ctx workflow.Context, input *neptune.DeleteDBClusterInput) *NeptuneDeleteDBClusterResult

       DeleteDBClusterParameterGroup(ctx workflow.Context, input *neptune.DeleteDBClusterParameterGroupInput) (*neptune.DeleteDBClusterParameterGroupOutput, error)
       DeleteDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.DeleteDBClusterParameterGroupInput) *NeptuneDeleteDBClusterParameterGroupResult

       DeleteDBClusterSnapshot(ctx workflow.Context, input *neptune.DeleteDBClusterSnapshotInput) (*neptune.DeleteDBClusterSnapshotOutput, error)
       DeleteDBClusterSnapshotAsync(ctx workflow.Context, input *neptune.DeleteDBClusterSnapshotInput) *NeptuneDeleteDBClusterSnapshotResult

       DeleteDBInstance(ctx workflow.Context, input *neptune.DeleteDBInstanceInput) (*neptune.DeleteDBInstanceOutput, error)
       DeleteDBInstanceAsync(ctx workflow.Context, input *neptune.DeleteDBInstanceInput) *NeptuneDeleteDBInstanceResult

       DeleteDBParameterGroup(ctx workflow.Context, input *neptune.DeleteDBParameterGroupInput) (*neptune.DeleteDBParameterGroupOutput, error)
       DeleteDBParameterGroupAsync(ctx workflow.Context, input *neptune.DeleteDBParameterGroupInput) *NeptuneDeleteDBParameterGroupResult

       DeleteDBSubnetGroup(ctx workflow.Context, input *neptune.DeleteDBSubnetGroupInput) (*neptune.DeleteDBSubnetGroupOutput, error)
       DeleteDBSubnetGroupAsync(ctx workflow.Context, input *neptune.DeleteDBSubnetGroupInput) *NeptuneDeleteDBSubnetGroupResult

       DeleteEventSubscription(ctx workflow.Context, input *neptune.DeleteEventSubscriptionInput) (*neptune.DeleteEventSubscriptionOutput, error)
       DeleteEventSubscriptionAsync(ctx workflow.Context, input *neptune.DeleteEventSubscriptionInput) *NeptuneDeleteEventSubscriptionResult

       DescribeDBClusterParameterGroups(ctx workflow.Context, input *neptune.DescribeDBClusterParameterGroupsInput) (*neptune.DescribeDBClusterParameterGroupsOutput, error)
       DescribeDBClusterParameterGroupsAsync(ctx workflow.Context, input *neptune.DescribeDBClusterParameterGroupsInput) *NeptuneDescribeDBClusterParameterGroupsResult

       DescribeDBClusterParameters(ctx workflow.Context, input *neptune.DescribeDBClusterParametersInput) (*neptune.DescribeDBClusterParametersOutput, error)
       DescribeDBClusterParametersAsync(ctx workflow.Context, input *neptune.DescribeDBClusterParametersInput) *NeptuneDescribeDBClusterParametersResult

       DescribeDBClusterSnapshotAttributes(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error)
       DescribeDBClusterSnapshotAttributesAsync(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput) *NeptuneDescribeDBClusterSnapshotAttributesResult

       DescribeDBClusterSnapshots(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotsInput) (*neptune.DescribeDBClusterSnapshotsOutput, error)
       DescribeDBClusterSnapshotsAsync(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotsInput) *NeptuneDescribeDBClusterSnapshotsResult

       DescribeDBClusters(ctx workflow.Context, input *neptune.DescribeDBClustersInput) (*neptune.DescribeDBClustersOutput, error)
       DescribeDBClustersAsync(ctx workflow.Context, input *neptune.DescribeDBClustersInput) *NeptuneDescribeDBClustersResult

       DescribeDBEngineVersions(ctx workflow.Context, input *neptune.DescribeDBEngineVersionsInput) (*neptune.DescribeDBEngineVersionsOutput, error)
       DescribeDBEngineVersionsAsync(ctx workflow.Context, input *neptune.DescribeDBEngineVersionsInput) *NeptuneDescribeDBEngineVersionsResult

       DescribeDBInstances(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) (*neptune.DescribeDBInstancesOutput, error)
       DescribeDBInstancesAsync(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) *NeptuneDescribeDBInstancesResult

       DescribeDBParameterGroups(ctx workflow.Context, input *neptune.DescribeDBParameterGroupsInput) (*neptune.DescribeDBParameterGroupsOutput, error)
       DescribeDBParameterGroupsAsync(ctx workflow.Context, input *neptune.DescribeDBParameterGroupsInput) *NeptuneDescribeDBParameterGroupsResult

       DescribeDBParameters(ctx workflow.Context, input *neptune.DescribeDBParametersInput) (*neptune.DescribeDBParametersOutput, error)
       DescribeDBParametersAsync(ctx workflow.Context, input *neptune.DescribeDBParametersInput) *NeptuneDescribeDBParametersResult

       DescribeDBSubnetGroups(ctx workflow.Context, input *neptune.DescribeDBSubnetGroupsInput) (*neptune.DescribeDBSubnetGroupsOutput, error)
       DescribeDBSubnetGroupsAsync(ctx workflow.Context, input *neptune.DescribeDBSubnetGroupsInput) *NeptuneDescribeDBSubnetGroupsResult

       DescribeEngineDefaultClusterParameters(ctx workflow.Context, input *neptune.DescribeEngineDefaultClusterParametersInput) (*neptune.DescribeEngineDefaultClusterParametersOutput, error)
       DescribeEngineDefaultClusterParametersAsync(ctx workflow.Context, input *neptune.DescribeEngineDefaultClusterParametersInput) *NeptuneDescribeEngineDefaultClusterParametersResult

       DescribeEngineDefaultParameters(ctx workflow.Context, input *neptune.DescribeEngineDefaultParametersInput) (*neptune.DescribeEngineDefaultParametersOutput, error)
       DescribeEngineDefaultParametersAsync(ctx workflow.Context, input *neptune.DescribeEngineDefaultParametersInput) *NeptuneDescribeEngineDefaultParametersResult

       DescribeEventCategories(ctx workflow.Context, input *neptune.DescribeEventCategoriesInput) (*neptune.DescribeEventCategoriesOutput, error)
       DescribeEventCategoriesAsync(ctx workflow.Context, input *neptune.DescribeEventCategoriesInput) *NeptuneDescribeEventCategoriesResult

       DescribeEventSubscriptions(ctx workflow.Context, input *neptune.DescribeEventSubscriptionsInput) (*neptune.DescribeEventSubscriptionsOutput, error)
       DescribeEventSubscriptionsAsync(ctx workflow.Context, input *neptune.DescribeEventSubscriptionsInput) *NeptuneDescribeEventSubscriptionsResult

       DescribeEvents(ctx workflow.Context, input *neptune.DescribeEventsInput) (*neptune.DescribeEventsOutput, error)
       DescribeEventsAsync(ctx workflow.Context, input *neptune.DescribeEventsInput) *NeptuneDescribeEventsResult

       DescribeOrderableDBInstanceOptions(ctx workflow.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error)
       DescribeOrderableDBInstanceOptionsAsync(ctx workflow.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput) *NeptuneDescribeOrderableDBInstanceOptionsResult

       DescribePendingMaintenanceActions(ctx workflow.Context, input *neptune.DescribePendingMaintenanceActionsInput) (*neptune.DescribePendingMaintenanceActionsOutput, error)
       DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *neptune.DescribePendingMaintenanceActionsInput) *NeptuneDescribePendingMaintenanceActionsResult

       DescribeValidDBInstanceModifications(ctx workflow.Context, input *neptune.DescribeValidDBInstanceModificationsInput) (*neptune.DescribeValidDBInstanceModificationsOutput, error)
       DescribeValidDBInstanceModificationsAsync(ctx workflow.Context, input *neptune.DescribeValidDBInstanceModificationsInput) *NeptuneDescribeValidDBInstanceModificationsResult

       FailoverDBCluster(ctx workflow.Context, input *neptune.FailoverDBClusterInput) (*neptune.FailoverDBClusterOutput, error)
       FailoverDBClusterAsync(ctx workflow.Context, input *neptune.FailoverDBClusterInput) *NeptuneFailoverDBClusterResult

       ListTagsForResource(ctx workflow.Context, input *neptune.ListTagsForResourceInput) (*neptune.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *neptune.ListTagsForResourceInput) *NeptuneListTagsForResourceResult

       ModifyDBCluster(ctx workflow.Context, input *neptune.ModifyDBClusterInput) (*neptune.ModifyDBClusterOutput, error)
       ModifyDBClusterAsync(ctx workflow.Context, input *neptune.ModifyDBClusterInput) *NeptuneModifyDBClusterResult

       ModifyDBClusterParameterGroup(ctx workflow.Context, input *neptune.ModifyDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error)
       ModifyDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.ModifyDBClusterParameterGroupInput) *NeptuneModifyDBClusterParameterGroupResult

       ModifyDBClusterSnapshotAttribute(ctx workflow.Context, input *neptune.ModifyDBClusterSnapshotAttributeInput) (*neptune.ModifyDBClusterSnapshotAttributeOutput, error)
       ModifyDBClusterSnapshotAttributeAsync(ctx workflow.Context, input *neptune.ModifyDBClusterSnapshotAttributeInput) *NeptuneModifyDBClusterSnapshotAttributeResult

       ModifyDBInstance(ctx workflow.Context, input *neptune.ModifyDBInstanceInput) (*neptune.ModifyDBInstanceOutput, error)
       ModifyDBInstanceAsync(ctx workflow.Context, input *neptune.ModifyDBInstanceInput) *NeptuneModifyDBInstanceResult

       ModifyDBParameterGroup(ctx workflow.Context, input *neptune.ModifyDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error)
       ModifyDBParameterGroupAsync(ctx workflow.Context, input *neptune.ModifyDBParameterGroupInput) *NeptuneModifyDBParameterGroupResult

       ModifyDBSubnetGroup(ctx workflow.Context, input *neptune.ModifyDBSubnetGroupInput) (*neptune.ModifyDBSubnetGroupOutput, error)
       ModifyDBSubnetGroupAsync(ctx workflow.Context, input *neptune.ModifyDBSubnetGroupInput) *NeptuneModifyDBSubnetGroupResult

       ModifyEventSubscription(ctx workflow.Context, input *neptune.ModifyEventSubscriptionInput) (*neptune.ModifyEventSubscriptionOutput, error)
       ModifyEventSubscriptionAsync(ctx workflow.Context, input *neptune.ModifyEventSubscriptionInput) *NeptuneModifyEventSubscriptionResult

       PromoteReadReplicaDBCluster(ctx workflow.Context, input *neptune.PromoteReadReplicaDBClusterInput) (*neptune.PromoteReadReplicaDBClusterOutput, error)
       PromoteReadReplicaDBClusterAsync(ctx workflow.Context, input *neptune.PromoteReadReplicaDBClusterInput) *NeptunePromoteReadReplicaDBClusterResult

       RebootDBInstance(ctx workflow.Context, input *neptune.RebootDBInstanceInput) (*neptune.RebootDBInstanceOutput, error)
       RebootDBInstanceAsync(ctx workflow.Context, input *neptune.RebootDBInstanceInput) *NeptuneRebootDBInstanceResult

       RemoveRoleFromDBCluster(ctx workflow.Context, input *neptune.RemoveRoleFromDBClusterInput) (*neptune.RemoveRoleFromDBClusterOutput, error)
       RemoveRoleFromDBClusterAsync(ctx workflow.Context, input *neptune.RemoveRoleFromDBClusterInput) *NeptuneRemoveRoleFromDBClusterResult

       RemoveSourceIdentifierFromSubscription(ctx workflow.Context, input *neptune.RemoveSourceIdentifierFromSubscriptionInput) (*neptune.RemoveSourceIdentifierFromSubscriptionOutput, error)
       RemoveSourceIdentifierFromSubscriptionAsync(ctx workflow.Context, input *neptune.RemoveSourceIdentifierFromSubscriptionInput) *NeptuneRemoveSourceIdentifierFromSubscriptionResult

       RemoveTagsFromResource(ctx workflow.Context, input *neptune.RemoveTagsFromResourceInput) (*neptune.RemoveTagsFromResourceOutput, error)
       RemoveTagsFromResourceAsync(ctx workflow.Context, input *neptune.RemoveTagsFromResourceInput) *NeptuneRemoveTagsFromResourceResult

       ResetDBClusterParameterGroup(ctx workflow.Context, input *neptune.ResetDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error)
       ResetDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.ResetDBClusterParameterGroupInput) *NeptuneResetDBClusterParameterGroupResult

       ResetDBParameterGroup(ctx workflow.Context, input *neptune.ResetDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error)
       ResetDBParameterGroupAsync(ctx workflow.Context, input *neptune.ResetDBParameterGroupInput) *NeptuneResetDBParameterGroupResult

       RestoreDBClusterFromSnapshot(ctx workflow.Context, input *neptune.RestoreDBClusterFromSnapshotInput) (*neptune.RestoreDBClusterFromSnapshotOutput, error)
       RestoreDBClusterFromSnapshotAsync(ctx workflow.Context, input *neptune.RestoreDBClusterFromSnapshotInput) *NeptuneRestoreDBClusterFromSnapshotResult

       RestoreDBClusterToPointInTime(ctx workflow.Context, input *neptune.RestoreDBClusterToPointInTimeInput) (*neptune.RestoreDBClusterToPointInTimeOutput, error)
       RestoreDBClusterToPointInTimeAsync(ctx workflow.Context, input *neptune.RestoreDBClusterToPointInTimeInput) *NeptuneRestoreDBClusterToPointInTimeResult

       StartDBCluster(ctx workflow.Context, input *neptune.StartDBClusterInput) (*neptune.StartDBClusterOutput, error)
       StartDBClusterAsync(ctx workflow.Context, input *neptune.StartDBClusterInput) *NeptuneStartDBClusterResult

       StopDBCluster(ctx workflow.Context, input *neptune.StopDBClusterInput) (*neptune.StopDBClusterOutput, error)
       StopDBClusterAsync(ctx workflow.Context, input *neptune.StopDBClusterInput) *NeptuneStopDBClusterResult

       WaitUntilDBInstanceAvailable(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) error
       WaitUntilDBInstanceDeleted(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) error}

type NeptuneStub struct {
}

func NewNeptuneStub() NeptuneClient {
    return &NeptuneStub{}
}

type NeptuneAddRoleToDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneAddRoleToDBClusterResult) Get(ctx workflow.Context) (*neptune.AddRoleToDBClusterOutput, error) {
    var output neptune.AddRoleToDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneAddSourceIdentifierToSubscriptionResult struct {
	Result workflow.Future
}

func (r *NeptuneAddSourceIdentifierToSubscriptionResult) Get(ctx workflow.Context) (*neptune.AddSourceIdentifierToSubscriptionOutput, error) {
    var output neptune.AddSourceIdentifierToSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *NeptuneAddTagsToResourceResult) Get(ctx workflow.Context) (*neptune.AddTagsToResourceOutput, error) {
    var output neptune.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneApplyPendingMaintenanceActionResult struct {
	Result workflow.Future
}

func (r *NeptuneApplyPendingMaintenanceActionResult) Get(ctx workflow.Context) (*neptune.ApplyPendingMaintenanceActionOutput, error) {
    var output neptune.ApplyPendingMaintenanceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCopyDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneCopyDBClusterParameterGroupResult) Get(ctx workflow.Context) (*neptune.CopyDBClusterParameterGroupOutput, error) {
    var output neptune.CopyDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCopyDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *NeptuneCopyDBClusterSnapshotResult) Get(ctx workflow.Context) (*neptune.CopyDBClusterSnapshotOutput, error) {
    var output neptune.CopyDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCopyDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneCopyDBParameterGroupResult) Get(ctx workflow.Context) (*neptune.CopyDBParameterGroupOutput, error) {
    var output neptune.CopyDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCreateDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneCreateDBClusterResult) Get(ctx workflow.Context) (*neptune.CreateDBClusterOutput, error) {
    var output neptune.CreateDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCreateDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneCreateDBClusterParameterGroupResult) Get(ctx workflow.Context) (*neptune.CreateDBClusterParameterGroupOutput, error) {
    var output neptune.CreateDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCreateDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *NeptuneCreateDBClusterSnapshotResult) Get(ctx workflow.Context) (*neptune.CreateDBClusterSnapshotOutput, error) {
    var output neptune.CreateDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCreateDBInstanceResult struct {
	Result workflow.Future
}

func (r *NeptuneCreateDBInstanceResult) Get(ctx workflow.Context) (*neptune.CreateDBInstanceOutput, error) {
    var output neptune.CreateDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCreateDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneCreateDBParameterGroupResult) Get(ctx workflow.Context) (*neptune.CreateDBParameterGroupOutput, error) {
    var output neptune.CreateDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCreateDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneCreateDBSubnetGroupResult) Get(ctx workflow.Context) (*neptune.CreateDBSubnetGroupOutput, error) {
    var output neptune.CreateDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneCreateEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *NeptuneCreateEventSubscriptionResult) Get(ctx workflow.Context) (*neptune.CreateEventSubscriptionOutput, error) {
    var output neptune.CreateEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDeleteDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneDeleteDBClusterResult) Get(ctx workflow.Context) (*neptune.DeleteDBClusterOutput, error) {
    var output neptune.DeleteDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDeleteDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneDeleteDBClusterParameterGroupResult) Get(ctx workflow.Context) (*neptune.DeleteDBClusterParameterGroupOutput, error) {
    var output neptune.DeleteDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDeleteDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *NeptuneDeleteDBClusterSnapshotResult) Get(ctx workflow.Context) (*neptune.DeleteDBClusterSnapshotOutput, error) {
    var output neptune.DeleteDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDeleteDBInstanceResult struct {
	Result workflow.Future
}

func (r *NeptuneDeleteDBInstanceResult) Get(ctx workflow.Context) (*neptune.DeleteDBInstanceOutput, error) {
    var output neptune.DeleteDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDeleteDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneDeleteDBParameterGroupResult) Get(ctx workflow.Context) (*neptune.DeleteDBParameterGroupOutput, error) {
    var output neptune.DeleteDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDeleteDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneDeleteDBSubnetGroupResult) Get(ctx workflow.Context) (*neptune.DeleteDBSubnetGroupOutput, error) {
    var output neptune.DeleteDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDeleteEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *NeptuneDeleteEventSubscriptionResult) Get(ctx workflow.Context) (*neptune.DeleteEventSubscriptionOutput, error) {
    var output neptune.DeleteEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBClusterParameterGroupsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBClusterParameterGroupsResult) Get(ctx workflow.Context) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
    var output neptune.DescribeDBClusterParameterGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBClusterParametersResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBClusterParametersResult) Get(ctx workflow.Context) (*neptune.DescribeDBClusterParametersOutput, error) {
    var output neptune.DescribeDBClusterParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBClusterSnapshotAttributesResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBClusterSnapshotAttributesResult) Get(ctx workflow.Context) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
    var output neptune.DescribeDBClusterSnapshotAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBClusterSnapshotsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBClusterSnapshotsResult) Get(ctx workflow.Context) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
    var output neptune.DescribeDBClusterSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBClustersResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBClustersResult) Get(ctx workflow.Context) (*neptune.DescribeDBClustersOutput, error) {
    var output neptune.DescribeDBClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBEngineVersionsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBEngineVersionsResult) Get(ctx workflow.Context) (*neptune.DescribeDBEngineVersionsOutput, error) {
    var output neptune.DescribeDBEngineVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBInstancesResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBInstancesResult) Get(ctx workflow.Context) (*neptune.DescribeDBInstancesOutput, error) {
    var output neptune.DescribeDBInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBParameterGroupsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBParameterGroupsResult) Get(ctx workflow.Context) (*neptune.DescribeDBParameterGroupsOutput, error) {
    var output neptune.DescribeDBParameterGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBParametersResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBParametersResult) Get(ctx workflow.Context) (*neptune.DescribeDBParametersOutput, error) {
    var output neptune.DescribeDBParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeDBSubnetGroupsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeDBSubnetGroupsResult) Get(ctx workflow.Context) (*neptune.DescribeDBSubnetGroupsOutput, error) {
    var output neptune.DescribeDBSubnetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeEngineDefaultClusterParametersResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeEngineDefaultClusterParametersResult) Get(ctx workflow.Context) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
    var output neptune.DescribeEngineDefaultClusterParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeEngineDefaultParametersResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeEngineDefaultParametersResult) Get(ctx workflow.Context) (*neptune.DescribeEngineDefaultParametersOutput, error) {
    var output neptune.DescribeEngineDefaultParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeEventCategoriesResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeEventCategoriesResult) Get(ctx workflow.Context) (*neptune.DescribeEventCategoriesOutput, error) {
    var output neptune.DescribeEventCategoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeEventSubscriptionsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeEventSubscriptionsResult) Get(ctx workflow.Context) (*neptune.DescribeEventSubscriptionsOutput, error) {
    var output neptune.DescribeEventSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeEventsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeEventsResult) Get(ctx workflow.Context) (*neptune.DescribeEventsOutput, error) {
    var output neptune.DescribeEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeOrderableDBInstanceOptionsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeOrderableDBInstanceOptionsResult) Get(ctx workflow.Context) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
    var output neptune.DescribeOrderableDBInstanceOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribePendingMaintenanceActionsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribePendingMaintenanceActionsResult) Get(ctx workflow.Context) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
    var output neptune.DescribePendingMaintenanceActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneDescribeValidDBInstanceModificationsResult struct {
	Result workflow.Future
}

func (r *NeptuneDescribeValidDBInstanceModificationsResult) Get(ctx workflow.Context) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
    var output neptune.DescribeValidDBInstanceModificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneFailoverDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneFailoverDBClusterResult) Get(ctx workflow.Context) (*neptune.FailoverDBClusterOutput, error) {
    var output neptune.FailoverDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *NeptuneListTagsForResourceResult) Get(ctx workflow.Context) (*neptune.ListTagsForResourceOutput, error) {
    var output neptune.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneModifyDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneModifyDBClusterResult) Get(ctx workflow.Context) (*neptune.ModifyDBClusterOutput, error) {
    var output neptune.ModifyDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneModifyDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneModifyDBClusterParameterGroupResult) Get(ctx workflow.Context) (*neptune.ResetDBClusterParameterGroupOutput, error) {
    var output neptune.ResetDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneModifyDBClusterSnapshotAttributeResult struct {
	Result workflow.Future
}

func (r *NeptuneModifyDBClusterSnapshotAttributeResult) Get(ctx workflow.Context) (*neptune.ModifyDBClusterSnapshotAttributeOutput, error) {
    var output neptune.ModifyDBClusterSnapshotAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneModifyDBInstanceResult struct {
	Result workflow.Future
}

func (r *NeptuneModifyDBInstanceResult) Get(ctx workflow.Context) (*neptune.ModifyDBInstanceOutput, error) {
    var output neptune.ModifyDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneModifyDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneModifyDBParameterGroupResult) Get(ctx workflow.Context) (*neptune.ResetDBParameterGroupOutput, error) {
    var output neptune.ResetDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneModifyDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneModifyDBSubnetGroupResult) Get(ctx workflow.Context) (*neptune.ModifyDBSubnetGroupOutput, error) {
    var output neptune.ModifyDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneModifyEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *NeptuneModifyEventSubscriptionResult) Get(ctx workflow.Context) (*neptune.ModifyEventSubscriptionOutput, error) {
    var output neptune.ModifyEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptunePromoteReadReplicaDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptunePromoteReadReplicaDBClusterResult) Get(ctx workflow.Context) (*neptune.PromoteReadReplicaDBClusterOutput, error) {
    var output neptune.PromoteReadReplicaDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneRebootDBInstanceResult struct {
	Result workflow.Future
}

func (r *NeptuneRebootDBInstanceResult) Get(ctx workflow.Context) (*neptune.RebootDBInstanceOutput, error) {
    var output neptune.RebootDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneRemoveRoleFromDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneRemoveRoleFromDBClusterResult) Get(ctx workflow.Context) (*neptune.RemoveRoleFromDBClusterOutput, error) {
    var output neptune.RemoveRoleFromDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneRemoveSourceIdentifierFromSubscriptionResult struct {
	Result workflow.Future
}

func (r *NeptuneRemoveSourceIdentifierFromSubscriptionResult) Get(ctx workflow.Context) (*neptune.RemoveSourceIdentifierFromSubscriptionOutput, error) {
    var output neptune.RemoveSourceIdentifierFromSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *NeptuneRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*neptune.RemoveTagsFromResourceOutput, error) {
    var output neptune.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneResetDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneResetDBClusterParameterGroupResult) Get(ctx workflow.Context) (*neptune.ResetDBClusterParameterGroupOutput, error) {
    var output neptune.ResetDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneResetDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *NeptuneResetDBParameterGroupResult) Get(ctx workflow.Context) (*neptune.ResetDBParameterGroupOutput, error) {
    var output neptune.ResetDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneRestoreDBClusterFromSnapshotResult struct {
	Result workflow.Future
}

func (r *NeptuneRestoreDBClusterFromSnapshotResult) Get(ctx workflow.Context) (*neptune.RestoreDBClusterFromSnapshotOutput, error) {
    var output neptune.RestoreDBClusterFromSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneRestoreDBClusterToPointInTimeResult struct {
	Result workflow.Future
}

func (r *NeptuneRestoreDBClusterToPointInTimeResult) Get(ctx workflow.Context) (*neptune.RestoreDBClusterToPointInTimeOutput, error) {
    var output neptune.RestoreDBClusterToPointInTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneStartDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneStartDBClusterResult) Get(ctx workflow.Context) (*neptune.StartDBClusterOutput, error) {
    var output neptune.StartDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type NeptuneStopDBClusterResult struct {
	Result workflow.Future
}

func (r *NeptuneStopDBClusterResult) Get(ctx workflow.Context) (*neptune.StopDBClusterOutput, error) {
    var output neptune.StopDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) AddRoleToDBCluster(ctx workflow.Context, input *neptune.AddRoleToDBClusterInput) (*neptune.AddRoleToDBClusterOutput, error) {
    var output neptune.AddRoleToDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.AddRoleToDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) AddRoleToDBClusterAsync(ctx workflow.Context, input *neptune.AddRoleToDBClusterInput) *NeptuneAddRoleToDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.AddRoleToDBCluster", input)
    return &NeptuneAddRoleToDBClusterResult{Result: future}
}

func (a *NeptuneStub) AddSourceIdentifierToSubscription(ctx workflow.Context, input *neptune.AddSourceIdentifierToSubscriptionInput) (*neptune.AddSourceIdentifierToSubscriptionOutput, error) {
    var output neptune.AddSourceIdentifierToSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.AddSourceIdentifierToSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) AddSourceIdentifierToSubscriptionAsync(ctx workflow.Context, input *neptune.AddSourceIdentifierToSubscriptionInput) *NeptuneAddSourceIdentifierToSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.AddSourceIdentifierToSubscription", input)
    return &NeptuneAddSourceIdentifierToSubscriptionResult{Result: future}
}

func (a *NeptuneStub) AddTagsToResource(ctx workflow.Context, input *neptune.AddTagsToResourceInput) (*neptune.AddTagsToResourceOutput, error) {
    var output neptune.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.AddTagsToResource", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) AddTagsToResourceAsync(ctx workflow.Context, input *neptune.AddTagsToResourceInput) *NeptuneAddTagsToResourceResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.AddTagsToResource", input)
    return &NeptuneAddTagsToResourceResult{Result: future}
}

func (a *NeptuneStub) ApplyPendingMaintenanceAction(ctx workflow.Context, input *neptune.ApplyPendingMaintenanceActionInput) (*neptune.ApplyPendingMaintenanceActionOutput, error) {
    var output neptune.ApplyPendingMaintenanceActionOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ApplyPendingMaintenanceAction", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *neptune.ApplyPendingMaintenanceActionInput) *NeptuneApplyPendingMaintenanceActionResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ApplyPendingMaintenanceAction", input)
    return &NeptuneApplyPendingMaintenanceActionResult{Result: future}
}

func (a *NeptuneStub) CopyDBClusterParameterGroup(ctx workflow.Context, input *neptune.CopyDBClusterParameterGroupInput) (*neptune.CopyDBClusterParameterGroupOutput, error) {
    var output neptune.CopyDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CopyDBClusterParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CopyDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.CopyDBClusterParameterGroupInput) *NeptuneCopyDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CopyDBClusterParameterGroup", input)
    return &NeptuneCopyDBClusterParameterGroupResult{Result: future}
}

func (a *NeptuneStub) CopyDBClusterSnapshot(ctx workflow.Context, input *neptune.CopyDBClusterSnapshotInput) (*neptune.CopyDBClusterSnapshotOutput, error) {
    var output neptune.CopyDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CopyDBClusterSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CopyDBClusterSnapshotAsync(ctx workflow.Context, input *neptune.CopyDBClusterSnapshotInput) *NeptuneCopyDBClusterSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CopyDBClusterSnapshot", input)
    return &NeptuneCopyDBClusterSnapshotResult{Result: future}
}

func (a *NeptuneStub) CopyDBParameterGroup(ctx workflow.Context, input *neptune.CopyDBParameterGroupInput) (*neptune.CopyDBParameterGroupOutput, error) {
    var output neptune.CopyDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CopyDBParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CopyDBParameterGroupAsync(ctx workflow.Context, input *neptune.CopyDBParameterGroupInput) *NeptuneCopyDBParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CopyDBParameterGroup", input)
    return &NeptuneCopyDBParameterGroupResult{Result: future}
}

func (a *NeptuneStub) CreateDBCluster(ctx workflow.Context, input *neptune.CreateDBClusterInput) (*neptune.CreateDBClusterOutput, error) {
    var output neptune.CreateDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CreateDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CreateDBClusterAsync(ctx workflow.Context, input *neptune.CreateDBClusterInput) *NeptuneCreateDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CreateDBCluster", input)
    return &NeptuneCreateDBClusterResult{Result: future}
}

func (a *NeptuneStub) CreateDBClusterParameterGroup(ctx workflow.Context, input *neptune.CreateDBClusterParameterGroupInput) (*neptune.CreateDBClusterParameterGroupOutput, error) {
    var output neptune.CreateDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CreateDBClusterParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CreateDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.CreateDBClusterParameterGroupInput) *NeptuneCreateDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CreateDBClusterParameterGroup", input)
    return &NeptuneCreateDBClusterParameterGroupResult{Result: future}
}

func (a *NeptuneStub) CreateDBClusterSnapshot(ctx workflow.Context, input *neptune.CreateDBClusterSnapshotInput) (*neptune.CreateDBClusterSnapshotOutput, error) {
    var output neptune.CreateDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CreateDBClusterSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CreateDBClusterSnapshotAsync(ctx workflow.Context, input *neptune.CreateDBClusterSnapshotInput) *NeptuneCreateDBClusterSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CreateDBClusterSnapshot", input)
    return &NeptuneCreateDBClusterSnapshotResult{Result: future}
}

func (a *NeptuneStub) CreateDBInstance(ctx workflow.Context, input *neptune.CreateDBInstanceInput) (*neptune.CreateDBInstanceOutput, error) {
    var output neptune.CreateDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CreateDBInstance", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CreateDBInstanceAsync(ctx workflow.Context, input *neptune.CreateDBInstanceInput) *NeptuneCreateDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CreateDBInstance", input)
    return &NeptuneCreateDBInstanceResult{Result: future}
}

func (a *NeptuneStub) CreateDBParameterGroup(ctx workflow.Context, input *neptune.CreateDBParameterGroupInput) (*neptune.CreateDBParameterGroupOutput, error) {
    var output neptune.CreateDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CreateDBParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CreateDBParameterGroupAsync(ctx workflow.Context, input *neptune.CreateDBParameterGroupInput) *NeptuneCreateDBParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CreateDBParameterGroup", input)
    return &NeptuneCreateDBParameterGroupResult{Result: future}
}

func (a *NeptuneStub) CreateDBSubnetGroup(ctx workflow.Context, input *neptune.CreateDBSubnetGroupInput) (*neptune.CreateDBSubnetGroupOutput, error) {
    var output neptune.CreateDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CreateDBSubnetGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CreateDBSubnetGroupAsync(ctx workflow.Context, input *neptune.CreateDBSubnetGroupInput) *NeptuneCreateDBSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CreateDBSubnetGroup", input)
    return &NeptuneCreateDBSubnetGroupResult{Result: future}
}

func (a *NeptuneStub) CreateEventSubscription(ctx workflow.Context, input *neptune.CreateEventSubscriptionInput) (*neptune.CreateEventSubscriptionOutput, error) {
    var output neptune.CreateEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.CreateEventSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) CreateEventSubscriptionAsync(ctx workflow.Context, input *neptune.CreateEventSubscriptionInput) *NeptuneCreateEventSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.CreateEventSubscription", input)
    return &NeptuneCreateEventSubscriptionResult{Result: future}
}

func (a *NeptuneStub) DeleteDBCluster(ctx workflow.Context, input *neptune.DeleteDBClusterInput) (*neptune.DeleteDBClusterOutput, error) {
    var output neptune.DeleteDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DeleteDBClusterAsync(ctx workflow.Context, input *neptune.DeleteDBClusterInput) *NeptuneDeleteDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBCluster", input)
    return &NeptuneDeleteDBClusterResult{Result: future}
}

func (a *NeptuneStub) DeleteDBClusterParameterGroup(ctx workflow.Context, input *neptune.DeleteDBClusterParameterGroupInput) (*neptune.DeleteDBClusterParameterGroupOutput, error) {
    var output neptune.DeleteDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBClusterParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DeleteDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.DeleteDBClusterParameterGroupInput) *NeptuneDeleteDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBClusterParameterGroup", input)
    return &NeptuneDeleteDBClusterParameterGroupResult{Result: future}
}

func (a *NeptuneStub) DeleteDBClusterSnapshot(ctx workflow.Context, input *neptune.DeleteDBClusterSnapshotInput) (*neptune.DeleteDBClusterSnapshotOutput, error) {
    var output neptune.DeleteDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBClusterSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DeleteDBClusterSnapshotAsync(ctx workflow.Context, input *neptune.DeleteDBClusterSnapshotInput) *NeptuneDeleteDBClusterSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBClusterSnapshot", input)
    return &NeptuneDeleteDBClusterSnapshotResult{Result: future}
}

func (a *NeptuneStub) DeleteDBInstance(ctx workflow.Context, input *neptune.DeleteDBInstanceInput) (*neptune.DeleteDBInstanceOutput, error) {
    var output neptune.DeleteDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBInstance", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DeleteDBInstanceAsync(ctx workflow.Context, input *neptune.DeleteDBInstanceInput) *NeptuneDeleteDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBInstance", input)
    return &NeptuneDeleteDBInstanceResult{Result: future}
}

func (a *NeptuneStub) DeleteDBParameterGroup(ctx workflow.Context, input *neptune.DeleteDBParameterGroupInput) (*neptune.DeleteDBParameterGroupOutput, error) {
    var output neptune.DeleteDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DeleteDBParameterGroupAsync(ctx workflow.Context, input *neptune.DeleteDBParameterGroupInput) *NeptuneDeleteDBParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBParameterGroup", input)
    return &NeptuneDeleteDBParameterGroupResult{Result: future}
}

func (a *NeptuneStub) DeleteDBSubnetGroup(ctx workflow.Context, input *neptune.DeleteDBSubnetGroupInput) (*neptune.DeleteDBSubnetGroupOutput, error) {
    var output neptune.DeleteDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBSubnetGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DeleteDBSubnetGroupAsync(ctx workflow.Context, input *neptune.DeleteDBSubnetGroupInput) *NeptuneDeleteDBSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DeleteDBSubnetGroup", input)
    return &NeptuneDeleteDBSubnetGroupResult{Result: future}
}

func (a *NeptuneStub) DeleteEventSubscription(ctx workflow.Context, input *neptune.DeleteEventSubscriptionInput) (*neptune.DeleteEventSubscriptionOutput, error) {
    var output neptune.DeleteEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DeleteEventSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DeleteEventSubscriptionAsync(ctx workflow.Context, input *neptune.DeleteEventSubscriptionInput) *NeptuneDeleteEventSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DeleteEventSubscription", input)
    return &NeptuneDeleteEventSubscriptionResult{Result: future}
}

func (a *NeptuneStub) DescribeDBClusterParameterGroups(ctx workflow.Context, input *neptune.DescribeDBClusterParameterGroupsInput) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
    var output neptune.DescribeDBClusterParameterGroupsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterParameterGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBClusterParameterGroupsAsync(ctx workflow.Context, input *neptune.DescribeDBClusterParameterGroupsInput) *NeptuneDescribeDBClusterParameterGroupsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterParameterGroups", input)
    return &NeptuneDescribeDBClusterParameterGroupsResult{Result: future}
}

func (a *NeptuneStub) DescribeDBClusterParameters(ctx workflow.Context, input *neptune.DescribeDBClusterParametersInput) (*neptune.DescribeDBClusterParametersOutput, error) {
    var output neptune.DescribeDBClusterParametersOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBClusterParametersAsync(ctx workflow.Context, input *neptune.DescribeDBClusterParametersInput) *NeptuneDescribeDBClusterParametersResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterParameters", input)
    return &NeptuneDescribeDBClusterParametersResult{Result: future}
}

func (a *NeptuneStub) DescribeDBClusterSnapshotAttributes(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
    var output neptune.DescribeDBClusterSnapshotAttributesOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterSnapshotAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBClusterSnapshotAttributesAsync(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput) *NeptuneDescribeDBClusterSnapshotAttributesResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterSnapshotAttributes", input)
    return &NeptuneDescribeDBClusterSnapshotAttributesResult{Result: future}
}

func (a *NeptuneStub) DescribeDBClusterSnapshots(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotsInput) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
    var output neptune.DescribeDBClusterSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterSnapshots", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBClusterSnapshotsAsync(ctx workflow.Context, input *neptune.DescribeDBClusterSnapshotsInput) *NeptuneDescribeDBClusterSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusterSnapshots", input)
    return &NeptuneDescribeDBClusterSnapshotsResult{Result: future}
}

func (a *NeptuneStub) DescribeDBClusters(ctx workflow.Context, input *neptune.DescribeDBClustersInput) (*neptune.DescribeDBClustersOutput, error) {
    var output neptune.DescribeDBClustersOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusters", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBClustersAsync(ctx workflow.Context, input *neptune.DescribeDBClustersInput) *NeptuneDescribeDBClustersResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBClusters", input)
    return &NeptuneDescribeDBClustersResult{Result: future}
}

func (a *NeptuneStub) DescribeDBEngineVersions(ctx workflow.Context, input *neptune.DescribeDBEngineVersionsInput) (*neptune.DescribeDBEngineVersionsOutput, error) {
    var output neptune.DescribeDBEngineVersionsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBEngineVersions", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBEngineVersionsAsync(ctx workflow.Context, input *neptune.DescribeDBEngineVersionsInput) *NeptuneDescribeDBEngineVersionsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBEngineVersions", input)
    return &NeptuneDescribeDBEngineVersionsResult{Result: future}
}

func (a *NeptuneStub) DescribeDBInstances(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) (*neptune.DescribeDBInstancesOutput, error) {
    var output neptune.DescribeDBInstancesOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBInstances", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBInstancesAsync(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) *NeptuneDescribeDBInstancesResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBInstances", input)
    return &NeptuneDescribeDBInstancesResult{Result: future}
}

func (a *NeptuneStub) DescribeDBParameterGroups(ctx workflow.Context, input *neptune.DescribeDBParameterGroupsInput) (*neptune.DescribeDBParameterGroupsOutput, error) {
    var output neptune.DescribeDBParameterGroupsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBParameterGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBParameterGroupsAsync(ctx workflow.Context, input *neptune.DescribeDBParameterGroupsInput) *NeptuneDescribeDBParameterGroupsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBParameterGroups", input)
    return &NeptuneDescribeDBParameterGroupsResult{Result: future}
}

func (a *NeptuneStub) DescribeDBParameters(ctx workflow.Context, input *neptune.DescribeDBParametersInput) (*neptune.DescribeDBParametersOutput, error) {
    var output neptune.DescribeDBParametersOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBParametersAsync(ctx workflow.Context, input *neptune.DescribeDBParametersInput) *NeptuneDescribeDBParametersResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBParameters", input)
    return &NeptuneDescribeDBParametersResult{Result: future}
}

func (a *NeptuneStub) DescribeDBSubnetGroups(ctx workflow.Context, input *neptune.DescribeDBSubnetGroupsInput) (*neptune.DescribeDBSubnetGroupsOutput, error) {
    var output neptune.DescribeDBSubnetGroupsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBSubnetGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeDBSubnetGroupsAsync(ctx workflow.Context, input *neptune.DescribeDBSubnetGroupsInput) *NeptuneDescribeDBSubnetGroupsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeDBSubnetGroups", input)
    return &NeptuneDescribeDBSubnetGroupsResult{Result: future}
}

func (a *NeptuneStub) DescribeEngineDefaultClusterParameters(ctx workflow.Context, input *neptune.DescribeEngineDefaultClusterParametersInput) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
    var output neptune.DescribeEngineDefaultClusterParametersOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeEngineDefaultClusterParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeEngineDefaultClusterParametersAsync(ctx workflow.Context, input *neptune.DescribeEngineDefaultClusterParametersInput) *NeptuneDescribeEngineDefaultClusterParametersResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeEngineDefaultClusterParameters", input)
    return &NeptuneDescribeEngineDefaultClusterParametersResult{Result: future}
}

func (a *NeptuneStub) DescribeEngineDefaultParameters(ctx workflow.Context, input *neptune.DescribeEngineDefaultParametersInput) (*neptune.DescribeEngineDefaultParametersOutput, error) {
    var output neptune.DescribeEngineDefaultParametersOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeEngineDefaultParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeEngineDefaultParametersAsync(ctx workflow.Context, input *neptune.DescribeEngineDefaultParametersInput) *NeptuneDescribeEngineDefaultParametersResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeEngineDefaultParameters", input)
    return &NeptuneDescribeEngineDefaultParametersResult{Result: future}
}

func (a *NeptuneStub) DescribeEventCategories(ctx workflow.Context, input *neptune.DescribeEventCategoriesInput) (*neptune.DescribeEventCategoriesOutput, error) {
    var output neptune.DescribeEventCategoriesOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeEventCategories", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeEventCategoriesAsync(ctx workflow.Context, input *neptune.DescribeEventCategoriesInput) *NeptuneDescribeEventCategoriesResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeEventCategories", input)
    return &NeptuneDescribeEventCategoriesResult{Result: future}
}

func (a *NeptuneStub) DescribeEventSubscriptions(ctx workflow.Context, input *neptune.DescribeEventSubscriptionsInput) (*neptune.DescribeEventSubscriptionsOutput, error) {
    var output neptune.DescribeEventSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeEventSubscriptions", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeEventSubscriptionsAsync(ctx workflow.Context, input *neptune.DescribeEventSubscriptionsInput) *NeptuneDescribeEventSubscriptionsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeEventSubscriptions", input)
    return &NeptuneDescribeEventSubscriptionsResult{Result: future}
}

func (a *NeptuneStub) DescribeEvents(ctx workflow.Context, input *neptune.DescribeEventsInput) (*neptune.DescribeEventsOutput, error) {
    var output neptune.DescribeEventsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeEvents", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeEventsAsync(ctx workflow.Context, input *neptune.DescribeEventsInput) *NeptuneDescribeEventsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeEvents", input)
    return &NeptuneDescribeEventsResult{Result: future}
}

func (a *NeptuneStub) DescribeOrderableDBInstanceOptions(ctx workflow.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
    var output neptune.DescribeOrderableDBInstanceOptionsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeOrderableDBInstanceOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeOrderableDBInstanceOptionsAsync(ctx workflow.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput) *NeptuneDescribeOrderableDBInstanceOptionsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeOrderableDBInstanceOptions", input)
    return &NeptuneDescribeOrderableDBInstanceOptionsResult{Result: future}
}

func (a *NeptuneStub) DescribePendingMaintenanceActions(ctx workflow.Context, input *neptune.DescribePendingMaintenanceActionsInput) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
    var output neptune.DescribePendingMaintenanceActionsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribePendingMaintenanceActions", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *neptune.DescribePendingMaintenanceActionsInput) *NeptuneDescribePendingMaintenanceActionsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribePendingMaintenanceActions", input)
    return &NeptuneDescribePendingMaintenanceActionsResult{Result: future}
}

func (a *NeptuneStub) DescribeValidDBInstanceModifications(ctx workflow.Context, input *neptune.DescribeValidDBInstanceModificationsInput) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
    var output neptune.DescribeValidDBInstanceModificationsOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.DescribeValidDBInstanceModifications", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) DescribeValidDBInstanceModificationsAsync(ctx workflow.Context, input *neptune.DescribeValidDBInstanceModificationsInput) *NeptuneDescribeValidDBInstanceModificationsResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.DescribeValidDBInstanceModifications", input)
    return &NeptuneDescribeValidDBInstanceModificationsResult{Result: future}
}

func (a *NeptuneStub) FailoverDBCluster(ctx workflow.Context, input *neptune.FailoverDBClusterInput) (*neptune.FailoverDBClusterOutput, error) {
    var output neptune.FailoverDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.FailoverDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) FailoverDBClusterAsync(ctx workflow.Context, input *neptune.FailoverDBClusterInput) *NeptuneFailoverDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.FailoverDBCluster", input)
    return &NeptuneFailoverDBClusterResult{Result: future}
}

func (a *NeptuneStub) ListTagsForResource(ctx workflow.Context, input *neptune.ListTagsForResourceInput) (*neptune.ListTagsForResourceOutput, error) {
    var output neptune.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ListTagsForResourceAsync(ctx workflow.Context, input *neptune.ListTagsForResourceInput) *NeptuneListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ListTagsForResource", input)
    return &NeptuneListTagsForResourceResult{Result: future}
}

func (a *NeptuneStub) ModifyDBCluster(ctx workflow.Context, input *neptune.ModifyDBClusterInput) (*neptune.ModifyDBClusterOutput, error) {
    var output neptune.ModifyDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ModifyDBClusterAsync(ctx workflow.Context, input *neptune.ModifyDBClusterInput) *NeptuneModifyDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBCluster", input)
    return &NeptuneModifyDBClusterResult{Result: future}
}

func (a *NeptuneStub) ModifyDBClusterParameterGroup(ctx workflow.Context, input *neptune.ModifyDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
    var output neptune.ResetDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBClusterParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ModifyDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.ModifyDBClusterParameterGroupInput) *NeptuneModifyDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBClusterParameterGroup", input)
    return &NeptuneModifyDBClusterParameterGroupResult{Result: future}
}

func (a *NeptuneStub) ModifyDBClusterSnapshotAttribute(ctx workflow.Context, input *neptune.ModifyDBClusterSnapshotAttributeInput) (*neptune.ModifyDBClusterSnapshotAttributeOutput, error) {
    var output neptune.ModifyDBClusterSnapshotAttributeOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBClusterSnapshotAttribute", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ModifyDBClusterSnapshotAttributeAsync(ctx workflow.Context, input *neptune.ModifyDBClusterSnapshotAttributeInput) *NeptuneModifyDBClusterSnapshotAttributeResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBClusterSnapshotAttribute", input)
    return &NeptuneModifyDBClusterSnapshotAttributeResult{Result: future}
}

func (a *NeptuneStub) ModifyDBInstance(ctx workflow.Context, input *neptune.ModifyDBInstanceInput) (*neptune.ModifyDBInstanceOutput, error) {
    var output neptune.ModifyDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBInstance", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ModifyDBInstanceAsync(ctx workflow.Context, input *neptune.ModifyDBInstanceInput) *NeptuneModifyDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBInstance", input)
    return &NeptuneModifyDBInstanceResult{Result: future}
}

func (a *NeptuneStub) ModifyDBParameterGroup(ctx workflow.Context, input *neptune.ModifyDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
    var output neptune.ResetDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ModifyDBParameterGroupAsync(ctx workflow.Context, input *neptune.ModifyDBParameterGroupInput) *NeptuneModifyDBParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBParameterGroup", input)
    return &NeptuneModifyDBParameterGroupResult{Result: future}
}

func (a *NeptuneStub) ModifyDBSubnetGroup(ctx workflow.Context, input *neptune.ModifyDBSubnetGroupInput) (*neptune.ModifyDBSubnetGroupOutput, error) {
    var output neptune.ModifyDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBSubnetGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ModifyDBSubnetGroupAsync(ctx workflow.Context, input *neptune.ModifyDBSubnetGroupInput) *NeptuneModifyDBSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ModifyDBSubnetGroup", input)
    return &NeptuneModifyDBSubnetGroupResult{Result: future}
}

func (a *NeptuneStub) ModifyEventSubscription(ctx workflow.Context, input *neptune.ModifyEventSubscriptionInput) (*neptune.ModifyEventSubscriptionOutput, error) {
    var output neptune.ModifyEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ModifyEventSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ModifyEventSubscriptionAsync(ctx workflow.Context, input *neptune.ModifyEventSubscriptionInput) *NeptuneModifyEventSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ModifyEventSubscription", input)
    return &NeptuneModifyEventSubscriptionResult{Result: future}
}

func (a *NeptuneStub) PromoteReadReplicaDBCluster(ctx workflow.Context, input *neptune.PromoteReadReplicaDBClusterInput) (*neptune.PromoteReadReplicaDBClusterOutput, error) {
    var output neptune.PromoteReadReplicaDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.PromoteReadReplicaDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) PromoteReadReplicaDBClusterAsync(ctx workflow.Context, input *neptune.PromoteReadReplicaDBClusterInput) *NeptunePromoteReadReplicaDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.PromoteReadReplicaDBCluster", input)
    return &NeptunePromoteReadReplicaDBClusterResult{Result: future}
}

func (a *NeptuneStub) RebootDBInstance(ctx workflow.Context, input *neptune.RebootDBInstanceInput) (*neptune.RebootDBInstanceOutput, error) {
    var output neptune.RebootDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.RebootDBInstance", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) RebootDBInstanceAsync(ctx workflow.Context, input *neptune.RebootDBInstanceInput) *NeptuneRebootDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.RebootDBInstance", input)
    return &NeptuneRebootDBInstanceResult{Result: future}
}

func (a *NeptuneStub) RemoveRoleFromDBCluster(ctx workflow.Context, input *neptune.RemoveRoleFromDBClusterInput) (*neptune.RemoveRoleFromDBClusterOutput, error) {
    var output neptune.RemoveRoleFromDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.RemoveRoleFromDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) RemoveRoleFromDBClusterAsync(ctx workflow.Context, input *neptune.RemoveRoleFromDBClusterInput) *NeptuneRemoveRoleFromDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.RemoveRoleFromDBCluster", input)
    return &NeptuneRemoveRoleFromDBClusterResult{Result: future}
}

func (a *NeptuneStub) RemoveSourceIdentifierFromSubscription(ctx workflow.Context, input *neptune.RemoveSourceIdentifierFromSubscriptionInput) (*neptune.RemoveSourceIdentifierFromSubscriptionOutput, error) {
    var output neptune.RemoveSourceIdentifierFromSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.RemoveSourceIdentifierFromSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) RemoveSourceIdentifierFromSubscriptionAsync(ctx workflow.Context, input *neptune.RemoveSourceIdentifierFromSubscriptionInput) *NeptuneRemoveSourceIdentifierFromSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.RemoveSourceIdentifierFromSubscription", input)
    return &NeptuneRemoveSourceIdentifierFromSubscriptionResult{Result: future}
}

func (a *NeptuneStub) RemoveTagsFromResource(ctx workflow.Context, input *neptune.RemoveTagsFromResourceInput) (*neptune.RemoveTagsFromResourceOutput, error) {
    var output neptune.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.RemoveTagsFromResource", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *neptune.RemoveTagsFromResourceInput) *NeptuneRemoveTagsFromResourceResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.RemoveTagsFromResource", input)
    return &NeptuneRemoveTagsFromResourceResult{Result: future}
}

func (a *NeptuneStub) ResetDBClusterParameterGroup(ctx workflow.Context, input *neptune.ResetDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
    var output neptune.ResetDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ResetDBClusterParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ResetDBClusterParameterGroupAsync(ctx workflow.Context, input *neptune.ResetDBClusterParameterGroupInput) *NeptuneResetDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ResetDBClusterParameterGroup", input)
    return &NeptuneResetDBClusterParameterGroupResult{Result: future}
}

func (a *NeptuneStub) ResetDBParameterGroup(ctx workflow.Context, input *neptune.ResetDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
    var output neptune.ResetDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.ResetDBParameterGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) ResetDBParameterGroupAsync(ctx workflow.Context, input *neptune.ResetDBParameterGroupInput) *NeptuneResetDBParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.ResetDBParameterGroup", input)
    return &NeptuneResetDBParameterGroupResult{Result: future}
}

func (a *NeptuneStub) RestoreDBClusterFromSnapshot(ctx workflow.Context, input *neptune.RestoreDBClusterFromSnapshotInput) (*neptune.RestoreDBClusterFromSnapshotOutput, error) {
    var output neptune.RestoreDBClusterFromSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.RestoreDBClusterFromSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) RestoreDBClusterFromSnapshotAsync(ctx workflow.Context, input *neptune.RestoreDBClusterFromSnapshotInput) *NeptuneRestoreDBClusterFromSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.RestoreDBClusterFromSnapshot", input)
    return &NeptuneRestoreDBClusterFromSnapshotResult{Result: future}
}

func (a *NeptuneStub) RestoreDBClusterToPointInTime(ctx workflow.Context, input *neptune.RestoreDBClusterToPointInTimeInput) (*neptune.RestoreDBClusterToPointInTimeOutput, error) {
    var output neptune.RestoreDBClusterToPointInTimeOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.RestoreDBClusterToPointInTime", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) RestoreDBClusterToPointInTimeAsync(ctx workflow.Context, input *neptune.RestoreDBClusterToPointInTimeInput) *NeptuneRestoreDBClusterToPointInTimeResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.RestoreDBClusterToPointInTime", input)
    return &NeptuneRestoreDBClusterToPointInTimeResult{Result: future}
}

func (a *NeptuneStub) StartDBCluster(ctx workflow.Context, input *neptune.StartDBClusterInput) (*neptune.StartDBClusterOutput, error) {
    var output neptune.StartDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.StartDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) StartDBClusterAsync(ctx workflow.Context, input *neptune.StartDBClusterInput) *NeptuneStartDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.StartDBCluster", input)
    return &NeptuneStartDBClusterResult{Result: future}
}

func (a *NeptuneStub) StopDBCluster(ctx workflow.Context, input *neptune.StopDBClusterInput) (*neptune.StopDBClusterOutput, error) {
    var output neptune.StopDBClusterOutput
    err := workflow.ExecuteActivity(ctx, "Neptune.StopDBCluster", input).Get(ctx, &output)
    return &output, err
}

func (a *NeptuneStub) StopDBClusterAsync(ctx workflow.Context, input *neptune.StopDBClusterInput) *NeptuneStopDBClusterResult {
    future := workflow.ExecuteActivity(ctx, "Neptune.StopDBCluster", input)
    return &NeptuneStopDBClusterResult{Result: future}
}

func (a *NeptuneStub) WaitUntilDBInstanceAvailable(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) error {
    return workflow.ExecuteActivity(ctx, "Neptune.WaitUntilDBInstanceAvailable", input).Get(ctx, nil)
}

func (a *NeptuneStub) WaitUntilDBInstanceAvailableAsync(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "Neptune.WaitUntilDBInstanceAvailable", input)
}


func (a *NeptuneStub) WaitUntilDBInstanceDeleted(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) error {
    return workflow.ExecuteActivity(ctx, "Neptune.WaitUntilDBInstanceDeleted", input).Get(ctx, nil)
}

func (a *NeptuneStub) WaitUntilDBInstanceDeletedAsync(ctx workflow.Context, input *neptune.DescribeDBInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "Neptune.WaitUntilDBInstanceDeleted", input)
}

