package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
)

type RDSActivities struct {
	client rdsiface.RDSAPI
}

func NewRDSActivities(client rdsiface.RDSAPI) *RDSActivities {
    return &RDSActivities{client: client}
}


func (a *RDSActivities) AddRoleToDBCluster(input *rds.AddRoleToDBClusterInput) (*rds.AddRoleToDBClusterOutput, error) {
    return a.client.AddRoleToDBCluster(input)
}



func (a *RDSActivities) AddRoleToDBInstance(input *rds.AddRoleToDBInstanceInput) (*rds.AddRoleToDBInstanceOutput, error) {
    return a.client.AddRoleToDBInstance(input)
}



func (a *RDSActivities) AddSourceIdentifierToSubscription(input *rds.AddSourceIdentifierToSubscriptionInput) (*rds.AddSourceIdentifierToSubscriptionOutput, error) {
    return a.client.AddSourceIdentifierToSubscription(input)
}



func (a *RDSActivities) AddTagsToResource(input *rds.AddTagsToResourceInput) (*rds.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResource(input)
}



func (a *RDSActivities) ApplyPendingMaintenanceAction(input *rds.ApplyPendingMaintenanceActionInput) (*rds.ApplyPendingMaintenanceActionOutput, error) {
    return a.client.ApplyPendingMaintenanceAction(input)
}



func (a *RDSActivities) AuthorizeDBSecurityGroupIngress(input *rds.AuthorizeDBSecurityGroupIngressInput) (*rds.AuthorizeDBSecurityGroupIngressOutput, error) {
    return a.client.AuthorizeDBSecurityGroupIngress(input)
}



func (a *RDSActivities) BacktrackDBCluster(input *rds.BacktrackDBClusterInput) (*rds.BacktrackDBClusterOutput, error) {
    return a.client.BacktrackDBCluster(input)
}



func (a *RDSActivities) CancelExportTask(input *rds.CancelExportTaskInput) (*rds.CancelExportTaskOutput, error) {
    return a.client.CancelExportTask(input)
}



func (a *RDSActivities) CopyDBClusterParameterGroup(input *rds.CopyDBClusterParameterGroupInput) (*rds.CopyDBClusterParameterGroupOutput, error) {
    return a.client.CopyDBClusterParameterGroup(input)
}



func (a *RDSActivities) CopyDBClusterSnapshot(input *rds.CopyDBClusterSnapshotInput) (*rds.CopyDBClusterSnapshotOutput, error) {
    return a.client.CopyDBClusterSnapshot(input)
}



func (a *RDSActivities) CopyDBParameterGroup(input *rds.CopyDBParameterGroupInput) (*rds.CopyDBParameterGroupOutput, error) {
    return a.client.CopyDBParameterGroup(input)
}



func (a *RDSActivities) CopyDBSnapshot(input *rds.CopyDBSnapshotInput) (*rds.CopyDBSnapshotOutput, error) {
    return a.client.CopyDBSnapshot(input)
}



func (a *RDSActivities) CopyOptionGroup(input *rds.CopyOptionGroupInput) (*rds.CopyOptionGroupOutput, error) {
    return a.client.CopyOptionGroup(input)
}



func (a *RDSActivities) CreateCustomAvailabilityZone(input *rds.CreateCustomAvailabilityZoneInput) (*rds.CreateCustomAvailabilityZoneOutput, error) {
    return a.client.CreateCustomAvailabilityZone(input)
}



func (a *RDSActivities) CreateDBCluster(input *rds.CreateDBClusterInput) (*rds.CreateDBClusterOutput, error) {
    return a.client.CreateDBCluster(input)
}



func (a *RDSActivities) CreateDBClusterEndpoint(input *rds.CreateDBClusterEndpointInput) (*rds.CreateDBClusterEndpointOutput, error) {
    return a.client.CreateDBClusterEndpoint(input)
}



func (a *RDSActivities) CreateDBClusterParameterGroup(input *rds.CreateDBClusterParameterGroupInput) (*rds.CreateDBClusterParameterGroupOutput, error) {
    return a.client.CreateDBClusterParameterGroup(input)
}



