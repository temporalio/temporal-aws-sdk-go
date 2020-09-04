package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/redshift/redshiftiface"
)

type RedshiftActivities struct {
	client redshiftiface.RedshiftAPI
}

func NewRedshiftActivities(client redshiftiface.RedshiftAPI) *RedshiftActivities {
    return &RedshiftActivities{client: client}
}


func (a *RedshiftActivities) AcceptReservedNodeExchange(input *redshift.AcceptReservedNodeExchangeInput) (*redshift.AcceptReservedNodeExchangeOutput, error) {
    return a.client.AcceptReservedNodeExchange(input)
}



func (a *RedshiftActivities) AuthorizeClusterSecurityGroupIngress(input *redshift.AuthorizeClusterSecurityGroupIngressInput) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error) {
    return a.client.AuthorizeClusterSecurityGroupIngress(input)
}



func (a *RedshiftActivities) AuthorizeSnapshotAccess(input *redshift.AuthorizeSnapshotAccessInput) (*redshift.AuthorizeSnapshotAccessOutput, error) {
    return a.client.AuthorizeSnapshotAccess(input)
}



func (a *RedshiftActivities) BatchDeleteClusterSnapshots(input *redshift.BatchDeleteClusterSnapshotsInput) (*redshift.BatchDeleteClusterSnapshotsOutput, error) {
    return a.client.BatchDeleteClusterSnapshots(input)
}



func (a *RedshiftActivities) BatchModifyClusterSnapshots(input *redshift.BatchModifyClusterSnapshotsInput) (*redshift.BatchModifyClusterSnapshotsOutput, error) {
    return a.client.BatchModifyClusterSnapshots(input)
}



func (a *RedshiftActivities) CancelResize(input *redshift.CancelResizeInput) (*redshift.CancelResizeOutput, error) {
    return a.client.CancelResize(input)
}



func (a *RedshiftActivities) CopyClusterSnapshot(input *redshift.CopyClusterSnapshotInput) (*redshift.CopyClusterSnapshotOutput, error) {
    return a.client.CopyClusterSnapshot(input)
}



func (a *RedshiftActivities) CreateCluster(input *redshift.CreateClusterInput) (*redshift.CreateClusterOutput, error) {
    return a.client.CreateCluster(input)
}



func (a *RedshiftActivities) CreateClusterParameterGroup(input *redshift.CreateClusterParameterGroupInput) (*redshift.CreateClusterParameterGroupOutput, error) {
    return a.client.CreateClusterParameterGroup(input)
}



func (a *RedshiftActivities) CreateClusterSecurityGroup(input *redshift.CreateClusterSecurityGroupInput) (*redshift.CreateClusterSecurityGroupOutput, error) {
    return a.client.CreateClusterSecurityGroup(input)
}



func (a *RedshiftActivities) CreateClusterSnapshot(input *redshift.CreateClusterSnapshotInput) (*redshift.CreateClusterSnapshotOutput, error) {
    return a.client.CreateClusterSnapshot(input)
}



func (a *RedshiftActivities) CreateClusterSubnetGroup(input *redshift.CreateClusterSubnetGroupInput) (*redshift.CreateClusterSubnetGroupOutput, error) {
    return a.client.CreateClusterSubnetGroup(input)
}



func (a *RedshiftActivities) CreateEventSubscription(input *redshift.CreateEventSubscriptionInput) (*redshift.CreateEventSubscriptionOutput, error) {
    return a.client.CreateEventSubscription(input)
}



func (a *RedshiftActivities) CreateHsmClientCertificate(input *redshift.CreateHsmClientCertificateInput) (*redshift.CreateHsmClientCertificateOutput, error) {
    return a.client.CreateHsmClientCertificate(input)
}



func (a *RedshiftActivities) CreateHsmConfiguration(input *redshift.CreateHsmConfigurationInput) (*redshift.CreateHsmConfigurationOutput, error) {
    return a.client.CreateHsmConfiguration(input)
}



func (a *RedshiftActivities) CreateScheduledAction(input *redshift.CreateScheduledActionInput) (*redshift.CreateScheduledActionOutput, error) {
    return a.client.CreateScheduledAction(input)
}



