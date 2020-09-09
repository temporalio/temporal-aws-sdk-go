
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/eventbridge/eventbridgeiface"
)

type EventBridgeActivities struct {
    client eventbridgeiface.EventBridgeAPI
}

func NewEventBridgeActivities(session *session.Session, config ...*aws.Config) *EventBridgeActivities {
    client := eventbridge.New(session, config...)
    return &EventBridgeActivities{client: client}
}

func (a *EventBridgeActivities) ActivateEventSource(input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error) {
    return a.client.ActivateEventSource(input)
}

func (a *EventBridgeActivities) CreateEventBus(input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error) {
    return a.client.CreateEventBus(input)
}

func (a *EventBridgeActivities) CreatePartnerEventSource(input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error) {
    return a.client.CreatePartnerEventSource(input)
}

func (a *EventBridgeActivities) DeactivateEventSource(input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error) {
    return a.client.DeactivateEventSource(input)
}

func (a *EventBridgeActivities) DeleteEventBus(input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error) {
    return a.client.DeleteEventBus(input)
}

func (a *EventBridgeActivities) DeletePartnerEventSource(input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error) {
    return a.client.DeletePartnerEventSource(input)
}

func (a *EventBridgeActivities) DeleteRule(input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error) {
    return a.client.DeleteRule(input)
}

func (a *EventBridgeActivities) DescribeEventBus(input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error) {
    return a.client.DescribeEventBus(input)
}

func (a *EventBridgeActivities) DescribeEventSource(input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error) {
    return a.client.DescribeEventSource(input)
}

func (a *EventBridgeActivities) DescribePartnerEventSource(input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error) {
    return a.client.DescribePartnerEventSource(input)
}

func (a *EventBridgeActivities) DescribeRule(input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error) {
    return a.client.DescribeRule(input)
}

func (a *EventBridgeActivities) DisableRule(input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error) {
    return a.client.DisableRule(input)
}

func (a *EventBridgeActivities) EnableRule(input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error) {
    return a.client.EnableRule(input)
}

func (a *EventBridgeActivities) ListEventBuses(input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error) {
    return a.client.ListEventBuses(input)
}

func (a *EventBridgeActivities) ListEventSources(input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error) {
    return a.client.ListEventSources(input)
}

func (a *EventBridgeActivities) ListPartnerEventSourceAccounts(input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
    return a.client.ListPartnerEventSourceAccounts(input)
}

func (a *EventBridgeActivities) ListPartnerEventSources(input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error) {
    return a.client.ListPartnerEventSources(input)
}

func (a *EventBridgeActivities) ListRuleNamesByTarget(input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error) {
    return a.client.ListRuleNamesByTarget(input)
}

func (a *EventBridgeActivities) ListRules(input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error) {
    return a.client.ListRules(input)
}

func (a *EventBridgeActivities) ListTagsForResource(input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *EventBridgeActivities) ListTargetsByRule(input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error) {
    return a.client.ListTargetsByRule(input)
}

func (a *EventBridgeActivities) PutEvents(input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error) {
    return a.client.PutEvents(input)
}

func (a *EventBridgeActivities) PutPartnerEvents(input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error) {
    return a.client.PutPartnerEvents(input)
}

func (a *EventBridgeActivities) PutPermission(input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error) {
    return a.client.PutPermission(input)
}

func (a *EventBridgeActivities) PutRule(input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error) {
    return a.client.PutRule(input)
}

func (a *EventBridgeActivities) PutTargets(input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error) {
    return a.client.PutTargets(input)
}

func (a *EventBridgeActivities) RemovePermission(input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error) {
    return a.client.RemovePermission(input)
}

func (a *EventBridgeActivities) RemoveTargets(input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error) {
    return a.client.RemoveTargets(input)
}

func (a *EventBridgeActivities) TagResource(input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *EventBridgeActivities) TestEventPattern(input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error) {
    return a.client.TestEventPattern(input)
}

func (a *EventBridgeActivities) UntagResource(input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
