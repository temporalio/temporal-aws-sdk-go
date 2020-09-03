package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
)

type CloudTrailActivities struct {
	client cloudtrailiface.CloudTrailAPI
}

func NewCloudTrailActivities(client cloudtrailiface.CloudTrailAPI) *CloudTrailActivities {
    return &CloudTrailActivities{client: client}
}

func (a *CloudTrailActivities) AddTags(input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *CloudTrailActivities) CreateTrail(input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error) {
    return a.client.CreateTrail(input)
}

func (a *CloudTrailActivities) DeleteTrail(input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error) {
    return a.client.DeleteTrail(input)
}

func (a *CloudTrailActivities) DescribeTrails(input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error) {
    return a.client.DescribeTrails(input)
}

func (a *CloudTrailActivities) GetEventSelectors(input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error) {
    return a.client.GetEventSelectors(input)
}

func (a *CloudTrailActivities) GetInsightSelectors(input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error) {
    return a.client.GetInsightSelectors(input)
}

func (a *CloudTrailActivities) GetTrail(input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error) {
    return a.client.GetTrail(input)
}

func (a *CloudTrailActivities) GetTrailStatus(input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error) {
    return a.client.GetTrailStatus(input)
}

func (a *CloudTrailActivities) ListPublicKeys(input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error) {
    return a.client.ListPublicKeys(input)
}

func (a *CloudTrailActivities) ListTags(input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *CloudTrailActivities) ListTrails(input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error) {
    return a.client.ListTrails(input)
}

func (a *CloudTrailActivities) LookupEvents(input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error) {
    return a.client.LookupEvents(input)
}

func (a *CloudTrailActivities) PutEventSelectors(input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error) {
    return a.client.PutEventSelectors(input)
}

func (a *CloudTrailActivities) PutInsightSelectors(input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error) {
    return a.client.PutInsightSelectors(input)
}

func (a *CloudTrailActivities) RemoveTags(input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error) {
    return a.client.RemoveTags(input)
}

func (a *CloudTrailActivities) StartLogging(input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error) {
    return a.client.StartLogging(input)
}

func (a *CloudTrailActivities) StopLogging(input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error) {
    return a.client.StopLogging(input)
}

func (a *CloudTrailActivities) UpdateTrail(input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error) {
    return a.client.UpdateTrail(input)
}
