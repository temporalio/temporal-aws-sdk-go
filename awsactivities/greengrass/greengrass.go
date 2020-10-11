// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package greengrass

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/greengrass/greengrassiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client greengrassiface.GreengrassAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := greengrass.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (greengrassiface.GreengrassAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return greengrass.New(sess), nil
}

func (a *Activities) AssociateRoleToGroup(ctx context.Context, input *greengrass.AssociateRoleToGroupInput) (*greengrass.AssociateRoleToGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateRoleToGroupWithContext(ctx, input)
}

func (a *Activities) AssociateServiceRoleToAccount(ctx context.Context, input *greengrass.AssociateServiceRoleToAccountInput) (*greengrass.AssociateServiceRoleToAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateServiceRoleToAccountWithContext(ctx, input)
}

func (a *Activities) CreateConnectorDefinition(ctx context.Context, input *greengrass.CreateConnectorDefinitionInput) (*greengrass.CreateConnectorDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConnectorDefinitionWithContext(ctx, input)
}

func (a *Activities) CreateConnectorDefinitionVersion(ctx context.Context, input *greengrass.CreateConnectorDefinitionVersionInput) (*greengrass.CreateConnectorDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConnectorDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) CreateCoreDefinition(ctx context.Context, input *greengrass.CreateCoreDefinitionInput) (*greengrass.CreateCoreDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCoreDefinitionWithContext(ctx, input)
}

func (a *Activities) CreateCoreDefinitionVersion(ctx context.Context, input *greengrass.CreateCoreDefinitionVersionInput) (*greengrass.CreateCoreDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCoreDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) CreateDeployment(ctx context.Context, input *greengrass.CreateDeploymentInput) (*greengrass.CreateDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDeploymentWithContext(ctx, input)
}

func (a *Activities) CreateDeviceDefinition(ctx context.Context, input *greengrass.CreateDeviceDefinitionInput) (*greengrass.CreateDeviceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDeviceDefinitionWithContext(ctx, input)
}

func (a *Activities) CreateDeviceDefinitionVersion(ctx context.Context, input *greengrass.CreateDeviceDefinitionVersionInput) (*greengrass.CreateDeviceDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDeviceDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) CreateFunctionDefinition(ctx context.Context, input *greengrass.CreateFunctionDefinitionInput) (*greengrass.CreateFunctionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateFunctionDefinitionWithContext(ctx, input)
}

func (a *Activities) CreateFunctionDefinitionVersion(ctx context.Context, input *greengrass.CreateFunctionDefinitionVersionInput) (*greengrass.CreateFunctionDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateFunctionDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) CreateGroup(ctx context.Context, input *greengrass.CreateGroupInput) (*greengrass.CreateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGroupWithContext(ctx, input)
}

func (a *Activities) CreateGroupCertificateAuthority(ctx context.Context, input *greengrass.CreateGroupCertificateAuthorityInput) (*greengrass.CreateGroupCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGroupCertificateAuthorityWithContext(ctx, input)
}

func (a *Activities) CreateGroupVersion(ctx context.Context, input *greengrass.CreateGroupVersionInput) (*greengrass.CreateGroupVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGroupVersionWithContext(ctx, input)
}

func (a *Activities) CreateLoggerDefinition(ctx context.Context, input *greengrass.CreateLoggerDefinitionInput) (*greengrass.CreateLoggerDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLoggerDefinitionWithContext(ctx, input)
}

func (a *Activities) CreateLoggerDefinitionVersion(ctx context.Context, input *greengrass.CreateLoggerDefinitionVersionInput) (*greengrass.CreateLoggerDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLoggerDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) CreateResourceDefinition(ctx context.Context, input *greengrass.CreateResourceDefinitionInput) (*greengrass.CreateResourceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResourceDefinitionWithContext(ctx, input)
}

