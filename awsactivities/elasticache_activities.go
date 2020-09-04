
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
)

type ElastiCacheActivities struct {
	client elasticacheiface.ElastiCacheAPI
}

func NewElastiCacheActivities(client elasticacheiface.ElastiCacheAPI) *ElastiCacheActivities {
    return &ElastiCacheActivities{client: client}
}

func (a *ElastiCacheActivities) AddTagsToResource(input *elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error) {
    return a.client.AddTagsToResource(input)
}

func (a *ElastiCacheActivities) AuthorizeCacheSecurityGroupIngress(input *elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error) {
    return a.client.AuthorizeCacheSecurityGroupIngress(input)
}

func (a *ElastiCacheActivities) BatchApplyUpdateAction(input *elasticache.BatchApplyUpdateActionInput) (*elasticache.BatchApplyUpdateActionOutput, error) {
    return a.client.BatchApplyUpdateAction(input)
}

func (a *ElastiCacheActivities) BatchStopUpdateAction(input *elasticache.BatchStopUpdateActionInput) (*elasticache.BatchStopUpdateActionOutput, error) {
    return a.client.BatchStopUpdateAction(input)
}

func (a *ElastiCacheActivities) CompleteMigration(input *elasticache.CompleteMigrationInput) (*elasticache.CompleteMigrationOutput, error) {
    return a.client.CompleteMigration(input)
}

func (a *ElastiCacheActivities) CopySnapshot(input *elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error) {
    return a.client.CopySnapshot(input)
}

func (a *ElastiCacheActivities) CreateCacheCluster(input *elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error) {
    return a.client.CreateCacheCluster(input)
}

func (a *ElastiCacheActivities) CreateCacheParameterGroup(input *elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error) {
    return a.client.CreateCacheParameterGroup(input)
}

func (a *ElastiCacheActivities) CreateCacheSecurityGroup(input *elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error) {
    return a.client.CreateCacheSecurityGroup(input)
}

func (a *ElastiCacheActivities) CreateCacheSubnetGroup(input *elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error) {
    return a.client.CreateCacheSubnetGroup(input)
}

func (a *ElastiCacheActivities) CreateGlobalReplicationGroup(input *elasticache.CreateGlobalReplicationGroupInput) (*elasticache.CreateGlobalReplicationGroupOutput, error) {
    return a.client.CreateGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) CreateReplicationGroup(input *elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error) {
    return a.client.CreateReplicationGroup(input)
}

func (a *ElastiCacheActivities) CreateSnapshot(input *elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error) {
    return a.client.CreateSnapshot(input)
}

func (a *ElastiCacheActivities) DecreaseNodeGroupsInGlobalReplicationGroup(input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
    return a.client.DecreaseNodeGroupsInGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) DecreaseReplicaCount(input *elasticache.DecreaseReplicaCountInput) (*elasticache.DecreaseReplicaCountOutput, error) {
    return a.client.DecreaseReplicaCount(input)
}

func (a *ElastiCacheActivities) DeleteCacheCluster(input *elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error) {
    return a.client.DeleteCacheCluster(input)
}

func (a *ElastiCacheActivities) DeleteCacheParameterGroup(input *elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error) {
    return a.client.DeleteCacheParameterGroup(input)
}

func (a *ElastiCacheActivities) DeleteCacheSecurityGroup(input *elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error) {
    return a.client.DeleteCacheSecurityGroup(input)
}

func (a *ElastiCacheActivities) DeleteCacheSubnetGroup(input *elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error) {
    return a.client.DeleteCacheSubnetGroup(input)
}

func (a *ElastiCacheActivities) DeleteGlobalReplicationGroup(input *elasticache.DeleteGlobalReplicationGroupInput) (*elasticache.DeleteGlobalReplicationGroupOutput, error) {
    return a.client.DeleteGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) DeleteReplicationGroup(input *elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error) {
    return a.client.DeleteReplicationGroup(input)
}

func (a *ElastiCacheActivities) DeleteSnapshot(input *elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error) {
    return a.client.DeleteSnapshot(input)
}

func (a *ElastiCacheActivities) DescribeCacheClusters(input *elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error) {
    return a.client.DescribeCacheClusters(input)
}

func (a *ElastiCacheActivities) DescribeCacheEngineVersions(input *elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
    return a.client.DescribeCacheEngineVersions(input)
}

func (a *ElastiCacheActivities) DescribeCacheParameterGroups(input *elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
    return a.client.DescribeCacheParameterGroups(input)
}

func (a *ElastiCacheActivities) DescribeCacheParameters(input *elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error) {
    return a.client.DescribeCacheParameters(input)
}

func (a *ElastiCacheActivities) DescribeCacheSecurityGroups(input *elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
    return a.client.DescribeCacheSecurityGroups(input)
}

func (a *ElastiCacheActivities) DescribeCacheSubnetGroups(input *elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
    return a.client.DescribeCacheSubnetGroups(input)
}

func (a *ElastiCacheActivities) DescribeEngineDefaultParameters(input *elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
    return a.client.DescribeEngineDefaultParameters(input)
}

func (a *ElastiCacheActivities) DescribeEvents(input *elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}

func (a *ElastiCacheActivities) DescribeGlobalReplicationGroups(input *elasticache.DescribeGlobalReplicationGroupsInput) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
    return a.client.DescribeGlobalReplicationGroups(input)
}

