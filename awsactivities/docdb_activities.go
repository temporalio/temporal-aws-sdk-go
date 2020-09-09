
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/docdb/docdbiface"
)

type DocDBActivities struct {
    client docdbiface.DocDBAPI
}

func NewDocDBActivities(session *session.Session, config ...*aws.Config) *DocDBActivities {
    client := docdb.New(session, config...)
    return &DocDBActivities{client: client}
}

func (a *DocDBActivities) AddTagsToResource(input *docdb.AddTagsToResourceInput) (*docdb.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResource(input)
}

func (a *DocDBActivities) ApplyPendingMaintenanceAction(input *docdb.ApplyPendingMaintenanceActionInput) (*docdb.ApplyPendingMaintenanceActionOutput, error) {
    return a.client.ApplyPendingMaintenanceAction(input)
}

func (a *DocDBActivities) CopyDBClusterParameterGroup(input *docdb.CopyDBClusterParameterGroupInput) (*docdb.CopyDBClusterParameterGroupOutput, error) {
    return a.client.CopyDBClusterParameterGroup(input)
}

func (a *DocDBActivities) CopyDBClusterSnapshot(input *docdb.CopyDBClusterSnapshotInput) (*docdb.CopyDBClusterSnapshotOutput, error) {
    return a.client.CopyDBClusterSnapshot(input)
}

func (a *DocDBActivities) CreateDBCluster(input *docdb.CreateDBClusterInput) (*docdb.CreateDBClusterOutput, error) {
    return a.client.CreateDBCluster(input)
}

func (a *DocDBActivities) CreateDBClusterParameterGroup(input *docdb.CreateDBClusterParameterGroupInput) (*docdb.CreateDBClusterParameterGroupOutput, error) {
    return a.client.CreateDBClusterParameterGroup(input)
}

func (a *DocDBActivities) CreateDBClusterSnapshot(input *docdb.CreateDBClusterSnapshotInput) (*docdb.CreateDBClusterSnapshotOutput, error) {
    return a.client.CreateDBClusterSnapshot(input)
}

func (a *DocDBActivities) CreateDBInstance(input *docdb.CreateDBInstanceInput) (*docdb.CreateDBInstanceOutput, error) {
    return a.client.CreateDBInstance(input)
}

func (a *DocDBActivities) CreateDBSubnetGroup(input *docdb.CreateDBSubnetGroupInput) (*docdb.CreateDBSubnetGroupOutput, error) {
    return a.client.CreateDBSubnetGroup(input)
}

func (a *DocDBActivities) DeleteDBCluster(input *docdb.DeleteDBClusterInput) (*docdb.DeleteDBClusterOutput, error) {
    return a.client.DeleteDBCluster(input)
}

func (a *DocDBActivities) DeleteDBClusterParameterGroup(input *docdb.DeleteDBClusterParameterGroupInput) (*docdb.DeleteDBClusterParameterGroupOutput, error) {
    return a.client.DeleteDBClusterParameterGroup(input)
}

func (a *DocDBActivities) DeleteDBClusterSnapshot(input *docdb.DeleteDBClusterSnapshotInput) (*docdb.DeleteDBClusterSnapshotOutput, error) {
    return a.client.DeleteDBClusterSnapshot(input)
}

func (a *DocDBActivities) DeleteDBInstance(input *docdb.DeleteDBInstanceInput) (*docdb.DeleteDBInstanceOutput, error) {
    return a.client.DeleteDBInstance(input)
}

func (a *DocDBActivities) DeleteDBSubnetGroup(input *docdb.DeleteDBSubnetGroupInput) (*docdb.DeleteDBSubnetGroupOutput, error) {
    return a.client.DeleteDBSubnetGroup(input)
}

func (a *DocDBActivities) DescribeCertificates(input *docdb.DescribeCertificatesInput) (*docdb.DescribeCertificatesOutput, error) {
    return a.client.DescribeCertificates(input)
}

func (a *DocDBActivities) DescribeDBClusterParameterGroups(input *docdb.DescribeDBClusterParameterGroupsInput) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
    return a.client.DescribeDBClusterParameterGroups(input)
}

func (a *DocDBActivities) DescribeDBClusterParameters(input *docdb.DescribeDBClusterParametersInput) (*docdb.DescribeDBClusterParametersOutput, error) {
    return a.client.DescribeDBClusterParameters(input)
}

func (a *DocDBActivities) DescribeDBClusterSnapshotAttributes(input *docdb.DescribeDBClusterSnapshotAttributesInput) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
    return a.client.DescribeDBClusterSnapshotAttributes(input)
}

func (a *DocDBActivities) DescribeDBClusterSnapshots(input *docdb.DescribeDBClusterSnapshotsInput) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
    return a.client.DescribeDBClusterSnapshots(input)
}

func (a *DocDBActivities) DescribeDBClusters(input *docdb.DescribeDBClustersInput) (*docdb.DescribeDBClustersOutput, error) {
    return a.client.DescribeDBClusters(input)
}