func (a *RedshiftActivities) CreateSnapshotCopyGrant(input *redshift.CreateSnapshotCopyGrantInput) (*redshift.CreateSnapshotCopyGrantOutput, error) {
    return a.client.CreateSnapshotCopyGrant(input)
}



func (a *RedshiftActivities) CreateSnapshotSchedule(input *redshift.CreateSnapshotScheduleInput) (*redshift.CreateSnapshotScheduleOutput, error) {
    return a.client.CreateSnapshotSchedule(input)
}



func (a *RedshiftActivities) CreateTags(input *redshift.CreateTagsInput) (*redshift.CreateTagsOutput, error) {
    return a.client.CreateTags(input)
}



func (a *RedshiftActivities) CreateUsageLimit(input *redshift.CreateUsageLimitInput) (*redshift.CreateUsageLimitOutput, error) {
    return a.client.CreateUsageLimit(input)
}



func (a *RedshiftActivities) DeleteCluster(input *redshift.DeleteClusterInput) (*redshift.DeleteClusterOutput, error) {
    return a.client.DeleteCluster(input)
}



func (a *RedshiftActivities) DeleteClusterParameterGroup(input *redshift.DeleteClusterParameterGroupInput) (*redshift.DeleteClusterParameterGroupOutput, error) {
    return a.client.DeleteClusterParameterGroup(input)
}



func (a *RedshiftActivities) DeleteClusterSecurityGroup(input *redshift.DeleteClusterSecurityGroupInput) (*redshift.DeleteClusterSecurityGroupOutput, error) {
    return a.client.DeleteClusterSecurityGroup(input)
}



func (a *RedshiftActivities) DeleteClusterSnapshot(input *redshift.DeleteClusterSnapshotInput) (*redshift.DeleteClusterSnapshotOutput, error) {
    return a.client.DeleteClusterSnapshot(input)
}



func (a *RedshiftActivities) DeleteClusterSubnetGroup(input *redshift.DeleteClusterSubnetGroupInput) (*redshift.DeleteClusterSubnetGroupOutput, error) {
    return a.client.DeleteClusterSubnetGroup(input)
}



func (a *RedshiftActivities) DeleteEventSubscription(input *redshift.DeleteEventSubscriptionInput) (*redshift.DeleteEventSubscriptionOutput, error) {
    return a.client.DeleteEventSubscription(input)
}



func (a *RedshiftActivities) DeleteHsmClientCertificate(input *redshift.DeleteHsmClientCertificateInput) (*redshift.DeleteHsmClientCertificateOutput, error) {
    return a.client.DeleteHsmClientCertificate(input)
}



func (a *RedshiftActivities) DeleteHsmConfiguration(input *redshift.DeleteHsmConfigurationInput) (*redshift.DeleteHsmConfigurationOutput, error) {
    return a.client.DeleteHsmConfiguration(input)
}



func (a *RedshiftActivities) DeleteScheduledAction(input *redshift.DeleteScheduledActionInput) (*redshift.DeleteScheduledActionOutput, error) {
    return a.client.DeleteScheduledAction(input)
}



func (a *RedshiftActivities) DeleteSnapshotCopyGrant(input *redshift.DeleteSnapshotCopyGrantInput) (*redshift.DeleteSnapshotCopyGrantOutput, error) {
    return a.client.DeleteSnapshotCopyGrant(input)
}



func (a *RedshiftActivities) DeleteSnapshotSchedule(input *redshift.DeleteSnapshotScheduleInput) (*redshift.DeleteSnapshotScheduleOutput, error) {
    return a.client.DeleteSnapshotSchedule(input)
}



func (a *RedshiftActivities) DeleteTags(input *redshift.DeleteTagsInput) (*redshift.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}



func (a *RedshiftActivities) DeleteUsageLimit(input *redshift.DeleteUsageLimitInput) (*redshift.DeleteUsageLimitOutput, error) {
    return a.client.DeleteUsageLimit(input)
}



func (a *RedshiftActivities) DescribeAccountAttributes(input *redshift.DescribeAccountAttributesInput) (*redshift.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributes(input)
}



