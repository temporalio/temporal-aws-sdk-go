package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
)

type CloudWatchEventsActivities struct {
	client cloudwatcheventsiface.CloudWatchEventsAPI
}

func NewCloudWatchEventsActivities(session *session.Session, config ...*aws.Config) *CloudWatchEventsActivities {
	client := cloudwatchevents.New(session, config...)
	return &CloudWatchEventsActivities{client: client}
}

func (a *CloudWatchEventsActivities) ActivateEventSource(input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error) {
	return a.client.ActivateEventSource(input)
}

func (a *CloudWatchEventsActivities) CreateEventBus(input *cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error) {
	return a.client.CreateEventBus(input)
}

func (a *CloudWatchEventsActivities) CreatePartnerEventSource(input *cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error) {
	return a.client.CreatePartnerEventSource(input)
}

func (a *CloudWatchEventsActivities) DeactivateEventSource(input *cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error) {
	return a.client.DeactivateEventSource(input)
}

func (a *CloudWatchEventsActivities) DeleteEventBus(input *cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error) {
	return a.client.DeleteEventBus(input)
}

func (a *CloudWatchEventsActivities) DeletePartnerEventSource(input *cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error) {
	return a.client.DeletePartnerEventSource(input)
}

func (a *CloudWatchEventsActivities) DeleteRule(input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error) {
	return a.client.DeleteRule(input)
}

func (a *CloudWatchEventsActivities) DescribeEventBus(input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error) {
	return a.client.DescribeEventBus(input)
}

func (a *CloudWatchEventsActivities) DescribeEventSource(input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	return a.client.DescribeEventSource(input)
}

func (a *CloudWatchEventsActivities) DescribePartnerEventSource(input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	return a.client.DescribePartnerEventSource(input)
}

func (a *CloudWatchEventsActivities) DescribeRule(input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error) {
	return a.client.DescribeRule(input)
}

func (a *CloudWatchEventsActivities) DisableRule(input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error) {
	return a.client.DisableRule(input)
}

func (a *CloudWatchEventsActivities) EnableRule(input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error) {
	return a.client.EnableRule(input)
}

func (a *CloudWatchEventsActivities) ListEventBuses(input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error) {
	return a.client.ListEventBuses(input)
}

func (a *CloudWatchEventsActivities) ListEventSources(input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error) {
	return a.client.ListEventSources(input)
}

func (a *CloudWatchEventsActivities) ListPartnerEventSourceAccounts(input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	return a.client.ListPartnerEventSourceAccounts(input)
}

func (a *CloudWatchEventsActivities) ListPartnerEventSources(input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	return a.client.ListPartnerEventSources(input)
}

func (a *CloudWatchEventsActivities) ListRuleNamesByTarget(input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	return a.client.ListRuleNamesByTarget(input)
}

func (a *CloudWatchEventsActivities) ListRules(input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error) {
	return a.client.ListRules(input)
}

func (a *CloudWatchEventsActivities) ListTagsForResource(input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *CloudWatchEventsActivities) ListTargetsByRule(input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	return a.client.ListTargetsByRule(input)
}

func (a *CloudWatchEventsActivities) PutEvents(input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error) {
	return a.client.PutEvents(input)
}

func (a *CloudWatchEventsActivities) PutPartnerEvents(input *cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error) {
	return a.client.PutPartnerEvents(input)
}

func (a *CloudWatchEventsActivities) PutPermission(input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error) {
	return a.client.PutPermission(input)
}

func (a *CloudWatchEventsActivities) PutRule(input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error) {
	return a.client.PutRule(input)
}

func (a *CloudWatchEventsActivities) PutTargets(input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error) {
	return a.client.PutTargets(input)
}

func (a *CloudWatchEventsActivities) RemovePermission(input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error) {
	return a.client.RemovePermission(input)
}

func (a *CloudWatchEventsActivities) RemoveTargets(input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error) {
	return a.client.RemoveTargets(input)
}

func (a *CloudWatchEventsActivities) TagResource(input *cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *CloudWatchEventsActivities) TestEventPattern(input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error) {
	return a.client.TestEventPattern(input)
}

func (a *CloudWatchEventsActivities) UntagResource(input *cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}