func (a *Activities) CreateResourceDefinitionVersion(ctx context.Context, input *greengrass.CreateResourceDefinitionVersionInput) (*greengrass.CreateResourceDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResourceDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) CreateSoftwareUpdateJob(ctx context.Context, input *greengrass.CreateSoftwareUpdateJobInput) (*greengrass.CreateSoftwareUpdateJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSoftwareUpdateJobWithContext(ctx, input)
}

func (a *Activities) CreateSubscriptionDefinition(ctx context.Context, input *greengrass.CreateSubscriptionDefinitionInput) (*greengrass.CreateSubscriptionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSubscriptionDefinitionWithContext(ctx, input)
}

func (a *Activities) CreateSubscriptionDefinitionVersion(ctx context.Context, input *greengrass.CreateSubscriptionDefinitionVersionInput) (*greengrass.CreateSubscriptionDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSubscriptionDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) DeleteConnectorDefinition(ctx context.Context, input *greengrass.DeleteConnectorDefinitionInput) (*greengrass.DeleteConnectorDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConnectorDefinitionWithContext(ctx, input)
}

func (a *Activities) DeleteCoreDefinition(ctx context.Context, input *greengrass.DeleteCoreDefinitionInput) (*greengrass.DeleteCoreDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCoreDefinitionWithContext(ctx, input)
}

func (a *Activities) DeleteDeviceDefinition(ctx context.Context, input *greengrass.DeleteDeviceDefinitionInput) (*greengrass.DeleteDeviceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDeviceDefinitionWithContext(ctx, input)
}

func (a *Activities) DeleteFunctionDefinition(ctx context.Context, input *greengrass.DeleteFunctionDefinitionInput) (*greengrass.DeleteFunctionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFunctionDefinitionWithContext(ctx, input)
}

func (a *Activities) DeleteGroup(ctx context.Context, input *greengrass.DeleteGroupInput) (*greengrass.DeleteGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGroupWithContext(ctx, input)
}

func (a *Activities) DeleteLoggerDefinition(ctx context.Context, input *greengrass.DeleteLoggerDefinitionInput) (*greengrass.DeleteLoggerDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLoggerDefinitionWithContext(ctx, input)
}

func (a *Activities) DeleteResourceDefinition(ctx context.Context, input *greengrass.DeleteResourceDefinitionInput) (*greengrass.DeleteResourceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResourceDefinitionWithContext(ctx, input)
}

func (a *Activities) DeleteSubscriptionDefinition(ctx context.Context, input *greengrass.DeleteSubscriptionDefinitionInput) (*greengrass.DeleteSubscriptionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSubscriptionDefinitionWithContext(ctx, input)
}

func (a *Activities) DisassociateRoleFromGroup(ctx context.Context, input *greengrass.DisassociateRoleFromGroupInput) (*greengrass.DisassociateRoleFromGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateRoleFromGroupWithContext(ctx, input)
}

func (a *Activities) DisassociateServiceRoleFromAccount(ctx context.Context, input *greengrass.DisassociateServiceRoleFromAccountInput) (*greengrass.DisassociateServiceRoleFromAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateServiceRoleFromAccountWithContext(ctx, input)
}

func (a *Activities) GetAssociatedRole(ctx context.Context, input *greengrass.GetAssociatedRoleInput) (*greengrass.GetAssociatedRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAssociatedRoleWithContext(ctx, input)
}

func (a *Activities) GetBulkDeploymentStatus(ctx context.Context, input *greengrass.GetBulkDeploymentStatusInput) (*greengrass.GetBulkDeploymentStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBulkDeploymentStatusWithContext(ctx, input)
}

func (a *Activities) GetConnectivityInfo(ctx context.Context, input *greengrass.GetConnectivityInfoInput) (*greengrass.GetConnectivityInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConnectivityInfoWithContext(ctx, input)
}

func (a *Activities) GetConnectorDefinition(ctx context.Context, input *greengrass.GetConnectorDefinitionInput) (*greengrass.GetConnectorDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConnectorDefinitionWithContext(ctx, input)
}