func (a *RDSActivities) CreateDBClusterSnapshot(input *rds.CreateDBClusterSnapshotInput) (*rds.CreateDBClusterSnapshotOutput, error) {
    return a.client.CreateDBClusterSnapshot(input)
}



func (a *RDSActivities) CreateDBInstance(input *rds.CreateDBInstanceInput) (*rds.CreateDBInstanceOutput, error) {
    return a.client.CreateDBInstance(input)
}



func (a *RDSActivities) CreateDBInstanceReadReplica(input *rds.CreateDBInstanceReadReplicaInput) (*rds.CreateDBInstanceReadReplicaOutput, error) {
    return a.client.CreateDBInstanceReadReplica(input)
}



func (a *RDSActivities) CreateDBParameterGroup(input *rds.CreateDBParameterGroupInput) (*rds.CreateDBParameterGroupOutput, error) {
    return a.client.CreateDBParameterGroup(input)
}



func (a *RDSActivities) CreateDBProxy(input *rds.CreateDBProxyInput) (*rds.CreateDBProxyOutput, error) {
    return a.client.CreateDBProxy(input)
}



func (a *RDSActivities) CreateDBSecurityGroup(input *rds.CreateDBSecurityGroupInput) (*rds.CreateDBSecurityGroupOutput, error) {
    return a.client.CreateDBSecurityGroup(input)
}



func (a *RDSActivities) CreateDBSnapshot(input *rds.CreateDBSnapshotInput) (*rds.CreateDBSnapshotOutput, error) {
    return a.client.CreateDBSnapshot(input)
}



func (a *RDSActivities) CreateDBSubnetGroup(input *rds.CreateDBSubnetGroupInput) (*rds.CreateDBSubnetGroupOutput, error) {
    return a.client.CreateDBSubnetGroup(input)
}



func (a *RDSActivities) CreateEventSubscription(input *rds.CreateEventSubscriptionInput) (*rds.CreateEventSubscriptionOutput, error) {
    return a.client.CreateEventSubscription(input)
}



func (a *RDSActivities) CreateGlobalCluster(input *rds.CreateGlobalClusterInput) (*rds.CreateGlobalClusterOutput, error) {
    return a.client.CreateGlobalCluster(input)
}



func (a *RDSActivities) CreateOptionGroup(input *rds.CreateOptionGroupInput) (*rds.CreateOptionGroupOutput, error) {
    return a.client.CreateOptionGroup(input)
}



func (a *RDSActivities) DeleteCustomAvailabilityZone(input *rds.DeleteCustomAvailabilityZoneInput) (*rds.DeleteCustomAvailabilityZoneOutput, error) {
    return a.client.DeleteCustomAvailabilityZone(input)
}



func (a *RDSActivities) DeleteDBCluster(input *rds.DeleteDBClusterInput) (*rds.DeleteDBClusterOutput, error) {
    return a.client.DeleteDBCluster(input)
}



func (a *RDSActivities) DeleteDBClusterEndpoint(input *rds.DeleteDBClusterEndpointInput) (*rds.DeleteDBClusterEndpointOutput, error) {
    return a.client.DeleteDBClusterEndpoint(input)
}



func (a *RDSActivities) DeleteDBClusterParameterGroup(input *rds.DeleteDBClusterParameterGroupInput) (*rds.DeleteDBClusterParameterGroupOutput, error) {
    return a.client.DeleteDBClusterParameterGroup(input)
}



func (a *RDSActivities) DeleteDBClusterSnapshot(input *rds.DeleteDBClusterSnapshotInput) (*rds.DeleteDBClusterSnapshotOutput, error) {
    return a.client.DeleteDBClusterSnapshot(input)
}



func (a *RDSActivities) DeleteDBInstance(input *rds.DeleteDBInstanceInput) (*rds.DeleteDBInstanceOutput, error) {
    return a.client.DeleteDBInstance(input)
}



func (a *RDSActivities) DeleteDBInstanceAutomatedBackup(input *rds.DeleteDBInstanceAutomatedBackupInput) (*rds.DeleteDBInstanceAutomatedBackupOutput, error) {
    return a.client.DeleteDBInstanceAutomatedBackup(input)
}



