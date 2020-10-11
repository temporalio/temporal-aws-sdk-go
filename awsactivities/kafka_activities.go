// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kafka/kafkaiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type KafkaActivities struct {
	client kafkaiface.KafkaAPI

	sessionFactory SessionFactory
}

func NewKafkaActivities(sess *session.Session, config ...*aws.Config) *KafkaActivities {
	client := kafka.New(sess, config...)
	return &KafkaActivities{client: client}
}

func NewKafkaActivitiesWithSessionFactory(sessionFactory SessionFactory) *KafkaActivities {
	return &KafkaActivities{sessionFactory: sessionFactory}
}

func (a *KafkaActivities) getClient(ctx context.Context) (kafkaiface.KafkaAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return kafka.New(sess), nil
}

func (a *KafkaActivities) CreateCluster(ctx context.Context, input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateClusterWithContext(ctx, input)
}

func (a *KafkaActivities) CreateConfiguration(ctx context.Context, input *kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) DeleteCluster(ctx context.Context, input *kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteClusterWithContext(ctx, input)
}

func (a *KafkaActivities) DeleteConfiguration(ctx context.Context, input *kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeCluster(ctx context.Context, input *kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeClusterWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeClusterOperation(ctx context.Context, input *kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeClusterOperationWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeConfiguration(ctx context.Context, input *kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeConfigurationRevision(ctx context.Context, input *kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConfigurationRevisionWithContext(ctx, input)
}

func (a *KafkaActivities) GetBootstrapBrokers(ctx context.Context, input *kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBootstrapBrokersWithContext(ctx, input)
}

func (a *KafkaActivities) GetCompatibleKafkaVersions(ctx context.Context, input *kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCompatibleKafkaVersionsWithContext(ctx, input)
}

func (a *KafkaActivities) ListClusterOperations(ctx context.Context, input *kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListClusterOperationsWithContext(ctx, input)
}

func (a *KafkaActivities) ListClusters(ctx context.Context, input *kafka.ListClustersInput) (*kafka.ListClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListClustersWithContext(ctx, input)
}

func (a *KafkaActivities) ListConfigurationRevisions(ctx context.Context, input *kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationRevisionsWithContext(ctx, input)
}

func (a *KafkaActivities) ListConfigurations(ctx context.Context, input *kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationsWithContext(ctx, input)
}

func (a *KafkaActivities) ListKafkaVersions(ctx context.Context, input *kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListKafkaVersionsWithContext(ctx, input)
}

func (a *KafkaActivities) ListNodes(ctx context.Context, input *kafka.ListNodesInput) (*kafka.ListNodesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNodesWithContext(ctx, input)
}

func (a *KafkaActivities) ListTagsForResource(ctx context.Context, input *kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *KafkaActivities) RebootBroker(ctx context.Context, input *kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootBrokerWithContext(ctx, input)
}

func (a *KafkaActivities) TagResource(ctx context.Context, input *kafka.TagResourceInput) (*kafka.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *KafkaActivities) UntagResource(ctx context.Context, input *kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateBrokerCount(ctx context.Context, input *kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateBrokerCountWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateBrokerStorage(ctx context.Context, input *kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateBrokerStorageWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateClusterConfiguration(ctx context.Context, input *kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateClusterConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateClusterKafkaVersion(ctx context.Context, input *kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateClusterKafkaVersionWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateConfiguration(ctx context.Context, input *kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateMonitoring(ctx context.Context, input *kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMonitoringWithContext(ctx, input)
}
