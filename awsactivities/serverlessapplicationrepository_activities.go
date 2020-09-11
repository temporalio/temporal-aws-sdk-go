
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository/serverlessapplicationrepositoryiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ServerlessApplicationRepositoryActivities struct {
    client serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI
}

func NewServerlessApplicationRepositoryActivities(session *session.Session, config ...*aws.Config) *ServerlessApplicationRepositoryActivities {
    client := serverlessapplicationrepository.New(session, config...)
    return &ServerlessApplicationRepositoryActivities{client: client}
}

func (a *ServerlessApplicationRepositoryActivities) CreateApplication(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
    return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) CreateApplicationVersion(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
    return a.client.CreateApplicationVersionWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) CreateCloudFormationChangeSet(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateCloudFormationChangeSetWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) CreateCloudFormationTemplate(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
    return a.client.CreateCloudFormationTemplateWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) DeleteApplication(ctx context.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
    return a.client.DeleteApplicationWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) GetApplication(ctx context.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error) {
    return a.client.GetApplicationWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) GetApplicationPolicy(ctx context.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
    return a.client.GetApplicationPolicyWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) GetCloudFormationTemplate(ctx context.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
    return a.client.GetCloudFormationTemplateWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) ListApplicationDependencies(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
    return a.client.ListApplicationDependenciesWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) ListApplicationVersions(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
    return a.client.ListApplicationVersionsWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) ListApplications(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
    return a.client.ListApplicationsWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) PutApplicationPolicy(ctx context.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
    return a.client.PutApplicationPolicyWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) UnshareApplication(ctx context.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
    return a.client.UnshareApplicationWithContext(ctx, input)
}

func (a *ServerlessApplicationRepositoryActivities) UpdateApplication(ctx context.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
    return a.client.UpdateApplicationWithContext(ctx, input)
}