func (a *RDSActivities) DeleteDBParameterGroup(input *rds.DeleteDBParameterGroupInput) (*rds.DeleteDBParameterGroupOutput, error) {
    return a.client.DeleteDBParameterGroup(input)
}



func (a *RDSActivities) DeleteDBProxy(input *rds.DeleteDBProxyInput) (*rds.DeleteDBProxyOutput, error) {
    return a.client.DeleteDBProxy(input)
}



func (a *RDSActivities) DeleteDBSecurityGroup(input *rds.DeleteDBSecurityGroupInput) (*rds.DeleteDBSecurityGroupOutput, error) {
    return a.client.DeleteDBSecurityGroup(input)
}



func (a *RDSActivities) DeleteDBSnapshot(input *rds.DeleteDBSnapshotInput) (*rds.DeleteDBSnapshotOutput, error) {
    return a.client.DeleteDBSnapshot(input)
}



func (a *RDSActivities) DeleteDBSubnetGroup(input *rds.DeleteDBSubnetGroupInput) (*rds.DeleteDBSubnetGroupOutput, error) {
    return a.client.DeleteDBSubnetGroup(input)
}



func (a *RDSActivities) DeleteEventSubscription(input *rds.DeleteEventSubscriptionInput) (*rds.DeleteEventSubscriptionOutput, error) {
    return a.client.DeleteEventSubscription(input)
}



func (a *RDSActivities) DeleteGlobalCluster(input *rds.DeleteGlobalClusterInput) (*rds.DeleteGlobalClusterOutput, error) {
    return a.client.DeleteGlobalCluster(input)
}



func (a *RDSActivities) DeleteInstallationMedia(input *rds.DeleteInstallationMediaInput) (*rds.DeleteInstallationMediaOutput, error) {
    return a.client.DeleteInstallationMedia(input)
}



func (a *RDSActivities) DeleteOptionGroup(input *rds.DeleteOptionGroupInput) (*rds.DeleteOptionGroupOutput, error) {
    return a.client.DeleteOptionGroup(input)
}



func (a *RDSActivities) DeregisterDBProxyTargets(input *rds.DeregisterDBProxyTargetsInput) (*rds.DeregisterDBProxyTargetsOutput, error) {
    return a.client.DeregisterDBProxyTargets(input)
}



func (a *RDSActivities) DescribeAccountAttributes(input *rds.DescribeAccountAttributesInput) (*rds.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributes(input)
}



func (a *RDSActivities) DescribeCertificates(input *rds.DescribeCertificatesInput) (*rds.DescribeCertificatesOutput, error) {
    return a.client.DescribeCertificates(input)
}



func (a *RDSActivities) DescribeCustomAvailabilityZones(input *rds.DescribeCustomAvailabilityZonesInput) (*rds.DescribeCustomAvailabilityZonesOutput, error) {
    return a.client.DescribeCustomAvailabilityZones(input)
}



func (a *RDSActivities) DescribeDBClusterBacktracks(input *rds.DescribeDBClusterBacktracksInput) (*rds.DescribeDBClusterBacktracksOutput, error) {
    return a.client.DescribeDBClusterBacktracks(input)
}



func (a *RDSActivities) DescribeDBClusterEndpoints(input *rds.DescribeDBClusterEndpointsInput) (*rds.DescribeDBClusterEndpointsOutput, error) {
    return a.client.DescribeDBClusterEndpoints(input)
}



func (a *RDSActivities) DescribeDBClusterParameterGroups(input *rds.DescribeDBClusterParameterGroupsInput) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
    return a.client.DescribeDBClusterParameterGroups(input)
}



func (a *RDSActivities) DescribeDBClusterParameters(input *rds.DescribeDBClusterParametersInput) (*rds.DescribeDBClusterParametersOutput, error) {
    return a.client.DescribeDBClusterParameters(input)
}



func (a *RDSActivities) DescribeDBClusterSnapshotAttributes(input *rds.DescribeDBClusterSnapshotAttributesInput) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
    return a.client.DescribeDBClusterSnapshotAttributes(input)
}



func (a *RDSActivities) DescribeDBClusterSnapshots(input *rds.DescribeDBClusterSnapshotsInput) (*rds.DescribeDBClusterSnapshotsOutput, error) {
    return a.client.DescribeDBClusterSnapshots(input)
}



