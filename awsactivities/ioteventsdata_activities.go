
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"github.com/aws/aws-sdk-go/service/ioteventsdata/ioteventsdataiface"
)

type IoTEventsDataActivities struct {
	client ioteventsdataiface.IoTEventsDataAPI
}

func NewIoTEventsDataActivities(client ioteventsdataiface.IoTEventsDataAPI) *IoTEventsDataActivities {
    return &IoTEventsDataActivities{client: client}
}

func (a *IoTEventsDataActivities) BatchPutMessage(input *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error) {
    return a.client.BatchPutMessage(input)
}

func (a *IoTEventsDataActivities) BatchUpdateDetector(input *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
    return a.client.BatchUpdateDetector(input)
}

func (a *IoTEventsDataActivities) DescribeDetector(input *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error) {
    return a.client.DescribeDetector(input)
}

func (a *IoTEventsDataActivities) ListDetectors(input *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error) {
    return a.client.ListDetectors(input)
}
