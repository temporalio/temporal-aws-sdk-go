// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elasticache

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client elasticacheiface.ElastiCacheAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := elasticache.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (elasticacheiface.ElastiCacheAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return elasticache.New(sess), nil
}

func (a *Activities) AddTagsToResource(ctx context.Context, input *elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToResourceWithContext(ctx, input)
}

func (a *Activities) AuthorizeCacheSecurityGroupIngress(ctx context.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AuthorizeCacheSecurityGroupIngressWithContext(ctx, input)
}

func (a *Activities) BatchApplyUpdateAction(ctx context.Context, input *elasticache.BatchApplyUpdateActionInput) (*elasticache.BatchApplyUpdateActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchApplyUpdateActionWithContext(ctx, input)
}

func (a *Activities) BatchStopUpdateAction(ctx context.Context, input *elasticache.BatchStopUpdateActionInput) (*elasticache.BatchStopUpdateActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchStopUpdateActionWithContext(ctx, input)
}

func (a *Activities) CompleteMigration(ctx context.Context, input *elasticache.CompleteMigrationInput) (*elasticache.CompleteMigrationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CompleteMigrationWithContext(ctx, input)
}

func (a *Activities) CopySnapshot(ctx context.Context, input *elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopySnapshotWithContext(ctx, input)
}

func (a *Activities) CreateCacheCluster(ctx context.Context, input *elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCacheClusterWithContext(ctx, input)
}

func (a *Activities) CreateCacheParameterGroup(ctx context.Context, input *elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCacheParameterGroupWithContext(ctx, input)
}

func (a *Activities) CreateCacheSecurityGroup(ctx context.Context, input *elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCacheSecurityGroupWithContext(ctx, input)
}

func (a *Activities) CreateCacheSubnetGroup(ctx context.Context, input *elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCacheSubnetGroupWithContext(ctx, input)
}

func (a *Activities) CreateGlobalReplicationGroup(ctx context.Context, input *elasticache.CreateGlobalReplicationGroupInput) (*elasticache.CreateGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) CreateReplicationGroup(ctx context.Context, input *elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateReplicationGroupWithContext(ctx, input)
}

func (a *Activities) CreateSnapshot(ctx context.Context, input *elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSnapshotWithContext(ctx, input)
}

func (a *Activities) CreateUser(ctx context.Context, input *elasticache.CreateUserInput) (*elasticache.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserWithContext(ctx, input)
}

func (a *Activities) CreateUserGroup(ctx context.Context, input *elasticache.CreateUserGroupInput) (*elasticache.CreateUserGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserGroupWithContext(ctx, input)
}

func (a *Activities) DecreaseNodeGroupsInGlobalReplicationGroup(ctx context.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DecreaseNodeGroupsInGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) DecreaseReplicaCount(ctx context.Context, input *elasticache.DecreaseReplicaCountInput) (*elasticache.DecreaseReplicaCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DecreaseReplicaCountWithContext(ctx, input)
}

func (a *Activities) DeleteCacheCluster(ctx context.Context, input *elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCacheClusterWithContext(ctx, input)
}

func (a *Activities) DeleteCacheParameterGroup(ctx context.Context, input *elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCacheParameterGroupWithContext(ctx, input)
}

func (a *Activities) DeleteCacheSecurityGroup(ctx context.Context, input *elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCacheSecurityGroupWithContext(ctx, input)
}

func (a *Activities) DeleteCacheSubnetGroup(ctx context.Context, input *elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCacheSubnetGroupWithContext(ctx, input)
}

func (a *Activities) DeleteGlobalReplicationGroup(ctx context.Context, input *elasticache.DeleteGlobalReplicationGroupInput) (*elasticache.DeleteGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) DeleteReplicationGroup(ctx context.Context, input *elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReplicationGroupWithContext(ctx, input)
}

func (a *Activities) DeleteSnapshot(ctx context.Context, input *elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSnapshotWithContext(ctx, input)
}

func (a *Activities) DeleteUser(ctx context.Context, input *elasticache.DeleteUserInput) (*elasticache.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserWithContext(ctx, input)
}

func (a *Activities) DeleteUserGroup(ctx context.Context, input *elasticache.DeleteUserGroupInput) (*elasticache.DeleteUserGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserGroupWithContext(ctx, input)
}

func (a *Activities) DescribeCacheClusters(ctx context.Context, input *elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCacheClustersWithContext(ctx, input)
}

func (a *Activities) DescribeCacheEngineVersions(ctx context.Context, input *elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCacheEngineVersionsWithContext(ctx, input)
}

