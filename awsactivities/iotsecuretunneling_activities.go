package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling/iotsecuretunnelingiface"
)

type IoTSecureTunnelingActivities struct {
	client iotsecuretunnelingiface.IoTSecureTunnelingAPI
}

func NewIoTSecureTunnelingActivities(session *session.Session, config ...*aws.Config) *IoTSecureTunnelingActivities {
	client := iotsecuretunneling.New(session, config...)
	return &IoTSecureTunnelingActivities{client: client}
}

func (a *IoTSecureTunnelingActivities) CloseTunnel(input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error) {
	return a.client.CloseTunnel(input)
}

func (a *IoTSecureTunnelingActivities) DescribeTunnel(input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	return a.client.DescribeTunnel(input)
}

func (a *IoTSecureTunnelingActivities) ListTagsForResource(input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *IoTSecureTunnelingActivities) ListTunnels(input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error) {
	return a.client.ListTunnels(input)
}

func (a *IoTSecureTunnelingActivities) OpenTunnel(input *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error) {
	return a.client.OpenTunnel(input)
}

func (a *IoTSecureTunnelingActivities) TagResource(input *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *IoTSecureTunnelingActivities) UntagResource(input *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}