func (a *RedshiftActivities) DescribeClusterDbRevisions(input *redshift.DescribeClusterDbRevisionsInput) (*redshift.DescribeClusterDbRevisionsOutput, error) {
    return a.client.DescribeClusterDbRevisions(input)
}



func (a *RedshiftActivities) DescribeClusterParameterGroups(input *redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error) {
    return a.client.DescribeClusterParameterGroups(input)
}



func (a *RedshiftActivities) DescribeClusterParameters(input *redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error) {
    return a.client.DescribeClusterParameters(input)
}



func (a *RedshiftActivities) DescribeClusterSecurityGroups(input *redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error) {
    return a.client.DescribeClusterSecurityGroups(input)
}



func (a *RedshiftActivities) DescribeClusterSnapshots(input *redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error) {
    return a.client.DescribeClusterSnapshots(input)
}



func (a *RedshiftActivities) DescribeClusterSubnetGroups(input *redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
    return a.client.DescribeClusterSubnetGroups(input)
}



func (a *RedshiftActivities) DescribeClusterTracks(input *redshift.DescribeClusterTracksInput) (*redshift.DescribeClusterTracksOutput, error) {
    return a.client.DescribeClusterTracks(input)
}



func (a *RedshiftActivities) DescribeClusterVersions(input *redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error) {
    return a.client.DescribeClusterVersions(input)
}



func (a *RedshiftActivities) DescribeClusters(input *redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error) {
    return a.client.DescribeClusters(input)
}



func (a *RedshiftActivities) DescribeDefaultClusterParameters(input *redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error) {
    return a.client.DescribeDefaultClusterParameters(input)
}



func (a *RedshiftActivities) DescribeEventCategories(input *redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error) {
    return a.client.DescribeEventCategories(input)
}



func (a *RedshiftActivities) DescribeEventSubscriptions(input *redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error) {
    return a.client.DescribeEventSubscriptions(input)
}



func (a *RedshiftActivities) DescribeEvents(input *redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}



func (a *RedshiftActivities) DescribeHsmClientCertificates(input *redshift.DescribeHsmClientCertificatesInput) (*redshift.DescribeHsmClientCertificatesOutput, error) {
    return a.client.DescribeHsmClientCertificates(input)
}



func (a *RedshiftActivities) DescribeHsmConfigurations(input *redshift.DescribeHsmConfigurationsInput) (*redshift.DescribeHsmConfigurationsOutput, error) {
    return a.client.DescribeHsmConfigurations(input)
}



func (a *RedshiftActivities) DescribeLoggingStatus(input *redshift.DescribeLoggingStatusInput) (*redshift.LoggingStatus, error) {
    return a.client.DescribeLoggingStatus(input)
}



func (a *RedshiftActivities) DescribeNodeConfigurationOptions(input *redshift.DescribeNodeConfigurationOptionsInput) (*redshift.DescribeNodeConfigurationOptionsOutput, error) {
    return a.client.DescribeNodeConfigurationOptions(input)
}



func (a *RedshiftActivities) DescribeOrderableClusterOptions(input *redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error) {
    return a.client.DescribeOrderableClusterOptions(input)
}



func (a *RedshiftActivities) DescribeReservedNodeOfferings(input *redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error) {
    return a.client.DescribeReservedNodeOfferings(input)
}



func (a *RedshiftActivities) DescribeReservedNodes(input *redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error) {
    return a.client.DescribeReservedNodes(input)
}



func (a *RedshiftActivities) DescribeResize(input *redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error) {
    return a.client.DescribeResize(input)
}



func (a *RedshiftActivities) DescribeScheduledActions(input *redshift.DescribeScheduledActionsInput) (*redshift.DescribeScheduledActionsOutput, error) {
    return a.client.DescribeScheduledActions(input)
}



func (a *RedshiftActivities) DescribeSnapshotCopyGrants(input *redshift.DescribeSnapshotCopyGrantsInput) (*redshift.DescribeSnapshotCopyGrantsOutput, error) {
    return a.client.DescribeSnapshotCopyGrants(input)
}



func (a *RedshiftActivities) DescribeSnapshotSchedules(input *redshift.DescribeSnapshotSchedulesInput) (*redshift.DescribeSnapshotSchedulesOutput, error) {
    return a.client.DescribeSnapshotSchedules(input)
}



