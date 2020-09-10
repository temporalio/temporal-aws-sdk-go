package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/iotevents/ioteventsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type IoTEventsActivities struct {
	client ioteventsiface.IoTEventsAPI
}

func NewIoTEventsActivities(session *session.Session, config ...*aws.Config) *IoTEventsActivities {
	client := iotevents.New(session, config...)
	return &IoTEventsActivities{client: client}
}

func (a *IoTEventsActivities) CreateDetectorModel(ctx context.Context, input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error) {
	return a.client.CreateDetectorModelWithContext(ctx, input)
}

func (a *IoTEventsActivities) CreateInput(ctx context.Context, input *iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error) {
	return a.client.CreateInputWithContext(ctx, input)
}

func (a *IoTEventsActivities) DeleteDetectorModel(ctx context.Context, input *iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error) {
	return a.client.DeleteDetectorModelWithContext(ctx, input)
}

func (a *IoTEventsActivities) DeleteInput(ctx context.Context, input *iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error) {
	return a.client.DeleteInputWithContext(ctx, input)
}

func (a *IoTEventsActivities) DescribeDetectorModel(ctx context.Context, input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error) {
	return a.client.DescribeDetectorModelWithContext(ctx, input)
}

func (a *IoTEventsActivities) DescribeInput(ctx context.Context, input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error) {
	return a.client.DescribeInputWithContext(ctx, input)
}

func (a *IoTEventsActivities) DescribeLoggingOptions(ctx context.Context, input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error) {
	return a.client.DescribeLoggingOptionsWithContext(ctx, input)
}

func (a *IoTEventsActivities) ListDetectorModelVersions(ctx context.Context, input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error) {
	return a.client.ListDetectorModelVersionsWithContext(ctx, input)
}

func (a *IoTEventsActivities) ListDetectorModels(ctx context.Context, input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error) {
	return a.client.ListDetectorModelsWithContext(ctx, input)
}

func (a *IoTEventsActivities) ListInputs(ctx context.Context, input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error) {
	return a.client.ListInputsWithContext(ctx, input)
}

func (a *IoTEventsActivities) ListTagsForResource(ctx context.Context, input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTEventsActivities) PutLoggingOptions(ctx context.Context, input *iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error) {
	return a.client.PutLoggingOptionsWithContext(ctx, input)
}

func (a *IoTEventsActivities) TagResource(ctx context.Context, input *iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoTEventsActivities) UntagResource(ctx context.Context, input *iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IoTEventsActivities) UpdateDetectorModel(ctx context.Context, input *iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error) {
	return a.client.UpdateDetectorModelWithContext(ctx, input)
}

func (a *IoTEventsActivities) UpdateInput(ctx context.Context, input *iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error) {
	return a.client.UpdateInputWithContext(ctx, input)
}
