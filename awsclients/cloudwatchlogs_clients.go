package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CloudWatchLogsClient interface {
    AssociateKmsKey(ctx workflow.Context, input *cloudwatchlogs.AssociateKmsKeyInput) (*cloudwatchlogs.AssociateKmsKeyOutput, error)
    AssociateKmsKeyAsync(ctx workflow.Context, input *cloudwatchlogs.AssociateKmsKeyInput) *CloudwatchlogsAssociateKmsKeyResult

    CancelExportTask(ctx workflow.Context, input *cloudwatchlogs.CancelExportTaskInput) (*cloudwatchlogs.CancelExportTaskOutput, error)
    CancelExportTaskAsync(ctx workflow.Context, input *cloudwatchlogs.CancelExportTaskInput) *CloudwatchlogsCancelExportTaskResult

    CreateExportTask(ctx workflow.Context, input *cloudwatchlogs.CreateExportTaskInput) (*cloudwatchlogs.CreateExportTaskOutput, error)
    CreateExportTaskAsync(ctx workflow.Context, input *cloudwatchlogs.CreateExportTaskInput) *CloudwatchlogsCreateExportTaskResult

    CreateLogGroup(ctx workflow.Context, input *cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error)
    CreateLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.CreateLogGroupInput) *CloudwatchlogsCreateLogGroupResult

    CreateLogStream(ctx workflow.Context, input *cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error)
    CreateLogStreamAsync(ctx workflow.Context, input *cloudwatchlogs.CreateLogStreamInput) *CloudwatchlogsCreateLogStreamResult

    DeleteDestination(ctx workflow.Context, input *cloudwatchlogs.DeleteDestinationInput) (*cloudwatchlogs.DeleteDestinationOutput, error)
    DeleteDestinationAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteDestinationInput) *CloudwatchlogsDeleteDestinationResult

    DeleteLogGroup(ctx workflow.Context, input *cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, error)
    DeleteLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteLogGroupInput) *CloudwatchlogsDeleteLogGroupResult

    DeleteLogStream(ctx workflow.Context, input *cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error)
    DeleteLogStreamAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteLogStreamInput) *CloudwatchlogsDeleteLogStreamResult

    DeleteMetricFilter(ctx workflow.Context, input *cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, error)
    DeleteMetricFilterAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteMetricFilterInput) *CloudwatchlogsDeleteMetricFilterResult

    DeleteQueryDefinition(ctx workflow.Context, input *cloudwatchlogs.DeleteQueryDefinitionInput) (*cloudwatchlogs.DeleteQueryDefinitionOutput, error)
    DeleteQueryDefinitionAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteQueryDefinitionInput) *CloudwatchlogsDeleteQueryDefinitionResult

    DeleteResourcePolicy(ctx workflow.Context, input *cloudwatchlogs.DeleteResourcePolicyInput) (*cloudwatchlogs.DeleteResourcePolicyOutput, error)
    DeleteResourcePolicyAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteResourcePolicyInput) *CloudwatchlogsDeleteResourcePolicyResult

    DeleteRetentionPolicy(ctx workflow.Context, input *cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error)
    DeleteRetentionPolicyAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteRetentionPolicyInput) *CloudwatchlogsDeleteRetentionPolicyResult

    DeleteSubscriptionFilter(ctx workflow.Context, input *cloudwatchlogs.DeleteSubscriptionFilterInput) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error)
    DeleteSubscriptionFilterAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteSubscriptionFilterInput) *CloudwatchlogsDeleteSubscriptionFilterResult

    DescribeDestinations(ctx workflow.Context, input *cloudwatchlogs.DescribeDestinationsInput) (*cloudwatchlogs.DescribeDestinationsOutput, error)
    DescribeDestinationsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeDestinationsInput) *CloudwatchlogsDescribeDestinationsResult

    DescribeExportTasks(ctx workflow.Context, input *cloudwatchlogs.DescribeExportTasksInput) (*cloudwatchlogs.DescribeExportTasksOutput, error)
    DescribeExportTasksAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeExportTasksInput) *CloudwatchlogsDescribeExportTasksResult

    DescribeLogGroups(ctx workflow.Context, input *cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error)
    DescribeLogGroupsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeLogGroupsInput) *CloudwatchlogsDescribeLogGroupsResult

    DescribeLogStreams(ctx workflow.Context, input *cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error)
    DescribeLogStreamsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeLogStreamsInput) *CloudwatchlogsDescribeLogStreamsResult

    DescribeMetricFilters(ctx workflow.Context, input *cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)
    DescribeMetricFiltersAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeMetricFiltersInput) *CloudwatchlogsDescribeMetricFiltersResult

    DescribeQueries(ctx workflow.Context, input *cloudwatchlogs.DescribeQueriesInput) (*cloudwatchlogs.DescribeQueriesOutput, error)
    DescribeQueriesAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeQueriesInput) *CloudwatchlogsDescribeQueriesResult

    DescribeQueryDefinitions(ctx workflow.Context, input *cloudwatchlogs.DescribeQueryDefinitionsInput) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error)
    DescribeQueryDefinitionsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeQueryDefinitionsInput) *CloudwatchlogsDescribeQueryDefinitionsResult

    DescribeResourcePolicies(ctx workflow.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error)
    DescribeResourcePoliciesAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput) *CloudwatchlogsDescribeResourcePoliciesResult

    DescribeSubscriptionFilters(ctx workflow.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error)
    DescribeSubscriptionFiltersAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput) *CloudwatchlogsDescribeSubscriptionFiltersResult

    DisassociateKmsKey(ctx workflow.Context, input *cloudwatchlogs.DisassociateKmsKeyInput) (*cloudwatchlogs.DisassociateKmsKeyOutput, error)
    DisassociateKmsKeyAsync(ctx workflow.Context, input *cloudwatchlogs.DisassociateKmsKeyInput) *CloudwatchlogsDisassociateKmsKeyResult

    FilterLogEvents(ctx workflow.Context, input *cloudwatchlogs.FilterLogEventsInput) (*cloudwatchlogs.FilterLogEventsOutput, error)
    FilterLogEventsAsync(ctx workflow.Context, input *cloudwatchlogs.FilterLogEventsInput) *CloudwatchlogsFilterLogEventsResult

    GetLogEvents(ctx workflow.Context, input *cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error)
    GetLogEventsAsync(ctx workflow.Context, input *cloudwatchlogs.GetLogEventsInput) *CloudwatchlogsGetLogEventsResult

    GetLogGroupFields(ctx workflow.Context, input *cloudwatchlogs.GetLogGroupFieldsInput) (*cloudwatchlogs.GetLogGroupFieldsOutput, error)
    GetLogGroupFieldsAsync(ctx workflow.Context, input *cloudwatchlogs.GetLogGroupFieldsInput) *CloudwatchlogsGetLogGroupFieldsResult

    GetLogRecord(ctx workflow.Context, input *cloudwatchlogs.GetLogRecordInput) (*cloudwatchlogs.GetLogRecordOutput, error)
    GetLogRecordAsync(ctx workflow.Context, input *cloudwatchlogs.GetLogRecordInput) *CloudwatchlogsGetLogRecordResult

    GetQueryResults(ctx workflow.Context, input *cloudwatchlogs.GetQueryResultsInput) (*cloudwatchlogs.GetQueryResultsOutput, error)
    GetQueryResultsAsync(ctx workflow.Context, input *cloudwatchlogs.GetQueryResultsInput) *CloudwatchlogsGetQueryResultsResult

    ListTagsLogGroup(ctx workflow.Context, input *cloudwatchlogs.ListTagsLogGroupInput) (*cloudwatchlogs.ListTagsLogGroupOutput, error)
    ListTagsLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.ListTagsLogGroupInput) *CloudwatchlogsListTagsLogGroupResult

    PutDestination(ctx workflow.Context, input *cloudwatchlogs.PutDestinationInput) (*cloudwatchlogs.PutDestinationOutput, error)
    PutDestinationAsync(ctx workflow.Context, input *cloudwatchlogs.PutDestinationInput) *CloudwatchlogsPutDestinationResult

    PutDestinationPolicy(ctx workflow.Context, input *cloudwatchlogs.PutDestinationPolicyInput) (*cloudwatchlogs.PutDestinationPolicyOutput, error)
    PutDestinationPolicyAsync(ctx workflow.Context, input *cloudwatchlogs.PutDestinationPolicyInput) *CloudwatchlogsPutDestinationPolicyResult

    PutLogEvents(ctx workflow.Context, input *cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error)
    PutLogEventsAsync(ctx workflow.Context, input *cloudwatchlogs.PutLogEventsInput) *CloudwatchlogsPutLogEventsResult

    PutMetricFilter(ctx workflow.Context, input *cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, error)
    PutMetricFilterAsync(ctx workflow.Context, input *cloudwatchlogs.PutMetricFilterInput) *CloudwatchlogsPutMetricFilterResult

    PutQueryDefinition(ctx workflow.Context, input *cloudwatchlogs.PutQueryDefinitionInput) (*cloudwatchlogs.PutQueryDefinitionOutput, error)
    PutQueryDefinitionAsync(ctx workflow.Context, input *cloudwatchlogs.PutQueryDefinitionInput) *CloudwatchlogsPutQueryDefinitionResult

    PutResourcePolicy(ctx workflow.Context, input *cloudwatchlogs.PutResourcePolicyInput) (*cloudwatchlogs.PutResourcePolicyOutput, error)
    PutResourcePolicyAsync(ctx workflow.Context, input *cloudwatchlogs.PutResourcePolicyInput) *CloudwatchlogsPutResourcePolicyResult

    PutRetentionPolicy(ctx workflow.Context, input *cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error)
    PutRetentionPolicyAsync(ctx workflow.Context, input *cloudwatchlogs.PutRetentionPolicyInput) *CloudwatchlogsPutRetentionPolicyResult

    PutSubscriptionFilter(ctx workflow.Context, input *cloudwatchlogs.PutSubscriptionFilterInput) (*cloudwatchlogs.PutSubscriptionFilterOutput, error)
    PutSubscriptionFilterAsync(ctx workflow.Context, input *cloudwatchlogs.PutSubscriptionFilterInput) *CloudwatchlogsPutSubscriptionFilterResult

    StartQuery(ctx workflow.Context, input *cloudwatchlogs.StartQueryInput) (*cloudwatchlogs.StartQueryOutput, error)
    StartQueryAsync(ctx workflow.Context, input *cloudwatchlogs.StartQueryInput) *CloudwatchlogsStartQueryResult

    StopQuery(ctx workflow.Context, input *cloudwatchlogs.StopQueryInput) (*cloudwatchlogs.StopQueryOutput, error)
    StopQueryAsync(ctx workflow.Context, input *cloudwatchlogs.StopQueryInput) *CloudwatchlogsStopQueryResult

    TagLogGroup(ctx workflow.Context, input *cloudwatchlogs.TagLogGroupInput) (*cloudwatchlogs.TagLogGroupOutput, error)
    TagLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.TagLogGroupInput) *CloudwatchlogsTagLogGroupResult

    TestMetricFilter(ctx workflow.Context, input *cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, error)
    TestMetricFilterAsync(ctx workflow.Context, input *cloudwatchlogs.TestMetricFilterInput) *CloudwatchlogsTestMetricFilterResult

    UntagLogGroup(ctx workflow.Context, input *cloudwatchlogs.UntagLogGroupInput) (*cloudwatchlogs.UntagLogGroupOutput, error)
    UntagLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.UntagLogGroupInput) *CloudwatchlogsUntagLogGroupResult
}
type CloudwatchlogsAssociateKmsKeyResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsAssociateKmsKeyResult) Get(ctx workflow.Context) (*cloudwatchlogs.AssociateKmsKeyOutput, error) {
    var output cloudwatchlogs.AssociateKmsKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsCancelExportTaskResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsCancelExportTaskResult) Get(ctx workflow.Context) (*cloudwatchlogs.CancelExportTaskOutput, error) {
    var output cloudwatchlogs.CancelExportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsCreateExportTaskResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsCreateExportTaskResult) Get(ctx workflow.Context) (*cloudwatchlogs.CreateExportTaskOutput, error) {
    var output cloudwatchlogs.CreateExportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsCreateLogGroupResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsCreateLogGroupResult) Get(ctx workflow.Context) (*cloudwatchlogs.CreateLogGroupOutput, error) {
    var output cloudwatchlogs.CreateLogGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsCreateLogStreamResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsCreateLogStreamResult) Get(ctx workflow.Context) (*cloudwatchlogs.CreateLogStreamOutput, error) {
    var output cloudwatchlogs.CreateLogStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteDestinationResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteDestinationResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteDestinationOutput, error) {
    var output cloudwatchlogs.DeleteDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteLogGroupResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteLogGroupResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteLogGroupOutput, error) {
    var output cloudwatchlogs.DeleteLogGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteLogStreamResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteLogStreamResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteLogStreamOutput, error) {
    var output cloudwatchlogs.DeleteLogStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteMetricFilterResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteMetricFilterResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteMetricFilterOutput, error) {
    var output cloudwatchlogs.DeleteMetricFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteQueryDefinitionResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteQueryDefinitionResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteQueryDefinitionOutput, error) {
    var output cloudwatchlogs.DeleteQueryDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteResourcePolicyResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteResourcePolicyResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteResourcePolicyOutput, error) {
    var output cloudwatchlogs.DeleteResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteRetentionPolicyResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteRetentionPolicyResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error) {
    var output cloudwatchlogs.DeleteRetentionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDeleteSubscriptionFilterResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDeleteSubscriptionFilterResult) Get(ctx workflow.Context) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error) {
    var output cloudwatchlogs.DeleteSubscriptionFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeDestinationsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeDestinationsResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
    var output cloudwatchlogs.DescribeDestinationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeExportTasksResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeExportTasksResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
    var output cloudwatchlogs.DescribeExportTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeLogGroupsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeLogGroupsResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
    var output cloudwatchlogs.DescribeLogGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeLogStreamsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeLogStreamsResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
    var output cloudwatchlogs.DescribeLogStreamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeMetricFiltersResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeMetricFiltersResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
    var output cloudwatchlogs.DescribeMetricFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeQueriesResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeQueriesResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeQueriesOutput, error) {
    var output cloudwatchlogs.DescribeQueriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeQueryDefinitionsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeQueryDefinitionsResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {
    var output cloudwatchlogs.DescribeQueryDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeResourcePoliciesResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeResourcePoliciesResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
    var output cloudwatchlogs.DescribeResourcePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDescribeSubscriptionFiltersResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDescribeSubscriptionFiltersResult) Get(ctx workflow.Context) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
    var output cloudwatchlogs.DescribeSubscriptionFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsDisassociateKmsKeyResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsDisassociateKmsKeyResult) Get(ctx workflow.Context) (*cloudwatchlogs.DisassociateKmsKeyOutput, error) {
    var output cloudwatchlogs.DisassociateKmsKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsFilterLogEventsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsFilterLogEventsResult) Get(ctx workflow.Context) (*cloudwatchlogs.FilterLogEventsOutput, error) {
    var output cloudwatchlogs.FilterLogEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsGetLogEventsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsGetLogEventsResult) Get(ctx workflow.Context) (*cloudwatchlogs.GetLogEventsOutput, error) {
    var output cloudwatchlogs.GetLogEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsGetLogGroupFieldsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsGetLogGroupFieldsResult) Get(ctx workflow.Context) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
    var output cloudwatchlogs.GetLogGroupFieldsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsGetLogRecordResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsGetLogRecordResult) Get(ctx workflow.Context) (*cloudwatchlogs.GetLogRecordOutput, error) {
    var output cloudwatchlogs.GetLogRecordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsGetQueryResultsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsGetQueryResultsResult) Get(ctx workflow.Context) (*cloudwatchlogs.GetQueryResultsOutput, error) {
    var output cloudwatchlogs.GetQueryResultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsListTagsLogGroupResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsListTagsLogGroupResult) Get(ctx workflow.Context) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
    var output cloudwatchlogs.ListTagsLogGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutDestinationResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutDestinationResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutDestinationOutput, error) {
    var output cloudwatchlogs.PutDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutDestinationPolicyResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutDestinationPolicyResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutDestinationPolicyOutput, error) {
    var output cloudwatchlogs.PutDestinationPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutLogEventsResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutLogEventsResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutLogEventsOutput, error) {
    var output cloudwatchlogs.PutLogEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutMetricFilterResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutMetricFilterResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutMetricFilterOutput, error) {
    var output cloudwatchlogs.PutMetricFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutQueryDefinitionResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutQueryDefinitionResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutQueryDefinitionOutput, error) {
    var output cloudwatchlogs.PutQueryDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutResourcePolicyResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutResourcePolicyResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutResourcePolicyOutput, error) {
    var output cloudwatchlogs.PutResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutRetentionPolicyResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutRetentionPolicyResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
    var output cloudwatchlogs.PutRetentionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsPutSubscriptionFilterResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsPutSubscriptionFilterResult) Get(ctx workflow.Context) (*cloudwatchlogs.PutSubscriptionFilterOutput, error) {
    var output cloudwatchlogs.PutSubscriptionFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsStartQueryResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsStartQueryResult) Get(ctx workflow.Context) (*cloudwatchlogs.StartQueryOutput, error) {
    var output cloudwatchlogs.StartQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsStopQueryResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsStopQueryResult) Get(ctx workflow.Context) (*cloudwatchlogs.StopQueryOutput, error) {
    var output cloudwatchlogs.StopQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsTagLogGroupResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsTagLogGroupResult) Get(ctx workflow.Context) (*cloudwatchlogs.TagLogGroupOutput, error) {
    var output cloudwatchlogs.TagLogGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsTestMetricFilterResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsTestMetricFilterResult) Get(ctx workflow.Context) (*cloudwatchlogs.TestMetricFilterOutput, error) {
    var output cloudwatchlogs.TestMetricFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudwatchlogsUntagLogGroupResult struct {
	Result workflow.Future
}

func (r *CloudwatchlogsUntagLogGroupResult) Get(ctx workflow.Context) (*cloudwatchlogs.UntagLogGroupOutput, error) {
    var output cloudwatchlogs.UntagLogGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CloudWatchLogsStub struct {
    activities awsactivities.CloudWatchLogsActivities
}

func NewCloudWatchLogsStub() CloudWatchLogsClient {
    return &CloudWatchLogsStub{}
}
func (a *CloudWatchLogsStub) AssociateKmsKey(ctx workflow.Context, input *cloudwatchlogs.AssociateKmsKeyInput) (*cloudwatchlogs.AssociateKmsKeyOutput, error) {
    var output cloudwatchlogs.AssociateKmsKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateKmsKey, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) AssociateKmsKeyAsync(ctx workflow.Context, input *cloudwatchlogs.AssociateKmsKeyInput) *CloudwatchlogsAssociateKmsKeyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateKmsKey, input)
    return &CloudwatchlogsAssociateKmsKeyResult{Result: future}
}
func (a *CloudWatchLogsStub) CancelExportTask(ctx workflow.Context, input *cloudwatchlogs.CancelExportTaskInput) (*cloudwatchlogs.CancelExportTaskOutput, error) {
    var output cloudwatchlogs.CancelExportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelExportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) CancelExportTaskAsync(ctx workflow.Context, input *cloudwatchlogs.CancelExportTaskInput) *CloudwatchlogsCancelExportTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelExportTask, input)
    return &CloudwatchlogsCancelExportTaskResult{Result: future}
}
func (a *CloudWatchLogsStub) CreateExportTask(ctx workflow.Context, input *cloudwatchlogs.CreateExportTaskInput) (*cloudwatchlogs.CreateExportTaskOutput, error) {
    var output cloudwatchlogs.CreateExportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateExportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) CreateExportTaskAsync(ctx workflow.Context, input *cloudwatchlogs.CreateExportTaskInput) *CloudwatchlogsCreateExportTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateExportTask, input)
    return &CloudwatchlogsCreateExportTaskResult{Result: future}
}
func (a *CloudWatchLogsStub) CreateLogGroup(ctx workflow.Context, input *cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error) {
    var output cloudwatchlogs.CreateLogGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLogGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) CreateLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.CreateLogGroupInput) *CloudwatchlogsCreateLogGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLogGroup, input)
    return &CloudwatchlogsCreateLogGroupResult{Result: future}
}
func (a *CloudWatchLogsStub) CreateLogStream(ctx workflow.Context, input *cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error) {
    var output cloudwatchlogs.CreateLogStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLogStream, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) CreateLogStreamAsync(ctx workflow.Context, input *cloudwatchlogs.CreateLogStreamInput) *CloudwatchlogsCreateLogStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLogStream, input)
    return &CloudwatchlogsCreateLogStreamResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteDestination(ctx workflow.Context, input *cloudwatchlogs.DeleteDestinationInput) (*cloudwatchlogs.DeleteDestinationOutput, error) {
    var output cloudwatchlogs.DeleteDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteDestinationAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteDestinationInput) *CloudwatchlogsDeleteDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDestination, input)
    return &CloudwatchlogsDeleteDestinationResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteLogGroup(ctx workflow.Context, input *cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, error) {
    var output cloudwatchlogs.DeleteLogGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLogGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteLogGroupInput) *CloudwatchlogsDeleteLogGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLogGroup, input)
    return &CloudwatchlogsDeleteLogGroupResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteLogStream(ctx workflow.Context, input *cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error) {
    var output cloudwatchlogs.DeleteLogStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLogStream, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteLogStreamAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteLogStreamInput) *CloudwatchlogsDeleteLogStreamResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLogStream, input)
    return &CloudwatchlogsDeleteLogStreamResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteMetricFilter(ctx workflow.Context, input *cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, error) {
    var output cloudwatchlogs.DeleteMetricFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMetricFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteMetricFilterAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteMetricFilterInput) *CloudwatchlogsDeleteMetricFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMetricFilter, input)
    return &CloudwatchlogsDeleteMetricFilterResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteQueryDefinition(ctx workflow.Context, input *cloudwatchlogs.DeleteQueryDefinitionInput) (*cloudwatchlogs.DeleteQueryDefinitionOutput, error) {
    var output cloudwatchlogs.DeleteQueryDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteQueryDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteQueryDefinitionAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteQueryDefinitionInput) *CloudwatchlogsDeleteQueryDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteQueryDefinition, input)
    return &CloudwatchlogsDeleteQueryDefinitionResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteResourcePolicy(ctx workflow.Context, input *cloudwatchlogs.DeleteResourcePolicyInput) (*cloudwatchlogs.DeleteResourcePolicyOutput, error) {
    var output cloudwatchlogs.DeleteResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteResourcePolicyAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteResourcePolicyInput) *CloudwatchlogsDeleteResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input)
    return &CloudwatchlogsDeleteResourcePolicyResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteRetentionPolicy(ctx workflow.Context, input *cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error) {
    var output cloudwatchlogs.DeleteRetentionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRetentionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteRetentionPolicyAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteRetentionPolicyInput) *CloudwatchlogsDeleteRetentionPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRetentionPolicy, input)
    return &CloudwatchlogsDeleteRetentionPolicyResult{Result: future}
}
func (a *CloudWatchLogsStub) DeleteSubscriptionFilter(ctx workflow.Context, input *cloudwatchlogs.DeleteSubscriptionFilterInput) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error) {
    var output cloudwatchlogs.DeleteSubscriptionFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSubscriptionFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DeleteSubscriptionFilterAsync(ctx workflow.Context, input *cloudwatchlogs.DeleteSubscriptionFilterInput) *CloudwatchlogsDeleteSubscriptionFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSubscriptionFilter, input)
    return &CloudwatchlogsDeleteSubscriptionFilterResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeDestinations(ctx workflow.Context, input *cloudwatchlogs.DescribeDestinationsInput) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
    var output cloudwatchlogs.DescribeDestinationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDestinations, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeDestinationsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeDestinationsInput) *CloudwatchlogsDescribeDestinationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDestinations, input)
    return &CloudwatchlogsDescribeDestinationsResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeExportTasks(ctx workflow.Context, input *cloudwatchlogs.DescribeExportTasksInput) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
    var output cloudwatchlogs.DescribeExportTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExportTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeExportTasksAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeExportTasksInput) *CloudwatchlogsDescribeExportTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeExportTasks, input)
    return &CloudwatchlogsDescribeExportTasksResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeLogGroups(ctx workflow.Context, input *cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
    var output cloudwatchlogs.DescribeLogGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLogGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeLogGroupsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeLogGroupsInput) *CloudwatchlogsDescribeLogGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLogGroups, input)
    return &CloudwatchlogsDescribeLogGroupsResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeLogStreams(ctx workflow.Context, input *cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
    var output cloudwatchlogs.DescribeLogStreamsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLogStreams, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeLogStreamsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeLogStreamsInput) *CloudwatchlogsDescribeLogStreamsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLogStreams, input)
    return &CloudwatchlogsDescribeLogStreamsResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeMetricFilters(ctx workflow.Context, input *cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
    var output cloudwatchlogs.DescribeMetricFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMetricFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeMetricFiltersAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeMetricFiltersInput) *CloudwatchlogsDescribeMetricFiltersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMetricFilters, input)
    return &CloudwatchlogsDescribeMetricFiltersResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeQueries(ctx workflow.Context, input *cloudwatchlogs.DescribeQueriesInput) (*cloudwatchlogs.DescribeQueriesOutput, error) {
    var output cloudwatchlogs.DescribeQueriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeQueries, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeQueriesAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeQueriesInput) *CloudwatchlogsDescribeQueriesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeQueries, input)
    return &CloudwatchlogsDescribeQueriesResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeQueryDefinitions(ctx workflow.Context, input *cloudwatchlogs.DescribeQueryDefinitionsInput) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {
    var output cloudwatchlogs.DescribeQueryDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeQueryDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeQueryDefinitionsAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeQueryDefinitionsInput) *CloudwatchlogsDescribeQueryDefinitionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeQueryDefinitions, input)
    return &CloudwatchlogsDescribeQueryDefinitionsResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeResourcePolicies(ctx workflow.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
    var output cloudwatchlogs.DescribeResourcePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeResourcePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeResourcePoliciesAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput) *CloudwatchlogsDescribeResourcePoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeResourcePolicies, input)
    return &CloudwatchlogsDescribeResourcePoliciesResult{Result: future}
}
func (a *CloudWatchLogsStub) DescribeSubscriptionFilters(ctx workflow.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
    var output cloudwatchlogs.DescribeSubscriptionFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSubscriptionFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DescribeSubscriptionFiltersAsync(ctx workflow.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput) *CloudwatchlogsDescribeSubscriptionFiltersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSubscriptionFilters, input)
    return &CloudwatchlogsDescribeSubscriptionFiltersResult{Result: future}
}
func (a *CloudWatchLogsStub) DisassociateKmsKey(ctx workflow.Context, input *cloudwatchlogs.DisassociateKmsKeyInput) (*cloudwatchlogs.DisassociateKmsKeyOutput, error) {
    var output cloudwatchlogs.DisassociateKmsKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateKmsKey, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) DisassociateKmsKeyAsync(ctx workflow.Context, input *cloudwatchlogs.DisassociateKmsKeyInput) *CloudwatchlogsDisassociateKmsKeyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateKmsKey, input)
    return &CloudwatchlogsDisassociateKmsKeyResult{Result: future}
}
func (a *CloudWatchLogsStub) FilterLogEvents(ctx workflow.Context, input *cloudwatchlogs.FilterLogEventsInput) (*cloudwatchlogs.FilterLogEventsOutput, error) {
    var output cloudwatchlogs.FilterLogEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.FilterLogEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) FilterLogEventsAsync(ctx workflow.Context, input *cloudwatchlogs.FilterLogEventsInput) *CloudwatchlogsFilterLogEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.FilterLogEvents, input)
    return &CloudwatchlogsFilterLogEventsResult{Result: future}
}
func (a *CloudWatchLogsStub) GetLogEvents(ctx workflow.Context, input *cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error) {
    var output cloudwatchlogs.GetLogEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLogEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) GetLogEventsAsync(ctx workflow.Context, input *cloudwatchlogs.GetLogEventsInput) *CloudwatchlogsGetLogEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLogEvents, input)
    return &CloudwatchlogsGetLogEventsResult{Result: future}
}
func (a *CloudWatchLogsStub) GetLogGroupFields(ctx workflow.Context, input *cloudwatchlogs.GetLogGroupFieldsInput) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
    var output cloudwatchlogs.GetLogGroupFieldsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLogGroupFields, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) GetLogGroupFieldsAsync(ctx workflow.Context, input *cloudwatchlogs.GetLogGroupFieldsInput) *CloudwatchlogsGetLogGroupFieldsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLogGroupFields, input)
    return &CloudwatchlogsGetLogGroupFieldsResult{Result: future}
}
func (a *CloudWatchLogsStub) GetLogRecord(ctx workflow.Context, input *cloudwatchlogs.GetLogRecordInput) (*cloudwatchlogs.GetLogRecordOutput, error) {
    var output cloudwatchlogs.GetLogRecordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLogRecord, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) GetLogRecordAsync(ctx workflow.Context, input *cloudwatchlogs.GetLogRecordInput) *CloudwatchlogsGetLogRecordResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLogRecord, input)
    return &CloudwatchlogsGetLogRecordResult{Result: future}
}
func (a *CloudWatchLogsStub) GetQueryResults(ctx workflow.Context, input *cloudwatchlogs.GetQueryResultsInput) (*cloudwatchlogs.GetQueryResultsOutput, error) {
    var output cloudwatchlogs.GetQueryResultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQueryResults, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) GetQueryResultsAsync(ctx workflow.Context, input *cloudwatchlogs.GetQueryResultsInput) *CloudwatchlogsGetQueryResultsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQueryResults, input)
    return &CloudwatchlogsGetQueryResultsResult{Result: future}
}
func (a *CloudWatchLogsStub) ListTagsLogGroup(ctx workflow.Context, input *cloudwatchlogs.ListTagsLogGroupInput) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
    var output cloudwatchlogs.ListTagsLogGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsLogGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) ListTagsLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.ListTagsLogGroupInput) *CloudwatchlogsListTagsLogGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsLogGroup, input)
    return &CloudwatchlogsListTagsLogGroupResult{Result: future}
}
func (a *CloudWatchLogsStub) PutDestination(ctx workflow.Context, input *cloudwatchlogs.PutDestinationInput) (*cloudwatchlogs.PutDestinationOutput, error) {
    var output cloudwatchlogs.PutDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutDestinationAsync(ctx workflow.Context, input *cloudwatchlogs.PutDestinationInput) *CloudwatchlogsPutDestinationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutDestination, input)
    return &CloudwatchlogsPutDestinationResult{Result: future}
}
func (a *CloudWatchLogsStub) PutDestinationPolicy(ctx workflow.Context, input *cloudwatchlogs.PutDestinationPolicyInput) (*cloudwatchlogs.PutDestinationPolicyOutput, error) {
    var output cloudwatchlogs.PutDestinationPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDestinationPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutDestinationPolicyAsync(ctx workflow.Context, input *cloudwatchlogs.PutDestinationPolicyInput) *CloudwatchlogsPutDestinationPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutDestinationPolicy, input)
    return &CloudwatchlogsPutDestinationPolicyResult{Result: future}
}
func (a *CloudWatchLogsStub) PutLogEvents(ctx workflow.Context, input *cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error) {
    var output cloudwatchlogs.PutLogEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLogEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutLogEventsAsync(ctx workflow.Context, input *cloudwatchlogs.PutLogEventsInput) *CloudwatchlogsPutLogEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutLogEvents, input)
    return &CloudwatchlogsPutLogEventsResult{Result: future}
}
func (a *CloudWatchLogsStub) PutMetricFilter(ctx workflow.Context, input *cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, error) {
    var output cloudwatchlogs.PutMetricFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutMetricFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutMetricFilterAsync(ctx workflow.Context, input *cloudwatchlogs.PutMetricFilterInput) *CloudwatchlogsPutMetricFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutMetricFilter, input)
    return &CloudwatchlogsPutMetricFilterResult{Result: future}
}
func (a *CloudWatchLogsStub) PutQueryDefinition(ctx workflow.Context, input *cloudwatchlogs.PutQueryDefinitionInput) (*cloudwatchlogs.PutQueryDefinitionOutput, error) {
    var output cloudwatchlogs.PutQueryDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutQueryDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutQueryDefinitionAsync(ctx workflow.Context, input *cloudwatchlogs.PutQueryDefinitionInput) *CloudwatchlogsPutQueryDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutQueryDefinition, input)
    return &CloudwatchlogsPutQueryDefinitionResult{Result: future}
}
func (a *CloudWatchLogsStub) PutResourcePolicy(ctx workflow.Context, input *cloudwatchlogs.PutResourcePolicyInput) (*cloudwatchlogs.PutResourcePolicyOutput, error) {
    var output cloudwatchlogs.PutResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutResourcePolicyAsync(ctx workflow.Context, input *cloudwatchlogs.PutResourcePolicyInput) *CloudwatchlogsPutResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input)
    return &CloudwatchlogsPutResourcePolicyResult{Result: future}
}
func (a *CloudWatchLogsStub) PutRetentionPolicy(ctx workflow.Context, input *cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
    var output cloudwatchlogs.PutRetentionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRetentionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutRetentionPolicyAsync(ctx workflow.Context, input *cloudwatchlogs.PutRetentionPolicyInput) *CloudwatchlogsPutRetentionPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutRetentionPolicy, input)
    return &CloudwatchlogsPutRetentionPolicyResult{Result: future}
}
func (a *CloudWatchLogsStub) PutSubscriptionFilter(ctx workflow.Context, input *cloudwatchlogs.PutSubscriptionFilterInput) (*cloudwatchlogs.PutSubscriptionFilterOutput, error) {
    var output cloudwatchlogs.PutSubscriptionFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSubscriptionFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) PutSubscriptionFilterAsync(ctx workflow.Context, input *cloudwatchlogs.PutSubscriptionFilterInput) *CloudwatchlogsPutSubscriptionFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutSubscriptionFilter, input)
    return &CloudwatchlogsPutSubscriptionFilterResult{Result: future}
}
func (a *CloudWatchLogsStub) StartQuery(ctx workflow.Context, input *cloudwatchlogs.StartQueryInput) (*cloudwatchlogs.StartQueryOutput, error) {
    var output cloudwatchlogs.StartQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) StartQueryAsync(ctx workflow.Context, input *cloudwatchlogs.StartQueryInput) *CloudwatchlogsStartQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartQuery, input)
    return &CloudwatchlogsStartQueryResult{Result: future}
}
func (a *CloudWatchLogsStub) StopQuery(ctx workflow.Context, input *cloudwatchlogs.StopQueryInput) (*cloudwatchlogs.StopQueryOutput, error) {
    var output cloudwatchlogs.StopQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) StopQueryAsync(ctx workflow.Context, input *cloudwatchlogs.StopQueryInput) *CloudwatchlogsStopQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopQuery, input)
    return &CloudwatchlogsStopQueryResult{Result: future}
}
func (a *CloudWatchLogsStub) TagLogGroup(ctx workflow.Context, input *cloudwatchlogs.TagLogGroupInput) (*cloudwatchlogs.TagLogGroupOutput, error) {
    var output cloudwatchlogs.TagLogGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagLogGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) TagLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.TagLogGroupInput) *CloudwatchlogsTagLogGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagLogGroup, input)
    return &CloudwatchlogsTagLogGroupResult{Result: future}
}
func (a *CloudWatchLogsStub) TestMetricFilter(ctx workflow.Context, input *cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, error) {
    var output cloudwatchlogs.TestMetricFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestMetricFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) TestMetricFilterAsync(ctx workflow.Context, input *cloudwatchlogs.TestMetricFilterInput) *CloudwatchlogsTestMetricFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestMetricFilter, input)
    return &CloudwatchlogsTestMetricFilterResult{Result: future}
}
func (a *CloudWatchLogsStub) UntagLogGroup(ctx workflow.Context, input *cloudwatchlogs.UntagLogGroupInput) (*cloudwatchlogs.UntagLogGroupOutput, error) {
    var output cloudwatchlogs.UntagLogGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagLogGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudWatchLogsStub) UntagLogGroupAsync(ctx workflow.Context, input *cloudwatchlogs.UntagLogGroupInput) *CloudwatchlogsUntagLogGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagLogGroup, input)
    return &CloudwatchlogsUntagLogGroupResult{Result: future}
}