func (a *RDSActivities) DescribeDBClusters(input *rds.DescribeDBClustersInput) (*rds.DescribeDBClustersOutput, error) {
    return a.client.DescribeDBClusters(input)
}



func (a *RDSActivities) DescribeDBEngineVersions(input *rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error) {
    return a.client.DescribeDBEngineVersions(input)
}



func (a *RDSActivities) DescribeDBInstanceAutomatedBackups(input *rds.DescribeDBInstanceAutomatedBackupsInput) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error) {
    return a.client.DescribeDBInstanceAutomatedBackups(input)
}



func (a *RDSActivities) DescribeDBInstances(input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
    return a.client.DescribeDBInstances(input)
}



func (a *RDSActivities) DescribeDBLogFiles(input *rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error) {
    return a.client.DescribeDBLogFiles(input)
}



func (a *RDSActivities) DescribeDBParameterGroups(input *rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error) {
    return a.client.DescribeDBParameterGroups(input)
}



func (a *RDSActivities) DescribeDBParameters(input *rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error) {
    return a.client.DescribeDBParameters(input)
}



func (a *RDSActivities) DescribeDBProxies(input *rds.DescribeDBProxiesInput) (*rds.DescribeDBProxiesOutput, error) {
    return a.client.DescribeDBProxies(input)
}



func (a *RDSActivities) DescribeDBProxyTargetGroups(input *rds.DescribeDBProxyTargetGroupsInput) (*rds.DescribeDBProxyTargetGroupsOutput, error) {
    return a.client.DescribeDBProxyTargetGroups(input)
}



func (a *RDSActivities) DescribeDBProxyTargets(input *rds.DescribeDBProxyTargetsInput) (*rds.DescribeDBProxyTargetsOutput, error) {
    return a.client.DescribeDBProxyTargets(input)
}



func (a *RDSActivities) DescribeDBSecurityGroups(input *rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error) {
    return a.client.DescribeDBSecurityGroups(input)
}



func (a *RDSActivities) DescribeDBSnapshotAttributes(input *rds.DescribeDBSnapshotAttributesInput) (*rds.DescribeDBSnapshotAttributesOutput, error) {
    return a.client.DescribeDBSnapshotAttributes(input)
}



func (a *RDSActivities) DescribeDBSnapshots(input *rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error) {
    return a.client.DescribeDBSnapshots(input)
}



func (a *RDSActivities) DescribeDBSubnetGroups(input *rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error) {
    return a.client.DescribeDBSubnetGroups(input)
}



func (a *RDSActivities) DescribeEngineDefaultClusterParameters(input *rds.DescribeEngineDefaultClusterParametersInput) (*rds.DescribeEngineDefaultClusterParametersOutput, error) {
    return a.client.DescribeEngineDefaultClusterParameters(input)
}



func (a *RDSActivities) DescribeEngineDefaultParameters(input *rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error) {
    return a.client.DescribeEngineDefaultParameters(input)
}



func (a *RDSActivities) DescribeEventCategories(input *rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error) {
    return a.client.DescribeEventCategories(input)
}



func (a *RDSActivities) DescribeEventSubscriptions(input *rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error) {
    return a.client.DescribeEventSubscriptions(input)
}



func (a *RDSActivities) DescribeEvents(input *rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}



func (a *RDSActivities) DescribeExportTasks(input *rds.DescribeExportTasksInput) (*rds.DescribeExportTasksOutput, error) {
    return a.client.DescribeExportTasks(input)
}



func (a *RDSActivities) DescribeGlobalClusters(input *rds.DescribeGlobalClustersInput) (*rds.DescribeGlobalClustersOutput, error) {
    return a.client.DescribeGlobalClusters(input)
}



func (a *RDSActivities) DescribeInstallationMedia(input *rds.DescribeInstallationMediaInput) (*rds.DescribeInstallationMediaOutput, error) {
    return a.client.DescribeInstallationMedia(input)
}



func (a *RDSActivities) DescribeOptionGroupOptions(input *rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error) {
    return a.client.DescribeOptionGroupOptions(input)
}



func (a *RDSActivities) DescribeOptionGroups(input *rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error) {
    return a.client.DescribeOptionGroups(input)
}



