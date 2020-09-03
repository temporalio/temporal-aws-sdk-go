package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/greengrass"
	"go.temporal.io/sdk/workflow"
)

type GreengrassClient interface {
    AssociateRoleToGroup(ctx workflow.Context, input *greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error)
    AssociateRoleToGroupAsync(ctx workflow.Context, input *greengrass.AssociateRoleToGroupInput) *GreengrassAssociateRoleToGroupResult

    AssociateServiceRoleToAccount(ctx workflow.Context, input *greengrass.AssociateServiceRoleToAccountInput) (*greengrass.AssociateServiceRoleToAccountOutput, error)
    AssociateServiceRoleToAccountAsync(ctx workflow.Context, input *greengrass.AssociateServiceRoleToAccountInput) *GreengrassAssociateServiceRoleToAccountResult

    CreateConnectorDefinition(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionInput) (*greengrass.CreateConnectorDefinitionOutput, error)
    CreateConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionInput) *GreengrassCreateConnectorDefinitionResult

    CreateConnectorDefinitionVersion(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionVersionInput) (*greengrass.CreateConnectorDefinitionVersionOutput, error)
    CreateConnectorDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionVersionInput) *GreengrassCreateConnectorDefinitionVersionResult

    CreateCoreDefinition(ctx workflow.Context, input *greengrass.CreateCoreDefinitionInput) (*greengrass.CreateCoreDefinitionOutput, error)
    CreateCoreDefinitionAsync(ctx workflow.Context, input *greengrass.CreateCoreDefinitionInput) *GreengrassCreateCoreDefinitionResult

    CreateCoreDefinitionVersion(ctx workflow.Context, input *greengrass.CreateCoreDefinitionVersionInput) (*greengrass.CreateCoreDefinitionVersionOutput, error)
    CreateCoreDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateCoreDefinitionVersionInput) *GreengrassCreateCoreDefinitionVersionResult

    CreateDeployment(ctx workflow.Context, input *greengrass.CreateDeploymentInput) (*greengrass.CreateDeploymentOutput, error)
    CreateDeploymentAsync(ctx workflow.Context, input *greengrass.CreateDeploymentInput) *GreengrassCreateDeploymentResult

    CreateDeviceDefinition(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionInput) (*greengrass.CreateDeviceDefinitionOutput, error)
    CreateDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionInput) *GreengrassCreateDeviceDefinitionResult

    CreateDeviceDefinitionVersion(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionVersionInput) (*greengrass.CreateDeviceDefinitionVersionOutput, error)
    CreateDeviceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionVersionInput) *GreengrassCreateDeviceDefinitionVersionResult

    CreateFunctionDefinition(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionInput) (*greengrass.CreateFunctionDefinitionOutput, error)
    CreateFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionInput) *GreengrassCreateFunctionDefinitionResult

    CreateFunctionDefinitionVersion(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionVersionInput) (*greengrass.CreateFunctionDefinitionVersionOutput, error)
    CreateFunctionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionVersionInput) *GreengrassCreateFunctionDefinitionVersionResult

    CreateGroup(ctx workflow.Context, input *greengrass.CreateGroupInput) (*greengrass.CreateGroupOutput, error)
    CreateGroupAsync(ctx workflow.Context, input *greengrass.CreateGroupInput) *GreengrassCreateGroupResult

    CreateGroupCertificateAuthority(ctx workflow.Context, input *greengrass.CreateGroupCertificateAuthorityInput) (*greengrass.CreateGroupCertificateAuthorityOutput, error)
    CreateGroupCertificateAuthorityAsync(ctx workflow.Context, input *greengrass.CreateGroupCertificateAuthorityInput) *GreengrassCreateGroupCertificateAuthorityResult

    CreateGroupVersion(ctx workflow.Context, input *greengrass.CreateGroupVersionInput) (*greengrass.CreateGroupVersionOutput, error)
    CreateGroupVersionAsync(ctx workflow.Context, input *greengrass.CreateGroupVersionInput) *GreengrassCreateGroupVersionResult

    CreateLoggerDefinition(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionInput) (*greengrass.CreateLoggerDefinitionOutput, error)
    CreateLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionInput) *GreengrassCreateLoggerDefinitionResult

    CreateLoggerDefinitionVersion(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionVersionInput) (*greengrass.CreateLoggerDefinitionVersionOutput, error)
    CreateLoggerDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionVersionInput) *GreengrassCreateLoggerDefinitionVersionResult

    CreateResourceDefinition(ctx workflow.Context, input *greengrass.CreateResourceDefinitionInput) (*greengrass.CreateResourceDefinitionOutput, error)
    CreateResourceDefinitionAsync(ctx workflow.Context, input *greengrass.CreateResourceDefinitionInput) *GreengrassCreateResourceDefinitionResult

    CreateResourceDefinitionVersion(ctx workflow.Context, input *greengrass.CreateResourceDefinitionVersionInput) (*greengrass.CreateResourceDefinitionVersionOutput, error)
    CreateResourceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateResourceDefinitionVersionInput) *GreengrassCreateResourceDefinitionVersionResult

    CreateSoftwareUpdateJob(ctx workflow.Context, input *greengrass.CreateSoftwareUpdateJobInput) (*greengrass.CreateSoftwareUpdateJobOutput, error)
    CreateSoftwareUpdateJobAsync(ctx workflow.Context, input *greengrass.CreateSoftwareUpdateJobInput) *GreengrassCreateSoftwareUpdateJobResult

    CreateSubscriptionDefinition(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionInput) (*greengrass.CreateSubscriptionDefinitionOutput, error)
    CreateSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionInput) *GreengrassCreateSubscriptionDefinitionResult

    CreateSubscriptionDefinitionVersion(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionVersionInput) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error)
    CreateSubscriptionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionVersionInput) *GreengrassCreateSubscriptionDefinitionVersionResult

    DeleteConnectorDefinition(ctx workflow.Context, input *greengrass.DeleteConnectorDefinitionInput) (*greengrass.DeleteConnectorDefinitionOutput, error)
    DeleteConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteConnectorDefinitionInput) *GreengrassDeleteConnectorDefinitionResult

    DeleteCoreDefinition(ctx workflow.Context, input *greengrass.DeleteCoreDefinitionInput) (*greengrass.DeleteCoreDefinitionOutput, error)
    DeleteCoreDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteCoreDefinitionInput) *GreengrassDeleteCoreDefinitionResult

    DeleteDeviceDefinition(ctx workflow.Context, input *greengrass.DeleteDeviceDefinitionInput) (*greengrass.DeleteDeviceDefinitionOutput, error)
    DeleteDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteDeviceDefinitionInput) *GreengrassDeleteDeviceDefinitionResult

    DeleteFunctionDefinition(ctx workflow.Context, input *greengrass.DeleteFunctionDefinitionInput) (*greengrass.DeleteFunctionDefinitionOutput, error)
    DeleteFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteFunctionDefinitionInput) *GreengrassDeleteFunctionDefinitionResult

    DeleteGroup(ctx workflow.Context, input *greengrass.DeleteGroupInput) (*greengrass.DeleteGroupOutput, error)
    DeleteGroupAsync(ctx workflow.Context, input *greengrass.DeleteGroupInput) *GreengrassDeleteGroupResult

    DeleteLoggerDefinition(ctx workflow.Context, input *greengrass.DeleteLoggerDefinitionInput) (*greengrass.DeleteLoggerDefinitionOutput, error)
    DeleteLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteLoggerDefinitionInput) *GreengrassDeleteLoggerDefinitionResult

    DeleteResourceDefinition(ctx workflow.Context, input *greengrass.DeleteResourceDefinitionInput) (*greengrass.DeleteResourceDefinitionOutput, error)
    DeleteResourceDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteResourceDefinitionInput) *GreengrassDeleteResourceDefinitionResult

    DeleteSubscriptionDefinition(ctx workflow.Context, input *greengrass.DeleteSubscriptionDefinitionInput) (*greengrass.DeleteSubscriptionDefinitionOutput, error)
    DeleteSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.DeleteSubscriptionDefinitionInput) *GreengrassDeleteSubscriptionDefinitionResult

    DisassociateRoleFromGroup(ctx workflow.Context, input *greengrass.DisassociateRoleFromGroupInput) (*greengrass.DisassociateRoleFromGroupOutput, error)
    DisassociateRoleFromGroupAsync(ctx workflow.Context, input *greengrass.DisassociateRoleFromGroupInput) *GreengrassDisassociateRoleFromGroupResult

    DisassociateServiceRoleFromAccount(ctx workflow.Context, input *greengrass.DisassociateServiceRoleFromAccountInput) (*greengrass.DisassociateServiceRoleFromAccountOutput, error)
    DisassociateServiceRoleFromAccountAsync(ctx workflow.Context, input *greengrass.DisassociateServiceRoleFromAccountInput) *GreengrassDisassociateServiceRoleFromAccountResult

    GetAssociatedRole(ctx workflow.Context, input *greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error)
    GetAssociatedRoleAsync(ctx workflow.Context, input *greengrass.GetAssociatedRoleInput) *GreengrassGetAssociatedRoleResult

    GetBulkDeploymentStatus(ctx workflow.Context, input *greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error)
    GetBulkDeploymentStatusAsync(ctx workflow.Context, input *greengrass.GetBulkDeploymentStatusInput) *GreengrassGetBulkDeploymentStatusResult

    GetConnectivityInfo(ctx workflow.Context, input *greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error)
    GetConnectivityInfoAsync(ctx workflow.Context, input *greengrass.GetConnectivityInfoInput) *GreengrassGetConnectivityInfoResult

    GetConnectorDefinition(ctx workflow.Context, input *greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error)
    GetConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.GetConnectorDefinitionInput) *GreengrassGetConnectorDefinitionResult

    GetConnectorDefinitionVersion(ctx workflow.Context, input *greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error)
    GetConnectorDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetConnectorDefinitionVersionInput) *GreengrassGetConnectorDefinitionVersionResult

    GetCoreDefinition(ctx workflow.Context, input *greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error)
    GetCoreDefinitionAsync(ctx workflow.Context, input *greengrass.GetCoreDefinitionInput) *GreengrassGetCoreDefinitionResult

    GetCoreDefinitionVersion(ctx workflow.Context, input *greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error)
    GetCoreDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetCoreDefinitionVersionInput) *GreengrassGetCoreDefinitionVersionResult

    GetDeploymentStatus(ctx workflow.Context, input *greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error)
    GetDeploymentStatusAsync(ctx workflow.Context, input *greengrass.GetDeploymentStatusInput) *GreengrassGetDeploymentStatusResult

    GetDeviceDefinition(ctx workflow.Context, input *greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error)
    GetDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.GetDeviceDefinitionInput) *GreengrassGetDeviceDefinitionResult

    GetDeviceDefinitionVersion(ctx workflow.Context, input *greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error)
    GetDeviceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetDeviceDefinitionVersionInput) *GreengrassGetDeviceDefinitionVersionResult

    GetFunctionDefinition(ctx workflow.Context, input *greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error)
    GetFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.GetFunctionDefinitionInput) *GreengrassGetFunctionDefinitionResult

    GetFunctionDefinitionVersion(ctx workflow.Context, input *greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error)
    GetFunctionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetFunctionDefinitionVersionInput) *GreengrassGetFunctionDefinitionVersionResult

    GetGroup(ctx workflow.Context, input *greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error)
    GetGroupAsync(ctx workflow.Context, input *greengrass.GetGroupInput) *GreengrassGetGroupResult

    GetGroupCertificateAuthority(ctx workflow.Context, input *greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error)
    GetGroupCertificateAuthorityAsync(ctx workflow.Context, input *greengrass.GetGroupCertificateAuthorityInput) *GreengrassGetGroupCertificateAuthorityResult

    GetGroupCertificateConfiguration(ctx workflow.Context, input *greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error)
    GetGroupCertificateConfigurationAsync(ctx workflow.Context, input *greengrass.GetGroupCertificateConfigurationInput) *GreengrassGetGroupCertificateConfigurationResult

    GetGroupVersion(ctx workflow.Context, input *greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error)
    GetGroupVersionAsync(ctx workflow.Context, input *greengrass.GetGroupVersionInput) *GreengrassGetGroupVersionResult

    GetLoggerDefinition(ctx workflow.Context, input *greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error)
    GetLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.GetLoggerDefinitionInput) *GreengrassGetLoggerDefinitionResult

    GetLoggerDefinitionVersion(ctx workflow.Context, input *greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error)
    GetLoggerDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetLoggerDefinitionVersionInput) *GreengrassGetLoggerDefinitionVersionResult

    GetResourceDefinition(ctx workflow.Context, input *greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error)
    GetResourceDefinitionAsync(ctx workflow.Context, input *greengrass.GetResourceDefinitionInput) *GreengrassGetResourceDefinitionResult

    GetResourceDefinitionVersion(ctx workflow.Context, input *greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error)
    GetResourceDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetResourceDefinitionVersionInput) *GreengrassGetResourceDefinitionVersionResult

    GetServiceRoleForAccount(ctx workflow.Context, input *greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error)
    GetServiceRoleForAccountAsync(ctx workflow.Context, input *greengrass.GetServiceRoleForAccountInput) *GreengrassGetServiceRoleForAccountResult

    GetSubscriptionDefinition(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error)
    GetSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionInput) *GreengrassGetSubscriptionDefinitionResult

    GetSubscriptionDefinitionVersion(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error)
    GetSubscriptionDefinitionVersionAsync(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionVersionInput) *GreengrassGetSubscriptionDefinitionVersionResult

    ListBulkDeploymentDetailedReports(ctx workflow.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error)
    ListBulkDeploymentDetailedReportsAsync(ctx workflow.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput) *GreengrassListBulkDeploymentDetailedReportsResult

    ListBulkDeployments(ctx workflow.Context, input *greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error)
    ListBulkDeploymentsAsync(ctx workflow.Context, input *greengrass.ListBulkDeploymentsInput) *GreengrassListBulkDeploymentsResult

    ListConnectorDefinitionVersions(ctx workflow.Context, input *greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error)
    ListConnectorDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListConnectorDefinitionVersionsInput) *GreengrassListConnectorDefinitionVersionsResult

    ListConnectorDefinitions(ctx workflow.Context, input *greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error)
    ListConnectorDefinitionsAsync(ctx workflow.Context, input *greengrass.ListConnectorDefinitionsInput) *GreengrassListConnectorDefinitionsResult

    ListCoreDefinitionVersions(ctx workflow.Context, input *greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error)
    ListCoreDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListCoreDefinitionVersionsInput) *GreengrassListCoreDefinitionVersionsResult

    ListCoreDefinitions(ctx workflow.Context, input *greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error)
    ListCoreDefinitionsAsync(ctx workflow.Context, input *greengrass.ListCoreDefinitionsInput) *GreengrassListCoreDefinitionsResult

    ListDeployments(ctx workflow.Context, input *greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error)
    ListDeploymentsAsync(ctx workflow.Context, input *greengrass.ListDeploymentsInput) *GreengrassListDeploymentsResult

    ListDeviceDefinitionVersions(ctx workflow.Context, input *greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error)
    ListDeviceDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListDeviceDefinitionVersionsInput) *GreengrassListDeviceDefinitionVersionsResult

    ListDeviceDefinitions(ctx workflow.Context, input *greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error)
    ListDeviceDefinitionsAsync(ctx workflow.Context, input *greengrass.ListDeviceDefinitionsInput) *GreengrassListDeviceDefinitionsResult

    ListFunctionDefinitionVersions(ctx workflow.Context, input *greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error)
    ListFunctionDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListFunctionDefinitionVersionsInput) *GreengrassListFunctionDefinitionVersionsResult

    ListFunctionDefinitions(ctx workflow.Context, input *greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error)
    ListFunctionDefinitionsAsync(ctx workflow.Context, input *greengrass.ListFunctionDefinitionsInput) *GreengrassListFunctionDefinitionsResult

    ListGroupCertificateAuthorities(ctx workflow.Context, input *greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error)
    ListGroupCertificateAuthoritiesAsync(ctx workflow.Context, input *greengrass.ListGroupCertificateAuthoritiesInput) *GreengrassListGroupCertificateAuthoritiesResult

    ListGroupVersions(ctx workflow.Context, input *greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error)
    ListGroupVersionsAsync(ctx workflow.Context, input *greengrass.ListGroupVersionsInput) *GreengrassListGroupVersionsResult

    ListGroups(ctx workflow.Context, input *greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error)
    ListGroupsAsync(ctx workflow.Context, input *greengrass.ListGroupsInput) *GreengrassListGroupsResult

    ListLoggerDefinitionVersions(ctx workflow.Context, input *greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error)
    ListLoggerDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListLoggerDefinitionVersionsInput) *GreengrassListLoggerDefinitionVersionsResult

    ListLoggerDefinitions(ctx workflow.Context, input *greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error)
    ListLoggerDefinitionsAsync(ctx workflow.Context, input *greengrass.ListLoggerDefinitionsInput) *GreengrassListLoggerDefinitionsResult

    ListResourceDefinitionVersions(ctx workflow.Context, input *greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error)
    ListResourceDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListResourceDefinitionVersionsInput) *GreengrassListResourceDefinitionVersionsResult

    ListResourceDefinitions(ctx workflow.Context, input *greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error)
    ListResourceDefinitionsAsync(ctx workflow.Context, input *greengrass.ListResourceDefinitionsInput) *GreengrassListResourceDefinitionsResult

    ListSubscriptionDefinitionVersions(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error)
    ListSubscriptionDefinitionVersionsAsync(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput) *GreengrassListSubscriptionDefinitionVersionsResult

    ListSubscriptionDefinitions(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error)
    ListSubscriptionDefinitionsAsync(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionsInput) *GreengrassListSubscriptionDefinitionsResult

    ListTagsForResource(ctx workflow.Context, input *greengrass.ListTagsForResourceInput) (*greengrass.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *greengrass.ListTagsForResourceInput) *GreengrassListTagsForResourceResult

    ResetDeployments(ctx workflow.Context, input *greengrass.ResetDeploymentsInput) (*greengrass.ResetDeploymentsOutput, error)
    ResetDeploymentsAsync(ctx workflow.Context, input *greengrass.ResetDeploymentsInput) *GreengrassResetDeploymentsResult

    StartBulkDeployment(ctx workflow.Context, input *greengrass.StartBulkDeploymentInput) (*greengrass.StartBulkDeploymentOutput, error)
    StartBulkDeploymentAsync(ctx workflow.Context, input *greengrass.StartBulkDeploymentInput) *GreengrassStartBulkDeploymentResult

    StopBulkDeployment(ctx workflow.Context, input *greengrass.StopBulkDeploymentInput) (*greengrass.StopBulkDeploymentOutput, error)
    StopBulkDeploymentAsync(ctx workflow.Context, input *greengrass.StopBulkDeploymentInput) *GreengrassStopBulkDeploymentResult

    TagResource(ctx workflow.Context, input *greengrass.TagResourceInput) (*greengrass.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *greengrass.TagResourceInput) *GreengrassTagResourceResult

    UntagResource(ctx workflow.Context, input *greengrass.UntagResourceInput) (*greengrass.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *greengrass.UntagResourceInput) *GreengrassUntagResourceResult

    UpdateConnectivityInfo(ctx workflow.Context, input *greengrass.UpdateConnectivityInfoInput) (*greengrass.UpdateConnectivityInfoOutput, error)
    UpdateConnectivityInfoAsync(ctx workflow.Context, input *greengrass.UpdateConnectivityInfoInput) *GreengrassUpdateConnectivityInfoResult

    UpdateConnectorDefinition(ctx workflow.Context, input *greengrass.UpdateConnectorDefinitionInput) (*greengrass.UpdateConnectorDefinitionOutput, error)
    UpdateConnectorDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateConnectorDefinitionInput) *GreengrassUpdateConnectorDefinitionResult

    UpdateCoreDefinition(ctx workflow.Context, input *greengrass.UpdateCoreDefinitionInput) (*greengrass.UpdateCoreDefinitionOutput, error)
    UpdateCoreDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateCoreDefinitionInput) *GreengrassUpdateCoreDefinitionResult

    UpdateDeviceDefinition(ctx workflow.Context, input *greengrass.UpdateDeviceDefinitionInput) (*greengrass.UpdateDeviceDefinitionOutput, error)
    UpdateDeviceDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateDeviceDefinitionInput) *GreengrassUpdateDeviceDefinitionResult

    UpdateFunctionDefinition(ctx workflow.Context, input *greengrass.UpdateFunctionDefinitionInput) (*greengrass.UpdateFunctionDefinitionOutput, error)
    UpdateFunctionDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateFunctionDefinitionInput) *GreengrassUpdateFunctionDefinitionResult

    UpdateGroup(ctx workflow.Context, input *greengrass.UpdateGroupInput) (*greengrass.UpdateGroupOutput, error)
    UpdateGroupAsync(ctx workflow.Context, input *greengrass.UpdateGroupInput) *GreengrassUpdateGroupResult

    UpdateGroupCertificateConfiguration(ctx workflow.Context, input *greengrass.UpdateGroupCertificateConfigurationInput) (*greengrass.UpdateGroupCertificateConfigurationOutput, error)
    UpdateGroupCertificateConfigurationAsync(ctx workflow.Context, input *greengrass.UpdateGroupCertificateConfigurationInput) *GreengrassUpdateGroupCertificateConfigurationResult

    UpdateLoggerDefinition(ctx workflow.Context, input *greengrass.UpdateLoggerDefinitionInput) (*greengrass.UpdateLoggerDefinitionOutput, error)
    UpdateLoggerDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateLoggerDefinitionInput) *GreengrassUpdateLoggerDefinitionResult

    UpdateResourceDefinition(ctx workflow.Context, input *greengrass.UpdateResourceDefinitionInput) (*greengrass.UpdateResourceDefinitionOutput, error)
    UpdateResourceDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateResourceDefinitionInput) *GreengrassUpdateResourceDefinitionResult

    UpdateSubscriptionDefinition(ctx workflow.Context, input *greengrass.UpdateSubscriptionDefinitionInput) (*greengrass.UpdateSubscriptionDefinitionOutput, error)
    UpdateSubscriptionDefinitionAsync(ctx workflow.Context, input *greengrass.UpdateSubscriptionDefinitionInput) *GreengrassUpdateSubscriptionDefinitionResult
}
type GreengrassAssociateRoleToGroupResult struct {
	Result workflow.Future
}

