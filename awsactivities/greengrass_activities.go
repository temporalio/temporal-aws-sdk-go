
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/greengrass/greengrassiface"
)

type GreengrassActivities struct {
	client greengrassiface.GreengrassAPI
}

func NewGreengrassActivities(client greengrassiface.GreengrassAPI) *GreengrassActivities {
    return &GreengrassActivities{client: client}
}

func (a *GreengrassActivities) AssociateRoleToGroup(input *greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error) {
    return a.client.AssociateRoleToGroup(input)
}

func (a *GreengrassActivities) AssociateServiceRoleToAccount(input *greengrass.AssociateServiceRoleToAccountInput) (*greengrass.AssociateServiceRoleToAccountOutput, error) {
    return a.client.AssociateServiceRoleToAccount(input)
}

func (a *GreengrassActivities) CreateConnectorDefinition(input *greengrass.CreateConnectorDefinitionInput) (*greengrass.CreateConnectorDefinitionOutput, error) {
    return a.client.CreateConnectorDefinition(input)
}

func (a *GreengrassActivities) CreateConnectorDefinitionVersion(input *greengrass.CreateConnectorDefinitionVersionInput) (*greengrass.CreateConnectorDefinitionVersionOutput, error) {
    return a.client.CreateConnectorDefinitionVersion(input)
}

func (a *GreengrassActivities) CreateCoreDefinition(input *greengrass.CreateCoreDefinitionInput) (*greengrass.CreateCoreDefinitionOutput, error) {
    return a.client.CreateCoreDefinition(input)
}

func (a *GreengrassActivities) CreateCoreDefinitionVersion(input *greengrass.CreateCoreDefinitionVersionInput) (*greengrass.CreateCoreDefinitionVersionOutput, error) {
    return a.client.CreateCoreDefinitionVersion(input)
}

func (a *GreengrassActivities) CreateDeployment(input *greengrass.CreateDeploymentInput) (*greengrass.CreateDeploymentOutput, error) {
    return a.client.CreateDeployment(input)
}

func (a *GreengrassActivities) CreateDeviceDefinition(input *greengrass.CreateDeviceDefinitionInput) (*greengrass.CreateDeviceDefinitionOutput, error) {
    return a.client.CreateDeviceDefinition(input)
}

func (a *GreengrassActivities) CreateDeviceDefinitionVersion(input *greengrass.CreateDeviceDefinitionVersionInput) (*greengrass.CreateDeviceDefinitionVersionOutput, error) {
    return a.client.CreateDeviceDefinitionVersion(input)
}

func (a *GreengrassActivities) CreateFunctionDefinition(input *greengrass.CreateFunctionDefinitionInput) (*greengrass.CreateFunctionDefinitionOutput, error) {
    return a.client.CreateFunctionDefinition(input)
}

func (a *GreengrassActivities) CreateFunctionDefinitionVersion(input *greengrass.CreateFunctionDefinitionVersionInput) (*greengrass.CreateFunctionDefinitionVersionOutput, error) {
    return a.client.CreateFunctionDefinitionVersion(input)
}

func (a *GreengrassActivities) CreateGroup(input *greengrass.CreateGroupInput) (*greengrass.CreateGroupOutput, error) {
    return a.client.CreateGroup(input)
}

func (a *GreengrassActivities) CreateGroupCertificateAuthority(input *greengrass.CreateGroupCertificateAuthorityInput) (*greengrass.CreateGroupCertificateAuthorityOutput, error) {
    return a.client.CreateGroupCertificateAuthority(input)
}

func (a *GreengrassActivities) CreateGroupVersion(input *greengrass.CreateGroupVersionInput) (*greengrass.CreateGroupVersionOutput, error) {
    return a.client.CreateGroupVersion(input)
}

func (a *GreengrassActivities) CreateLoggerDefinition(input *greengrass.CreateLoggerDefinitionInput) (*greengrass.CreateLoggerDefinitionOutput, error) {
    return a.client.CreateLoggerDefinition(input)
}

func (a *GreengrassActivities) CreateLoggerDefinitionVersion(input *greengrass.CreateLoggerDefinitionVersionInput) (*greengrass.CreateLoggerDefinitionVersionOutput, error) {
    return a.client.CreateLoggerDefinitionVersion(input)
}

func (a *GreengrassActivities) CreateResourceDefinition(input *greengrass.CreateResourceDefinitionInput) (*greengrass.CreateResourceDefinitionOutput, error) {
    return a.client.CreateResourceDefinition(input)
}