func (a *RDSActivities) DescribeOrderableDBInstanceOptions(input *rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error) {
    return a.client.DescribeOrderableDBInstanceOptions(input)
}



func (a *RDSActivities) DescribePendingMaintenanceActions(input *rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error) {
    return a.client.DescribePendingMaintenanceActions(input)
}



func (a *RDSActivities) DescribeReservedDBInstances(input *rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error) {
    return a.client.DescribeReservedDBInstances(input)
}



func (a *RDSActivities) DescribeReservedDBInstancesOfferings(input *rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error) {
    return a.client.DescribeReservedDBInstancesOfferings(input)
}



func (a *RDSActivities) DescribeSourceRegions(input *rds.DescribeSourceRegionsInput) (*rds.DescribeSourceRegionsOutput, error) {
    return a.client.DescribeSourceRegions(input)
}



func (a *RDSActivities) DescribeValidDBInstanceModifications(input *rds.DescribeValidDBInstanceModificationsInput) (*rds.DescribeValidDBInstanceModificationsOutput, error) {
    return a.client.DescribeValidDBInstanceModifications(input)
}



func (a *RDSActivities) DownloadDBLogFilePortion(input *rds.DownloadDBLogFilePortionInput) (*rds.DownloadDBLogFilePortionOutput, error) {
    return a.client.DownloadDBLogFilePortion(input)
}



func (a *RDSActivities) FailoverDBCluster(input *rds.FailoverDBClusterInput) (*rds.FailoverDBClusterOutput, error) {
    return a.client.FailoverDBCluster(input)
}



func (a *RDSActivities) ImportInstallationMedia(input *rds.ImportInstallationMediaInput) (*rds.ImportInstallationMediaOutput, error) {
    return a.client.ImportInstallationMedia(input)
}



func (a *RDSActivities) ListTagsForResource(input *rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *RDSActivities) ModifyCertificates(input *rds.ModifyCertificatesInput) (*rds.ModifyCertificatesOutput, error) {
    return a.client.ModifyCertificates(input)
}



func (a *RDSActivities) ModifyCurrentDBClusterCapacity(input *rds.ModifyCurrentDBClusterCapacityInput) (*rds.ModifyCurrentDBClusterCapacityOutput, error) {
    return a.client.ModifyCurrentDBClusterCapacity(input)
}



func (a *RDSActivities) ModifyDBCluster(input *rds.ModifyDBClusterInput) (*rds.ModifyDBClusterOutput, error) {
    return a.client.ModifyDBCluster(input)
}



func (a *RDSActivities) ModifyDBClusterEndpoint(input *rds.ModifyDBClusterEndpointInput) (*rds.ModifyDBClusterEndpointOutput, error) {
    return a.client.ModifyDBClusterEndpoint(input)
}



func (a *RDSActivities) ModifyDBClusterParameterGroup(input *rds.ModifyDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error) {
    return a.client.ModifyDBClusterParameterGroup(input)
}



func (a *RDSActivities) ModifyDBClusterSnapshotAttribute(input *rds.ModifyDBClusterSnapshotAttributeInput) (*rds.ModifyDBClusterSnapshotAttributeOutput, error) {
    return a.client.ModifyDBClusterSnapshotAttribute(input)
}



func (a *RDSActivities) ModifyDBInstance(input *rds.ModifyDBInstanceInput) (*rds.ModifyDBInstanceOutput, error) {
    return a.client.ModifyDBInstance(input)
}



func (a *RDSActivities) ModifyDBParameterGroup(input *rds.ModifyDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error) {
    return a.client.ModifyDBParameterGroup(input)
}



func (a *RDSActivities) ModifyDBProxy(input *rds.ModifyDBProxyInput) (*rds.ModifyDBProxyOutput, error) {
    return a.client.ModifyDBProxy(input)
}



func (a *RDSActivities) ModifyDBProxyTargetGroup(input *rds.ModifyDBProxyTargetGroupInput) (*rds.ModifyDBProxyTargetGroupOutput, error) {
    return a.client.ModifyDBProxyTargetGroup(input)
}