func (a *Activities) DescribeCacheParameterGroups(ctx context.Context, input *elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCacheParameterGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeCacheParameters(ctx context.Context, input *elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCacheParametersWithContext(ctx, input)
}

func (a *Activities) DescribeCacheSecurityGroups(ctx context.Context, input *elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCacheSecurityGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeCacheSubnetGroups(ctx context.Context, input *elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCacheSubnetGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeEngineDefaultParameters(ctx context.Context, input *elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEngineDefaultParametersWithContext(ctx, input)
}

func (a *Activities) DescribeEvents(ctx context.Context, input *elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventsWithContext(ctx, input)
}

func (a *Activities) DescribeGlobalReplicationGroups(ctx context.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeGlobalReplicationGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeReplicationGroups(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReplicationGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeReservedCacheNodes(ctx context.Context, input *elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReservedCacheNodesWithContext(ctx, input)
}

func (a *Activities) DescribeReservedCacheNodesOfferings(ctx context.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeReservedCacheNodesOfferingsWithContext(ctx, input)
}

func (a *Activities) DescribeServiceUpdates(ctx context.Context, input *elasticache.DescribeServiceUpdatesInput) (*elasticache.DescribeServiceUpdatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeServiceUpdatesWithContext(ctx, input)
}

func (a *Activities) DescribeSnapshots(ctx context.Context, input *elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSnapshotsWithContext(ctx, input)
}

func (a *Activities) DescribeUpdateActions(ctx context.Context, input *elasticache.DescribeUpdateActionsInput) (*elasticache.DescribeUpdateActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUpdateActionsWithContext(ctx, input)
}

func (a *Activities) DescribeUserGroups(ctx context.Context, input *elasticache.DescribeUserGroupsInput) (*elasticache.DescribeUserGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeUsers(ctx context.Context, input *elasticache.DescribeUsersInput) (*elasticache.DescribeUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUsersWithContext(ctx, input)
}

func (a *Activities) DisassociateGlobalReplicationGroup(ctx context.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) (*elasticache.DisassociateGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) FailoverGlobalReplicationGroup(ctx context.Context, input *elasticache.FailoverGlobalReplicationGroupInput) (*elasticache.FailoverGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.FailoverGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) IncreaseNodeGroupsInGlobalReplicationGroup(ctx context.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.IncreaseNodeGroupsInGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) IncreaseReplicaCount(ctx context.Context, input *elasticache.IncreaseReplicaCountInput) (*elasticache.IncreaseReplicaCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.IncreaseReplicaCountWithContext(ctx, input)
}

func (a *Activities) ListAllowedNodeTypeModifications(ctx context.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAllowedNodeTypeModificationsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ModifyCacheCluster(ctx context.Context, input *elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyCacheClusterWithContext(ctx, input)
}

func (a *Activities) ModifyCacheParameterGroup(ctx context.Context, input *elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyCacheParameterGroupWithContext(ctx, input)
}

func (a *Activities) ModifyCacheSubnetGroup(ctx context.Context, input *elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyCacheSubnetGroupWithContext(ctx, input)
}

func (a *Activities) ModifyGlobalReplicationGroup(ctx context.Context, input *elasticache.ModifyGlobalReplicationGroupInput) (*elasticache.ModifyGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) ModifyReplicationGroup(ctx context.Context, input *elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyReplicationGroupWithContext(ctx, input)
}

func (a *Activities) ModifyReplicationGroupShardConfiguration(ctx context.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyReplicationGroupShardConfigurationWithContext(ctx, input)
}

func (a *Activities) ModifyUser(ctx context.Context, input *elasticache.ModifyUserInput) (*elasticache.ModifyUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyUserWithContext(ctx, input)
}

func (a *Activities) ModifyUserGroup(ctx context.Context, input *elasticache.ModifyUserGroupInput) (*elasticache.ModifyUserGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyUserGroupWithContext(ctx, input)
}

func (a *Activities) PurchaseReservedCacheNodesOffering(ctx context.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PurchaseReservedCacheNodesOfferingWithContext(ctx, input)
}

func (a *Activities) RebalanceSlotsInGlobalReplicationGroup(ctx context.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) (*elasticache.RebalanceSlotsInGlobalReplicationGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebalanceSlotsInGlobalReplicationGroupWithContext(ctx, input)
}

func (a *Activities) RebootCacheCluster(ctx context.Context, input *elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootCacheClusterWithContext(ctx, input)
}

func (a *Activities) RemoveTagsFromResource(ctx context.Context, input *elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *Activities) ResetCacheParameterGroup(ctx context.Context, input *elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetCacheParameterGroupWithContext(ctx, input)
}

func (a *Activities) RevokeCacheSecurityGroupIngress(ctx context.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RevokeCacheSecurityGroupIngressWithContext(ctx, input)
}

func (a *Activities) StartMigration(ctx context.Context, input *elasticache.StartMigrationInput) (*elasticache.StartMigrationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartMigrationWithContext(ctx, input)
}

func (a *Activities) TestFailover(ctx context.Context, input *elasticache.TestFailoverInput) (*elasticache.TestFailoverOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestFailoverWithContext(ctx, input)
}

func (a *Activities) WaitUntilCacheClusterAvailable(ctx context.Context, input *elasticache.DescribeCacheClustersInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilCacheClusterAvailableWithContext(ctx, input, options...)
	})
}

func (a *Activities) WaitUntilCacheClusterDeleted(ctx context.Context, input *elasticache.DescribeCacheClustersInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilCacheClusterDeletedWithContext(ctx, input, options...)
	})
}

func (a *Activities) WaitUntilReplicationGroupAvailable(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationGroupAvailableWithContext(ctx, input, options...)
	})
}

func (a *Activities) WaitUntilReplicationGroupDeleted(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilReplicationGroupDeletedWithContext(ctx, input, options...)
	})
}
