package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kafka/kafkaiface"
)

type KafkaActivities struct {
	client kafkaiface.KafkaAPI
}

func NewKafkaActivities(client kafkaiface.KafkaAPI) *KafkaActivities {
    return &KafkaActivities{client: client}
}


func (a *KafkaActivities) CreateCluster(input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error) {
    return a.client.CreateCluster(input)
}



func (a *KafkaActivities) CreateConfiguration(input *kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error) {
    return a.client.CreateConfiguration(input)
}



func (a *KafkaActivities) DeleteCluster(input *kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error) {
    return a.client.DeleteCluster(input)
}



func (a *KafkaActivities) DeleteConfiguration(input *kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error) {
    return a.client.DeleteConfiguration(input)
}



func (a *KafkaActivities) DescribeCluster(input *kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error) {
    return a.client.DescribeCluster(input)
}



func (a *KafkaActivities) DescribeClusterOperation(input *kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error) {
    return a.client.DescribeClusterOperation(input)
}



func (a *KafkaActivities) DescribeConfiguration(input *kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error) {
    return a.client.DescribeConfiguration(input)
}



func (a *KafkaActivities) DescribeConfigurationRevision(input *kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error) {
    return a.client.DescribeConfigurationRevision(input)
}



func (a *KafkaActivities) GetBootstrapBrokers(input *kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error) {
    return a.client.GetBootstrapBrokers(input)
}



func (a *KafkaActivities) GetCompatibleKafkaVersions(input *kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
    return a.client.GetCompatibleKafkaVersions(input)
}



func (a *KafkaActivities) ListClusterOperations(input *kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error) {
    return a.client.ListClusterOperations(input)
}



func (a *KafkaActivities) ListClusters(input *kafka.ListClustersInput) (*kafka.ListClustersOutput, error) {
    return a.client.ListClusters(input)
}



func (a *KafkaActivities) ListConfigurationRevisions(input *kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error) {
    return a.client.ListConfigurationRevisions(input)
}



func (a *KafkaActivities) ListConfigurations(input *kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error) {
    return a.client.ListConfigurations(input)
}



func (a *KafkaActivities) ListKafkaVersions(input *kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error) {
    return a.client.ListKafkaVersions(input)
}



func (a *KafkaActivities) ListNodes(input *kafka.ListNodesInput) (*kafka.ListNodesOutput, error) {
    return a.client.ListNodes(input)
}



func (a *KafkaActivities) ListTagsForResource(input *kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *KafkaActivities) RebootBroker(input *kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error) {
    return a.client.RebootBroker(input)
}



func (a *KafkaActivities) TagResource(input *kafka.TagResourceInput) (*kafka.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *KafkaActivities) UntagResource(input *kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *KafkaActivities) UpdateBrokerCount(input *kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error) {
    return a.client.UpdateBrokerCount(input)
}



func (a *KafkaActivities) UpdateBrokerStorage(input *kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error) {
    return a.client.UpdateBrokerStorage(input)
}



func (a *KafkaActivities) UpdateClusterConfiguration(input *kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error) {
    return a.client.UpdateClusterConfiguration(input)
}



func (a *KafkaActivities) UpdateClusterKafkaVersion(input *kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error) {
    return a.client.UpdateClusterKafkaVersion(input)
}



func (a *KafkaActivities) UpdateConfiguration(input *kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error) {
    return a.client.UpdateConfiguration(input)
}



func (a *KafkaActivities) UpdateMonitoring(input *kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error) {
    return a.client.UpdateMonitoring(input)
}