func (a *RedshiftActivities) DescribeStorage(input *redshift.DescribeStorageInput) (*redshift.DescribeStorageOutput, error) {
    return a.client.DescribeStorage(input)
}



func (a *RedshiftActivities) DescribeTableRestoreStatus(input *redshift.DescribeTableRestoreStatusInput) (*redshift.DescribeTableRestoreStatusOutput, error) {
    return a.client.DescribeTableRestoreStatus(input)
}



func (a *RedshiftActivities) DescribeTags(input *redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}



func (a *RedshiftActivities) DescribeUsageLimits(input *redshift.DescribeUsageLimitsInput) (*redshift.DescribeUsageLimitsOutput, error) {
    return a.client.DescribeUsageLimits(input)
}



func (a *RedshiftActivities) DisableLogging(input *redshift.DisableLoggingInput) (*redshift.LoggingStatus, error) {
    return a.client.DisableLogging(input)
}



func (a *RedshiftActivities) DisableSnapshotCopy(input *redshift.DisableSnapshotCopyInput) (*redshift.DisableSnapshotCopyOutput, error) {
    return a.client.DisableSnapshotCopy(input)
}



func (a *RedshiftActivities) EnableLogging(input *redshift.EnableLoggingInput) (*redshift.LoggingStatus, error) {
    return a.client.EnableLogging(input)
}



func (a *RedshiftActivities) EnableSnapshotCopy(input *redshift.EnableSnapshotCopyInput) (*redshift.EnableSnapshotCopyOutput, error) {
    return a.client.EnableSnapshotCopy(input)
}



func (a *RedshiftActivities) GetClusterCredentials(input *redshift.GetClusterCredentialsInput) (*redshift.GetClusterCredentialsOutput, error) {
    return a.client.GetClusterCredentials(input)
}



func (a *RedshiftActivities) GetReservedNodeExchangeOfferings(input *redshift.GetReservedNodeExchangeOfferingsInput) (*redshift.GetReservedNodeExchangeOfferingsOutput, error) {
    return a.client.GetReservedNodeExchangeOfferings(input)
}



func (a *RedshiftActivities) ModifyCluster(input *redshift.ModifyClusterInput) (*redshift.ModifyClusterOutput, error) {
    return a.client.ModifyCluster(input)
}



func (a *RedshiftActivities) ModifyClusterDbRevision(input *redshift.ModifyClusterDbRevisionInput) (*redshift.ModifyClusterDbRevisionOutput, error) {
    return a.client.ModifyClusterDbRevision(input)
}



func (a *RedshiftActivities) ModifyClusterIamRoles(input *redshift.ModifyClusterIamRolesInput) (*redshift.ModifyClusterIamRolesOutput, error) {
    return a.client.ModifyClusterIamRoles(input)
}



func (a *RedshiftActivities) ModifyClusterMaintenance(input *redshift.ModifyClusterMaintenanceInput) (*redshift.ModifyClusterMaintenanceOutput, error) {
    return a.client.ModifyClusterMaintenance(input)
}



func (a *RedshiftActivities) ModifyClusterParameterGroup(input *redshift.ModifyClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error) {
    return a.client.ModifyClusterParameterGroup(input)
}



func (a *RedshiftActivities) ModifyClusterSnapshot(input *redshift.ModifyClusterSnapshotInput) (*redshift.ModifyClusterSnapshotOutput, error) {
    return a.client.ModifyClusterSnapshot(input)
}



func (a *RedshiftActivities) ModifyClusterSnapshotSchedule(input *redshift.ModifyClusterSnapshotScheduleInput) (*redshift.ModifyClusterSnapshotScheduleOutput, error) {
    return a.client.ModifyClusterSnapshotSchedule(input)
}



func (a *RedshiftActivities) ModifyClusterSubnetGroup(input *redshift.ModifyClusterSubnetGroupInput) (*redshift.ModifyClusterSubnetGroupOutput, error) {
    return a.client.ModifyClusterSubnetGroup(input)
}



func (a *RedshiftActivities) ModifyEventSubscription(input *redshift.ModifyEventSubscriptionInput) (*redshift.ModifyEventSubscriptionOutput, error) {
    return a.client.ModifyEventSubscription(input)
}



