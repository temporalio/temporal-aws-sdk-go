package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CloudTrailActivities struct {
	client cloudtrailiface.CloudTrailAPI
}

func NewCloudTrailActivities(session *session.Session, config ...*aws.Config) *CloudTrailActivities {
	client := cloudtrail.New(session, config...)
	return &CloudTrailActivities{client: client}
}

func (a *CloudTrailActivities) AddTags(ctx context.Context, input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
	return a.client.AddTagsWithContext(ctx, input)
}

func (a *CloudTrailActivities) CreateTrail(ctx context.Context, input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error) {
	return a.client.CreateTrailWithContext(ctx, input)
}

func (a *CloudTrailActivities) DeleteTrail(ctx context.Context, input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error) {
	return a.client.DeleteTrailWithContext(ctx, input)
}

func (a *CloudTrailActivities) DescribeTrails(ctx context.Context, input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error) {
	return a.client.DescribeTrailsWithContext(ctx, input)
}

func (a *CloudTrailActivities) GetEventSelectors(ctx context.Context, input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error) {
	return a.client.GetEventSelectorsWithContext(ctx, input)
}

func (a *CloudTrailActivities) GetInsightSelectors(ctx context.Context, input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error) {
	return a.client.GetInsightSelectorsWithContext(ctx, input)
}

func (a *CloudTrailActivities) GetTrail(ctx context.Context, input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error) {
	return a.client.GetTrailWithContext(ctx, input)
}

func (a *CloudTrailActivities) GetTrailStatus(ctx context.Context, input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error) {
	return a.client.GetTrailStatusWithContext(ctx, input)
}

func (a *CloudTrailActivities) ListPublicKeys(ctx context.Context, input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error) {
	return a.client.ListPublicKeysWithContext(ctx, input)
}

func (a *CloudTrailActivities) ListTags(ctx context.Context, input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *CloudTrailActivities) ListTrails(ctx context.Context, input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error) {
	return a.client.ListTrailsWithContext(ctx, input)
}

func (a *CloudTrailActivities) LookupEvents(ctx context.Context, input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error) {
	return a.client.LookupEventsWithContext(ctx, input)
}

func (a *CloudTrailActivities) PutEventSelectors(ctx context.Context, input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error) {
	return a.client.PutEventSelectorsWithContext(ctx, input)
}

func (a *CloudTrailActivities) PutInsightSelectors(ctx context.Context, input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error) {
	return a.client.PutInsightSelectorsWithContext(ctx, input)
}

func (a *CloudTrailActivities) RemoveTags(ctx context.Context, input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error) {
	return a.client.RemoveTagsWithContext(ctx, input)
}

func (a *CloudTrailActivities) StartLogging(ctx context.Context, input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error) {
	return a.client.StartLoggingWithContext(ctx, input)
}

func (a *CloudTrailActivities) StopLogging(ctx context.Context, input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error) {
	return a.client.StopLoggingWithContext(ctx, input)
}

func (a *CloudTrailActivities) UpdateTrail(ctx context.Context, input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error) {
	return a.client.UpdateTrailWithContext(ctx, input)
}
