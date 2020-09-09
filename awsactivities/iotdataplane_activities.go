
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface"
)

type IoTDataPlaneActivities struct {
	client iotdataplaneiface.IoTDataPlaneAPI
}

func NewIoTDataPlaneActivities(session *session.Session, config... *aws.Config) *IoTDataPlaneActivities {
    client := iotdataplane.New(session, config...)
    return &IoTDataPlaneActivities{client: client}
}

func (a *IoTDataPlaneActivities) DeleteThingShadow(input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
    return a.client.DeleteThingShadow(input)
}

func (a *IoTDataPlaneActivities) GetThingShadow(input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
    return a.client.GetThingShadow(input)
}

func (a *IoTDataPlaneActivities) ListNamedShadowsForThing(input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
    return a.client.ListNamedShadowsForThing(input)
}

func (a *IoTDataPlaneActivities) Publish(input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {
    return a.client.Publish(input)
}

func (a *IoTDataPlaneActivities) UpdateThingShadow(input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error) {
    return a.client.UpdateThingShadow(input)
}