func (a *DocDBActivities) DescribeDBEngineVersions(input *docdb.DescribeDBEngineVersionsInput) (*docdb.DescribeDBEngineVersionsOutput, error) {
    return a.client.DescribeDBEngineVersions(input)
}

func (a *DocDBActivities) DescribeDBInstances(input *docdb.DescribeDBInstancesInput) (*docdb.DescribeDBInstancesOutput, error) {
    return a.client.DescribeDBInstances(input)
}

func (a *DocDBActivities) DescribeDBSubnetGroups(input *docdb.DescribeDBSubnetGroupsInput) (*docdb.DescribeDBSubnetGroupsOutput, error) {
    return a.client.DescribeDBSubnetGroups(input)
}

func (a *DocDBActivities) DescribeEngineDefaultClusterParameters(input *docdb.DescribeEngineDefaultClusterParametersInput) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
    return a.client.DescribeEngineDefaultClusterParameters(input)
}

func (a *DocDBActivities) DescribeEventCategories(input *docdb.DescribeEventCategoriesInput) (*docdb.DescribeEventCategoriesOutput, error) {
    return a.client.DescribeEventCategories(input)
}

func (a *DocDBActivities) DescribeEvents(input *docdb.DescribeEventsInput) (*docdb.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}

func (a *DocDBActivities) DescribeOrderableDBInstanceOptions(input *docdb.DescribeOrderableDBInstanceOptionsInput) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
    return a.client.DescribeOrderableDBInstanceOptions(input)
}

func (a *DocDBActivities) DescribePendingMaintenanceActions(input *docdb.DescribePendingMaintenanceActionsInput) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
    return a.client.DescribePendingMaintenanceActions(input)
}

func (a *DocDBActivities) FailoverDBCluster(input *docdb.FailoverDBClusterInput) (*docdb.FailoverDBClusterOutput, error) {
    return a.client.FailoverDBCluster(input)
}

func (a *DocDBActivities) ListTagsForResource(input *docdb.ListTagsForResourceInput) (*docdb.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *DocDBActivities) ModifyDBCluster(input *docdb.ModifyDBClusterInput) (*docdb.ModifyDBClusterOutput, error) {
    return a.client.ModifyDBCluster(input)
}

func (a *DocDBActivities) ModifyDBClusterParameterGroup(input *docdb.ModifyDBClusterParameterGroupInput) (*docdb.ModifyDBClusterParameterGroupOutput, error) {
    return a.client.ModifyDBClusterParameterGroup(input)
}

func (a *DocDBActivities) ModifyDBClusterSnapshotAttribute(input *docdb.ModifyDBClusterSnapshotAttributeInput) (*docdb.ModifyDBClusterSnapshotAttributeOutput, error) {
    return a.client.ModifyDBClusterSnapshotAttribute(input)
}

func (a *DocDBActivities) ModifyDBInstance(input *docdb.ModifyDBInstanceInput) (*docdb.ModifyDBInstanceOutput, error) {
    return a.client.ModifyDBInstance(input)
}

func (a *DocDBActivities) ModifyDBSubnetGroup(input *docdb.ModifyDBSubnetGroupInput) (*docdb.ModifyDBSubnetGroupOutput, error) {
    return a.client.ModifyDBSubnetGroup(input)
}

func (a *DocDBActivities) RebootDBInstance(input *docdb.RebootDBInstanceInput) (*docdb.RebootDBInstanceOutput, error) {
    return a.client.RebootDBInstance(input)
}

func (a *DocDBActivities) RemoveTagsFromResource(input *docdb.RemoveTagsFromResourceInput) (*docdb.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResource(input)
}

func (a *DocDBActivities) ResetDBClusterParameterGroup(input *docdb.ResetDBClusterParameterGroupInput) (*docdb.ResetDBClusterParameterGroupOutput, error) {
    return a.client.ResetDBClusterParameterGroup(input)
}

func (a *DocDBActivities) RestoreDBClusterFromSnapshot(input *docdb.RestoreDBClusterFromSnapshotInput) (*docdb.RestoreDBClusterFromSnapshotOutput, error) {
    return a.client.RestoreDBClusterFromSnapshot(input)
}

func (a *DocDBActivities) RestoreDBClusterToPointInTime(input *docdb.RestoreDBClusterToPointInTimeInput) (*docdb.RestoreDBClusterToPointInTimeOutput, error) {
    return a.client.RestoreDBClusterToPointInTime(input)
}

func (a *DocDBActivities) StartDBCluster(input *docdb.StartDBClusterInput) (*docdb.StartDBClusterOutput, error) {
    return a.client.StartDBCluster(input)
}

func (a *DocDBActivities) StopDBCluster(input *docdb.StopDBClusterInput) (*docdb.StopDBClusterOutput, error) {
    return a.client.StopDBCluster(input)
}

func (a *DocDBActivities) WaitUntilDBInstanceAvailable(input *docdb.DescribeDBInstancesInput) error {
    return a.client.WaitUntilDBInstanceAvailable(input)
}

func (a *DocDBActivities) WaitUntilDBInstanceDeleted(input *docdb.DescribeDBInstancesInput) error {
    return a.client.WaitUntilDBInstanceDeleted(input)
}