func (a *GreengrassActivities) CreateResourceDefinitionVersion(input *greengrass.CreateResourceDefinitionVersionInput) (*greengrass.CreateResourceDefinitionVersionOutput, error) {
    return a.client.CreateResourceDefinitionVersion(input)
}

func (a *GreengrassActivities) CreateSoftwareUpdateJob(input *greengrass.CreateSoftwareUpdateJobInput) (*greengrass.CreateSoftwareUpdateJobOutput, error) {
    return a.client.CreateSoftwareUpdateJob(input)
}

func (a *GreengrassActivities) CreateSubscriptionDefinition(input *greengrass.CreateSubscriptionDefinitionInput) (*greengrass.CreateSubscriptionDefinitionOutput, error) {
    return a.client.CreateSubscriptionDefinition(input)
}

func (a *GreengrassActivities) CreateSubscriptionDefinitionVersion(input *greengrass.CreateSubscriptionDefinitionVersionInput) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error) {
    return a.client.CreateSubscriptionDefinitionVersion(input)
}

func (a *GreengrassActivities) DeleteConnectorDefinition(input *greengrass.DeleteConnectorDefinitionInput) (*greengrass.DeleteConnectorDefinitionOutput, error) {
    return a.client.DeleteConnectorDefinition(input)
}

func (a *GreengrassActivities) DeleteCoreDefinition(input *greengrass.DeleteCoreDefinitionInput) (*greengrass.DeleteCoreDefinitionOutput, error) {
    return a.client.DeleteCoreDefinition(input)
}

func (a *GreengrassActivities) DeleteDeviceDefinition(input *greengrass.DeleteDeviceDefinitionInput) (*greengrass.DeleteDeviceDefinitionOutput, error) {
    return a.client.DeleteDeviceDefinition(input)
}

func (a *GreengrassActivities) DeleteFunctionDefinition(input *greengrass.DeleteFunctionDefinitionInput) (*greengrass.DeleteFunctionDefinitionOutput, error) {
    return a.client.DeleteFunctionDefinition(input)
}

func (a *GreengrassActivities) DeleteGroup(input *greengrass.DeleteGroupInput) (*greengrass.DeleteGroupOutput, error) {
    return a.client.DeleteGroup(input)
}

func (a *GreengrassActivities) DeleteLoggerDefinition(input *greengrass.DeleteLoggerDefinitionInput) (*greengrass.DeleteLoggerDefinitionOutput, error) {
    return a.client.DeleteLoggerDefinition(input)
}

func (a *GreengrassActivities) DeleteResourceDefinition(input *greengrass.DeleteResourceDefinitionInput) (*greengrass.DeleteResourceDefinitionOutput, error) {
    return a.client.DeleteResourceDefinition(input)
}

func (a *GreengrassActivities) DeleteSubscriptionDefinition(input *greengrass.DeleteSubscriptionDefinitionInput) (*greengrass.DeleteSubscriptionDefinitionOutput, error) {
    return a.client.DeleteSubscriptionDefinition(input)
}

func (a *GreengrassActivities) DisassociateRoleFromGroup(input *greengrass.DisassociateRoleFromGroupInput) (*greengrass.DisassociateRoleFromGroupOutput, error) {
    return a.client.DisassociateRoleFromGroup(input)
}

func (a *GreengrassActivities) DisassociateServiceRoleFromAccount(input *greengrass.DisassociateServiceRoleFromAccountInput) (*greengrass.DisassociateServiceRoleFromAccountOutput, error) {
    return a.client.DisassociateServiceRoleFromAccount(input)
}

func (a *GreengrassActivities) GetAssociatedRole(input *greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error) {
    return a.client.GetAssociatedRole(input)
}

func (a *GreengrassActivities) GetBulkDeploymentStatus(input *greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error) {
    return a.client.GetBulkDeploymentStatus(input)
}

func (a *GreengrassActivities) GetConnectivityInfo(input *greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error) {
    return a.client.GetConnectivityInfo(input)
}

func (a *GreengrassActivities) GetConnectorDefinition(input *greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error) {
    return a.client.GetConnectorDefinition(input)
}

func (a *GreengrassActivities) GetConnectorDefinitionVersion(input *greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error) {
    return a.client.GetConnectorDefinitionVersion(input)
}

func (a *GreengrassActivities) GetCoreDefinition(input *greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error) {
    return a.client.GetCoreDefinition(input)
}

