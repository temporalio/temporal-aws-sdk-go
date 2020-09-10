package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CloudWatchActivities struct {
	client cloudwatchiface.CloudWatchAPI
}

func NewCloudWatchActivities(session *session.Session, config ...*aws.Config) *CloudWatchActivities {
	client := cloudwatch.New(session, config...)
	return &CloudWatchActivities{client: client}
}

func (a *CloudWatchActivities) DeleteAlarms(ctx context.Context, input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	return a.client.DeleteAlarmsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DeleteAnomalyDetector(ctx context.Context, input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	return a.client.DeleteAnomalyDetectorWithContext(ctx, input)
}

func (a *CloudWatchActivities) DeleteDashboards(ctx context.Context, input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	return a.client.DeleteDashboardsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DeleteInsightRules(ctx context.Context, input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error) {
	return a.client.DeleteInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAlarmHistory(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	return a.client.DescribeAlarmHistoryWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAlarms(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	return a.client.DescribeAlarmsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAlarmsForMetric(ctx context.Context, input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	return a.client.DescribeAlarmsForMetricWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAnomalyDetectors(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	return a.client.DescribeAnomalyDetectorsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeInsightRules(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error) {
	return a.client.DescribeInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) DisableAlarmActions(ctx context.Context, input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	return a.client.DisableAlarmActionsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DisableInsightRules(ctx context.Context, input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error) {
	return a.client.DisableInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) EnableAlarmActions(ctx context.Context, input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	return a.client.EnableAlarmActionsWithContext(ctx, input)
}

func (a *CloudWatchActivities) EnableInsightRules(ctx context.Context, input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error) {
	return a.client.EnableInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetDashboard(ctx context.Context, input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	return a.client.GetDashboardWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetInsightRuleReport(ctx context.Context, input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error) {
	return a.client.GetInsightRuleReportWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetMetricData(ctx context.Context, input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	return a.client.GetMetricDataWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetMetricStatistics(ctx context.Context, input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	return a.client.GetMetricStatisticsWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetMetricWidgetImage(ctx context.Context, input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	return a.client.GetMetricWidgetImageWithContext(ctx, input)
}

func (a *CloudWatchActivities) ListDashboards(ctx context.Context, input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	return a.client.ListDashboardsWithContext(ctx, input)
}

func (a *CloudWatchActivities) ListMetrics(ctx context.Context, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	return a.client.ListMetricsWithContext(ctx, input)
}

func (a *CloudWatchActivities) ListTagsForResource(ctx context.Context, input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutAnomalyDetector(ctx context.Context, input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	return a.client.PutAnomalyDetectorWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutCompositeAlarm(ctx context.Context, input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error) {
	return a.client.PutCompositeAlarmWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutDashboard(ctx context.Context, input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	return a.client.PutDashboardWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutInsightRule(ctx context.Context, input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error) {
	return a.client.PutInsightRuleWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutMetricAlarm(ctx context.Context, input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	return a.client.PutMetricAlarmWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutMetricData(ctx context.Context, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	return a.client.PutMetricDataWithContext(ctx, input)
}

func (a *CloudWatchActivities) SetAlarmState(ctx context.Context, input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	return a.client.SetAlarmStateWithContext(ctx, input)
}

func (a *CloudWatchActivities) TagResource(ctx context.Context, input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CloudWatchActivities) UntagResource(ctx context.Context, input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CloudWatchActivities) WaitUntilAlarmExists(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) error {
	return a.client.WaitUntilAlarmExistsWithContext(ctx, input)

}

func (a *CloudWatchActivities) WaitUntilCompositeAlarmExists(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) error {
	return a.client.WaitUntilCompositeAlarmExistsWithContext(ctx, input)

}
