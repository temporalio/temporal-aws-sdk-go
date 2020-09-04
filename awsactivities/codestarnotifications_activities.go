
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"github.com/aws/aws-sdk-go/service/codestarnotifications/codestarnotificationsiface"
)

type CodeStarNotificationsActivities struct {
	client codestarnotificationsiface.CodeStarNotificationsAPI
}

func NewCodeStarNotificationsActivities(client codestarnotificationsiface.CodeStarNotificationsAPI) *CodeStarNotificationsActivities {
    return &CodeStarNotificationsActivities{client: client}
}

func (a *CodeStarNotificationsActivities) CreateNotificationRule(input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error) {
    return a.client.CreateNotificationRule(input)
}

func (a *CodeStarNotificationsActivities) DeleteNotificationRule(input *codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error) {
    return a.client.DeleteNotificationRule(input)
}

func (a *CodeStarNotificationsActivities) DeleteTarget(input *codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error) {
    return a.client.DeleteTarget(input)
}

func (a *CodeStarNotificationsActivities) DescribeNotificationRule(input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
    return a.client.DescribeNotificationRule(input)
}

func (a *CodeStarNotificationsActivities) ListEventTypes(input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error) {
    return a.client.ListEventTypes(input)
}

func (a *CodeStarNotificationsActivities) ListNotificationRules(input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error) {
    return a.client.ListNotificationRules(input)
}

func (a *CodeStarNotificationsActivities) ListTagsForResource(input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CodeStarNotificationsActivities) ListTargets(input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error) {
    return a.client.ListTargets(input)
}

func (a *CodeStarNotificationsActivities) Subscribe(input *codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error) {
    return a.client.Subscribe(input)
}

func (a *CodeStarNotificationsActivities) TagResource(input *codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CodeStarNotificationsActivities) Unsubscribe(input *codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error) {
    return a.client.Unsubscribe(input)
}

func (a *CodeStarNotificationsActivities) UntagResource(input *codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CodeStarNotificationsActivities) UpdateNotificationRule(input *codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error) {
    return a.client.UpdateNotificationRule(input)
}