func (a *RedshiftActivities) ModifyScheduledAction(input *redshift.ModifyScheduledActionInput) (*redshift.ModifyScheduledActionOutput, error) {
    return a.client.ModifyScheduledAction(input)
}



func (a *RedshiftActivities) ModifySnapshotCopyRetentionPeriod(input *redshift.ModifySnapshotCopyRetentionPeriodInput) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error) {
    return a.client.ModifySnapshotCopyRetentionPeriod(input)
}



func (a *RedshiftActivities) ModifySnapshotSchedule(input *redshift.ModifySnapshotScheduleInput) (*redshift.ModifySnapshotScheduleOutput, error) {
    return a.client.ModifySnapshotSchedule(input)
}



func (a *RedshiftActivities) ModifyUsageLimit(input *redshift.ModifyUsageLimitInput) (*redshift.ModifyUsageLimitOutput, error) {
    return a.client.ModifyUsageLimit(input)
}



func (a *RedshiftActivities) PauseCluster(input *redshift.PauseClusterInput) (*redshift.PauseClusterOutput, error) {
    return a.client.PauseCluster(input)
}



func (a *RedshiftActivities) PurchaseReservedNodeOffering(input *redshift.PurchaseReservedNodeOfferingInput) (*redshift.PurchaseReservedNodeOfferingOutput, error) {
    return a.client.PurchaseReservedNodeOffering(input)
}



func (a *RedshiftActivities) RebootCluster(input *redshift.RebootClusterInput) (*redshift.RebootClusterOutput, error) {
    return a.client.RebootCluster(input)
}



func (a *RedshiftActivities) ResetClusterParameterGroup(input *redshift.ResetClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error) {
    return a.client.ResetClusterParameterGroup(input)
}



func (a *RedshiftActivities) ResizeCluster(input *redshift.ResizeClusterInput) (*redshift.ResizeClusterOutput, error) {
    return a.client.ResizeCluster(input)
}



func (a *RedshiftActivities) RestoreFromClusterSnapshot(input *redshift.RestoreFromClusterSnapshotInput) (*redshift.RestoreFromClusterSnapshotOutput, error) {
    return a.client.RestoreFromClusterSnapshot(input)
}



func (a *RedshiftActivities) RestoreTableFromClusterSnapshot(input *redshift.RestoreTableFromClusterSnapshotInput) (*redshift.RestoreTableFromClusterSnapshotOutput, error) {
    return a.client.RestoreTableFromClusterSnapshot(input)
}



func (a *RedshiftActivities) ResumeCluster(input *redshift.ResumeClusterInput) (*redshift.ResumeClusterOutput, error) {
    return a.client.ResumeCluster(input)
}



func (a *RedshiftActivities) RevokeClusterSecurityGroupIngress(input *redshift.RevokeClusterSecurityGroupIngressInput) (*redshift.RevokeClusterSecurityGroupIngressOutput, error) {
    return a.client.RevokeClusterSecurityGroupIngress(input)
}



func (a *RedshiftActivities) RevokeSnapshotAccess(input *redshift.RevokeSnapshotAccessInput) (*redshift.RevokeSnapshotAccessOutput, error) {
    return a.client.RevokeSnapshotAccess(input)
}



func (a *RedshiftActivities) RotateEncryptionKey(input *redshift.RotateEncryptionKeyInput) (*redshift.RotateEncryptionKeyOutput, error) {
    return a.client.RotateEncryptionKey(input)
}



func (a *RedshiftActivities) WaitUntilClusterAvailable(input *redshift.DescribeClustersInput) error {
    return a.client.WaitUntilClusterAvailable(input)
}



func (a *RedshiftActivities) WaitUntilClusterDeleted(input *redshift.DescribeClustersInput) error {
    return a.client.WaitUntilClusterDeleted(input)
}



func (a *RedshiftActivities) WaitUntilClusterRestored(input *redshift.DescribeClustersInput) error {
    return a.client.WaitUntilClusterRestored(input)
}



func (a *RedshiftActivities) WaitUntilSnapshotAvailable(input *redshift.DescribeClusterSnapshotsInput) error {
    return a.client.WaitUntilSnapshotAvailable(input)
}