func (r *GreengrassAssociateRoleToGroupResult) Get(ctx workflow.Context) (*greengrass.AssociateRoleToGroupOutput, error) {
    var output greengrass.AssociateRoleToGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassAssociateServiceRoleToAccountResult struct {
	Result workflow.Future
}

func (r *GreengrassAssociateServiceRoleToAccountResult) Get(ctx workflow.Context) (*greengrass.AssociateServiceRoleToAccountOutput, error) {
    var output greengrass.AssociateServiceRoleToAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateConnectorDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateConnectorDefinitionResult) Get(ctx workflow.Context) (*greengrass.CreateConnectorDefinitionOutput, error) {
    var output greengrass.CreateConnectorDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateConnectorDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateConnectorDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.CreateConnectorDefinitionVersionOutput, error) {
    var output greengrass.CreateConnectorDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateCoreDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateCoreDefinitionResult) Get(ctx workflow.Context) (*greengrass.CreateCoreDefinitionOutput, error) {
    var output greengrass.CreateCoreDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateCoreDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateCoreDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.CreateCoreDefinitionVersionOutput, error) {
    var output greengrass.CreateCoreDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateDeploymentResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateDeploymentResult) Get(ctx workflow.Context) (*greengrass.CreateDeploymentOutput, error) {
    var output greengrass.CreateDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateDeviceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateDeviceDefinitionResult) Get(ctx workflow.Context) (*greengrass.CreateDeviceDefinitionOutput, error) {
    var output greengrass.CreateDeviceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateDeviceDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateDeviceDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.CreateDeviceDefinitionVersionOutput, error) {
    var output greengrass.CreateDeviceDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateFunctionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateFunctionDefinitionResult) Get(ctx workflow.Context) (*greengrass.CreateFunctionDefinitionOutput, error) {
    var output greengrass.CreateFunctionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateFunctionDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateFunctionDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.CreateFunctionDefinitionVersionOutput, error) {
    var output greengrass.CreateFunctionDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateGroupResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateGroupResult) Get(ctx workflow.Context) (*greengrass.CreateGroupOutput, error) {
    var output greengrass.CreateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateGroupCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateGroupCertificateAuthorityResult) Get(ctx workflow.Context) (*greengrass.CreateGroupCertificateAuthorityOutput, error) {
    var output greengrass.CreateGroupCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateGroupVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateGroupVersionResult) Get(ctx workflow.Context) (*greengrass.CreateGroupVersionOutput, error) {
    var output greengrass.CreateGroupVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateLoggerDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateLoggerDefinitionResult) Get(ctx workflow.Context) (*greengrass.CreateLoggerDefinitionOutput, error) {
    var output greengrass.CreateLoggerDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateLoggerDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateLoggerDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.CreateLoggerDefinitionVersionOutput, error) {
    var output greengrass.CreateLoggerDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateResourceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateResourceDefinitionResult) Get(ctx workflow.Context) (*greengrass.CreateResourceDefinitionOutput, error) {
    var output greengrass.CreateResourceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateResourceDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateResourceDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.CreateResourceDefinitionVersionOutput, error) {
    var output greengrass.CreateResourceDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateSoftwareUpdateJobResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateSoftwareUpdateJobResult) Get(ctx workflow.Context) (*greengrass.CreateSoftwareUpdateJobOutput, error) {
    var output greengrass.CreateSoftwareUpdateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateSubscriptionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateSubscriptionDefinitionResult) Get(ctx workflow.Context) (*greengrass.CreateSubscriptionDefinitionOutput, error) {
    var output greengrass.CreateSubscriptionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassCreateSubscriptionDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassCreateSubscriptionDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error) {
    var output greengrass.CreateSubscriptionDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteConnectorDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteConnectorDefinitionResult) Get(ctx workflow.Context) (*greengrass.DeleteConnectorDefinitionOutput, error) {
    var output greengrass.DeleteConnectorDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteCoreDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteCoreDefinitionResult) Get(ctx workflow.Context) (*greengrass.DeleteCoreDefinitionOutput, error) {
    var output greengrass.DeleteCoreDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteDeviceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteDeviceDefinitionResult) Get(ctx workflow.Context) (*greengrass.DeleteDeviceDefinitionOutput, error) {
    var output greengrass.DeleteDeviceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteFunctionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteFunctionDefinitionResult) Get(ctx workflow.Context) (*greengrass.DeleteFunctionDefinitionOutput, error) {
    var output greengrass.DeleteFunctionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteGroupResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteGroupResult) Get(ctx workflow.Context) (*greengrass.DeleteGroupOutput, error) {
    var output greengrass.DeleteGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteLoggerDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteLoggerDefinitionResult) Get(ctx workflow.Context) (*greengrass.DeleteLoggerDefinitionOutput, error) {
    var output greengrass.DeleteLoggerDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteResourceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteResourceDefinitionResult) Get(ctx workflow.Context) (*greengrass.DeleteResourceDefinitionOutput, error) {
    var output greengrass.DeleteResourceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDeleteSubscriptionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassDeleteSubscriptionDefinitionResult) Get(ctx workflow.Context) (*greengrass.DeleteSubscriptionDefinitionOutput, error) {
    var output greengrass.DeleteSubscriptionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDisassociateRoleFromGroupResult struct {
	Result workflow.Future
}

func (r *GreengrassDisassociateRoleFromGroupResult) Get(ctx workflow.Context) (*greengrass.DisassociateRoleFromGroupOutput, error) {
    var output greengrass.DisassociateRoleFromGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassDisassociateServiceRoleFromAccountResult struct {
	Result workflow.Future
}

func (r *GreengrassDisassociateServiceRoleFromAccountResult) Get(ctx workflow.Context) (*greengrass.DisassociateServiceRoleFromAccountOutput, error) {
    var output greengrass.DisassociateServiceRoleFromAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetAssociatedRoleResult struct {
	Result workflow.Future
}

func (r *GreengrassGetAssociatedRoleResult) Get(ctx workflow.Context) (*greengrass.GetAssociatedRoleOutput, error) {
    var output greengrass.GetAssociatedRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetBulkDeploymentStatusResult struct {
	Result workflow.Future
}

func (r *GreengrassGetBulkDeploymentStatusResult) Get(ctx workflow.Context) (*greengrass.GetBulkDeploymentStatusOutput, error) {
    var output greengrass.GetBulkDeploymentStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetConnectivityInfoResult struct {
	Result workflow.Future
}

func (r *GreengrassGetConnectivityInfoResult) Get(ctx workflow.Context) (*greengrass.GetConnectivityInfoOutput, error) {
    var output greengrass.GetConnectivityInfoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetConnectorDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetConnectorDefinitionResult) Get(ctx workflow.Context) (*greengrass.GetConnectorDefinitionOutput, error) {
    var output greengrass.GetConnectorDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetConnectorDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetConnectorDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.GetConnectorDefinitionVersionOutput, error) {
    var output greengrass.GetConnectorDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetCoreDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetCoreDefinitionResult) Get(ctx workflow.Context) (*greengrass.GetCoreDefinitionOutput, error) {
    var output greengrass.GetCoreDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetCoreDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetCoreDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.GetCoreDefinitionVersionOutput, error) {
    var output greengrass.GetCoreDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetDeploymentStatusResult struct {
	Result workflow.Future
}

func (r *GreengrassGetDeploymentStatusResult) Get(ctx workflow.Context) (*greengrass.GetDeploymentStatusOutput, error) {
    var output greengrass.GetDeploymentStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetDeviceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetDeviceDefinitionResult) Get(ctx workflow.Context) (*greengrass.GetDeviceDefinitionOutput, error) {
    var output greengrass.GetDeviceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetDeviceDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetDeviceDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.GetDeviceDefinitionVersionOutput, error) {
    var output greengrass.GetDeviceDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetFunctionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetFunctionDefinitionResult) Get(ctx workflow.Context) (*greengrass.GetFunctionDefinitionOutput, error) {
    var output greengrass.GetFunctionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetFunctionDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetFunctionDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.GetFunctionDefinitionVersionOutput, error) {
    var output greengrass.GetFunctionDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetGroupResult struct {
	Result workflow.Future
}

func (r *GreengrassGetGroupResult) Get(ctx workflow.Context) (*greengrass.GetGroupOutput, error) {
    var output greengrass.GetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetGroupCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *GreengrassGetGroupCertificateAuthorityResult) Get(ctx workflow.Context) (*greengrass.GetGroupCertificateAuthorityOutput, error) {
    var output greengrass.GetGroupCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetGroupCertificateConfigurationResult struct {
	Result workflow.Future
}

func (r *GreengrassGetGroupCertificateConfigurationResult) Get(ctx workflow.Context) (*greengrass.GetGroupCertificateConfigurationOutput, error) {
    var output greengrass.GetGroupCertificateConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetGroupVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetGroupVersionResult) Get(ctx workflow.Context) (*greengrass.GetGroupVersionOutput, error) {
    var output greengrass.GetGroupVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetLoggerDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetLoggerDefinitionResult) Get(ctx workflow.Context) (*greengrass.GetLoggerDefinitionOutput, error) {
    var output greengrass.GetLoggerDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetLoggerDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetLoggerDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.GetLoggerDefinitionVersionOutput, error) {
    var output greengrass.GetLoggerDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetResourceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetResourceDefinitionResult) Get(ctx workflow.Context) (*greengrass.GetResourceDefinitionOutput, error) {
    var output greengrass.GetResourceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetResourceDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetResourceDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.GetResourceDefinitionVersionOutput, error) {
    var output greengrass.GetResourceDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetServiceRoleForAccountResult struct {
	Result workflow.Future
}

func (r *GreengrassGetServiceRoleForAccountResult) Get(ctx workflow.Context) (*greengrass.GetServiceRoleForAccountOutput, error) {
    var output greengrass.GetServiceRoleForAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetSubscriptionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetSubscriptionDefinitionResult) Get(ctx workflow.Context) (*greengrass.GetSubscriptionDefinitionOutput, error) {
    var output greengrass.GetSubscriptionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassGetSubscriptionDefinitionVersionResult struct {
	Result workflow.Future
}

func (r *GreengrassGetSubscriptionDefinitionVersionResult) Get(ctx workflow.Context) (*greengrass.GetSubscriptionDefinitionVersionOutput, error) {
    var output greengrass.GetSubscriptionDefinitionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListBulkDeploymentDetailedReportsResult struct {
	Result workflow.Future
}

func (r *GreengrassListBulkDeploymentDetailedReportsResult) Get(ctx workflow.Context) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error) {
    var output greengrass.ListBulkDeploymentDetailedReportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListBulkDeploymentsResult struct {
	Result workflow.Future
}

func (r *GreengrassListBulkDeploymentsResult) Get(ctx workflow.Context) (*greengrass.ListBulkDeploymentsOutput, error) {
    var output greengrass.ListBulkDeploymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListConnectorDefinitionVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListConnectorDefinitionVersionsResult) Get(ctx workflow.Context) (*greengrass.ListConnectorDefinitionVersionsOutput, error) {
    var output greengrass.ListConnectorDefinitionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListConnectorDefinitionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListConnectorDefinitionsResult) Get(ctx workflow.Context) (*greengrass.ListConnectorDefinitionsOutput, error) {
    var output greengrass.ListConnectorDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListCoreDefinitionVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListCoreDefinitionVersionsResult) Get(ctx workflow.Context) (*greengrass.ListCoreDefinitionVersionsOutput, error) {
    var output greengrass.ListCoreDefinitionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListCoreDefinitionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListCoreDefinitionsResult) Get(ctx workflow.Context) (*greengrass.ListCoreDefinitionsOutput, error) {
    var output greengrass.ListCoreDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListDeploymentsResult struct {
	Result workflow.Future
}

func (r *GreengrassListDeploymentsResult) Get(ctx workflow.Context) (*greengrass.ListDeploymentsOutput, error) {
    var output greengrass.ListDeploymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListDeviceDefinitionVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListDeviceDefinitionVersionsResult) Get(ctx workflow.Context) (*greengrass.ListDeviceDefinitionVersionsOutput, error) {
    var output greengrass.ListDeviceDefinitionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListDeviceDefinitionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListDeviceDefinitionsResult) Get(ctx workflow.Context) (*greengrass.ListDeviceDefinitionsOutput, error) {
    var output greengrass.ListDeviceDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListFunctionDefinitionVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListFunctionDefinitionVersionsResult) Get(ctx workflow.Context) (*greengrass.ListFunctionDefinitionVersionsOutput, error) {
    var output greengrass.ListFunctionDefinitionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListFunctionDefinitionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListFunctionDefinitionsResult) Get(ctx workflow.Context) (*greengrass.ListFunctionDefinitionsOutput, error) {
    var output greengrass.ListFunctionDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListGroupCertificateAuthoritiesResult struct {
	Result workflow.Future
}

func (r *GreengrassListGroupCertificateAuthoritiesResult) Get(ctx workflow.Context) (*greengrass.ListGroupCertificateAuthoritiesOutput, error) {
    var output greengrass.ListGroupCertificateAuthoritiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListGroupVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListGroupVersionsResult) Get(ctx workflow.Context) (*greengrass.ListGroupVersionsOutput, error) {
    var output greengrass.ListGroupVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListGroupsResult struct {
	Result workflow.Future
}

func (r *GreengrassListGroupsResult) Get(ctx workflow.Context) (*greengrass.ListGroupsOutput, error) {
    var output greengrass.ListGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListLoggerDefinitionVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListLoggerDefinitionVersionsResult) Get(ctx workflow.Context) (*greengrass.ListLoggerDefinitionVersionsOutput, error) {
    var output greengrass.ListLoggerDefinitionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListLoggerDefinitionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListLoggerDefinitionsResult) Get(ctx workflow.Context) (*greengrass.ListLoggerDefinitionsOutput, error) {
    var output greengrass.ListLoggerDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListResourceDefinitionVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListResourceDefinitionVersionsResult) Get(ctx workflow.Context) (*greengrass.ListResourceDefinitionVersionsOutput, error) {
    var output greengrass.ListResourceDefinitionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListResourceDefinitionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListResourceDefinitionsResult) Get(ctx workflow.Context) (*greengrass.ListResourceDefinitionsOutput, error) {
    var output greengrass.ListResourceDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListSubscriptionDefinitionVersionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListSubscriptionDefinitionVersionsResult) Get(ctx workflow.Context) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error) {
    var output greengrass.ListSubscriptionDefinitionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListSubscriptionDefinitionsResult struct {
	Result workflow.Future
}

func (r *GreengrassListSubscriptionDefinitionsResult) Get(ctx workflow.Context) (*greengrass.ListSubscriptionDefinitionsOutput, error) {
    var output greengrass.ListSubscriptionDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *GreengrassListTagsForResourceResult) Get(ctx workflow.Context) (*greengrass.ListTagsForResourceOutput, error) {
    var output greengrass.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassResetDeploymentsResult struct {
	Result workflow.Future
}

func (r *GreengrassResetDeploymentsResult) Get(ctx workflow.Context) (*greengrass.ResetDeploymentsOutput, error) {
    var output greengrass.ResetDeploymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassStartBulkDeploymentResult struct {
	Result workflow.Future
}

func (r *GreengrassStartBulkDeploymentResult) Get(ctx workflow.Context) (*greengrass.StartBulkDeploymentOutput, error) {
    var output greengrass.StartBulkDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassStopBulkDeploymentResult struct {
	Result workflow.Future
}

func (r *GreengrassStopBulkDeploymentResult) Get(ctx workflow.Context) (*greengrass.StopBulkDeploymentOutput, error) {
    var output greengrass.StopBulkDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassTagResourceResult struct {
	Result workflow.Future
}

func (r *GreengrassTagResourceResult) Get(ctx workflow.Context) (*greengrass.TagResourceOutput, error) {
    var output greengrass.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUntagResourceResult struct {
	Result workflow.Future
}

func (r *GreengrassUntagResourceResult) Get(ctx workflow.Context) (*greengrass.UntagResourceOutput, error) {
    var output greengrass.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateConnectivityInfoResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateConnectivityInfoResult) Get(ctx workflow.Context) (*greengrass.UpdateConnectivityInfoOutput, error) {
    var output greengrass.UpdateConnectivityInfoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateConnectorDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateConnectorDefinitionResult) Get(ctx workflow.Context) (*greengrass.UpdateConnectorDefinitionOutput, error) {
    var output greengrass.UpdateConnectorDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateCoreDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateCoreDefinitionResult) Get(ctx workflow.Context) (*greengrass.UpdateCoreDefinitionOutput, error) {
    var output greengrass.UpdateCoreDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateDeviceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateDeviceDefinitionResult) Get(ctx workflow.Context) (*greengrass.UpdateDeviceDefinitionOutput, error) {
    var output greengrass.UpdateDeviceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateFunctionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateFunctionDefinitionResult) Get(ctx workflow.Context) (*greengrass.UpdateFunctionDefinitionOutput, error) {
    var output greengrass.UpdateFunctionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateGroupResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateGroupResult) Get(ctx workflow.Context) (*greengrass.UpdateGroupOutput, error) {
    var output greengrass.UpdateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateGroupCertificateConfigurationResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateGroupCertificateConfigurationResult) Get(ctx workflow.Context) (*greengrass.UpdateGroupCertificateConfigurationOutput, error) {
    var output greengrass.UpdateGroupCertificateConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateLoggerDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateLoggerDefinitionResult) Get(ctx workflow.Context) (*greengrass.UpdateLoggerDefinitionOutput, error) {
    var output greengrass.UpdateLoggerDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateResourceDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateResourceDefinitionResult) Get(ctx workflow.Context) (*greengrass.UpdateResourceDefinitionOutput, error) {
    var output greengrass.UpdateResourceDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GreengrassUpdateSubscriptionDefinitionResult struct {
	Result workflow.Future
}

func (r *GreengrassUpdateSubscriptionDefinitionResult) Get(ctx workflow.Context) (*greengrass.UpdateSubscriptionDefinitionOutput, error) {
    var output greengrass.UpdateSubscriptionDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type GreengrassStub struct {
    activities GreengrassClient
}

func NewGreengrassStub() GreengrassClient {
    return &GreengrassStub{}
}

func (a *GreengrassStub) AssociateRoleToGroup(ctx workflow.Context, input *greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error) {
    var output greengrass.AssociateRoleToGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateRoleToGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) AssociateServiceRoleToAccount(ctx workflow.Context, input *greengrass.AssociateServiceRoleToAccountInput) (*greengrass.AssociateServiceRoleToAccountOutput, error) {
    var output greengrass.AssociateServiceRoleToAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateServiceRoleToAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateConnectorDefinition(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionInput) (*greengrass.CreateConnectorDefinitionOutput, error) {
    var output greengrass.CreateConnectorDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConnectorDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateConnectorDefinitionVersion(ctx workflow.Context, input *greengrass.CreateConnectorDefinitionVersionInput) (*greengrass.CreateConnectorDefinitionVersionOutput, error) {
    var output greengrass.CreateConnectorDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConnectorDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateCoreDefinition(ctx workflow.Context, input *greengrass.CreateCoreDefinitionInput) (*greengrass.CreateCoreDefinitionOutput, error) {
    var output greengrass.CreateCoreDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCoreDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateCoreDefinitionVersion(ctx workflow.Context, input *greengrass.CreateCoreDefinitionVersionInput) (*greengrass.CreateCoreDefinitionVersionOutput, error) {
    var output greengrass.CreateCoreDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCoreDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateDeployment(ctx workflow.Context, input *greengrass.CreateDeploymentInput) (*greengrass.CreateDeploymentOutput, error) {
    var output greengrass.CreateDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateDeviceDefinition(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionInput) (*greengrass.CreateDeviceDefinitionOutput, error) {
    var output greengrass.CreateDeviceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeviceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateDeviceDefinitionVersion(ctx workflow.Context, input *greengrass.CreateDeviceDefinitionVersionInput) (*greengrass.CreateDeviceDefinitionVersionOutput, error) {
    var output greengrass.CreateDeviceDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeviceDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateFunctionDefinition(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionInput) (*greengrass.CreateFunctionDefinitionOutput, error) {
    var output greengrass.CreateFunctionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFunctionDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateFunctionDefinitionVersion(ctx workflow.Context, input *greengrass.CreateFunctionDefinitionVersionInput) (*greengrass.CreateFunctionDefinitionVersionOutput, error) {
    var output greengrass.CreateFunctionDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFunctionDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateGroup(ctx workflow.Context, input *greengrass.CreateGroupInput) (*greengrass.CreateGroupOutput, error) {
    var output greengrass.CreateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateGroupCertificateAuthority(ctx workflow.Context, input *greengrass.CreateGroupCertificateAuthorityInput) (*greengrass.CreateGroupCertificateAuthorityOutput, error) {
    var output greengrass.CreateGroupCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroupCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateGroupVersion(ctx workflow.Context, input *greengrass.CreateGroupVersionInput) (*greengrass.CreateGroupVersionOutput, error) {
    var output greengrass.CreateGroupVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroupVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateLoggerDefinition(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionInput) (*greengrass.CreateLoggerDefinitionOutput, error) {
    var output greengrass.CreateLoggerDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLoggerDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateLoggerDefinitionVersion(ctx workflow.Context, input *greengrass.CreateLoggerDefinitionVersionInput) (*greengrass.CreateLoggerDefinitionVersionOutput, error) {
    var output greengrass.CreateLoggerDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLoggerDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateResourceDefinition(ctx workflow.Context, input *greengrass.CreateResourceDefinitionInput) (*greengrass.CreateResourceDefinitionOutput, error) {
    var output greengrass.CreateResourceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResourceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateResourceDefinitionVersion(ctx workflow.Context, input *greengrass.CreateResourceDefinitionVersionInput) (*greengrass.CreateResourceDefinitionVersionOutput, error) {
    var output greengrass.CreateResourceDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResourceDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateSoftwareUpdateJob(ctx workflow.Context, input *greengrass.CreateSoftwareUpdateJobInput) (*greengrass.CreateSoftwareUpdateJobOutput, error) {
    var output greengrass.CreateSoftwareUpdateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSoftwareUpdateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateSubscriptionDefinition(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionInput) (*greengrass.CreateSubscriptionDefinitionOutput, error) {
    var output greengrass.CreateSubscriptionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSubscriptionDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) CreateSubscriptionDefinitionVersion(ctx workflow.Context, input *greengrass.CreateSubscriptionDefinitionVersionInput) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error) {
    var output greengrass.CreateSubscriptionDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSubscriptionDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteConnectorDefinition(ctx workflow.Context, input *greengrass.DeleteConnectorDefinitionInput) (*greengrass.DeleteConnectorDefinitionOutput, error) {
    var output greengrass.DeleteConnectorDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConnectorDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteCoreDefinition(ctx workflow.Context, input *greengrass.DeleteCoreDefinitionInput) (*greengrass.DeleteCoreDefinitionOutput, error) {
    var output greengrass.DeleteCoreDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCoreDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteDeviceDefinition(ctx workflow.Context, input *greengrass.DeleteDeviceDefinitionInput) (*greengrass.DeleteDeviceDefinitionOutput, error) {
    var output greengrass.DeleteDeviceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeviceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteFunctionDefinition(ctx workflow.Context, input *greengrass.DeleteFunctionDefinitionInput) (*greengrass.DeleteFunctionDefinitionOutput, error) {
    var output greengrass.DeleteFunctionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFunctionDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteGroup(ctx workflow.Context, input *greengrass.DeleteGroupInput) (*greengrass.DeleteGroupOutput, error) {
    var output greengrass.DeleteGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteLoggerDefinition(ctx workflow.Context, input *greengrass.DeleteLoggerDefinitionInput) (*greengrass.DeleteLoggerDefinitionOutput, error) {
    var output greengrass.DeleteLoggerDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLoggerDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteResourceDefinition(ctx workflow.Context, input *greengrass.DeleteResourceDefinitionInput) (*greengrass.DeleteResourceDefinitionOutput, error) {
    var output greengrass.DeleteResourceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DeleteSubscriptionDefinition(ctx workflow.Context, input *greengrass.DeleteSubscriptionDefinitionInput) (*greengrass.DeleteSubscriptionDefinitionOutput, error) {
    var output greengrass.DeleteSubscriptionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSubscriptionDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DisassociateRoleFromGroup(ctx workflow.Context, input *greengrass.DisassociateRoleFromGroupInput) (*greengrass.DisassociateRoleFromGroupOutput, error) {
    var output greengrass.DisassociateRoleFromGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateRoleFromGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) DisassociateServiceRoleFromAccount(ctx workflow.Context, input *greengrass.DisassociateServiceRoleFromAccountInput) (*greengrass.DisassociateServiceRoleFromAccountOutput, error) {
    var output greengrass.DisassociateServiceRoleFromAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateServiceRoleFromAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetAssociatedRole(ctx workflow.Context, input *greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error) {
    var output greengrass.GetAssociatedRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAssociatedRole, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetBulkDeploymentStatus(ctx workflow.Context, input *greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error) {
    var output greengrass.GetBulkDeploymentStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBulkDeploymentStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetConnectivityInfo(ctx workflow.Context, input *greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error) {
    var output greengrass.GetConnectivityInfoOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnectivityInfo, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetConnectorDefinition(ctx workflow.Context, input *greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error) {
    var output greengrass.GetConnectorDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnectorDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetConnectorDefinitionVersion(ctx workflow.Context, input *greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error) {
    var output greengrass.GetConnectorDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnectorDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetCoreDefinition(ctx workflow.Context, input *greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error) {
    var output greengrass.GetCoreDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCoreDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetCoreDefinitionVersion(ctx workflow.Context, input *greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error) {
    var output greengrass.GetCoreDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCoreDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetDeploymentStatus(ctx workflow.Context, input *greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error) {
    var output greengrass.GetDeploymentStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetDeviceDefinition(ctx workflow.Context, input *greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error) {
    var output greengrass.GetDeviceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeviceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetDeviceDefinitionVersion(ctx workflow.Context, input *greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error) {
    var output greengrass.GetDeviceDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeviceDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetFunctionDefinition(ctx workflow.Context, input *greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error) {
    var output greengrass.GetFunctionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFunctionDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetFunctionDefinitionVersion(ctx workflow.Context, input *greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error) {
    var output greengrass.GetFunctionDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFunctionDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetGroup(ctx workflow.Context, input *greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error) {
    var output greengrass.GetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetGroupCertificateAuthority(ctx workflow.Context, input *greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error) {
    var output greengrass.GetGroupCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroupCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetGroupCertificateConfiguration(ctx workflow.Context, input *greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error) {
    var output greengrass.GetGroupCertificateConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroupCertificateConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetGroupVersion(ctx workflow.Context, input *greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error) {
    var output greengrass.GetGroupVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroupVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetLoggerDefinition(ctx workflow.Context, input *greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error) {
    var output greengrass.GetLoggerDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLoggerDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetLoggerDefinitionVersion(ctx workflow.Context, input *greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error) {
    var output greengrass.GetLoggerDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLoggerDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetResourceDefinition(ctx workflow.Context, input *greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error) {
    var output greengrass.GetResourceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetResourceDefinitionVersion(ctx workflow.Context, input *greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error) {
    var output greengrass.GetResourceDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourceDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetServiceRoleForAccount(ctx workflow.Context, input *greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error) {
    var output greengrass.GetServiceRoleForAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceRoleForAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetSubscriptionDefinition(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error) {
    var output greengrass.GetSubscriptionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSubscriptionDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) GetSubscriptionDefinitionVersion(ctx workflow.Context, input *greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error) {
    var output greengrass.GetSubscriptionDefinitionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSubscriptionDefinitionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListBulkDeploymentDetailedReports(ctx workflow.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error) {
    var output greengrass.ListBulkDeploymentDetailedReportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBulkDeploymentDetailedReports, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListBulkDeployments(ctx workflow.Context, input *greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error) {
    var output greengrass.ListBulkDeploymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBulkDeployments, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListConnectorDefinitionVersions(ctx workflow.Context, input *greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error) {
    var output greengrass.ListConnectorDefinitionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConnectorDefinitionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListConnectorDefinitions(ctx workflow.Context, input *greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error) {
    var output greengrass.ListConnectorDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConnectorDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListCoreDefinitionVersions(ctx workflow.Context, input *greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error) {
    var output greengrass.ListCoreDefinitionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCoreDefinitionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListCoreDefinitions(ctx workflow.Context, input *greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error) {
    var output greengrass.ListCoreDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCoreDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListDeployments(ctx workflow.Context, input *greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error) {
    var output greengrass.ListDeploymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeployments, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListDeviceDefinitionVersions(ctx workflow.Context, input *greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error) {
    var output greengrass.ListDeviceDefinitionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeviceDefinitionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListDeviceDefinitions(ctx workflow.Context, input *greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error) {
    var output greengrass.ListDeviceDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeviceDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListFunctionDefinitionVersions(ctx workflow.Context, input *greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error) {
    var output greengrass.ListFunctionDefinitionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFunctionDefinitionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListFunctionDefinitions(ctx workflow.Context, input *greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error) {
    var output greengrass.ListFunctionDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFunctionDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListGroupCertificateAuthorities(ctx workflow.Context, input *greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error) {
    var output greengrass.ListGroupCertificateAuthoritiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroupCertificateAuthorities, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListGroupVersions(ctx workflow.Context, input *greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error) {
    var output greengrass.ListGroupVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroupVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListGroups(ctx workflow.Context, input *greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error) {
    var output greengrass.ListGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListLoggerDefinitionVersions(ctx workflow.Context, input *greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error) {
    var output greengrass.ListLoggerDefinitionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLoggerDefinitionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListLoggerDefinitions(ctx workflow.Context, input *greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error) {
    var output greengrass.ListLoggerDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLoggerDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListResourceDefinitionVersions(ctx workflow.Context, input *greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error) {
    var output greengrass.ListResourceDefinitionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceDefinitionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListResourceDefinitions(ctx workflow.Context, input *greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error) {
    var output greengrass.ListResourceDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListSubscriptionDefinitionVersions(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error) {
    var output greengrass.ListSubscriptionDefinitionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSubscriptionDefinitionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListSubscriptionDefinitions(ctx workflow.Context, input *greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error) {
    var output greengrass.ListSubscriptionDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSubscriptionDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ListTagsForResource(ctx workflow.Context, input *greengrass.ListTagsForResourceInput) (*greengrass.ListTagsForResourceOutput, error) {
    var output greengrass.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) ResetDeployments(ctx workflow.Context, input *greengrass.ResetDeploymentsInput) (*greengrass.ResetDeploymentsOutput, error) {
    var output greengrass.ResetDeploymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetDeployments, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) StartBulkDeployment(ctx workflow.Context, input *greengrass.StartBulkDeploymentInput) (*greengrass.StartBulkDeploymentOutput, error) {
    var output greengrass.StartBulkDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartBulkDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) StopBulkDeployment(ctx workflow.Context, input *greengrass.StopBulkDeploymentInput) (*greengrass.StopBulkDeploymentOutput, error) {
    var output greengrass.StopBulkDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopBulkDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) TagResource(ctx workflow.Context, input *greengrass.TagResourceInput) (*greengrass.TagResourceOutput, error) {
    var output greengrass.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UntagResource(ctx workflow.Context, input *greengrass.UntagResourceInput) (*greengrass.UntagResourceOutput, error) {
    var output greengrass.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateConnectivityInfo(ctx workflow.Context, input *greengrass.UpdateConnectivityInfoInput) (*greengrass.UpdateConnectivityInfoOutput, error) {
    var output greengrass.UpdateConnectivityInfoOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConnectivityInfo, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateConnectorDefinition(ctx workflow.Context, input *greengrass.UpdateConnectorDefinitionInput) (*greengrass.UpdateConnectorDefinitionOutput, error) {
    var output greengrass.UpdateConnectorDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConnectorDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateCoreDefinition(ctx workflow.Context, input *greengrass.UpdateCoreDefinitionInput) (*greengrass.UpdateCoreDefinitionOutput, error) {
    var output greengrass.UpdateCoreDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCoreDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateDeviceDefinition(ctx workflow.Context, input *greengrass.UpdateDeviceDefinitionInput) (*greengrass.UpdateDeviceDefinitionOutput, error) {
    var output greengrass.UpdateDeviceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeviceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateFunctionDefinition(ctx workflow.Context, input *greengrass.UpdateFunctionDefinitionInput) (*greengrass.UpdateFunctionDefinitionOutput, error) {
    var output greengrass.UpdateFunctionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFunctionDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateGroup(ctx workflow.Context, input *greengrass.UpdateGroupInput) (*greengrass.UpdateGroupOutput, error) {
    var output greengrass.UpdateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateGroupCertificateConfiguration(ctx workflow.Context, input *greengrass.UpdateGroupCertificateConfigurationInput) (*greengrass.UpdateGroupCertificateConfigurationOutput, error) {
    var output greengrass.UpdateGroupCertificateConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroupCertificateConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateLoggerDefinition(ctx workflow.Context, input *greengrass.UpdateLoggerDefinitionInput) (*greengrass.UpdateLoggerDefinitionOutput, error) {
    var output greengrass.UpdateLoggerDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLoggerDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateResourceDefinition(ctx workflow.Context, input *greengrass.UpdateResourceDefinitionInput) (*greengrass.UpdateResourceDefinitionOutput, error) {
    var output greengrass.UpdateResourceDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateResourceDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *GreengrassStub) UpdateSubscriptionDefinition(ctx workflow.Context, input *greengrass.UpdateSubscriptionDefinitionInput) (*greengrass.UpdateSubscriptionDefinitionOutput, error) {
    var output greengrass.UpdateSubscriptionDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSubscriptionDefinition, input).Get(ctx, &output)
    return &output, err
}