func (a *ElastiCacheActivities) DescribeReplicationGroups(input *elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error) {
    return a.client.DescribeReplicationGroups(input)
}

func (a *ElastiCacheActivities) DescribeReservedCacheNodes(input *elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error) {
    return a.client.DescribeReservedCacheNodes(input)
}

func (a *ElastiCacheActivities) DescribeReservedCacheNodesOfferings(input *elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
    return a.client.DescribeReservedCacheNodesOfferings(input)
}

func (a *ElastiCacheActivities) DescribeServiceUpdates(input *elasticache.DescribeServiceUpdatesInput) (*elasticache.DescribeServiceUpdatesOutput, error) {
    return a.client.DescribeServiceUpdates(input)
}

func (a *ElastiCacheActivities) DescribeSnapshots(input *elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error) {
    return a.client.DescribeSnapshots(input)
}

func (a *ElastiCacheActivities) DescribeUpdateActions(input *elasticache.DescribeUpdateActionsInput) (*elasticache.DescribeUpdateActionsOutput, error) {
    return a.client.DescribeUpdateActions(input)
}

func (a *ElastiCacheActivities) DisassociateGlobalReplicationGroup(input *elasticache.DisassociateGlobalReplicationGroupInput) (*elasticache.DisassociateGlobalReplicationGroupOutput, error) {
    return a.client.DisassociateGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) FailoverGlobalReplicationGroup(input *elasticache.FailoverGlobalReplicationGroupInput) (*elasticache.FailoverGlobalReplicationGroupOutput, error) {
    return a.client.FailoverGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) IncreaseNodeGroupsInGlobalReplicationGroup(input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
    return a.client.IncreaseNodeGroupsInGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) IncreaseReplicaCount(input *elasticache.IncreaseReplicaCountInput) (*elasticache.IncreaseReplicaCountOutput, error) {
    return a.client.IncreaseReplicaCount(input)
}

func (a *ElastiCacheActivities) ListAllowedNodeTypeModifications(input *elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
    return a.client.ListAllowedNodeTypeModifications(input)
}

func (a *ElastiCacheActivities) ListTagsForResource(input *elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ElastiCacheActivities) ModifyCacheCluster(input *elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error) {
    return a.client.ModifyCacheCluster(input)
}

func (a *ElastiCacheActivities) ModifyCacheParameterGroup(input *elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
    return a.client.ModifyCacheParameterGroup(input)
}

func (a *ElastiCacheActivities) ModifyCacheSubnetGroup(input *elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error) {
    return a.client.ModifyCacheSubnetGroup(input)
}

func (a *ElastiCacheActivities) ModifyGlobalReplicationGroup(input *elasticache.ModifyGlobalReplicationGroupInput) (*elasticache.ModifyGlobalReplicationGroupOutput, error) {
    return a.client.ModifyGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) ModifyReplicationGroup(input *elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error) {
    return a.client.ModifyReplicationGroup(input)
}

func (a *ElastiCacheActivities) ModifyReplicationGroupShardConfiguration(input *elasticache.ModifyReplicationGroupShardConfigurationInput) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error) {
    return a.client.ModifyReplicationGroupShardConfiguration(input)
}

func (a *ElastiCacheActivities) PurchaseReservedCacheNodesOffering(input *elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error) {
    return a.client.PurchaseReservedCacheNodesOffering(input)
}

func (a *ElastiCacheActivities) RebalanceSlotsInGlobalReplicationGroup(input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) (*elasticache.RebalanceSlotsInGlobalReplicationGroupOutput, error) {
    return a.client.RebalanceSlotsInGlobalReplicationGroup(input)
}

func (a *ElastiCacheActivities) RebootCacheCluster(input *elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error) {
    return a.client.RebootCacheCluster(input)
}

func (a *ElastiCacheActivities) RemoveTagsFromResource(input *elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error) {
    return a.client.RemoveTagsFromResource(input)
}

func (a *ElastiCacheActivities) ResetCacheParameterGroup(input *elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
    return a.client.ResetCacheParameterGroup(input)
}

func (a *ElastiCacheActivities) RevokeCacheSecurityGroupIngress(input *elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error) {
    return a.client.RevokeCacheSecurityGroupIngress(input)
}

func (a *ElastiCacheActivities) StartMigration(input *elasticache.StartMigrationInput) (*elasticache.StartMigrationOutput, error) {
    return a.client.StartMigration(input)
}

func (a *ElastiCacheActivities) TestFailover(input *elasticache.TestFailoverInput) (*elasticache.TestFailoverOutput, error) {
    return a.client.TestFailover(input)
}

func (a *ElastiCacheActivities) WaitUntilCacheClusterAvailable(input *elasticache.DescribeCacheClustersInput) error {
    return a.client.WaitUntilCacheClusterAvailable(input)
}

func (a *ElastiCacheActivities) WaitUntilCacheClusterDeleted(input *elasticache.DescribeCacheClustersInput) error {
    return a.client.WaitUntilCacheClusterDeleted(input)
}

func (a *ElastiCacheActivities) WaitUntilReplicationGroupAvailable(input *elasticache.DescribeReplicationGroupsInput) error {
    return a.client.WaitUntilReplicationGroupAvailable(input)
}

func (a *ElastiCacheActivities) WaitUntilReplicationGroupDeleted(input *elasticache.DescribeReplicationGroupsInput) error {
    return a.client.WaitUntilReplicationGroupDeleted(input)
}