func (a *Activities) GetConnectorDefinitionVersion(ctx context.Context, input *greengrass.GetConnectorDefinitionVersionInput) (*greengrass.GetConnectorDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConnectorDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) GetCoreDefinition(ctx context.Context, input *greengrass.GetCoreDefinitionInput) (*greengrass.GetCoreDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCoreDefinitionWithContext(ctx, input)
}

func (a *Activities) GetCoreDefinitionVersion(ctx context.Context, input *greengrass.GetCoreDefinitionVersionInput) (*greengrass.GetCoreDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCoreDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) GetDeploymentStatus(ctx context.Context, input *greengrass.GetDeploymentStatusInput) (*greengrass.GetDeploymentStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeploymentStatusWithContext(ctx, input)
}

func (a *Activities) GetDeviceDefinition(ctx context.Context, input *greengrass.GetDeviceDefinitionInput) (*greengrass.GetDeviceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeviceDefinitionWithContext(ctx, input)
}

func (a *Activities) GetDeviceDefinitionVersion(ctx context.Context, input *greengrass.GetDeviceDefinitionVersionInput) (*greengrass.GetDeviceDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeviceDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) GetFunctionDefinition(ctx context.Context, input *greengrass.GetFunctionDefinitionInput) (*greengrass.GetFunctionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFunctionDefinitionWithContext(ctx, input)
}

func (a *Activities) GetFunctionDefinitionVersion(ctx context.Context, input *greengrass.GetFunctionDefinitionVersionInput) (*greengrass.GetFunctionDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFunctionDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) GetGroup(ctx context.Context, input *greengrass.GetGroupInput) (*greengrass.GetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGroupWithContext(ctx, input)
}

func (a *Activities) GetGroupCertificateAuthority(ctx context.Context, input *greengrass.GetGroupCertificateAuthorityInput) (*greengrass.GetGroupCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGroupCertificateAuthorityWithContext(ctx, input)
}

func (a *Activities) GetGroupCertificateConfiguration(ctx context.Context, input *greengrass.GetGroupCertificateConfigurationInput) (*greengrass.GetGroupCertificateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGroupCertificateConfigurationWithContext(ctx, input)
}

func (a *Activities) GetGroupVersion(ctx context.Context, input *greengrass.GetGroupVersionInput) (*greengrass.GetGroupVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGroupVersionWithContext(ctx, input)
}

func (a *Activities) GetLoggerDefinition(ctx context.Context, input *greengrass.GetLoggerDefinitionInput) (*greengrass.GetLoggerDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLoggerDefinitionWithContext(ctx, input)
}

func (a *Activities) GetLoggerDefinitionVersion(ctx context.Context, input *greengrass.GetLoggerDefinitionVersionInput) (*greengrass.GetLoggerDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLoggerDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) GetResourceDefinition(ctx context.Context, input *greengrass.GetResourceDefinitionInput) (*greengrass.GetResourceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResourceDefinitionWithContext(ctx, input)
}

func (a *Activities) GetResourceDefinitionVersion(ctx context.Context, input *greengrass.GetResourceDefinitionVersionInput) (*greengrass.GetResourceDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResourceDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) GetServiceRoleForAccount(ctx context.Context, input *greengrass.GetServiceRoleForAccountInput) (*greengrass.GetServiceRoleForAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetServiceRoleForAccountWithContext(ctx, input)
}

func (a *Activities) GetSubscriptionDefinition(ctx context.Context, input *greengrass.GetSubscriptionDefinitionInput) (*greengrass.GetSubscriptionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSubscriptionDefinitionWithContext(ctx, input)
}

func (a *Activities) GetSubscriptionDefinitionVersion(ctx context.Context, input *greengrass.GetSubscriptionDefinitionVersionInput) (*greengrass.GetSubscriptionDefinitionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSubscriptionDefinitionVersionWithContext(ctx, input)
}

func (a *Activities) GetThingRuntimeConfiguration(ctx context.Context, input *greengrass.GetThingRuntimeConfigurationInput) (*greengrass.GetThingRuntimeConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetThingRuntimeConfigurationWithContext(ctx, input)
}

