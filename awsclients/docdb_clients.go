package awsclients

import (
	"github.com/aws/aws-sdk-go/service/docdb"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DocDBClient interface {
    AddTagsToResource(ctx workflow.Context, input *docdb.AddTagsToResourceInput) (*docdb.AddTagsToResourceOutput, error)
    AddTagsToResourceAsync(ctx workflow.Context, input *docdb.AddTagsToResourceInput) *DocdbAddTagsToResourceResult

    ApplyPendingMaintenanceAction(ctx workflow.Context, input *docdb.ApplyPendingMaintenanceActionInput) (*docdb.ApplyPendingMaintenanceActionOutput, error)
    ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *docdb.ApplyPendingMaintenanceActionInput) *DocdbApplyPendingMaintenanceActionResult

    CopyDBClusterParameterGroup(ctx workflow.Context, input *docdb.CopyDBClusterParameterGroupInput) (*docdb.CopyDBClusterParameterGroupOutput, error)
    CopyDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.CopyDBClusterParameterGroupInput) *DocdbCopyDBClusterParameterGroupResult

    CopyDBClusterSnapshot(ctx workflow.Context, input *docdb.CopyDBClusterSnapshotInput) (*docdb.CopyDBClusterSnapshotOutput, error)
    CopyDBClusterSnapshotAsync(ctx workflow.Context, input *docdb.CopyDBClusterSnapshotInput) *DocdbCopyDBClusterSnapshotResult

    CreateDBCluster(ctx workflow.Context, input *docdb.CreateDBClusterInput) (*docdb.CreateDBClusterOutput, error)
    CreateDBClusterAsync(ctx workflow.Context, input *docdb.CreateDBClusterInput) *DocdbCreateDBClusterResult

    CreateDBClusterParameterGroup(ctx workflow.Context, input *docdb.CreateDBClusterParameterGroupInput) (*docdb.CreateDBClusterParameterGroupOutput, error)
    CreateDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.CreateDBClusterParameterGroupInput) *DocdbCreateDBClusterParameterGroupResult

    CreateDBClusterSnapshot(ctx workflow.Context, input *docdb.CreateDBClusterSnapshotInput) (*docdb.CreateDBClusterSnapshotOutput, error)
    CreateDBClusterSnapshotAsync(ctx workflow.Context, input *docdb.CreateDBClusterSnapshotInput) *DocdbCreateDBClusterSnapshotResult

    CreateDBInstance(ctx workflow.Context, input *docdb.CreateDBInstanceInput) (*docdb.CreateDBInstanceOutput, error)
    CreateDBInstanceAsync(ctx workflow.Context, input *docdb.CreateDBInstanceInput) *DocdbCreateDBInstanceResult

    CreateDBSubnetGroup(ctx workflow.Context, input *docdb.CreateDBSubnetGroupInput) (*docdb.CreateDBSubnetGroupOutput, error)
    CreateDBSubnetGroupAsync(ctx workflow.Context, input *docdb.CreateDBSubnetGroupInput) *DocdbCreateDBSubnetGroupResult

    DeleteDBCluster(ctx workflow.Context, input *docdb.DeleteDBClusterInput) (*docdb.DeleteDBClusterOutput, error)
    DeleteDBClusterAsync(ctx workflow.Context, input *docdb.DeleteDBClusterInput) *DocdbDeleteDBClusterResult

    DeleteDBClusterParameterGroup(ctx workflow.Context, input *docdb.DeleteDBClusterParameterGroupInput) (*docdb.DeleteDBClusterParameterGroupOutput, error)
    DeleteDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.DeleteDBClusterParameterGroupInput) *DocdbDeleteDBClusterParameterGroupResult

    DeleteDBClusterSnapshot(ctx workflow.Context, input *docdb.DeleteDBClusterSnapshotInput) (*docdb.DeleteDBClusterSnapshotOutput, error)
    DeleteDBClusterSnapshotAsync(ctx workflow.Context, input *docdb.DeleteDBClusterSnapshotInput) *DocdbDeleteDBClusterSnapshotResult

    DeleteDBInstance(ctx workflow.Context, input *docdb.DeleteDBInstanceInput) (*docdb.DeleteDBInstanceOutput, error)
    DeleteDBInstanceAsync(ctx workflow.Context, input *docdb.DeleteDBInstanceInput) *DocdbDeleteDBInstanceResult

    DeleteDBSubnetGroup(ctx workflow.Context, input *docdb.DeleteDBSubnetGroupInput) (*docdb.DeleteDBSubnetGroupOutput, error)
    DeleteDBSubnetGroupAsync(ctx workflow.Context, input *docdb.DeleteDBSubnetGroupInput) *DocdbDeleteDBSubnetGroupResult

    DescribeCertificates(ctx workflow.Context, input *docdb.DescribeCertificatesInput) (*docdb.DescribeCertificatesOutput, error)
    DescribeCertificatesAsync(ctx workflow.Context, input *docdb.DescribeCertificatesInput) *DocdbDescribeCertificatesResult

    DescribeDBClusterParameterGroups(ctx workflow.Context, input *docdb.DescribeDBClusterParameterGroupsInput) (*docdb.DescribeDBClusterParameterGroupsOutput, error)
    DescribeDBClusterParameterGroupsAsync(ctx workflow.Context, input *docdb.DescribeDBClusterParameterGroupsInput) *DocdbDescribeDBClusterParameterGroupsResult

    DescribeDBClusterParameters(ctx workflow.Context, input *docdb.DescribeDBClusterParametersInput) (*docdb.DescribeDBClusterParametersOutput, error)
    DescribeDBClusterParametersAsync(ctx workflow.Context, input *docdb.DescribeDBClusterParametersInput) *DocdbDescribeDBClusterParametersResult

    DescribeDBClusterSnapshotAttributes(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotAttributesInput) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error)
    DescribeDBClusterSnapshotAttributesAsync(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotAttributesInput) *DocdbDescribeDBClusterSnapshotAttributesResult

    DescribeDBClusterSnapshots(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotsInput) (*docdb.DescribeDBClusterSnapshotsOutput, error)
    DescribeDBClusterSnapshotsAsync(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotsInput) *DocdbDescribeDBClusterSnapshotsResult

    DescribeDBClusters(ctx workflow.Context, input *docdb.DescribeDBClustersInput) (*docdb.DescribeDBClustersOutput, error)
    DescribeDBClustersAsync(ctx workflow.Context, input *docdb.DescribeDBClustersInput) *DocdbDescribeDBClustersResult

    DescribeDBEngineVersions(ctx workflow.Context, input *docdb.DescribeDBEngineVersionsInput) (*docdb.DescribeDBEngineVersionsOutput, error)
    DescribeDBEngineVersionsAsync(ctx workflow.Context, input *docdb.DescribeDBEngineVersionsInput) *DocdbDescribeDBEngineVersionsResult

    DescribeDBInstances(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) (*docdb.DescribeDBInstancesOutput, error)
    DescribeDBInstancesAsync(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) *DocdbDescribeDBInstancesResult

    DescribeDBSubnetGroups(ctx workflow.Context, input *docdb.DescribeDBSubnetGroupsInput) (*docdb.DescribeDBSubnetGroupsOutput, error)
    DescribeDBSubnetGroupsAsync(ctx workflow.Context, input *docdb.DescribeDBSubnetGroupsInput) *DocdbDescribeDBSubnetGroupsResult

    DescribeEngineDefaultClusterParameters(ctx workflow.Context, input *docdb.DescribeEngineDefaultClusterParametersInput) (*docdb.DescribeEngineDefaultClusterParametersOutput, error)
    DescribeEngineDefaultClusterParametersAsync(ctx workflow.Context, input *docdb.DescribeEngineDefaultClusterParametersInput) *DocdbDescribeEngineDefaultClusterParametersResult

    DescribeEventCategories(ctx workflow.Context, input *docdb.DescribeEventCategoriesInput) (*docdb.DescribeEventCategoriesOutput, error)
    DescribeEventCategoriesAsync(ctx workflow.Context, input *docdb.DescribeEventCategoriesInput) *DocdbDescribeEventCategoriesResult

    DescribeEvents(ctx workflow.Context, input *docdb.DescribeEventsInput) (*docdb.DescribeEventsOutput, error)
    DescribeEventsAsync(ctx workflow.Context, input *docdb.DescribeEventsInput) *DocdbDescribeEventsResult

    DescribeOrderableDBInstanceOptions(ctx workflow.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error)
    DescribeOrderableDBInstanceOptionsAsync(ctx workflow.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput) *DocdbDescribeOrderableDBInstanceOptionsResult

    DescribePendingMaintenanceActions(ctx workflow.Context, input *docdb.DescribePendingMaintenanceActionsInput) (*docdb.DescribePendingMaintenanceActionsOutput, error)
    DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *docdb.DescribePendingMaintenanceActionsInput) *DocdbDescribePendingMaintenanceActionsResult

    FailoverDBCluster(ctx workflow.Context, input *docdb.FailoverDBClusterInput) (*docdb.FailoverDBClusterOutput, error)
    FailoverDBClusterAsync(ctx workflow.Context, input *docdb.FailoverDBClusterInput) *DocdbFailoverDBClusterResult

    ListTagsForResource(ctx workflow.Context, input *docdb.ListTagsForResourceInput) (*docdb.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *docdb.ListTagsForResourceInput) *DocdbListTagsForResourceResult

    ModifyDBCluster(ctx workflow.Context, input *docdb.ModifyDBClusterInput) (*docdb.ModifyDBClusterOutput, error)
    ModifyDBClusterAsync(ctx workflow.Context, input *docdb.ModifyDBClusterInput) *DocdbModifyDBClusterResult

    ModifyDBClusterParameterGroup(ctx workflow.Context, input *docdb.ModifyDBClusterParameterGroupInput) (*docdb.ModifyDBClusterParameterGroupOutput, error)
    ModifyDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.ModifyDBClusterParameterGroupInput) *DocdbModifyDBClusterParameterGroupResult

    ModifyDBClusterSnapshotAttribute(ctx workflow.Context, input *docdb.ModifyDBClusterSnapshotAttributeInput) (*docdb.ModifyDBClusterSnapshotAttributeOutput, error)
    ModifyDBClusterSnapshotAttributeAsync(ctx workflow.Context, input *docdb.ModifyDBClusterSnapshotAttributeInput) *DocdbModifyDBClusterSnapshotAttributeResult

    ModifyDBInstance(ctx workflow.Context, input *docdb.ModifyDBInstanceInput) (*docdb.ModifyDBInstanceOutput, error)
    ModifyDBInstanceAsync(ctx workflow.Context, input *docdb.ModifyDBInstanceInput) *DocdbModifyDBInstanceResult

    ModifyDBSubnetGroup(ctx workflow.Context, input *docdb.ModifyDBSubnetGroupInput) (*docdb.ModifyDBSubnetGroupOutput, error)
    ModifyDBSubnetGroupAsync(ctx workflow.Context, input *docdb.ModifyDBSubnetGroupInput) *DocdbModifyDBSubnetGroupResult

    RebootDBInstance(ctx workflow.Context, input *docdb.RebootDBInstanceInput) (*docdb.RebootDBInstanceOutput, error)
    RebootDBInstanceAsync(ctx workflow.Context, input *docdb.RebootDBInstanceInput) *DocdbRebootDBInstanceResult

    RemoveTagsFromResource(ctx workflow.Context, input *docdb.RemoveTagsFromResourceInput) (*docdb.RemoveTagsFromResourceOutput, error)
    RemoveTagsFromResourceAsync(ctx workflow.Context, input *docdb.RemoveTagsFromResourceInput) *DocdbRemoveTagsFromResourceResult

    ResetDBClusterParameterGroup(ctx workflow.Context, input *docdb.ResetDBClusterParameterGroupInput) (*docdb.ResetDBClusterParameterGroupOutput, error)
    ResetDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.ResetDBClusterParameterGroupInput) *DocdbResetDBClusterParameterGroupResult

    RestoreDBClusterFromSnapshot(ctx workflow.Context, input *docdb.RestoreDBClusterFromSnapshotInput) (*docdb.RestoreDBClusterFromSnapshotOutput, error)
    RestoreDBClusterFromSnapshotAsync(ctx workflow.Context, input *docdb.RestoreDBClusterFromSnapshotInput) *DocdbRestoreDBClusterFromSnapshotResult

    RestoreDBClusterToPointInTime(ctx workflow.Context, input *docdb.RestoreDBClusterToPointInTimeInput) (*docdb.RestoreDBClusterToPointInTimeOutput, error)
    RestoreDBClusterToPointInTimeAsync(ctx workflow.Context, input *docdb.RestoreDBClusterToPointInTimeInput) *DocdbRestoreDBClusterToPointInTimeResult

    StartDBCluster(ctx workflow.Context, input *docdb.StartDBClusterInput) (*docdb.StartDBClusterOutput, error)
    StartDBClusterAsync(ctx workflow.Context, input *docdb.StartDBClusterInput) *DocdbStartDBClusterResult

    StopDBCluster(ctx workflow.Context, input *docdb.StopDBClusterInput) (*docdb.StopDBClusterOutput, error)
    StopDBClusterAsync(ctx workflow.Context, input *docdb.StopDBClusterInput) *DocdbStopDBClusterResult

    WaitUntilDBInstanceAvailable(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) error
    WaitUntilDBInstanceDeleted(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) error}
type DocdbAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *DocdbAddTagsToResourceResult) Get(ctx workflow.Context) (*docdb.AddTagsToResourceOutput, error) {
    var output docdb.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbApplyPendingMaintenanceActionResult struct {
	Result workflow.Future
}

func (r *DocdbApplyPendingMaintenanceActionResult) Get(ctx workflow.Context) (*docdb.ApplyPendingMaintenanceActionOutput, error) {
    var output docdb.ApplyPendingMaintenanceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbCopyDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *DocdbCopyDBClusterParameterGroupResult) Get(ctx workflow.Context) (*docdb.CopyDBClusterParameterGroupOutput, error) {
    var output docdb.CopyDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbCopyDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *DocdbCopyDBClusterSnapshotResult) Get(ctx workflow.Context) (*docdb.CopyDBClusterSnapshotOutput, error) {
    var output docdb.CopyDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbCreateDBClusterResult struct {
	Result workflow.Future
}

func (r *DocdbCreateDBClusterResult) Get(ctx workflow.Context) (*docdb.CreateDBClusterOutput, error) {
    var output docdb.CreateDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbCreateDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *DocdbCreateDBClusterParameterGroupResult) Get(ctx workflow.Context) (*docdb.CreateDBClusterParameterGroupOutput, error) {
    var output docdb.CreateDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbCreateDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *DocdbCreateDBClusterSnapshotResult) Get(ctx workflow.Context) (*docdb.CreateDBClusterSnapshotOutput, error) {
    var output docdb.CreateDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbCreateDBInstanceResult struct {
	Result workflow.Future
}

func (r *DocdbCreateDBInstanceResult) Get(ctx workflow.Context) (*docdb.CreateDBInstanceOutput, error) {
    var output docdb.CreateDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbCreateDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DocdbCreateDBSubnetGroupResult) Get(ctx workflow.Context) (*docdb.CreateDBSubnetGroupOutput, error) {
    var output docdb.CreateDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDeleteDBClusterResult struct {
	Result workflow.Future
}

func (r *DocdbDeleteDBClusterResult) Get(ctx workflow.Context) (*docdb.DeleteDBClusterOutput, error) {
    var output docdb.DeleteDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDeleteDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *DocdbDeleteDBClusterParameterGroupResult) Get(ctx workflow.Context) (*docdb.DeleteDBClusterParameterGroupOutput, error) {
    var output docdb.DeleteDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDeleteDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *DocdbDeleteDBClusterSnapshotResult) Get(ctx workflow.Context) (*docdb.DeleteDBClusterSnapshotOutput, error) {
    var output docdb.DeleteDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDeleteDBInstanceResult struct {
	Result workflow.Future
}

func (r *DocdbDeleteDBInstanceResult) Get(ctx workflow.Context) (*docdb.DeleteDBInstanceOutput, error) {
    var output docdb.DeleteDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDeleteDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DocdbDeleteDBSubnetGroupResult) Get(ctx workflow.Context) (*docdb.DeleteDBSubnetGroupOutput, error) {
    var output docdb.DeleteDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeCertificatesResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeCertificatesResult) Get(ctx workflow.Context) (*docdb.DescribeCertificatesOutput, error) {
    var output docdb.DescribeCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBClusterParameterGroupsResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBClusterParameterGroupsResult) Get(ctx workflow.Context) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
    var output docdb.DescribeDBClusterParameterGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBClusterParametersResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBClusterParametersResult) Get(ctx workflow.Context) (*docdb.DescribeDBClusterParametersOutput, error) {
    var output docdb.DescribeDBClusterParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBClusterSnapshotAttributesResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBClusterSnapshotAttributesResult) Get(ctx workflow.Context) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
    var output docdb.DescribeDBClusterSnapshotAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBClusterSnapshotsResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBClusterSnapshotsResult) Get(ctx workflow.Context) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
    var output docdb.DescribeDBClusterSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBClustersResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBClustersResult) Get(ctx workflow.Context) (*docdb.DescribeDBClustersOutput, error) {
    var output docdb.DescribeDBClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBEngineVersionsResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBEngineVersionsResult) Get(ctx workflow.Context) (*docdb.DescribeDBEngineVersionsOutput, error) {
    var output docdb.DescribeDBEngineVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBInstancesResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBInstancesResult) Get(ctx workflow.Context) (*docdb.DescribeDBInstancesOutput, error) {
    var output docdb.DescribeDBInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeDBSubnetGroupsResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeDBSubnetGroupsResult) Get(ctx workflow.Context) (*docdb.DescribeDBSubnetGroupsOutput, error) {
    var output docdb.DescribeDBSubnetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeEngineDefaultClusterParametersResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeEngineDefaultClusterParametersResult) Get(ctx workflow.Context) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
    var output docdb.DescribeEngineDefaultClusterParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeEventCategoriesResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeEventCategoriesResult) Get(ctx workflow.Context) (*docdb.DescribeEventCategoriesOutput, error) {
    var output docdb.DescribeEventCategoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeEventsResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeEventsResult) Get(ctx workflow.Context) (*docdb.DescribeEventsOutput, error) {
    var output docdb.DescribeEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribeOrderableDBInstanceOptionsResult struct {
	Result workflow.Future
}

func (r *DocdbDescribeOrderableDBInstanceOptionsResult) Get(ctx workflow.Context) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
    var output docdb.DescribeOrderableDBInstanceOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbDescribePendingMaintenanceActionsResult struct {
	Result workflow.Future
}

func (r *DocdbDescribePendingMaintenanceActionsResult) Get(ctx workflow.Context) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
    var output docdb.DescribePendingMaintenanceActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbFailoverDBClusterResult struct {
	Result workflow.Future
}

func (r *DocdbFailoverDBClusterResult) Get(ctx workflow.Context) (*docdb.FailoverDBClusterOutput, error) {
    var output docdb.FailoverDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *DocdbListTagsForResourceResult) Get(ctx workflow.Context) (*docdb.ListTagsForResourceOutput, error) {
    var output docdb.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbModifyDBClusterResult struct {
	Result workflow.Future
}

func (r *DocdbModifyDBClusterResult) Get(ctx workflow.Context) (*docdb.ModifyDBClusterOutput, error) {
    var output docdb.ModifyDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbModifyDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *DocdbModifyDBClusterParameterGroupResult) Get(ctx workflow.Context) (*docdb.ModifyDBClusterParameterGroupOutput, error) {
    var output docdb.ModifyDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbModifyDBClusterSnapshotAttributeResult struct {
	Result workflow.Future
}

func (r *DocdbModifyDBClusterSnapshotAttributeResult) Get(ctx workflow.Context) (*docdb.ModifyDBClusterSnapshotAttributeOutput, error) {
    var output docdb.ModifyDBClusterSnapshotAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbModifyDBInstanceResult struct {
	Result workflow.Future
}

func (r *DocdbModifyDBInstanceResult) Get(ctx workflow.Context) (*docdb.ModifyDBInstanceOutput, error) {
    var output docdb.ModifyDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbModifyDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DocdbModifyDBSubnetGroupResult) Get(ctx workflow.Context) (*docdb.ModifyDBSubnetGroupOutput, error) {
    var output docdb.ModifyDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbRebootDBInstanceResult struct {
	Result workflow.Future
}

func (r *DocdbRebootDBInstanceResult) Get(ctx workflow.Context) (*docdb.RebootDBInstanceOutput, error) {
    var output docdb.RebootDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *DocdbRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*docdb.RemoveTagsFromResourceOutput, error) {
    var output docdb.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbResetDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *DocdbResetDBClusterParameterGroupResult) Get(ctx workflow.Context) (*docdb.ResetDBClusterParameterGroupOutput, error) {
    var output docdb.ResetDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbRestoreDBClusterFromSnapshotResult struct {
	Result workflow.Future
}

func (r *DocdbRestoreDBClusterFromSnapshotResult) Get(ctx workflow.Context) (*docdb.RestoreDBClusterFromSnapshotOutput, error) {
    var output docdb.RestoreDBClusterFromSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbRestoreDBClusterToPointInTimeResult struct {
	Result workflow.Future
}

func (r *DocdbRestoreDBClusterToPointInTimeResult) Get(ctx workflow.Context) (*docdb.RestoreDBClusterToPointInTimeOutput, error) {
    var output docdb.RestoreDBClusterToPointInTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbStartDBClusterResult struct {
	Result workflow.Future
}

func (r *DocdbStartDBClusterResult) Get(ctx workflow.Context) (*docdb.StartDBClusterOutput, error) {
    var output docdb.StartDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DocdbStopDBClusterResult struct {
	Result workflow.Future
}

func (r *DocdbStopDBClusterResult) Get(ctx workflow.Context) (*docdb.StopDBClusterOutput, error) {
    var output docdb.StopDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type DocDBStub struct {
    activities awsactivities.DocDBActivities
}

func NewDocDBStub() DocDBClient {
    return &DocDBStub{}
}
func (a *DocDBStub) AddTagsToResource(ctx workflow.Context, input *docdb.AddTagsToResourceInput) (*docdb.AddTagsToResourceOutput, error) {
    var output docdb.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) AddTagsToResourceAsync(ctx workflow.Context, input *docdb.AddTagsToResourceInput) *DocdbAddTagsToResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input)
    return &DocdbAddTagsToResourceResult{Result: future}
}
func (a *DocDBStub) ApplyPendingMaintenanceAction(ctx workflow.Context, input *docdb.ApplyPendingMaintenanceActionInput) (*docdb.ApplyPendingMaintenanceActionOutput, error) {
    var output docdb.ApplyPendingMaintenanceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ApplyPendingMaintenanceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *docdb.ApplyPendingMaintenanceActionInput) *DocdbApplyPendingMaintenanceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ApplyPendingMaintenanceAction, input)
    return &DocdbApplyPendingMaintenanceActionResult{Result: future}
}
func (a *DocDBStub) CopyDBClusterParameterGroup(ctx workflow.Context, input *docdb.CopyDBClusterParameterGroupInput) (*docdb.CopyDBClusterParameterGroupOutput, error) {
    var output docdb.CopyDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) CopyDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.CopyDBClusterParameterGroupInput) *DocdbCopyDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyDBClusterParameterGroup, input)
    return &DocdbCopyDBClusterParameterGroupResult{Result: future}
}
func (a *DocDBStub) CopyDBClusterSnapshot(ctx workflow.Context, input *docdb.CopyDBClusterSnapshotInput) (*docdb.CopyDBClusterSnapshotOutput, error) {
    var output docdb.CopyDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyDBClusterSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) CopyDBClusterSnapshotAsync(ctx workflow.Context, input *docdb.CopyDBClusterSnapshotInput) *DocdbCopyDBClusterSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyDBClusterSnapshot, input)
    return &DocdbCopyDBClusterSnapshotResult{Result: future}
}
func (a *DocDBStub) CreateDBCluster(ctx workflow.Context, input *docdb.CreateDBClusterInput) (*docdb.CreateDBClusterOutput, error) {
    var output docdb.CreateDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) CreateDBClusterAsync(ctx workflow.Context, input *docdb.CreateDBClusterInput) *DocdbCreateDBClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDBCluster, input)
    return &DocdbCreateDBClusterResult{Result: future}
}
func (a *DocDBStub) CreateDBClusterParameterGroup(ctx workflow.Context, input *docdb.CreateDBClusterParameterGroupInput) (*docdb.CreateDBClusterParameterGroupOutput, error) {
    var output docdb.CreateDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) CreateDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.CreateDBClusterParameterGroupInput) *DocdbCreateDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDBClusterParameterGroup, input)
    return &DocdbCreateDBClusterParameterGroupResult{Result: future}
}
func (a *DocDBStub) CreateDBClusterSnapshot(ctx workflow.Context, input *docdb.CreateDBClusterSnapshotInput) (*docdb.CreateDBClusterSnapshotOutput, error) {
    var output docdb.CreateDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBClusterSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) CreateDBClusterSnapshotAsync(ctx workflow.Context, input *docdb.CreateDBClusterSnapshotInput) *DocdbCreateDBClusterSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDBClusterSnapshot, input)
    return &DocdbCreateDBClusterSnapshotResult{Result: future}
}
func (a *DocDBStub) CreateDBInstance(ctx workflow.Context, input *docdb.CreateDBInstanceInput) (*docdb.CreateDBInstanceOutput, error) {
    var output docdb.CreateDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) CreateDBInstanceAsync(ctx workflow.Context, input *docdb.CreateDBInstanceInput) *DocdbCreateDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDBInstance, input)
    return &DocdbCreateDBInstanceResult{Result: future}
}
func (a *DocDBStub) CreateDBSubnetGroup(ctx workflow.Context, input *docdb.CreateDBSubnetGroupInput) (*docdb.CreateDBSubnetGroupOutput, error) {
    var output docdb.CreateDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) CreateDBSubnetGroupAsync(ctx workflow.Context, input *docdb.CreateDBSubnetGroupInput) *DocdbCreateDBSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDBSubnetGroup, input)
    return &DocdbCreateDBSubnetGroupResult{Result: future}
}
func (a *DocDBStub) DeleteDBCluster(ctx workflow.Context, input *docdb.DeleteDBClusterInput) (*docdb.DeleteDBClusterOutput, error) {
    var output docdb.DeleteDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DeleteDBClusterAsync(ctx workflow.Context, input *docdb.DeleteDBClusterInput) *DocdbDeleteDBClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDBCluster, input)
    return &DocdbDeleteDBClusterResult{Result: future}
}
func (a *DocDBStub) DeleteDBClusterParameterGroup(ctx workflow.Context, input *docdb.DeleteDBClusterParameterGroupInput) (*docdb.DeleteDBClusterParameterGroupOutput, error) {
    var output docdb.DeleteDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DeleteDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.DeleteDBClusterParameterGroupInput) *DocdbDeleteDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDBClusterParameterGroup, input)
    return &DocdbDeleteDBClusterParameterGroupResult{Result: future}
}
func (a *DocDBStub) DeleteDBClusterSnapshot(ctx workflow.Context, input *docdb.DeleteDBClusterSnapshotInput) (*docdb.DeleteDBClusterSnapshotOutput, error) {
    var output docdb.DeleteDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBClusterSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DeleteDBClusterSnapshotAsync(ctx workflow.Context, input *docdb.DeleteDBClusterSnapshotInput) *DocdbDeleteDBClusterSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDBClusterSnapshot, input)
    return &DocdbDeleteDBClusterSnapshotResult{Result: future}
}
func (a *DocDBStub) DeleteDBInstance(ctx workflow.Context, input *docdb.DeleteDBInstanceInput) (*docdb.DeleteDBInstanceOutput, error) {
    var output docdb.DeleteDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DeleteDBInstanceAsync(ctx workflow.Context, input *docdb.DeleteDBInstanceInput) *DocdbDeleteDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDBInstance, input)
    return &DocdbDeleteDBInstanceResult{Result: future}
}
func (a *DocDBStub) DeleteDBSubnetGroup(ctx workflow.Context, input *docdb.DeleteDBSubnetGroupInput) (*docdb.DeleteDBSubnetGroupOutput, error) {
    var output docdb.DeleteDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DeleteDBSubnetGroupAsync(ctx workflow.Context, input *docdb.DeleteDBSubnetGroupInput) *DocdbDeleteDBSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDBSubnetGroup, input)
    return &DocdbDeleteDBSubnetGroupResult{Result: future}
}
func (a *DocDBStub) DescribeCertificates(ctx workflow.Context, input *docdb.DescribeCertificatesInput) (*docdb.DescribeCertificatesOutput, error) {
    var output docdb.DescribeCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeCertificatesAsync(ctx workflow.Context, input *docdb.DescribeCertificatesInput) *DocdbDescribeCertificatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificates, input)
    return &DocdbDescribeCertificatesResult{Result: future}
}
func (a *DocDBStub) DescribeDBClusterParameterGroups(ctx workflow.Context, input *docdb.DescribeDBClusterParameterGroupsInput) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
    var output docdb.DescribeDBClusterParameterGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterParameterGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBClusterParameterGroupsAsync(ctx workflow.Context, input *docdb.DescribeDBClusterParameterGroupsInput) *DocdbDescribeDBClusterParameterGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterParameterGroups, input)
    return &DocdbDescribeDBClusterParameterGroupsResult{Result: future}
}
func (a *DocDBStub) DescribeDBClusterParameters(ctx workflow.Context, input *docdb.DescribeDBClusterParametersInput) (*docdb.DescribeDBClusterParametersOutput, error) {
    var output docdb.DescribeDBClusterParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBClusterParametersAsync(ctx workflow.Context, input *docdb.DescribeDBClusterParametersInput) *DocdbDescribeDBClusterParametersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterParameters, input)
    return &DocdbDescribeDBClusterParametersResult{Result: future}
}
func (a *DocDBStub) DescribeDBClusterSnapshotAttributes(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotAttributesInput) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
    var output docdb.DescribeDBClusterSnapshotAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterSnapshotAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBClusterSnapshotAttributesAsync(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotAttributesInput) *DocdbDescribeDBClusterSnapshotAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterSnapshotAttributes, input)
    return &DocdbDescribeDBClusterSnapshotAttributesResult{Result: future}
}
func (a *DocDBStub) DescribeDBClusterSnapshots(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotsInput) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
    var output docdb.DescribeDBClusterSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterSnapshots, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBClusterSnapshotsAsync(ctx workflow.Context, input *docdb.DescribeDBClusterSnapshotsInput) *DocdbDescribeDBClusterSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterSnapshots, input)
    return &DocdbDescribeDBClusterSnapshotsResult{Result: future}
}
func (a *DocDBStub) DescribeDBClusters(ctx workflow.Context, input *docdb.DescribeDBClustersInput) (*docdb.DescribeDBClustersOutput, error) {
    var output docdb.DescribeDBClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBClustersAsync(ctx workflow.Context, input *docdb.DescribeDBClustersInput) *DocdbDescribeDBClustersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusters, input)
    return &DocdbDescribeDBClustersResult{Result: future}
}
func (a *DocDBStub) DescribeDBEngineVersions(ctx workflow.Context, input *docdb.DescribeDBEngineVersionsInput) (*docdb.DescribeDBEngineVersionsOutput, error) {
    var output docdb.DescribeDBEngineVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBEngineVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBEngineVersionsAsync(ctx workflow.Context, input *docdb.DescribeDBEngineVersionsInput) *DocdbDescribeDBEngineVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBEngineVersions, input)
    return &DocdbDescribeDBEngineVersionsResult{Result: future}
}
func (a *DocDBStub) DescribeDBInstances(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) (*docdb.DescribeDBInstancesOutput, error) {
    var output docdb.DescribeDBInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBInstancesAsync(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) *DocdbDescribeDBInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBInstances, input)
    return &DocdbDescribeDBInstancesResult{Result: future}
}
func (a *DocDBStub) DescribeDBSubnetGroups(ctx workflow.Context, input *docdb.DescribeDBSubnetGroupsInput) (*docdb.DescribeDBSubnetGroupsOutput, error) {
    var output docdb.DescribeDBSubnetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBSubnetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeDBSubnetGroupsAsync(ctx workflow.Context, input *docdb.DescribeDBSubnetGroupsInput) *DocdbDescribeDBSubnetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDBSubnetGroups, input)
    return &DocdbDescribeDBSubnetGroupsResult{Result: future}
}
func (a *DocDBStub) DescribeEngineDefaultClusterParameters(ctx workflow.Context, input *docdb.DescribeEngineDefaultClusterParametersInput) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
    var output docdb.DescribeEngineDefaultClusterParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEngineDefaultClusterParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeEngineDefaultClusterParametersAsync(ctx workflow.Context, input *docdb.DescribeEngineDefaultClusterParametersInput) *DocdbDescribeEngineDefaultClusterParametersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEngineDefaultClusterParameters, input)
    return &DocdbDescribeEngineDefaultClusterParametersResult{Result: future}
}
func (a *DocDBStub) DescribeEventCategories(ctx workflow.Context, input *docdb.DescribeEventCategoriesInput) (*docdb.DescribeEventCategoriesOutput, error) {
    var output docdb.DescribeEventCategoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventCategories, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeEventCategoriesAsync(ctx workflow.Context, input *docdb.DescribeEventCategoriesInput) *DocdbDescribeEventCategoriesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventCategories, input)
    return &DocdbDescribeEventCategoriesResult{Result: future}
}
func (a *DocDBStub) DescribeEvents(ctx workflow.Context, input *docdb.DescribeEventsInput) (*docdb.DescribeEventsOutput, error) {
    var output docdb.DescribeEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeEventsAsync(ctx workflow.Context, input *docdb.DescribeEventsInput) *DocdbDescribeEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input)
    return &DocdbDescribeEventsResult{Result: future}
}
func (a *DocDBStub) DescribeOrderableDBInstanceOptions(ctx workflow.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
    var output docdb.DescribeOrderableDBInstanceOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrderableDBInstanceOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribeOrderableDBInstanceOptionsAsync(ctx workflow.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput) *DocdbDescribeOrderableDBInstanceOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrderableDBInstanceOptions, input)
    return &DocdbDescribeOrderableDBInstanceOptionsResult{Result: future}
}
func (a *DocDBStub) DescribePendingMaintenanceActions(ctx workflow.Context, input *docdb.DescribePendingMaintenanceActionsInput) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
    var output docdb.DescribePendingMaintenanceActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePendingMaintenanceActions, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *docdb.DescribePendingMaintenanceActionsInput) *DocdbDescribePendingMaintenanceActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePendingMaintenanceActions, input)
    return &DocdbDescribePendingMaintenanceActionsResult{Result: future}
}
func (a *DocDBStub) FailoverDBCluster(ctx workflow.Context, input *docdb.FailoverDBClusterInput) (*docdb.FailoverDBClusterOutput, error) {
    var output docdb.FailoverDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.FailoverDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) FailoverDBClusterAsync(ctx workflow.Context, input *docdb.FailoverDBClusterInput) *DocdbFailoverDBClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.FailoverDBCluster, input)
    return &DocdbFailoverDBClusterResult{Result: future}
}
func (a *DocDBStub) ListTagsForResource(ctx workflow.Context, input *docdb.ListTagsForResourceInput) (*docdb.ListTagsForResourceOutput, error) {
    var output docdb.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ListTagsForResourceAsync(ctx workflow.Context, input *docdb.ListTagsForResourceInput) *DocdbListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &DocdbListTagsForResourceResult{Result: future}
}
func (a *DocDBStub) ModifyDBCluster(ctx workflow.Context, input *docdb.ModifyDBClusterInput) (*docdb.ModifyDBClusterOutput, error) {
    var output docdb.ModifyDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ModifyDBClusterAsync(ctx workflow.Context, input *docdb.ModifyDBClusterInput) *DocdbModifyDBClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyDBCluster, input)
    return &DocdbModifyDBClusterResult{Result: future}
}
func (a *DocDBStub) ModifyDBClusterParameterGroup(ctx workflow.Context, input *docdb.ModifyDBClusterParameterGroupInput) (*docdb.ModifyDBClusterParameterGroupOutput, error) {
    var output docdb.ModifyDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ModifyDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.ModifyDBClusterParameterGroupInput) *DocdbModifyDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyDBClusterParameterGroup, input)
    return &DocdbModifyDBClusterParameterGroupResult{Result: future}
}
func (a *DocDBStub) ModifyDBClusterSnapshotAttribute(ctx workflow.Context, input *docdb.ModifyDBClusterSnapshotAttributeInput) (*docdb.ModifyDBClusterSnapshotAttributeOutput, error) {
    var output docdb.ModifyDBClusterSnapshotAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBClusterSnapshotAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ModifyDBClusterSnapshotAttributeAsync(ctx workflow.Context, input *docdb.ModifyDBClusterSnapshotAttributeInput) *DocdbModifyDBClusterSnapshotAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyDBClusterSnapshotAttribute, input)
    return &DocdbModifyDBClusterSnapshotAttributeResult{Result: future}
}
func (a *DocDBStub) ModifyDBInstance(ctx workflow.Context, input *docdb.ModifyDBInstanceInput) (*docdb.ModifyDBInstanceOutput, error) {
    var output docdb.ModifyDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ModifyDBInstanceAsync(ctx workflow.Context, input *docdb.ModifyDBInstanceInput) *DocdbModifyDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyDBInstance, input)
    return &DocdbModifyDBInstanceResult{Result: future}
}
func (a *DocDBStub) ModifyDBSubnetGroup(ctx workflow.Context, input *docdb.ModifyDBSubnetGroupInput) (*docdb.ModifyDBSubnetGroupOutput, error) {
    var output docdb.ModifyDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ModifyDBSubnetGroupAsync(ctx workflow.Context, input *docdb.ModifyDBSubnetGroupInput) *DocdbModifyDBSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyDBSubnetGroup, input)
    return &DocdbModifyDBSubnetGroupResult{Result: future}
}
func (a *DocDBStub) RebootDBInstance(ctx workflow.Context, input *docdb.RebootDBInstanceInput) (*docdb.RebootDBInstanceOutput, error) {
    var output docdb.RebootDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) RebootDBInstanceAsync(ctx workflow.Context, input *docdb.RebootDBInstanceInput) *DocdbRebootDBInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebootDBInstance, input)
    return &DocdbRebootDBInstanceResult{Result: future}
}
func (a *DocDBStub) RemoveTagsFromResource(ctx workflow.Context, input *docdb.RemoveTagsFromResourceInput) (*docdb.RemoveTagsFromResourceOutput, error) {
    var output docdb.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *docdb.RemoveTagsFromResourceInput) *DocdbRemoveTagsFromResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input)
    return &DocdbRemoveTagsFromResourceResult{Result: future}
}
func (a *DocDBStub) ResetDBClusterParameterGroup(ctx workflow.Context, input *docdb.ResetDBClusterParameterGroupInput) (*docdb.ResetDBClusterParameterGroupOutput, error) {
    var output docdb.ResetDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) ResetDBClusterParameterGroupAsync(ctx workflow.Context, input *docdb.ResetDBClusterParameterGroupInput) *DocdbResetDBClusterParameterGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetDBClusterParameterGroup, input)
    return &DocdbResetDBClusterParameterGroupResult{Result: future}
}
func (a *DocDBStub) RestoreDBClusterFromSnapshot(ctx workflow.Context, input *docdb.RestoreDBClusterFromSnapshotInput) (*docdb.RestoreDBClusterFromSnapshotOutput, error) {
    var output docdb.RestoreDBClusterFromSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBClusterFromSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) RestoreDBClusterFromSnapshotAsync(ctx workflow.Context, input *docdb.RestoreDBClusterFromSnapshotInput) *DocdbRestoreDBClusterFromSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreDBClusterFromSnapshot, input)
    return &DocdbRestoreDBClusterFromSnapshotResult{Result: future}
}
func (a *DocDBStub) RestoreDBClusterToPointInTime(ctx workflow.Context, input *docdb.RestoreDBClusterToPointInTimeInput) (*docdb.RestoreDBClusterToPointInTimeOutput, error) {
    var output docdb.RestoreDBClusterToPointInTimeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBClusterToPointInTime, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) RestoreDBClusterToPointInTimeAsync(ctx workflow.Context, input *docdb.RestoreDBClusterToPointInTimeInput) *DocdbRestoreDBClusterToPointInTimeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreDBClusterToPointInTime, input)
    return &DocdbRestoreDBClusterToPointInTimeResult{Result: future}
}
func (a *DocDBStub) StartDBCluster(ctx workflow.Context, input *docdb.StartDBClusterInput) (*docdb.StartDBClusterOutput, error) {
    var output docdb.StartDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) StartDBClusterAsync(ctx workflow.Context, input *docdb.StartDBClusterInput) *DocdbStartDBClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartDBCluster, input)
    return &DocdbStartDBClusterResult{Result: future}
}
func (a *DocDBStub) StopDBCluster(ctx workflow.Context, input *docdb.StopDBClusterInput) (*docdb.StopDBClusterOutput, error) {
    var output docdb.StopDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *DocDBStub) StopDBClusterAsync(ctx workflow.Context, input *docdb.StopDBClusterInput) *DocdbStopDBClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopDBCluster, input)
    return &DocdbStopDBClusterResult{Result: future}
}

func (a *DocDBStub) WaitUntilDBInstanceAvailable(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDBInstanceAvailable, input).Get(ctx, nil)
}

func (a *DocDBStub) WaitUntilDBInstanceAvailableAsync(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDBInstanceAvailable, input)
}

func (a *DocDBStub) WaitUntilDBInstanceDeleted(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDBInstanceDeleted, input).Get(ctx, nil)
}

func (a *DocDBStub) WaitUntilDBInstanceDeletedAsync(ctx workflow.Context, input *docdb.DescribeDBInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDBInstanceDeleted, input)
}