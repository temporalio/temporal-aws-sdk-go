package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ElastiCacheActivities struct {
	client elasticacheiface.ElastiCacheAPI
}

func NewElastiCacheActivities(session *session.Session, config ...*aws.Config) *ElastiCacheActivities {
	client := elasticache.New(session, config...)
	return &ElastiCacheActivities{client: client}
}

func (a *ElastiCacheActivities) AddTagsToResource(ctx context.Context, input *elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *ElastiCacheActivities) AuthorizeCacheSecurityGroupIngress(ctx context.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error) {
	return a.client.AuthorizeCacheSecurityGroupIngressWithContext(ctx, input)
}

func (a *ElastiCacheActivities) BatchApplyUpdateAction(ctx context.Context, input *elasticache.BatchApplyUpdateActionInput) (*elasticache.BatchApplyUpdateActionOutput, error) {
	return a.client.BatchApplyUpdateActionWithContext(ctx, input)
}

func (a *ElastiCacheActivities) BatchStopUpdateAction(ctx context.Context, input *elasticache.BatchStopUpdateActionInput) (*elasticache.BatchStopUpdateActionOutput, error) {
	return a.client.BatchStopUpdateActionWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CompleteMigration(ctx context.Context, input *elasticache.CompleteMigrationInput) (*elasticache.CompleteMigrationOutput, error) {
	return a.client.CompleteMigrationWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CopySnapshot(ctx context.Context, input *elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error) {
	return a.client.CopySnapshotWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CreateCacheCluster(ctx context.Context, input *elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error) {
	return a.client.CreateCacheClusterWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CreateCacheParameterGroup(ctx context.Context, input *elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error) {
	return a.client.CreateCacheParameterGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CreateCacheSecurityGroup(ctx context.Context, input *elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error) {
	return a.client.CreateCacheSecurityGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CreateCacheSubnetGroup(ctx context.Context, input *elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error) {
	return a.client.CreateCacheSubnetGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CreateGlobalReplicationGroup(ctx context.Context, input *elasticache.CreateGlobalReplicationGroupInput) (*elasticache.CreateGlobalReplicationGroupOutput, error) {
	return a.client.CreateGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CreateReplicationGroup(ctx context.Context, input *elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error) {
	return a.client.CreateReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) CreateSnapshot(ctx context.Context, input *elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error) {
	return a.client.CreateSnapshotWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DecreaseNodeGroupsInGlobalReplicationGroup(ctx context.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
	return a.client.DecreaseNodeGroupsInGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DecreaseReplicaCount(ctx context.Context, input *elasticache.DecreaseReplicaCountInput) (*elasticache.DecreaseReplicaCountOutput, error) {
	return a.client.DecreaseReplicaCountWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DeleteCacheCluster(ctx context.Context, input *elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error) {
	return a.client.DeleteCacheClusterWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DeleteCacheParameterGroup(ctx context.Context, input *elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error) {
	return a.client.DeleteCacheParameterGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DeleteCacheSecurityGroup(ctx context.Context, input *elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error) {
	return a.client.DeleteCacheSecurityGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DeleteCacheSubnetGroup(ctx context.Context, input *elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error) {
	return a.client.DeleteCacheSubnetGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DeleteGlobalReplicationGroup(ctx context.Context, input *elasticache.DeleteGlobalReplicationGroupInput) (*elasticache.DeleteGlobalReplicationGroupOutput, error) {
	return a.client.DeleteGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DeleteReplicationGroup(ctx context.Context, input *elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error) {
	return a.client.DeleteReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DeleteSnapshot(ctx context.Context, input *elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error) {
	return a.client.DeleteSnapshotWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeCacheClusters(ctx context.Context, input *elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error) {
	return a.client.DescribeCacheClustersWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeCacheEngineVersions(ctx context.Context, input *elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
	return a.client.DescribeCacheEngineVersionsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeCacheParameterGroups(ctx context.Context, input *elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
	return a.client.DescribeCacheParameterGroupsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeCacheParameters(ctx context.Context, input *elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error) {
	return a.client.DescribeCacheParametersWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeCacheSecurityGroups(ctx context.Context, input *elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
	return a.client.DescribeCacheSecurityGroupsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeCacheSubnetGroups(ctx context.Context, input *elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
	return a.client.DescribeCacheSubnetGroupsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeEngineDefaultParameters(ctx context.Context, input *elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
	return a.client.DescribeEngineDefaultParametersWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeEvents(ctx context.Context, input *elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error) {
	return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeGlobalReplicationGroups(ctx context.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
	return a.client.DescribeGlobalReplicationGroupsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeReplicationGroups(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error) {
	return a.client.DescribeReplicationGroupsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeReservedCacheNodes(ctx context.Context, input *elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error) {
	return a.client.DescribeReservedCacheNodesWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeReservedCacheNodesOfferings(ctx context.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
	return a.client.DescribeReservedCacheNodesOfferingsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeServiceUpdates(ctx context.Context, input *elasticache.DescribeServiceUpdatesInput) (*elasticache.DescribeServiceUpdatesOutput, error) {
	return a.client.DescribeServiceUpdatesWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeSnapshots(ctx context.Context, input *elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error) {
	return a.client.DescribeSnapshotsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DescribeUpdateActions(ctx context.Context, input *elasticache.DescribeUpdateActionsInput) (*elasticache.DescribeUpdateActionsOutput, error) {
	return a.client.DescribeUpdateActionsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) DisassociateGlobalReplicationGroup(ctx context.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) (*elasticache.DisassociateGlobalReplicationGroupOutput, error) {
	return a.client.DisassociateGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) FailoverGlobalReplicationGroup(ctx context.Context, input *elasticache.FailoverGlobalReplicationGroupInput) (*elasticache.FailoverGlobalReplicationGroupOutput, error) {
	return a.client.FailoverGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) IncreaseNodeGroupsInGlobalReplicationGroup(ctx context.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
	return a.client.IncreaseNodeGroupsInGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) IncreaseReplicaCount(ctx context.Context, input *elasticache.IncreaseReplicaCountInput) (*elasticache.IncreaseReplicaCountOutput, error) {
	return a.client.IncreaseReplicaCountWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ListAllowedNodeTypeModifications(ctx context.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
	return a.client.ListAllowedNodeTypeModificationsWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ListTagsForResource(ctx context.Context, input *elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ModifyCacheCluster(ctx context.Context, input *elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error) {
	return a.client.ModifyCacheClusterWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ModifyCacheParameterGroup(ctx context.Context, input *elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
	return a.client.ModifyCacheParameterGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ModifyCacheSubnetGroup(ctx context.Context, input *elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error) {
	return a.client.ModifyCacheSubnetGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ModifyGlobalReplicationGroup(ctx context.Context, input *elasticache.ModifyGlobalReplicationGroupInput) (*elasticache.ModifyGlobalReplicationGroupOutput, error) {
	return a.client.ModifyGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ModifyReplicationGroup(ctx context.Context, input *elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error) {
	return a.client.ModifyReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ModifyReplicationGroupShardConfiguration(ctx context.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error) {
	return a.client.ModifyReplicationGroupShardConfigurationWithContext(ctx, input)
}

func (a *ElastiCacheActivities) PurchaseReservedCacheNodesOffering(ctx context.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error) {
	return a.client.PurchaseReservedCacheNodesOfferingWithContext(ctx, input)
}

func (a *ElastiCacheActivities) RebalanceSlotsInGlobalReplicationGroup(ctx context.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) (*elasticache.RebalanceSlotsInGlobalReplicationGroupOutput, error) {
	return a.client.RebalanceSlotsInGlobalReplicationGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) RebootCacheCluster(ctx context.Context, input *elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error) {
	return a.client.RebootCacheClusterWithContext(ctx, input)
}

func (a *ElastiCacheActivities) RemoveTagsFromResource(ctx context.Context, input *elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *ElastiCacheActivities) ResetCacheParameterGroup(ctx context.Context, input *elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
	return a.client.ResetCacheParameterGroupWithContext(ctx, input)
}

func (a *ElastiCacheActivities) RevokeCacheSecurityGroupIngress(ctx context.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error) {
	return a.client.RevokeCacheSecurityGroupIngressWithContext(ctx, input)
}

func (a *ElastiCacheActivities) StartMigration(ctx context.Context, input *elasticache.StartMigrationInput) (*elasticache.StartMigrationOutput, error) {
	return a.client.StartMigrationWithContext(ctx, input)
}

func (a *ElastiCacheActivities) TestFailover(ctx context.Context, input *elasticache.TestFailoverInput) (*elasticache.TestFailoverOutput, error) {
	return a.client.TestFailoverWithContext(ctx, input)
}

func (a *ElastiCacheActivities) WaitUntilCacheClusterAvailable(ctx context.Context, input *elasticache.DescribeCacheClustersInput) error {
	return a.client.WaitUntilCacheClusterAvailableWithContext(ctx, input)

}

func (a *ElastiCacheActivities) WaitUntilCacheClusterDeleted(ctx context.Context, input *elasticache.DescribeCacheClustersInput) error {
	return a.client.WaitUntilCacheClusterDeletedWithContext(ctx, input)

}

func (a *ElastiCacheActivities) WaitUntilReplicationGroupAvailable(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput) error {
	return a.client.WaitUntilReplicationGroupAvailableWithContext(ctx, input)

}

func (a *ElastiCacheActivities) WaitUntilReplicationGroupDeleted(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput) error {
	return a.client.WaitUntilReplicationGroupDeletedWithContext(ctx, input)

}
