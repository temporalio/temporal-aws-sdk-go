package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository/serverlessapplicationrepositoryiface"
)

type ServerlessApplicationRepositoryActivities struct {
	client serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI
}

func NewServerlessApplicationRepositoryActivities(client serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI) *ServerlessApplicationRepositoryActivities {
    return &ServerlessApplicationRepositoryActivities{client: client}
}

func (a *ServerlessApplicationRepositoryActivities) CreateApplication(input *serverlessapplicationrepository.CreateApplicationInput) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
    return a.client.CreateApplication(input)
}

func (a *ServerlessApplicationRepositoryActivities) CreateApplicationVersion(input *serverlessapplicationrepository.CreateApplicationVersionInput) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
    return a.client.CreateApplicationVersion(input)
}

func (a *ServerlessApplicationRepositoryActivities) CreateCloudFormationChangeSet(input *serverlessapplicationrepository.CreateCloudFormationChangeSetInput) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
    return a.client.CreateCloudFormationChangeSet(input)
}

func (a *ServerlessApplicationRepositoryActivities) CreateCloudFormationTemplate(input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
    return a.client.CreateCloudFormationTemplate(input)
}

func (a *ServerlessApplicationRepositoryActivities) DeleteApplication(input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
    return a.client.DeleteApplication(input)
}

func (a *ServerlessApplicationRepositoryActivities) GetApplication(input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error) {
    return a.client.GetApplication(input)
}

func (a *ServerlessApplicationRepositoryActivities) GetApplicationPolicy(input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
    return a.client.GetApplicationPolicy(input)
}

func (a *ServerlessApplicationRepositoryActivities) GetCloudFormationTemplate(input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
    return a.client.GetCloudFormationTemplate(input)
}

func (a *ServerlessApplicationRepositoryActivities) ListApplicationDependencies(input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
    return a.client.ListApplicationDependencies(input)
}

func (a *ServerlessApplicationRepositoryActivities) ListApplicationVersions(input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
    return a.client.ListApplicationVersions(input)
}

func (a *ServerlessApplicationRepositoryActivities) ListApplications(input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
    return a.client.ListApplications(input)
}

func (a *ServerlessApplicationRepositoryActivities) PutApplicationPolicy(input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
    return a.client.PutApplicationPolicy(input)
}

func (a *ServerlessApplicationRepositoryActivities) UnshareApplication(input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
    return a.client.UnshareApplication(input)
}

func (a *ServerlessApplicationRepositoryActivities) UpdateApplication(input *serverlessapplicationrepository.UpdateApplicationInput) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
    return a.client.UpdateApplication(input)
}
