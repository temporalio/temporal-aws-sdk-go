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
}

func NewKafkaActivities(session *session.Session, config ...*aws.Config) *KafkaActivities {
	client := kafka.New(session, config...)
	return &KafkaActivities{client: client}
}

func (a *KafkaActivities) CreateCluster(ctx context.Context, input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error) {
	return a.client.CreateClusterWithContext(ctx, input)
}

func (a *KafkaActivities) CreateConfiguration(ctx context.Context, input *kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error) {
	return a.client.CreateConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) DeleteCluster(ctx context.Context, input *kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error) {
	return a.client.DeleteClusterWithContext(ctx, input)
}

func (a *KafkaActivities) DeleteConfiguration(ctx context.Context, input *kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error) {
	return a.client.DeleteConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeCluster(ctx context.Context, input *kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error) {
	return a.client.DescribeClusterWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeClusterOperation(ctx context.Context, input *kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error) {
	return a.client.DescribeClusterOperationWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeConfiguration(ctx context.Context, input *kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error) {
	return a.client.DescribeConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) DescribeConfigurationRevision(ctx context.Context, input *kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error) {
	return a.client.DescribeConfigurationRevisionWithContext(ctx, input)
}

func (a *KafkaActivities) GetBootstrapBrokers(ctx context.Context, input *kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error) {
	return a.client.GetBootstrapBrokersWithContext(ctx, input)
}

func (a *KafkaActivities) GetCompatibleKafkaVersions(ctx context.Context, input *kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
	return a.client.GetCompatibleKafkaVersionsWithContext(ctx, input)
}

func (a *KafkaActivities) ListClusterOperations(ctx context.Context, input *kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error) {
	return a.client.ListClusterOperationsWithContext(ctx, input)
}

func (a *KafkaActivities) ListClusters(ctx context.Context, input *kafka.ListClustersInput) (*kafka.ListClustersOutput, error) {
	return a.client.ListClustersWithContext(ctx, input)
}

func (a *KafkaActivities) ListConfigurationRevisions(ctx context.Context, input *kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error) {
	return a.client.ListConfigurationRevisionsWithContext(ctx, input)
}

func (a *KafkaActivities) ListConfigurations(ctx context.Context, input *kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error) {
	return a.client.ListConfigurationsWithContext(ctx, input)
}

func (a *KafkaActivities) ListKafkaVersions(ctx context.Context, input *kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error) {
	return a.client.ListKafkaVersionsWithContext(ctx, input)
}

func (a *KafkaActivities) ListNodes(ctx context.Context, input *kafka.ListNodesInput) (*kafka.ListNodesOutput, error) {
	return a.client.ListNodesWithContext(ctx, input)
}

func (a *KafkaActivities) ListTagsForResource(ctx context.Context, input *kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *KafkaActivities) RebootBroker(ctx context.Context, input *kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error) {
	return a.client.RebootBrokerWithContext(ctx, input)
}

func (a *KafkaActivities) TagResource(ctx context.Context, input *kafka.TagResourceInput) (*kafka.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *KafkaActivities) UntagResource(ctx context.Context, input *kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateBrokerCount(ctx context.Context, input *kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error) {
	return a.client.UpdateBrokerCountWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateBrokerStorage(ctx context.Context, input *kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error) {
	return a.client.UpdateBrokerStorageWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateClusterConfiguration(ctx context.Context, input *kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error) {
	return a.client.UpdateClusterConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateClusterKafkaVersion(ctx context.Context, input *kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error) {
	return a.client.UpdateClusterKafkaVersionWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateConfiguration(ctx context.Context, input *kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error) {
	return a.client.UpdateConfigurationWithContext(ctx, input)
}

func (a *KafkaActivities) UpdateMonitoring(ctx context.Context, input *kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error) {
	return a.client.UpdateMonitoringWithContext(ctx, input)
}
