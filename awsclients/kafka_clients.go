package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kafka"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KafkaClient interface {
    CreateCluster(ctx workflow.Context, input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error)
    CreateClusterAsync(ctx workflow.Context, input *kafka.CreateClusterInput) *KafkaCreateClusterResult

    CreateConfiguration(ctx workflow.Context, input *kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error)
    CreateConfigurationAsync(ctx workflow.Context, input *kafka.CreateConfigurationInput) *KafkaCreateConfigurationResult

    DeleteCluster(ctx workflow.Context, input *kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error)
    DeleteClusterAsync(ctx workflow.Context, input *kafka.DeleteClusterInput) *KafkaDeleteClusterResult

    DeleteConfiguration(ctx workflow.Context, input *kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error)
    DeleteConfigurationAsync(ctx workflow.Context, input *kafka.DeleteConfigurationInput) *KafkaDeleteConfigurationResult

    DescribeCluster(ctx workflow.Context, input *kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error)
    DescribeClusterAsync(ctx workflow.Context, input *kafka.DescribeClusterInput) *KafkaDescribeClusterResult

    DescribeClusterOperation(ctx workflow.Context, input *kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error)
    DescribeClusterOperationAsync(ctx workflow.Context, input *kafka.DescribeClusterOperationInput) *KafkaDescribeClusterOperationResult

    DescribeConfiguration(ctx workflow.Context, input *kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error)
    DescribeConfigurationAsync(ctx workflow.Context, input *kafka.DescribeConfigurationInput) *KafkaDescribeConfigurationResult

    DescribeConfigurationRevision(ctx workflow.Context, input *kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error)
    DescribeConfigurationRevisionAsync(ctx workflow.Context, input *kafka.DescribeConfigurationRevisionInput) *KafkaDescribeConfigurationRevisionResult

    GetBootstrapBrokers(ctx workflow.Context, input *kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error)
    GetBootstrapBrokersAsync(ctx workflow.Context, input *kafka.GetBootstrapBrokersInput) *KafkaGetBootstrapBrokersResult

    GetCompatibleKafkaVersions(ctx workflow.Context, input *kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error)
    GetCompatibleKafkaVersionsAsync(ctx workflow.Context, input *kafka.GetCompatibleKafkaVersionsInput) *KafkaGetCompatibleKafkaVersionsResult

    ListClusterOperations(ctx workflow.Context, input *kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error)
    ListClusterOperationsAsync(ctx workflow.Context, input *kafka.ListClusterOperationsInput) *KafkaListClusterOperationsResult

    ListClusters(ctx workflow.Context, input *kafka.ListClustersInput) (*kafka.ListClustersOutput, error)
    ListClustersAsync(ctx workflow.Context, input *kafka.ListClustersInput) *KafkaListClustersResult

    ListConfigurationRevisions(ctx workflow.Context, input *kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error)
    ListConfigurationRevisionsAsync(ctx workflow.Context, input *kafka.ListConfigurationRevisionsInput) *KafkaListConfigurationRevisionsResult

    ListConfigurations(ctx workflow.Context, input *kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error)
    ListConfigurationsAsync(ctx workflow.Context, input *kafka.ListConfigurationsInput) *KafkaListConfigurationsResult

    ListKafkaVersions(ctx workflow.Context, input *kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error)
    ListKafkaVersionsAsync(ctx workflow.Context, input *kafka.ListKafkaVersionsInput) *KafkaListKafkaVersionsResult

    ListNodes(ctx workflow.Context, input *kafka.ListNodesInput) (*kafka.ListNodesOutput, error)
    ListNodesAsync(ctx workflow.Context, input *kafka.ListNodesInput) *KafkaListNodesResult

    ListTagsForResource(ctx workflow.Context, input *kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *kafka.ListTagsForResourceInput) *KafkaListTagsForResourceResult

    RebootBroker(ctx workflow.Context, input *kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error)
    RebootBrokerAsync(ctx workflow.Context, input *kafka.RebootBrokerInput) *KafkaRebootBrokerResult

    TagResource(ctx workflow.Context, input *kafka.TagResourceInput) (*kafka.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *kafka.TagResourceInput) *KafkaTagResourceResult

    UntagResource(ctx workflow.Context, input *kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *kafka.UntagResourceInput) *KafkaUntagResourceResult

    UpdateBrokerCount(ctx workflow.Context, input *kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error)
    UpdateBrokerCountAsync(ctx workflow.Context, input *kafka.UpdateBrokerCountInput) *KafkaUpdateBrokerCountResult

    UpdateBrokerStorage(ctx workflow.Context, input *kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error)
    UpdateBrokerStorageAsync(ctx workflow.Context, input *kafka.UpdateBrokerStorageInput) *KafkaUpdateBrokerStorageResult

    UpdateClusterConfiguration(ctx workflow.Context, input *kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error)
    UpdateClusterConfigurationAsync(ctx workflow.Context, input *kafka.UpdateClusterConfigurationInput) *KafkaUpdateClusterConfigurationResult

    UpdateClusterKafkaVersion(ctx workflow.Context, input *kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error)
    UpdateClusterKafkaVersionAsync(ctx workflow.Context, input *kafka.UpdateClusterKafkaVersionInput) *KafkaUpdateClusterKafkaVersionResult

    UpdateConfiguration(ctx workflow.Context, input *kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error)
    UpdateConfigurationAsync(ctx workflow.Context, input *kafka.UpdateConfigurationInput) *KafkaUpdateConfigurationResult

    UpdateMonitoring(ctx workflow.Context, input *kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error)
    UpdateMonitoringAsync(ctx workflow.Context, input *kafka.UpdateMonitoringInput) *KafkaUpdateMonitoringResult
}
type KafkaCreateClusterResult struct {
	Result workflow.Future
}

