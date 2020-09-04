package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

type CloudWatchActivities struct {
	client cloudwatchiface.CloudWatchAPI
}

func NewCloudWatchActivities(client cloudwatchiface.CloudWatchAPI) *CloudWatchActivities {
    return &CloudWatchActivities{client: client}
}


func (a *CloudWatchActivities) DeleteAlarms(input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
    return a.client.DeleteAlarms(input)
}



func (a *CloudWatchActivities) DeleteAnomalyDetector(input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
    return a.client.DeleteAnomalyDetector(input)
}



func (a *CloudWatchActivities) DeleteDashboards(input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
    return a.client.DeleteDashboards(input)
}



func (a *CloudWatchActivities) DeleteInsightRules(input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error) {
    return a.client.DeleteInsightRules(input)
}



func (a *CloudWatchActivities) DescribeAlarmHistory(input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
    return a.client.DescribeAlarmHistory(input)
}



func (a *CloudWatchActivities) DescribeAlarms(input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
    return a.client.DescribeAlarms(input)
}



func (a *CloudWatchActivities) DescribeAlarmsForMetric(input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
    return a.client.DescribeAlarmsForMetric(input)
}



func (a *CloudWatchActivities) DescribeAnomalyDetectors(input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
    return a.client.DescribeAnomalyDetectors(input)
}



func (a *CloudWatchActivities) DescribeInsightRules(input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error) {
    return a.client.DescribeInsightRules(input)
}



func (a *CloudWatchActivities) DisableAlarmActions(input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
    return a.client.DisableAlarmActions(input)
}



func (a *CloudWatchActivities) DisableInsightRules(input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error) {
    return a.client.DisableInsightRules(input)
}



func (a *CloudWatchActivities) EnableAlarmActions(input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
    return a.client.EnableAlarmActions(input)
}



func (a *CloudWatchActivities) EnableInsightRules(input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error) {
    return a.client.EnableInsightRules(input)
}



func (a *CloudWatchActivities) GetDashboard(input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
    return a.client.GetDashboard(input)
}



func (a *CloudWatchActivities) GetInsightRuleReport(input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error) {
    return a.client.GetInsightRuleReport(input)
}



func (a *CloudWatchActivities) GetMetricData(input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
    return a.client.GetMetricData(input)
}



func (a *CloudWatchActivities) GetMetricStatistics(input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
    return a.client.GetMetricStatistics(input)
}



func (a *CloudWatchActivities) GetMetricWidgetImage(input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error) {
    return a.client.GetMetricWidgetImage(input)
}



func (a *CloudWatchActivities) ListDashboards(input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
    return a.client.ListDashboards(input)
}



func (a *CloudWatchActivities) ListMetrics(input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
    return a.client.ListMetrics(input)
}



func (a *CloudWatchActivities) ListTagsForResource(input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *CloudWatchActivities) PutAnomalyDetector(input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error) {
    return a.client.PutAnomalyDetector(input)
}



func (a *CloudWatchActivities) PutCompositeAlarm(input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error) {
    return a.client.PutCompositeAlarm(input)
}



func (a *CloudWatchActivities) PutDashboard(input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
    return a.client.PutDashboard(input)
}



func (a *CloudWatchActivities) PutInsightRule(input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error) {
    return a.client.PutInsightRule(input)
}



func (a *CloudWatchActivities) PutMetricAlarm(input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
    return a.client.PutMetricAlarm(input)
}



func (a *CloudWatchActivities) PutMetricData(input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
    return a.client.PutMetricData(input)
}



func (a *CloudWatchActivities) SetAlarmState(input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
    return a.client.SetAlarmState(input)
}



func (a *CloudWatchActivities) TagResource(input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *CloudWatchActivities) UntagResource(input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *CloudWatchActivities) WaitUntilAlarmExists(input *cloudwatch.DescribeAlarmsInput) error {
    return a.client.WaitUntilAlarmExists(input)
}



func (a *CloudWatchActivities) WaitUntilCompositeAlarmExists(input *cloudwatch.DescribeAlarmsInput) error {
    return a.client.WaitUntilCompositeAlarmExists(input)
}

