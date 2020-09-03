package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/appflow"
	"github.com/aws/aws-sdk-go/service/appflow/appflowiface"
)

type AppflowActivities struct {
	client appflowiface.AppflowAPI
}

func NewAppflowActivities(client appflowiface.AppflowAPI) *AppflowActivities {
    return &AppflowActivities{client: client}
}


func (a *AppflowActivities) CreateConnectorProfile(input *appflow.CreateConnectorProfileInput) (*appflow.CreateConnectorProfileOutput, error) {
    return a.client.CreateConnectorProfile(input)
}



func (a *AppflowActivities) CreateFlow(input *appflow.CreateFlowInput) (*appflow.CreateFlowOutput, error) {
    return a.client.CreateFlow(input)
}



func (a *AppflowActivities) DeleteConnectorProfile(input *appflow.DeleteConnectorProfileInput) (*appflow.DeleteConnectorProfileOutput, error) {
    return a.client.DeleteConnectorProfile(input)
}



func (a *AppflowActivities) DeleteFlow(input *appflow.DeleteFlowInput) (*appflow.DeleteFlowOutput, error) {
    return a.client.DeleteFlow(input)
}



func (a *AppflowActivities) DescribeConnectorEntity(input *appflow.DescribeConnectorEntityInput) (*appflow.DescribeConnectorEntityOutput, error) {
    return a.client.DescribeConnectorEntity(input)
}



func (a *AppflowActivities) DescribeConnectorProfiles(input *appflow.DescribeConnectorProfilesInput) (*appflow.DescribeConnectorProfilesOutput, error) {
    return a.client.DescribeConnectorProfiles(input)
}



func (a *AppflowActivities) DescribeConnectors(input *appflow.DescribeConnectorsInput) (*appflow.DescribeConnectorsOutput, error) {
    return a.client.DescribeConnectors(input)
}



func (a *AppflowActivities) DescribeFlow(input *appflow.DescribeFlowInput) (*appflow.DescribeFlowOutput, error) {
    return a.client.DescribeFlow(input)
}



func (a *AppflowActivities) DescribeFlowExecutionRecords(input *appflow.DescribeFlowExecutionRecordsInput) (*appflow.DescribeFlowExecutionRecordsOutput, error) {
    return a.client.DescribeFlowExecutionRecords(input)
}



func (a *AppflowActivities) ListConnectorEntities(input *appflow.ListConnectorEntitiesInput) (*appflow.ListConnectorEntitiesOutput, error) {
    return a.client.ListConnectorEntities(input)
}



func (a *AppflowActivities) ListFlows(input *appflow.ListFlowsInput) (*appflow.ListFlowsOutput, error) {
    return a.client.ListFlows(input)
}



func (a *AppflowActivities) ListTagsForResource(input *appflow.ListTagsForResourceInput) (*appflow.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *AppflowActivities) StartFlow(input *appflow.StartFlowInput) (*appflow.StartFlowOutput, error) {
    return a.client.StartFlow(input)
}



func (a *AppflowActivities) StopFlow(input *appflow.StopFlowInput) (*appflow.StopFlowOutput, error) {
    return a.client.StopFlow(input)
}



func (a *AppflowActivities) TagResource(input *appflow.TagResourceInput) (*appflow.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *AppflowActivities) UntagResource(input *appflow.UntagResourceInput) (*appflow.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *AppflowActivities) UpdateConnectorProfile(input *appflow.UpdateConnectorProfileInput) (*appflow.UpdateConnectorProfileOutput, error) {
    return a.client.UpdateConnectorProfile(input)
}



func (a *AppflowActivities) UpdateFlow(input *appflow.UpdateFlowInput) (*appflow.UpdateFlowOutput, error) {
    return a.client.UpdateFlow(input)
}

