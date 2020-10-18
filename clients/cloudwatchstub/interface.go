// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudwatchstub

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteAlarms(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAlarmsAsync(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) *CloudWatchDeleteAlarmsFuture

	DeleteAnomalyDetector(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error)
	DeleteAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) *CloudWatchDeleteAnomalyDetectorFuture

	DeleteDashboards(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteDashboardsAsync(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) *CloudWatchDeleteDashboardsFuture

	DeleteInsightRules(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error)
	DeleteInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) *CloudWatchDeleteInsightRulesFuture

	DescribeAlarmHistory(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) *CloudWatchDescribeAlarmHistoryFuture

	DescribeAlarms(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *CloudWatchDescribeAlarmsFuture

	DescribeAlarmsForMetric(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsForMetricAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) *CloudWatchDescribeAlarmsForMetricFuture

	DescribeAnomalyDetectors(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error)
	DescribeAnomalyDetectorsAsync(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) *CloudWatchDescribeAnomalyDetectorsFuture

	DescribeInsightRules(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error)
	DescribeInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) *CloudWatchDescribeInsightRulesFuture

	DisableAlarmActions(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) *CloudWatchDisableAlarmActionsFuture

	DisableInsightRules(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error)
	DisableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) *CloudWatchDisableInsightRulesFuture

	EnableAlarmActions(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) *CloudWatchEnableAlarmActionsFuture

	EnableInsightRules(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error)
	EnableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) *CloudWatchEnableInsightRulesFuture

	GetDashboard(ctx workflow.Context, input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error)
	GetDashboardAsync(ctx workflow.Context, input *cloudwatch.GetDashboardInput) *CloudWatchGetDashboardFuture

	GetInsightRuleReport(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error)
	GetInsightRuleReportAsync(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) *CloudWatchGetInsightRuleReportFuture

	GetMetricData(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricDataAsync(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) *CloudWatchGetMetricDataFuture

	GetMetricStatistics(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStatisticsAsync(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) *CloudWatchGetMetricStatisticsFuture

	GetMetricWidgetImage(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error)
	GetMetricWidgetImageAsync(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) *CloudWatchGetMetricWidgetImageFuture

	ListDashboards(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsAsync(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) *CloudWatchListDashboardsFuture

	ListMetrics(ctx workflow.Context, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsAsync(ctx workflow.Context, input *cloudwatch.ListMetricsInput) *CloudWatchListMetricsFuture

	ListTagsForResource(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) *CloudWatchListTagsForResourceFuture

	PutAnomalyDetector(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error)
	PutAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) *CloudWatchPutAnomalyDetectorFuture

	PutCompositeAlarm(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error)
	PutCompositeAlarmAsync(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) *CloudWatchPutCompositeAlarmFuture

	PutDashboard(ctx workflow.Context, input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error)
	PutDashboardAsync(ctx workflow.Context, input *cloudwatch.PutDashboardInput) *CloudWatchPutDashboardFuture

	PutInsightRule(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error)
	PutInsightRuleAsync(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) *CloudWatchPutInsightRuleFuture

	PutMetricAlarm(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricAlarmAsync(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) *CloudWatchPutMetricAlarmFuture

	PutMetricData(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricDataAsync(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) *CloudWatchPutMetricDataFuture

	SetAlarmState(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error)
	SetAlarmStateAsync(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) *CloudWatchSetAlarmStateFuture

	TagResource(ctx workflow.Context, input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloudwatch.TagResourceInput) *CloudWatchTagResourceFuture

	UntagResource(ctx workflow.Context, input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloudwatch.UntagResourceInput) *CloudWatchUntagResourceFuture

	WaitUntilAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error
	WaitUntilAlarmExistsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *clients.VoidFuture

	WaitUntilCompositeAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error
	WaitUntilCompositeAlarmExistsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
