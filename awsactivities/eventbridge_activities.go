
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/eventbridge/eventbridgeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type EventBridgeActivities struct {
    client eventbridgeiface.EventBridgeAPI
}

func NewEventBridgeActivities(session *session.Session, config ...*aws.Config) *EventBridgeActivities {
    client := eventbridge.New(session, config...)
    return &EventBridgeActivities{client: client}
}

func (a *EventBridgeActivities) ActivateEventSource(ctx context.Context, input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error) {
    return a.client.ActivateEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) CreateEventBus(ctx context.Context, input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error) {
    return a.client.CreateEventBusWithContext(ctx, input)
}

func (a *EventBridgeActivities) CreatePartnerEventSource(ctx context.Context, input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error) {
    return a.client.CreatePartnerEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeactivateEventSource(ctx context.Context, input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error) {
    return a.client.DeactivateEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeleteEventBus(ctx context.Context, input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error) {
    return a.client.DeleteEventBusWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeletePartnerEventSource(ctx context.Context, input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error) {
    return a.client.DeletePartnerEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeleteRule(ctx context.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error) {
    return a.client.DeleteRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribeEventBus(ctx context.Context, input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error) {
    return a.client.DescribeEventBusWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribeEventSource(ctx context.Context, input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error) {
    return a.client.DescribeEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribePartnerEventSource(ctx context.Context, input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error) {
    return a.client.DescribePartnerEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribeRule(ctx context.Context, input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error) {
    return a.client.DescribeRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) DisableRule(ctx context.Context, input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error) {
    return a.client.DisableRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) EnableRule(ctx context.Context, input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error) {
    return a.client.EnableRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListEventBuses(ctx context.Context, input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error) {
    return a.client.ListEventBusesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListEventSources(ctx context.Context, input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error) {
    return a.client.ListEventSourcesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListPartnerEventSourceAccounts(ctx context.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
    return a.client.ListPartnerEventSourceAccountsWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListPartnerEventSources(ctx context.Context, input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error) {
    return a.client.ListPartnerEventSourcesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListRuleNamesByTarget(ctx context.Context, input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error) {
    return a.client.ListRuleNamesByTargetWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListRules(ctx context.Context, input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error) {
    return a.client.ListRulesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListTagsForResource(ctx context.Context, input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListTargetsByRule(ctx context.Context, input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error) {
    return a.client.ListTargetsByRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutEvents(ctx context.Context, input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error) {
    return a.client.PutEventsWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutPartnerEvents(ctx context.Context, input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error) {
    return a.client.PutPartnerEventsWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutPermission(ctx context.Context, input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error) {
    return a.client.PutPermissionWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutRule(ctx context.Context, input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error) {
    return a.client.PutRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutTargets(ctx context.Context, input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error) {
    return a.client.PutTargetsWithContext(ctx, input)
}

func (a *EventBridgeActivities) RemovePermission(ctx context.Context, input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error) {
    return a.client.RemovePermissionWithContext(ctx, input)
}

func (a *EventBridgeActivities) RemoveTargets(ctx context.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error) {
    return a.client.RemoveTargetsWithContext(ctx, input)
}

func (a *EventBridgeActivities) TagResource(ctx context.Context, input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) TestEventPattern(ctx context.Context, input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error) {
    return a.client.TestEventPatternWithContext(ctx, input)
}

func (a *EventBridgeActivities) UntagResource(ctx context.Context, input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}
