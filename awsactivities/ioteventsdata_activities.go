package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"github.com/aws/aws-sdk-go/service/ioteventsdata/ioteventsdataiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type IoTEventsDataActivities struct {
	client ioteventsdataiface.IoTEventsDataAPI
}

func NewIoTEventsDataActivities(session *session.Session, config ...*aws.Config) *IoTEventsDataActivities {
	client := ioteventsdata.New(session, config...)
	return &IoTEventsDataActivities{client: client}
}

func (a *IoTEventsDataActivities) BatchPutMessage(ctx context.Context, input *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error) {
	return a.client.BatchPutMessageWithContext(ctx, input)
}

func (a *IoTEventsDataActivities) BatchUpdateDetector(ctx context.Context, input *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	return a.client.BatchUpdateDetectorWithContext(ctx, input)
}

func (a *IoTEventsDataActivities) DescribeDetector(ctx context.Context, input *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error) {
	return a.client.DescribeDetectorWithContext(ctx, input)
}

func (a *IoTEventsDataActivities) ListDetectors(ctx context.Context, input *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error) {
	return a.client.ListDetectorsWithContext(ctx, input)
}