func (a *Activities) ListBulkDeploymentDetailedReports(ctx context.Context, input *greengrass.ListBulkDeploymentDetailedReportsInput) (*greengrass.ListBulkDeploymentDetailedReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBulkDeploymentDetailedReportsWithContext(ctx, input)
}

func (a *Activities) ListBulkDeployments(ctx context.Context, input *greengrass.ListBulkDeploymentsInput) (*greengrass.ListBulkDeploymentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBulkDeploymentsWithContext(ctx, input)
}

func (a *Activities) ListConnectorDefinitionVersions(ctx context.Context, input *greengrass.ListConnectorDefinitionVersionsInput) (*greengrass.ListConnectorDefinitionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConnectorDefinitionVersionsWithContext(ctx, input)
}

func (a *Activities) ListConnectorDefinitions(ctx context.Context, input *greengrass.ListConnectorDefinitionsInput) (*greengrass.ListConnectorDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConnectorDefinitionsWithContext(ctx, input)
}

func (a *Activities) ListCoreDefinitionVersions(ctx context.Context, input *greengrass.ListCoreDefinitionVersionsInput) (*greengrass.ListCoreDefinitionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCoreDefinitionVersionsWithContext(ctx, input)
}

func (a *Activities) ListCoreDefinitions(ctx context.Context, input *greengrass.ListCoreDefinitionsInput) (*greengrass.ListCoreDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCoreDefinitionsWithContext(ctx, input)
}

func (a *Activities) ListDeployments(ctx context.Context, input *greengrass.ListDeploymentsInput) (*greengrass.ListDeploymentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeploymentsWithContext(ctx, input)
}

func (a *Activities) ListDeviceDefinitionVersions(ctx context.Context, input *greengrass.ListDeviceDefinitionVersionsInput) (*greengrass.ListDeviceDefinitionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeviceDefinitionVersionsWithContext(ctx, input)
}

func (a *Activities) ListDeviceDefinitions(ctx context.Context, input *greengrass.ListDeviceDefinitionsInput) (*greengrass.ListDeviceDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeviceDefinitionsWithContext(ctx, input)
}

func (a *Activities) ListFunctionDefinitionVersions(ctx context.Context, input *greengrass.ListFunctionDefinitionVersionsInput) (*greengrass.ListFunctionDefinitionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFunctionDefinitionVersionsWithContext(ctx, input)
}

func (a *Activities) ListFunctionDefinitions(ctx context.Context, input *greengrass.ListFunctionDefinitionsInput) (*greengrass.ListFunctionDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFunctionDefinitionsWithContext(ctx, input)
}

func (a *Activities) ListGroupCertificateAuthorities(ctx context.Context, input *greengrass.ListGroupCertificateAuthoritiesInput) (*greengrass.ListGroupCertificateAuthoritiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGroupCertificateAuthoritiesWithContext(ctx, input)
}

func (a *Activities) ListGroupVersions(ctx context.Context, input *greengrass.ListGroupVersionsInput) (*greengrass.ListGroupVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGroupVersionsWithContext(ctx, input)
}

func (a *Activities) ListGroups(ctx context.Context, input *greengrass.ListGroupsInput) (*greengrass.ListGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGroupsWithContext(ctx, input)
}

func (a *Activities) ListLoggerDefinitionVersions(ctx context.Context, input *greengrass.ListLoggerDefinitionVersionsInput) (*greengrass.ListLoggerDefinitionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLoggerDefinitionVersionsWithContext(ctx, input)
}

func (a *Activities) ListLoggerDefinitions(ctx context.Context, input *greengrass.ListLoggerDefinitionsInput) (*greengrass.ListLoggerDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLoggerDefinitionsWithContext(ctx, input)
}

func (a *Activities) ListResourceDefinitionVersions(ctx context.Context, input *greengrass.ListResourceDefinitionVersionsInput) (*greengrass.ListResourceDefinitionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourceDefinitionVersionsWithContext(ctx, input)
}