func (a *RDSActivities) ModifyDBSnapshot(input *rds.ModifyDBSnapshotInput) (*rds.ModifyDBSnapshotOutput, error) {
    return a.client.ModifyDBSnapshot(input)
}



func (a *RDSActivities) ModifyDBSnapshotAttribute(input *rds.ModifyDBSnapshotAttributeInput) (*rds.ModifyDBSnapshotAttributeOutput, error) {
    return a.client.ModifyDBSnapshotAttribute(input)
}



func (a *RDSActivities) ModifyDBSubnetGroup(input *rds.ModifyDBSubnetGroupInput) (*rds.ModifyDBSubnetGroupOutput, error) {
    return a.client.ModifyDBSubnetGroup(input)
}



func (a *RDSActivities) ModifyEventSubscription(input *rds.ModifyEventSubscriptionInput) (*rds.ModifyEventSubscriptionOutput, error) {
    return a.client.ModifyEventSubscription(input)
}



func (a *RDSActivities) ModifyGlobalCluster(input *rds.ModifyGlobalClusterInput) (*rds.ModifyGlobalClusterOutput, error) {
    return a.client.ModifyGlobalCluster(input)
}



func (a *RDSActivities) ModifyOptionGroup(input *rds.ModifyOptionGroupInput) (*rds.ModifyOptionGroupOutput, error) {
    return a.client.ModifyOptionGroup(input)
}



func (a *RDSActivities) PromoteReadReplica(input *rds.PromoteReadReplicaInput) (*rds.PromoteReadReplicaOutput, error) {
    return a.client.PromoteReadReplica(input)
}



func (a *RDSActivities) PromoteReadReplicaDBCluster(input *rds.PromoteReadReplicaDBClusterInput) (*rds.PromoteReadReplicaDBClusterOutput, error) {
    return a.client.PromoteReadReplicaDBCluster(input)
}



func (a *RDSActivities) PurchaseReservedDBInstancesOffering(input *rds.PurchaseReservedDBInstancesOfferingInput) (*rds.PurchaseReservedDBInstancesOfferingOutput, error) {
    return a.client.PurchaseReservedDBInstancesOffering(input)
}



func (a *RDSActivities) RebootDBInstance(input *rds.RebootDBInstanceInput) (*rds.RebootDBInstanceOutput, error) {
    return a.client.RebootDBInstance(input)
}



func (a *RDSActivities) RegisterDBProxyTargets(input *rds.RegisterDBProxyTargetsInput) (*rds.RegisterDBProxyTargetsOutput, error) {
    return a.client.RegisterDBProxyTargets(input)
}



func (a *RDSActivities) RemoveFromGlobalCluster(input *rds.RemoveFromGlobalClusterInput) (*rds.RemoveFromGlobalClusterOutput, error) {
    return a.client.RemoveFromGlobalCluster(input)
}



func (a *RDSActivities) RemoveRoleFromDBCluster(input *rds.RemoveRoleFromDBClusterInput) (*rds.RemoveRoleFromDBClusterOutput, error) {
    return a.client.RemoveRoleFromDBCluster(input)
}



func (a *RDSActivities) RemoveRoleFromDBInstance(input *rds.RemoveRoleFromDBInstanceInput) (*rds.RemoveRoleFromDBInstanceOutput, error) {
    return a.client.RemoveRoleFromDBInstance(input)
}



func (a *RDSActivities) RemoveSourceIdentifierFromSubscription(input *rds.RemoveSourceIdentifierFromSubscriptionInput) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error) {
    return a.client.RemoveSourceIdentifierFromSubscription(input)
}



func (a *RDSActivities) RemoveTagsFromResource(input *rds.RemoveTagsFromResourceInput) (*rds.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResource(input)
}



func (a *RDSActivities) ResetDBClusterParameterGroup(input *rds.ResetDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error) {
    return a.client.ResetDBClusterParameterGroup(input)
}



func (a *RDSActivities) ResetDBParameterGroup(input *rds.ResetDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error) {
    return a.client.ResetDBParameterGroup(input)
}



func (a *RDSActivities) RestoreDBClusterFromS3(input *rds.RestoreDBClusterFromS3Input) (*rds.RestoreDBClusterFromS3Output, error) {
    return a.client.RestoreDBClusterFromS3(input)
}



