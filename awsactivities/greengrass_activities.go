package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/greengrass/greengrassiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type GreengrassActivities struct {
	client greengrassiface.GreengrassAPI
}

func NewGreengrassActivities(session *session.Session, config ...*aws.Config) *GreengrassActivities {
	client := greengrass.New(session, config...)
	return &GreengrassActivities{client: client}
}

func (a *GreengrassActivities) AssociateRoleToGroup(ctx context.Context, input *greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error) {
	return a.client.AssociateRoleToGroupWithContext(ctx, input)
}

func (a *GreengrassActivities) AssociateServiceRoleToAccount(ctx context.Context, input *greengrass.AssociateServiceRoleToAccountInput) (*greengrass.AssociateServiceRoleToAccountOutput, error) {
	return a.client.AssociateServiceRoleToAccountWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateConnectorDefinition(ctx context.Context, input *greengrass.CreateConnectorDefinitionInput) (*greengrass.CreateConnectorDefinitionOutput, error) {
	return a.client.CreateConnectorDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateConnectorDefinitionVersion(ctx context.Context, input *greengrass.CreateConnectorDefinitionVersionInput) (*greengrass.CreateConnectorDefinitionVersionOutput, error) {
	return a.client.CreateConnectorDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateCoreDefinition(ctx context.Context, input *greengrass.CreateCoreDefinitionInput) (*greengrass.CreateCoreDefinitionOutput, error) {
	return a.client.CreateCoreDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateCoreDefinitionVersion(ctx context.Context, input *greengrass.CreateCoreDefinitionVersionInput) (*greengrass.CreateCoreDefinitionVersionOutput, error) {
	return a.client.CreateCoreDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateDeployment(ctx context.Context, input *greengrass.CreateDeploymentInput) (*greengrass.CreateDeploymentOutput, error) {
	return a.client.CreateDeploymentWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateDeviceDefinition(ctx context.Context, input *greengrass.CreateDeviceDefinitionInput) (*greengrass.CreateDeviceDefinitionOutput, error) {
	return a.client.CreateDeviceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateDeviceDefinitionVersion(ctx context.Context, input *greengrass.CreateDeviceDefinitionVersionInput) (*greengrass.CreateDeviceDefinitionVersionOutput, error) {
	return a.client.CreateDeviceDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateFunctionDefinition(ctx context.Context, input *greengrass.CreateFunctionDefinitionInput) (*greengrass.CreateFunctionDefinitionOutput, error) {
	return a.client.CreateFunctionDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateFunctionDefinitionVersion(ctx context.Context, input *greengrass.CreateFunctionDefinitionVersionInput) (*greengrass.CreateFunctionDefinitionVersionOutput, error) {
	return a.client.CreateFunctionDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateGroup(ctx context.Context, input *greengrass.CreateGroupInput) (*greengrass.CreateGroupOutput, error) {
	return a.client.CreateGroupWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateGroupCertificateAuthority(ctx context.Context, input *greengrass.CreateGroupCertificateAuthorityInput) (*greengrass.CreateGroupCertificateAuthorityOutput, error) {
	return a.client.CreateGroupCertificateAuthorityWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateGroupVersion(ctx context.Context, input *greengrass.CreateGroupVersionInput) (*greengrass.CreateGroupVersionOutput, error) {
	return a.client.CreateGroupVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateLoggerDefinition(ctx context.Context, input *greengrass.CreateLoggerDefinitionInput) (*greengrass.CreateLoggerDefinitionOutput, error) {
	return a.client.CreateLoggerDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateLoggerDefinitionVersion(ctx context.Context, input *greengrass.CreateLoggerDefinitionVersionInput) (*greengrass.CreateLoggerDefinitionVersionOutput, error) {
	return a.client.CreateLoggerDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateResourceDefinition(ctx context.Context, input *greengrass.CreateResourceDefinitionInput) (*greengrass.CreateResourceDefinitionOutput, error) {
	return a.client.CreateResourceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateResourceDefinitionVersion(ctx context.Context, input *greengrass.CreateResourceDefinitionVersionInput) (*greengrass.CreateResourceDefinitionVersionOutput, error) {
	return a.client.CreateResourceDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateSoftwareUpdateJob(ctx context.Context, input *greengrass.CreateSoftwareUpdateJobInput) (*greengrass.CreateSoftwareUpdateJobOutput, error) {
	return a.client.CreateSoftwareUpdateJobWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateSubscriptionDefinition(ctx context.Context, input *greengrass.CreateSubscriptionDefinitionInput) (*greengrass.CreateSubscriptionDefinitionOutput, error) {
	return a.client.CreateSubscriptionDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) CreateSubscriptionDefinitionVersion(ctx context.Context, input *greengrass.CreateSubscriptionDefinitionVersionInput) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error) {
	return a.client.CreateSubscriptionDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteConnectorDefinition(ctx context.Context, input *greengrass.DeleteConnectorDefinitionInput) (*greengrass.DeleteConnectorDefinitionOutput, error) {
	return a.client.DeleteConnectorDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteCoreDefinition(ctx context.Context, input *greengrass.DeleteCoreDefinitionInput) (*greengrass.DeleteCoreDefinitionOutput, error) {
	return a.client.DeleteCoreDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteDeviceDefinition(ctx context.Context, input *greengrass.DeleteDeviceDefinitionInput) (*greengrass.DeleteDeviceDefinitionOutput, error) {
	return a.client.DeleteDeviceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteFunctionDefinition(ctx context.Context, input *greengrass.DeleteFunctionDefinitionInput) (*greengrass.DeleteFunctionDefinitionOutput, error) {
	return a.client.DeleteFunctionDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteGroup(ctx context.Context, input *greengrass.DeleteGroupInput) (*greengrass.DeleteGroupOutput, error) {
	return a.client.DeleteGroupWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteLoggerDefinition(ctx context.Context, input *greengrass.DeleteLoggerDefinitionInput) (*greengrass.DeleteLoggerDefinitionOutput, error) {
	return a.client.DeleteLoggerDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteResourceDefinition(ctx context.Context, input *greengrass.DeleteResourceDefinitionInput) (*greengrass.DeleteResourceDefinitionOutput, error) {
	return a.client.DeleteResourceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) DeleteSubscriptionDefinition(ctx context.Context, input *greengrass.DeleteSubscriptionDefinitionInput) (*greengrass.DeleteSubscriptionDefinitionOutput, error) {
	return a.client.DeleteSubscriptionDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) DisassociateRoleFromGroup(ctx context.Context, input *greengrass.DisassociateRoleFromGroupInput) (*greengrass.DisassociateRoleFromGroupOutput, error) {
	return a.client.DisassociateRoleFromGroupWithContext(ctx, input)
}

func (a *GreengrassActivities) DisassociateServiceRoleFromAccount(ctx context.Context, input *greengrass.DisassociateServiceRoleFromAccountInput) (*greengrass.DisassociateServiceRoleFromAccountOutput, error) {
	return a.client.DisassociateServiceRoleFromAccountWithContext(ctx, input)
}

func (a *GreengrassActivities) GetAssociatedRole(ctx context.Context, input *greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error) {
	return a.client.GetAssociatedRoleWithContext(ctx, input)
}

func (a *GreengrassActivities) GetBulkDeploymentStatus(ctx context.Context, input *greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error) {
	return a.client.GetBulkDeploymentStatusWithContext(ctx, input)
}

func (a *GreengrassActivities) GetConnectivityInfo(ctx context.Context, input *greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error) {
	return a.client.GetConnectivityInfoWithContext(ctx, input)
}

func (a *GreengrassActivities) GetConnectorDefinition(ctx context.Context, input *greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error) {
	return a.client.GetConnectorDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetConnectorDefinitionVersion(ctx context.Context, input *greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error) {
	return a.client.GetConnectorDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetCoreDefinition(ctx context.Context, input *greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error) {
	return a.client.GetCoreDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetCoreDefinitionVersion(ctx context.Context, input *greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error) {
	return a.client.GetCoreDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetDeploymentStatus(ctx context.Context, input *greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error) {
	return a.client.GetDeploymentStatusWithContext(ctx, input)
}

func (a *GreengrassActivities) GetDeviceDefinition(ctx context.Context, input *greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error) {
	return a.client.GetDeviceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetDeviceDefinitionVersion(ctx context.Context, input *greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error) {
	return a.client.GetDeviceDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetFunctionDefinition(ctx context.Context, input *greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error) {
	return a.client.GetFunctionDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetFunctionDefinitionVersion(ctx context.Context, input *greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error) {
	return a.client.GetFunctionDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetGroup(ctx context.Context, input *greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error) {
	return a.client.GetGroupWithContext(ctx, input)
}

func (a *GreengrassActivities) GetGroupCertificateAuthority(ctx context.Context, input *greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error) {
	return a.client.GetGroupCertificateAuthorityWithContext(ctx, input)
}

func (a *GreengrassActivities) GetGroupCertificateConfiguration(ctx context.Context, input *greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error) {
	return a.client.GetGroupCertificateConfigurationWithContext(ctx, input)
}

func (a *GreengrassActivities) GetGroupVersion(ctx context.Context, input *greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error) {
	return a.client.GetGroupVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetLoggerDefinition(ctx context.Context, input *greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error) {
	return a.client.GetLoggerDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetLoggerDefinitionVersion(ctx context.Context, input *greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error) {
	return a.client.GetLoggerDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetResourceDefinition(ctx context.Context, input *greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error) {
	return a.client.GetResourceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetResourceDefinitionVersion(ctx context.Context, input *greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error) {
	return a.client.GetResourceDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetServiceRoleForAccount(ctx context.Context, input *greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error) {
	return a.client.GetServiceRoleForAccountWithContext(ctx, input)
}

func (a *GreengrassActivities) GetSubscriptionDefinition(ctx context.Context, input *greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error) {
	return a.client.GetSubscriptionDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) GetSubscriptionDefinitionVersion(ctx context.Context, input *greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error) {
	return a.client.GetSubscriptionDefinitionVersionWithContext(ctx, input)
}

func (a *GreengrassActivities) ListBulkDeploymentDetailedReports(ctx context.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error) {
	return a.client.ListBulkDeploymentDetailedReportsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListBulkDeployments(ctx context.Context, input *greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error) {
	return a.client.ListBulkDeploymentsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListConnectorDefinitionVersions(ctx context.Context, input *greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error) {
	return a.client.ListConnectorDefinitionVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListConnectorDefinitions(ctx context.Context, input *greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error) {
	return a.client.ListConnectorDefinitionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListCoreDefinitionVersions(ctx context.Context, input *greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error) {
	return a.client.ListCoreDefinitionVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListCoreDefinitions(ctx context.Context, input *greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error) {
	return a.client.ListCoreDefinitionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListDeployments(ctx context.Context, input *greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error) {
	return a.client.ListDeploymentsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListDeviceDefinitionVersions(ctx context.Context, input *greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error) {
	return a.client.ListDeviceDefinitionVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListDeviceDefinitions(ctx context.Context, input *greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error) {
	return a.client.ListDeviceDefinitionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListFunctionDefinitionVersions(ctx context.Context, input *greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error) {
	return a.client.ListFunctionDefinitionVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListFunctionDefinitions(ctx context.Context, input *greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error) {
	return a.client.ListFunctionDefinitionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListGroupCertificateAuthorities(ctx context.Context, input *greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error) {
	return a.client.ListGroupCertificateAuthoritiesWithContext(ctx, input)
}

func (a *GreengrassActivities) ListGroupVersions(ctx context.Context, input *greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error) {
	return a.client.ListGroupVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListGroups(ctx context.Context, input *greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error) {
	return a.client.ListGroupsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListLoggerDefinitionVersions(ctx context.Context, input *greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error) {
	return a.client.ListLoggerDefinitionVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListLoggerDefinitions(ctx context.Context, input *greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error) {
	return a.client.ListLoggerDefinitionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListResourceDefinitionVersions(ctx context.Context, input *greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error) {
	return a.client.ListResourceDefinitionVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListResourceDefinitions(ctx context.Context, input *greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error) {
	return a.client.ListResourceDefinitionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListSubscriptionDefinitionVersions(ctx context.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error) {
	return a.client.ListSubscriptionDefinitionVersionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListSubscriptionDefinitions(ctx context.Context, input *greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error) {
	return a.client.ListSubscriptionDefinitionsWithContext(ctx, input)
}

func (a *GreengrassActivities) ListTagsForResource(ctx context.Context, input *greengrass.ListTagsForResourceInput) (*greengrass.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *GreengrassActivities) ResetDeployments(ctx context.Context, input *greengrass.ResetDeploymentsInput) (*greengrass.ResetDeploymentsOutput, error) {
	return a.client.ResetDeploymentsWithContext(ctx, input)
}

func (a *GreengrassActivities) StartBulkDeployment(ctx context.Context, input *greengrass.StartBulkDeploymentInput) (*greengrass.StartBulkDeploymentOutput, error) {
	return a.client.StartBulkDeploymentWithContext(ctx, input)
}

func (a *GreengrassActivities) StopBulkDeployment(ctx context.Context, input *greengrass.StopBulkDeploymentInput) (*greengrass.StopBulkDeploymentOutput, error) {
	return a.client.StopBulkDeploymentWithContext(ctx, input)
}

func (a *GreengrassActivities) TagResource(ctx context.Context, input *greengrass.TagResourceInput) (*greengrass.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *GreengrassActivities) UntagResource(ctx context.Context, input *greengrass.UntagResourceInput) (*greengrass.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateConnectivityInfo(ctx context.Context, input *greengrass.UpdateConnectivityInfoInput) (*greengrass.UpdateConnectivityInfoOutput, error) {
	return a.client.UpdateConnectivityInfoWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateConnectorDefinition(ctx context.Context, input *greengrass.UpdateConnectorDefinitionInput) (*greengrass.UpdateConnectorDefinitionOutput, error) {
	return a.client.UpdateConnectorDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateCoreDefinition(ctx context.Context, input *greengrass.UpdateCoreDefinitionInput) (*greengrass.UpdateCoreDefinitionOutput, error) {
	return a.client.UpdateCoreDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateDeviceDefinition(ctx context.Context, input *greengrass.UpdateDeviceDefinitionInput) (*greengrass.UpdateDeviceDefinitionOutput, error) {
	return a.client.UpdateDeviceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateFunctionDefinition(ctx context.Context, input *greengrass.UpdateFunctionDefinitionInput) (*greengrass.UpdateFunctionDefinitionOutput, error) {
	return a.client.UpdateFunctionDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateGroup(ctx context.Context, input *greengrass.UpdateGroupInput) (*greengrass.UpdateGroupOutput, error) {
	return a.client.UpdateGroupWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateGroupCertificateConfiguration(ctx context.Context, input *greengrass.UpdateGroupCertificateConfigurationInput) (*greengrass.UpdateGroupCertificateConfigurationOutput, error) {
	return a.client.UpdateGroupCertificateConfigurationWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateLoggerDefinition(ctx context.Context, input *greengrass.UpdateLoggerDefinitionInput) (*greengrass.UpdateLoggerDefinitionOutput, error) {
	return a.client.UpdateLoggerDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateResourceDefinition(ctx context.Context, input *greengrass.UpdateResourceDefinitionInput) (*greengrass.UpdateResourceDefinitionOutput, error) {
	return a.client.UpdateResourceDefinitionWithContext(ctx, input)
}

func (a *GreengrassActivities) UpdateSubscriptionDefinition(ctx context.Context, input *greengrass.UpdateSubscriptionDefinitionInput) (*greengrass.UpdateSubscriptionDefinitionOutput, error) {
	return a.client.UpdateSubscriptionDefinitionWithContext(ctx, input)
}
