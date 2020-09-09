package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/aws/aws-sdk-go/service/globalaccelerator/globalacceleratoriface"
)

type GlobalAcceleratorActivities struct {
	client globalacceleratoriface.GlobalAcceleratorAPI
}

func NewGlobalAcceleratorActivities(session *session.Session, config ...*aws.Config) *GlobalAcceleratorActivities {
	client := globalaccelerator.New(session, config...)
	return &GlobalAcceleratorActivities{client: client}
}

func (a *GlobalAcceleratorActivities) AdvertiseByoipCidr(input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
	return a.client.AdvertiseByoipCidr(input)
}

func (a *GlobalAcceleratorActivities) CreateAccelerator(input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error) {
	return a.client.CreateAccelerator(input)
}

func (a *GlobalAcceleratorActivities) CreateEndpointGroup(input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error) {
	return a.client.CreateEndpointGroup(input)
}

func (a *GlobalAcceleratorActivities) CreateListener(input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error) {
	return a.client.CreateListener(input)
}

func (a *GlobalAcceleratorActivities) DeleteAccelerator(input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error) {
	return a.client.DeleteAccelerator(input)
}

func (a *GlobalAcceleratorActivities) DeleteEndpointGroup(input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
	return a.client.DeleteEndpointGroup(input)
}

func (a *GlobalAcceleratorActivities) DeleteListener(input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error) {
	return a.client.DeleteListener(input)
}

func (a *GlobalAcceleratorActivities) DeprovisionByoipCidr(input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
	return a.client.DeprovisionByoipCidr(input)
}

func (a *GlobalAcceleratorActivities) DescribeAccelerator(input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	return a.client.DescribeAccelerator(input)
}

func (a *GlobalAcceleratorActivities) DescribeAcceleratorAttributes(input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	return a.client.DescribeAcceleratorAttributes(input)
}

func (a *GlobalAcceleratorActivities) DescribeEndpointGroup(input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	return a.client.DescribeEndpointGroup(input)
}

func (a *GlobalAcceleratorActivities) DescribeListener(input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error) {
	return a.client.DescribeListener(input)
}

func (a *GlobalAcceleratorActivities) ListAccelerators(input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error) {
	return a.client.ListAccelerators(input)
}

func (a *GlobalAcceleratorActivities) ListByoipCidrs(input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error) {
	return a.client.ListByoipCidrs(input)
}

func (a *GlobalAcceleratorActivities) ListEndpointGroups(input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	return a.client.ListEndpointGroups(input)
}

func (a *GlobalAcceleratorActivities) ListListeners(input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error) {
	return a.client.ListListeners(input)
}

func (a *GlobalAcceleratorActivities) ListTagsForResource(input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *GlobalAcceleratorActivities) ProvisionByoipCidr(input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
	return a.client.ProvisionByoipCidr(input)
}

func (a *GlobalAcceleratorActivities) TagResource(input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *GlobalAcceleratorActivities) UntagResource(input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *GlobalAcceleratorActivities) UpdateAccelerator(input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error) {
	return a.client.UpdateAccelerator(input)
}

func (a *GlobalAcceleratorActivities) UpdateAcceleratorAttributes(input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
	return a.client.UpdateAcceleratorAttributes(input)
}

func (a *GlobalAcceleratorActivities) UpdateEndpointGroup(input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
	return a.client.UpdateEndpointGroup(input)
}

func (a *GlobalAcceleratorActivities) UpdateListener(input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error) {
	return a.client.UpdateListener(input)
}

func (a *GlobalAcceleratorActivities) WithdrawByoipCidr(input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
	return a.client.WithdrawByoipCidr(input)
}
