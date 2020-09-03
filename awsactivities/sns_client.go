package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"go.temporal.io/sdk/workflow"
)

type SNSClient interface {
    AddPermission(ctx workflow.Context, input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error)
    AddPermissionAsync(ctx workflow.Context, input *sns.AddPermissionInput) *SnsAddPermissionResult

    CheckIfPhoneNumberIsOptedOut(ctx workflow.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
    CheckIfPhoneNumberIsOptedOutAsync(ctx workflow.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) *SnsCheckIfPhoneNumberIsOptedOutResult

    ConfirmSubscription(ctx workflow.Context, input *sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error)
    ConfirmSubscriptionAsync(ctx workflow.Context, input *sns.ConfirmSubscriptionInput) *SnsConfirmSubscriptionResult

    CreatePlatformApplication(ctx workflow.Context, input *sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error)
    CreatePlatformApplicationAsync(ctx workflow.Context, input *sns.CreatePlatformApplicationInput) *SnsCreatePlatformApplicationResult

    CreatePlatformEndpoint(ctx workflow.Context, input *sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error)
    CreatePlatformEndpointAsync(ctx workflow.Context, input *sns.CreatePlatformEndpointInput) *SnsCreatePlatformEndpointResult

    CreateTopic(ctx workflow.Context, input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error)
    CreateTopicAsync(ctx workflow.Context, input *sns.CreateTopicInput) *SnsCreateTopicResult

    DeleteEndpoint(ctx workflow.Context, input *sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error)
    DeleteEndpointAsync(ctx workflow.Context, input *sns.DeleteEndpointInput) *SnsDeleteEndpointResult

    DeletePlatformApplication(ctx workflow.Context, input *sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error)
    DeletePlatformApplicationAsync(ctx workflow.Context, input *sns.DeletePlatformApplicationInput) *SnsDeletePlatformApplicationResult

    DeleteTopic(ctx workflow.Context, input *sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error)
    DeleteTopicAsync(ctx workflow.Context, input *sns.DeleteTopicInput) *SnsDeleteTopicResult

    GetEndpointAttributes(ctx workflow.Context, input *sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error)
    GetEndpointAttributesAsync(ctx workflow.Context, input *sns.GetEndpointAttributesInput) *SnsGetEndpointAttributesResult

    GetPlatformApplicationAttributes(ctx workflow.Context, input *sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error)
    GetPlatformApplicationAttributesAsync(ctx workflow.Context, input *sns.GetPlatformApplicationAttributesInput) *SnsGetPlatformApplicationAttributesResult

    GetSMSAttributes(ctx workflow.Context, input *sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error)
    GetSMSAttributesAsync(ctx workflow.Context, input *sns.GetSMSAttributesInput) *SnsGetSMSAttributesResult

    GetSubscriptionAttributes(ctx workflow.Context, input *sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error)
    GetSubscriptionAttributesAsync(ctx workflow.Context, input *sns.GetSubscriptionAttributesInput) *SnsGetSubscriptionAttributesResult

    GetTopicAttributes(ctx workflow.Context, input *sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error)
    GetTopicAttributesAsync(ctx workflow.Context, input *sns.GetTopicAttributesInput) *SnsGetTopicAttributesResult

    ListEndpointsByPlatformApplication(ctx workflow.Context, input *sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error)
    ListEndpointsByPlatformApplicationAsync(ctx workflow.Context, input *sns.ListEndpointsByPlatformApplicationInput) *SnsListEndpointsByPlatformApplicationResult

    ListPhoneNumbersOptedOut(ctx workflow.Context, input *sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error)
    ListPhoneNumbersOptedOutAsync(ctx workflow.Context, input *sns.ListPhoneNumbersOptedOutInput) *SnsListPhoneNumbersOptedOutResult

    ListPlatformApplications(ctx workflow.Context, input *sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error)
    ListPlatformApplicationsAsync(ctx workflow.Context, input *sns.ListPlatformApplicationsInput) *SnsListPlatformApplicationsResult

    ListSubscriptions(ctx workflow.Context, input *sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error)
    ListSubscriptionsAsync(ctx workflow.Context, input *sns.ListSubscriptionsInput) *SnsListSubscriptionsResult

    ListSubscriptionsByTopic(ctx workflow.Context, input *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error)
    ListSubscriptionsByTopicAsync(ctx workflow.Context, input *sns.ListSubscriptionsByTopicInput) *SnsListSubscriptionsByTopicResult

    ListTagsForResource(ctx workflow.Context, input *sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *sns.ListTagsForResourceInput) *SnsListTagsForResourceResult

    ListTopics(ctx workflow.Context, input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error)
    ListTopicsAsync(ctx workflow.Context, input *sns.ListTopicsInput) *SnsListTopicsResult

    OptInPhoneNumber(ctx workflow.Context, input *sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error)
    OptInPhoneNumberAsync(ctx workflow.Context, input *sns.OptInPhoneNumberInput) *SnsOptInPhoneNumberResult

    Publish(ctx workflow.Context, input *sns.PublishInput) (*sns.PublishOutput, error)
    PublishAsync(ctx workflow.Context, input *sns.PublishInput) *SnsPublishResult

    RemovePermission(ctx workflow.Context, input *sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error)
    RemovePermissionAsync(ctx workflow.Context, input *sns.RemovePermissionInput) *SnsRemovePermissionResult

    SetEndpointAttributes(ctx workflow.Context, input *sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error)
    SetEndpointAttributesAsync(ctx workflow.Context, input *sns.SetEndpointAttributesInput) *SnsSetEndpointAttributesResult

    SetPlatformApplicationAttributes(ctx workflow.Context, input *sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error)
    SetPlatformApplicationAttributesAsync(ctx workflow.Context, input *sns.SetPlatformApplicationAttributesInput) *SnsSetPlatformApplicationAttributesResult

    SetSMSAttributes(ctx workflow.Context, input *sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error)
    SetSMSAttributesAsync(ctx workflow.Context, input *sns.SetSMSAttributesInput) *SnsSetSMSAttributesResult

    SetSubscriptionAttributes(ctx workflow.Context, input *sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error)
    SetSubscriptionAttributesAsync(ctx workflow.Context, input *sns.SetSubscriptionAttributesInput) *SnsSetSubscriptionAttributesResult

    SetTopicAttributes(ctx workflow.Context, input *sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error)
    SetTopicAttributesAsync(ctx workflow.Context, input *sns.SetTopicAttributesInput) *SnsSetTopicAttributesResult

    Subscribe(ctx workflow.Context, input *sns.SubscribeInput) (*sns.SubscribeOutput, error)
    SubscribeAsync(ctx workflow.Context, input *sns.SubscribeInput) *SnsSubscribeResult

    TagResource(ctx workflow.Context, input *sns.TagResourceInput) (*sns.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *sns.TagResourceInput) *SnsTagResourceResult

    Unsubscribe(ctx workflow.Context, input *sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error)
    UnsubscribeAsync(ctx workflow.Context, input *sns.UnsubscribeInput) *SnsUnsubscribeResult

    UntagResource(ctx workflow.Context, input *sns.UntagResourceInput) (*sns.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *sns.UntagResourceInput) *SnsUntagResourceResult
}
type SnsAddPermissionResult struct {
	Result workflow.Future
}

func (r *SnsAddPermissionResult) Get(ctx workflow.Context) (*sns.AddPermissionOutput, error) {
    var output sns.AddPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsCheckIfPhoneNumberIsOptedOutResult struct {
	Result workflow.Future
}

func (r *SnsCheckIfPhoneNumberIsOptedOutResult) Get(ctx workflow.Context) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
    var output sns.CheckIfPhoneNumberIsOptedOutOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsConfirmSubscriptionResult struct {
	Result workflow.Future
}

func (r *SnsConfirmSubscriptionResult) Get(ctx workflow.Context) (*sns.ConfirmSubscriptionOutput, error) {
    var output sns.ConfirmSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsCreatePlatformApplicationResult struct {
	Result workflow.Future
}

func (r *SnsCreatePlatformApplicationResult) Get(ctx workflow.Context) (*sns.CreatePlatformApplicationOutput, error) {
    var output sns.CreatePlatformApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsCreatePlatformEndpointResult struct {
	Result workflow.Future
}

func (r *SnsCreatePlatformEndpointResult) Get(ctx workflow.Context) (*sns.CreatePlatformEndpointOutput, error) {
    var output sns.CreatePlatformEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsCreateTopicResult struct {
	Result workflow.Future
}

func (r *SnsCreateTopicResult) Get(ctx workflow.Context) (*sns.CreateTopicOutput, error) {
    var output sns.CreateTopicOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsDeleteEndpointResult struct {
	Result workflow.Future
}

func (r *SnsDeleteEndpointResult) Get(ctx workflow.Context) (*sns.DeleteEndpointOutput, error) {
    var output sns.DeleteEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsDeletePlatformApplicationResult struct {
	Result workflow.Future
}

func (r *SnsDeletePlatformApplicationResult) Get(ctx workflow.Context) (*sns.DeletePlatformApplicationOutput, error) {
    var output sns.DeletePlatformApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsDeleteTopicResult struct {
	Result workflow.Future
}

func (r *SnsDeleteTopicResult) Get(ctx workflow.Context) (*sns.DeleteTopicOutput, error) {
    var output sns.DeleteTopicOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsGetEndpointAttributesResult struct {
	Result workflow.Future
}

func (r *SnsGetEndpointAttributesResult) Get(ctx workflow.Context) (*sns.GetEndpointAttributesOutput, error) {
    var output sns.GetEndpointAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsGetPlatformApplicationAttributesResult struct {
	Result workflow.Future
}

func (r *SnsGetPlatformApplicationAttributesResult) Get(ctx workflow.Context) (*sns.GetPlatformApplicationAttributesOutput, error) {
    var output sns.GetPlatformApplicationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsGetSMSAttributesResult struct {
	Result workflow.Future
}

func (r *SnsGetSMSAttributesResult) Get(ctx workflow.Context) (*sns.GetSMSAttributesOutput, error) {
    var output sns.GetSMSAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsGetSubscriptionAttributesResult struct {
	Result workflow.Future
}

func (r *SnsGetSubscriptionAttributesResult) Get(ctx workflow.Context) (*sns.GetSubscriptionAttributesOutput, error) {
    var output sns.GetSubscriptionAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsGetTopicAttributesResult struct {
	Result workflow.Future
}

func (r *SnsGetTopicAttributesResult) Get(ctx workflow.Context) (*sns.GetTopicAttributesOutput, error) {
    var output sns.GetTopicAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsListEndpointsByPlatformApplicationResult struct {
	Result workflow.Future
}

func (r *SnsListEndpointsByPlatformApplicationResult) Get(ctx workflow.Context) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
    var output sns.ListEndpointsByPlatformApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsListPhoneNumbersOptedOutResult struct {
	Result workflow.Future
}

func (r *SnsListPhoneNumbersOptedOutResult) Get(ctx workflow.Context) (*sns.ListPhoneNumbersOptedOutOutput, error) {
    var output sns.ListPhoneNumbersOptedOutOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsListPlatformApplicationsResult struct {
	Result workflow.Future
}

func (r *SnsListPlatformApplicationsResult) Get(ctx workflow.Context) (*sns.ListPlatformApplicationsOutput, error) {
    var output sns.ListPlatformApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsListSubscriptionsResult struct {
	Result workflow.Future
}

func (r *SnsListSubscriptionsResult) Get(ctx workflow.Context) (*sns.ListSubscriptionsOutput, error) {
    var output sns.ListSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsListSubscriptionsByTopicResult struct {
	Result workflow.Future
}

func (r *SnsListSubscriptionsByTopicResult) Get(ctx workflow.Context) (*sns.ListSubscriptionsByTopicOutput, error) {
    var output sns.ListSubscriptionsByTopicOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SnsListTagsForResourceResult) Get(ctx workflow.Context) (*sns.ListTagsForResourceOutput, error) {
    var output sns.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsListTopicsResult struct {
	Result workflow.Future
}

func (r *SnsListTopicsResult) Get(ctx workflow.Context) (*sns.ListTopicsOutput, error) {
    var output sns.ListTopicsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsOptInPhoneNumberResult struct {
	Result workflow.Future
}

func (r *SnsOptInPhoneNumberResult) Get(ctx workflow.Context) (*sns.OptInPhoneNumberOutput, error) {
    var output sns.OptInPhoneNumberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsPublishResult struct {
	Result workflow.Future
}

func (r *SnsPublishResult) Get(ctx workflow.Context) (*sns.PublishOutput, error) {
    var output sns.PublishOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsRemovePermissionResult struct {
	Result workflow.Future
}

func (r *SnsRemovePermissionResult) Get(ctx workflow.Context) (*sns.RemovePermissionOutput, error) {
    var output sns.RemovePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsSetEndpointAttributesResult struct {
	Result workflow.Future
}

func (r *SnsSetEndpointAttributesResult) Get(ctx workflow.Context) (*sns.SetEndpointAttributesOutput, error) {
    var output sns.SetEndpointAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsSetPlatformApplicationAttributesResult struct {
	Result workflow.Future
}

func (r *SnsSetPlatformApplicationAttributesResult) Get(ctx workflow.Context) (*sns.SetPlatformApplicationAttributesOutput, error) {
    var output sns.SetPlatformApplicationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsSetSMSAttributesResult struct {
	Result workflow.Future
}

func (r *SnsSetSMSAttributesResult) Get(ctx workflow.Context) (*sns.SetSMSAttributesOutput, error) {
    var output sns.SetSMSAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsSetSubscriptionAttributesResult struct {
	Result workflow.Future
}

func (r *SnsSetSubscriptionAttributesResult) Get(ctx workflow.Context) (*sns.SetSubscriptionAttributesOutput, error) {
    var output sns.SetSubscriptionAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsSetTopicAttributesResult struct {
	Result workflow.Future
}

func (r *SnsSetTopicAttributesResult) Get(ctx workflow.Context) (*sns.SetTopicAttributesOutput, error) {
    var output sns.SetTopicAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsSubscribeResult struct {
	Result workflow.Future
}

func (r *SnsSubscribeResult) Get(ctx workflow.Context) (*sns.SubscribeOutput, error) {
    var output sns.SubscribeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsTagResourceResult struct {
	Result workflow.Future
}

func (r *SnsTagResourceResult) Get(ctx workflow.Context) (*sns.TagResourceOutput, error) {
    var output sns.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsUnsubscribeResult struct {
	Result workflow.Future
}

func (r *SnsUnsubscribeResult) Get(ctx workflow.Context) (*sns.UnsubscribeOutput, error) {
    var output sns.UnsubscribeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SnsUntagResourceResult struct {
	Result workflow.Future
}

func (r *SnsUntagResourceResult) Get(ctx workflow.Context) (*sns.UntagResourceOutput, error) {
    var output sns.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SNSStub struct {
    activities SNSClient
}

func NewSNSStub() SNSClient {
    return &SNSStub{}
}

func (a *SNSStub) AddPermission(ctx workflow.Context, input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error) {
    var output sns.AddPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) CheckIfPhoneNumberIsOptedOut(ctx workflow.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
    var output sns.CheckIfPhoneNumberIsOptedOutOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CheckIfPhoneNumberIsOptedOut, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ConfirmSubscription(ctx workflow.Context, input *sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error) {
    var output sns.ConfirmSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) CreatePlatformApplication(ctx workflow.Context, input *sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error) {
    var output sns.CreatePlatformApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePlatformApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) CreatePlatformEndpoint(ctx workflow.Context, input *sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error) {
    var output sns.CreatePlatformEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePlatformEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) CreateTopic(ctx workflow.Context, input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
    var output sns.CreateTopicOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTopic, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) DeleteEndpoint(ctx workflow.Context, input *sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error) {
    var output sns.DeleteEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) DeletePlatformApplication(ctx workflow.Context, input *sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error) {
    var output sns.DeletePlatformApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePlatformApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) DeleteTopic(ctx workflow.Context, input *sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error) {
    var output sns.DeleteTopicOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTopic, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) GetEndpointAttributes(ctx workflow.Context, input *sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error) {
    var output sns.GetEndpointAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEndpointAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) GetPlatformApplicationAttributes(ctx workflow.Context, input *sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error) {
    var output sns.GetPlatformApplicationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPlatformApplicationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) GetSMSAttributes(ctx workflow.Context, input *sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error) {
    var output sns.GetSMSAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSMSAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) GetSubscriptionAttributes(ctx workflow.Context, input *sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error) {
    var output sns.GetSubscriptionAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSubscriptionAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) GetTopicAttributes(ctx workflow.Context, input *sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error) {
    var output sns.GetTopicAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTopicAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ListEndpointsByPlatformApplication(ctx workflow.Context, input *sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
    var output sns.ListEndpointsByPlatformApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEndpointsByPlatformApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ListPhoneNumbersOptedOut(ctx workflow.Context, input *sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error) {
    var output sns.ListPhoneNumbersOptedOutOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPhoneNumbersOptedOut, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ListPlatformApplications(ctx workflow.Context, input *sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error) {
    var output sns.ListPlatformApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPlatformApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ListSubscriptions(ctx workflow.Context, input *sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error) {
    var output sns.ListSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSubscriptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ListSubscriptionsByTopic(ctx workflow.Context, input *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error) {
    var output sns.ListSubscriptionsByTopicOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSubscriptionsByTopic, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ListTagsForResource(ctx workflow.Context, input *sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error) {
    var output sns.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) ListTopics(ctx workflow.Context, input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error) {
    var output sns.ListTopicsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTopics, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) OptInPhoneNumber(ctx workflow.Context, input *sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error) {
    var output sns.OptInPhoneNumberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.OptInPhoneNumber, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) Publish(ctx workflow.Context, input *sns.PublishInput) (*sns.PublishOutput, error) {
    var output sns.PublishOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Publish, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) RemovePermission(ctx workflow.Context, input *sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error) {
    var output sns.RemovePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) SetEndpointAttributes(ctx workflow.Context, input *sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error) {
    var output sns.SetEndpointAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetEndpointAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) SetPlatformApplicationAttributes(ctx workflow.Context, input *sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error) {
    var output sns.SetPlatformApplicationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetPlatformApplicationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) SetSMSAttributes(ctx workflow.Context, input *sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error) {
    var output sns.SetSMSAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetSMSAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) SetSubscriptionAttributes(ctx workflow.Context, input *sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error) {
    var output sns.SetSubscriptionAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetSubscriptionAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) SetTopicAttributes(ctx workflow.Context, input *sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error) {
    var output sns.SetTopicAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetTopicAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) Subscribe(ctx workflow.Context, input *sns.SubscribeInput) (*sns.SubscribeOutput, error) {
    var output sns.SubscribeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Subscribe, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) TagResource(ctx workflow.Context, input *sns.TagResourceInput) (*sns.TagResourceOutput, error) {
    var output sns.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) Unsubscribe(ctx workflow.Context, input *sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error) {
    var output sns.UnsubscribeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Unsubscribe, input).Get(ctx, &output)
    return &output, err
}

func (a *SNSStub) UntagResource(ctx workflow.Context, input *sns.UntagResourceInput) (*sns.UntagResourceOutput, error) {
    var output sns.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}