func (a *Activities) ListResourceDefinitions(ctx context.Context, input *greengrass.ListResourceDefinitionsInput) (*greengrass.ListResourceDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourceDefinitionsWithContext(ctx, input)
}

func (a *Activities) ListSubscriptionDefinitionVersions(ctx context.Context, input *greengrass.ListSubscriptionDefinitionVersionsInput) (*greengrass.ListSubscriptionDefinitionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSubscriptionDefinitionVersionsWithContext(ctx, input)
}

func (a *Activities) ListSubscriptionDefinitions(ctx context.Context, input *greengrass.ListSubscriptionDefinitionsInput) (*greengrass.ListSubscriptionDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSubscriptionDefinitionsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *greengrass.ListTagsForResourceInput) (*greengrass.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ResetDeployments(ctx context.Context, input *greengrass.ResetDeploymentsInput) (*greengrass.ResetDeploymentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetDeploymentsWithContext(ctx, input)
}

func (a *Activities) StartBulkDeployment(ctx context.Context, input *greengrass.StartBulkDeploymentInput) (*greengrass.StartBulkDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartBulkDeploymentWithContext(ctx, input)
}

func (a *Activities) StopBulkDeployment(ctx context.Context, input *greengrass.StopBulkDeploymentInput) (*greengrass.StopBulkDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopBulkDeploymentWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *greengrass.TagResourceInput) (*greengrass.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *greengrass.UntagResourceInput) (*greengrass.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateConnectivityInfo(ctx context.Context, input *greengrass.UpdateConnectivityInfoInput) (*greengrass.UpdateConnectivityInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConnectivityInfoWithContext(ctx, input)
}

func (a *Activities) UpdateConnectorDefinition(ctx context.Context, input *greengrass.UpdateConnectorDefinitionInput) (*greengrass.UpdateConnectorDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConnectorDefinitionWithContext(ctx, input)
}

func (a *Activities) UpdateCoreDefinition(ctx context.Context, input *greengrass.UpdateCoreDefinitionInput) (*greengrass.UpdateCoreDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateCoreDefinitionWithContext(ctx, input)
}

func (a *Activities) UpdateDeviceDefinition(ctx context.Context, input *greengrass.UpdateDeviceDefinitionInput) (*greengrass.UpdateDeviceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDeviceDefinitionWithContext(ctx, input)
}

func (a *Activities) UpdateFunctionDefinition(ctx context.Context, input *greengrass.UpdateFunctionDefinitionInput) (*greengrass.UpdateFunctionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFunctionDefinitionWithContext(ctx, input)
}

func (a *Activities) UpdateGroup(ctx context.Context, input *greengrass.UpdateGroupInput) (*greengrass.UpdateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGroupWithContext(ctx, input)
}

func (a *Activities) UpdateGroupCertificateConfiguration(ctx context.Context, input *greengrass.UpdateGroupCertificateConfigurationInput) (*greengrass.UpdateGroupCertificateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGroupCertificateConfigurationWithContext(ctx, input)
}

func (a *Activities) UpdateLoggerDefinition(ctx context.Context, input *greengrass.UpdateLoggerDefinitionInput) (*greengrass.UpdateLoggerDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateLoggerDefinitionWithContext(ctx, input)
}

func (a *Activities) UpdateResourceDefinition(ctx context.Context, input *greengrass.UpdateResourceDefinitionInput) (*greengrass.UpdateResourceDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateResourceDefinitionWithContext(ctx, input)
}

func (a *Activities) UpdateSubscriptionDefinition(ctx context.Context, input *greengrass.UpdateSubscriptionDefinitionInput) (*greengrass.UpdateSubscriptionDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSubscriptionDefinitionWithContext(ctx, input)
}

func (a *Activities) UpdateThingRuntimeConfiguration(ctx context.Context, input *greengrass.UpdateThingRuntimeConfigurationInput) (*greengrass.UpdateThingRuntimeConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateThingRuntimeConfigurationWithContext(ctx, input)
}