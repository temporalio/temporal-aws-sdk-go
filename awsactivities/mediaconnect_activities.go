package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/mediaconnect/mediaconnectiface"
)

type MediaConnectActivities struct {
	client mediaconnectiface.MediaConnectAPI
}

func NewMediaConnectActivities(session *session.Session, config ...*aws.Config) *MediaConnectActivities {
	client := mediaconnect.New(session, config...)
	return &MediaConnectActivities{client: client}
}

func (a *MediaConnectActivities) AddFlowOutputs(input *mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error) {
	return a.client.AddFlowOutputs(input)
}

func (a *MediaConnectActivities) AddFlowSources(input *mediaconnect.AddFlowSourcesInput) (*mediaconnect.AddFlowSourcesOutput, error) {
	return a.client.AddFlowSources(input)
}

func (a *MediaConnectActivities) AddFlowVpcInterfaces(input *mediaconnect.AddFlowVpcInterfacesInput) (*mediaconnect.AddFlowVpcInterfacesOutput, error) {
	return a.client.AddFlowVpcInterfaces(input)
}

func (a *MediaConnectActivities) CreateFlow(input *mediaconnect.CreateFlowInput) (*mediaconnect.CreateFlowOutput, error) {
	return a.client.CreateFlow(input)
}

func (a *MediaConnectActivities) DeleteFlow(input *mediaconnect.DeleteFlowInput) (*mediaconnect.DeleteFlowOutput, error) {
	return a.client.DeleteFlow(input)
}

func (a *MediaConnectActivities) DescribeFlow(input *mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error) {
	return a.client.DescribeFlow(input)
}

func (a *MediaConnectActivities) GrantFlowEntitlements(input *mediaconnect.GrantFlowEntitlementsInput) (*mediaconnect.GrantFlowEntitlementsOutput, error) {
	return a.client.GrantFlowEntitlements(input)
}

func (a *MediaConnectActivities) ListEntitlements(input *mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error) {
	return a.client.ListEntitlements(input)
}

func (a *MediaConnectActivities) ListFlows(input *mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error) {
	return a.client.ListFlows(input)
}

func (a *MediaConnectActivities) ListTagsForResource(input *mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *MediaConnectActivities) RemoveFlowOutput(input *mediaconnect.RemoveFlowOutputInput) (*mediaconnect.RemoveFlowOutputOutput, error) {
	return a.client.RemoveFlowOutput(input)
}

func (a *MediaConnectActivities) RemoveFlowSource(input *mediaconnect.RemoveFlowSourceInput) (*mediaconnect.RemoveFlowSourceOutput, error) {
	return a.client.RemoveFlowSource(input)
}

func (a *MediaConnectActivities) RemoveFlowVpcInterface(input *mediaconnect.RemoveFlowVpcInterfaceInput) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error) {
	return a.client.RemoveFlowVpcInterface(input)
}

func (a *MediaConnectActivities) RevokeFlowEntitlement(input *mediaconnect.RevokeFlowEntitlementInput) (*mediaconnect.RevokeFlowEntitlementOutput, error) {
	return a.client.RevokeFlowEntitlement(input)
}

func (a *MediaConnectActivities) StartFlow(input *mediaconnect.StartFlowInput) (*mediaconnect.StartFlowOutput, error) {
	return a.client.StartFlow(input)
}

func (a *MediaConnectActivities) StopFlow(input *mediaconnect.StopFlowInput) (*mediaconnect.StopFlowOutput, error) {
	return a.client.StopFlow(input)
}

func (a *MediaConnectActivities) TagResource(input *mediaconnect.TagResourceInput) (*mediaconnect.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *MediaConnectActivities) UntagResource(input *mediaconnect.UntagResourceInput) (*mediaconnect.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *MediaConnectActivities) UpdateFlow(input *mediaconnect.UpdateFlowInput) (*mediaconnect.UpdateFlowOutput, error) {
	return a.client.UpdateFlow(input)
}

func (a *MediaConnectActivities) UpdateFlowEntitlement(input *mediaconnect.UpdateFlowEntitlementInput) (*mediaconnect.UpdateFlowEntitlementOutput, error) {
	return a.client.UpdateFlowEntitlement(input)
}

func (a *MediaConnectActivities) UpdateFlowOutput(input *mediaconnect.UpdateFlowOutputInput) (*mediaconnect.UpdateFlowOutputOutput, error) {
	return a.client.UpdateFlowOutput(input)
}

func (a *MediaConnectActivities) UpdateFlowSource(input *mediaconnect.UpdateFlowSourceInput) (*mediaconnect.UpdateFlowSourceOutput, error) {
	return a.client.UpdateFlowSource(input)
}