func (a *GreengrassActivities) GetCoreDefinitionVersion(input *greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error) {
    return a.client.GetCoreDefinitionVersion(input)
}

func (a *GreengrassActivities) GetDeploymentStatus(input *greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error) {
    return a.client.GetDeploymentStatus(input)
}

func (a *GreengrassActivities) GetDeviceDefinition(input *greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error) {
    return a.client.GetDeviceDefinition(input)
}

func (a *GreengrassActivities) GetDeviceDefinitionVersion(input *greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error) {
    return a.client.GetDeviceDefinitionVersion(input)
}

func (a *GreengrassActivities) GetFunctionDefinition(input *greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error) {
    return a.client.GetFunctionDefinition(input)
}

func (a *GreengrassActivities) GetFunctionDefinitionVersion(input *greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error) {
    return a.client.GetFunctionDefinitionVersion(input)
}

func (a *GreengrassActivities) GetGroup(input *greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error) {
    return a.client.GetGroup(input)
}

func (a *GreengrassActivities) GetGroupCertificateAuthority(input *greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error) {
    return a.client.GetGroupCertificateAuthority(input)
}

func (a *GreengrassActivities) GetGroupCertificateConfiguration(input *greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error) {
    return a.client.GetGroupCertificateConfiguration(input)
}

func (a *GreengrassActivities) GetGroupVersion(input *greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error) {
    return a.client.GetGroupVersion(input)
}

func (a *GreengrassActivities) GetLoggerDefinition(input *greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error) {
    return a.client.GetLoggerDefinition(input)
}

func (a *GreengrassActivities) GetLoggerDefinitionVersion(input *greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error) {
    return a.client.GetLoggerDefinitionVersion(input)
}

func (a *GreengrassActivities) GetResourceDefinition(input *greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error) {
    return a.client.GetResourceDefinition(input)
}

func (a *GreengrassActivities) GetResourceDefinitionVersion(input *greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error) {
    return a.client.GetResourceDefinitionVersion(input)
}

func (a *GreengrassActivities) GetServiceRoleForAccount(input *greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error) {
    return a.client.GetServiceRoleForAccount(input)
}

func (a *GreengrassActivities) GetSubscriptionDefinition(input *greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error) {
    return a.client.GetSubscriptionDefinition(input)
}

func (a *GreengrassActivities) GetSubscriptionDefinitionVersion(input *greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error) {
    return a.client.GetSubscriptionDefinitionVersion(input)
}

func (a *GreengrassActivities) ListBulkDeploymentDetailedReports(input *greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error) {
    return a.client.ListBulkDeploymentDetailedReports(input)
}

func (a *GreengrassActivities) ListBulkDeployments(input *greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error) {
    return a.client.ListBulkDeployments(input)
}

func (a *GreengrassActivities) ListConnectorDefinitionVersions(input *greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error) {
    return a.client.ListConnectorDefinitionVersions(input)
}

func (a *GreengrassActivities) ListConnectorDefinitions(input *greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error) {
    return a.client.ListConnectorDefinitions(input)
}

func (a *GreengrassActivities) ListCoreDefinitionVersions(input *greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error) {
    return a.client.ListCoreDefinitionVersions(input)
}

func (a *GreengrassActivities) ListCoreDefinitions(input *greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error) {
    return a.client.ListCoreDefinitions(input)
}

func (a *GreengrassActivities) ListDeployments(input *greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error) {
    return a.client.ListDeployments(input)
}

func (a *GreengrassActivities) ListDeviceDefinitionVersions(input *greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error) {
    return a.client.ListDeviceDefinitionVersions(input)
}

func (a *GreengrassActivities) ListDeviceDefinitions(input *greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error) {
    return a.client.ListDeviceDefinitions(input)
}

func (a *GreengrassActivities) ListFunctionDefinitionVersions(input *greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error) {
    return a.client.ListFunctionDefinitionVersions(input)
}

func (a *GreengrassActivities) ListFunctionDefinitions(input *greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error) {
    return a.client.ListFunctionDefinitions(input)
}

func (a *GreengrassActivities) ListGroupCertificateAuthorities(input *greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error) {
    return a.client.ListGroupCertificateAuthorities(input)
}

func (a *GreengrassActivities) ListGroupVersions(input *greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error) {
    return a.client.ListGroupVersions(input)
}

func (a *GreengrassActivities) ListGroups(input *greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error) {
    return a.client.ListGroups(input)
}

