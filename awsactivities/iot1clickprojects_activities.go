package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects/iot1clickprojectsiface"
)

type IoT1ClickProjectsActivities struct {
	client iot1clickprojectsiface.IoT1ClickProjectsAPI
}

func NewIoT1ClickProjectsActivities(session *session.Session, config ...*aws.Config) *IoT1ClickProjectsActivities {
	client := iot1clickprojects.New(session, config...)
	return &IoT1ClickProjectsActivities{client: client}
}

func (a *IoT1ClickProjectsActivities) AssociateDeviceWithPlacement(input *iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
	return a.client.AssociateDeviceWithPlacement(input)
}

func (a *IoT1ClickProjectsActivities) CreatePlacement(input *iot1clickprojects.CreatePlacementInput) (*iot1clickprojects.CreatePlacementOutput, error) {
	return a.client.CreatePlacement(input)
}

func (a *IoT1ClickProjectsActivities) CreateProject(input *iot1clickprojects.CreateProjectInput) (*iot1clickprojects.CreateProjectOutput, error) {
	return a.client.CreateProject(input)
}

func (a *IoT1ClickProjectsActivities) DeletePlacement(input *iot1clickprojects.DeletePlacementInput) (*iot1clickprojects.DeletePlacementOutput, error) {
	return a.client.DeletePlacement(input)
}

func (a *IoT1ClickProjectsActivities) DeleteProject(input *iot1clickprojects.DeleteProjectInput) (*iot1clickprojects.DeleteProjectOutput, error) {
	return a.client.DeleteProject(input)
}

func (a *IoT1ClickProjectsActivities) DescribePlacement(input *iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error) {
	return a.client.DescribePlacement(input)
}

func (a *IoT1ClickProjectsActivities) DescribeProject(input *iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error) {
	return a.client.DescribeProject(input)
}

func (a *IoT1ClickProjectsActivities) DisassociateDeviceFromPlacement(input *iot1clickprojects.DisassociateDeviceFromPlacementInput) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
	return a.client.DisassociateDeviceFromPlacement(input)
}

func (a *IoT1ClickProjectsActivities) GetDevicesInPlacement(input *iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	return a.client.GetDevicesInPlacement(input)
}

func (a *IoT1ClickProjectsActivities) ListPlacements(input *iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error) {
	return a.client.ListPlacements(input)
}

func (a *IoT1ClickProjectsActivities) ListProjects(input *iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error) {
	return a.client.ListProjects(input)
}

func (a *IoT1ClickProjectsActivities) ListTagsForResource(input *iot1clickprojects.ListTagsForResourceInput) (*iot1clickprojects.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *IoT1ClickProjectsActivities) TagResource(input *iot1clickprojects.TagResourceInput) (*iot1clickprojects.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *IoT1ClickProjectsActivities) UntagResource(input *iot1clickprojects.UntagResourceInput) (*iot1clickprojects.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *IoT1ClickProjectsActivities) UpdatePlacement(input *iot1clickprojects.UpdatePlacementInput) (*iot1clickprojects.UpdatePlacementOutput, error) {
	return a.client.UpdatePlacement(input)
}

func (a *IoT1ClickProjectsActivities) UpdateProject(input *iot1clickprojects.UpdateProjectInput) (*iot1clickprojects.UpdateProjectOutput, error) {
	return a.client.UpdateProject(input)
}