func (r *KafkaCreateClusterResult) Get(ctx workflow.Context) (*kafka.CreateClusterOutput, error) {
    var output kafka.CreateClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaCreateConfigurationResult struct {
	Result workflow.Future
}

func (r *KafkaCreateConfigurationResult) Get(ctx workflow.Context) (*kafka.CreateConfigurationOutput, error) {
    var output kafka.CreateConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaDeleteClusterResult struct {
	Result workflow.Future
}

func (r *KafkaDeleteClusterResult) Get(ctx workflow.Context) (*kafka.DeleteClusterOutput, error) {
    var output kafka.DeleteClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaDeleteConfigurationResult struct {
	Result workflow.Future
}

func (r *KafkaDeleteConfigurationResult) Get(ctx workflow.Context) (*kafka.DeleteConfigurationOutput, error) {
    var output kafka.DeleteConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaDescribeClusterResult struct {
	Result workflow.Future
}

func (r *KafkaDescribeClusterResult) Get(ctx workflow.Context) (*kafka.DescribeClusterOutput, error) {
    var output kafka.DescribeClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaDescribeClusterOperationResult struct {
	Result workflow.Future
}

func (r *KafkaDescribeClusterOperationResult) Get(ctx workflow.Context) (*kafka.DescribeClusterOperationOutput, error) {
    var output kafka.DescribeClusterOperationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaDescribeConfigurationResult struct {
	Result workflow.Future
}

func (r *KafkaDescribeConfigurationResult) Get(ctx workflow.Context) (*kafka.DescribeConfigurationOutput, error) {
    var output kafka.DescribeConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaDescribeConfigurationRevisionResult struct {
	Result workflow.Future
}

func (r *KafkaDescribeConfigurationRevisionResult) Get(ctx workflow.Context) (*kafka.DescribeConfigurationRevisionOutput, error) {
    var output kafka.DescribeConfigurationRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaGetBootstrapBrokersResult struct {
	Result workflow.Future
}

func (r *KafkaGetBootstrapBrokersResult) Get(ctx workflow.Context) (*kafka.GetBootstrapBrokersOutput, error) {
    var output kafka.GetBootstrapBrokersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaGetCompatibleKafkaVersionsResult struct {
	Result workflow.Future
}

func (r *KafkaGetCompatibleKafkaVersionsResult) Get(ctx workflow.Context) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
    var output kafka.GetCompatibleKafkaVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaListClusterOperationsResult struct {
	Result workflow.Future
}

func (r *KafkaListClusterOperationsResult) Get(ctx workflow.Context) (*kafka.ListClusterOperationsOutput, error) {
    var output kafka.ListClusterOperationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaListClustersResult struct {
	Result workflow.Future
}

func (r *KafkaListClustersResult) Get(ctx workflow.Context) (*kafka.ListClustersOutput, error) {
    var output kafka.ListClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaListConfigurationRevisionsResult struct {
	Result workflow.Future
}

func (r *KafkaListConfigurationRevisionsResult) Get(ctx workflow.Context) (*kafka.ListConfigurationRevisionsOutput, error) {
    var output kafka.ListConfigurationRevisionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaListConfigurationsResult struct {
	Result workflow.Future
}

func (r *KafkaListConfigurationsResult) Get(ctx workflow.Context) (*kafka.ListConfigurationsOutput, error) {
    var output kafka.ListConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaListKafkaVersionsResult struct {
	Result workflow.Future
}

func (r *KafkaListKafkaVersionsResult) Get(ctx workflow.Context) (*kafka.ListKafkaVersionsOutput, error) {
    var output kafka.ListKafkaVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaListNodesResult struct {
	Result workflow.Future
}

func (r *KafkaListNodesResult) Get(ctx workflow.Context) (*kafka.ListNodesOutput, error) {
    var output kafka.ListNodesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *KafkaListTagsForResourceResult) Get(ctx workflow.Context) (*kafka.ListTagsForResourceOutput, error) {
    var output kafka.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaRebootBrokerResult struct {
	Result workflow.Future
}

func (r *KafkaRebootBrokerResult) Get(ctx workflow.Context) (*kafka.RebootBrokerOutput, error) {
    var output kafka.RebootBrokerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaTagResourceResult struct {
	Result workflow.Future
}

func (r *KafkaTagResourceResult) Get(ctx workflow.Context) (*kafka.TagResourceOutput, error) {
    var output kafka.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaUntagResourceResult struct {
	Result workflow.Future
}

func (r *KafkaUntagResourceResult) Get(ctx workflow.Context) (*kafka.UntagResourceOutput, error) {
    var output kafka.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaUpdateBrokerCountResult struct {
	Result workflow.Future
}

func (r *KafkaUpdateBrokerCountResult) Get(ctx workflow.Context) (*kafka.UpdateBrokerCountOutput, error) {
    var output kafka.UpdateBrokerCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaUpdateBrokerStorageResult struct {
	Result workflow.Future
}

func (r *KafkaUpdateBrokerStorageResult) Get(ctx workflow.Context) (*kafka.UpdateBrokerStorageOutput, error) {
    var output kafka.UpdateBrokerStorageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaUpdateClusterConfigurationResult struct {
	Result workflow.Future
}

func (r *KafkaUpdateClusterConfigurationResult) Get(ctx workflow.Context) (*kafka.UpdateClusterConfigurationOutput, error) {
    var output kafka.UpdateClusterConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaUpdateClusterKafkaVersionResult struct {
	Result workflow.Future
}

func (r *KafkaUpdateClusterKafkaVersionResult) Get(ctx workflow.Context) (*kafka.UpdateClusterKafkaVersionOutput, error) {
    var output kafka.UpdateClusterKafkaVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaUpdateConfigurationResult struct {
	Result workflow.Future
}

func (r *KafkaUpdateConfigurationResult) Get(ctx workflow.Context) (*kafka.UpdateConfigurationOutput, error) {
    var output kafka.UpdateConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KafkaUpdateMonitoringResult struct {
	Result workflow.Future
}

func (r *KafkaUpdateMonitoringResult) Get(ctx workflow.Context) (*kafka.UpdateMonitoringOutput, error) {
    var output kafka.UpdateMonitoringOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type KafkaStub struct {
    activities awsactivities.KafkaActivities
}

func NewKafkaStub() KafkaClient {
    return &KafkaStub{}
}
func (a *KafkaStub) CreateCluster(ctx workflow.Context, input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error) {
    var output kafka.CreateClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) CreateClusterAsync(ctx workflow.Context, input *kafka.CreateClusterInput) *KafkaCreateClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input)
    return &KafkaCreateClusterResult{Result: future}
}
func (a *KafkaStub) CreateConfiguration(ctx workflow.Context, input *kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error) {
    var output kafka.CreateConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) CreateConfigurationAsync(ctx workflow.Context, input *kafka.CreateConfigurationInput) *KafkaCreateConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateConfiguration, input)
    return &KafkaCreateConfigurationResult{Result: future}
}
func (a *KafkaStub) DeleteCluster(ctx workflow.Context, input *kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error) {
    var output kafka.DeleteClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) DeleteClusterAsync(ctx workflow.Context, input *kafka.DeleteClusterInput) *KafkaDeleteClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input)
    return &KafkaDeleteClusterResult{Result: future}
}
func (a *KafkaStub) DeleteConfiguration(ctx workflow.Context, input *kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error) {
    var output kafka.DeleteConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) DeleteConfigurationAsync(ctx workflow.Context, input *kafka.DeleteConfigurationInput) *KafkaDeleteConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfiguration, input)
    return &KafkaDeleteConfigurationResult{Result: future}
}
func (a *KafkaStub) DescribeCluster(ctx workflow.Context, input *kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error) {
    var output kafka.DescribeClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) DescribeClusterAsync(ctx workflow.Context, input *kafka.DescribeClusterInput) *KafkaDescribeClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCluster, input)
    return &KafkaDescribeClusterResult{Result: future}
}
func (a *KafkaStub) DescribeClusterOperation(ctx workflow.Context, input *kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error) {
    var output kafka.DescribeClusterOperationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterOperation, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) DescribeClusterOperationAsync(ctx workflow.Context, input *kafka.DescribeClusterOperationInput) *KafkaDescribeClusterOperationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusterOperation, input)
    return &KafkaDescribeClusterOperationResult{Result: future}
}
func (a *KafkaStub) DescribeConfiguration(ctx workflow.Context, input *kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error) {
    var output kafka.DescribeConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) DescribeConfigurationAsync(ctx workflow.Context, input *kafka.DescribeConfigurationInput) *KafkaDescribeConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfiguration, input)
    return &KafkaDescribeConfigurationResult{Result: future}
}
func (a *KafkaStub) DescribeConfigurationRevision(ctx workflow.Context, input *kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error) {
    var output kafka.DescribeConfigurationRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) DescribeConfigurationRevisionAsync(ctx workflow.Context, input *kafka.DescribeConfigurationRevisionInput) *KafkaDescribeConfigurationRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationRevision, input)
    return &KafkaDescribeConfigurationRevisionResult{Result: future}
}
func (a *KafkaStub) GetBootstrapBrokers(ctx workflow.Context, input *kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error) {
    var output kafka.GetBootstrapBrokersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBootstrapBrokers, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) GetBootstrapBrokersAsync(ctx workflow.Context, input *kafka.GetBootstrapBrokersInput) *KafkaGetBootstrapBrokersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBootstrapBrokers, input)
    return &KafkaGetBootstrapBrokersResult{Result: future}
}
func (a *KafkaStub) GetCompatibleKafkaVersions(ctx workflow.Context, input *kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
    var output kafka.GetCompatibleKafkaVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCompatibleKafkaVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) GetCompatibleKafkaVersionsAsync(ctx workflow.Context, input *kafka.GetCompatibleKafkaVersionsInput) *KafkaGetCompatibleKafkaVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCompatibleKafkaVersions, input)
    return &KafkaGetCompatibleKafkaVersionsResult{Result: future}
}
func (a *KafkaStub) ListClusterOperations(ctx workflow.Context, input *kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error) {
    var output kafka.ListClusterOperationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListClusterOperations, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) ListClusterOperationsAsync(ctx workflow.Context, input *kafka.ListClusterOperationsInput) *KafkaListClusterOperationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListClusterOperations, input)
    return &KafkaListClusterOperationsResult{Result: future}
}
func (a *KafkaStub) ListClusters(ctx workflow.Context, input *kafka.ListClustersInput) (*kafka.ListClustersOutput, error) {
    var output kafka.ListClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) ListClustersAsync(ctx workflow.Context, input *kafka.ListClustersInput) *KafkaListClustersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListClusters, input)
    return &KafkaListClustersResult{Result: future}
}
func (a *KafkaStub) ListConfigurationRevisions(ctx workflow.Context, input *kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error) {
    var output kafka.ListConfigurationRevisionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationRevisions, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) ListConfigurationRevisionsAsync(ctx workflow.Context, input *kafka.ListConfigurationRevisionsInput) *KafkaListConfigurationRevisionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationRevisions, input)
    return &KafkaListConfigurationRevisionsResult{Result: future}
}
func (a *KafkaStub) ListConfigurations(ctx workflow.Context, input *kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error) {
    var output kafka.ListConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) ListConfigurationsAsync(ctx workflow.Context, input *kafka.ListConfigurationsInput) *KafkaListConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListConfigurations, input)
    return &KafkaListConfigurationsResult{Result: future}
}
func (a *KafkaStub) ListKafkaVersions(ctx workflow.Context, input *kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error) {
    var output kafka.ListKafkaVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListKafkaVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) ListKafkaVersionsAsync(ctx workflow.Context, input *kafka.ListKafkaVersionsInput) *KafkaListKafkaVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListKafkaVersions, input)
    return &KafkaListKafkaVersionsResult{Result: future}
}
func (a *KafkaStub) ListNodes(ctx workflow.Context, input *kafka.ListNodesInput) (*kafka.ListNodesOutput, error) {
    var output kafka.ListNodesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNodes, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) ListNodesAsync(ctx workflow.Context, input *kafka.ListNodesInput) *KafkaListNodesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListNodes, input)
    return &KafkaListNodesResult{Result: future}
}
func (a *KafkaStub) ListTagsForResource(ctx workflow.Context, input *kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error) {
    var output kafka.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) ListTagsForResourceAsync(ctx workflow.Context, input *kafka.ListTagsForResourceInput) *KafkaListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &KafkaListTagsForResourceResult{Result: future}
}
func (a *KafkaStub) RebootBroker(ctx workflow.Context, input *kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error) {
    var output kafka.RebootBrokerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootBroker, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) RebootBrokerAsync(ctx workflow.Context, input *kafka.RebootBrokerInput) *KafkaRebootBrokerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebootBroker, input)
    return &KafkaRebootBrokerResult{Result: future}
}
func (a *KafkaStub) TagResource(ctx workflow.Context, input *kafka.TagResourceInput) (*kafka.TagResourceOutput, error) {
    var output kafka.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) TagResourceAsync(ctx workflow.Context, input *kafka.TagResourceInput) *KafkaTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &KafkaTagResourceResult{Result: future}
}
func (a *KafkaStub) UntagResource(ctx workflow.Context, input *kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error) {
    var output kafka.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) UntagResourceAsync(ctx workflow.Context, input *kafka.UntagResourceInput) *KafkaUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &KafkaUntagResourceResult{Result: future}
}
func (a *KafkaStub) UpdateBrokerCount(ctx workflow.Context, input *kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error) {
    var output kafka.UpdateBrokerCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBrokerCount, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) UpdateBrokerCountAsync(ctx workflow.Context, input *kafka.UpdateBrokerCountInput) *KafkaUpdateBrokerCountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBrokerCount, input)
    return &KafkaUpdateBrokerCountResult{Result: future}
}
func (a *KafkaStub) UpdateBrokerStorage(ctx workflow.Context, input *kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error) {
    var output kafka.UpdateBrokerStorageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBrokerStorage, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) UpdateBrokerStorageAsync(ctx workflow.Context, input *kafka.UpdateBrokerStorageInput) *KafkaUpdateBrokerStorageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBrokerStorage, input)
    return &KafkaUpdateBrokerStorageResult{Result: future}
}
func (a *KafkaStub) UpdateClusterConfiguration(ctx workflow.Context, input *kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error) {
    var output kafka.UpdateClusterConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateClusterConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) UpdateClusterConfigurationAsync(ctx workflow.Context, input *kafka.UpdateClusterConfigurationInput) *KafkaUpdateClusterConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateClusterConfiguration, input)
    return &KafkaUpdateClusterConfigurationResult{Result: future}
}
func (a *KafkaStub) UpdateClusterKafkaVersion(ctx workflow.Context, input *kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error) {
    var output kafka.UpdateClusterKafkaVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateClusterKafkaVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) UpdateClusterKafkaVersionAsync(ctx workflow.Context, input *kafka.UpdateClusterKafkaVersionInput) *KafkaUpdateClusterKafkaVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateClusterKafkaVersion, input)
    return &KafkaUpdateClusterKafkaVersionResult{Result: future}
}
func (a *KafkaStub) UpdateConfiguration(ctx workflow.Context, input *kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error) {
    var output kafka.UpdateConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) UpdateConfigurationAsync(ctx workflow.Context, input *kafka.UpdateConfigurationInput) *KafkaUpdateConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateConfiguration, input)
    return &KafkaUpdateConfigurationResult{Result: future}
}
func (a *KafkaStub) UpdateMonitoring(ctx workflow.Context, input *kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error) {
    var output kafka.UpdateMonitoringOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMonitoring, input).Get(ctx, &output)
    return &output, err
}

func (a *KafkaStub) UpdateMonitoringAsync(ctx workflow.Context, input *kafka.UpdateMonitoringInput) *KafkaUpdateMonitoringResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMonitoring, input)
    return &KafkaUpdateMonitoringResult{Result: future}
}