func (a *GreengrassActivities) ListLoggerDefinitionVersions(input *greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error) {
    return a.client.ListLoggerDefinitionVersions(input)
}

func (a *GreengrassActivities) ListLoggerDefinitions(input *greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error) {
    return a.client.ListLoggerDefinitions(input)
}

func (a *GreengrassActivities) ListResourceDefinitionVersions(input *greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error) {
    return a.client.ListResourceDefinitionVersions(input)
}

func (a *GreengrassActivities) ListResourceDefinitions(input *greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error) {
    return a.client.ListResourceDefinitions(input)
}

func (a *GreengrassActivities) ListSubscriptionDefinitionVersions(input *greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error) {
    return a.client.ListSubscriptionDefinitionVersions(input)
}

func (a *GreengrassActivities) ListSubscriptionDefinitions(input *greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error) {
    return a.client.ListSubscriptionDefinitions(input)
}

func (a *GreengrassActivities) ListTagsForResource(input *greengrass.ListTagsForResourceInput) (*greengrass.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *GreengrassActivities) ResetDeployments(input *greengrass.ResetDeploymentsInput) (*greengrass.ResetDeploymentsOutput, error) {
    return a.client.ResetDeployments(input)
}

func (a *GreengrassActivities) StartBulkDeployment(input *greengrass.StartBulkDeploymentInput) (*greengrass.StartBulkDeploymentOutput, error) {
    return a.client.StartBulkDeployment(input)
}

func (a *GreengrassActivities) StopBulkDeployment(input *greengrass.StopBulkDeploymentInput) (*greengrass.StopBulkDeploymentOutput, error) {
    return a.client.StopBulkDeployment(input)
}

func (a *GreengrassActivities) TagResource(input *greengrass.TagResourceInput) (*greengrass.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *GreengrassActivities) UntagResource(input *greengrass.UntagResourceInput) (*greengrass.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *GreengrassActivities) UpdateConnectivityInfo(input *greengrass.UpdateConnectivityInfoInput) (*greengrass.UpdateConnectivityInfoOutput, error) {
    return a.client.UpdateConnectivityInfo(input)
}

func (a *GreengrassActivities) UpdateConnectorDefinition(input *greengrass.UpdateConnectorDefinitionInput) (*greengrass.UpdateConnectorDefinitionOutput, error) {
    return a.client.UpdateConnectorDefinition(input)
}

func (a *GreengrassActivities) UpdateCoreDefinition(input *greengrass.UpdateCoreDefinitionInput) (*greengrass.UpdateCoreDefinitionOutput, error) {
    return a.client.UpdateCoreDefinition(input)
}

func (a *GreengrassActivities) UpdateDeviceDefinition(input *greengrass.UpdateDeviceDefinitionInput) (*greengrass.UpdateDeviceDefinitionOutput, error) {
    return a.client.UpdateDeviceDefinition(input)
}

func (a *GreengrassActivities) UpdateFunctionDefinition(input *greengrass.UpdateFunctionDefinitionInput) (*greengrass.UpdateFunctionDefinitionOutput, error) {
    return a.client.UpdateFunctionDefinition(input)
}

func (a *GreengrassActivities) UpdateGroup(input *greengrass.UpdateGroupInput) (*greengrass.UpdateGroupOutput, error) {
    return a.client.UpdateGroup(input)
}

func (a *GreengrassActivities) UpdateGroupCertificateConfiguration(input *greengrass.UpdateGroupCertificateConfigurationInput) (*greengrass.UpdateGroupCertificateConfigurationOutput, error) {
    return a.client.UpdateGroupCertificateConfiguration(input)
}

func (a *GreengrassActivities) UpdateLoggerDefinition(input *greengrass.UpdateLoggerDefinitionInput) (*greengrass.UpdateLoggerDefinitionOutput, error) {
    return a.client.UpdateLoggerDefinition(input)
}

func (a *GreengrassActivities) UpdateResourceDefinition(input *greengrass.UpdateResourceDefinitionInput) (*greengrass.UpdateResourceDefinitionOutput, error) {
    return a.client.UpdateResourceDefinition(input)
}

func (a *GreengrassActivities) UpdateSubscriptionDefinition(input *greengrass.UpdateSubscriptionDefinitionInput) (*greengrass.UpdateSubscriptionDefinitionOutput, error) {
    return a.client.UpdateSubscriptionDefinition(input)
}
