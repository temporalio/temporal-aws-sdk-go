// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client cloudwatchiface.CloudWatchAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := cloudwatch.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (cloudwatchiface.CloudWatchAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return cloudwatch.New(sess), nil
}

func (a *Activities) DeleteAlarms(ctx context.Context, input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAlarmsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAnomalyDetector(ctx context.Context, input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAnomalyDetectorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDashboards(ctx context.Context, input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDashboardsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteInsightRules(ctx context.Context, input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteInsightRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAlarmHistory(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAlarmHistoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAlarms(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAlarmsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAlarmsForMetric(ctx context.Context, input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAlarmsForMetricWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAnomalyDetectors(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAnomalyDetectorsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInsightRules(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInsightRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisableAlarmActions(ctx context.Context, input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisableAlarmActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisableInsightRules(ctx context.Context, input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisableInsightRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) EnableAlarmActions(ctx context.Context, input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.EnableAlarmActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) EnableInsightRules(ctx context.Context, input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.EnableInsightRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDashboard(ctx context.Context, input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDashboardWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetInsightRuleReport(ctx context.Context, input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetInsightRuleReportWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMetricData(ctx context.Context, input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMetricDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMetricStatistics(ctx context.Context, input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMetricStatisticsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMetricWidgetImage(ctx context.Context, input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMetricWidgetImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDashboards(ctx context.Context, input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDashboardsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMetrics(ctx context.Context, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMetricsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAnomalyDetector(ctx context.Context, input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAnomalyDetectorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutCompositeAlarm(ctx context.Context, input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutCompositeAlarmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutDashboard(ctx context.Context, input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutDashboardWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutInsightRule(ctx context.Context, input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutInsightRuleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutMetricAlarm(ctx context.Context, input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutMetricAlarmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutMetricData(ctx context.Context, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutMetricDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetAlarmState(ctx context.Context, input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetAlarmStateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilAlarmExists(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilAlarmExistsWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilCompositeAlarmExists(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilCompositeAlarmExistsWithContext(ctx, input, options...))
	})
}