func (a *RDSActivities) RestoreDBClusterFromSnapshot(input *rds.RestoreDBClusterFromSnapshotInput) (*rds.RestoreDBClusterFromSnapshotOutput, error) {
    return a.client.RestoreDBClusterFromSnapshot(input)
}



func (a *RDSActivities) RestoreDBClusterToPointInTime(input *rds.RestoreDBClusterToPointInTimeInput) (*rds.RestoreDBClusterToPointInTimeOutput, error) {
    return a.client.RestoreDBClusterToPointInTime(input)
}



func (a *RDSActivities) RestoreDBInstanceFromDBSnapshot(input *rds.RestoreDBInstanceFromDBSnapshotInput) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error) {
    return a.client.RestoreDBInstanceFromDBSnapshot(input)
}



func (a *RDSActivities) RestoreDBInstanceFromS3(input *rds.RestoreDBInstanceFromS3Input) (*rds.RestoreDBInstanceFromS3Output, error) {
    return a.client.RestoreDBInstanceFromS3(input)
}



func (a *RDSActivities) RestoreDBInstanceToPointInTime(input *rds.RestoreDBInstanceToPointInTimeInput) (*rds.RestoreDBInstanceToPointInTimeOutput, error) {
    return a.client.RestoreDBInstanceToPointInTime(input)
}



func (a *RDSActivities) RevokeDBSecurityGroupIngress(input *rds.RevokeDBSecurityGroupIngressInput) (*rds.RevokeDBSecurityGroupIngressOutput, error) {
    return a.client.RevokeDBSecurityGroupIngress(input)
}



func (a *RDSActivities) StartActivityStream(input *rds.StartActivityStreamInput) (*rds.StartActivityStreamOutput, error) {
    return a.client.StartActivityStream(input)
}



func (a *RDSActivities) StartDBCluster(input *rds.StartDBClusterInput) (*rds.StartDBClusterOutput, error) {
    return a.client.StartDBCluster(input)
}



func (a *RDSActivities) StartDBInstance(input *rds.StartDBInstanceInput) (*rds.StartDBInstanceOutput, error) {
    return a.client.StartDBInstance(input)
}



func (a *RDSActivities) StartExportTask(input *rds.StartExportTaskInput) (*rds.StartExportTaskOutput, error) {
    return a.client.StartExportTask(input)
}



func (a *RDSActivities) StopActivityStream(input *rds.StopActivityStreamInput) (*rds.StopActivityStreamOutput, error) {
    return a.client.StopActivityStream(input)
}



func (a *RDSActivities) StopDBCluster(input *rds.StopDBClusterInput) (*rds.StopDBClusterOutput, error) {
    return a.client.StopDBCluster(input)
}



func (a *RDSActivities) StopDBInstance(input *rds.StopDBInstanceInput) (*rds.StopDBInstanceOutput, error) {
    return a.client.StopDBInstance(input)
}



func (a *RDSActivities) WaitUntilDBClusterSnapshotAvailable(input *rds.DescribeDBClusterSnapshotsInput) error {
    return a.client.WaitUntilDBClusterSnapshotAvailable(input)
}



func (a *RDSActivities) WaitUntilDBClusterSnapshotDeleted(input *rds.DescribeDBClusterSnapshotsInput) error {
    return a.client.WaitUntilDBClusterSnapshotDeleted(input)
}



func (a *RDSActivities) WaitUntilDBInstanceAvailable(input *rds.DescribeDBInstancesInput) error {
    return a.client.WaitUntilDBInstanceAvailable(input)
}



func (a *RDSActivities) WaitUntilDBInstanceDeleted(input *rds.DescribeDBInstancesInput) error {
    return a.client.WaitUntilDBInstanceDeleted(input)
}



func (a *RDSActivities) WaitUntilDBSnapshotAvailable(input *rds.DescribeDBSnapshotsInput) error {
    return a.client.WaitUntilDBSnapshotAvailable(input)
}



func (a *RDSActivities) WaitUntilDBSnapshotDeleted(input *rds.DescribeDBSnapshotsInput) error {
    return a.client.WaitUntilDBSnapshotDeleted(input)
}

