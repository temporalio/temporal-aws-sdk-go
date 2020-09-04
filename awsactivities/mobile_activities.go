
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mobile"
	"github.com/aws/aws-sdk-go/service/mobile/mobileiface"
)

type MobileActivities struct {
	client mobileiface.MobileAPI
}

func NewMobileActivities(client mobileiface.MobileAPI) *MobileActivities {
    return &MobileActivities{client: client}
}

func (a *MobileActivities) CreateProject(input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
    return a.client.CreateProject(input)
}

func (a *MobileActivities) DeleteProject(input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error) {
    return a.client.DeleteProject(input)
}

func (a *MobileActivities) DescribeBundle(input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error) {
    return a.client.DescribeBundle(input)
}

func (a *MobileActivities) DescribeProject(input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error) {
    return a.client.DescribeProject(input)
}

func (a *MobileActivities) ExportBundle(input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error) {
    return a.client.ExportBundle(input)
}

func (a *MobileActivities) ExportProject(input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error) {
    return a.client.ExportProject(input)
}

func (a *MobileActivities) ListBundles(input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error) {
    return a.client.ListBundles(input)
}

func (a *MobileActivities) ListProjects(input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error) {
    return a.client.ListProjects(input)
}

func (a *MobileActivities) UpdateProject(input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error) {
    return a.client.UpdateProject(input)
}
