
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
)

type CloudWatchLogsActivities struct {
	client cloudwatchlogsiface.CloudWatchLogsAPI
}

func NewCloudWatchLogsActivities(session *session.Session, config... *aws.Config) *CloudWatchLogsActivities {
    client := cloudwatchlogs.New(session, config...)
    return &CloudWatchLogsActivities{client: client}
}

func (a *CloudWatchLogsActivities) AssociateKmsKey(input *cloudwatchlogs.AssociateKmsKeyInput) (*cloudwatchlogs.AssociateKmsKeyOutput, error) {
    return a.client.AssociateKmsKey(input)
}

func (a *CloudWatchLogsActivities) CancelExportTask(input *cloudwatchlogs.CancelExportTaskInput) (*cloudwatchlogs.CancelExportTaskOutput, error) {
    return a.client.CancelExportTask(input)
}

func (a *CloudWatchLogsActivities) CreateExportTask(input *cloudwatchlogs.CreateExportTaskInput) (*cloudwatchlogs.CreateExportTaskOutput, error) {
    return a.client.CreateExportTask(input)
}

func (a *CloudWatchLogsActivities) CreateLogGroup(input *cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error) {
    return a.client.CreateLogGroup(input)
}

func (a *CloudWatchLogsActivities) CreateLogStream(input *cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error) {
    return a.client.CreateLogStream(input)
}

func (a *CloudWatchLogsActivities) DeleteDestination(input *cloudwatchlogs.DeleteDestinationInput) (*cloudwatchlogs.DeleteDestinationOutput, error) {
    return a.client.DeleteDestination(input)
}

func (a *CloudWatchLogsActivities) DeleteLogGroup(input *cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, error) {
    return a.client.DeleteLogGroup(input)
}

func (a *CloudWatchLogsActivities) DeleteLogStream(input *cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error) {
    return a.client.DeleteLogStream(input)
}

func (a *CloudWatchLogsActivities) DeleteMetricFilter(input *cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, error) {
    return a.client.DeleteMetricFilter(input)
}

func (a *CloudWatchLogsActivities) DeleteQueryDefinition(input *cloudwatchlogs.DeleteQueryDefinitionInput) (*cloudwatchlogs.DeleteQueryDefinitionOutput, error) {
    return a.client.DeleteQueryDefinition(input)
}

func (a *CloudWatchLogsActivities) DeleteResourcePolicy(input *cloudwatchlogs.DeleteResourcePolicyInput) (*cloudwatchlogs.DeleteResourcePolicyOutput, error) {
    return a.client.DeleteResourcePolicy(input)
}

func (a *CloudWatchLogsActivities) DeleteRetentionPolicy(input *cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error) {
    return a.client.DeleteRetentionPolicy(input)
}

func (a *CloudWatchLogsActivities) DeleteSubscriptionFilter(input *cloudwatchlogs.DeleteSubscriptionFilterInput) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error) {
    return a.client.DeleteSubscriptionFilter(input)
}

func (a *CloudWatchLogsActivities) DescribeDestinations(input *cloudwatchlogs.DescribeDestinationsInput) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
    return a.client.DescribeDestinations(input)
}

func (a *CloudWatchLogsActivities) DescribeExportTasks(input *cloudwatchlogs.DescribeExportTasksInput) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
    return a.client.DescribeExportTasks(input)
}

func (a *CloudWatchLogsActivities) DescribeLogGroups(input *cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
    return a.client.DescribeLogGroups(input)
}

func (a *CloudWatchLogsActivities) DescribeLogStreams(input *cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
    return a.client.DescribeLogStreams(input)
}

func (a *CloudWatchLogsActivities) DescribeMetricFilters(input *cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
    return a.client.DescribeMetricFilters(input)
}

func (a *CloudWatchLogsActivities) DescribeQueries(input *cloudwatchlogs.DescribeQueriesInput) (*cloudwatchlogs.DescribeQueriesOutput, error) {
    return a.client.DescribeQueries(input)
}

func (a *CloudWatchLogsActivities) DescribeQueryDefinitions(input *cloudwatchlogs.DescribeQueryDefinitionsInput) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {
    return a.client.DescribeQueryDefinitions(input)
}

func (a *CloudWatchLogsActivities) DescribeResourcePolicies(input *cloudwatchlogs.DescribeResourcePoliciesInput) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
    return a.client.DescribeResourcePolicies(input)
}

