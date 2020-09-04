
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type SNSActivities struct {
	client snsiface.SNSAPI
}

func NewSNSActivities(client snsiface.SNSAPI) *SNSActivities {
    return &SNSActivities{client: client}
}

func (a *SNSActivities) AddPermission(input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error) {
    return a.client.AddPermission(input)
}

func (a *SNSActivities) CheckIfPhoneNumberIsOptedOut(input *sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
    return a.client.CheckIfPhoneNumberIsOptedOut(input)
}

func (a *SNSActivities) ConfirmSubscription(input *sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error) {
    return a.client.ConfirmSubscription(input)
}

func (a *SNSActivities) CreatePlatformApplication(input *sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error) {
    return a.client.CreatePlatformApplication(input)
}

func (a *SNSActivities) CreatePlatformEndpoint(input *sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error) {
    return a.client.CreatePlatformEndpoint(input)
}

func (a *SNSActivities) CreateTopic(input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
    return a.client.CreateTopic(input)
}

func (a *SNSActivities) DeleteEndpoint(input *sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error) {
    return a.client.DeleteEndpoint(input)
}

func (a *SNSActivities) DeletePlatformApplication(input *sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error) {
    return a.client.DeletePlatformApplication(input)
}

func (a *SNSActivities) DeleteTopic(input *sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error) {
    return a.client.DeleteTopic(input)
}

func (a *SNSActivities) GetEndpointAttributes(input *sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error) {
    return a.client.GetEndpointAttributes(input)
}

func (a *SNSActivities) GetPlatformApplicationAttributes(input *sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error) {
    return a.client.GetPlatformApplicationAttributes(input)
}

func (a *SNSActivities) GetSMSAttributes(input *sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error) {
    return a.client.GetSMSAttributes(input)
}

func (a *SNSActivities) GetSubscriptionAttributes(input *sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error) {
    return a.client.GetSubscriptionAttributes(input)
}

func (a *SNSActivities) GetTopicAttributes(input *sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error) {
    return a.client.GetTopicAttributes(input)
}

func (a *SNSActivities) ListEndpointsByPlatformApplication(input *sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
    return a.client.ListEndpointsByPlatformApplication(input)
}

func (a *SNSActivities) ListPhoneNumbersOptedOut(input *sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error) {
    return a.client.ListPhoneNumbersOptedOut(input)
}

func (a *SNSActivities) ListPlatformApplications(input *sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error) {
    return a.client.ListPlatformApplications(input)
}

func (a *SNSActivities) ListSubscriptions(input *sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error) {
    return a.client.ListSubscriptions(input)
}

func (a *SNSActivities) ListSubscriptionsByTopic(input *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error) {
    return a.client.ListSubscriptionsByTopic(input)
}

func (a *SNSActivities) ListTagsForResource(input *sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SNSActivities) ListTopics(input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error) {
    return a.client.ListTopics(input)
}

func (a *SNSActivities) OptInPhoneNumber(input *sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error) {
    return a.client.OptInPhoneNumber(input)
}

func (a *SNSActivities) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
    return a.client.Publish(input)
}

func (a *SNSActivities) RemovePermission(input *sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error) {
    return a.client.RemovePermission(input)
}

func (a *SNSActivities) SetEndpointAttributes(input *sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error) {
    return a.client.SetEndpointAttributes(input)
}

func (a *SNSActivities) SetPlatformApplicationAttributes(input *sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error) {
    return a.client.SetPlatformApplicationAttributes(input)
}

func (a *SNSActivities) SetSMSAttributes(input *sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error) {
    return a.client.SetSMSAttributes(input)
}

func (a *SNSActivities) SetSubscriptionAttributes(input *sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error) {
    return a.client.SetSubscriptionAttributes(input)
}

func (a *SNSActivities) SetTopicAttributes(input *sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error) {
    return a.client.SetTopicAttributes(input)
}

func (a *SNSActivities) Subscribe(input *sns.SubscribeInput) (*sns.SubscribeOutput, error) {
    return a.client.Subscribe(input)
}

func (a *SNSActivities) TagResource(input *sns.TagResourceInput) (*sns.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SNSActivities) Unsubscribe(input *sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error) {
    return a.client.Unsubscribe(input)
}

func (a *SNSActivities) UntagResource(input *sns.UntagResourceInput) (*sns.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
