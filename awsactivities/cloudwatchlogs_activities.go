
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CloudWatchLogsActivities struct {
    client cloudwatchlogsiface.CloudWatchLogsAPI
}

func NewCloudWatchLogsActivities(session *session.Session, config ...*aws.Config) *CloudWatchLogsActivities {
    client := cloudwatchlogs.New(session, config...)
    return &CloudWatchLogsActivities{client: client}
}

func (a *CloudWatchLogsActivities) AssociateKmsKey(ctx context.Context, input *cloudwatchlogs.AssociateKmsKeyInput) (*cloudwatchlogs.AssociateKmsKeyOutput, error) {
    return a.client.AssociateKmsKeyWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) CancelExportTask(ctx context.Context, input *cloudwatchlogs.CancelExportTaskInput) (*cloudwatchlogs.CancelExportTaskOutput, error) {
    return a.client.CancelExportTaskWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) CreateExportTask(ctx context.Context, input *cloudwatchlogs.CreateExportTaskInput) (*cloudwatchlogs.CreateExportTaskOutput, error) {
    return a.client.CreateExportTaskWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) CreateLogGroup(ctx context.Context, input *cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error) {
    return a.client.CreateLogGroupWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) CreateLogStream(ctx context.Context, input *cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error) {
    return a.client.CreateLogStreamWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteDestination(ctx context.Context, input *cloudwatchlogs.DeleteDestinationInput) (*cloudwatchlogs.DeleteDestinationOutput, error) {
    return a.client.DeleteDestinationWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteLogGroup(ctx context.Context, input *cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, error) {
    return a.client.DeleteLogGroupWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteLogStream(ctx context.Context, input *cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error) {
    return a.client.DeleteLogStreamWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteMetricFilter(ctx context.Context, input *cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, error) {
    return a.client.DeleteMetricFilterWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteQueryDefinition(ctx context.Context, input *cloudwatchlogs.DeleteQueryDefinitionInput) (*cloudwatchlogs.DeleteQueryDefinitionOutput, error) {
    return a.client.DeleteQueryDefinitionWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteResourcePolicy(ctx context.Context, input *cloudwatchlogs.DeleteResourcePolicyInput) (*cloudwatchlogs.DeleteResourcePolicyOutput, error) {
    return a.client.DeleteResourcePolicyWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteRetentionPolicy(ctx context.Context, input *cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error) {
    return a.client.DeleteRetentionPolicyWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DeleteSubscriptionFilter(ctx context.Context, input *cloudwatchlogs.DeleteSubscriptionFilterInput) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error) {
    return a.client.DeleteSubscriptionFilterWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeDestinations(ctx context.Context, input *cloudwatchlogs.DescribeDestinationsInput) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
    return a.client.DescribeDestinationsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeExportTasks(ctx context.Context, input *cloudwatchlogs.DescribeExportTasksInput) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
    return a.client.DescribeExportTasksWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeLogGroups(ctx context.Context, input *cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
    return a.client.DescribeLogGroupsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeLogStreams(ctx context.Context, input *cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
    return a.client.DescribeLogStreamsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeMetricFilters(ctx context.Context, input *cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
    return a.client.DescribeMetricFiltersWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeQueries(ctx context.Context, input *cloudwatchlogs.DescribeQueriesInput) (*cloudwatchlogs.DescribeQueriesOutput, error) {
    return a.client.DescribeQueriesWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeQueryDefinitions(ctx context.Context, input *cloudwatchlogs.DescribeQueryDefinitionsInput) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {
    return a.client.DescribeQueryDefinitionsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeResourcePolicies(ctx context.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
    return a.client.DescribeResourcePoliciesWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DescribeSubscriptionFilters(ctx context.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
    return a.client.DescribeSubscriptionFiltersWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) DisassociateKmsKey(ctx context.Context, input *cloudwatchlogs.DisassociateKmsKeyInput) (*cloudwatchlogs.DisassociateKmsKeyOutput, error) {
    return a.client.DisassociateKmsKeyWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) FilterLogEvents(ctx context.Context, input *cloudwatchlogs.FilterLogEventsInput) (*cloudwatchlogs.FilterLogEventsOutput, error) {
    return a.client.FilterLogEventsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) GetLogEvents(ctx context.Context, input *cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error) {
    return a.client.GetLogEventsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) GetLogGroupFields(ctx context.Context, input *cloudwatchlogs.GetLogGroupFieldsInput) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
    return a.client.GetLogGroupFieldsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) GetLogRecord(ctx context.Context, input *cloudwatchlogs.GetLogRecordInput) (*cloudwatchlogs.GetLogRecordOutput, error) {
    return a.client.GetLogRecordWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) GetQueryResults(ctx context.Context, input *cloudwatchlogs.GetQueryResultsInput) (*cloudwatchlogs.GetQueryResultsOutput, error) {
    return a.client.GetQueryResultsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) ListTagsLogGroup(ctx context.Context, input *cloudwatchlogs.ListTagsLogGroupInput) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
    return a.client.ListTagsLogGroupWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutDestination(ctx context.Context, input *cloudwatchlogs.PutDestinationInput) (*cloudwatchlogs.PutDestinationOutput, error) {
    return a.client.PutDestinationWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutDestinationPolicy(ctx context.Context, input *cloudwatchlogs.PutDestinationPolicyInput) (*cloudwatchlogs.PutDestinationPolicyOutput, error) {
    return a.client.PutDestinationPolicyWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutLogEvents(ctx context.Context, input *cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error) {
    return a.client.PutLogEventsWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutMetricFilter(ctx context.Context, input *cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, error) {
    return a.client.PutMetricFilterWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutQueryDefinition(ctx context.Context, input *cloudwatchlogs.PutQueryDefinitionInput) (*cloudwatchlogs.PutQueryDefinitionOutput, error) {
    return a.client.PutQueryDefinitionWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutResourcePolicy(ctx context.Context, input *cloudwatchlogs.PutResourcePolicyInput) (*cloudwatchlogs.PutResourcePolicyOutput, error) {
    return a.client.PutResourcePolicyWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutRetentionPolicy(ctx context.Context, input *cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
    return a.client.PutRetentionPolicyWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) PutSubscriptionFilter(ctx context.Context, input *cloudwatchlogs.PutSubscriptionFilterInput) (*cloudwatchlogs.PutSubscriptionFilterOutput, error) {
    return a.client.PutSubscriptionFilterWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) StartQuery(ctx context.Context, input *cloudwatchlogs.StartQueryInput) (*cloudwatchlogs.StartQueryOutput, error) {
    return a.client.StartQueryWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) StopQuery(ctx context.Context, input *cloudwatchlogs.StopQueryInput) (*cloudwatchlogs.StopQueryOutput, error) {
    return a.client.StopQueryWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) TagLogGroup(ctx context.Context, input *cloudwatchlogs.TagLogGroupInput) (*cloudwatchlogs.TagLogGroupOutput, error) {
    return a.client.TagLogGroupWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) TestMetricFilter(ctx context.Context, input *cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, error) {
    return a.client.TestMetricFilterWithContext(ctx, input)
}

func (a *CloudWatchLogsActivities) UntagLogGroup(ctx context.Context, input *cloudwatchlogs.UntagLogGroupInput) (*cloudwatchlogs.UntagLogGroupOutput, error) {
    return a.client.UntagLogGroupWithContext(ctx, input)
}