func (a *CloudWatchLogsActivities) DescribeSubscriptionFilters(input *cloudwatchlogs.DescribeSubscriptionFiltersInput) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
    return a.client.DescribeSubscriptionFilters(input)
}

func (a *CloudWatchLogsActivities) DisassociateKmsKey(input *cloudwatchlogs.DisassociateKmsKeyInput) (*cloudwatchlogs.DisassociateKmsKeyOutput, error) {
    return a.client.DisassociateKmsKey(input)
}

func (a *CloudWatchLogsActivities) FilterLogEvents(input *cloudwatchlogs.FilterLogEventsInput) (*cloudwatchlogs.FilterLogEventsOutput, error) {
    return a.client.FilterLogEvents(input)
}

func (a *CloudWatchLogsActivities) GetLogEvents(input *cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error) {
    return a.client.GetLogEvents(input)
}

func (a *CloudWatchLogsActivities) GetLogGroupFields(input *cloudwatchlogs.GetLogGroupFieldsInput) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
    return a.client.GetLogGroupFields(input)
}

func (a *CloudWatchLogsActivities) GetLogRecord(input *cloudwatchlogs.GetLogRecordInput) (*cloudwatchlogs.GetLogRecordOutput, error) {
    return a.client.GetLogRecord(input)
}

func (a *CloudWatchLogsActivities) GetQueryResults(input *cloudwatchlogs.GetQueryResultsInput) (*cloudwatchlogs.GetQueryResultsOutput, error) {
    return a.client.GetQueryResults(input)
}

func (a *CloudWatchLogsActivities) ListTagsLogGroup(input *cloudwatchlogs.ListTagsLogGroupInput) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
    return a.client.ListTagsLogGroup(input)
}

func (a *CloudWatchLogsActivities) PutDestination(input *cloudwatchlogs.PutDestinationInput) (*cloudwatchlogs.PutDestinationOutput, error) {
    return a.client.PutDestination(input)
}

func (a *CloudWatchLogsActivities) PutDestinationPolicy(input *cloudwatchlogs.PutDestinationPolicyInput) (*cloudwatchlogs.PutDestinationPolicyOutput, error) {
    return a.client.PutDestinationPolicy(input)
}

func (a *CloudWatchLogsActivities) PutLogEvents(input *cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error) {
    return a.client.PutLogEvents(input)
}

func (a *CloudWatchLogsActivities) PutMetricFilter(input *cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, error) {
    return a.client.PutMetricFilter(input)
}

func (a *CloudWatchLogsActivities) PutQueryDefinition(input *cloudwatchlogs.PutQueryDefinitionInput) (*cloudwatchlogs.PutQueryDefinitionOutput, error) {
    return a.client.PutQueryDefinition(input)
}

func (a *CloudWatchLogsActivities) PutResourcePolicy(input *cloudwatchlogs.PutResourcePolicyInput) (*cloudwatchlogs.PutResourcePolicyOutput, error) {
    return a.client.PutResourcePolicy(input)
}

func (a *CloudWatchLogsActivities) PutRetentionPolicy(input *cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
    return a.client.PutRetentionPolicy(input)
}

func (a *CloudWatchLogsActivities) PutSubscriptionFilter(input *cloudwatchlogs.PutSubscriptionFilterInput) (*cloudwatchlogs.PutSubscriptionFilterOutput, error) {
    return a.client.PutSubscriptionFilter(input)
}

func (a *CloudWatchLogsActivities) StartQuery(input *cloudwatchlogs.StartQueryInput) (*cloudwatchlogs.StartQueryOutput, error) {
    return a.client.StartQuery(input)
}

func (a *CloudWatchLogsActivities) StopQuery(input *cloudwatchlogs.StopQueryInput) (*cloudwatchlogs.StopQueryOutput, error) {
    return a.client.StopQuery(input)
}

func (a *CloudWatchLogsActivities) TagLogGroup(input *cloudwatchlogs.TagLogGroupInput) (*cloudwatchlogs.TagLogGroupOutput, error) {
    return a.client.TagLogGroup(input)
}

func (a *CloudWatchLogsActivities) TestMetricFilter(input *cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, error) {
    return a.client.TestMetricFilter(input)
}

func (a *CloudWatchLogsActivities) UntagLogGroup(input *cloudwatchlogs.UntagLogGroupInput) (*cloudwatchlogs.UntagLogGroupOutput, error) {
    return a.client.UntagLogGroup(input)
}
